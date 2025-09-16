<template>
    <div class="content">
        <div class="tools">
            <el-button type="primary" :icon="Refresh" @click="getList">刷新</el-button>
            <el-button type="primary" :icon="Plus" @click="addVisible = true">添加</el-button>
        </div>
        <div class="data">
            <el-table :data="list" border style="width: 100%">
                <el-table-column align="center" prop="id" label="ID" />
                <el-table-column align="center" prop="name" label="名称" />
                <el-table-column align="center" prop="host" label="远程地址" />
                <el-table-column align="center" prop="user" label="用户" />

                <el-table-column align="center" fixed="right" label="操作" min-width="150">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="editRow(scope.row)">编辑</el-button>
                        <el-button link type="primary" size="small" @click="delRow(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <AddVue v-if="addVisible" @change="change" />
        <EditVue v-if="editVisible" @change="change" :data="editData" />
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
