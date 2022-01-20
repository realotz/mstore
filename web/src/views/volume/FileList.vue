<template>
  <div class="p-2 h-full">
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
          <Button value="small">
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
                    <img :src="imageShow(item)" />
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
  </div>
</template>

<script lang="ts" setup>
  import Breadcrumb from './Breadcrumb.vue';
  import { computed, onMounted, ref, watch, reactive, onUnmounted } from 'vue';
  import { BasicTable, useTable } from '/@/components/Table';
  import { useContextMenu } from '/@/hooks/web/useContextMenu';
  import { SelectArea, closeArea } from './components/SelectArea';
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
  import { volumeList, fileRename } from '/@/api/mstore/volume';
  import { formatUnixToTime } from '/@/utils/dateUtil';
  import { getPathInfo } from '/@/utils/filepath';
  import { sizeFmt } from '/@/utils/fmt';
  import { columns } from './FileData';
  import FlieShowType from './components/FileShowType.vue';
  import RenameModel from './components/RenameModel.vue';
  import FileSort from './components/FileSort.vue';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  import { useModal } from '/@/components/Modal';
  import { useMessage } from '/@/hooks/web/useMessage';
  const [registerRename, { openModal: openRenameModal }] = useModal();
  const { createMessage } = useMessage();
  const volumeStore = useVolumeStoreWithOut();
  //每行个数
  const [createContextMenu] = useContextMenu();
  const grid = ref(12);
  const selectKey = ref([]);
  const loading = ref(false);
  const ListItem = List.Item;
  //数据
  const data = ref([]);
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
    (newv) => {
      pathState.value = newv;
      fetch();
    },
  );
  //暴露内部方法
  const emit = defineEmits(['selectDir', 'resetDir']);
  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
    document.body.onselectstart = new Function('return false');
    document.addEventListener('mousedown', handleMouseDown);
    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
  });

  onUnmounted(() => {
    document.removeEventListener('mousedown', handleMouseDown);
    document.removeEventListener('mousemove', handleMouseMove);
    document.removeEventListener('mouseup', handleMouseUp);
  });

  // 视图显示
  const handleShowSelect = (key) => {
    showType.value = key;
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

  const handleMouseDown = (e: Event) => {
    mouseTime = setTimeout(function () {
      mouseDown.value = true;
      console.log('鼠标按下');
      closeArea();
      selectProps.startPoint.x = e.clientX - 2;
      selectProps.startPoint.y = e.clientY - 2;
      SelectArea(selectProps);
      mouseDown.value = true;
      mouseComplete.value = false;
      mouseTime = undefined;
    }, 250);
  };

  let allSelectKey = [];

  const handleMouseUp = (e: Event) => {
    if (mouseTime) {
      clearTimeout(mouseTime);
    }
    console.log('鼠标松开');
    console.log(allSelectKey);
    mouseDown.value = false;
    mouseComplete.value = true;
    selectProps.startPoint.x = 0;
    selectProps.startPoint.y = 0;
    selectProps.endPoint.x = 0;
    selectProps.endPoint.y = 0;
    closeArea();
    selectKey.value = allSelectKey;
    allSelectKey = [];
  };

  const handleMouseMove = (e: Event) => {
    if (mouseDown.value && !mouseComplete.value) {
      selectProps.endPoint.x = e.clientX - 2;
      selectProps.endPoint.y = e.clientY - 2;
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

  // 右键绑定
  const handleBodyContext = (e) => {
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
            label: `上传到此目录`,
            icon: 'bi:cloud-upload-fill',
            handler: () => {},
          },
        ],
      });
    }
  };

  // 有元素的时候右键
  const handleItemContext = (e, key) => {
    const item = data.value[key - 1];
    let items = [
      {
        label: '打开',
        icon: 'bx:bxs-folder-open',
        handler: () => {},
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
        handler: () => {},
      },
      {
        label: '粘贴',
        icon: 'bx:bx-copy',
        handler: () => {},
      },
      {
        label: '删除',
        icon: 'ant-design:delete-filled',
        handler: () => {},
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
    console.log(selectKey.value);
    clickTimes.value++;
    if (clickTimes.value === 2) {
      clickTimes.value = 0;
      selectKey.value = [];
      const item = data.value[key - 1];
      if (item.is_dir) {
        volumeStore.addBackPath(props.path);
        volumeStore.resetAdvancePath();
        if (item.path == '/') {
          emit('selectDir', '/' + item.volume_id + item.path + item.name);
        } else {
          emit('selectDir', '/' + item.volume_id + item.path + '/' + item.name);
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
    emit('resetDir', `/${item.volume_id}${item.path}/${data.name}`);
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
      const info = getPathInfo(path);
      const res = await volumeList(info[0], {
        path: info[1],
        option: params.value,
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
    // border: 1px solid #a6daff;
  }
  .file-item1 {
    height: 80px;
    width: 66px;
    display: block;
  }
  .file-item1:hover {
    background: #e6f5ff;
    // border: 1px solid #a6daff;
  }
  .file-item-box1 {
    height: 80px;
    width: 66px;
    text-align: center;
    vertical-align: bottom;
    display: table-cell;
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
    width: 70px;
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
