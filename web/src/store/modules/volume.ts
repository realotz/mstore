import { defineStore } from 'pinia';
import { store } from '/@/store';
import { getVolumeList } from '/@/api/mstore/volume';

interface VolumeState {
  volumes: [];
  backPaths: Array<String>;
  advancePaths: Array<String>;
}

export const useVolumeStore = defineStore({
  id: 'app-volume',
  state: (): VolumeState => ({
    volumes: [],
    backPaths: [],
    advancePaths: [],
  }),
  getters: {
    getVolumes(): any {
      return this.volumes;
    },
    getAdvancePaths(): Array<String> {
      return this.advancePaths;
    },
  },
  actions: {
    setVolumeList(list) {
      this.volumes = list;
    },
    getBackPath(): any {
      return this.backPaths.pop();
    },
    getAdvancePath(): any {
      return this.advancePaths.pop();
    },
    addBackPath(path) {
      this.backPaths.push(path);
    },
    addAdvancePath(path) {
      this.advancePaths.push(path);
    },
    resetAdvancePath() {
      this.advancePaths.splice(0, this.advancePaths.length);
    },
    async volumeList() {
      try {
        const res = await getVolumeList();
        this.setVolumeList(res.list);
        return res;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  },
});

// Need to be used outside the setup
export function useVolumeStoreWithOut() {
  return useVolumeStore(store);
}
