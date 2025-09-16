<template>
    <div class="Menu allow-move" @dblclick="ToggleMaximise">
        <div class="right">
            <div type="button" class="btn" @click="WindowMinimise">
                <ClarityWindowMinLine />
            </div>
            <div type="button" class="btn" @click="ToggleMaximise">
                <ClarityWindowMaxLine v-show="!isMaximised" />
                <ClarityWindowRestoreLine v-show="isMaximised" />
            </div>
            <div type="button" class="btn close" @click="WindowHide">
                <ClarityWindowCloseLine />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Component from "./index"
const components = new Component();
export default components.vue();
</script>

<style lang="scss" scoped>
@use "./index.scss";
</style>
