<template>
    <div class="content">
        <div class="tools">
            <el-button type="primary" :icon="Refresh" @click="getList">刷新</el-button>
            <el-button type="primary" :icon="Plus" @click="addVisible = true">添加</el-button>
        </div>
        <div class="data">
            <el-table :data="list" border style="width: 100%;" :height="height">
                <el-table-column align="center" prop="id" :label="fieldName['id']" />
                <el-table-column align="center" prop="name" :label="fieldName['name']" />
                <el-table-column align="center" prop="ssh_server_name" :label="fieldName['ssh_server_name']" />
                <el-table-column align="center" :label="fieldName['tag']" min-width="120">
                    <template #default="scope">
                        {{ tagField[scope.row.tag] }}
                    </template>
                </el-table-column>
                <el-table-column align="center" prop="local_addr" :label="fieldName['local_addr']" />
                <el-table-column align="center" prop="remote_addr" :label="fieldName['remote_addr']" />
                <el-table-column align="center" :label="fieldName['status']" min-width="120">
                    <template #default="scope">
                        {{ scope.row.status ? '运行中' : '未运行' }}
                    </template>
                </el-table-column>
                <el-table-column align="center" fixed="right" :label="fieldName['operate']" min-width="150">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click="editRow(scope.row)">编辑</el-button>
                        <el-button link type="primary" size="small" @click="delRow(scope.row.id)">删除</el-button>

                        <template v-if="scope.row.status === true">
                            <el-button link type="primary" size="small" @click="close(scope.row.id)">
                                停止
                            </el-button>
                        </template>
                        <template v-else>
                            <el-button link type="primary" size="small" @click="RunById(scope.row.id)">启动</el-button>
                        </template>
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
