<template>
  <BasicDrawer v-bind="$attrs" @register="register" :isDetail="true" :title="title">
    <CodeEditor v-model:value="value" :mode="modeValue" />
    <template #titleToolbar> <div></div></template>
  </BasicDrawer>
</template>
<script lang="ts">
  import { defineComponent, ref, unref } from 'vue';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { CodeEditor, JsonPreview, MODE } from '/@/components/CodeEditor';
  import { downUrl } from '/@/api/mstore/volume';
  import { pathFmt } from '/@/utils/filepath';
  import { defHttp } from '/@/utils/http/axios';
  export default defineComponent({
    components: { BasicDrawer, CodeEditor },
    setup() {
      const modeValue = ref<MODE>(MODE.JSON);
      const value = ref('');
      const title = ref('编辑器');
      const [register, { closeDrawer, changeLoading, setDrawerProps }] = useDrawerInner((data) => {
        data && onDataReceive(data);
      });
      function onDataReceive(data) {
        value.value = '';
        if (!data.volume_id) {
          return;
        }
        changeLoading(true);
        if (data.ext == '.json') {
          modeValue.value = MODE.JSON;
        }
        if (data.ext == '.html') {
          modeValue.value = MODE.HTML;
        }
        if (data.ext == '.js' || data.ext == '.tx' || data.ext == '.vue') {
          modeValue.value = MODE.JS;
        }
        title.value = data.name ?? '编辑器';
        defHttp
          .get({
            url: '/v1/files/' + data.volume_id,
            params: { path: pathFmt(data.path + '/' + data.name) },
          })
          .then((res) => {
            changeLoading(false);
            value.value = res.data;
          });
      }
      return {
        modeValue,
        value,
        title,
        register,
      };
    },
  });
</script>
