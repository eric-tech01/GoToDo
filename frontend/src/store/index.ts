import { createPinia } from 'pinia';
import useAppStore from './modules/app';
import useSettingStore from './modules/setting';
// import useUserStore from './modules/user';

const pinia = createPinia();

export { useAppStore, useSettingStore };
export default pinia;
