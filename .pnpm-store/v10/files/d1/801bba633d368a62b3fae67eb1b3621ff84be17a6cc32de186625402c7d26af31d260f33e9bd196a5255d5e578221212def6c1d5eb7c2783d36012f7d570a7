import { SvelteComponentTyped } from "svelte";
import { slide } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type SlideTransition = typeof slide;
declare class __sveltets_Render<TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> {
    props(): {
        [x: string]: any;
        /** Set the auto-collapse mode.*/
        autocollapse?: boolean | undefined;
        /** Provide classes to set the accordion width.*/
        width?: string | undefined;
        /** Provide classes to set the vertical spacing between items.*/
        spacing?: string | undefined;
        /** Set the accordion disabled state for all items.*/
        disabled?: boolean | undefined;
        /** Provide classes to set the accordion item padding styles.*/
        padding?: string | undefined;
        /** Provide classes to set the accordion item hover styles.*/
        hover?: string | undefined;
        /** Provide classes to set the accordion item rounded styles.*/
        rounded?: string | undefined;
        /** Set the rotation of the item caret in the open state.*/
        caretOpen?: string | undefined;
        /** Set the rotation of the item caret in the closed state.*/
        caretClosed?: string | undefined;
        /** Provide arbitrary classes to the trigger button region.*/
        regionControl?: string | undefined;
        /** Provide arbitrary classes to the content panel region.*/
        regionPanel?: string | undefined;
        /** Provide arbitrary classes to the caret icon region.*/
        regionCaret?: string | undefined;
        /** Enable/Disable transitions*/
        transitions?: boolean | undefined;
        /** Provide the transition to used on entry.*/
        transitionIn?: TransitionIn | undefined;
        /** Transition params provided to `transitionIn`.*/
        transitionInParams?: TransitionParams<TransitionIn> | undefined;
        /** Provide the transition to used on exit.*/
        transitionOut?: TransitionOut | undefined;
        /** Transition params provided to `transitionOut`.*/
        transitionOutParams?: TransitionParams<TransitionOut> | undefined;
    };
    events(): {} & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        default: {};
    };
}
export type AccordionProps<TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['props']>;
export type AccordionEvents<TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['events']>;
export type AccordionSlots<TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['slots']>;
/** The Accordion parent element. */
export default class Accordion<TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> extends SvelteComponentTyped<AccordionProps<TransitionIn, TransitionOut>, AccordionEvents<TransitionIn, TransitionOut>, AccordionSlots<TransitionIn, TransitionOut>> {
}
export {};
