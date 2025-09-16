import { defineStore } from 'pinia';
import { List } from '@wails/go/bll/ForwardRuleBll';
import { Status } from '@wails/go/bll/forwardBll';
export type CuForwardRule = {
    name: string;
    remote_addr: string;
    local_addr: string;
    tag: string;
    sort: number;
    css_id: number;
}

export type ForwardRule = CuForwardRule & {
    id: number;
    ssh_server_name: string;
    status?: boolean;
}

type ForwarStore = {
    page: number;
    pageSize: number;
    count: number;
    list: ForwardRule[];
    fieldName: FieldNameMap;
    tagField: FieldNameMap;
    tagSelect: ElSelectType[];
}

export const useForwardStore = defineStore('forwar', {
    state: (): ForwarStore => {
        let tagField = {
            "L": "远程转发到本地",
            "R": "本地转发到远程",
        };
        let tagSelect: ElSelectType[] = Object.entries(tagField).map(([key, value]) => ({
            label: value,
            value: key,
        }));

        return {
            page: 1,
            pageSize: 1000,
            count: 0,
            list: [],
            fieldName: {
                id: 'ID',
                name: '名称',
                ssh_server_name: '服务器',
                remote_addr: '远程地址',
                local_addr: '本地地址',
                tag: '转发模式',
                sort: '排序',
                status: '状态',
                css_id: '服务器ID',
                operate: '操作',
            },
            tagField,
            tagSelect,
        };
    },
    actions: {
        async getList() {
            const result = await List(this.page, this.pageSize);
            if (result.result) {
                this.list = result.data;
            }
        },
        async check() {
            let res = await Status();
            let getStatus = (id: number) => {
                let status = res[id];
                if (status == true) {
                    return true;
                }
                return false;
            };
            let list: ForwardRule[] = [];
            this.list.forEach((item) => {
                list.push({
                    ...item,
                    status: getStatus(item.id),
                });
            });
            this.list = list;
        },
    }
});