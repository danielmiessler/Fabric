import mean from "./mean.js";
import sampleStandardDeviation from "./sample_standard_deviation.js";

/**
 * The`coefficient of variation`_ is the ratio of the standard deviation to the mean.
 * .._`coefficient of variation`: https://en.wikipedia.org/wiki/Coefficient_of_variation
 *
 *
 * @param {Array} x input
 * @returns {number} coefficient of variation
 * @example
 * coefficientOfVariation([1, 2, 3, 4]).toFixed(3); // => 0.516
 * coefficientOfVariation([1, 2, 3, 4, 5]).toFixed(3); // => 0.527
 * coefficientOfVariation([-1, 0, 1, 2, 3, 4]).toFixed(3); // => 1.247
 */
function coefficientOfVariation(x) {
    return sampleStandardDeviation(x) / mean(x);
}

export default coefficientOfVariation;
