<template>
  <PageWrapper dense contentFullHeight fixedHeight contentClass="flex">
    <VolumeTree class="w-1/4 xl:w-1/5" @select="handleSelect" />
    <FileList
      class="w-3/4 xl:w-4/5"
      :params="params"
      :api="demoListApi"
      :path="pathRef"
      :volumeId="volumeIdRef"
      @getMethod="getMethod"
      @delete="handleDel"
    >
      <template #header>
        <Button type="primary" color="error"> 按钮1 </Button>
        <Button type="primary" color="success"> 按钮2 </Button>
      </template>
    </FileList>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { h, ref } from 'vue';
  import { Tag } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';
  import { Description, DescItem, useDescription } from '/@/components/Description/index';

  import FileList from './FileList.vue';
  import VolumeTree from './VolumeTree.vue';
  const pathRef = ref('');
  const volumeIdRef = ref('');
  const handleSelect = (key) => {
    const paths = key.split('/');
    let id = '';
    let path = '';
    if (paths.length == 1) {
      id = paths[0];
    } else {
      id = paths[1];
      paths.splice(0, 2);
      path = '/' + paths.join('/');
    }
    pathRef.value = path;
    volumeIdRef.value = id;
  };
</script>
