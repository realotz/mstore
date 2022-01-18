<template>
  <div>
    <span class="ant-cascader-picker" tabindex="0" style="width: 100%">
      <a-breadcrumb class="breadcrumb-picker-label" :routes="routes">
        <template #itemRender="{ route, paths }">
          <router-link to="" @click="handleClick(route, paths, $event)">
            {{ route.name }}
          </router-link>
        </template>
      </a-breadcrumb>
      <input autocomplete="off" type="text" class="ant-input ant-cascader-input" />
    </span>
  </div>
</template>
<script lang="ts">
  import type { RouteLocationMatched } from 'vue-router';
  import { useRouter } from 'vue-router';
  import { defineComponent, ref, watch } from 'vue';
  import { Breadcrumb } from 'ant-design-vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useVolumeStoreWithOut } from '/@/store/modules/volume';
  const volumeStore = useVolumeStoreWithOut();
  export default defineComponent({
    name: 'Breadcrumb',
    components: { [Breadcrumb.name]: Breadcrumb },
    props: {
      path: {
        type: String,
        require: true,
      },
    },
    emits: ['select'],
    setup(props, { emit }) {
      const routes = ref<RouteLocationMatched[]>([]);
      const { currentRoute } = useRouter();
      const { prefixCls } = useDesign('layout-breadcrumb');
      watch(
        () => props.path,
        (v: string) => {
          if (!v) {
            return;
          }
          const paths = v.split('/');
          const r = paths.map((item, index) => {
            if (index != 1) {
              return {
                name: item,
                key: paths.slice(0, index + 1).join('/'),
              };
            }
            for (let i = 0; i < volumeStore.getVolumes.length; i++) {
              if (volumeStore.getVolumes[i].id == item) {
                item = volumeStore.getVolumes[i].name;
                return {
                  name: item,
                  key: paths.slice(0, index + 1).join('/'),
                };
              }
            }
          });
          routes.value = r;
        },
      );
      function handleClick(route: RouteLocationMatched, paths: string[], e: Event) {
        const { key } = route;
        emit('select', key);
      }
      return { routes, handleClick };
    },
  });
</script>
<style lang="less">
  .ant-breadcrumb-separator {
    margin: 0 2px;
  }
  .breadcrumb-picker-label {
    position: absolute;
    top: 50%;
    left: 0;
    width: 100%;
    height: 20px;
    margin-top: -10px;
    padding: 0 20px 0 12px;
    overflow: hidden;
    line-height: 20px;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
</style>
