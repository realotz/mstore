import { createVNode, render } from 'vue';
import ContextMsgConstructor from './ContextMsg.vue';
let instence: HTMLElement | undefined;
let instenceIsExit = false;
const ShowMsg = function (options: any) {
  if (instenceIsExit) {
    return;
  }
  const vm = createVNode(ContextMsgConstructor, options);
  const container = document.createElement('div');
  render(vm, container);
  instence = container.firstElementChild as HTMLElement;
  document.body.appendChild(instence);
  instenceIsExit = true;
  return instence;
};

const closeMsg = () => {
  if (instenceIsExit) {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    instenceIsExit = false;
    document.body.removeChild(instence as HTMLElement);
    instence = undefined;
  }
};
export { ShowMsg, closeMsg };
