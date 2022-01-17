<template>
  <div :class="[prefixCls, `${prefixCls}--${theme}`]">
    <a-breadcrumb :routes="routes">
      <template #itemRender="{ route, paths }">
        <router-link to="" @click="handleClick(route, paths, $event)">
          {{ route.name }}
        </router-link>
      </template>
    </a-breadcrumb>
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
  @prefix-cls: ~'@{namespace}-layout-breadcrumb';

  .@{prefix-cls} {
    display: flex;
    padding: 0 8px;
    align-items: center;

    .ant-breadcrumb-link {
      .anticon {
        margin-right: 4px;
        margin-bottom: 2px;
      }
    }

    &--light {
      .ant-breadcrumb-link {
        color: @breadcrumb-item-normal-color;

        a {
          color: rgb(0 0 0 / 65%);

          &:hover {
            color: @primary-color;
          }
        }
      }

      .ant-breadcrumb-separator {
        color: @breadcrumb-item-normal-color;
      }
    }

    &--dark {
      .ant-breadcrumb-link {
        color: rgb(255 255 255 / 60%);

        a {
          color: rgb(255 255 255 / 80%);

          &:hover {
            color: @white;
          }
        }
      }

      .ant-breadcrumb-separator,
      .anticon {
        color: rgb(255 255 255 / 80%);
      }
    }
  }
</style>
