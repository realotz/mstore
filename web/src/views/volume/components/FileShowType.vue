<template>
  <Dropdown>
    <Button value="small">
      <template #icon> <AppstoreOutlined /> </template>
    </Button>
    <template #overlay>
      <Menu selectable @click="handleSelect" v-model:selectedKeys="selectedShowKeys">
        <MenuItem :key="0">
          <span>列表视图</span>
        </MenuItem>
        <MenuItem :key="1">
          <span>小图标</span>
        </MenuItem>
        <MenuItem :key="2">
          <span>大图标</span>
        </MenuItem>
      </Menu>
    </template>
  </Dropdown>
</template>
<script lang="ts">
  import { AppstoreOutlined } from '@ant-design/icons-vue';
  import { Dropdown, Menu, MenuItem, Button } from 'ant-design-vue';
  import { defineComponent, ref, watch } from 'vue';
  export default defineComponent({
    name: 'FlieShowType',
    components: { Dropdown, Menu, MenuItem, Button, AppstoreOutlined },
    props: {},
    emits: ['select'],
    setup(props, { emit }) {
      const selectedShowKeys = ref([1]);
      const handleSelect = (item) => {
        emit('select', item.key);
      };
      return { selectedShowKeys, handleSelect };
    },
  });
</script>
<style lang="less"></style>
