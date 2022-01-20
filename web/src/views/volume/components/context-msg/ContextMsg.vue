<template>
  <div
    class="context-msg ant-message-notice v-leave-from ant-move-up ant-move-up-leave"
    :style="[{ top: Point.y + 'px' }, { left: Point.x + 'px' }]"
  >
    <div class="ant-message-notice-content">
      <div class="ant-message-custom-content ant-message-success">
        <span role="img" aria-label="check-circle">
          <CheckCircleFilled style="padding: 5px" />
        </span>
        <span>{{ Text.text }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  import { CheckCircleFilled } from '@ant-design/icons-vue';
  import { computed, defineComponent, PropType } from 'vue';
  interface Point {
    x: number;
    y: number;
  }
  export default defineComponent({
    name: 'ContextMsg',
    components: { CheckCircleFilled },
    props: {
      text: {
        type: String,
        required: true,
      },
      point: {
        type: Object as PropType<Point>,
        required: true,
      },
    },
    setup(props) {
      let Point = computed(() => {
        let x = props.point.x;
        let y = props.point.y;
        return {
          x,
          y,
        };
      });
      let Text = computed(() => {
        let text = props.text;
        return {
          text,
        };
      });
      return {
        Point,
        Text,
      };
    },
  });
</script>
<style lang="less" scoped>
  .context-msg {
    position: fixed;
    z-index: 9;
  }
</style>
