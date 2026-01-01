import type { ActionReturn } from 'svelte/action';
type ClipboardArgs = string | {
    element: string;
} | {
    input: string;
};
type ClipboardAttributes = {
    'on:copyComplete'?: () => void;
};
export declare function clipboard(node: HTMLElement, args: ClipboardArgs): ActionReturn<ClipboardArgs, ClipboardAttributes>;
export {};
