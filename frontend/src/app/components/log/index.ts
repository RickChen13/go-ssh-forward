import { useEventBus } from "@/app/plugins/mitt/EventBus";
import Time from "@/common/Time";
import BaseView from "@/fast/base/base.view";
import { defineComponent, nextTick, reactive } from "vue";
import {
    Delete,
} from '@element-plus/icons-vue';
class Component extends BaseView {
    constructor() {
        super();
    }

    public vue() {
        const vue = defineComponent({
            setup() {
                let logData = reactive<String[]>([]);
                return {
                    logData,
                    Delete,
                };
            },

            created() {
                useEventBus().on("log", (data: string) => {
                    this.logData.push(`<span class="time">${Time.getYmdHis()}</span> ${data}`);
                    nextTick(() => {
                        document.querySelector(`#log-${this.logData.length - 1}`)?.scrollIntoView(false);
                    });
                });
            },
            methods: {
                clear() {
                    this.logData.length = 0;
                },
            },

        });
        return vue;
    }
}

export default Component;
