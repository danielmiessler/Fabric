import type { TransitionConfig } from 'svelte/transition';
export declare function dynamicTransition<T extends Transition>(node: Element, dynParams: DynamicTransitionParams<T>): TransitionConfig;
type DynamicTransitionParams<T extends Transition> = {
    transition: T;
    params: TransitionParams<T>;
    enabled: boolean;
};
export type Transition = (node: Element, params?: any) => TransitionConfig;
export type TransitionParams<T extends Transition> = Parameters<T>[1];
export {};
