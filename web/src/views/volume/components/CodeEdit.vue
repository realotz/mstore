<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="register"
    :isDetail="true"
    :closeFunc="closeFunc"
    :title="title"
  >
    <CodeEditor v-model:value="value" :mode="modeValue" @update:value="editUpdateHandle" />
    <template #titleToolbar> <div></div></template>
  </BasicDrawer>
</template>
<script lang="ts">
  import { defineComponent, ref, unref } from 'vue';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { CodeEditor, JsonPreview, MODE } from '/@/components/CodeEditor';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { downUrl, saveFile } from '/@/api/mstore/volume';
  import md5 from 'js-md5';
  import { pathFmt } from '/@/utils/filepath';
  import { defHttp } from '/@/utils/http/axios';

  export default defineComponent({
    components: { BasicDrawer, CodeEditor },
    setup() {
      const { createMessage, createConfirm } = useMessage();
      const modeValue = ref<MODE>(MODE.JSON);
      const value = ref('');
      const title = ref('编辑器');
      let dataRow;
      let md5Value = '';
      let outOnkeydown;
      const handleKeydown = (e) => {
        let key = window.event.keyCode;
        if (!e.altKey && !e.shiftKey && key === 83 && (e.metaKey || e.ctrlKey)) {
          window.event.preventDefault();
          saveFileHandle();
        }
      };
      const [register, { closeDrawer, changeLoading, setDrawerProps }] = useDrawerInner((data) => {
        outOnkeydown = document.onkeydown;
        document.onkeydown = handleKeydown;
        data && onDataReceive(data);
      });
      function saveFileHandle() {
        saveFile(dataRow.volume_id, {
          path: pathFmt(dataRow.path + '/' + dataRow.name),
          data: value.value,
        }).then((res) => {
          md5Value = md5(value.value);
          title.value = dataRow.name ?? '编辑器';
          createMessage.success('已保存');
        });
      }
      function closeFunc() {
        if (md5Value != md5(value.value)) {
          let flag = false;
          createConfirm({
            iconType: 'warning',
            title: '确认',
            content: `文件未保存，是否退出`,
            onOk: () => {
              document.onkeydown = outOnkeydown;
              closeDrawer();
            },
          });
          return false;
        } else {
          document.onkeydown = outOnkeydown;
          return true;
        }
      }
      function editUpdateHandle(v) {
        if (md5Value == '') {
          md5Value = md5(v);
        }
        if (md5Value == md5(v)) {
          title.value = dataRow.name ?? '编辑器';
        } else {
          title.value = '*' + (dataRow.name ?? '编辑器');
        }
      }
      function onDataReceive(data) {
        dataRow = data;
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
            md5Value = md5(value.value);
          });
      }
      return {
        modeValue,
        editUpdateHandle,
        value,
        title,
        register,
        closeFunc,
      };
    },
  });
</script>
