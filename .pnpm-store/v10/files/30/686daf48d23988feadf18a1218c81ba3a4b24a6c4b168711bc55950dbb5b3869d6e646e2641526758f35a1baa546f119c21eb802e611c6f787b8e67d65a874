/** Placement https://floating-ui.com/docs/computePosition#placement */
type Direction = 'top' | 'bottom' | 'left' | 'right';
type Placement = Direction | `${Direction}-start` | `${Direction}-end`;
export interface Middleware {
    /** Offset middleware settings: https://floating-ui.com/docs/offset */
    offset?: number | Record<string, any>;
    /** Shift middleware settings: https://floating-ui.com/docs/shift */
    shift?: Record<string, any>;
    /** Flip middleware settings: https://floating-ui.com/docs/flip */
    flip?: Record<string, any>;
    /** Arrow middleware settings: https://floating-ui.com/docs/arrow */
    arrow?: {
        element: string;
    } & Record<string, any>;
    /** Size middleware settings: https://floating-ui.com/docs/size */
    size?: Record<string, any>;
    /** Auto Placement middleware settings: https://floating-ui.com/docs/autoPlacement */
    autoPlacement?: Record<string, any>;
    /** Hide middleware settings: https://floating-ui.com/docs/hide */
    hide?: Record<string, any>;
    /** Inline middleware settings: https://floating-ui.com/docs/inline */
    inline?: Record<string, any>;
}
export interface PopupSettings {
    /** Provide the event type. */
    event: 'click' | 'hover' | 'focus-blur' | 'focus-click';
    /** Match the popup data value `data-popup="targetNameHere"` */
    target: string;
    /** Set the placement position. Defaults 'bottom'. */
    placement?: Placement;
    /** Query elements that close the popup when clicked. Defaults `'a[href], button'`. */
    closeQuery?: string;
    /** Optional callback function that reports state change. */
    state?: (event: {
        state: boolean;
    }) => void;
    /** Provide Floating UI middleware settings. */
    middleware?: Middleware;
}
export {};
