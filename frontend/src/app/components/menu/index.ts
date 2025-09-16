import BaseView from "@/fast/base/base.view";
import { defineComponent, ref } from "vue";
import {
    WindowHide,
    WindowMinimise,
    WindowIsMaximised, WindowToggleMaximise,
} from '@wails/runtime/runtime';
import ClarityWindowCloseLine from "@/app/components/icon/ClarityWindowCloseLine.vue";
import ClarityWindowMaxLine from "@/app/components/icon/ClarityWindowMaxLine.vue";
import ClarityWindowMinLine from "@/app/components/icon/ClarityWindowMinLine.vue";
import ClarityWindowRestoreLine from "@/app/components/icon/ClarityWindowRestoreLine.vue";

class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            components: {
                ClarityWindowCloseLine,
                ClarityWindowMaxLine,
                ClarityWindowMinLine,
                ClarityWindowRestoreLine,
            },
            setup() {
                const isMaximised = ref<Boolean>(false);
                const SetIsMaximised = async () => {
                    isMaximised.value = await WindowIsMaximised();
                };
                SetIsMaximised();

                const ToggleMaximise = async () => {
                    WindowToggleMaximise();
                    await SetIsMaximised();
                };
                return {
                    isMaximised,
                    WindowHide,
                    WindowMinimise,
                    ToggleMaximise,
                };
            },
            created() { },
            methods: {},

        });
        return vue;
    }
}

export default Component;
