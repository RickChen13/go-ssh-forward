import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import type { CuSshServerData } from "@/app/pinia/sshServer";
import { useSshServerStore } from "@/app/pinia/sshServer";
import { ElMessage } from "element-plus";
import { Add } from "@wails/go/bll/SshServerBll";

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            setup() {
                const sshServerStore = useSshServerStore();
                const fieldName = computed(() => sshServerStore.fieldName);

                let visible = ref(true);
                let data = ref<CuSshServerData>({
                    name: '',
                    host: '',
                    user: '',
                    pass: '',
                    key_path: '',
                    pass_phrase: '',
                    sort: 0,
                });

                return {
                    sshServerStore,
                    fieldName,
                    visible,
                    data,
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
                    if (this.data.name === '') {
                        ElMessage.error('请输入名称');
                        return;
                    }
                    if (this.data.host === '') {
                        ElMessage.error('请输入主机');
                        return;
                    }
                    if (this.data.user === '') {
                        ElMessage.error('请输入用户名');
                        return;
                    }
                    if (this.data.pass === '' && this.data.key_path === '') {
                        ElMessage.error('请输入密码或密钥');
                        return;
                    }
                    this.data.sort = Number(this.data.sort);
                    let res = await Add(this.data);
                    if (res.result) {
                        ElMessage.success('添加成功');
                        this.$emit('change', {
                            reload: true,
                            visible: false,
                        });
                    } else {
                        ElMessage.error(res.msg);
                    }

                }
            },
            components: {},
        });
        return vue;
    }
}

export default Component;
