/**
 * The [Logit](https://en.wikipedia.org/wiki/Logit)
 * is the inverse of cumulativeStdLogisticProbability,
 * and is also known as the logistic quantile function.
 *
 * @param {number} p
 * @returns {number} logit
 */
function logit(p) {
    if (p <= 0 || p >= 1) {
        throw new Error("p must be strictly between zero and one");
    }
    return Math.log(p / (1 - p));
}

export default logit;
