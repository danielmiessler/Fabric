export class CssFontInfo {
    static strings: string[];
    static write(info: any): ArrayBuffer;
    constructor(buffer: any);
    get fontFamily(): string;
    get fontWeight(): string;
    get italicAngle(): string;
    #private;
}
export class FontInfo {
    static bools: string[];
    static numbers: string[];
    static strings: string[];
    static "__#private@#OFFSET_NUMBERS": number;
    static "__#private@#OFFSET_BBOX": number;
    static "__#private@#OFFSET_FONT_MATRIX": number;
    static "__#private@#OFFSET_DEFAULT_VMETRICS": number;
    static "__#private@#OFFSET_STRINGS": number;
    static write(font: any): any;
    constructor({ data, extra }: {
        data: any;
        extra: any;
    });
    get black(): boolean | undefined;
    get bold(): boolean | undefined;
    get disableFontFace(): boolean | undefined;
    get fontExtraProperties(): boolean | undefined;
    get isInvalidPDFjsFont(): boolean | undefined;
    get isType3Font(): boolean | undefined;
    get italic(): boolean | undefined;
    get missingFile(): boolean | undefined;
    get remeasure(): boolean | undefined;
    get vertical(): boolean | undefined;
    get ascent(): number;
    get defaultWidth(): number;
    get descent(): number;
    get bbox(): number[] | undefined;
    get fontMatrix(): number[] | undefined;
    get defaultVMetrics(): number[] | undefined;
    get fallbackName(): string;
    get loadedName(): string;
    get mimetype(): string;
    get name(): string;
    get data(): Uint8Array<any> | undefined;
    clearData(): void;
    get cssFontInfo(): CssFontInfo | null;
    get systemFontInfo(): SystemFontInfo | null;
    #private;
}
export class FontPathInfo {
    static write(path: any): ArrayBuffer;
    constructor(buffer: any);
    get path(): Float32Array<any> | Float16Array<any>;
    #private;
}
export class PatternInfo {
    static "__#private@#KIND": number;
    static "__#private@#HAS_BBOX": number;
    static "__#private@#HAS_BACKGROUND": number;
    static "__#private@#SHADING_TYPE": number;
    static "__#private@#N_COORD": number;
    static "__#private@#N_COLOR": number;
    static "__#private@#N_STOP": number;
    static "__#private@#N_FIGURES": number;
    static write(ir: any): ArrayBuffer;
    constructor(buffer: any);
    buffer: any;
    view: DataView<any>;
    data: Uint8Array<any>;
    getIR(): (string | number | number[] | Uint8Array<any> | Float32Array<any> | {
        type: number;
        coords: Int32Array<any>;
        colors: Int32Array<any>;
    }[] | null)[] | (string | number | number[] | (string | number)[][] | null)[];
}
export class SystemFontInfo {
    static strings: string[];
    static write(info: any): any;
    constructor(buffer: any);
    get guessFallback(): boolean;
    get css(): string;
    get loadedName(): string;
    get baseFontName(): string;
    get src(): string;
    get style(): {
        style: string;
        weight: string;
    };
    #private;
}
