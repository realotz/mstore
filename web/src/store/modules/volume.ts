import { defineStore } from 'pinia';
interface VolumeState {
  volumes: [];
}

export const useUserStore = defineStore({
  id: 'app-user',
  state: (): VolumeState => ({
    volumes: [],
  }),
  getters: {},
  actions: {},
});
