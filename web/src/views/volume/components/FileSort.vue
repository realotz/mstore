<template>
  <Dropdown>
    <Button value="small">
      <template #icon> <FunnelPlotOutlined /> </template>
    </Button>
    <template #overlay>
      <Menu multiple @click="handleSelect" v-model:selectedKeys="selectedShowKeys">
        <MenuItem key="name">
          <span>名称</span>
        </MenuItem>
        <MenuItem key="size">
          <span>大小</span>
        </MenuItem>
        <MenuItem key="updated_at">
          <span>修改日期</span>
        </MenuItem>
        <MenuItem key="ext">
          <span>文件类型</span>
        </MenuItem>
        <MenuDivider />
        <MenuItem key="asc">
          <span>从小到大</span>
        </MenuItem>
        <MenuItem key="desc">
          <span>从大到小</span>
        </MenuItem>
      </Menu>
    </template>
  </Dropdown>
</template>
<script lang="ts">
  import { FunnelPlotOutlined } from '@ant-design/icons-vue';
  import { Dropdown, Menu, MenuItem, Button, MenuDivider } from 'ant-design-vue';
  import { defineComponent, ref, watch } from 'vue';
  export default defineComponent({
    name: 'FlieSort',
    components: { Dropdown, Menu, MenuItem, Button, MenuDivider, FunnelPlotOutlined },
    props: {},
    emits: ['select'],
    setup(props, { emit }) {
      const selectedShowKeys = ref(['name', 'asc']);
      const handleSelect = (item) => {
        let sort = 'asc';
        let sort_name = 'name';
        for (let i = 0; i < selectedShowKeys.value.length; i++) {
          if (selectedShowKeys.value[i] == 'asc' || selectedShowKeys.value[i] == 'desc') {
            sort = selectedShowKeys.value[i];
          } else {
            sort_name = selectedShowKeys.value[i];
          }
        }
        if (item.key == 'asc' || item.key == 'desc') {
          sort = item.key;
        } else {
          sort_name = item.key;
        }
        selectedShowKeys.value = [sort_name, sort];
        emit('select', sort_name, sort);
      };
      return { selectedShowKeys, handleSelect };
    },
  });
</script>
<style lang="less"></style>
