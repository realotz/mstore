<template>
  <div class="p-2 h-full">
    <div class="h-full bg-white">
      <div class="px-2 flex py-1.5 list-harder">
        <Space style="width: 90%">
          <Button value="small">
            <template #icon> <LeftOutlined /></template>
          </Button>
          <Button
            :disabled="!volumeStore.getAdvancePaths || volumeStore.getAdvancePaths.length == 0"
            value="small"
          >
            <template #icon> <RightOutlined /></template>
          </Button>
          <Button value="small">
            <template #icon> <RedoOutlined /></template>
          </Button>
          <Breadcrumb style="width: 600px" :path="pathState" @select="breadSelect" />
          <InputSearch enter-button />
        </Space>
        <Space style="width: 10%; flex-direction: row-reverse">
          <Dropdown>
            <Button value="small">
              <template #icon> <AppstoreOutlined /> </template>
            </Button>
            <template #overlay>
              <Menu>
                <MenuItem>
                  <span href="javascript:;">列表视图</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">小图标</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">大图标</span>
                </MenuItem>
              </Menu>
            </template>
          </Dropdown>
          <Dropdown>
            <Button value="small">
              <template #icon> <FunnelPlotOutlined /> </template>
            </Button>
            <template #overlay>
              <Menu>
                <MenuItem>
                  <span href="javascript:;">名称</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">大小</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">修改日期</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">文件类型</span>
                </MenuItem>
                <MenuDivider />
                <MenuItem>
                  <span href="javascript:;">从小到大</span>
                </MenuItem>
                <MenuItem>
                  <span href="javascript:;">从大到小</span>
                </MenuItem>
              </Menu>
            </template>
          </Dropdown>
        </Space>
      </div>
      <Table
        v-if="showType == 0"
        :columns="columns"
        :dataSource="data"
        :loading="loading"
        :pagination="{ pageSize: 100 }"
        rowKey="name"
        bordered
      />
      <div class="p-4 list-body" v-if="showType > 0">
        <List :grid="{ gutter: 16, column: 7 }" size="small" class="p-2" :data-source="data">
          <template #renderItem="{ item, index }">
            <ListItem style="text-align: center">
              <div
                :class="selectKey == index + 1 ? 'file-item-select file-item' : 'file-item'"
                @click="selectItem(index + 1)"
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
                  <div class="file-item-box">
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
  </div>
</template>

<script lang="ts" setup>
  import Breadcrumb from './Breadcrumb.vue';
  import { computed, onMounted, ref, watch } from 'vue';
  import { BasicTable, ColumnChangeParam } from '/@/components/Table';
  import { BasicColumn } from '/@/components/Table/src/types/table';
  import {
    EditOutlined,
    EllipsisOutlined,
    RedoOutlined,
    LeftOutlined,
    RightOutlined,
    TableOutlined,
    AppstoreOutlined,
    FunnelPlotOutlined,
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
    Dropdown,
    Menu,
    MenuItem,
    MenuDivider,
    Table,
  } from 'ant-design-vue';
  import { BasicForm, useForm } from '/@/components/Form';
  import { isFunction } from '/@/utils/is';
  import { volumeList } from '/@/api/mstore/volume';
  import { formatUnixToTime } from '/@/utils/dateUtil';
  import { getPathInfo } from '/@/utils/filepath';
  import { sizeFmt } from '/@/utils/fmt';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  const volumeStore = useVolumeStoreWithOut();
  const columns: BasicColumn[] = [
    {
      title: '名称',
      dataIndex: 'name',
      width: 150,
    },
    {
      title: '大小',
      dataIndex: 'size',
    },
    {
      title: '类型',
      dataIndex: 'ext',
      width: 150,
    },
    {
      title: '修改日期',
      width: 150,
      dataIndex: 'updated_at',
    },
  ];
  //每行个数
  const grid = ref(12);
  const selectKey = ref(0);
  const loading = ref(false);
  const ListItem = List.Item;
  //数据
  const data = ref([]);
  const pathState = ref('');
  // 前进按钮的栈格式
  const forwardStake = ref([]);
  // 展示类型
  const showType = ref(0);
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
  const emit = defineEmits(['selectDir', 'delete']);

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
  });

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
    return '/resource/img/misc.png';
  };

  // 双击判断
  const clickTimes = ref(0);

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

  // 选择文件夹
  const selectItem = (key: number) => {
    selectKey.value = key;
    clickTimes.value++;
    if (clickTimes.value === 2) {
      clickTimes.value = 0;
      selectKey.value = 0;
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
        selectKey.value = key;
      }
    }, 250);
  };

  //表单提交
  async function handleSubmit() {
    await fetch();
  }

  // 文件列表
  async function fetch() {
    const { path } = props;
    if (path) {
      const info = getPathInfo(path);
      const res = await volumeList(info[0], {
        path: info[1],
      });
      data.value = res.list;
    }
  }
</script>

<style lang="less">
  .file-item {
    width: 128px;
    height: 158px;
    display: block;
  }
  .file-item:hover {
    background: #e6f5ff;
    // border: 1px solid #a6daff;
  }
  .file-item-box {
    height: 130px;
    width: 128px;
    text-align: center;
    vertical-align: bottom;
    display: table-cell;
  }
  .file-item-select {
    background: #e6f5ff;
    border: 1px solid #a6daff;
  }
  .file-item-box img {
    max-height: 128px;
    max-width: 128px;
    width: auto;
    height: auto;
  }
  .file-item-box span {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    overflow: hidden;
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
