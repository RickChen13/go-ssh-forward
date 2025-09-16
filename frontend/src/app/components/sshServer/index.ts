import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import { Plus, Refresh } from '@element-plus/icons-vue';
import AddVue from './dialog/add/index.vue';
import EditVue from './dialog/edit/index.vue';
import type { SshServerData } from "@/app/pinia/sshServer";
import { useSshServerStore } from "@/app/pinia/sshServer";
import { ElMessage, ElMessageBox } from "element-plus";
import { CssIdCount } from "@wails/go/bll/ForwardRuleBll";
import { Delete } from "@wails/go/bll/SshServerBll";
type Change = {
    visible?: boolean;
    reload?: boolean;
}

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            setup() {
                let _any: any = null;

                const sshServerStore = useSshServerStore();
                let list = computed(() => sshServerStore.list);
                let editData = ref<SshServerData>(_any);
                const getList = async () => {
                    await sshServerStore.getList();
                };

                getList();
                return {
                    list,
                    editData,
                    getList,

                    Plus,
                    Refresh,
                };

            },
            data() {
                return {
                    addVisible: false,
                    editVisible: false,
                };
            },
            created() { },
            methods: {
                change(config: Change) {
                    if (config.reload != undefined) {
                        this.getList();
                    }
                    if (config.visible != undefined) {
                        this.addVisible = config.visible;
                        this.editVisible = config.visible;
                    }
                },
                editRow(data: SshServerData) {
                    this.editData = data;
                    this.editVisible = true;
                },
                async delRow(id: number) {
                    ElMessageBox.confirm(
                        '是否删除该配置?',
                        '提示',
                        {
                            confirmButtonText: '确定',
                            cancelButtonText: '取消',
                            type: 'warning',
                        }
                    ).then(async () => {
                        let countRes = await CssIdCount(id);
                        if (countRes.result && countRes.data > 0) {
                            ElMessage.error('该配置被使用中，无法删除');
                            return;
                        }
                        let delRes = await Delete(id);
                        if (delRes.result) {
                            ElMessage.success('删除成功');
                            this.getList();
                        } else {
                            ElMessage.error('删除失败');
                        }
                    }).catch(() => { });
                }
            },
            components: {
                AddVue,
                EditVue,
            },
        });
        return vue;
    }
}

export default Component;
