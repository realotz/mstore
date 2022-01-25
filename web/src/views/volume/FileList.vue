<template>
  <div class="flex h-full">
    <div class="p-2 h-full w-1/4 xl:w-1/5 m-2 mr-0 overflow-hidden bg-white">
      <BasicTree
        title="目录"
        :clickRowToExpand="false"
        :treeData="treeData"
        :load-data="onLoadData"
        v-model:selectedKeys="selectedKeys"
        v-model:expandedKeys="expandedKeys"
        ref="asyncTreeRef"
        @select="handleSelect"
      />
    </div>
    <div class="p-2 h-full w-3/4 xl:w-4/5">
      <div class="h-full bg-white">
        <div class="px-2 flex py-1.5 list-harder">
          <Space style="width: 90%">
            <Button value="small" @click="handleBack">
              <template #icon> <LeftOutlined /></template>
            </Button>
            <Button
              :disabled="!volumeStore.getAdvancePaths || volumeStore.getAdvancePaths.length == 0"
              value="small"
              @click="handleAdvance"
            >
              <template #icon> <RightOutlined /></template>
            </Button>
            <Button value="small" @click="fetch">
              <template #icon> <RedoOutlined /></template>
            </Button>
            <Breadcrumb style="width: 600px" :path="pathState" @select="breadSelect" />
            <!-- <InputSearch enter-button /> -->
          </Space>
          <Space style="width: 10%; flex-direction: row-reverse">
            <FileSort @select="handleSortSelect" />
            <FlieShowType @select="handleShowSelect" />
          </Space>
        </div>
        <BasicTable
          v-if="showType == 0"
          :columns="columns"
          :dataSource="data"
          :loading="loading"
          :pagination="{ pageSize: 100 }"
          :bordered="true"
          rowKey="name"
        >
          <template #name="{ record }">
            <Space>
              <img class="file-item-list-img" :src="imageShow(record)" />
              <span>{{ record.name }}</span>
            </Space>
          </template>
          <template #size="{ record }">{{ sizeFmt(record.size) }} </template>
          <template #ext="{ record }">{{ record.ext }} </template>
          <template #updated_at="{ record }">{{ formatUnixToTime(record.updated_at) }} </template>
        </BasicTable>
        <div class="p-4 list-body" @contextmenu="handleBodyContext" v-if="showType > 0">
          <List
            :loading="loading"
            :grid="{ gutter: 16, column: showType == 1 ? 12 : 7 }"
            size="small"
            class="p-2"
            :data-source="data"
          >
            <template #renderItem="{ item, index }">
              <ListItem style="text-align: center">
                <div
                  :indexkey="index + 1"
                  :class="
                    selectKey.indexOf(index + 1) > -1
                      ? `file-item-select file-item${showType}`
                      : `file-item${showType}`
                  "
                  @click="selectItem(index + 1)"
                  @mouseenter="suspensionItem(index + 1)"
                  @mouseleave="outSuspensionItem(index + 1)"
                >
                  <Tooltip placement="bottom" :mouseEnterDelay="0.8">
                    <template #title>
                      <div style="font-size: 10px">
                        <span>名称{{ item.name }}</span>
                        <br />
                        <span>大小：{{ sizeFmt(item.size) }}</span>
                        <br />
                        <span>修改日期: {{ formatUnixToTime(item.updated_at) }}</span>
                      </div>
                    </template>
                    <div :class="`file-item-box${showType}`">
                      <img draggable="false" :src="imageShow(item)" />
                      <span>{{ item.name }}</span>
                    </div>
                  </Tooltip>
                </div>
              </ListItem>
            </template>
          </List>
        </div>
      </div>
      <RenameModel
        :autoSubmitOnEnter="true"
        :height="150"
        :minHeight="10"
        @register="registerRename"
        @ok="renameHandle"
      />
      <CopyInfoModel
        :autoSubmitOnEnter="true"
        :height="200"
        :width="280"
        :minHeight="10"
        @register="registerCopyinfo"
        @ok="confirmCopyMoveHandle"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
  import Breadcrumb from './Breadcrumb.vue';
  import { computed, onMounted, ref, unref, watch, reactive, onUnmounted } from 'vue';
  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { BasicTable, useTable } from '/@/components/Table';
  import { useContextMenu } from '/@/hooks/web/useContextMenu';
  import { SelectArea, closeArea } from './components/selelct-area/SelectArea';
  import { ShowMsg, closeMsg } from './components/context-msg/ContextMsg';
  import {
    EditOutlined,
    EllipsisOutlined,
    RedoOutlined,
    LeftOutlined,
    RightOutlined,
  } from '@ant-design/icons-vue';
  import {
    List,
    Card,
    Image,
    Typography,
    Tooltip,
    Slider,
    Avatar,
    RadioGroup,
    InputSearch,
    Button,
    Space,
    Menu,
    Table,
  } from 'ant-design-vue';
  import { volumeList, fileRename, copyMove, delFile } from '/@/api/mstore/volume';
  import { formatUnixToTime } from '/@/utils/dateUtil';
  import { getPathInfo, pathFmt } from '/@/utils/filepath';
  import { sizeFmt } from '/@/utils/fmt';
  import { columns } from './FileData';
  import FlieShowType from './components/FileShowType.vue';
  import RenameModel from './components/RenameModel.vue';
  import CopyInfoModel from './components/CopyInfoModel.vue';
  import FileSort from './components/FileSort.vue';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  import { useModal } from '/@/components/Modal';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { FileListItem } from './model/volumeModel';
  const [registerRename, { openModal: openRenameModal }] = useModal();
  const [registerCopyinfo, { openModal: openCopyModal }] = useModal();
  const { createMessage, createConfirm, createInfoModal } = useMessage();
  const volumeStore = useVolumeStoreWithOut();
  //每行个数
  const [createContextMenu] = useContextMenu();
  const grid = ref(12);
  const selectKey = ref([]);
  const loading = ref(false);
  const ListItem = List.Item;
  //数据
  const data: Array<FileListItem> = ref([]);
  const pathState = ref('');
  // 展示类型
  const showType = ref(1);
  const params = ref({
    order_field: 'name',
    order_desc: false,
  });
  const suspensionKey = ref(0);
  // 组件接收参数
  const props = defineProps({
    path: {
      type: String,
      default: '',
    },
  });
  watch(
    () => props.path,
    (v) => {
      if (!v) {
        return;
      }
      pathState.value = v;
      fetch();
      const paths = v.split('/');
      let p = '';
      for (let i = 1; i < paths.length; i++) {
        p += '/' + paths[i];
        let f = true;
        for (let j = 0; j <= expandedKeys.value.length; j++) {
          if (expandedKeys.value[j] == p) {
            f = false;
          }
        }
        if (f) {
          loadTree(p);
          expandedKeys.value.push(p);
        }
      }
      selectedKeys.value = [v];
    },
  );
  watch(
    () => volumeStore.getVolumes,
    (v: string) => {
      fetchDir();
    },
  );
  //暴露内部方法
  const emit = defineEmits(['selectDir']);
  // 自动请求并暴露内部方法
  onMounted(() => {
    fetchDir();
    fetch();
    document.body.onselectstart = new Function('return false');
    document.addEventListener('mousedown', handleMouseDown);
    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
  });

  onUnmounted(() => {
    document.body.onselectstart = null;
    document.removeEventListener('mousedown', handleMouseDown);
    document.removeEventListener('mousemove', handleMouseMove);
    document.removeEventListener('mouseup', handleMouseUp);
  });

  const treeData = ref<TreeItem[]>([]);
  const asyncTreeRef = ref<Nullable<TreeActionType>>(null);
  const expandedKeys = ref<string[]>([]);
  const selectedKeys = ref<string[]>([]);

  // 存储卷列表
  async function fetchDir() {
    treeData.value = [];
    const volumes = volumeStore.getVolumes;
    volumes.map((item) => {
      treeData.value.push({
        title: item.name,
        key: '/' + item.id,
        isLeaf: false,
        children: [],
      });
    });
  }

  async function loadTree(key) {
    const info = getPathInfo(key);
    const res = await volumeList(info[0], {
      type: 2,
      path: info[1],
    });
    const children = res.list.map((item) => {
      return {
        title: item.name,
        key: pathFmt('/' + info[0] + '/' + item.path + '/' + item.name),
        isLeaf: false,
        children: [],
      };
    });
    const asyncTreeAction: TreeActionType | null = unref(asyncTreeRef);
    if (asyncTreeAction) {
      if (children.length > 0) {
        asyncTreeAction.updateNodeByKey(key, { children: children });
      } else {
        asyncTreeAction.updateNodeByKey(key, { isLeaf: true });
      }
      asyncTreeAction.setExpandedKeys([key, ...asyncTreeAction.getExpandedKeys()]);
    }
  }

  // 异步展开存储卷
  async function onLoadData(treeNode) {
    await loadTree(treeNode.eventKey);
  }

  function handleSelect(keys) {
    if (keys[0]) {
      if (pathState.value != '') {
        volumeStore.addBackPath(pathState.value);
        volumeStore.resetAdvancePath();
      }
      emit('selectDir', keys[0]);
    } else {
      selectedKeys.value = [oldPath.value];
    }
  }

  // 视图显示
  const handleShowSelect = (key) => {
    showType.value = key;
  };

  const handleSelectDir = (key) => {
    emit('selectDir', key);
  };

  // 筛选排序
  const handleSortSelect = (sortName, sortType) => {
    params.value = {
      order_field: sortName,
      order_desc: sortType == 'desc',
    };
    fetch();
  };

  // 选中文件
  const suspensionItem = (key) => {
    suspensionKey.value = key;
  };
  // 取消选中
  const outSuspensionItem = (key) => {
    suspensionKey.value = 0;
  };

  //鼠标是否按下

  const mouseDown = ref(false);
  const mouseComplete = ref(false);

  let selectProps = reactive({
    startPoint: {
      x: 0,
      y: 0,
    },
    endPoint: {
      x: 0,
      y: 0,
    },
  });

  /**
   * 获取该元素下可以被选中的元素集合
   * @param parentElement 父元素
   * @param selectBoxElement 选择框元素
   * @param keyCode 可选元素标识
   * @returns
   */
  function selectElement(parentElement: HTMLElement, selectBoxElement: HTMLElement) {
    const canCheckedElements = parentElement.querySelectorAll(`.file-item${showType.value}`);
    const containElements = judgeContainElement(selectBoxElement, canCheckedElements);
    return {
      containElements,
      canCheckedElements,
    };
  }

  /**
   *
   * 获取该元素下可以被选中的元素集合
   * @param parentElement 父元素
   * @param keyCode 可选元素标识
   * @returns
   */
  function getChildrens(parentElement: HTMLElement, keyCode: string) {
    const ary = [];
    const childs = parentElement.childNodes;
    for (let i = 0; i < childs.length; i++) {
      if (childs[i].nodeType === 1) {
        if ((childs[i] as HTMLElement).getAttribute(keyCode) !== null) {
          ary.push(childs[i]);
        }
      }
    }
    return ary as Array<HTMLElement>;
  }

  function judgeContainElement(
    selectBoxElement: HTMLElement,
    canCheckedElements: Array<HTMLElement>,
  ) {
    const ContainElement: Array<HTMLElement> = [];
    const { left, right, bottom, top } = selectBoxElement.getBoundingClientRect();
    canCheckedElements.forEach((item) => {
      const child = item.getBoundingClientRect();
      if (child.left > left && child.top > top && child.bottom < bottom && child.right < right) {
        ContainElement.push(item);
      }
    });
    return ContainElement;
  }
  let mouseTime;
  let selectMove;

  // 鼠标按下事件
  const handleMouseDown = (e: Event) => {
    if (e.which != 1) {
      return;
    }
    // 选区选中开始 设置个延迟
    mouseTime = setTimeout(function () {
      mouseDown.value = true;
      mouseComplete.value = false;
      mouseTime = undefined;
      // console.log('鼠标按下');
      // 拖拽移动
      if (selectKey.value.length > 0) {
        const selectFiles = document.querySelectorAll('.file-item-select');
        for (let i = 0; i < selectFiles.length; i++) {
          const child = selectFiles[i].getBoundingClientRect();
          if (
            e.clientY >= child.top &&
            e.clientY <= child.top + child.height &&
            e.clientX >= child.left &&
            e.clientY <= child.left + child.width
          ) {
            selectMove = true;
            return;
          }
        }
      }
      closeArea();
      selectProps.startPoint.x = e.clientX - 2;
      selectProps.startPoint.y = e.clientY - 2;
      SelectArea(selectProps);
    }, 250);
  };

  // 选区选中临时存储数据
  let allSelectKey = [];

  let resIndex = 0;

  // 松开鼠标
  const handleMouseUp = (e: Event) => {
    if (e.which != 1) {
      return;
    }
    if (mouseTime) {
      clearTimeout(mouseTime);
    }
    // console.log('松开鼠标');
    mouseDown.value = false;
    mouseComplete.value = true;
    selectProps.startPoint.x = 0;
    selectProps.startPoint.y = 0;
    selectProps.endPoint.x = 0;
    selectProps.endPoint.y = 0;
    closeArea();
    closeMsg();
    // console.log(selectMove);
    if (selectMove) {
      let to_id, to_path, to_name;
      let is_dir = false;
      if (suspensionKey.value) {
        if (selectKey.value.indexOf(suspensionKey.value) == -1) {
          const toItem = data.value[suspensionKey.value - 1] as FileListItem;
          to_id = toItem.volume_id;
          to_path = pathFmt(`${toItem.path}/${toItem.name}`);
          is_dir = toItem.is_dir;
          to_name = toItem.name;
        }
      } else if (e.target.getAttribute('indexkey')) {
        is_dir = true;
        const info = getPathInfo(e.target.getAttribute('indexkey'));
        to_id = info[0];
        to_path = info[1];
        let paths = e.target.getAttribute('indexkey').split('/');
        if (paths.length == 2) {
          for (let i = 0; i < volumeStore.getVolumes.length; i++) {
            if (volumeStore.getVolumes[i].id == paths[1]) {
              to_name = volumeStore.getVolumes[i]?.name;
              break;
            }
          }
        } else {
          to_name = paths[paths.length - 1];
        }
      }
      let indexs = selectKey.value;
      if (is_dir) {
        for (let i = 0; i < indexs.length; i++) {
          const item = data.value[indexs[i] - 1];
          if (
            item?.volume_id == to_id &&
            (item?.path == to_path || pathFmt(`${item.path}/${item.name}`) == to_path)
          ) {
            createMessage.error('错误的目标，无法移动到相同目录');
            return;
          }
        }
        createConfirm({
          iconType: 'warning',
          title: '确认',
          content: `是否将${indexs.length}个文件移动到 ${to_name}`,
          onOk: () => {
            let files = [];
            for (let i = 0; i < indexs.length; i++) {
              const item = data.value[indexs[i] - 1];
              files.push({
                id: item?.volume_id,
                path: pathFmt(`${item.path}/${item.name}`),
              });
            }
            const params = {
              files: files,
              is_delete: true,
              to_path: to_path,
              to_volume_id: to_id,
            };
            // todo move
            copyMove(params, 'none').then((data) => {
              if (data?.code == 200 || data?.code == undefined) {
                loadTree(pathFmt(pathState.value));
                loadTree(pathFmt(`/${to_id}/${to_path}`));
                fetch();
                return;
              }
              if (data?.code == 409) {
                openCopyModal(true, params);
              } else {
                createMessage.error(data?.message);
              }
            });
          },
        });
      }
    }
    selectKey.value = allSelectKey;
    allSelectKey = [];
    selectMove = false;
  };

  let msgPorp = reactive({
    text: '',
    point: {
      x: 0,
      y: 0,
    },
  });

  // 鼠标移动
  const handleMouseMove = (e: Event) => {
    if (e.which != 1) {
      return;
    }
    // 拖动
    if (selectMove) {
      if (suspensionKey.value > 0 && selectKey.value.indexOf(suspensionKey.value) == -1) {
        const item = data.value[suspensionKey.value - 1];
        if (item.is_dir) {
          msgPorp.text = `移动到 ${item?.name}`;
          ShowMsg(msgPorp);
          msgPorp.point.x = e.clientX + 2;
          msgPorp.point.y = e.clientY + 2;
        } else {
          closeMsg();
        }
      } else {
        if (e.target.getAttribute('indexkey')) {
          let path = e.target.getAttribute('indexkey');
          let paths = path.split('/');
          if (paths.length == 2) {
            for (let i = 0; i < volumeStore.getVolumes.length; i++) {
              if (volumeStore.getVolumes[i].id == paths[1]) {
                msgPorp.text = `移动到 ${volumeStore.getVolumes[i]?.name}`;
                break;
              }
            }
          } else {
            msgPorp.text = `移动到 ${paths[paths.length - 1]}`;
          }
          ShowMsg(msgPorp);
          msgPorp.point.x = e.clientX + 2;
          msgPorp.point.y = e.clientY + 2;
        } else {
          closeMsg();
        }
      }
      return;
    }
    // 选区
    if (mouseDown.value && !mouseComplete.value) {
      if (selectProps.startPoint.x > selectProps.endPoint.x) {
        selectProps.endPoint.x = e.clientX + 2;
      } else {
        selectProps.endPoint.x = e.clientX - 1;
      }
      if (selectProps.startPoint.y > selectProps.endPoint.y) {
        selectProps.endPoint.y = e.clientY + 2;
      } else {
        selectProps.endPoint.y = e.clientY - 1;
      }
      const div = document.querySelector('#select-area');
      const parent = document.querySelector('.list-body');
      const containDiv = selectElement(parent as HTMLElement, div as HTMLElement);
      containDiv.canCheckedElements.forEach((item) => {
        item.className = `file-item${showType.value}`;
      });
      allSelectKey = containDiv.containElements.map((item) => {
        item.className = `file-item-select file-item${showType.value}`;
        return parseInt(item.getAttribute('indexkey'));
      });
    }
  };

  const delFileHandle = (items) => {
    if (items.length <= 0) {
      return;
    }
    let content = `确认删除${items.length}个文件？`;
    if (items.length == 1) {
      content = `确认删除【${items[0].name}】?`;
    }
    createConfirm({
      iconType: 'warning',
      title: '确认',
      content: content,
      onOk: () => {
        let files = [];
        for (let i = 0; i < items.length; i++) {
          files.push({
            id: items[i]?.volume_id,
            path: pathFmt(`${items[i].path}/${items[i].name}`),
          });
        }
        delFile({
          files: files,
        }).then((e) => {
          loadTree(pathFmt(pathState.value));
          fetch();
        });
      },
    });
  };

  // 右键绑定
  const handleBodyContext = (e: Event) => {
    mouseDown.value = false;
    // 有选中元素
    if (suspensionKey.value != 0) {
      return handleItemContext(e, suspensionKey.value);
    }
    if (props.path) {
      createContextMenu({
        event: e,
        items: [
          {
            label: '新建文件',
            icon: 'bi:plus',
            handler: () => {
              createMessage.success('click new');
            },
          },
          {
            label: '新建文件夹',
            icon: 'bi:plus',
            handler: () => {
              createMessage.success('click new');
            },
          },
          {
            label: '粘贴',
            icon: 'bx:bx-copy',
            disabled: copyItems.length == 0,
            handler: pasteHandler,
          },
          {
            label: `上传到此目录`,
            icon: 'bi:cloud-upload-fill',
            handler: () => {},
          },
        ],
      });
    }
  };

  // 确认移动/粘贴
  const confirmCopyMoveHandle = (params, wireType) => {
    const res = copyMove({
      wire_type: wireType,
      ...params,
    }).then((data) => {
      loadTree(pathFmt(pathState.value));
      fetch();
    });
  };

  //粘贴操作
  const pasteHandler = () => {
    if (copyItems.length == 0) {
      return;
    }
    const info = getPathInfo(unref(pathState));
    let is_confirm = false;
    const files = copyItems.map((item) => {
      if (item.path != info[1]) {
        is_confirm = true;
      }
      return {
        id: item?.volume_id,
        path: pathFmt(`${item.path}/${item.name}`),
      };
    });
    let params = {
      files: files,
      is_delete: false,
      to_path: info[1],
      to_volume_id: info[0],
    };
    const res = copyMove(params, 'none').then((data) => {
      if (data?.code == 200 || data?.code == undefined) {
        loadTree(pathFmt(pathState.value));
        fetch();
        return;
      }
      if (data?.code == 409) {
        if (is_confirm) {
          //不同目录复制询问是否覆盖还是重命名
          openCopyModal(true, params);
        } else {
          //同目录复制直接重命名
          copyMoveHandle(params, 2);
        }
      } else {
        createMessage.error(data?.message);
      }
    });
  };

  let copyItems = [];

  // 有元素的时候右键
  const handleItemContext = (e: Event, key) => {
    if (selectKey.value.length > 1) {
      const items = selectKey.value.map((id) => {
        return unref(data)[id - 1];
      });
      console.log(items);
      console.log(selectKey.value);
      e?.stopPropagation();
      e?.preventDefault();
      createContextMenu({
        event: e,
        items: [
          {
            label: '下载',
            icon: 'bi:cloud-arrow-down-fill',
            handler: () => {},
          },
          {
            label: '复制',
            icon: 'bx:bx-copy-alt',
            handler: () => {
              copyItems = items;
              createMessage.success('复制成功');
            },
          },
          {
            label: '粘贴',
            icon: 'bx:bx-copy',
            handler: pasteHandler,
          },
          {
            label: '删除',
            icon: 'ant-design:delete-filled',
            handler: () => {
              delFileHandle(items);
            },
          },
        ],
      });
      return;
    }
    const item = data.value[key - 1];
    let items = [
      {
        label: '打开',
        icon: 'bx:bxs-folder-open',
        handler: () => {
          if (item.is_dir) {
            volumeStore.addBackPath(props.path);
            if (item.path == '/') {
              emit('selectDir', pathFmt('/' + item.volume_id + item.path + item.name));
            } else {
              emit('selectDir', pathFmt('/' + item.volume_id + item.path + '/' + item.name));
            }
          }
        },
      },
      {
        label: '下载',
        icon: 'bi:cloud-arrow-down-fill',
        handler: () => {},
      },
      {
        label: '重命名',
        icon: 'ic:baseline-drive-file-rename-outline',
        handler: () => {
          openRenameModal(true, item);
        },
      },
      {
        label: '复制',
        icon: 'bx:bx-copy-alt',
        handler: () => {
          copyItems = [item];
          createMessage.success('复制成功');
        },
      },
      {
        label: '粘贴',
        icon: 'bx:bx-copy',
        handler: pasteHandler,
      },
      {
        label: '删除',
        icon: 'ant-design:delete-filled',
        handler: () => {
          delFileHandle([item]);
        },
      },
    ];
    if (checkVedio(item.ext)) {
      items[0] = {
        label: '播放视频',
        icon: 'ant-design:play-circle-filled',
        handler: () => {},
      };
    }
    if (checkAudio(item.ext)) {
      items[0] = {
        label: '播放音乐',
        icon: 'ant-design:play-circle-outlined',
        handler: () => {},
      };
    }
    createContextMenu({
      event: e,
      items: items,
    });
  };

  // 前进操作
  const handleAdvance = () => {
    const path = volumeStore.getAdvancePath();
    if (path) {
      volumeStore.addBackPath(props.path);
      emit('selectDir', path);
    }
  };

  // 后退操作
  const handleBack = () => {
    const path = volumeStore.getBackPath();
    if (path) {
      volumeStore.addAdvancePath(props.path);
      emit('selectDir', path);
    } else {
      if (props.path) {
        const paths = props.path.split('/');
        if (paths.length > 2) {
          volumeStore.addAdvancePath(props.path);
        }
        paths.pop();
        emit('selectDir', paths.join('/'));
      }
    }
  };

  // 面包屑选择
  const breadSelect = (key) => {
    if (key) {
      volumeStore.addBackPath(props.path);
      volumeStore.resetAdvancePath();
      emit('selectDir', key);
    }
  };

  // 双击判断
  const clickTimes = ref(0);

  // 文件单击选中
  const selectItem = (key: number) => {
    selectKey.value = [key];
    clickTimes.value++;
    if (clickTimes.value === 2) {
      clickTimes.value = 0;
      selectKey.value = [];
      const item = data.value[key - 1];
      if (item.is_dir) {
        volumeStore.addBackPath(props.path);
        volumeStore.resetAdvancePath();
        if (item.path == '/') {
          emit('selectDir', pathFmt('/' + item.volume_id + item.path + item.name));
        } else {
          emit('selectDir', pathFmt('/' + item.volume_id + item.path + '/' + item.name));
        }
      }
    }
    setTimeout(function () {
      if (clickTimes.value === 1) {
        clickTimes.value = 0;
        selectKey.value = [key];
      }
    }, 250);
  };

  // 重命名
  async function renameHandle(item, data) {
    let path = item.path;
    if (path !== '/') {
      path = path + '/';
    }
    const res = await fileRename(item.volume_id, {
      path: path + item.name,
      new_path: path + data.name,
    });
    createMessage.success('文件重命名成功');
    fetch();
    openRenameModal(false, {});
    loadTree(pathFmt(pathState.value));
    loadTree(pathFmt(`/${toItem.volume_id}/${toItem.path}`));
  }

  //表单提交
  async function handleSubmit() {
    await fetch();
  }

  // 检查是否视频
  const checkVedio = (ext) => {
    if (
      ext == '.mp4' ||
      ext == '.mov' ||
      ext == '.qt' ||
      ext == '.mpg' ||
      ext == '.avi' ||
      ext == '.mod' ||
      ext == '.flv' ||
      ext == '.rmvb' ||
      ext == '.mkv' ||
      ext == '.rm' ||
      ext == '.wmv'
    ) {
      return true;
    }
    return false;
  };

  // 检查是否音频
  const checkAudio = (ext) => {
    if (ext == '.wav' || ext == '.mp3' || ext == '.ogg' || ext == '.acc' || ext == '.webm') {
      return true;
    }
    return false;
  };

  // 文件图标显示
  const imageShow = (item: any) => {
    if (item.is_dir) {
      return '/resource/img/folder.png';
    }
    if (item.ext === '.mp3') {
      return '/resource/img/flac.png';
    }
    if (item.ext === '.mp4') {
      return '/resource/img/mp4.png';
    }
    if (item.ext === '.exe') {
      return '/resource/img/exe.png';
    }
    if (item.ext === '.zip') {
      return '/resource/img/zip.png';
    }
    if (item.ext === '.pdf') {
      return '/resource/img/pdf.png';
    }
    if (item.ext === '.html') {
      return '/resource/img/html.png';
    }
    return '/resource/img/misc.png';
  };

  // 文件列表
  async function fetch() {
    const { path } = props;
    if (path) {
      loading.value = true;
      const info = getPathInfo(path);
      const res = await volumeList(info[0], {
        path: info[1],
        option: params.value,
      }).finally(() => {
        loading.value = false;
      });
      data.value = res.list;
    }
  }
</script>

<style lang="less">
  .file-item-select {
    background: #e6f5ff;
    border: 1px solid #a6daff;
  }
  .file-item2 {
    width: 128px;
    height: 158px;
    display: block;
  }
  .file-item2:hover {
    background: #e6f5ff;
  }
  .file-item1 {
    height: 90px;
    width: 70px;
    display: block;
  }
  .file-item1:hover {
    background: #e6f5ff;
  }
  .file-item-box1 {
    height: 90px;
    width: 70px;
    text-align: center;
    vertical-align: bottom;
  }
  .file-item-box1 img {
    max-height: 64px;
    max-width: 64px;
    width: auto;
    height: auto;
  }
  .file-item-box2 {
    height: 130px;
    width: 128px;
    text-align: center;
    vertical-align: bottom;
    display: table-cell;
  }
  .file-item-box2 img {
    max-height: 128px;
    max-width: 128px;
    width: auto;
    height: auto;
  }
  .file-item-box2 span {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    overflow: hidden;
  }
  .file-item-box1 span {
    width: 68px;
    font-size: 10px;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    overflow: hidden;
  }
  .file-item-list-img {
    max-height: 20px;
    max-width: 20px;
    width: auto;
    height: auto;
  }
  .list-harder {
    height: 48px;
    padding-left: 7px;
  }
  .list-body {
    overflow-y: scroll;
    height: calc(100% - 48px);
  }
</style>
