import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import lodash from 'lodash';
import type { ForwardRule } from "@/app/pinia/forwar";
import { useForwardStore } from "@/app/pinia/forwar";
import { useSshServerStore } from "@/app/pinia/sshServer";
import { ElMessage } from "element-plus";
import { Update } from "@wails/go/bll/ForwardRuleBll";
import { forwardRule } from "@wails/go/models";

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            props: {
                data: {
                    type: Object as () => ForwardRule,
                    required: true,
                },
            },
            setup() {
                const forwarStore = useForwardStore();
                const fieldName = computed(() => forwarStore.fieldName);
                let tagSelect = computed(() => forwarStore.tagSelect);
                const sshServerStore = useSshServerStore();
                let serverList = computed(() => sshServerStore.serverList);
                let visible = ref(true);
                return {
                    forwarStore,
                    fieldName,
                    tagSelect,
                    serverList,
                    visible,
                };
            },
            data() {
                return {
                    localData: lodash.cloneDeep(this.data),
                };
            },
            created() { },
            methods: {
                cancel() {
                    this.$emit('change', {
                        visible: false,
                    });
                },
                async submit() {
                    if (this.localData.name == "") {
                        ElMessage.error('请输入名称');
                        return;
                    }
                    if (this.localData.remote_addr == "") {
                        ElMessage.error('请输入远程地址');
                        return;
                    }
                    if (this.localData.local_addr == "") {
                        ElMessage.error('请输入本地地址');
                        return;
                    }
                    let updateData: forwardRule.ConfigForwardRuleUpdateData = {
                        name: this.localData.name,
                        remote_addr: this.localData.remote_addr,
                        local_addr: this.localData.local_addr,
                        tag: this.localData.tag,
                        sort: this.localData.sort,
                        css_id: this.localData.css_id,
                    };
                    let result = await Update(this.data.id, updateData);
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
