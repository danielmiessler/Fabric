import sumSimple from "./sum_simple.js";

/**
 * The mean, _also known as average_,
 * is the sum of all values over the number of values.
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * The simple mean uses the successive addition method internally
 * to calculate it's result. Errors in floating-point addition are
 * not accounted for, so if precision is required, the standard {@link mean}
 * method should be used instead.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 *
 * @param {Array<number>} x sample of one or more data points
 * @throws {Error} if the length of x is less than one
 * @returns {number} mean
 * @example
 * mean([0, 10]); // => 5
 */
function meanSimple(x) {
    if (x.length === 0) {
        throw new Error("meanSimple requires at least one data point");
    }

    return sumSimple(x) / x.length;
}

export default meanSimple;
