import { SvelteComponentTyped } from "svelte";
import { fly, scale } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type FlyTransition = typeof fly;
type ScaleTransition = typeof scale;
declare class __sveltets_Render<ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition> {
    props(): {
        [x: string]: any;
        /** Bind the input value.*/
        input?: string | undefined;
        /** Set a unique select input name.*/
        name: string;
        /** An array of values.*/
        value?: any[] | undefined;
        /** Provide a whitelist of accepted values.*/
        whitelist?: string[] | undefined;
        /** Maximum number of chips. Set -1 to disable.*/
        max?: number | undefined;
        /** Set the minimum character length.*/
        minlength?: number | undefined;
        /** Set the maximum character length.*/
        maxlength?: number | undefined;
        /** When enabled, allows for uppercase values.*/
        allowUpperCase?: boolean | undefined;
        /** When enabled, allows for duplicate values.*/
        allowDuplicates?: boolean | undefined;
        /** Provide a custom validator function.*/
        validation?: ((...args: any[]) => boolean) | undefined;
        /** The duration of the flip (first, last, invert, play) animation.*/
        duration?: number | undefined;
        /** Set the required state for this input field.*/
        required?: boolean | undefined;
        /** Provide classes or a variant to style the chips.*/
        chips?: string | undefined;
        /** {{ event: Event, input: any  }} invalid - Fires when the input value is invalid.*/
        invalid?: string | undefined;
        /** Provide classes to set padding styles.*/
        padding?: string | undefined;
        /** Provide classes to set border radius styles.*/
        rounded?: string | undefined;
        /** Provide arbitrary classes to style the chip wrapper region.*/
        regionChipWrapper?: string | undefined;
        /** Provide arbitrary classes to style the chip list region.*/
        regionChipList?: string | undefined;
        /** Provide arbitrary classes to style the input field region.*/
        regionInput?: string | undefined;
        /** Provide the ARIA label for the select input.*/
        label?: string | undefined;
        /** Enable/Disable transitions*/
        transitions?: boolean | undefined;
        /** Provide the transition used in list on entry.*/
        listTransitionIn?: ListTransitionIn | undefined;
        /** Transition params provided to `ListTransitionIn`.*/
        listTransitionInParams?: TransitionParams<ListTransitionIn> | undefined;
        /** Provide the transition used in list on exit.*/
        listTransitionOut?: ListTransitionOut | undefined;
        /** Transition params provided to `ListTransitionOut`.*/
        listTransitionOutParams?: TransitionParams<ListTransitionOut> | undefined;
        /** Provide the transition used in chip on entry.*/
        chipTransitionIn?: ChipTransitionIn | undefined;
        /** Transition params provided to `ChipTransitionIn`.*/
        chipTransitionInParams?: TransitionParams<ChipTransitionIn> | undefined;
        /** Provide the transition used in chip on exit.*/
        chipTransitionOut?: ChipTransitionOut | undefined;
        /** Transition params provided to `ChipTransitionOut`.*/
        chipTransitionOutParams?: TransitionParams<ChipTransitionOut> | undefined;
        addChip?: ((chip: string) => void) | undefined;
        removeChip?: ((chip: string) => void) | undefined;
    };
    events(): {
        input: Event;
        focus: FocusEvent;
        blur: FocusEvent;
        click: MouseEvent;
        keypress: KeyboardEvent;
        keydown: KeyboardEvent;
        keyup: KeyboardEvent;
        add: CustomEvent<{
            event: KeyboardEvent;
            chipIndex: number;
            chipValue: string;
        }>;
        remove: CustomEvent<{
            event: MouseEvent;
            chipIndex: number;
            chipValue: string;
        }>;
        addManually: CustomEvent<{
            chipIndex: number;
            chipValue: string;
        }>;
        removeManually: CustomEvent<{
            chipValue: string;
        }>;
        invalid: CustomEvent<{
            event: KeyboardEvent;
            input: string;
        }>;
        invalidManually: CustomEvent<{
            input: string;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {};
}
export type InputChipProps<ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition> = ReturnType<__sveltets_Render<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>['props']>;
export type InputChipEvents<ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition> = ReturnType<__sveltets_Render<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>['events']>;
export type InputChipSlots<ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition> = ReturnType<__sveltets_Render<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>['slots']>;
export default class InputChip<ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition> extends SvelteComponentTyped<InputChipProps<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>, InputChipEvents<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>, InputChipSlots<ListTransitionIn, ListTransitionOut, ChipTransitionIn, ChipTransitionOut>> {
    get addChip(): (chip: string) => void;
    get removeChip(): (chip: string) => void;
}
export {};
