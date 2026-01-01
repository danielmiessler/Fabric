import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        /** Current rating value.*/
        value?: number | undefined;
        /** Maximum rating value.*/
        max?: number | undefined;
        /** Enables interactive mode for each rating icon.*/
        interactive?: boolean | undefined;
        /** Provide classes to set the text color.*/
        text?: string | undefined;
        /** Provide classes to set the SVG fill color.*/
        fill?: string | undefined;
        /** Provide classes to set the flexbox justification.*/
        justify?: string | undefined;
        /** Provide classes to set the horizontal spacing style.*/
        spacing?: string | undefined;
        /** Provide arbitrary classes to the icon region.*/
        regionIcon?: string | undefined;
    };
    events: {
        /** {{ index: number  }} icon - Fires when an icons is clicked*/
        icon: CustomEvent<{
            index: number;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        full: {};
        half: {};
        empty: {};
    };
};
export type RatingsProps = typeof __propDef.props;
export type RatingsEvents = typeof __propDef.events;
export type RatingsSlots = typeof __propDef.slots;
export default class Ratings extends SvelteComponentTyped<RatingsProps, RatingsEvents, RatingsSlots> {
}
export {};
