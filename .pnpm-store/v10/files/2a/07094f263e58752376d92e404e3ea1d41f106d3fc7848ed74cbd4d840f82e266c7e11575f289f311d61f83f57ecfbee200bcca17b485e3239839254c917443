import { SvelteComponentTyped } from "svelte";
import { fade } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type FadeTransition = typeof fade;
import type { StepperButton, StepperState } from './types.js';
declare class __sveltets_Render<TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition> {
    props(): {
        [x: string]: any;
        /** Provide classes to style the stepper header gap.*/
        gap?: string | undefined;
        /** Provide the verbiage that represents "Step".*/
        stepTerm?: string | undefined;
        /** Provide classes to style the stepper header badges.*/
        badge?: string | undefined;
        /** Provide classes to style the stepper header active step badge.*/
        active?: string | undefined;
        /** Provide classes to style the stepper header border.*/
        border?: string | undefined;
        /** Provide the initially selected step*/
        start?: number | undefined;
        /** Set the justification for the step navigation buttons.*/
        justify?: string | undefined;
        /** Provide arbitrary classes to style the back button.*/
        buttonBack?: string | undefined;
        /** Set the type of the back button.*/
        buttonBackType?: StepperButton | undefined;
        /** Provide the HTML label content for the back button.*/
        buttonBackLabel?: string | undefined;
        /** Provide arbitrary classes to style the next button.*/
        buttonNext?: string | undefined;
        /** Set the type of the next button.*/
        buttonNextType?: StepperButton | undefined;
        /** Provide the HTML label content for the next button.*/
        buttonNextLabel?: string | undefined;
        /** Provide arbitrary classes to style the complete button.*/
        buttonComplete?: string | undefined;
        /** Set the type of the complete button.*/
        buttonCompleteType?: StepperButton | undefined;
        /** Provide the HTML label content for the complete button.*/
        buttonCompleteLabel?: string | undefined;
        /** Provide arbitrary classes to the stepper header region.*/
        regionHeader?: string | undefined;
        /** Provide arbitrary classes to the stepper content region.*/
        regionContent?: string | undefined;
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
    events(): {
        next: CustomEvent<{
            step: number;
            state: StepperState;
        }>;
        step: CustomEvent<{
            step: number;
            state: StepperState;
        }>;
        back: CustomEvent<{
            step: number;
            state: StepperState;
        }>;
        complete: CustomEvent<{
            step: number;
            state: StepperState;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        default: {};
    };
}
export type StepperProps<TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['props']>;
export type StepperEvents<TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['events']>;
export type StepperSlots<TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition> = ReturnType<__sveltets_Render<TransitionIn, TransitionOut>['slots']>;
export default class Stepper<TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition> extends SvelteComponentTyped<StepperProps<TransitionIn, TransitionOut>, StepperEvents<TransitionIn, TransitionOut>, StepperSlots<TransitionIn, TransitionOut>> {
}
export {};
