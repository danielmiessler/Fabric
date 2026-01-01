import { SvelteComponentTyped } from "svelte";
import type { Writable } from 'svelte/store';
import type { StepperState } from './types.js';
import type { Transition, TransitionParams } from '../../index.js';
declare class __sveltets_Render<TransitionIn extends Transition, TransitionOut extends Transition> {
    props(): {
        [x: string]: any;
        locked?: boolean | undefined;
        /** Provide arbitrary classes to the step header region.*/
        regionHeader?: string | undefined;
        /** Provide arbitrary classes to the step content region.*/
        regionContent?: string | undefined;
        /** Provide arbitrary classes to the step navigation region.*/
        regionNavigation?: string | undefined;
        state?: Writable<StepperState> | undefined;
        stepTerm?: string | undefined;
        gap?: string | undefined;
        justify?: string | undefined;
        onNext?: ((locked: boolean, stepIndex: number) => void) | undefined;
        onBack?: ((stepIndex: number) => void) | undefined;
        onComplete?: ((stepIndex: number) => void) | undefined;
        buttonBack?: string | undefined;
        buttonBackType?: "button" | "reset" | "submit" | undefined;
        buttonBackLabel?: string | undefined;
        buttonNext?: string | undefined;
        buttonNextType?: "button" | "reset" | "submit" | undefined;
        buttonNextLabel?: string | undefined;
        buttonComplete?: string | undefined;
        buttonCompleteType?: "button" | "reset" | "submit" | undefined;
        buttonCompleteLabel?: string | undefined;
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
        header: {};
        default: {};
        navigation: {};
    };
}
export type StepProps<TransitionIn extends Transition, TransitionOut extends Transition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['props']>;
export type StepEvents<TransitionIn extends Transition, TransitionOut extends Transition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['events']>;
export type StepSlots<TransitionIn extends Transition, TransitionOut extends Transition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['slots']>;
export default class Step<TransitionIn extends Transition, TransitionOut extends Transition> extends SvelteComponentTyped<StepProps<TransitionIn, TransitionOut>, StepEvents<TransitionIn, TransitionOut>, StepSlots<TransitionIn, TransitionOut>> {
}
export {};
