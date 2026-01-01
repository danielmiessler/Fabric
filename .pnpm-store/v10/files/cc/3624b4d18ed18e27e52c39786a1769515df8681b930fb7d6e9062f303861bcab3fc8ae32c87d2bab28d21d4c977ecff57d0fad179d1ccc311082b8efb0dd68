/**
 * Relative error.
 *
 * This is more difficult to calculate than it first appears [1,2].  The usual
 * formula for the relative error between an actual value A and an expected
 * value E is `|(A-E)/E|`, but:
 *
 * 1. If the expected value is 0, any other value has infinite relative error,
 *    which is counter-intuitive: if the expected voltage is 0, getting 1/10th
 *    of a volt doesn't feel like an infinitely large error.
 *
 * 2. This formula does not satisfy the mathematical definition of a metric [3].
 *    [4] solved this problem by defining the relative error as `|ln(|A/E|)|`,
 *    but that formula only works if all values are positive: for example, it
 *    reports the relative error of -10 and 10 as 0.
 *
 * Our implementation sticks with convention and returns:
 *
 * - 0 if the actual and expected values are both zero
 * - Infinity if the actual value is non-zero and the expected value is zero
 * - `|(A-E)/E|` in all other cases
 *
 * [1] https://math.stackexchange.com/questions/677852/how-to-calculate-relative-error-when-true-value-is-zero
 * [2] https://en.wikipedia.org/wiki/Relative_change_and_difference
 * [3] https://en.wikipedia.org/wiki/Metric_(mathematics)#Definition
 * [4] F.W.J. Olver: "A New Approach to Error Arithmetic." SIAM Journal on
 *     Numerical Analysis, 15(2), 1978, 10.1137/0715024.
 *
 * @param {number} actual The actual value.
 * @param {number} expected The expected value.
 * @return {number} The relative error.
 */
function relativeError(actual, expected) {
    // These lines are actually covered by tests, but it seems
    // like c8 has a bug that marks them as not covered.
    /* c8 ignore start */
    if (actual === 0 && expected === 0) {
        return 0;
    }
    /* c8 ignore end */
    return Math.abs((actual - expected) / expected);
}

export default relativeError;
