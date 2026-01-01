import { SvelteComponentTyped } from "svelte";
import { slide } from 'svelte/transition';
import { type Transition, type TransitionParams } from '../../index.js';
type SlideTransition = typeof slide;
import type { AutocompleteOption } from './types.js';
declare class __sveltets_Render<Value = unknown, Meta = unknown, TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> {
    props(): {
        [x: string]: any;
        /** Bind the input value.*/
        input?: unknown;
        /** Define values for the list.*/
        options?: AutocompleteOption<Value, Meta>[] | undefined;
        /** Limit the total number of suggestions.*/
        limit?: number | undefined;
        /** Provide allowlist values.*/
        allowlist?: Value[] | undefined;
        /** Provide denylist values.*/
        denylist?: Value[] | undefined;
        /** Provide a HTML markup to display when no match is found.*/
        emptyState?: string | undefined;
        /** Provide arbitrary classes to nav element.*/
        regionNav?: string | undefined;
        /** Provide arbitrary classes to each list.*/
        regionList?: string | undefined;
        /** Provide arbitrary classes to each list item.*/
        regionItem?: string | undefined;
        /** Provide arbitrary classes to each button.*/
        regionButton?: string | undefined;
        /** Provide arbitrary classes to empty message.*/
        regionEmpty?: string | undefined;
        /** Provide a custom filter function.*/
        filter?: (() => AutocompleteOption<Value, Meta>[]) | undefined;
        /** Enable/Disable transitions*/
        transitions?: boolean | undefined;
        /** Provide the transition used on entry.*/
        transitionIn?: TransitionIn | undefined;
        /** Transition params provided to `transitionIn`.*/
        transitionInParams?: TransitionParams<TransitionIn> | undefined;
        /** Provide the transition used on exit.*/
        transitionOut?: TransitionOut | undefined;
        /** Transition params provided to `transitionOut`.*/
        transitionOutParams?: TransitionParams<TransitionOut> | undefined;
    };
    events(): {
        click: MouseEvent;
        keypress: KeyboardEvent;
        selection: CustomEvent<AutocompleteOption<Value, Meta>>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {};
}
export type AutocompleteProps<Value = unknown, Meta = unknown, TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<Value, Meta, TransitionIn, TransitionOut>['props']>;
export type AutocompleteEvents<Value = unknown, Meta = unknown, TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<Value, Meta, TransitionIn, TransitionOut>['events']>;
export type AutocompleteSlots<Value = unknown, Meta = unknown, TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> = ReturnType<__sveltets_Render<Value, Meta, TransitionIn, TransitionOut>['slots']>;
export default class Autocomplete<Value = unknown, Meta = unknown, TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition> extends SvelteComponentTyped<AutocompleteProps<Value, Meta, TransitionIn, TransitionOut>, AutocompleteEvents<Value, Meta, TransitionIn, TransitionOut>, AutocompleteSlots<Value, Meta, TransitionIn, TransitionOut>> {
}
export {};
