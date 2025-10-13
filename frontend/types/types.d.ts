type Fn<T = any> = (data: T) => void;

type ElSelectType = {
    label: string;
    value: string | number;
};

type FieldNameMap = {
    [key: string]: string;
};

declare interface Window {
    wailsApi: {
        log: (data: string) => void;
    };
}