<template>
    <div class="container">
        <div class="menu">
            <el-button style="flex:1;border-style: dashed;" small @click="isAddNew = true" plain>
                <el-icon>
                    <Plus />
                </el-icon></el-button>
            <el-button style="border-width:0;" @click="handleHide" :icon="isShow ? Hide : View" />
        </div>
        <div class="input_area" v-if="isAddNew">
            <InputTodo v-model="todoValue" @submitEvent="submitEvent" @cancelEvent="cancelEvent">
            </InputTodo>
        </div>

        <div class="todo_area">
            <div class="todo_list" v-for="(item, index) in todoList">
                <Todo :key="item.id" :id="item.id" :title="item.title" :completed="item.completed"
                    @deleteEvent="deleteEvent" :is-hide="isShow" @contextmenu.prevent="onContextMenu($event, item.id)">
                </Todo>
                <LineDivider></LineDivider>
            </div>
        </div>
        <ContextMenu
            :menu="[{ key: 'edit', label: $t('common.edit') }, { key: 'delete', label: $t('common.delete') }, { key: 'finish', label: $t('common.finish') }]"
            :mouse-x="mouseX" :mouse-y="mouseY" :show-right-menu="showRightMenu" @contextmenuEvent="rightSelect">
        </ContextMenu>

    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref, toValue } from 'vue';
import Todo from './component/Todo.vue';
import { Plus, Hide, View } from '@element-plus/icons-vue';
import InputTodo from '@/components/InputTodo.vue';
import { Load, Save, Delete, Finish } from "../../../wailsjs/go/gateway/TodoController"
import { useSettingStore } from '@/store';
import { systemModel, todoModel } from '../../../wailsjs/go/models';
import ContextMenu from '@/components/ContextMenu.vue';
import LineDivider from '@/components/LineDivider.vue';
// import { useI18n } from 'vue-i18n' // 多语言
// const { t } = useI18n() // t方法取出，t('code')使用
const isAddNew = ref(false);
const settingStore = useSettingStore()
settingStore.fetch()

const isShow = computed(() => {
    return settingStore.isShow
})
const todoList = ref([] as todoModel.Todo[]);
const { updateSetting } = useSettingStore()
onMounted(() => {
    Load().then(rsp => {
        console.log(rsp)
        todoList.value.push(...rsp)
    })
    window.addEventListener("click", closeContextMenu, true)
    window.addEventListener("contextmenu", (e) => {
        e.preventDefault();
        closeContextMenu()
    }, true)
})

const handleHide = () => {
    updateSetting({ isShow: !isShow.value } as systemModel.Setting).then(r => {
        console.log("rrr: ", r)
    })
    // isHide.value = !isHide.value;
};

const handleDelete = (id: number) => {
    Delete(id).then(rsp => {
        console.log(rsp)
        Load().then(rsp => {
            console.log(rsp)
            todoList.value = []
            todoList.value.push(...rsp)
        })
    })
};
const handleFinish = (id: number) => {
    console.log(id);
    Finish(id).then(rsp => {
        Load().then(rsp => {
            todoList.value = []
            todoList.value.push(...rsp)
        })
    })
};
const handleEdit = (id: number) => {
    todoValue.value = todoList.value.find(item => item.id === id) as todoModel.Todo
    isAddNew.value = true;
}
const todoValue = ref<todoModel.Todo>({} as todoModel.Todo)
const submitEvent = (m: todoModel.Todo) => {
    Save(m).then(rsp => {
        Load().then(rsp => {
            console.log(rsp)
            todoList.value = []
            todoList.value.push(...rsp)
        })
    })
    isAddNew.value = false;
    console.log(m);
};
const cancelEvent = () => {
    isAddNew.value = false
    console.log('cancel');
};
const deleteEvent = (id: number) => {
    // todoList.value = todoList.value.filter(item => item.id !== id);
    Delete(id).then(rsp => {
        console.log(rsp)
        Load().then(rsp => {
            console.log(rsp)
            todoList.value = []
            todoList.value.push(...rsp)
        })
    })
};

const rightSelect = (key: string) => {
    console.log(key);
    switch (key) {
        case "edit":
            handleEdit(triggerItem.value)
            break
        case "delete":
            handleDelete(triggerItem.value)
            break;
        case "finish":
            console.log("finish: ", key)
            handleFinish(triggerItem.value)
            break;
        default:
            break;
    }
};

const showRightMenu = ref(false);
const mouseX = ref(0);
const mouseY = ref(0);
const triggerItem = ref(0);
const onContextMenu = (e: MouseEvent, id: number) => {
    e.preventDefault();
    console.log(e);
    showRightMenu.value = true;

    const viewportWidth = window.innerWidth;
    const menuWidth = 80
    if (e.clientX >= viewportWidth - menuWidth) {
        mouseX.value = viewportWidth - menuWidth
    } else {
        mouseX.value = e.clientX;
    }

    // mouseX.value = e.clientX;
    mouseY.value = e.clientY;
    triggerItem.value = id
};
const closeContextMenu = () => {
    showRightMenu.value = false;
    mouseX.value = 0;
    mouseY.value = 0;
};
</script>

<style scoped lang="less">
.container {
    height: 100vh;
    display: flex;
    flex-direction: column;

    overflow-x: hidden;

    background: linear-gradient(180deg, #BEDAFF 0%, #BEDAFF 0%, #F7F8FA 100%);

    .menu {
        // height: 30px;
        display: flex;
        justify-content: flex-end;
        margin-top: 12px;
        padding-left: 8px;
        padding-right: 8px;
    }

    .todo_area {
        margin-top: 20px;

        .todo_list {
            padding-top: 4px;
            padding-bottom: 4px;
        }
    }
}

.el-button+.el-button {
    margin-left: 4px;
}
</style>