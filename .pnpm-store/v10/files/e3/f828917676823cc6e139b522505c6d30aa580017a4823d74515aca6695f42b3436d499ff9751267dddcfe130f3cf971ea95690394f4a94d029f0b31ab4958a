/**
 * The simple [sum](https://en.wikipedia.org/wiki/Summation) of an array
 * is the result of adding all numbers together, starting from zero.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x input
 * @return {number} sum of all input numbers
 * @example
 * sumSimple([1, 2, 3]); // => 6
 */
function sumSimple(x) {
    let value = 0;
    for (let i = 0; i < x.length; i++) {
        if (typeof x[i] !== "number") {
            return Number.NaN;
        }
        value += x[i];
    }
    return value;
}

export default sumSimple;
