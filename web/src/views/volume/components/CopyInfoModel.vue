<template>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    title="文件冲突"
    @ok="handleSubmit"
    :draggable="false"
    :showOkBtn="false"
    :showCancelBtn="false"
    :canFullscreen="false"
  >
    <div class="h-full w-full p-1">
      <a-button class="w-full m-1" @click="handleOk" style="height: 50px" danger>
        <span>覆盖</span><br />
        <span style="font-size: 12px">文件将覆盖到目标目录中</span>
      </a-button>
      <a-button class="w-full m-1" @click="handleRname" style="height: 50px">
        <span>重命名</span><br />
        <span style="font-size: 12px">文件名以copy后缀进行重命名</span>
      </a-button>
      <a-button class="w-full m-1" @click="handleCancel" style="height: 50px">
        <span>取消</span>
      </a-button>
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, nextTick } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  export default defineComponent({
    components: { BasicModal },
    props: {
      data: { type: Object },
    },
    emits: ['ok'],
    setup(props, { emit }) {
      const dataRef = ref({});
      const [register, { closeModal, setModalProps }] = useModalInner((data) => {
        data && onDataReceive(data);
      });
      function onDataReceive(data) {
        dataRef.value = data;
      }
      function handleOk(v) {
        emit('ok', dataRef.value, 1);
        closeModal();
      }
      function handleRname(v) {
        console.log(dataRef.value);
        emit('ok', dataRef.value, 2);
        closeModal();
      }
      function handleCancel(v) {
        closeModal();
      }
      return {
        register,
        handleOk,
        handleRname,
        handleCancel,
      };
    },
  });
</script>
