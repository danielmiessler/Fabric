import type { EventDispatcher } from 'svelte';
export interface StepperState {
    current: number;
    total: number;
}
export type StepperButton = 'submit' | 'reset' | 'button';
export type StepperEvent = {
    next: {
        step: number;
        state: StepperState;
    };
    step: {
        step: number;
        state: StepperState;
    };
    back: {
        step: number;
        state: StepperState;
    };
    complete: {
        step: number;
        state: StepperState;
    };
};
export type StepperEventDispatcher = EventDispatcher<StepperEvent>;
