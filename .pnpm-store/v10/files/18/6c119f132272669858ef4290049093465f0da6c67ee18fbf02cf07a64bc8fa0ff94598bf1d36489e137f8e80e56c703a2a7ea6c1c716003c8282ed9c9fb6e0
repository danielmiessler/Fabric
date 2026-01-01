// Define series coefficients
const COEFFICIENTS = [
    0.99999999999999709182, 57.156235665862923517, -59.597960355475491248,
    14.136097974741747174, -0.49191381609762019978, 0.33994649984811888699e-4,
    0.46523628927048575665e-4, -0.98374475304879564677e-4,
    0.15808870322491248884e-3, -0.21026444172410488319e-3,
    0.2174396181152126432e-3, -0.16431810653676389022e-3,
    0.84418223983852743293e-4, -0.2619083840158140867e-4,
    0.36899182659531622704e-5
];

const g = 607 / 128;
const LOGSQRT2PI = Math.log(Math.sqrt(2 * Math.PI));

/**
 * Compute the logarithm of the [gamma function](https://en.wikipedia.org/wiki/Gamma_function) of a value using Lanczos' approximation.
 * This function takes as input any real-value n greater than 0.
 * This function is useful for values of n too large for the normal gamma function (n > 165).
 * The code is based on Lanczo's Gamma approximation, defined [here](http://my.fit.edu/~gabdo/gamma.txt).
 *
 * @param {number} n Any real number greater than zero.
 * @returns {number} The logarithm of gamma of the input value.
 *
 * @example
 * gammaln(500); // 2605.1158503617335
 * gammaln(2.4); // 0.21685932244884043
 */
function gammaln(n) {
    // Return infinity if value not in domain
    if (n <= 0) {
        return Number.POSITIVE_INFINITY;
    }

    // Decrement n, because approximation is defined for n - 1
    n--;

    // Create series approximation
    let a = COEFFICIENTS[0];

    for (let i = 1; i < 15; i++) {
        a += COEFFICIENTS[i] / (n + i);
    }

    const tmp = g + 0.5 + n;

    // Return natural logarithm of gamma(n)
    return LOGSQRT2PI + Math.log(a) - tmp + (n + 0.5) * Math.log(tmp);
}

export default gammaln;
