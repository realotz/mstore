<template>
  <PageWrapper dense contentFullHeight fixedHeight contentClass="flex">
    <VolumeTree />
    <FileList
      class="w-3/4 xl:w-4/5"
      :params="params"
      :api="demoListApi"
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
  import { h } from 'vue';
  import { Tag } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';
  import { Description, DescItem, useDescription } from '/@/components/Description/index';

  import FileList from './FileList.vue';
  import VolumeTree from './VolumeTree.vue';
</script>
