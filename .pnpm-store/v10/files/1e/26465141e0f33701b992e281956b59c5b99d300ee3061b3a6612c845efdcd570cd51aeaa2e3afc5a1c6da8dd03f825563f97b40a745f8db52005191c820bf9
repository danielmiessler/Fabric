import epsilon from "./epsilon.js";
import relativeError from "./relative_error.js";

/**
 * Approximate equality.
 *
 * @param {number} actual The value to be tested.
 * @param {number} expected The reference value.
 * @param {number} tolerance The acceptable relative difference.
 * @return {boolean} Whether numbers are within tolerance.
 */
function approxEqual(actual, expected, tolerance = epsilon) {
    return relativeError(actual, expected) <= tolerance;
}

export default approxEqual;
