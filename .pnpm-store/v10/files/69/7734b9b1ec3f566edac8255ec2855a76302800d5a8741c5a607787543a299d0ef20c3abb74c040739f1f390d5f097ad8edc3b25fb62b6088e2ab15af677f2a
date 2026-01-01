import mean from "./mean.js";

/**
 * [Kurtosis](http://en.wikipedia.org/wiki/Kurtosis) is
 * a measure of the heaviness of a distribution's tails relative to its
 * variance. The kurtosis value can be positive or negative, or even undefined.
 *
 * Implementation is based on Fisher's excess kurtosis definition and uses
 * unbiased moment estimators. This is the version found in Excel and available
 * in several statistical packages, including SAS and SciPy.
 *
 * @param {Array<number>} x a sample of 4 or more data points
 * @returns {number} sample kurtosis
 * @throws {Error} if x has length less than 4
 * @example
 * sampleKurtosis([1, 2, 2, 3, 5]); // => 1.4555765595463122
 */
function sampleKurtosis(x) {
    const n = x.length;

    if (n < 4) {
        throw new Error("sampleKurtosis requires at least four data points");
    }

    const meanValue = mean(x);
    let tempValue;
    let secondCentralMoment = 0;
    let fourthCentralMoment = 0;

    for (let i = 0; i < n; i++) {
        tempValue = x[i] - meanValue;
        secondCentralMoment += tempValue * tempValue;
        fourthCentralMoment += tempValue * tempValue * tempValue * tempValue;
    }

    return (
        ((n - 1) / ((n - 2) * (n - 3))) *
        ((n * (n + 1) * fourthCentralMoment) /
            (secondCentralMoment * secondCentralMoment) -
            3 * (n - 1))
    );
}

export default sampleKurtosis;
