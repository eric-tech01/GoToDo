<template>
    <div class="input_container">
        <el-input class="title" v-model="content" clerarable :autosize="{ minRows: 2, maxRows: 4 }"
            placeholder="请输入待办任务" type="textarea" show-word-limit maxlength="100" @keyup.enter="submit" />
        <div class="btn_ok">
            <el-button style="color: gray;" @click="remove" :icon="Delete" link></el-button>
            <el-button style="color: green;" @click="submit" :icon="Select" link></el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Select, Delete } from '@element-plus/icons-vue'
import { todoModel } from '../../wailsjs/go/models';
import { useI18n } from 'vue-i18n' // 多语言
const { t } = useI18n() // t方法取出，t('code')使用

const emits = defineEmits(['submitEvent', "cancelEvent"]);

const props = defineProps({
    modelValue: {
        type: todoModel.Todo,
        required: true,
    }
})
const content = ref(props.modelValue.title)
const submit = () => {
    emits('submitEvent', {
        ...props.modelValue,
        title: content.value
    })
}
const remove = () => {
    emits('cancelEvent')
}
</script>

<style lang="less" scoped>
.input_container {
    display: flex;
    flex-direction: row;
    padding-top: 10px;
    padding-bottom: 10px;

    .todo {
        // flex: 1;
    }

    .btn_ok {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
    }

    .el-button+.el-button {
        margin-left: 0;
    }
}
</style>