export type PDFPageProxy = import("./api").PDFPageProxy;
export type PageViewport = import("./display_utils").PageViewport;
export type TextAccessibilityManager = import("../../web/text_accessibility.js").TextAccessibilityManager;
export type IDownloadManager = import("../../web/interfaces").IDownloadManager;
export type IPDFLinkService = import("../../web/interfaces").IPDFLinkService;
export type AnnotationEditorUIManager = any;
export type StructTreeLayerBuilder = import("../../web/struct_tree_layer_builder.js").StructTreeLayerBuilder;
export type CommentManager = import("../../web/comment_manager.js").CommentManager;
export type AnnotationElementParameters = {
    data: Object;
    layer: HTMLDivElement;
    linkService: IPDFLinkService;
    downloadManager?: import("../../web/interfaces").IDownloadManager | undefined;
    annotationStorage?: AnnotationStorage | undefined;
    /**
     * - Path for image resources, mainly
     * for annotation icons. Include trailing slash.
     */
    imageResourcesPath?: string | undefined;
    renderForms: boolean;
    svgFactory: Object;
    enableScripting?: boolean | undefined;
    hasJSActions?: boolean | undefined;
    fieldObjects?: Object | undefined;
};
export type AnnotationLayerParameters = {
    viewport: PageViewport;
    div: HTMLDivElement;
    annotations: any[];
    page: PDFPageProxy;
    linkService: IPDFLinkService;
    downloadManager?: import("../../web/interfaces").IDownloadManager | undefined;
    annotationStorage?: AnnotationStorage | undefined;
    /**
     * - Path for image resources, mainly
     * for annotation icons. Include trailing slash.
     */
    imageResourcesPath?: string | undefined;
    renderForms: boolean;
    /**
     * - Enable embedded script execution.
     */
    enableScripting?: boolean | undefined;
    /**
     * - Some fields have JS actions.
     * The default value is `false`.
     */
    hasJSActions?: boolean | undefined;
    fieldObjects?: {
        [x: string]: Object[];
    } | null | undefined;
    annotationCanvasMap?: Map<string, HTMLCanvasElement> | undefined;
    accessibilityManager?: import("../../web/text_accessibility.js").TextAccessibilityManager | undefined;
    annotationEditorUIManager?: AnnotationEditorUIManager;
    structTreeLayer?: import("../../web/struct_tree_layer_builder.js").StructTreeLayerBuilder | undefined;
    /**
     * - The comment manager instance.
     */
    commentManager?: import("../../web/comment_manager.js").CommentManager | undefined;
};
/**
 * @typedef {Object} AnnotationLayerParameters
 * @property {PageViewport} viewport
 * @property {HTMLDivElement} div
 * @property {Array} annotations
 * @property {PDFPageProxy} page
 * @property {IPDFLinkService} linkService
 * @property {IDownloadManager} [downloadManager]
 * @property {AnnotationStorage} [annotationStorage]
 * @property {string} [imageResourcesPath] - Path for image resources, mainly
 *   for annotation icons. Include trailing slash.
 * @property {boolean} renderForms
 * @property {boolean} [enableScripting] - Enable embedded script execution.
 * @property {boolean} [hasJSActions] - Some fields have JS actions.
 *   The default value is `false`.
 * @property {Object<string, Array<Object>> | null} [fieldObjects]
 * @property {Map<string, HTMLCanvasElement>} [annotationCanvasMap]
 * @property {TextAccessibilityManager} [accessibilityManager]
 * @property {AnnotationEditorUIManager} [annotationEditorUIManager]
 * @property {StructTreeLayerBuilder} [structTreeLayer]
 * @property {CommentManager} [commentManager] - The comment manager instance.
 */
/**
 * Manage the layer containing all the annotations.
 */
export class AnnotationLayer {
    /**
     * @private
     */
    private static get _defaultBorderStyle();
    constructor({ div, accessibilityManager, annotationCanvasMap, annotationEditorUIManager, page, viewport, structTreeLayer, commentManager, linkService, annotationStorage, }: {
        div: any;
        accessibilityManager: any;
        annotationCanvasMap: any;
        annotationEditorUIManager: any;
        page: any;
        viewport: any;
        structTreeLayer: any;
        commentManager: any;
        linkService: any;
        annotationStorage: any;
    });
    div: any;
    page: any;
    viewport: any;
    zIndex: number;
    _annotationEditorUIManager: any;
    _commentManager: any;
    popupShow: any[] | undefined;
    hasEditableAnnotations(): boolean;
    /**
     * Render a new annotation layer with all annotation elements.
     *
     * @param {AnnotationLayerParameters} params
     * @memberof AnnotationLayer
     */
    render(params: AnnotationLayerParameters): Promise<void>;
    /**
     * Add link annotations to the annotation layer.
     *
     * @param {Array<Object>} annotations
     * @param {IPDFLinkService} linkService
     * @memberof AnnotationLayer
     */
    addLinkAnnotations(annotations: Array<Object>): Promise<void>;
    /**
     * Update the annotation elements on existing annotation layer.
     *
     * @param {AnnotationLayerParameters} viewport
     * @memberof AnnotationLayer
     */
    update({ viewport }: AnnotationLayerParameters): void;
    getEditableAnnotations(): any[];
    getEditableAnnotation(id: any): any;
    addFakeAnnotation(editor: any): EditorAnnotationElement;
    removeAnnotation(id: any): void;
    updateFakeAnnotations(editors: any): void;
    togglePointerEvents(enabled?: boolean): void;
    #private;
}
export class FreeTextAnnotationElement extends AnnotationElement {
    constructor(parameters: any);
    textContent: any;
    textPosition: any;
    annotationEditorType: number;
    render(): HTMLElement | undefined;
}
export class HighlightAnnotationElement extends AnnotationElement {
    constructor(parameters: any);
    annotationEditorType: number;
    render(): HTMLElement | undefined;
}
export class InkAnnotationElement extends AnnotationElement {
    constructor(parameters: any);
    containerClassName: string;
    svgElementName: string;
    annotationEditorType: number;
    render(): HTMLElement | undefined;
    getElementsToTriggerPopup(): any[];
    #private;
}
export class StampAnnotationElement extends AnnotationElement {
    constructor(parameters: any);
    annotationEditorType: number;
    render(): HTMLElement | undefined;
}
import { AnnotationStorage } from "./annotation_storage.js";
declare class EditorAnnotationElement extends AnnotationElement {
    constructor(parameters: any);
    editor: any;
    render(): HTMLElement | undefined;
    createOrUpdatePopup(): void;
    remove(): void;
}
declare class AnnotationElement {
    static _hasPopupData({ contentsObj, richText }: {
        contentsObj: any;
        richText: any;
    }): boolean;
    constructor(parameters: any, { isRenderable, ignoreBorder, createQuadrilaterals, }?: {
        isRenderable?: boolean | undefined;
        ignoreBorder?: boolean | undefined;
        createQuadrilaterals?: boolean | undefined;
    });
    isRenderable: boolean;
    data: any;
    layer: any;
    linkService: any;
    downloadManager: any;
    imageResourcesPath: any;
    renderForms: any;
    svgFactory: any;
    annotationStorage: any;
    enableComment: any;
    enableScripting: any;
    hasJSActions: any;
    _fieldObjects: any;
    parent: any;
    hasOwnCommentButton: boolean;
    contentElement: HTMLElement | undefined;
    container: HTMLElement | undefined;
    get _isEditable(): any;
    get hasPopupData(): any;
    get commentData(): any;
    get hasCommentButton(): any;
    get commentButtonPosition(): any;
    _normalizePoint(point: any): any;
    set commentText(text: any);
    get commentText(): any;
    removePopup(): void;
    popup: any;
    updateEdited(params: any): void;
    resetEdited(): void;
    /**
     * Create an empty container for the annotation's HTML element.
     *
     * @private
     * @param {boolean} ignoreBorder
     * @memberof AnnotationElement
     * @returns {HTMLElement} A section element.
     */
    private _createContainer;
    setRotation(angle: any, container?: HTMLElement | undefined): void;
    get _commonActions(): any;
    _dispatchEventFromSandbox(actions: any, jsEvent: any): void;
    _setDefaultPropertiesFromJS(element: any): void;
    /**
     * Create quadrilaterals from the annotation's quadpoints.
     *
     * @private
     * @memberof AnnotationElement
     */
    private _createQuadrilaterals;
    /**
     * Create a popup for the annotation's HTML element. This is used for
     * annotations that do not have a Popup entry in the dictionary, but
     * are of a type that works with popups (such as Highlight annotations).
     *
     * @param {Object} [popupData] - The data for the popup, if any.
     *
     * @private
     * @memberof AnnotationElement
     */
    private _createPopup;
    get hasPopupElement(): boolean;
    get extraPopupElement(): null;
    /**
     * Render the annotation's HTML element(s).
     *
     * @public
     * @memberof AnnotationElement
     */
    public render(): void;
    /**
     * @private
     * @returns {Array}
     */
    private _getElementsByName;
    show(): void;
    hide(): void;
    /**
     * Get the HTML element(s) which can trigger a popup when clicked or hovered.
     *
     * @public
     * @memberof AnnotationElement
     * @returns {Array<HTMLElement>|HTMLElement} An array of elements or an
     *          element.
     */
    public getElementsToTriggerPopup(): Array<HTMLElement> | HTMLElement;
    addHighlightArea(): void;
    _editOnDoubleClick(): void;
    get width(): number;
    get height(): number;
    #private;
}
export {};
