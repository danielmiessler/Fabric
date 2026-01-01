(function (global, factory) {
    typeof exports === 'object' && typeof module !== 'undefined' ? factory(exports) :
    typeof define === 'function' && define.amd ? define(['exports'], factory) :
    (global = typeof globalThis !== 'undefined' ? globalThis : global || self, factory(global.Inview = {}));
})(this, (function (exports) { 'use strict';

    function run(fn) {
        return fn();
    }
    function run_all(fns) {
        fns.forEach(run);
    }

    const dirty_components = [];
    const binding_callbacks = [];
    const render_callbacks = [];
    const flush_callbacks = [];
    const resolved_promise = Promise.resolve();
    let update_scheduled = false;
    function schedule_update() {
        if (!update_scheduled) {
            update_scheduled = true;
            resolved_promise.then(flush);
        }
    }
    function tick() {
        schedule_update();
        return resolved_promise;
    }
    function add_render_callback(fn) {
        render_callbacks.push(fn);
    }
    // flush() calls callbacks in this order:
    // 1. All beforeUpdate callbacks, in order: parents before children
    // 2. All bind:this callbacks, in reverse order: children before parents.
    // 3. All afterUpdate callbacks, in order: parents before children. EXCEPT
    //    for afterUpdates called during the initial onMount, which are called in
    //    reverse order: children before parents.
    // Since callbacks might update component values, which could trigger another
    // call to flush(), the following steps guard against this:
    // 1. During beforeUpdate, any updated components will be added to the
    //    dirty_components array and will cause a reentrant call to flush(). Because
    //    the flush index is kept outside the function, the reentrant call will pick
    //    up where the earlier call left off and go through all dirty components. The
    //    current_component value is saved and restored so that the reentrant call will
    //    not interfere with the "parent" flush() call.
    // 2. bind:this callbacks cannot trigger new flush() calls.
    // 3. During afterUpdate, any updated components will NOT have their afterUpdate
    //    callback called a second time; the seen_callbacks set, outside the flush()
    //    function, guarantees this behavior.
    const seen_callbacks = new Set();
    let flushidx = 0; // Do *not* move this inside the flush() function
    function flush() {
        do {
            // first, call beforeUpdate functions
            // and update components
            while (flushidx < dirty_components.length) {
                const component = dirty_components[flushidx];
                flushidx++;
                update(component.$$);
            }
            dirty_components.length = 0;
            flushidx = 0;
            while (binding_callbacks.length)
                binding_callbacks.pop()();
            // then, once components are updated, call
            // afterUpdate functions. This may cause
            // subsequent updates...
            for (let i = 0; i < render_callbacks.length; i += 1) {
                const callback = render_callbacks[i];
                if (!seen_callbacks.has(callback)) {
                    // ...so guard against infinite loops
                    seen_callbacks.add(callback);
                    callback();
                }
            }
            render_callbacks.length = 0;
        } while (dirty_components.length);
        while (flush_callbacks.length) {
            flush_callbacks.pop()();
        }
        update_scheduled = false;
        seen_callbacks.clear();
    }
    function update($$) {
        if ($$.fragment !== null) {
            $$.update();
            run_all($$.before_update);
            const dirty = $$.dirty;
            $$.dirty = [-1];
            $$.fragment && $$.fragment.p($$.ctx, dirty);
            $$.after_update.forEach(add_render_callback);
        }
    }

    const defaultOptions = {
        root: null,
        rootMargin: '0px',
        threshold: 0,
        unobserveOnEnter: false,
    };
    const createEvent = (name, detail) => new CustomEvent(name, { detail });
    function inview(node, options = {}) {
        const { root, rootMargin, threshold, unobserveOnEnter } = Object.assign(Object.assign({}, defaultOptions), options);
        let prevPos = {
            x: undefined,
            y: undefined,
        };
        let scrollDirection = {
            vertical: undefined,
            horizontal: undefined,
        };
        if (typeof IntersectionObserver !== 'undefined' && node) {
            const observer = new IntersectionObserver((entries, _observer) => {
                entries.forEach((singleEntry) => {
                    if (prevPos.y > singleEntry.boundingClientRect.y) {
                        scrollDirection.vertical = 'up';
                    }
                    else {
                        scrollDirection.vertical = 'down';
                    }
                    if (prevPos.x > singleEntry.boundingClientRect.x) {
                        scrollDirection.horizontal = 'left';
                    }
                    else {
                        scrollDirection.horizontal = 'right';
                    }
                    prevPos = {
                        y: singleEntry.boundingClientRect.y,
                        x: singleEntry.boundingClientRect.x,
                    };
                    const detail = {
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
                    }
                    else {
                        node.dispatchEvent(createEvent('inview_leave', detail));
                        //@ts-expect-error only for backward compatibility
                        node.dispatchEvent(createEvent('leave', detail));
                    }
                });
            }, {
                root,
                rootMargin,
                threshold,
            });
            tick().then(() => {
                node.dispatchEvent(createEvent('inview_init', { observer, node }));
                node.dispatchEvent(
                //@ts-expect-error only for backward compatibility
                createEvent('init', { observer, node }));
            });
            observer.observe(node);
            return {
                destroy() {
                    observer.unobserve(node);
                },
            };
        }
    }

    exports.inview = inview;

    Object.defineProperty(exports, '__esModule', { value: true });

}));
