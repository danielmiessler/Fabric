/**
 * The [log average](https://en.wikipedia.org/wiki/https://en.wikipedia.org/wiki/Geometric_mean#Relationship_with_logarithms)
 * is an equivalent way of computing the geometric mean of an array suitable for large or small products.
 *
 * It's found by calculating the average logarithm of the elements and exponentiating.
 *
 * @param {Array<number>} x sample of one or more data points
 * @returns {number} geometric mean
 * @throws {Error} if x is empty
 * @throws {Error} if x contains a negative number
 */
function logAverage(x) {
    if (x.length === 0) {
        throw new Error("logAverage requires at least one data point");
    }

    let value = 0;
    for (let i = 0; i < x.length; i++) {
        if (x[i] < 0) {
            throw new Error(
                "logAverage requires only non-negative numbers as input"
            );
        }
        value += Math.log(x[i]);
    }

    return Math.exp(value / x.length);
}

export default logAverage;
