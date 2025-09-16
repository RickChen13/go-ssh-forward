import BaseView from "@/fast/base/base.view";
import { computed, defineComponent, ref } from "vue";
import lodash from 'lodash';
import type { SshServerData, SshServerEditData } from "@/app/pinia/sshServer";
import { sshServer } from "@wails/go/models";
import { useSshServerStore } from "@/app/pinia/sshServer";
import { ElMessage } from "element-plus";
import { Update } from "@wails/go/bll/SshServerBll";

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            props: {
                data: {
                    type: Object as () => SshServerData,
                    required: true,
                },
            },
            setup() {
                const sshServerStore = useSshServerStore();
                const fieldName = computed(() => sshServerStore.fieldName);
                let visible = ref(true);
                return {
                    fieldName,
                    visible,
                };
            },
            data() {
                let localData: SshServerEditData = {
                    ...lodash.cloneDeep(this.data),
                    pass_new: '',
                    pass_phrase_new: '',
                    clear_pass: false,
                    clear_pass_phrase: false,
                };

                return {
                    localData,
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
                    if (this.localData.name === '') {
                        ElMessage.error('请输入名称');
                        return;
                    }
                    if (this.localData.host === '') {
                        ElMessage.error('请输入主机');
                        return;
                    }
                    if (this.localData.user === '') {
                        ElMessage.error('请输入用户名');
                        return;
                    }

                    let updateData: sshServer.ConfigSshServerU = {
                        name: this.localData.name,
                        host: this.localData.host,
                        user: this.localData.user,
                        key_path: this.localData.key_path,
                        sort: Number(this.localData.sort),
                        pass: "",
                        pass_phrase: "",
                        clear_pass: this.localData.clear_pass,
                        clear_pass_phrase: this.localData.clear_pass_phrase,
                    };
                    if (this.localData.pass_new !== '') {
                        updateData.pass = this.localData.pass_new;
                    }
                    if (this.localData.pass_phrase_new !== '') {
                        updateData.pass_phrase = this.localData.pass_phrase_new;
                    }
                    let result = await Update(this.data.id, updateData);
                    if (result.result) {
                        ElMessage.success('更新成功');
                        this.$emit('change', {
                            reload: true,
                            visible: false,
                        });
                    } else {
                        ElMessage.error('更新失败');
                    }
                },
            },
            components: {},
        });
        return vue;
    }
}

export default Component;
