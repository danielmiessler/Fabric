import { tick } from 'svelte';
import type { ActionReturn } from 'svelte/action';
import type {
  ObserverEventDetails,
  Options,
  Position,
  ScrollDirection,
  Event,
  LifecycleEventDetails,
  Attributes,
} from './types';

const defaultOptions: Options = {
  root: null,
  rootMargin: '0px',
  threshold: 0,
  unobserveOnEnter: false,
};

const createEvent = <T = ObserverEventDetails>(
  name: Event,
  detail: T
): CustomEvent<T> => new CustomEvent(name, { detail });

export function inview(
  node: HTMLElement,
  options: Options = {}
): ActionReturn<Options, Attributes> {
  const { root, rootMargin, threshold, unobserveOnEnter }: Options = {
    ...defaultOptions,
    ...options,
  };

  let prevPos: Position = {
    x: undefined,
    y: undefined,
  };

  let scrollDirection: ScrollDirection = {
    vertical: undefined,
    horizontal: undefined,
  };

  if (typeof IntersectionObserver !== 'undefined' && node) {
    const observer = new IntersectionObserver(
      (entries, _observer) => {
        entries.forEach((singleEntry) => {
          if (prevPos.y > singleEntry.boundingClientRect.y) {
            scrollDirection.vertical = 'up';
          } else {
            scrollDirection.vertical = 'down';
          }

          if (prevPos.x > singleEntry.boundingClientRect.x) {
            scrollDirection.horizontal = 'left';
          } else {
            scrollDirection.horizontal = 'right';
          }

          prevPos = {
            y: singleEntry.boundingClientRect.y,
            x: singleEntry.boundingClientRect.x,
          };

          const detail: ObserverEventDetails = {
            inView: singleEntry.isIntersecting,
            entry: singleEntry,
            scrollDirection,
            node,
            observer: _observer,
          };

          node.dispatchEvent(createEvent('inview_change', detail));
          //@ts-expect-error only for backward compatibility
          node.dispatchEvent(createEvent('change', detail));

          if (singleEntry.isIntersecting) {
            node.dispatchEvent(createEvent('inview_enter', detail));
            //@ts-expect-error only for backward compatibility
            node.dispatchEvent(createEvent('enter', detail));

            unobserveOnEnter && _observer.unobserve(node);
          } else {
            node.dispatchEvent(createEvent('inview_leave', detail));
            //@ts-expect-error only for backward compatibility
            node.dispatchEvent(createEvent('leave', detail));
          }
        });
      },
      {
        root,
        rootMargin,
        threshold,
      }
    );

    tick().then(() => {
      node.dispatchEvent(
        createEvent<LifecycleEventDetails>('inview_init', { observer, node })
      );
      node.dispatchEvent(
        //@ts-expect-error only for backward compatibility
        createEvent<LifecycleEventDetails>('init', { observer, node })
      );
    });

    observer.observe(node);

    return {
      destroy() {
        observer.unobserve(node);
      },
    };
  }
}
