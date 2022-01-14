<template>
  <div class="p-2 h-full">
    {{ sliderProp.width }}
    <div class="h-full">
      <div
        class="bg-white flex px-2 py-1.5 items-center basic-tree-header"
        style="height: 48px; padding-left: 7px"
      >
        <RadioGroup>
          <Tooltip @click="fetch">
            <template #title>刷新</template>
            <Button value="small"><LeftOutlined /></Button>
          </Tooltip>
          <Tooltip @click="fetch">
            <template #title>刷新</template>
            <Button value="small"><RightOutlined /></Button>
          </Tooltip>
          <Tooltip @click="fetch">
            <template #title>刷新</template>
            <Button value="small"><RedoOutlined /></Button>
          </Tooltip>
        </RadioGroup>
        <Breadcrumb :path="pathState" />
      </div>
      <div class="p-4 h-full bg-white">
        <List :grid="{ gutter: 16, column: 7 }" size="small" class="p-2 h-full" :data-source="data">
          <template #renderItem="{ item, index }">
            <ListItem style="text-align: center">
              <div
                :class="selectKey == index + 1 ? 'file-item-select file-item' : 'file-item'"
                @click="selectItem(index + 1)"
              >
                <Tooltip placement="bottom" mouseEnterDelay="0.8">
                  <template #title>
                    <div style="font-size: 10px">
                      <span>名称{{ item.name }}</span>
                      <br />
                      <span>大小：{{ sizeShow(item.size) }}</span>
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
  import {
    EditOutlined,
    EllipsisOutlined,
    RedoOutlined,
    LeftOutlined,
    RightOutlined,
    TableOutlined,
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
  } from 'ant-design-vue';
  import { Dropdown } from '/@/components/Dropdown';
  import { BasicForm, useForm } from '/@/components/Form';
  import { Button } from '/@/components/Button';
  import { isFunction } from '/@/utils/is';
  import { volumeList } from '/@/api/mstore/volume';
  import { formatUnixToTime } from '/@/utils/dateUtil';
  import { getPathInfo } from '/@/utils/filepath';

  //每行个数
  const grid = ref(12);
  const selectKey = ref(0);
  // slider属性
  const useSlider = (min = 6, max = 12) => {
    // 每行显示个数滑动条
    const getMarks = () => {
      const l = {};
      for (let i = min; i < max + 1; i++) {
        l[i] = {
          style: {
            color: '#fff',
          },
          label: i,
        };
      }
      return l;
    };
    return {
      min,
      max,
      marks: getMarks(),
      step: 1,
    };
  };

  const ListItem = List.Item;
  const CardMeta = Card.Meta;
  const TypographyText = Typography.Text;
  // 获取slider属性
  const sliderProp = computed(() => useSlider(4));
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
  //数据
  const data = ref([]);

  const pathState = ref('');

  const imageShow = (item: any) => {
    if (item.is_dir) {
      return '/resource/img/folder.png';
    }
    if (item.ext === 'mp3') {
      return '/resource/img/flac.png';
    }
    if (item.ext === 'mp4') {
      return '/resource/img/mp4.png';
    }
    if (item.ext === 'exe') {
      return '/resource/img/exe.png';
    }
    if (item.ext === 'exe') {
      return '/resource/img/exe.png';
    }
    if (item.ext === 'zip') {
      return '/resource/img/zip.png';
    }
    return '/resource/img/misc.png';
  };

  const sizeShow = (limit: number) => {
    var size = '';
    if (limit < 1 * 1024) {
      //小于0.1KB，则转化成B
      size = limit + 'B';
    } else if (limit < 1024 * 1024) {
      size = (limit / 1024).toFixed(2) + 'KB';
    } else if (limit < 1024 * 1024 * 1024) {
      size = (limit / (1024 * 1024)).toFixed(2) + 'MB';
    } else {
      size = (limit / (1024 * 1024 * 1024)).toFixed(2) + 'GB';
    }
    var sizeStr = size + '';
    var index = sizeStr.indexOf('.');
    var dou = sizeStr.substr(index + 1, 2);
    if (dou == '00') {
      return sizeStr.substring(0, index) + sizeStr.substr(index + 3, 2);
    }
    return size;
  };

  //表单
  const [registerForm, { validate }] = useForm({
    schemas: [{ field: 'type', component: 'Input', label: '类型' }],
    labelWidth: 80,
    baseColProps: { span: 6 },
    actionColOptions: { span: 24 },
    autoSubmitOnEnter: true,
    submitFunc: handleSubmit,
  });

  const clickTimes = ref(0);

  const selectItem = (key: number) => {
    selectKey.value = key;
    clickTimes.value++;
    if (clickTimes.value === 2) {
      clickTimes.value = 0;
      selectKey.value = 0;
      emit('selectDir', data.value[key - 1]);
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
    const data = await validate();
    await fetch(data);
  }

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
  });

  async function fetch(p = {}) {
    const { path, params } = props;
    if (path) {
      const info = getPathInfo(path);
      const res = await volumeList(info[0], {
        path: info[1],
      });
      data.value = res.list;
      total.value = res.total;
    }
  }
  const total = ref(0);

  async function handleDelete(id) {
    emit('selectDir', id);
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
</style>
