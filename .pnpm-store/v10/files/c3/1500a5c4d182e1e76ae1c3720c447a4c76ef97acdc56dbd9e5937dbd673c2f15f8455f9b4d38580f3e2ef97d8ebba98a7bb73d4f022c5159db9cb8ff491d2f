import { type Writable } from 'svelte/store';
interface Serializer<T> {
    parse(text: string): T;
    stringify(object: T): string;
}
type StorageType = 'local' | 'session';
interface Options<T> {
    serializer?: Serializer<T>;
    storage?: StorageType;
}
export declare function localStorageStore<T>(key: string, initialValue: T, options?: Options<T>): Writable<T>;
export {};
