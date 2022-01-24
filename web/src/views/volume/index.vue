<template>
  <PageWrapper dense contentFullHeight fixedHeight>
    <FileList :path="pathRef" @selectDir="handleSelectDir" />
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { h, ref, onMounted } from 'vue';
  import { Tag } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';
  import { Description, DescItem, useDescription } from '/@/components/Description/index';
  import { getVolumeList } from '/@/api/mstore/volume';
  import FileList from './FileList.vue';
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
</script>
