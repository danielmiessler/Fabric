import factorial from "./factorial.js";

/**
 * Compute the [gamma function](https://en.wikipedia.org/wiki/Gamma_function) of a value using Nemes' approximation.
 * The gamma of n is equivalent to (n-1)!, but unlike the factorial function, gamma is defined for all real n except zero
 * and negative integers (where NaN is returned). Note, the gamma function is also well-defined for complex numbers,
 * though this implementation currently does not handle complex numbers as input values.
 * Nemes' approximation is defined [here](https://arxiv.org/abs/1003.6020) as Theorem 2.2.
 * Negative values use [Euler's reflection formula](https://en.wikipedia.org/wiki/Gamma_function#Properties) for computation.
 *
 * @param {number} n Any real number except for zero and negative integers.
 * @returns {number} The gamma of the input value.
 *
 * @example
 * gamma(11.5); // 11899423.084037038
 * gamma(-11.5); // 2.29575810481609e-8
 * gamma(5); // 24
 */
function gamma(n) {
    if (Number.isInteger(n)) {
        if (n <= 0) {
            // gamma not defined for zero or negative integers
            return Number.NaN;
        } else {
            // use factorial for integer inputs
            return factorial(n - 1);
        }
    }

    // Decrement n, because approximation is defined for n - 1
    n--;

    if (n < 0) {
        // Use Euler's reflection formula for negative inputs
        // see:  https://en.wikipedia.org/wiki/Gamma_function#Properties
        return Math.PI / (Math.sin(Math.PI * -n) * gamma(-n));
    } else {
        // Nemes' expansion approximation
        const seriesCoefficient =
            Math.pow(n / Math.E, n) * Math.sqrt(2 * Math.PI * (n + 1 / 6));

        const seriesDenom = n + 1 / 4;

        const seriesExpansion =
            1 +
            1 / 144 / Math.pow(seriesDenom, 2) -
            1 / 12960 / Math.pow(seriesDenom, 3) -
            257 / 207360 / Math.pow(seriesDenom, 4) -
            52 / 2612736 / Math.pow(seriesDenom, 5) +
            5741173 / 9405849600 / Math.pow(seriesDenom, 6) +
            37529 / 18811699200 / Math.pow(seriesDenom, 7);

        return seriesCoefficient * seriesExpansion;
    }
}

export default gamma;
