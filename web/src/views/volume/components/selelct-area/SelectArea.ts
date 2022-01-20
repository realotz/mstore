import { createVNode, render } from 'vue';
import SelectAreaConstructor from './SelectArea.vue';
let instence: HTMLElement | undefined;
let instenceIsExit = false;
const SelectArea = function (options: any) {
  if (instenceIsExit) {
    document.body.removeChild(instence as HTMLElement);
    instenceIsExit = false;
  }
  const vm = createVNode(SelectAreaConstructor, options);
  const container = document.createElement('div');
  render(vm, container);
  instence = container.firstElementChild as HTMLElement;
  document.body.appendChild(instence);
  instenceIsExit = true;
  return instence;
};

const closeArea = () => {
  if (instenceIsExit) {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    instenceIsExit = false;
    document.body.removeChild(instence as HTMLElement);
    instence = undefined;
  }
};
export { SelectArea, closeArea };
