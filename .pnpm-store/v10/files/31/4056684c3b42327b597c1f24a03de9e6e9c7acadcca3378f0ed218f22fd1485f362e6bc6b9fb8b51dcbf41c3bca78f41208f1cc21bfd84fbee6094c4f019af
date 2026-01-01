import { SvelteComponentTyped } from "svelte";
import type { TableSource } from '../../index.js';
declare const __propDef: {
    props: {
        [x: string]: any;
        /** Provide the full set of table source data.*/
        source: TableSource;
        /** Enables row hover style and `on:selected` event when rows are clicked.*/
        interactive?: boolean | undefined;
        /** Override the Tailwind Element class. Replace this for a headless UI.*/
        element?: string | undefined;
        /** Provide classes to set the table text size.*/
        text?: string | undefined;
        /** Provide classes to set the table text color.*/
        color?: string | undefined;
        /** Provide arbitrary classes for the table head.*/
        regionHead?: string | undefined;
        /** Provide arbitrary classes for the table head cells.*/
        regionHeadCell?: string | undefined;
        /** Provide arbitrary classes for the table body.*/
        regionBody?: string | undefined;
        /** Provide arbitrary classes for the table cells.*/
        regionCell?: string | undefined;
        /** Provide arbitrary classes for the table foot.*/
        regionFoot?: string | undefined;
        /** Provide arbitrary classes for the table foot cells.*/
        regionFootCell?: string | undefined;
    };
    events: {
        /** {rowMetaData} selected - Fires when a table row is clicked.*/
        selected: CustomEvent<string[]>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export type TableProps = typeof __propDef.props;
export type TableEvents = typeof __propDef.events;
export type TableSlots = typeof __propDef.slots;
export default class Table extends SvelteComponentTyped<TableProps, TableEvents, TableSlots> {
}
export {};
