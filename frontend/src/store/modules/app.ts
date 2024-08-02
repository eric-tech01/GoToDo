import { defineStore } from 'pinia';

export interface AppState {
  theme: 'light' | 'dark';
  device: string;
  menuCollapse: boolean;
  hideMenu: boolean;
}

const useAppStore = defineStore('app', {
  state: (): AppState => ({
    device: 'desktop',
    theme: 'light',
    menuCollapse: false,
    hideMenu: false,
  }),

  getters: {
    appCurrentSetting(state: AppState): AppState {
      return { ...state };
    },
    appDevice(state: AppState) {
      return state.device;
    },
  },

  actions: {
    // Update app settings
    updateSettings(partial: Partial<AppState>) {
      // @ts-ignore-next-line
      this.$patch(partial);
    },

    // Change theme color
    toggleTheme(dark: boolean) {
      if (dark) {
        this.theme = 'dark';
        document.body.setAttribute('arco-theme', 'dark');
      } else {
        this.theme = 'light';
        document.body.removeAttribute('arco-theme');
      }
    },
    toggleDevice(device: string) {
      this.device = device;
    },
    toggleMenu(value: boolean) {
      this.hideMenu = value;
    },
  },
});

export default useAppStore;
