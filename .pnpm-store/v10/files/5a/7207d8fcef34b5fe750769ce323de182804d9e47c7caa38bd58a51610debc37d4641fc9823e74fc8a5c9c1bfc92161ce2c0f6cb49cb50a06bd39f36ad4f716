/**
 * The Inverse [Gaussian error function](http://en.wikipedia.org/wiki/Error_function)
 * returns a numerical approximation to the value that would have caused
 * `errorFunction()` to return x.
 *
 * @param {number} x value of error function
 * @returns {number} estimated inverted value
 */
function inverseErrorFunction(x) {
    const a = (8 * (Math.PI - 3)) / (3 * Math.PI * (4 - Math.PI));

    const inv = Math.sqrt(
        Math.sqrt(
            Math.pow(2 / (Math.PI * a) + Math.log(1 - x * x) / 2, 2) -
                Math.log(1 - x * x) / a
        ) -
            (2 / (Math.PI * a) + Math.log(1 - x * x) / 2)
    );

    if (x >= 0) {
        return inv;
    } else {
        return -inv;
    }
}

export default inverseErrorFunction;
