import { SvelteComponentTyped } from "svelte";
import TreeViewItem from './TreeViewItem.svelte';
import type { TreeViewNode } from './types.js';
declare const __propDef: {
    props: {
        /** Provide data-driven nodes. */ nodes?: TreeViewNode[] | undefined;
        /**
             * provides id's of expanded nodes
             * @type {string[]}
             */ expandedNodes?: string[] | undefined;
        /**
             * provides id's of disabled nodes
             * @type {string[]}
             */ disabledNodes?: string[] | undefined;
        /**
             * provides id's of checked nodes
             * @type {string[]}
             */ checkedNodes?: string[] | undefined;
        /**
             * provides id's of indeterminate nodes
             * @type {string[]}
             */ indeterminateNodes?: string[] | undefined;
        treeItems?: TreeViewItem[] | undefined;
    };
    events: {
        click: CustomEvent<any>;
        toggle: CustomEvent<any>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export type RecursiveTreeViewItemProps = typeof __propDef.props;
export type RecursiveTreeViewItemEvents = typeof __propDef.events;
export type RecursiveTreeViewItemSlots = typeof __propDef.slots;
export default class RecursiveTreeViewItem extends SvelteComponentTyped<RecursiveTreeViewItemProps, RecursiveTreeViewItemEvents, RecursiveTreeViewItemSlots> {
}
export {};
