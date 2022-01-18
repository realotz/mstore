<template>
  <div class="m-2 mr-0 overflow-hidden bg-white">
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
</template>
<script lang="ts">
  import { defineComponent, onMounted, ref, unref, watch } from 'vue';
  import { BasicTree, TreeItem } from '/@/components/Tree';
  import { getVolumeList, volumeList } from '/@/api/mstore/volume';
  import { pathFmt, getPathInfo } from '/@/utils/filepath';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  const volumeStore = useVolumeStoreWithOut();
  const oldPath = ref('');
  export default defineComponent({
    name: 'VolumeList',
    components: { BasicTree },
    props: {
      path: {
        type: String,
        require: true,
      },
    },
    emits: ['select'],
    setup(props, { emit }) {
      watch(
        () => volumeStore.getVolumes,
        (v: string) => {
          fetch();
        },
      ),
        watch(
          () => props.path,
          (v: string) => {
            if (!v) {
              return;
            }
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
            oldPath.value = v;
            selectedKeys.value = [v];
          },
        );
      const treeData = ref<TreeItem[]>([]);
      const asyncTreeRef = ref<Nullable<TreeActionType>>(null);
      const expandedKeys = ref<string[]>([]);
      const selectedKeys = ref<string[]>([]);

      // 存储卷列表
      async function fetch() {
        treeData.value = [];
        const volumes = volumeStore.getVolumes;
        volumes.map((item) => {
          treeData.value.push({
            title: item.name,
            key: '/' + item.id,
            icon: 'home|svg',
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
            icon: 'home|svg',
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
        console.log(keys);
        if (keys[0]) {
          if (oldPath.value != '') {
            volumeStore.addBackPath(oldPath.value);
            volumeStore.resetAdvancePath();
          }
          emit('select', keys[0]);
        } else {
          selectedKeys.value = [oldPath.value];
        }
      }
      onMounted(() => {
        fetch();
      });
      return { treeData, handleSelect, onLoadData, asyncTreeRef, expandedKeys, selectedKeys };
    },
  });
</script>
