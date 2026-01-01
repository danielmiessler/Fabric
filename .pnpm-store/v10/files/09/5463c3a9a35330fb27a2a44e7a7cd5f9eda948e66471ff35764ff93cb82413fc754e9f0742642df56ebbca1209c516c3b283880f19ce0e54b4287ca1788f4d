/**
 * Viewer control to display a sidebar with resizer functionality.
 */
export class Sidebar {
    /**
     * @typedef {Object} SidebarElements
     * @property {HTMLElement} sidebar - The sidebar element.
     * @property {HTMLElement} resizer - The sidebar resizer element.
     * @property {HTMLElement} toggleButton - The button used to toggle the
     *   sidebar.
     */
    /**
     * Create a sidebar with resizer functionality.
     * @param {SidebarElements} sidebarElements
     * @param {boolean} ltr
     * @param {boolean} isResizerOnTheLeft
     */
    constructor({ sidebar, resizer, toggleButton }: {
        /**
         * - The sidebar element.
         */
        sidebar: HTMLElement;
        /**
         * - The sidebar resizer element.
         */
        resizer: HTMLElement;
        /**
         * - The button used to toggle the
         * sidebar.
         */
        toggleButton: HTMLElement;
    }, ltr: boolean, isResizerOnTheLeft: boolean);
    _sidebar: HTMLElement;
    _isOpen: boolean;
    /**
     * Set the width of the sidebar in pixels.
     * @param {number} newWidth
     */
    set width(newWidth: number);
    /**
     * Get the current width of the sidebar in pixels.
     * @returns {number}
     */
    get width(): number;
    /**
     * Callback to be executed when the user starts resizing the sidebar.
     */
    onStartResizing(): void;
    /**
     * Callback to be executed when the user stops resizing the sidebar.
     */
    onStopResizing(): void;
    /**
     * Callback to be executed when the sidebar is being resized.
     * @param {number} newWidth - The new width of the sidebar in pixels.
     */
    onResizing(_newWidth: any): void;
    /**
     * Toggle the sidebar's visibility.
     * @param {boolean} [visibility] - The visibility state to set.
     */
    toggle(visibility?: boolean): void;
    destroy(): void;
    #private;
}
