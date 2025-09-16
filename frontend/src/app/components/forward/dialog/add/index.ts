import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import type { CuForwardRule } from "@/app/pinia/forwar";
import { useForwardStore } from "@/app/pinia/forwar";
import { useSshServerStore } from "@/app/pinia/sshServer";
import { ElMessage } from "element-plus";
import { Add } from "@wails/go/bll/ForwardRuleBll";

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            setup() {
                const forwarStore = useForwardStore();
                const fieldName = computed(() => forwarStore.fieldName);
                let tagSelect = computed(() => forwarStore.tagSelect);
                const sshServerStore = useSshServerStore();
                let serverList = computed(() => sshServerStore.serverList);
                let visible = ref(true);
                let data = ref<CuForwardRule>({
                    name: "",
                    remote_addr: "",
                    local_addr: "",
                    tag: "L",
                    sort: 0,
                    css_id: 0,
                });

                return {
                    forwarStore,
                    fieldName,
                    tagSelect,
                    serverList,
                    visible,
                    data,
                };
            },
            created() {
                if (this.serverList.length == 0) {
                    ElMessage.error('请先添加服务器');
                    this.cancel();
                } else {
                    this.data.css_id = Number(this.serverList[0].value);
                }
            },
            methods: {
                cancel() {
                    this.$emit('change', {
                        visible: false,
                    });
                },
                async submit() {
                    if (this.data.name == "") {
                        ElMessage.error('请输入名称');
                        return;
                    }
                    if (this.data.remote_addr == "") {
                        ElMessage.error('请输入远程地址');
                        return;
                    }
                    if (this.data.local_addr == "") {
                        ElMessage.error('请输入本地地址');
                        return;
                    }

                    let result = await Add(this.data);
                    if (result.result) {
                        ElMessage.success('添加成功');
                        this.$emit('change', {
                            reload: true,
                            visible: false,
                        });
                    } else {
                        ElMessage.error(result.msg);
                    }
                },
            },
            components: {},
        });
        return vue;
    }
}

export default Component;
