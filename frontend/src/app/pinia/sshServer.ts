import { defineStore } from 'pinia';
import { List } from '@wails/go/bll/SshServerBll';

export type CuSshServerData = {
    name: string;
    host: string;
    user: string;
    pass: string;
    key_path: string;
    pass_phrase: string;
    sort: number;
}
export type SshServerData = CuSshServerData & {
    id: number;
}
export type SshServerEditData = CuSshServerData & {
    id: number;
    pass_new: string;
    pass_phrase_new: string;
    clear_pass: boolean;
    clear_pass_phrase: boolean;
}


type ForwarStore = {
    page: number;
    pageSize: number;
    count: number;
    list: SshServerData[];
    serverList: ElSelectType[];
    fieldName: FieldNameMap;
}

export const useSshServerStore = defineStore('sshServer', {
    state: (): ForwarStore => {
        return {
            page: 1,
            pageSize: 1000,
            count: 0,
            list: [],
            serverList: [],
            fieldName: {
                id: 'ID',
                name: '名称',
                host: '主机',
                user: '用户',
                pass: '密码',
                key_path: '密钥路径',
                pass_phrase: '密钥密码',
                sort: '排序',
            }
        };
    },
    actions: {
        async getList() {
            let result = await List(this.page, this.pageSize);
            if (result.result) {
                let list = result.data as SshServerData[];
                let serverList = list.map((item) => {
                    return {
                        label: item.name,
                        value: item.id,
                    };
                });
                this.serverList = serverList;
                this.list = list;
            }

        },
    }
});