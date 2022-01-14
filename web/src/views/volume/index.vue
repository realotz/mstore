<template>
  <PageWrapper dense contentFullHeight fixedHeight>
    <div class="flex h-full">
      <VolumeTree
        class="w-1/4 xl:w-1/5"
        @select="handleSelect"
        :path="pathRef"
        :volumes="volumes"
      />
      <FileList
        class="w-3/4 xl:w-4/5"
        :path="pathRef"
        @selectDir="handleSelectDir"
        :volumes="volumes"
      >
        <!-- <template #header>
        <Button type="primary" color="error"> 按钮1 </Button>
        <Button type="primary" color="success"> 按钮2 </Button>
      </template> -->
      </FileList>
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

  const volumes = ref([]);

  // 存储卷列表
  async function fetch() {
    const res = await getVolumeList();
    volumes.value = res.list;
  }

  // 自动请求并暴露内部方法
  onMounted(() => {
    fetch();
  });

  const pathRef = ref('');
  const handleSelectDir = (item) => {
    if (item.path == '/') {
      handleSelect('/' + item.volume_id + item.path + item.name);
    } else {
      handleSelect('/' + item.volume_id + item.path + '/' + item.name);
    }
  };
  const handleSelect = (key) => {
    pathRef.value = key;
  };
</script>
