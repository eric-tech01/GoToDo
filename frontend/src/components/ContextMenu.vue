<template>
    <div>
        <Transition @before-enter="handleBeforeEnter" @enter="handleEnter" @after-enter="handleafterEnter">
            <div class="context-menu" v-if="showRightMenu"
                :style="{ position: 'absolute', left: mouseX + 'px', top: mouseY + 'px' }">
                <div class="context-menu-item" v-for="item in menu" :key="item.key"
                    @click="$emit('contextmenuEvent', item.key)">
                    {{ item.label }}
                </div>
            </div>
        </Transition>
    </div>
</template>
<script setup lang="ts">
import { RightMenuI } from '@/types/types';
import { onMounted, ref } from 'vue';

defineProps({
    menu: {
        type: Array<RightMenuI>,
        default: () => ([])
    },
    mouseX: {
        type: Number,
        default: 0
    },
    mouseY: {
        type: Number,
        default: 0
    },
    showRightMenu: {
        type: Boolean,
        default: false
    }
})

const handleBeforeEnter = (e: Element) => {
    const el = e as HTMLElement
    el.style.height = "0";
    // el.style.opacity = '0'
}

const handleEnter = (e: Element) => {
    const el = e as HTMLElement
    el.style.height = 'auto'
    const h = el.clientHeight
    el.style.height = "0"
    requestAnimationFrame(() => {
        el.style.height = h + 'px'
        el.style.transition = '.5s'
    })
}
const handleafterEnter = (e: Element) => {
    const el = e as HTMLElement
    el.style.transition = 'none'
}
</script>
<style lang="less" scoped>
.context-menu {
    background-color: #fff;
    padding-top: 8px;
    padding-left: 8px;
    padding-right: 8px;
    width: 60px;
    border-radius: 4px;

    .context-menu-item {
        cursor: pointer;
        // margin-bottom: 6px;
        border-radius: 4px;
        padding: 4px;
        font-size: 14px;
        font-weight: 500;
    }

    .context-menu-item:hover {
        background-color: #409eff;
    }
}
</style>