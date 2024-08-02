import { createI18n } from 'vue-i18n' // 引入i18n创建多语言对象
import zhLocale from './language/zh' // 引入中文对应语言配置
import enLocale from './language/en'; // 引入英文对应语言配置

const messages = {
    zh: zhLocale,
    en: enLocale
}

const localLang = navigator.language.split('-')[0]; // 获取浏览器语言环境
const storageLang = window.localStorage.getItem('locale')?.split('"')[1].split('"')[0].toLocaleLowerCase() || 'en'; // 获取localStorage本地语言环境
const c = (storageLang.toLocaleLowerCase() !== 'zh' && storageLang.toLocaleLowerCase() !== 'en') ? 'en' : storageLang; // 如果localStorage语言环境既不是中文也不是英文就取为英文

const i18n = createI18n({
    //这里是语种的持久化，刷新也存在
  locale:localStorage.getItem('language') || 'zhCn', // 默认是中文
  fallbackLocale: 'zh', // 语言切换的时候是英文
  globalInjection:true,//全局配置
  legacy:false,//vue3写法
  messages:messages//本地message，也就是你需要做国际化的语种
})

export default i18n
