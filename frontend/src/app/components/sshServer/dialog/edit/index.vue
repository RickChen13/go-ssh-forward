<template>
    <el-dialog v-model="visible" :close-on-click-modal="false" @closed="cancel">
        <div class="dialog-body">
            <el-form label-width="70px" :model="localData">
                <el-form-item :label="fieldName['name']">
                    <el-input v-model="localData.name" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['host']">
                    <el-input v-model="localData.host" class="form-input" placeholder="ip:port"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['user']">
                    <el-input v-model="localData.user" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['pass']">
                    <el-input v-model="localData.pass_new" class="form-input" type="password"
                        :placeholder="localData.pass != '' ? '不修改请留空' : '若有密码，请输入密码'"></el-input>
                    <el-checkbox v-model="localData.clear_pass" class="form-checkbox"
                        v-if="localData.pass != ''">清空密码</el-checkbox>
                </el-form-item>

                <el-form-item :label="fieldName['key_path']">
                    <el-input v-model="localData.key_path" class="form-input"></el-input>
                </el-form-item>

                <el-form-item :label="fieldName['pass_phrase']">
                    <el-input v-model="localData.pass_phrase_new" class="form-input" type="password"
                        :placeholder="localData.pass_phrase != '' ? '不修改请留空' : '若有密码，请输入密码'"></el-input>
                    <el-checkbox v-model="localData.clear_pass_phrase" class="form-checkbox"
                        v-if="localData.pass_phrase != ''">清空密码</el-checkbox>
                </el-form-item>

                <el-form-item :label="fieldName['sort']">
                    <el-input v-model="localData.sort" class="form-input" type="number" min="0" max="99999"></el-input>
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
