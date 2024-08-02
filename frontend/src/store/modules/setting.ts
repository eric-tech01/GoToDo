
import { set } from "nprogress";
import { defineStore } from "pinia";
import { Load, Save } from "../../../wailsjs/go/gateway/SettingController";

import { systemModel } from '../../../wailsjs/go/models';
const useSettingStore = defineStore("setting", {
  state: ():systemModel.Setting => {
    return {
        authOpened: false,
        isShow: false,
        id: 0,
    }

  },
  getters: {
    settingInfo(state: systemModel.Setting): systemModel.Setting {
        return {...state};
    }
  },
  actions:{
    setSettingInfo(setting: systemModel.Setting){
        this.$patch(setting);
    },
    async fetch() {
      Load().then((rsp) => {
          this.setSettingInfo(rsp.data)
      })
    },
    async updateSetting(s: systemModel.Setting) {
        try {
            Save(s).then((rsp) => {
                // if (rsp.)
                console.log(rsp)
                this.setSettingInfo(rsp.data)
            })
        } catch (error) {
            throw error
        }
    },
    
  }
})

export default useSettingStore