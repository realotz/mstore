<template>
  <PageWrapper dense contentFullHeight fixedHeight>
    <div class="flex h-full">
      <VolumeTree
        class="w-1/4 xl:w-1/5"
        :resetPath="resetPath"
        @select="handleSelect"
        :path="pathRef"
      />
      <FileList
        class="w-3/4 xl:w-4/5"
        @resetDir="resetPathHandle"
        :path="pathRef"
        @selectDir="handleSelectDir"
      />
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { h, ref, onMounted } from 'vue';
  import { Tag } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';
  import { Description, DescItem, useDescription } from '/@/components/Description/index';
  import { getVolumeList } from '/@/api/mstore/volume';
  import FileList from './FileList.vue';
  import VolumeTree from './VolumeTree.vue';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  const volumeStore = useVolumeStoreWithOut();

  // 自动请求并暴露内部方法
  onMounted(() => {
    volumeStore.volumeList();
  });

  const pathRef = ref('');
  const resetPath = ref('');
  const handleSelectDir = (path) => {
    handleSelect(path);
  };
  const handleSelect = (key) => {
    pathRef.value = key;
  };
  const resetPathHandle = (key) => {
    console.log(key);
    resetPath.value = key;
  };
</script>
