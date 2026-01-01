import { SvelteComponentTyped } from "svelte";
import { fly } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type FlyTransition = typeof fly;
declare class __sveltets_Render<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> {
    props(): {
        [x: string]: any;
        /** Set the toast position.*/
        position?: "b" | "br" | "tr" | "t" | "l" | "r" | "tl" | "bl" | undefined;
        /** Maximum toasts that can show at once.*/
        max?: number | undefined;
        /** Provide classes to set the background color.*/
        background?: string | undefined;
        /** Provide classes to set width styles.*/
        width?: string | undefined;
        /** Provide classes to set the text color.*/
        color?: string | undefined;
        /** Provide classes to set the padding.*/
        padding?: string | undefined;
        /** Provide classes to set toast horizontal spacing.*/
        spacing?: string | undefined;
        /** Provide classes to set the border radius styles.*/
        rounded?: string | undefined;
        /** Provide classes to set the border box shadow.*/
        shadow?: string | undefined;
        /** Provide a class to override the z-index*/
        zIndex?: string | undefined;
        /** Provide styles for the action button.*/
        buttonAction?: string | undefined;
        /** Provide styles for the dismiss button.*/
        buttonDismiss?: string | undefined;
        /** The button label text.*/
        buttonDismissLabel?: string | undefined;
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
    slots(): {};
}
export type ToastProps<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['props']>;
export type ToastEvents<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['events']>;
export type ToastSlots<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['slots']>;
export default class Toast<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> extends SvelteComponentTyped<ToastProps<TransitionIn, TransitionOut>, ToastEvents<TransitionIn, TransitionOut>, ToastSlots<TransitionIn, TransitionOut>> {
}
export {};
