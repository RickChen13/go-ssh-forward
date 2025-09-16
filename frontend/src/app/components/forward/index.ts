import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import { Plus, Refresh } from '@element-plus/icons-vue';
import { Stop, RunById } from '@wails/go/bll/forwardBll';
import AddVue from './dialog/add/index.vue';
import EditVue from './dialog/edit/index.vue';
import type { ForwardRule } from "@/app/pinia/forwar";
import { useForwardStore } from "@/app/pinia/forwar";
import { Delete } from "@wails/go/bll/ForwardRuleBll";
import { ElMessage, ElMessageBox } from "element-plus";

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

                const forwarStore = useForwardStore();
                let list = computed(() => forwarStore.list);
                let tagField = computed(() => forwarStore.tagField);
                const getList = async () => {
                    forwarStore.getList();
                };
                let editData = ref<ForwardRule>(_any);
                let fieldName = computed(() => {
                    return forwarStore.fieldName;
                });
                getList();
                return {
                    forwarStore,
                    list,
                    tagField,
                    getList,
                    editData,
                    fieldName,
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
            created() {
                let check = () => {
                    this.forwarStore.check();
                    setTimeout(() => {
                        check();
                    }, 1000);
                };
                check();
            },
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
                editRow(data: ForwardRule) {
                    this.editData = data;
                    this.editVisible = true;
                },
                delRow(data: ForwardRule) {
                    ElMessageBox.confirm(
                        '是否删除该配置?',
                        '提示',
                        {
                            confirmButtonText: '确定',
                            cancelButtonText: '取消',
                            type: 'warning',
                        }
                    ).then(async () => {
                        if (data.status == true) {
                            ElMessage.error('请先关闭该配置');
                            return;
                        }
                        let delRes = await Delete(data.id);
                        if (delRes.result) {
                            ElMessage.success('删除成功');
                            this.getList();
                        } else {
                            ElMessage.error('删除失败');
                        }
                    }).catch(() => { });
                },

                close(id: number) {
                    Stop(String(id));
                },
                RunById(id: number) {
                    RunById(id);
                },

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
