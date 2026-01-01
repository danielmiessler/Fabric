export class BasePDFPageView {
    constructor(options: any);
    canvas: null;
    /** @type {null | HTMLDivElement} */
    div: null | HTMLDivElement;
    enableOptimizedPartialRendering: boolean;
    eventBus: null;
    id: null;
    pageColors: null;
    recordedBBoxes: null;
    renderingQueue: null;
    renderTask: null;
    resume: null;
    set renderingState(state: number);
    get renderingState(): number;
    _createCanvas(onShow: any, hideUntilComplete?: boolean): {
        canvas: HTMLCanvasElement;
        prevCanvas: null;
    };
    _resetCanvas(): void;
    _drawCanvas(options: any, onCancel: any, onFinish: any): Promise<void>;
    cancelRendering({ cancelExtraDelay }?: {
        cancelExtraDelay?: number | undefined;
    }): void;
    dispatchPageRender(): void;
    dispatchPageRendered(cssTransform: any, isDetailView: any): void;
    #private;
}
