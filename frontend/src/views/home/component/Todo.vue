<template>
    <div class="todo-item">
        <el-text :tag="completed ? 'del' : 'p'" class="title" :class="completed ? 'completed' : ''" v-if="isHide"
            style="flex: 1;" size="large">{{ content
            }}</el-text>
        <el-text :tag="completed ? 'del' : 'p'" class="title" :class="completed ? 'completed' : ''" v-else
            size="large">{{ title }}</el-text>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue';

const props = defineProps({
    id: {
        type: Number,
        default: 0
    },
    title: {
        type: String,
        default: ""
    },
    completed: {
        type: Boolean,
        default: false
    },
    isHide: {
        type: Boolean,
        default: false
    }
})
const content = ref()

watch(() => props.isHide, (newVal) => {
    if (newVal) {
        content.value = "***" + props.title.substring(0, 2) + "***"
    }
})
onMounted(() => {
    if (props.isHide) {
        content.value = "***" + props.title.substring(0, 2) + "***"
    }
})

</script>

<style scoped lang="less">
.todo-item {
    display: flex;
    cursor: pointer;
    margin-bottom: 2px;

    :hover {
        background: rgba(255, 255, 255, 0.6);
    }

    .title {
        flex: 1;
        padding-left: 12px;
        padding-right: 8px;

        font-family: PingFang SC;
        font-size: 16px;
        font-weight: normal;
        line-height: 30px;
        display: flex;
        align-items: center;
        letter-spacing: 0px;

        color: #000000;
    }

    .completed {
        color: gray;
    }
}
</style>
