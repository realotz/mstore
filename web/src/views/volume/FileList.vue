<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
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
    //api
    api: {
      type: Function,
      default: null,
    },
  });
  //暴露内部方法
  const emit = defineEmits(['getMethod', 'delete']);
  //数据
  const data = ref([]);
  // 切换每行个数
  // cover图片自适应高度
  //修改pageSize并重新请求数据

  const height = computed(() => {
    return `h-${120 - grid.value * 6}`;
  });
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
  function sliderChange(n) {
    pageSize.value = n * 4;
    fetch();
  }

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
    emit('getMethod', fetch);
  });

  async function fetch(p = {}) {
    const { api, params } = props;
    if (api && isFunction(api)) {
      const res = await api({ ...params, page: page.value, pageSize: pageSize.value, ...p });
      data.value = res.items;
      total.value = res.total;
    }
  }
  //分页相关
  const page = ref(1);
  const pageSize = ref(36);
  const total = ref(0);
  const paginationProp = ref({
    showSizeChanger: false,
    showQuickJumper: true,
    pageSize,
    current: page,
    total,
    onChange: pageChange,
    onShowSizeChange: pageSizeChange,
  });

  function pageChange(p, pz) {
    page.value = p;
    pageSize.value = pz;
    fetch();
  }
  function pageSizeChange(_current, size) {
    pageSize.value = size;
    fetch();
  }

  async function handleDelete(id) {
    emit('delete', id);
  }
</script>

<template>
  <div class="p-2 h-full">
    {{ sliderProp.width }}
    <div class="p-2 bg-white h-full">
      <List
        class="h-full"
        :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: 6, xxl: grid }"
        :data-source="data"
        :pagination="paginationProp"
      >
        <template #header>
          <div class="flex justify-end space-x-2"
            ><slot name="header"></slot>
            <Tooltip>
              <template #title>
                <div class="w-50">每行显示数量</div
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
          <ListItem>
            <Card>
              <template #title></template>
              <template #cover>
                <div :class="height">
                  <Image :src="item.imgs[0]" />
                </div>
              </template>
              <template #actions>
                <EditOutlined key="edit" />
                <Dropdown
                  :trigger="['hover']"
                  :dropMenuList="[
                    {
                      text: '删除',
                      event: '1',
                      popConfirm: {
                        title: '是否确认删除',
                        confirm: handleDelete.bind(null, item.id),
                      },
                    },
                  ]"
                  popconfirm
                >
                  <EllipsisOutlined key="ellipsis" />
                </Dropdown>
              </template>

              <CardMeta>
                <template #title>
                  <TypographyText :content="item.name" :ellipsis="{ tooltip: item.address }" />
                </template>
                <template #avatar>
                  <Avatar :src="item.avatar" />
                </template>
                <template #description>{{ item.time }}</template>
              </CardMeta>
            </Card>
          </ListItem>
        </template>
      </List>
    </div>
  </div>
</template>
