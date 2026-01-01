/**
 * The events fired by the Svelte action.
 */
type RevealEvents = {
    /**
     * Function that gets fired when the node starts being revealed.
     */
    onRevealStart: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the node is fully revealed.
     */
    onRevealEnd: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the `reset` option is set to `true` and the node starts transitioning out.
     */
    onResetStart: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the `reset` option is set to `true` and the node has fully transitioned out.
     */
    onResetEnd: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the node is mounted on the DOM.
     */
    onMount: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the action options are updated.
     */
    onUpdate: (node: HTMLElement) => void;
    /**
     * Function that gets fired when the node is unmounted from the DOM.
     */
    onDestroy: (node: HTMLElement) => void;
};
/**
 * The return type of the Svelte action.
 */
type ActionReturn = {
    /**
     * Lifecycle function that is triggered when the action options are updated.
     */
    update: () => void;
    /**
     * Lifecycle function that is triggered when the node is unmounted from the DOM.
     */
    destroy: () => void;
};

/**
 * The name of a built-in easing function.
 */
type StandardEasingFunction = 'linear' | 'easeInSine' | 'easeOutSine' | 'easeInOutSine' | 'easeInQuad' | 'easeOutQuad' | 'easeInOutQuad' | 'easeInCubic' | 'easeOutCubic' | 'easeInOutCubic' | 'easeInQuart' | 'easeOutQuart' | 'easeInOutQuart' | 'easeInQuint' | 'easeOutQuint' | 'easeInOutQuint' | 'easeInExpo' | 'easeOutExpo' | 'easeInOutExpo' | 'easeInCirc' | 'easeOutCirc' | 'easeInOutCirc' | 'easeInBack' | 'easeOutBack' | 'easeInOutBack';
/**
 * Array of four numbers, representing a custom bezier curve.
 */
type EasingWeights = [number, number, number, number];
/**
 * The easing function to use. It can either be an existing one or an array of four numbers representing a custom bezier curve.
 */
type Easing = StandardEasingFunction | EasingWeights;

/**
 * Configuration of the Intersection Observer used by the library.
 */
type IntersectionObserverConfig = {
    /**
     * The root element used by the Intersection Observer.
     */
    root: IntersectionObserver['root'];
    /**
     * The root margin property of the Intersection Observer.
     */
    rootMargin: IntersectionObserver['rootMargin'];
    /**
     * The `threshold` (in percentage from 0.0 to 1.0) property used by the Intersection
     * Observer to know when its target element is considered visible.
     */
    threshold: number;
};

type RevealTransition = {
    /**
     * When set to false, the transition is disabled for the target element.
     */
    disable: boolean;
    /**
     * When set to true, the node transitions out when out of view, and is revealed again when back in view.
     *
     * ⚠️ Be careful not to overuse this option.
     */
    reset: boolean;
    /**
     * How long the transition lasts (in ms).
     */
    duration: number;
    /**
     * How long the transition is delayed (in ms) before being triggered.
     */
    delay: number;
    /**
     * The starting opacity value.
     */
    opacity: number;
    /**
     * The transition preset that should be applied.
     */
    preset: 'fade' | 'slide' | 'fly' | 'spin' | 'blur' | 'scale';
    /**
     * The starting offset position in pixels on the x-axis.
     */
    x: number;
    /**
     * The starting offset position in pixels on the y-axis.
     */
    y: number;
    /**
     * The starting rotation offset in degrees.
     */
    rotate: number;
    /**
     * The starting blur value in pixels.
     */
    blur: number;
    /**
     * The starting scale value in percentage. `1` corresponds to `100%`.
     */
    scale: number;
};

/**
 * Specifies how a single reveal instance behaves.
 */
type RevealOptions = Partial<IntersectionObserverConfig & RevealTransition & {
    easing: Easing;
} & RevealEvents>;

/**
 * Reveals a given HTML node element on scroll.
 * @param node The DOM node element to apply the reveal on scroll effect to.
 * @param userOptions User-provided options to tweak the scroll animation behavior for `node`.
 * @returns The action object containing the update and destroy functions for `node`.
 */
declare const reveal: (node: HTMLElement, userOptions?: RevealOptions) => Partial<ActionReturn>;

/**
 * The device type.
 */
type Device = 'mobile' | 'tablet' | 'laptop' | 'desktop';
/**
 * The configuration of a device.
 */
type DeviceConfig = {
    /**
     * Whether the reveal effect is performed on the device.
     */
    enabled: boolean;
    /**
     * The max viewport width of the device.
     */
    breakpoint: number;
};
/**
 * Specifies how the library handles responsiveness. It can be used to enable/disable the reveal effect on some devices.
 */
type Responsive = Record<Device, DeviceConfig>;

/**
 * Global configuration for all instances of this library.
 */
interface RevealConfig {
    /**
     * Whether the reveal effect runs only once (i.e. it doesn't re-run on page reload).
     */
    once: boolean;
    /**
     * Specifies how the library handles responsiveness. It can be used to enable/disable the reveal effect on some devices.
     */
    responsive: Responsive;
}

/**
 * Sets the value of the global `once` property.
 * @param once Whether the reveal effect runs only once (i.e. it doesn't re-run on page reload).
 * @returns The config object with the updated `once` property,
 */
declare const setOnce: (once: boolean) => RevealConfig;
/**
 * Sets the status of a device.
 * @param device The device to enable/disable.
 * @param status The new status for `device`.
 * @returns The config object with the updated `enabled` property.
 */
declare const setDeviceStatus: (device: Device, status: boolean) => RevealConfig;
/**
 * Sets the status of multiple devices.
 * @param devices The devices to enabled/disable.
 * @param status The devices' new status.
 * @returns The config object with the updated `enabled` properties.
 */
declare const setDevicesStatus: (devices: Device[], status: boolean) => RevealConfig;
/**
 * Sets the breakpoint of a device.
 * @param device The device to update with `breakpoint`.
 * @param breakpoint The new breakpoint for `device`.
 * @returns The config object with the updated `breakpoint` property.
 */
declare const setDeviceBreakpoint: (device: Device, breakpoint: number) => RevealConfig;
/**
 * Sets the settings of a device.
 * @param device The device to update with `settings`.
 * @param settings The new settings for `device`.
 * @returns The config object with the updated device settings.
 */
declare const setDevice: (device: Device, settings: DeviceConfig) => RevealConfig;
/**
 * Updates how responsiveness is handled by the library.
 * @param responsive Object that instructs the library how to handle responsiveness.
 * @returns The config object with the updated `responsive` property.
 */
declare const setResponsive: (responsive: Responsive) => RevealConfig;
/**
 * Sets the Intersection Observer `root` element.
 * @param root The new Intersection Observer `root` element.
 * @returns The Intersection Obsever configuration with the updated `root` property.
 */
declare const setObserverRoot: (root: IntersectionObserver['root']) => IntersectionObserverConfig;
/**
 * Sets the Intersection Observer `rootMargin` property.
 * @param rootMargin The new `rootMargin` used by the Intersection Observer.
 * @returns The Intersection Observer configuration with the updated `rootMargin` property.
 */
declare const setObserverRootMargin: (rootMargin: IntersectionObserver['rootMargin']) => IntersectionObserverConfig;
/**
 * Sets the Intersection Observer `threshold` property.
 * @param threshold The new `threshold` used by the Intersection Observer.
 * @returns The Intersection Observer configuration object with the updated `threshold` property.
 */
declare const setObserverThreshold: (threshold: number) => IntersectionObserverConfig;
/**
 * Sets the Intersection Observer configuration.
 * @param observerConfig The new Intersection Observer configuration.
 * @returns The updated configuration used to manage the Intersection Observer behavior.
 */
declare const setObserverConfig: (observerConfig: Partial<IntersectionObserverConfig>) => IntersectionObserverConfig;
/**
 * Updates the library global configuration.
 * @param userConfig The new custom configuration.
 * @returns The updated config object.
 */
declare const setConfig: (userConfig: RevealConfig) => RevealConfig;
/**
 * Updates the default options to be used for the reveal effect.
 * @param userOptions The new default options.
 * @returns The updated default options.
 */
declare const setDefaultOptions: (userOptions: RevealOptions) => Required<RevealOptions>;

export { type RevealConfig, type RevealOptions, reveal, setConfig, setDefaultOptions, setDevice, setDeviceBreakpoint, setDeviceStatus, setDevicesStatus, setObserverConfig, setObserverRoot, setObserverRootMargin, setObserverThreshold, setOnce, setResponsive };
