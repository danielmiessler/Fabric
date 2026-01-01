export type Kernels = "gaussian";
export type BandwidthMethods = "nrd";

/**
 * https://simple-statistics.github.io/docs/#kde
 */
declare function kernelDensityEstimation(
    X: number[],
    kernel?: Kernels | ((u: number) => number),
    bandwidthMethod?: BandwidthMethods | number
): (x: number) => number;

export default kernelDensityEstimation;
