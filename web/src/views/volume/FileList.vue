<template>
  <div class="p-2 h-full">
    {{ sliderProp.width }}
    <div class="p-2 bg-white h-full">
      <List :grid="{ gutter: 16, column: 7 }" size="small" class="h-full" :data-source="data">
        <template #header>
          <div class="flex justify-end space-x-2"
            ><slot name="header"></slot>
            <Tooltip>
              <template #title>
                <div class="w-25">每行显示数量</div
                ><Slider
                  id="slider"
                  v-bind="sliderProp"
                  v-model:value="grid"
                  @change="sliderChange"
              /></template>
              <Button><TableOutlined /></Button>
            </Tooltip>
            <Tooltip @click="fetch">
              <template #title>刷新</template>
              <Button><RedoOutlined /></Button>
            </Tooltip>
          </div>
        </template>
        <template #renderItem="{ item }">
          <ListItem style="text-align: center">
            <div class="file-item">
              <Tooltip placement="bottom">
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
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref, watch } from 'vue';
  import {
    EditOutlined,
    EllipsisOutlined,
    RedoOutlined,
    TableOutlined,
  } from '@ant-design/icons-vue';
  import { List, Card, Image, Typography, Tooltip, Slider, Avatar } from 'ant-design-vue';
  import { Dropdown } from '/@/components/Dropdown';
  import { BasicForm, useForm } from '/@/components/Form';
  import { Button } from '/@/components/Button';
  import { isFunction } from '/@/utils/is';
  import { volumeList } from '/@/api/mstore/volume';
  import { formatUnixToTime } from '/@/utils/dateUtil';

  //每行个数
  const grid = ref(12);
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
    // 请求API的参数
    params: {
      type: Object,
      default: () => ({}),
    },
    volumeId: {
      type: String,
      default: '',
    },
    path: {
      type: String,
      default: '',
    },
  });

  watch(
    () => props.volumeId,
    (newv, oldv) => {
      fetch();
    },
  );
  watch(
    () => props.path,
    (newv, oldv) => {
      fetch();
    },
  );

  //暴露内部方法
  const emit = defineEmits(['getMethod', 'delete']);
  //数据
  const data = ref([]);

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

  //表单提交
  async function handleSubmit() {
    const data = await validate();
    await fetch(data);
  }

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
    emit('getMethod', fetch);
  });

  async function fetch(p = {}) {
    const { path, params, volumeId } = props;
    if (volumeId) {
      const res = await volumeList(volumeId, {
        path: path,
      });
      data.value = res.list;
      console.log(data.value);
      total.value = res.total;
    }
  }
  const total = ref(0);

  async function handleDelete(id) {
    emit('delete', id);
  }
</script>

<style lang="less">
  .file-item {
    width: 128px;
    height: 158px;
    display: block;
  }
  .file-item-box {
    height: 130px;
    width: 128px;
    text-align: center;
    vertical-align: bottom;
    display: table-cell;
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
