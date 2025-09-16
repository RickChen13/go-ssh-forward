<template>
    <el-dialog v-model="visible" :close-on-click-modal="false" @closed="cancel">
        <div class="dialog-body">
            <el-form label-width="70px" :model="localData">
                <el-form-item :label="fieldName['name']">
                    <el-input v-model="localData.name" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['remote_addr']">
                    <el-input v-model="localData.remote_addr" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['local_addr']">
                    <el-input v-model="localData.local_addr" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['tag']">
                    <el-select v-model="localData.tag" placeholder="Select">
                        <el-option v-for="item in tagSelect" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>

                <el-form-item :label="fieldName['css_id']">
                    <el-select v-model="localData.css_id" placeholder="Select">
                        <el-option v-for="item in serverList" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>

                <el-form-item label="权重">
                    <el-input v-model="localData.sort" class="form-input"></el-input>
                </el-form-item>
            </el-form>
        </div>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="cancel">取消</el-button>
                <el-button type="primary" @click="submit">提交</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script lang="ts">
import Component from "./index"
const components = new Component();
export default components.vue();
</script>

<style lang="scss" scoped>
@use "./index.scss";
</style>
