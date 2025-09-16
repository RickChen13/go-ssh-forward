<template>
    <div class="main">
        <div class="head">
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="1">转发规则</el-menu-item>
                <el-menu-item index="2">服务器配置</el-menu-item>
            </el-menu>
        </div>
        <div class="content" v-show="activeIndex == '1'">
            <ForwardVue />
        </div>
        <div class="content" v-show="activeIndex == '2'">
            <SshServerVue />
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
