/**
 * Basic text editor in order to create a Signature annotation.
 */
export class SignatureExtractor {
    static "__#private@#PARAMETERS": {
        maxDim: number;
        sigmaSFactor: number;
        sigmaR: number;
        kernelSize: number;
    };
    static "__#private@#neighborIndexToId"(i0: any, j0: any, i: any, j: any): any;
    static "__#private@#neighborIdToIndex": Int32Array<ArrayBuffer>;
    static "__#private@#clockwiseNonZero"(buf: any, width: any, i0: any, j0: any, i: any, j: any, offset: any): number;
    static "__#private@#counterClockwiseNonZero"(buf: any, width: any, i0: any, j0: any, i: any, j: any, offset: any): number;
    static "__#private@#findContours"(buf: any, width: any, height: any, threshold: any): {
        isHole: boolean;
        points: number[];
        id: number;
        parent: number;
    }[];
    static "__#private@#douglasPeuckerHelper"(points: any, start: any, end: any, output: any): void;
    static "__#private@#douglasPeucker"(points: any): any[] | null;
    static "__#private@#bilateralFilter"(buf: any, width: any, height: any, sigmaS: any, sigmaR: any, kernelSize: any): (Uint32Array<ArrayBuffer> | Uint8Array<any>)[];
    static "__#private@#getHistogram"(buf: any): Uint32Array<ArrayBuffer>;
    static "__#private@#toUint8"(buf: any): Uint8ClampedArray<ArrayBuffer>;
    static "__#private@#guessThreshold"(histogram: any): number;
    static "__#private@#getGrayPixels"(bitmap: any): any[];
    static extractContoursFromText(text: any, { fontFamily, fontStyle, fontWeight }: {
        fontFamily: any;
        fontStyle: any;
        fontWeight: any;
    }, pageWidth: any, pageHeight: any, rotation: any, innerMargin: any): {
        outline: InkDrawOutline;
        newCurves: any[];
        areContours: any;
        thickness: any;
        width: any;
        height: any;
    } | null;
    static process(bitmap: any, pageWidth: any, pageHeight: any, rotation: any, innerMargin: any): {
        outline: InkDrawOutline;
        newCurves: any[];
        areContours: any;
        thickness: any;
        width: any;
        height: any;
    } | null;
    static processDrawnLines({ lines, pageWidth, pageHeight, rotation, innerMargin, mustSmooth, areContours, }: {
        lines: any;
        pageWidth: any;
        pageHeight: any;
        rotation: any;
        innerMargin: any;
        mustSmooth: any;
        areContours: any;
    }): {
        outline: InkDrawOutline;
        newCurves: any[];
        areContours: any;
        thickness: any;
        width: any;
        height: any;
    } | null;
    static compressSignature({ outlines, areContours, thickness, width, height, }: {
        outlines: any;
        areContours: any;
        thickness: any;
        width: any;
        height: any;
    }): Promise<any>;
    static decompressSignature(signatureData: any): Promise<{
        areContours: boolean;
        thickness: number;
        outlines: Float32Array<ArrayBuffer>[];
        width: number;
        height: number;
    } | null>;
}
import { InkDrawOutline } from "./inkdraw.js";
