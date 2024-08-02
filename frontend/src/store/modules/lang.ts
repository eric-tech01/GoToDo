import { defineStore } from 'pinia';

const useLangStore = defineStore("lang", {
    state: () => {
        return {
            lang: "zh"
        }
    },
    getters: {
        getLang: (state) => state.lang
    },
    actions: {
        setLang(lang: string) {
            this.lang = lang;
        }
    }
})

export default useLangStore;