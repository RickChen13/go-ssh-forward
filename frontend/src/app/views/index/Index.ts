import BaseViews from '@/fast/base/base.view';
import { defineComponent, ref } from 'vue';
import ForwardVue from "@/app/components/forward/index.vue";
import SshServerVue from "@/app/components/sshServer/index.vue";


class Component extends BaseViews {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            setup() {
                const activeIndex = ref('1');
                const handleSelect = (key: string) => {
                    activeIndex.value = key;
                };
                return {
                    activeIndex,
                    handleSelect,
                };
            },
            components: {
                ForwardVue,
                SshServerVue,
            },
        });
        return vue;
    }
}

export default Component;
