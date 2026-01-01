export type Prefix = 'inview';
export type EventBase = 'change' | 'leave' | 'enter' | 'init';
export type Event = `${Prefix}_${EventBase}`;
export type Options = {
    root?: HTMLElement | null;
    rootMargin?: string;
    threshold?: number | number[];
    unobserveOnEnter?: boolean;
};
export type Position = {
    x?: number;
    y?: number;
};
type Direction = 'up' | 'down' | 'left' | 'right';
export type ScrollDirection = {
    vertical?: Direction;
    horizontal?: Direction;
};
export type ObserverEventDetails = {
    inView: boolean;
    entry: IntersectionObserverEntry;
    scrollDirection: ScrollDirection;
    node: HTMLElement;
    observer: IntersectionObserver;
};
export type LifecycleEventDetails = {
    node: HTMLElement;
    observer: IntersectionObserver;
};
export interface Attributes {
    'on:inview_change'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_enter'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_leave'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_init'?: (e: CustomEvent<LifecycleEventDetails>) => void;
    oninview_change?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_enter?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_leave?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_init?: (e: CustomEvent<LifecycleEventDetails>) => void;
}
export {};
