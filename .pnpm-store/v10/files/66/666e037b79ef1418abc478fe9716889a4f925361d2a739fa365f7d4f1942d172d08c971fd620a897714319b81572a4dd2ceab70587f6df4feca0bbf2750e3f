// Due to an ambivalent nature of .d.ts files, we can't import or export
// anything in this file. That is the cause we need to manually put ScrollDirection, ObserverEventDetails
// and LifecycleEventDetails types to ensure correct typings in the app.

type Direction = 'up' | 'down' | 'left' | 'right';

type ScrollDirection = {
  vertical?: Direction;
  horizontal?: Direction;
};

type ObserverEventDetails = {
  inView: boolean;
  entry: IntersectionObserverEntry;
  scrollDirection: ScrollDirection;
  node: HTMLElement;
  observer: IntersectionObserver;
};

type LifecycleEventDetails = {
  node: HTMLElement;
  observer: IntersectionObserver;
};

declare namespace svelteHTML {
  interface HTMLAttributes<T> {
    // Backwards compatibility
    'on:inview_change'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_enter'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_leave'?: (e: CustomEvent<ObserverEventDetails>) => void;
    'on:inview_init'?: (e: CustomEvent<LifecycleEventDetails>) => void;

    // Svelte5 Syntax
    oninview_change?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_enter?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_leave?: (e: CustomEvent<ObserverEventDetails>) => void;
    oninview_init?: (e: CustomEvent<LifecycleEventDetails>) => void;
  }
}
