export type { ToastStore } from './stores.js';
export interface ToastSettings {
    /** Provide the toast message. Supports HTML. */
    message: string;
    /** Provide CSS classes to set the background color. */
    background?: string;
    /** Enables auto-hide after the timeout duration. */
    autohide?: boolean;
    /** Set the auto-hide timeout duration. */
    timeout?: number;
    /** Generate a custom action button UI. */
    action?: {
        /** The button label. Supports HTML. */
        label: string;
        /** The function triggered when the button is pressed. */
        response: () => void;
    };
    /** Hide dismiss button */
    hideDismiss?: boolean;
    /** Remain visible on Hover */
    hoverable?: boolean;
    /** Provide arbitrary CSS classes to style the toast. */
    classes?: string;
    /** Callback function that fires on trigger and close. */
    callback?: (response: {
        id: string;
        status: 'queued' | 'closed';
    }) => void;
}
export interface Toast extends ToastSettings {
    /** A UUID will be auto-assigned on `.trigger()`. */
    id: string;
    /** The id of the `setTimeout` if `autohide` is enabled  */
    timeoutId?: ReturnType<typeof setTimeout>;
}
