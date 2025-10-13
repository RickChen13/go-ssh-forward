import BaseViews from '@/fast/base/base.view';
import { defineComponent, ref } from 'vue';
import ForwardVue from "@/app/components/forward/index.vue";
import SshServerVue from "@/app/components/sshServer/index.vue";
import LogVue from "@/app/components/log/index.vue";
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

                const containerRef = ref();
                const splitterRef = ref();
                let hegiht = 400;
                const topPaneHeight = ref(`${hegiht}px`); // 初始高度
                const topPaneHeightNum = ref(hegiht - 102); // 初始高度
                const isResizing = ref(false);
                const MIN_HEIGHT = 200; // 最小高度限制 (px)
                return {
                    activeIndex,
                    handleSelect,

                    containerRef,
                    splitterRef,
                    topPaneHeight,
                    topPaneHeightNum,
                    isResizing,
                    MIN_HEIGHT,
                };
            },
            mounted() {
            },
            beforeUnmount() {
                document.removeEventListener('mousemove', this.handleMouseMove);
                document.removeEventListener('mouseup', this.stopResize);
            },
            methods: {
                startResize(e: { preventDefault: () => void; }) {
                    e.preventDefault();
                    this.isResizing = true;

                    // 在 document 上添加全局事件监听器
                    document.addEventListener('mousemove', this.handleMouseMove);
                    document.addEventListener('mouseup', this.stopResize);
                },
                handleMouseMove(e: { clientY: number; }) {
                    if (!this.isResizing) return;

                    const container = this.containerRef;
                    if (!container) return;

                    const containerRect = container.getBoundingClientRect();
                    const mouseY = e.clientY - containerRect.top;

                    const newHeight = mouseY;
                    const containerHeight = containerRect.height;

                    // 限制拖动范围
                    if (newHeight > this.MIN_HEIGHT && newHeight < containerHeight - this.MIN_HEIGHT) {
                        // 设置顶部面板的新高度 (使用像素值)
                        this.topPaneHeight = `${newHeight}px`;
                        this.topPaneHeightNum = newHeight - 102;
                    }
                },
                stopResize() {
                    this.isResizing = false;

                    // 移除全局事件监听器
                    document.removeEventListener('mousemove', this.handleMouseMove);
                    document.removeEventListener('mouseup', this.stopResize);
                },
            },
            components: {
                ForwardVue,
                SshServerVue,
                LogVue,
            },
        });
        return vue;
    }
}

export default Component;
