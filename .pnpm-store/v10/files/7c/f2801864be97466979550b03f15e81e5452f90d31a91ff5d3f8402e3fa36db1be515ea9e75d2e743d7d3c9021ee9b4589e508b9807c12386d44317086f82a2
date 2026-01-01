import mean from "./mean.js";

/**
 * [Skewness](http://en.wikipedia.org/wiki/Skewness) is
 * a measure of the extent to which a probability distribution of a
 * real-valued random variable "leans" to one side of the mean.
 * The skewness value can be positive or negative, or even undefined.
 *
 * Implementation is based on the adjusted Fisher-Pearson standardized
 * moment coefficient, which is the version found in Excel and several
 * statistical packages including Minitab, SAS and SPSS.
 *
 * @since 4.1.0
 * @param {Array<number>} x a sample of 3 or more data points
 * @returns {number} sample skewness
 * @throws {Error} if x has length less than 3
 * @example
 * sampleSkewness([2, 4, 6, 3, 1]); // => 0.590128656384365
 */
function sampleSkewness(x) {
    if (x.length < 3) {
        throw new Error("sampleSkewness requires at least three data points");
    }

    const meanValue = mean(x);
    let tempValue;
    let sumSquaredDeviations = 0;
    let sumCubedDeviations = 0;

    for (let i = 0; i < x.length; i++) {
        tempValue = x[i] - meanValue;
        sumSquaredDeviations += tempValue * tempValue;
        sumCubedDeviations += tempValue * tempValue * tempValue;
    }

    // this is Bessels' Correction: an adjustment made to sample statistics
    // that allows for the reduced degree of freedom entailed in calculating
    // values from samples rather than complete populations.
    const besselsCorrection = x.length - 1;

    // Find the mean value of that list
    const theSampleStandardDeviation = Math.sqrt(
        sumSquaredDeviations / besselsCorrection
    );

    const n = x.length;
    const cubedS = Math.pow(theSampleStandardDeviation, 3);

    return (n * sumCubedDeviations) / ((n - 1) * (n - 2) * cubedS);
}

export default sampleSkewness;
