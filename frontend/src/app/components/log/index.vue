<template>
    <div class="fixed-bottom-right" v-show="logData.length > 0">
        <el-button type="danger" :icon="Delete" circle @click="clear" />
    </div>
    <div class="log-container">
        <div v-for="(data, key) in logData" :key="key" :id="`log-${key}`" v-html="data">
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
