<template>
  <div class="m-2 mr-0 overflow-hidden bg-white">
    <BasicTree
      title="目录"
      :clickRowToExpand="false"
      :treeData="treeData"
      :load-data="onLoadData"
      ref="asyncTreeRef"
      @select="handleSelect"
    />
  </div>
</template>
<script lang="ts">
  import { defineComponent, onMounted, ref, unref } from 'vue';
  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { getVolumeList, volumeList } from '/@/api/mstore/volume';
  export default defineComponent({
    name: 'VolumeList',
    components: { BasicTree },
    emits: ['select'],
    setup(_, { emit }) {
      const treeData = ref<TreeItem[]>([]);
      const asyncTreeRef = ref<Nullable<TreeActionType>>(null);
      // 存储卷列表
      async function fetch() {
        const res = await getVolumeList();
        res.list.map((item) => {
          treeData.value.push({
            title: item.name,
            key: item.id,
            icon: 'home|svg',
            isLeaf: false,
            children: [],
          });
        });
      }
      // 异步展开存储卷
      async function onLoadData(treeNode) {
        console.log(treeNode.eventKey);
        const paths = treeNode.eventKey.split('/');
        let id = '';
        let path = '';
        if (paths.length == 1) {
          id = paths[0];
        } else {
          id = paths[1];
          paths.splice(0, 2);
          path = '/' + paths.join('/');
        }
        const res = await volumeList(id, {
          type: 2,
          path: path,
        });
        treeNode.children = res.list.map((item) => {
          console.log(item);
          if (item.path !== '') {
            item.path = '/' + item.path;
          }
          return {
            title: item.name,
            key: '/' + id + item.path + '/' + item.name,
            icon: 'home|svg',
            isLeaf: false,
            children: [],
          };
        });
        const asyncTreeAction: TreeActionType | null = unref(asyncTreeRef);
        if (asyncTreeAction) {
          if (treeNode.children.length > 0) {
            asyncTreeAction.updateNodeByKey(treeNode.eventKey, { children: treeNode.children });
          } else {
            asyncTreeAction.updateNodeByKey(treeNode.eventKey, { isLeaf: true });
          }
          asyncTreeAction.setExpandedKeys([
            treeNode.eventKey,
            ...asyncTreeAction.getExpandedKeys(),
          ]);
        }
      }
      function handleSelect(keys) {
        emit('select', keys[0]);
      }
      onMounted(() => {
        fetch();
      });
      return { treeData, handleSelect, onLoadData, asyncTreeRef };
    },
  });
</script>
