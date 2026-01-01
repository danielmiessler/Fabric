import { SvelteComponentTyped } from "svelte";
import { fly } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type FlyTransition = typeof fly;
import type { ModalComponent } from './types.js';
declare class __sveltets_Render<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> {
    props(): {
        [x: string]: any;
        /** Register a list of reusable component modals.*/
        components?: Record<string, ModalComponent> | undefined;
        /** Set the modal position within the backdrop container*/
        position?: string | undefined;
        /** Provide classes to style the modal background.*/
        background?: string | undefined;
        /** Provide classes to style the modal width.*/
        width?: string | undefined;
        /** Provide classes to style the modal height.*/
        height?: string | undefined;
        /** Provide classes to style the modal padding.*/
        padding?: string | undefined;
        /** Provide classes to style the modal spacing.*/
        spacing?: string | undefined;
        /** Provide classes to style the modal border radius.*/
        rounded?: string | undefined;
        /** Provide classes to style modal box shadow.*/
        shadow?: string | undefined;
        /** Provide a class to override the z-index*/
        zIndex?: string | undefined;
        /** Provide classes for neutral buttons, such as Cancel.*/
        buttonNeutral?: string | undefined;
        /** Provide classes for positive actions, such as Confirm or Submit.*/
        buttonPositive?: string | undefined;
        /** Override the text for the Cancel button.*/
        buttonTextCancel?: string | undefined;
        /** Override the text for the Confirm button.*/
        buttonTextConfirm?: string | undefined;
        /** Override the text for the Submit button.*/
        buttonTextSubmit?: string | undefined;
        /** Provide arbitrary classes to the backdrop region.*/
        regionBackdrop?: string | undefined;
        /** Provide arbitrary classes to modal header region.*/
        regionHeader?: string | undefined;
        /** Provide arbitrary classes to modal body region.*/
        regionBody?: string | undefined;
        /** Provide arbitrary classes to modal footer region.*/
        regionFooter?: string | undefined;
        /** Enable/Disable transitions*/
        transitions?: boolean | undefined;
        /** Provide the transition used on entry.*/
        transitionIn?: TransitionIn | undefined;
        /** Transition params provided to `TransitionIn`.*/
        transitionInParams?: TransitionParams<TransitionIn> | undefined;
        /** Provide the transition used on exit.*/
        transitionOut?: TransitionOut | undefined;
        /** Transition params provided to `TransitionOut`.*/
        transitionOutParams?: TransitionParams<TransitionOut> | undefined;
    };
    events(): {
        touchstart: TouchEvent;
        touchend: TouchEvent;
        backdrop: CustomEvent<MouseEvent>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {};
}
export type ModalProps<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['props']>;
export type ModalEvents<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['events']>;
export type ModalSlots<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['slots']>;
export default class Modal<TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition> extends SvelteComponentTyped<ModalProps<TransitionIn, TransitionOut>, ModalEvents<TransitionIn, TransitionOut>, ModalSlots<TransitionIn, TransitionOut>> {
}
export {};
