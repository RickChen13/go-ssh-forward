<template>
    <div class="main">
        <div class="head">
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="1">转发规则</el-menu-item>
                <el-menu-item index="2">服务器配置</el-menu-item>
            </el-menu>
        </div>
        <div class="container" ref="containerRef">
            <div class="top-pane" :style="{ height: topPaneHeight }">
                <div class="content" v-show="activeIndex == '1'">
                    <ForwardVue :height="topPaneHeightNum" />
                </div>
                <div class="content" v-show="activeIndex == '2'">
                    <SshServerVue :height="topPaneHeightNum" />
                </div>
            </div>

            <div class="splitter" ref="splitterRef" @mousedown="startResize" :class="{ 'resizing': isResizing }"></div>

            <div class="bottom-pane">
                <LogVue />
            </div>
        </div>

    </div>
</template>

<script lang="ts">
import Component from './Index';
const components = new Component();
export default components.vue();
</script>

<style lang="scss" scoped>
@use './Index.scss';
</style>
