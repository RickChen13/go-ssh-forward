import mitt from 'mitt';
import { onBeforeUnmount } from 'vue';

type Event = {
    back: void;
}
type Option<K extends keyof Event> = {
    name: K; // 对应 Event 的键
    callback: Fn<Event[K]>; // 对应 Event 的值作为回调函数的参数类型
};

const EventBus = mitt<Event>();
export const useEventBus = <K extends keyof Event>(option?: Option<K>) => {
    if (option) {
        EventBus.on(option.name, option.callback);
        onBeforeUnmount(() => {
            EventBus.off(option.name, option.callback);
        });
    }

    return {
        on: EventBus.on,
        off: EventBus.off,
        emit: EventBus.emit,
        all: EventBus.all
    };
};