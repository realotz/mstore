<template>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    :title="title"
    @visible-change="handleVisibleChange"
    @ok="handleSubmit"
  >
    <div class="pt-3px pr-3px">
      <BasicForm @register="registerForm" :model="model" />
    </div>
  </BasicModal>
</template>
<script lang="ts">
  import { defineComponent, ref, nextTick } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, FormSchema, useForm } from '/@/components/Form/index';
  const schemas: FormSchema[] = [
    {
      field: 'name',
      component: 'Input',
      label: '名称',
      colProps: {
        span: 24,
      },
      defaultValue: '',
    },
  ];
  export default defineComponent({
    components: { BasicModal, BasicForm },
    props: {
      userData: { type: Object },
    },
    emits: ['ok'],

    setup(props, { emit }) {
      const modelRef = ref({});
      const dataRef = ref({});
      const title = ref('重命名');
      const [registerForm, { getFieldsValue }] = useForm({
        labelWidth: 80,
        schemas,
        showActionButtonGroup: false,
        actionColOptions: {
          span: 24,
        },
      });
      const [register] = useModalInner((data) => {
        data && onDataReceive(data);
      });
      function onDataReceive(data) {
        modelRef.value = { name: data.name };
        dataRef.value = data;
        title.value = data.title ?? '重命名';
      }
      function handleVisibleChange(v) {
        v && props.userData && nextTick(() => onDataReceive(props.userData));
      }
      function handleSubmit(v) {
        let data = getFieldsValue();
        emit('ok', dataRef.value, data);
      }
      return {
        register,
        schemas,
        registerForm,
        model: modelRef,
        handleVisibleChange,
        handleSubmit,
        title,
      };
    },
  });
</script>
