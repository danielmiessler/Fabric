/**
 * [Simple linear regression](http://en.wikipedia.org/wiki/Simple_linear_regression)
 * is a simple way to find a fitted line
 * between a set of coordinates. This algorithm finds the slope and y-intercept of a regression line
 * using the least sum of squares.
 *
 * @param {Array<Array<number>>} data an array of two-element of arrays,
 * like `[[0, 1], [2, 3]]`
 * @returns {Object} object containing slope and intersect of regression line
 * @example
 * linearRegression([[0, 0], [1, 1]]); // => { m: 1, b: 0 }
 */
function linearRegression(data) {
    var m;
    var b;

    // Store data length in a local variable to reduce
    // repeated object property lookups
    var dataLength = data.length;

    //if there's only one point, arbitrarily choose a slope of 0
    //and a y-intercept of whatever the y of the initial point is
    if (dataLength === 1) {
        m = 0;
        b = data[0][1];
    } else {
        // Initialize our sums and scope the `m` and `b`
        // variables that define the line.
        var sumX = 0;
        var sumY = 0;
        var sumXX = 0;
        var sumXY = 0;

        // Use local variables to grab point values
        // with minimal object property lookups
        var point;
        var x;
        var y;

        // Gather the sum of all x values, the sum of all
        // y values, and the sum of x^2 and (x*y) for each
        // value.
        //
        // In math notation, these would be SS_x, SS_y, SS_xx, and SS_xy
        for (var i = 0; i < dataLength; i++) {
            point = data[i];
            x = point[0];
            y = point[1];

            sumX += x;
            sumY += y;

            sumXX += x * x;
            sumXY += x * y;
        }

        // `m` is the slope of the regression line
        m =
            (dataLength * sumXY - sumX * sumY) /
            (dataLength * sumXX - sumX * sumX);

        // `b` is the y-intercept of the line.
        b = sumY / dataLength - (m * sumX) / dataLength;
    }

    // Return both values as an object.
    return {
        m: m,
        b: b
    };
}

/**
 * Given the output of `linearRegression`: an object
 * with `m` and `b` values indicating slope and intercept,
 * respectively, generate a line function that translates
 * x values into y values.
 *
 * @param {Object} mb object with `m` and `b` members, representing
 * slope and intersect of desired line
 * @returns {Function} method that computes y-value at any given
 * x-value on the line.
 * @example
 * var l = linearRegressionLine(linearRegression([[0, 0], [1, 1]]));
 * l(0) // = 0
 * l(2) // = 2
 * linearRegressionLine({ b: 0, m: 1 })(1); // => 1
 * linearRegressionLine({ b: 1, m: 1 })(1); // => 2
 */
function linearRegressionLine(mb /*: { b: number, m: number }*/) {
    // Return a function that computes a `y` value for each
    // x value it is given, based on the values of `b` and `a`
    // that we just computed.
    return function (x) {
        return mb.b + mb.m * x;
    };
}

/**
 * Our default sum is the [Kahan-Babuska algorithm](https://pdfs.semanticscholar.org/1760/7d467cda1d0277ad272deb2113533131dc09.pdf).
 * This method is an improvement over the classical
 * [Kahan summation algorithm](https://en.wikipedia.org/wiki/Kahan_summation_algorithm).
 * It aims at computing the sum of a list of numbers while correcting for
 * floating-point errors. Traditionally, sums are calculated as many
 * successive additions, each one with its own floating-point roundoff. These
 * losses in precision add up as the number of numbers increases. This alternative
 * algorithm is more accurate than the simple way of calculating sums by simple
 * addition.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x input
 * @return {number} sum of all input numbers
 * @example
 * sum([1, 2, 3]); // => 6
 */
function sum(x) {
    // If the array is empty, we needn't bother computing its sum
    if (x.length === 0) {
        return 0;
    }

    // Initializing the sum as the first number in the array
    var sum = x[0];

    // Keeping track of the floating-point error correction
    var correction = 0;

    var transition;

    if (typeof sum !== "number") {
        return Number.NaN;
    }

    for (var i = 1; i < x.length; i++) {
        if (typeof x[i] !== "number") {
            return Number.NaN;
        }
        transition = sum + x[i];

        // Here we need to update the correction in a different fashion
        // if the new absolute value is greater than the absolute sum
        if (Math.abs(sum) >= Math.abs(x[i])) {
            correction += sum - transition + x[i];
        } else {
            correction += x[i] - transition + sum;
        }

        sum = transition;
    }

    // Returning the corrected sum
    return sum + correction;
}

/**
 * The mean, _also known as average_,
 * is the sum of all values over the number of values.
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @throws {Error} if the length of x is less than one
 * @returns {number} mean
 * @example
 * mean([0, 10]); // => 5
 */
function mean(x) {
    if (x.length === 0) {
        throw new Error("mean requires at least one data point");
    }

    return sum(x) / x.length;
}

/**
 * The sum of deviations to the Nth power.
 * When n=2 it's the sum of squared deviations.
 * When n=3 it's the sum of cubed deviations.
 *
 * @param {Array<number>} x
 * @param {number} n power
 * @returns {number} sum of nth power deviations
 *
 * @example
 * var input = [1, 2, 3];
 * // since the variance of a set is the mean squared
 * // deviations, we can calculate that with sumNthPowerDeviations:
 * sumNthPowerDeviations(input, 2) / input.length;
 */
function sumNthPowerDeviations(x, n) {
    var meanValue = mean(x);
    var sum = 0;
    var tempValue;
    var i;

    // This is an optimization: when n is 2 (we're computing a number squared),
    // multiplying the number by itself is significantly faster than using
    // the Math.pow method.
    if (n === 2) {
        for (i = 0; i < x.length; i++) {
            tempValue = x[i] - meanValue;
            sum += tempValue * tempValue;
        }
    } else {
        for (i = 0; i < x.length; i++) {
            sum += Math.pow(x[i] - meanValue, n);
        }
    }

    return sum;
}

/**
 * The [variance](http://en.wikipedia.org/wiki/Variance)
 * is the sum of squared deviations from the mean.
 *
 * This is an implementation of variance, not sample variance:
 * see the `sampleVariance` method if you want a sample measure.
 *
 * @param {Array<number>} x a population of one or more data points
 * @returns {number} variance: a value greater than or equal to zero.
 * zero indicates that all values are identical.
 * @throws {Error} if x's length is 0
 * @example
 * variance([1, 2, 3, 4, 5, 6]); // => 2.9166666666666665
 */
function variance(x) {
    if (x.length === 0) {
        throw new Error("variance requires at least one data point");
    }

    // Find the mean of squared deviations between the
    // mean value and each value.
    return sumNthPowerDeviations(x, 2) / x.length;
}

/**
 * The [standard deviation](http://en.wikipedia.org/wiki/Standard_deviation)
 * is the square root of the variance. This is also known as the population
 * standard deviation. It's useful for measuring the amount
 * of variation or dispersion in a set of values.
 *
 * Standard deviation is only appropriate for full-population knowledge: for
 * samples of a population, {@link sampleStandardDeviation} is
 * more appropriate.
 *
 * @param {Array<number>} x input
 * @returns {number} standard deviation
 * @example
 * variance([2, 4, 4, 4, 5, 5, 7, 9]); // => 4
 * standardDeviation([2, 4, 4, 4, 5, 5, 7, 9]); // => 2
 */
function standardDeviation(x) {
    if (x.length === 1) {
        return 0;
    }
    var v = variance(x);
    return Math.sqrt(v);
}

/**
 * The [R Squared](http://en.wikipedia.org/wiki/Coefficient_of_determination)
 * value of data compared with a function `f`
 * is the sum of the squared differences between the prediction
 * and the actual value.
 *
 * @param {Array<Array<number>>} x input data: this should be doubly-nested
 * @param {Function} func function called on `[i][0]` values within the dataset
 * @returns {number} r-squared value
 * @example
 * var samples = [[0, 0], [1, 1]];
 * var regressionLine = linearRegressionLine(linearRegression(samples));
 * rSquared(samples, regressionLine); // = 1 this line is a perfect fit
 */
function rSquared(x, func) {
    if (x.length < 2) {
        return 1;
    }

    // Compute the average y value for the actual
    // data set in order to compute the
    // _total sum of squares_
    var sum = 0;
    for (var i = 0; i < x.length; i++) {
        sum += x[i][1];
    }
    var average = sum / x.length;

    // Compute the total sum of squares - the
    // squared difference between each point
    // and the average of all points.
    var sumOfSquares = 0;
    for (var j = 0; j < x.length; j++) {
        sumOfSquares += Math.pow(average - x[j][1], 2);
    }

    // Finally estimate the error: the squared
    // difference between the estimate and the actual data
    // value at each point.
    var err = 0;
    for (var k = 0; k < x.length; k++) {
        err += Math.pow(x[k][1] - func(x[k][0]), 2);
    }

    // As the error grows larger, its ratio to the
    // sum of squares increases and the r squared
    // value grows lower.
    return 1 - err / sumOfSquares;
}

/**
 * The [mode](https://en.wikipedia.org/wiki/Mode_%28statistics%29) is the number
 * that appears in a list the highest number of times.
 * There can be multiple modes in a list: in the event of a tie, this
 * algorithm will return the most recently seen mode.
 *
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * This runs in `O(n)` because the input is sorted.
 *
 * @param {Array<number>} sorted a sample of one or more data points
 * @returns {number} mode
 * @throws {Error} if sorted is empty
 * @example
 * modeSorted([0, 0, 1]); // => 0
 */
function modeSorted(sorted) {
    // Handle edge cases:
    // The mode of an empty list is undefined
    if (sorted.length === 0) {
        throw new Error("mode requires at least one data point");
    }
    if (sorted.length === 1) {
        return sorted[0];
    }

    // This assumes it is dealing with an array of size > 1, since size
    // 0 and 1 are handled immediately. Hence it starts at index 1 in the
    // array.
    var last = sorted[0];
    // store the mode as we find new modes
    var value = Number.NaN;
    // store how many times we've seen the mode
    var maxSeen = 0;
    // how many times the current candidate for the mode
    // has been seen
    var seenThis = 1;

    // end at sorted.length + 1 to fix the case in which the mode is
    // the highest number that occurs in the sequence. the last iteration
    // compares sorted[i], which is undefined, to the highest number
    // in the series
    for (var i = 1; i < sorted.length + 1; i++) {
        // we're seeing a new number pass by
        if (sorted[i] !== last) {
            // the last number is the new mode since we saw it more
            // often than the old one
            if (seenThis > maxSeen) {
                maxSeen = seenThis;
                value = last;
            }
            seenThis = 1;
            last = sorted[i];
            // if this isn't a new number, it's one more occurrence of
            // the potential mode
        } else {
            seenThis++;
        }
    }
    return value;
}

/**
 * Sort an array of numbers by their numeric value, ensuring that the
 * array is not changed in place.
 *
 * This is necessary because the default behavior of .sort
 * in JavaScript is to sort arrays as string values
 *
 *     [1, 10, 12, 102, 20].sort()
 *     // output
 *     [1, 10, 102, 12, 20]
 *
 * @param {Array<number>} x input array
 * @return {Array<number>} sorted array
 * @private
 * @example
 * numericSort([3, 2, 1]) // => [1, 2, 3]
 */
function numericSort(x) {
    return (
        x
            // ensure the array is not changed in-place
            .slice()
            // comparator function that treats input as numeric
            .sort(function (a, b) {
                return a - b;
            })
    );
}

/**
 * The [mode](https://en.wikipedia.org/wiki/Mode_%28statistics%29) is the number
 * that appears in a list the highest number of times.
 * There can be multiple modes in a list: in the event of a tie, this
 * algorithm will return the most recently seen mode.
 *
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * This runs in `O(n log(n))` because it needs to sort the array internally
 * before running an `O(n)` search to find the mode.
 *
 * @param {Array<number>} x input
 * @returns {number} mode
 * @example
 * mode([0, 0, 1]); // => 0
 */
function mode(x) {
    // Sorting the array lets us iterate through it below and be sure
    // that every time we see a new number it's new and we'll never
    // see the same number twice
    return modeSorted(numericSort(x));
}

/* globals Map: false */

/**
 * The [mode](https://en.wikipedia.org/wiki/Mode_%28statistics%29) is the number
 * that appears in a list the highest number of times.
 * There can be multiple modes in a list: in the event of a tie, this
 * algorithm will return the most recently seen mode.
 *
 * modeFast uses a Map object to keep track of the mode, instead of the approach
 * used with `mode`, a sorted array. As a result, it is faster
 * than `mode` and supports any data type that can be compared with `==`.
 * It also requires a
 * [JavaScript environment with support for Map](https://kangax.github.io/compat-table/es6/#test-Map),
 * and will throw an error if Map is not available.
 *
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * @param {Array<*>} x a sample of one or more data points
 * @returns {?*} mode
 * @throws {ReferenceError} if the JavaScript environment doesn't support Map
 * @throws {Error} if x is empty
 * @example
 * modeFast(['rabbits', 'rabbits', 'squirrels']); // => 'rabbits'
 */
function modeFast(x) {
    // This index will reflect the incidence of different values, indexing
    // them like
    // { value: count }
    var index = new Map();

    // A running `mode` and the number of times it has been encountered.
    var mode;
    var modeCount = 0;

    for (var i = 0; i < x.length; i++) {
        var newCount = index.get(x[i]);
        if (newCount === undefined) {
            newCount = 1;
        } else {
            newCount++;
        }
        if (newCount > modeCount) {
            mode = x[i];
            modeCount = newCount;
        }
        index.set(x[i], newCount);
    }

    if (modeCount === 0) {
        throw new Error("mode requires at last one data point");
    }

    return mode;
}

/**
 * The min is the lowest number in the array.
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @throws {Error} if the length of x is less than one
 * @returns {number} minimum value
 * @example
 * min([1, 5, -10, 100, 2]); // => -10
 */
function min(x) {
    if (x.length === 0) {
        throw new Error("min requires at least one data point");
    }

    var value = x[0];
    for (var i = 1; i < x.length; i++) {
        if (x[i] < value) {
            value = x[i];
        }
    }
    return value;
}

/**
 * This computes the maximum number in an array.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @returns {number} maximum value
 * @throws {Error} if the length of x is less than one
 * @example
 * max([1, 2, 3, 4]);
 * // => 4
 */
function max(x) {
    if (x.length === 0) {
        throw new Error("max requires at least one data point");
    }

    var value = x[0];
    for (var i = 1; i < x.length; i++) {
        if (x[i] > value) {
            value = x[i];
        }
    }
    return value;
}

/**
 * This computes the minimum & maximum number in an array.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @returns {Array<number>} minimum & maximum value
 * @throws {Error} if the length of x is less than one
 * @example
 * extent([1, 2, 3, 4]);
 * // => [1, 4]
 */
function extent(x) {
    if (x.length === 0) {
        throw new Error("extent requires at least one data point");
    }

    var min = x[0];
    var max = x[0];
    for (var i = 1; i < x.length; i++) {
        if (x[i] > max) {
            max = x[i];
        }
        if (x[i] < min) {
            min = x[i];
        }
    }
    return [min, max];
}

/**
 * The minimum is the lowest number in the array. With a sorted array,
 * the first element in the array is always the smallest, so this calculation
 * can be done in one step, or constant time.
 *
 * @param {Array<number>} x input
 * @returns {number} minimum value
 * @example
 * minSorted([-100, -10, 1, 2, 5]); // => -100
 */
function minSorted(x) {
    return x[0];
}

/**
 * The maximum is the highest number in the array. With a sorted array,
 * the last element in the array is always the largest, so this calculation
 * can be done in one step, or constant time.
 *
 * @param {Array<number>} x input
 * @returns {number} maximum value
 * @example
 * maxSorted([-100, -10, 1, 2, 5]); // => 5
 */
function maxSorted(x) {
    return x[x.length - 1];
}

/**
 * The extent is the lowest & highest number in the array. With a sorted array,
 * the first element in the array is always the lowest while the last element is always the largest, so this calculation
 * can be done in one step, or constant time.
 *
 * @param {Array<number>} x input
 * @returns {Array<number>} minimum & maximum value
 * @example
 * extentSorted([-100, -10, 1, 2, 5]); // => [-100, 5]
 */
function extentSorted(x) {
    return [x[0], x[x.length - 1]];
}

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
    var value = 0;
    for (var i = 0; i < x.length; i++) {
        if (typeof x[i] !== "number") {
            return Number.NaN;
        }
        value += x[i];
    }
    return value;
}

/**
 * The [product](https://en.wikipedia.org/wiki/Product_(mathematics)) of an array
 * is the result of multiplying all numbers together, starting using one as the multiplicative identity.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x input
 * @return {number} product of all input numbers
 * @example
 * product([1, 2, 3, 4]); // => 24
 */
function product(x) {
    var value = 1;
    for (var i = 0; i < x.length; i++) {
        value *= x[i];
    }
    return value;
}

/**
 * This is the internal implementation of quantiles: when you know
 * that the order is sorted, you don't need to re-sort it, and the computations
 * are faster.
 *
 * @param {Array<number>} x sample of one or more data points
 * @param {number} p desired quantile: a number between 0 to 1, inclusive
 * @returns {number} quantile value
 * @throws {Error} if p ix outside of the range from 0 to 1
 * @throws {Error} if x is empty
 * @example
 * quantileSorted([3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20], 0.5); // => 9
 */
function quantileSorted(x, p) {
    var idx = x.length * p;
    if (x.length === 0) {
        throw new Error("quantile requires at least one data point.");
    } else if (p < 0 || p > 1) {
        throw new Error("quantiles must be between 0 and 1");
    } else if (p === 1) {
        // If p is 1, directly return the last element
        return x[x.length - 1];
    } else if (p === 0) {
        // If p is 0, directly return the first element
        return x[0];
    } else if (idx % 1 !== 0) {
        // If p is not integer, return the next element in array
        return x[Math.ceil(idx) - 1];
    } else if (x.length % 2 === 0) {
        // If the list has even-length, we'll take the average of this number
        // and the next value, if there is one
        return (x[idx - 1] + x[idx]) / 2;
    } else {
        // Finally, in the simple case of an integer value
        // with an odd-length list, return the x value at the index.
        return x[idx];
    }
}

/**
 * Rearrange items in `arr` so that all items in `[left, k]` range are the smallest.
 * The `k`-th element will have the `(k - left + 1)`-th smallest value in `[left, right]`.
 *
 * Implements Floyd-Rivest selection algorithm https://en.wikipedia.org/wiki/Floyd-Rivest_algorithm
 *
 * @param {Array<number>} arr input array
 * @param {number} k pivot index
 * @param {number} [left] left index
 * @param {number} [right] right index
 * @returns {void} mutates input array
 * @example
 * var arr = [65, 28, 59, 33, 21, 56, 22, 95, 50, 12, 90, 53, 28, 77, 39];
 * quickselect(arr, 8);
 * // = [39, 28, 28, 33, 21, 12, 22, 50, 53, 56, 59, 65, 90, 77, 95]
 */
function quickselect(arr, k, left, right) {
    left = left || 0;
    right = right || arr.length - 1;

    while (right > left) {
        // 600 and 0.5 are arbitrary constants chosen in the original paper to minimize execution time
        if (right - left > 600) {
            var n = right - left + 1;
            var m = k - left + 1;
            var z = Math.log(n);
            var s = 0.5 * Math.exp((2 * z) / 3);
            var sd = 0.5 * Math.sqrt((z * s * (n - s)) / n);
            if (m - n / 2 < 0) { sd *= -1; }
            var newLeft = Math.max(left, Math.floor(k - (m * s) / n + sd));
            var newRight = Math.min(
                right,
                Math.floor(k + ((n - m) * s) / n + sd)
            );
            quickselect(arr, k, newLeft, newRight);
        }

        var t = arr[k];
        var i = left;
        var j = right;

        swap(arr, left, k);
        if (arr[right] > t) { swap(arr, left, right); }

        while (i < j) {
            swap(arr, i, j);
            i++;
            j--;
            while (arr[i] < t) { i++; }
            while (arr[j] > t) { j--; }
        }

        if (arr[left] === t) { swap(arr, left, j); }
        else {
            j++;
            swap(arr, j, right);
        }

        if (j <= k) { left = j + 1; }
        if (k <= j) { right = j - 1; }
    }
}

function swap(arr, i, j) {
    var tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

/**
 * The [quantile](https://en.wikipedia.org/wiki/Quantile):
 * this is a population quantile, since we assume to know the entire
 * dataset in this library. This is an implementation of the
 * [Quantiles of a Population](http://en.wikipedia.org/wiki/Quantile#Quantiles_of_a_population)
 * algorithm from wikipedia.
 *
 * Sample is a one-dimensional array of numbers,
 * and p is either a decimal number from 0 to 1 or an array of decimal
 * numbers from 0 to 1.
 * In terms of a k/q quantile, p = k/q - it's just dealing with fractions or dealing
 * with decimal values.
 * When p is an array, the result of the function is also an array containing the appropriate
 * quantiles in input order
 *
 * @param {Array<number>} x sample of one or more numbers
 * @param {Array<number> | number} p the desired quantile, as a number between 0 and 1
 * @returns {number} quantile
 * @example
 * quantile([3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20], 0.5); // => 9
 */
function quantile(x, p) {
    var copy = x.slice();

    if (Array.isArray(p)) {
        // rearrange elements so that each element corresponding to a requested
        // quantile is on a place it would be if the array was fully sorted
        multiQuantileSelect(copy, p);
        // Initialize the result array
        var results = [];
        // For each requested quantile
        for (var i = 0; i < p.length; i++) {
            results[i] = quantileSorted(copy, p[i]);
        }
        return results;
    } else {
        var idx = quantileIndex(copy.length, p);
        quantileSelect(copy, idx, 0, copy.length - 1);
        return quantileSorted(copy, p);
    }
}

function quantileSelect(arr, k, left, right) {
    if (k % 1 === 0) {
        quickselect(arr, k, left, right);
    } else {
        k = Math.floor(k);
        quickselect(arr, k, left, right);
        quickselect(arr, k + 1, k + 1, right);
    }
}

function multiQuantileSelect(arr, p) {
    var indices = [0];
    for (var i = 0; i < p.length; i++) {
        indices.push(quantileIndex(arr.length, p[i]));
    }
    indices.push(arr.length - 1);
    indices.sort(compare);

    var stack = [0, indices.length - 1];

    while (stack.length) {
        var r = Math.ceil(stack.pop());
        var l = Math.floor(stack.pop());
        if (r - l <= 1) { continue; }

        var m = Math.floor((l + r) / 2);
        quantileSelect(
            arr,
            indices[m],
            Math.floor(indices[l]),
            Math.ceil(indices[r])
        );

        stack.push(l, m, m, r);
    }
}

function compare(a, b) {
    return a - b;
}

function quantileIndex(len, p) {
    var idx = len * p;
    if (p === 1) {
        // If p is 1, directly return the last index
        return len - 1;
    } else if (p === 0) {
        // If p is 0, directly return the first index
        return 0;
    } else if (idx % 1 !== 0) {
        // If index is not integer, return the next index in array
        return Math.ceil(idx) - 1;
    } else if (len % 2 === 0) {
        // If the list has even-length, we'll return the middle of two indices
        // around quantile to indicate that we need an average value of the two
        return idx - 0.5;
    } else {
        // Finally, in the simple case of an integer index
        // with an odd-length list, return the index
        return idx;
    }
}

/* eslint no-bitwise: 0 */

/**
 * This function returns the quantile in which one would find the given value in
 * the given array. With a sorted array, leveraging binary search, we can find
 * this information in logarithmic time.
 *
 * @param {Array<number>} x input
 * @returns {number} value value
 * @example
 * quantileRankSorted([1, 2, 3, 4], 3); // => 0.75
 * quantileRankSorted([1, 2, 3, 3, 4], 3); // => 0.7
 * quantileRankSorted([1, 2, 3, 4], 6); // => 1
 * quantileRankSorted([1, 2, 3, 3, 5], 4); // => 0.8
 */
function quantileRankSorted(x, value) {
    // Value is lesser than any value in the array
    if (value < x[0]) {
        return 0;
    }

    // Value is greater than any value in the array
    if (value > x[x.length - 1]) {
        return 1;
    }

    var l = lowerBound(x, value);

    // Value is not in the array
    if (x[l] !== value) {
        return l / x.length;
    }

    l++;

    var u = upperBound(x, value);

    // The value exists only once in the array
    if (u === l) {
        return l / x.length;
    }

    // Here, we are basically computing the mean of the range of indices
    // containing our searched value. But, instead, of initializing an
    // array and looping over it, there is a dedicated math formula that
    // we apply below to get the result.
    var r = u - l + 1;
    var sum = (r * (u + l)) / 2;
    var mean = sum / r;

    return mean / x.length;
}

function lowerBound(x, value) {
    var mid = 0;
    var lo = 0;
    var hi = x.length;

    while (lo < hi) {
        mid = (lo + hi) >>> 1;

        if (value <= x[mid]) {
            hi = mid;
        } else {
            lo = -~mid;
        }
    }

    return lo;
}

function upperBound(x, value) {
    var mid = 0;
    var lo = 0;
    var hi = x.length;

    while (lo < hi) {
        mid = (lo + hi) >>> 1;

        if (value >= x[mid]) {
            lo = -~mid;
        } else {
            hi = mid;
        }
    }

    return lo;
}

/**
 * This function returns the quantile in which one would find the given value in
 * the given array. It will copy and sort your array before each run, so
 * if you know your array is already sorted, you should use `quantileRankSorted`
 * instead.
 *
 * @param {Array<number>} x input
 * @returns {number} value value
 * @example
 * quantileRank([4, 3, 1, 2], 3); // => 0.75
 * quantileRank([4, 3, 2, 3, 1], 3); // => 0.7
 * quantileRank([2, 4, 1, 3], 6); // => 1
 * quantileRank([5, 3, 1, 2, 3], 4); // => 0.8
 */
function quantileRank(x, value) {
    // Cloning and sorting the array
    var sortedCopy = numericSort(x);

    return quantileRankSorted(sortedCopy, value);
}

/**
 * The [Interquartile range](http://en.wikipedia.org/wiki/Interquartile_range) is
 * a measure of statistical dispersion, or how scattered, spread, or
 * concentrated a distribution is. It's computed as the difference between
 * the third quartile and first quartile.
 *
 * @param {Array<number>} x sample of one or more numbers
 * @returns {number} interquartile range: the span between lower and upper quartile,
 * 0.25 and 0.75
 * @example
 * interquartileRange([0, 1, 2, 3]); // => 2
 */
function interquartileRange(x) {
    // Interquartile range is the span between the upper quartile,
    // at `0.75`, and lower quartile, `0.25`
    var q1 = quantile(x, 0.75);
    var q2 = quantile(x, 0.25);

    if (typeof q1 === "number" && typeof q2 === "number") {
        return q1 - q2;
    }
}

/**
 * The [median](http://en.wikipedia.org/wiki/Median) is
 * the middle number of a list. This is often a good indicator of 'the middle'
 * when there are outliers that skew the `mean()` value.
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * The median isn't necessarily one of the elements in the list: the value
 * can be the average of two elements if the list has an even length
 * and the two central values are different.
 *
 * @param {Array<number>} x input
 * @returns {number} median value
 * @example
 * median([10, 2, 5, 100, 2, 1]); // => 3.5
 */
function median(x) {
    return +quantile(x, 0.5);
}

/**
 * The [Median Absolute Deviation](http://en.wikipedia.org/wiki/Median_absolute_deviation) is
 * a robust measure of statistical
 * dispersion. It is more resilient to outliers than the standard deviation.
 *
 * @param {Array<number>} x input array
 * @returns {number} median absolute deviation
 * @example
 * medianAbsoluteDeviation([1, 1, 2, 2, 4, 6, 9]); // => 1
 */
function medianAbsoluteDeviation(x) {
    var medianValue = median(x);
    var medianAbsoluteDeviations = [];

    // Make a list of absolute deviations from the median
    for (var i = 0; i < x.length; i++) {
        medianAbsoluteDeviations.push(Math.abs(x[i] - medianValue));
    }

    // Find the median value of that list
    return median(medianAbsoluteDeviations);
}

/**
 * Split an array into chunks of a specified size. This function
 * has the same behavior as [PHP's array_chunk](http://php.net/manual/en/function.array-chunk.php)
 * function, and thus will insert smaller-sized chunks at the end if
 * the input size is not divisible by the chunk size.
 *
 * `x` is expected to be an array, and `chunkSize` a number.
 * The `x` array can contain any kind of data.
 *
 * @param {Array} x a sample
 * @param {number} chunkSize size of each output array. must be a positive integer
 * @returns {Array<Array>} a chunked array
 * @throws {Error} if chunk size is less than 1 or not an integer
 * @example
 * chunk([1, 2, 3, 4, 5, 6], 2);
 * // => [[1, 2], [3, 4], [5, 6]]
 */
function chunk(x, chunkSize) {
    // a list of result chunks, as arrays in an array
    var output = [];

    // `chunkSize` must be zero or higher - otherwise the loop below,
    // in which we call `start += chunkSize`, will loop infinitely.
    // So, we'll detect and throw in that case to indicate
    // invalid input.
    if (chunkSize < 1) {
        throw new Error("chunk size must be a positive number");
    }

    if (Math.floor(chunkSize) !== chunkSize) {
        throw new Error("chunk size must be an integer");
    }

    // `start` is the index at which `.slice` will start selecting
    // new array elements
    for (var start = 0; start < x.length; start += chunkSize) {
        // for each chunk, slice that part of the array and add it
        // to the output. The `.slice` function does not change
        // the original array.
        output.push(x.slice(start, start + chunkSize));
    }
    return output;
}

/**
 * Sampling with replacement is a type of sampling that allows the same
 * item to be picked out of a population more than once.
 *
 * @param {Array<*>} x an array of any kind of value
 * @param {number} n count of how many elements to take
 * @param {Function} [randomSource=Math.random] an optional entropy source that
 * returns numbers between 0 inclusive and 1 exclusive: the range [0, 1)
 * @return {Array} n sampled items from the population
 * @example
 * var values = [1, 2, 3, 4];
 * sampleWithReplacement(values, 2); // returns 2 random values, like [2, 4];
 */
function sampleWithReplacement(x, n, randomSource) {
    if (x.length === 0) {
        return [];
    }

    // a custom random number source can be provided if you want to use
    // a fixed seed or another random number generator, like
    // [random-js](https://www.npmjs.org/package/random-js)
    randomSource = randomSource || Math.random;

    var length = x.length;
    var sample = [];

    for (var i = 0; i < n; i++) {
        var index = Math.floor(randomSource() * length);

        sample.push(x[index]);
    }

    return sample;
}

/**
 * A [Fisher-Yates shuffle](http://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle)
 * in-place - which means that it **will change the order of the original
 * array by reference**.
 *
 * This is an algorithm that generates a random [permutation](https://en.wikipedia.org/wiki/Permutation)
 * of a set.
 *
 * @param {Array} x sample of one or more numbers
 * @param {Function} [randomSource=Math.random] an optional entropy source that
 * returns numbers between 0 inclusive and 1 exclusive: the range [0, 1)
 * @returns {Array} x
 * @example
 * var x = [1, 2, 3, 4];
 * shuffleInPlace(x);
 * // x is shuffled to a value like [2, 1, 4, 3]
 */
function shuffleInPlace(x, randomSource) {
    // a custom random number source can be provided if you want to use
    // a fixed seed or another random number generator, like
    // [random-js](https://www.npmjs.org/package/random-js)
    randomSource = randomSource || Math.random;

    // store the current length of the x to determine
    // when no elements remain to shuffle.
    var length = x.length;

    // temporary is used to hold an item when it is being
    // swapped between indices.
    var temporary;

    // The index to swap at each stage.
    var index;

    // While there are still items to shuffle
    while (length > 0) {
        // choose a random index within the subset of the array
        // that is not yet shuffled
        index = Math.floor(randomSource() * length--);

        // store the value that we'll move temporarily
        temporary = x[length];

        // swap the value at `x[length]` with `x[index]`
        x[length] = x[index];
        x[index] = temporary;
    }

    return x;
}

/**
 * A [Fisher-Yates shuffle](http://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle)
 * is a fast way to create a random permutation of a finite set. This is
 * a function around `shuffle_in_place` that adds the guarantee that
 * it will not modify its input.
 *
 * @param {Array} x sample of 0 or more numbers
 * @param {Function} [randomSource=Math.random] an optional entropy source that
 * returns numbers between 0 inclusive and 1 exclusive: the range [0, 1)
 * @return {Array} shuffled version of input
 * @example
 * var shuffled = shuffle([1, 2, 3, 4]);
 * shuffled; // = [2, 3, 1, 4] or any other random permutation
 */
function shuffle(x, randomSource) {
    // slice the original array so that it is not modified
    var sample = x.slice();

    // and then shuffle that shallow-copied array, in place
    return shuffleInPlace(sample, randomSource);
}

/**
 * Create a [simple random sample](http://en.wikipedia.org/wiki/Simple_random_sample)
 * from a given array of `n` elements.
 *
 * The sampled values will be in any order, not necessarily the order
 * they appear in the input.
 *
 * @param {Array<any>} x input array. can contain any type
 * @param {number} n count of how many elements to take
 * @param {Function} [randomSource=Math.random] an optional entropy source that
 * returns numbers between 0 inclusive and 1 exclusive: the range [0, 1)
 * @return {Array} subset of n elements in original array
 *
 * @example
 * var values = [1, 2, 4, 5, 6, 7, 8, 9];
 * sample(values, 3); // returns 3 random values, like [2, 5, 8];
 */
function sample(x, n, randomSource) {
    // shuffle the original array using a fisher-yates shuffle
    var shuffled = shuffle(x, randomSource);

    // and then return a subset of it - the first `n` elements.
    return shuffled.slice(0, n);
}

/**
 * Create a new column x row matrix.
 *
 * @private
 * @param {number} columns
 * @param {number} rows
 * @return {Array<Array<number>>} matrix
 * @example
 * makeMatrix(10, 10);
 */
function makeMatrix(columns, rows) {
    var matrix = [];
    for (var i = 0; i < columns; i++) {
        var column = [];
        for (var j = 0; j < rows; j++) {
            column.push(0);
        }
        matrix.push(column);
    }
    return matrix;
}

/**
 * For a sorted input, counting the number of unique values
 * is possible in constant time and constant memory. This is
 * a simple implementation of the algorithm.
 *
 * Values are compared with `===`, so objects and non-primitive objects
 * are not handled in any special way.
 *
 * @param {Array<*>} x an array of any kind of value
 * @returns {number} count of unique values
 * @example
 * uniqueCountSorted([1, 2, 3]); // => 3
 * uniqueCountSorted([1, 1, 1]); // => 1
 */
function uniqueCountSorted(x) {
    var uniqueValueCount = 0;
    var lastSeenValue;
    for (var i = 0; i < x.length; i++) {
        if (i === 0 || x[i] !== lastSeenValue) {
            lastSeenValue = x[i];
            uniqueValueCount++;
        }
    }
    return uniqueValueCount;
}

/**
 * Generates incrementally computed values based on the sums and sums of
 * squares for the data array
 *
 * @private
 * @param {number} j
 * @param {number} i
 * @param {Array<number>} sums
 * @param {Array<number>} sumsOfSquares
 * @return {number}
 * @example
 * ssq(0, 1, [-1, 0, 2], [1, 1, 5]);
 */
function ssq(j, i, sums, sumsOfSquares) {
    var sji; // s(j, i)
    if (j > 0) {
        var muji = (sums[i] - sums[j - 1]) / (i - j + 1); // mu(j, i)
        sji =
            sumsOfSquares[i] - sumsOfSquares[j - 1] - (i - j + 1) * muji * muji;
    } else {
        sji = sumsOfSquares[i] - (sums[i] * sums[i]) / (i + 1);
    }
    if (sji < 0) {
        return 0;
    }
    return sji;
}

/**
 * Function that recursively divides and conquers computations
 * for cluster j
 *
 * @private
 * @param {number} iMin Minimum index in cluster to be computed
 * @param {number} iMax Maximum index in cluster to be computed
 * @param {number} cluster Index of the cluster currently being computed
 * @param {Array<Array<number>>} matrix
 * @param {Array<Array<number>>} backtrackMatrix
 * @param {Array<number>} sums
 * @param {Array<number>} sumsOfSquares
 */
function fillMatrixColumn(
    iMin,
    iMax,
    cluster,
    matrix,
    backtrackMatrix,
    sums,
    sumsOfSquares
) {
    if (iMin > iMax) {
        return;
    }

    // Start at midpoint between iMin and iMax
    var i = Math.floor((iMin + iMax) / 2);

    matrix[cluster][i] = matrix[cluster - 1][i - 1];
    backtrackMatrix[cluster][i] = i;

    var jlow = cluster; // the lower end for j

    if (iMin > cluster) {
        jlow = Math.max(jlow, backtrackMatrix[cluster][iMin - 1] || 0);
    }
    jlow = Math.max(jlow, backtrackMatrix[cluster - 1][i] || 0);

    var jhigh = i - 1; // the upper end for j
    if (iMax < matrix[0].length - 1) {
        /* c8 ignore start */
        jhigh = Math.min(jhigh, backtrackMatrix[cluster][iMax + 1] || 0);
        /* c8 ignore end */
    }

    var sji;
    var sjlowi;
    var ssqjlow;
    var ssqj;
    for (var j = jhigh; j >= jlow; --j) {
        sji = ssq(j, i, sums, sumsOfSquares);

        if (sji + matrix[cluster - 1][jlow - 1] >= matrix[cluster][i]) {
            break;
        }

        // Examine the lower bound of the cluster border
        sjlowi = ssq(jlow, i, sums, sumsOfSquares);

        ssqjlow = sjlowi + matrix[cluster - 1][jlow - 1];

        if (ssqjlow < matrix[cluster][i]) {
            // Shrink the lower bound
            matrix[cluster][i] = ssqjlow;
            backtrackMatrix[cluster][i] = jlow;
        }
        jlow++;

        ssqj = sji + matrix[cluster - 1][j - 1];
        if (ssqj < matrix[cluster][i]) {
            matrix[cluster][i] = ssqj;
            backtrackMatrix[cluster][i] = j;
        }
    }

    fillMatrixColumn(
        iMin,
        i - 1,
        cluster,
        matrix,
        backtrackMatrix,
        sums,
        sumsOfSquares
    );
    fillMatrixColumn(
        i + 1,
        iMax,
        cluster,
        matrix,
        backtrackMatrix,
        sums,
        sumsOfSquares
    );
}

/**
 * Initializes the main matrices used in Ckmeans and kicks
 * off the divide and conquer cluster computation strategy
 *
 * @private
 * @param {Array<number>} data sorted array of values
 * @param {Array<Array<number>>} matrix
 * @param {Array<Array<number>>} backtrackMatrix
 */
function fillMatrices(data, matrix, backtrackMatrix) {
    var nValues = matrix[0].length;

    // Shift values by the median to improve numeric stability
    var shift = data[Math.floor(nValues / 2)];

    // Cumulative sum and cumulative sum of squares for all values in data array
    var sums = [];
    var sumsOfSquares = [];

    // Initialize first column in matrix & backtrackMatrix
    for (var i = 0, shiftedValue = (void 0); i < nValues; ++i) {
        shiftedValue = data[i] - shift;
        if (i === 0) {
            sums.push(shiftedValue);
            sumsOfSquares.push(shiftedValue * shiftedValue);
        } else {
            sums.push(sums[i - 1] + shiftedValue);
            sumsOfSquares.push(
                sumsOfSquares[i - 1] + shiftedValue * shiftedValue
            );
        }

        // Initialize for cluster = 0
        matrix[0][i] = ssq(0, i, sums, sumsOfSquares);
        backtrackMatrix[0][i] = 0;
    }

    // Initialize the rest of the columns
    var iMin;
    for (var cluster = 1; cluster < matrix.length; ++cluster) {
        if (cluster < matrix.length - 1) {
            iMin = cluster;
        } else {
            // No need to compute matrix[K-1][0] ... matrix[K-1][N-2]
            iMin = nValues - 1;
        }

        fillMatrixColumn(
            iMin,
            nValues - 1,
            cluster,
            matrix,
            backtrackMatrix,
            sums,
            sumsOfSquares
        );
    }
}

/**
 * Ckmeans clustering is an improvement on heuristic-based clustering
 * approaches like Jenks. The algorithm was developed in
 * [Haizhou Wang and Mingzhou Song](http://journal.r-project.org/archive/2011-2/RJournal_2011-2_Wang+Song.pdf)
 * as a [dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming) approach
 * to the problem of clustering numeric data into groups with the least
 * within-group sum-of-squared-deviations.
 *
 * Minimizing the difference within groups - what Wang & Song refer to as
 * `withinss`, or within sum-of-squares, means that groups are optimally
 * homogenous within and the data is split into representative groups.
 * This is very useful for visualization, where you may want to represent
 * a continuous variable in discrete color or style groups. This function
 * can provide groups that emphasize differences between data.
 *
 * Being a dynamic approach, this algorithm is based on two matrices that
 * store incrementally-computed values for squared deviations and backtracking
 * indexes.
 *
 * This implementation is based on Ckmeans 3.4.6, which introduced a new divide
 * and conquer approach that improved runtime from O(kn^2) to O(kn log(n)).
 *
 * Unlike the [original implementation](https://cran.r-project.org/web/packages/Ckmeans.1d.dp/index.html),
 * this implementation does not include any code to automatically determine
 * the optimal number of clusters: this information needs to be explicitly
 * provided.
 *
 * ### References
 * _Ckmeans.1d.dp: Optimal k-means Clustering in One Dimension by Dynamic
 * Programming_ Haizhou Wang and Mingzhou Song ISSN 2073-4859
 *
 * from The R Journal Vol. 3/2, December 2011
 * @param {Array<number>} x input data, as an array of number values
 * @param {number} nClusters number of desired classes. This cannot be
 * greater than the number of values in the data array.
 * @returns {Array<Array<number>>} clustered input
 * @throws {Error} if the number of requested clusters is higher than the size of the data
 * @example
 * ckmeans([-1, 2, -1, 2, 4, 5, 6, -1, 2, -1], 3);
 * // The input, clustered into groups of similar numbers.
 * //= [[-1, -1, -1, -1], [2, 2, 2], [4, 5, 6]]);
 */
function ckmeans(x, nClusters) {
    if (nClusters > x.length) {
        throw new Error(
            "cannot generate more classes than there are data values"
        );
    }

    var sorted = numericSort(x);
    // we'll use this as the maximum number of clusters
    var uniqueCount = uniqueCountSorted(sorted);

    // if all of the input values are identical, there's one cluster
    // with all of the input in it.
    if (uniqueCount === 1) {
        return [sorted];
    }

    // named 'S' originally
    var matrix = makeMatrix(nClusters, sorted.length);
    // named 'J' originally
    var backtrackMatrix = makeMatrix(nClusters, sorted.length);

    // This is a dynamic programming way to solve the problem of minimizing
    // within-cluster sum of squares. It's similar to linear regression
    // in this way, and this calculation incrementally computes the
    // sum of squares that are later read.
    fillMatrices(sorted, matrix, backtrackMatrix);

    // The real work of Ckmeans clustering happens in the matrix generation:
    // the generated matrices encode all possible clustering combinations, and
    // once they're generated we can solve for the best clustering groups
    // very quickly.
    var clusters = [];
    var clusterRight = backtrackMatrix[0].length - 1;

    // Backtrack the clusters from the dynamic programming matrix. This
    // starts at the bottom-right corner of the matrix (if the top-left is 0, 0),
    // and moves the cluster target with the loop.
    for (var cluster = backtrackMatrix.length - 1; cluster >= 0; cluster--) {
        var clusterLeft = backtrackMatrix[cluster][clusterRight];

        // fill the cluster from the sorted input by taking a slice of the
        // array. the backtrack matrix makes this easy - it stores the
        // indexes where the cluster should start and end.
        clusters[cluster] = sorted.slice(clusterLeft, clusterRight + 1);

        if (cluster > 0) {
            clusterRight = clusterLeft - 1;
        }
    }

    return clusters;
}

/*
 * Pull Breaks Values for Jenks
 *
 * the second part of the jenks recipe: take the calculated matrices
 * and derive an array of n breaks.
 *
 * @private
 */
function jenksBreaks(data, lowerClassLimits, nClasses) {
    var k = data.length;
    var kclass = [];
    var countNum = nClasses;

    // the calculation of classes will never include the upper
    // bound, so we need to explicitly set it
    kclass[nClasses] = data[data.length - 1];

    // the lowerClassLimits matrix is used as indices into itself
    // here: the `k` variable is reused in each iteration.
    while (countNum > 0) {
        kclass[countNum - 1] = data[lowerClassLimits[k][countNum] - 1];
        k = lowerClassLimits[k][countNum] - 1;
        countNum--;
    }

    return kclass;
}

/*
 * Compute Matrices for Jenks
 *
 * Compute the matrices required for Jenks breaks. These matrices
 * can be used for any classing of data with `classes <= nClasses`
 *
 * @private
 */
function jenksMatrices(data, nClasses) {
    // in the original implementation, these matrices are referred to
    // as `LC` and `OP`
    //
    // * lowerClassLimits (LC): optimal lower class limits
    // * varianceCombinations (OP): optimal variance combinations for all classes
    var lowerClassLimits = [];
    var varianceCombinations = [];
    // loop counters
    var i;
    var j;
    // the variance, as computed at each step in the calculation
    var variance = 0;

    // Initialize and fill each matrix with zeroes
    for (i = 0; i < data.length + 1; i++) {
        var tmp1 = [];
        var tmp2 = [];
        // despite these arrays having the same values, we need
        // to keep them separate so that changing one does not change
        // the other
        for (j = 0; j < nClasses + 1; j++) {
            tmp1.push(0);
            tmp2.push(0);
        }
        lowerClassLimits.push(tmp1);
        varianceCombinations.push(tmp2);
    }

    for (i = 1; i < nClasses + 1; i++) {
        lowerClassLimits[1][i] = 1;
        varianceCombinations[1][i] = 0;
        // in the original implementation, 9999999 is used but
        // since Javascript has `Infinity`, we use that.
        for (j = 2; j < data.length + 1; j++) {
            varianceCombinations[j][i] = Number.POSITIVE_INFINITY;
        }
    }

    for (var l = 2; l < data.length + 1; l++) {
        // `SZ` originally. this is the sum of the values seen thus
        // far when calculating variance.
        var sum = 0;
        // `ZSQ` originally. the sum of squares of values seen
        // thus far
        var sumSquares = 0;
        // `WT` originally. This is the number of
        var w = 0;
        // `IV` originally
        var i4 = 0;

        // in several instances, you could say `Math.pow(x, 2)`
        // instead of `x * x`, but this is slower in some browsers
        // introduces an unnecessary concept.
        for (var m = 1; m < l + 1; m++) {
            // `III` originally
            var lowerClassLimit = l - m + 1;
            var val = data[lowerClassLimit - 1];

            // here we're estimating variance for each potential classing
            // of the data, for each potential number of classes. `w`
            // is the number of data points considered so far.
            w++;

            // increase the current sum and sum-of-squares
            sum += val;
            sumSquares += val * val;

            // the variance at this point in the sequence is the difference
            // between the sum of squares and the total x 2, over the number
            // of samples.
            variance = sumSquares - (sum * sum) / w;

            i4 = lowerClassLimit - 1;

            if (i4 !== 0) {
                for (j = 2; j < nClasses + 1; j++) {
                    // if adding this element to an existing class
                    // will increase its variance beyond the limit, break
                    // the class at this point, setting the `lowerClassLimit`
                    // at this point.
                    if (
                        varianceCombinations[l][j] >=
                        variance + varianceCombinations[i4][j - 1]
                    ) {
                        lowerClassLimits[l][j] = lowerClassLimit;
                        varianceCombinations[l][j] =
                            variance + varianceCombinations[i4][j - 1];
                    }
                }
            }
        }

        lowerClassLimits[l][1] = 1;
        varianceCombinations[l][1] = variance;
    }

    // return the two matrices. for just providing breaks, only
    // `lowerClassLimits` is needed, but variances can be useful to
    // evaluate goodness of fit.
    return {
        lowerClassLimits: lowerClassLimits,
        varianceCombinations: varianceCombinations
    };
}

/**
 * The **[jenks natural breaks optimization](http://en.wikipedia.org/wiki/Jenks_natural_breaks_optimization)**
 * is an algorithm commonly used in cartography and visualization to decide
 * upon groupings of data values that minimize variance within themselves
 * and maximize variation between themselves.
 *
 * For instance, cartographers often use jenks in order to choose which
 * values are assigned to which colors in a [choropleth](https://en.wikipedia.org/wiki/Choropleth_map)
 * map.
 *
 * @param {Array<number>} data input data, as an array of number values
 * @param {number} nClasses number of desired classes
 * @returns {Array<number>} array of class break positions
 * // split data into 3 break points
 * jenks([1, 2, 4, 5, 7, 9, 10, 20], 3) // = [1, 7, 20, 20]
 */
function jenks(data, nClasses) {
    if (nClasses > data.length) {
        return null;
    }

    // sort data in numerical order, since this is expected
    // by the matrices function
    data = data.slice().sort(function (a, b) {
        return a - b;
    });

    // get our basic matrices
    var matrices = jenksMatrices(data, nClasses);
    // we only need lower class limits here
    var lowerClassLimits = matrices.lowerClassLimits;

    // extract nClasses out of the computed matrices
    return jenksBreaks(data, lowerClassLimits, nClasses);
}

/**
 * Given an array of x, this will find the extent of the
 * x and return an array of breaks that can be used
 * to categorize the x into a number of classes. The
 * returned array will always be 1 longer than the number of
 * classes because it includes the minimum value.
 *
 * @param {Array<number>} x an array of number values
 * @param {number} nClasses number of desired classes
 * @returns {Array<number>} array of class break positions
 * @example
 * equalIntervalBreaks([1, 2, 3, 4, 5, 6], 4); // => [1, 2.25, 3.5, 4.75, 6]
 */
function equalIntervalBreaks(x, nClasses) {
    if (x.length < 2) {
        return x;
    }

    var theMin = min(x);
    var theMax = max(x);

    // the first break will always be the minimum value
    // in the xset
    var breaks = [theMin];

    // The size of each break is the full range of the x
    // divided by the number of classes requested
    var breakSize = (theMax - theMin) / nClasses;

    // In the case of nClasses = 1, this loop won't run
    // and the returned breaks will be [min, max]
    for (var i = 1; i < nClasses; i++) {
        breaks.push(breaks[0] + breakSize * i);
    }

    // the last break will always be the
    // maximum.
    breaks.push(theMax);

    return breaks;
}

/**
 * [Sample covariance](https://en.wikipedia.org/wiki/Sample_mean_and_covariance) of two datasets:
 * how much do the two datasets move together?
 * x and y are two datasets, represented as arrays of numbers.
 *
 * @param {Array<number>} x a sample of two or more data points
 * @param {Array<number>} y a sample of two or more data points
 * @throws {Error} if x and y do not have equal lengths
 * @throws {Error} if x or y have length of one or less
 * @returns {number} sample covariance
 * @example
 * sampleCovariance([1, 2, 3, 4, 5, 6], [6, 5, 4, 3, 2, 1]); // => -3.5
 */
function sampleCovariance(x, y) {
    // The two datasets must have the same length which must be more than 1
    if (x.length !== y.length) {
        throw new Error("sampleCovariance requires samples with equal lengths");
    }

    if (x.length < 2) {
        throw new Error(
            "sampleCovariance requires at least two data points in each sample"
        );
    }

    // determine the mean of each dataset so that we can judge each
    // value of the dataset fairly as the difference from the mean. this
    // way, if one dataset is [1, 2, 3] and [2, 3, 4], their covariance
    // does not suffer because of the difference in absolute values
    var xmean = mean(x);
    var ymean = mean(y);
    var sum = 0;

    // for each pair of values, the covariance increases when their
    // difference from the mean is associated - if both are well above
    // or if both are well below
    // the mean, the covariance increases significantly.
    for (var i = 0; i < x.length; i++) {
        sum += (x[i] - xmean) * (y[i] - ymean);
    }

    // this is Bessels' Correction: an adjustment made to sample statistics
    // that allows for the reduced degree of freedom entailed in calculating
    // values from samples rather than complete populations.
    var besselsCorrection = x.length - 1;

    // the covariance is weighted by the length of the datasets.
    return sum / besselsCorrection;
}

/**
 * The [sample variance](https://en.wikipedia.org/wiki/Variance#Sample_variance)
 * is the sum of squared deviations from the mean. The sample variance
 * is distinguished from the variance by the usage of [Bessel's Correction](https://en.wikipedia.org/wiki/Bessel's_correction):
 * instead of dividing the sum of squared deviations by the length of the input,
 * it is divided by the length minus one. This corrects the bias in estimating
 * a value from a set that you don't know if full.
 *
 * References:
 * * [Wolfram MathWorld on Sample Variance](http://mathworld.wolfram.com/SampleVariance.html)
 *
 * @param {Array<number>} x a sample of two or more data points
 * @throws {Error} if the length of x is less than 2
 * @return {number} sample variance
 * @example
 * sampleVariance([1, 2, 3, 4, 5]); // => 2.5
 */
function sampleVariance(x) {
    if (x.length < 2) {
        throw new Error("sampleVariance requires at least two data points");
    }

    var sumSquaredDeviationsValue = sumNthPowerDeviations(x, 2);

    // this is Bessels' Correction: an adjustment made to sample statistics
    // that allows for the reduced degree of freedom entailed in calculating
    // values from samples rather than complete populations.
    var besselsCorrection = x.length - 1;

    // Find the mean value of that list
    return sumSquaredDeviationsValue / besselsCorrection;
}

/**
 * The [sample standard deviation](http://en.wikipedia.org/wiki/Standard_deviation#Sample_standard_deviation)
 * is the square root of the sample variance.
 *
 * @param {Array<number>} x input array
 * @returns {number} sample standard deviation
 * @example
 * sampleStandardDeviation([2, 4, 4, 4, 5, 5, 7, 9]).toFixed(2);
 * // => '2.14'
 */
function sampleStandardDeviation(x) {
    var sampleVarianceX = sampleVariance(x);
    return Math.sqrt(sampleVarianceX);
}

/**
 * The [correlation](http://en.wikipedia.org/wiki/Correlation_and_dependence) is
 * a measure of how correlated two datasets are, between -1 and 1
 *
 * @param {Array<number>} x first input
 * @param {Array<number>} y second input
 * @returns {number} sample correlation
 * @example
 * sampleCorrelation([1, 2, 3, 4, 5, 6], [2, 2, 3, 4, 5, 60]).toFixed(2);
 * // => '0.69'
 */
function sampleCorrelation(x, y) {
    var cov = sampleCovariance(x, y);
    var xstd = sampleStandardDeviation(x);
    var ystd = sampleStandardDeviation(y);

    return cov / xstd / ystd;
}

/**
 * The [rank correlation](https://en.wikipedia.org/wiki/Rank_correlation) is
 * a measure of the strength of monotonic relationship between two arrays
 *
 * @param {Array<number>} x first input
 * @param {Array<number>} y second input
 * @returns {number} sample rank correlation
 */
function sampleRankCorrelation(x, y) {
    var xIndexes = x
        .map(function (value, index) { return [value, index]; })
        .sort(function (a, b) { return a[0] - b[0]; })
        .map(function (pair) { return pair[1]; });
    var yIndexes = y
        .map(function (value, index) { return [value, index]; })
        .sort(function (a, b) { return a[0] - b[0]; })
        .map(function (pair) { return pair[1]; });

    // At this step, we have an array of indexes
    // that map from sorted numbers to their original indexes. We reverse
    // that so that it is an array of the sorted destination index.
    var xRanks = Array(xIndexes.length);
    var yRanks = Array(xIndexes.length);
    for (var i = 0; i < xIndexes.length; i++) {
        xRanks[xIndexes[i]] = i;
        yRanks[yIndexes[i]] = i;
    }

    return sampleCorrelation(xRanks, yRanks);
}

/**
 * [Skewness](http://en.wikipedia.org/wiki/Skewness) is
 * a measure of the extent to which a probability distribution of a
 * real-valued random variable "leans" to one side of the mean.
 * The skewness value can be positive or negative, or even undefined.
 *
 * Implementation is based on the adjusted Fisher-Pearson standardized
 * moment coefficient, which is the version found in Excel and several
 * statistical packages including Minitab, SAS and SPSS.
 *
 * @since 4.1.0
 * @param {Array<number>} x a sample of 3 or more data points
 * @returns {number} sample skewness
 * @throws {Error} if x has length less than 3
 * @example
 * sampleSkewness([2, 4, 6, 3, 1]); // => 0.590128656384365
 */
function sampleSkewness(x) {
    if (x.length < 3) {
        throw new Error("sampleSkewness requires at least three data points");
    }

    var meanValue = mean(x);
    var tempValue;
    var sumSquaredDeviations = 0;
    var sumCubedDeviations = 0;

    for (var i = 0; i < x.length; i++) {
        tempValue = x[i] - meanValue;
        sumSquaredDeviations += tempValue * tempValue;
        sumCubedDeviations += tempValue * tempValue * tempValue;
    }

    // this is Bessels' Correction: an adjustment made to sample statistics
    // that allows for the reduced degree of freedom entailed in calculating
    // values from samples rather than complete populations.
    var besselsCorrection = x.length - 1;

    // Find the mean value of that list
    var theSampleStandardDeviation = Math.sqrt(
        sumSquaredDeviations / besselsCorrection
    );

    var n = x.length;
    var cubedS = Math.pow(theSampleStandardDeviation, 3);

    return (n * sumCubedDeviations) / ((n - 1) * (n - 2) * cubedS);
}

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
    var n = x.length;

    if (n < 4) {
        throw new Error("sampleKurtosis requires at least four data points");
    }

    var meanValue = mean(x);
    var tempValue;
    var secondCentralMoment = 0;
    var fourthCentralMoment = 0;

    for (var i = 0; i < n; i++) {
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

/**
 * Implementation of [Heap's Algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm)
 * for generating permutations.
 *
 * @param {Array} elements any type of data
 * @returns {Array<Array>} array of permutations
 */
function permutationsHeap(elements) {
    var indexes = new Array(elements.length);
    var permutations = [elements.slice()];

    for (var i = 0; i < elements.length; i++) {
        indexes[i] = 0;
    }

    for (var i$1 = 0; i$1 < elements.length; ) {
        if (indexes[i$1] < i$1) {
            // At odd indexes, swap from indexes[i] instead
            // of from the beginning of the array
            var swapFrom = 0;
            if (i$1 % 2 !== 0) {
                swapFrom = indexes[i$1];
            }

            // swap between swapFrom and i, using
            // a temporary variable as storage.
            var temp = elements[swapFrom];
            elements[swapFrom] = elements[i$1];
            elements[i$1] = temp;

            permutations.push(elements.slice());
            indexes[i$1]++;
            i$1 = 0;
        } else {
            indexes[i$1] = 0;
            i$1++;
        }
    }

    return permutations;
}

/**
 * Implementation of Combinations
 * Combinations are unique subsets of a collection - in this case, k x from a collection at a time.
 * https://en.wikipedia.org/wiki/Combination
 * @param {Array} x any type of data
 * @param {int} k the number of objects in each group (without replacement)
 * @returns {Array<Array>} array of permutations
 * @example
 * combinations([1, 2, 3], 2); // => [[1,2], [1,3], [2,3]]
 */

function combinations(x, k) {
    var i;
    var subI;
    var combinationList = [];
    var subsetCombinations;
    var next;

    for (i = 0; i < x.length; i++) {
        if (k === 1) {
            combinationList.push([x[i]]);
        } else {
            subsetCombinations = combinations(x.slice(i + 1, x.length), k - 1);
            for (subI = 0; subI < subsetCombinations.length; subI++) {
                next = subsetCombinations[subI];
                next.unshift(x[i]);
                combinationList.push(next);
            }
        }
    }
    return combinationList;
}

/**
 * Implementation of [Combinations](https://en.wikipedia.org/wiki/Combination) with replacement
 * Combinations are unique subsets of a collection - in this case, k x from a collection at a time.
 * 'With replacement' means that a given element can be chosen multiple times.
 * Unlike permutation, order doesn't matter for combinations.
 *
 * @param {Array} x any type of data
 * @param {int} k the number of objects in each group (without replacement)
 * @returns {Array<Array>} array of permutations
 * @example
 * combinationsReplacement([1, 2], 2); // => [[1, 1], [1, 2], [2, 2]]
 */
function combinationsReplacement(x, k) {
    var combinationList = [];

    for (var i = 0; i < x.length; i++) {
        if (k === 1) {
            // If we're requested to find only one element, we don't need
            // to recurse: just push `x[i]` onto the list of combinations.
            combinationList.push([x[i]]);
        } else {
            // Otherwise, recursively find combinations, given `k - 1`. Note that
            // we request `k - 1`, so if you were looking for k=3 combinations, we're
            // requesting k=2. This -1 gets reversed in the for loop right after this
            // code, since we concatenate `x[i]` onto the selected combinations,
            // bringing `k` back up to your requested level.
            // This recursion may go many levels deep, since it only stops once
            // k=1.
            var subsetCombinations = combinationsReplacement(
                x.slice(i, x.length),
                k - 1
            );

            for (var j = 0; j < subsetCombinations.length; j++) {
                combinationList.push([x[i]].concat(subsetCombinations[j]));
            }
        }
    }

    return combinationList;
}

/**
 * When adding a new value to a list, one does not have to necessary
 * recompute the mean of the list in linear time. They can instead use
 * this function to compute the new mean by providing the current mean,
 * the number of elements in the list that produced it and the new
 * value to add.
 *
 * @since 2.5.0
 * @param {number} mean current mean
 * @param {number} n number of items in the list
 * @param {number} newValue the added value
 * @returns {number} the new mean
 *
 * @example
 * addToMean(14, 5, 53); // => 20.5
 */
function addToMean(mean, n, newValue) {
    return mean + (newValue - mean) / (n + 1);
}

/**
 * When combining two lists of values for which one already knows the means,
 * one does not have to necessary recompute the mean of the combined lists in
 * linear time. They can instead use this function to compute the combined
 * mean by providing the mean & number of values of the first list and the mean
 * & number of values of the second list.
 *
 * @since 3.0.0
 * @param {number} mean1 mean of the first list
 * @param {number} n1 number of items in the first list
 * @param {number} mean2 mean of the second list
 * @param {number} n2 number of items in the second list
 * @returns {number} the combined mean
 *
 * @example
 * combineMeans(5, 3, 4, 3); // => 4.5
 */
function combineMeans(mean1, n1, mean2, n2) {
    return (mean1 * n1 + mean2 * n2) / (n1 + n2);
}

/**
 * When combining two lists of values for which one already knows the variances,
 * one does not have to necessary recompute the variance of the combined lists
 * in linear time. They can instead use this function to compute the combined
 * variance by providing the variance, mean & number of values of the first list
 * and the variance, mean & number of values of the second list.
 *
 * @since 3.0.0
 * @param {number} variance1 variance of the first list
 * @param {number} mean1 mean of the first list
 * @param {number} n1 number of items in the first list
 * @param {number} variance2 variance of the second list
 * @param {number} mean2 mean of the second list
 * @param {number} n2 number of items in the second list
 * @returns {number} the combined mean
 *
 * @example
 * combineVariances(14 / 3, 5, 3, 8 / 3, 4, 3); // => 47 / 12
 */
function combineVariances(variance1, mean1, n1, variance2, mean2, n2) {
    var newMean = combineMeans(mean1, n1, mean2, n2);

    return (
        (n1 * (variance1 + Math.pow(mean1 - newMean, 2)) +
            n2 * (variance2 + Math.pow(mean2 - newMean, 2))) /
        (n1 + n2)
    );
}

/**
 * The [Geometric Mean](https://en.wikipedia.org/wiki/Geometric_mean) is
 * a mean function that is more useful for numbers in different
 * ranges.
 *
 * This is the nth root of the input numbers multiplied by each other.
 *
 * The geometric mean is often useful for
 * **[proportional growth](https://en.wikipedia.org/wiki/Geometric_mean#Proportional_growth)**: given
 * growth rates for multiple years, like _80%, 16.66% and 42.85%_, a simple
 * mean will incorrectly estimate an average growth rate, whereas a geometric
 * mean will correctly estimate a growth rate that, over those years,
 * will yield the same end value.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @returns {number} geometric mean
 * @throws {Error} if x is empty
 * @throws {Error} if x contains a negative number
 * @example
 * var growthRates = [1.80, 1.166666, 1.428571];
 * var averageGrowth = ss.geometricMean(growthRates);
 * var averageGrowthRates = [averageGrowth, averageGrowth, averageGrowth];
 * var startingValue = 10;
 * var startingValueMean = 10;
 * growthRates.forEach(function(rate) {
 *   startingValue *= rate;
 * });
 * averageGrowthRates.forEach(function(rate) {
 *   startingValueMean *= rate;
 * });
 * startingValueMean === startingValue;
 */
function geometricMean(x) {
    if (x.length === 0) {
        throw new Error("geometricMean requires at least one data point");
    }

    // the starting value.
    var value = 1;

    for (var i = 0; i < x.length; i++) {
        // the geometric mean is only valid for positive numbers
        if (x[i] < 0) {
            throw new Error(
                "geometricMean requires only non-negative numbers as input"
            );
        }

        // repeatedly multiply the value by each number
        value *= x[i];
    }

    return Math.pow(value, 1 / x.length);
}

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

    var value = 0;
    for (var i = 0; i < x.length; i++) {
        if (x[i] < 0) {
            throw new Error(
                "logAverage requires only non-negative numbers as input"
            );
        }
        value += Math.log(x[i]);
    }

    return Math.exp(value / x.length);
}

/**
 * The [Harmonic Mean](https://en.wikipedia.org/wiki/Harmonic_mean) is
 * a mean function typically used to find the average of rates.
 * This mean is calculated by taking the reciprocal of the arithmetic mean
 * of the reciprocals of the input numbers.
 *
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x sample of one or more data points
 * @returns {number} harmonic mean
 * @throws {Error} if x is empty
 * @throws {Error} if x contains a negative number
 * @example
 * harmonicMean([2, 3]).toFixed(2) // => '2.40'
 */
function harmonicMean(x) {
    if (x.length === 0) {
        throw new Error("harmonicMean requires at least one data point");
    }

    var reciprocalSum = 0;

    for (var i = 0; i < x.length; i++) {
        // the harmonic mean is only valid for positive numbers
        if (x[i] <= 0) {
            throw new Error(
                "harmonicMean requires only positive numbers as input"
            );
        }

        reciprocalSum += 1 / x[i];
    }

    // divide n by the reciprocal sum
    return x.length / reciprocalSum;
}

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

/**
 * The [median](http://en.wikipedia.org/wiki/Median) is
 * the middle number of a list. This is often a good indicator of 'the middle'
 * when there are outliers that skew the `mean()` value.
 * This is a [measure of central tendency](https://en.wikipedia.org/wiki/Central_tendency):
 * a method of finding a typical or central value of a set of numbers.
 *
 * The median isn't necessarily one of the elements in the list: the value
 * can be the average of two elements if the list has an even length
 * and the two central values are different.
 *
 * @param {Array<number>} sorted input
 * @returns {number} median value
 * @example
 * medianSorted([10, 2, 5, 100, 2, 1]); // => 52.5
 */
function medianSorted(sorted) {
    return quantileSorted(sorted, 0.5);
}

/**
 * When removing a value from a list, one does not have to necessary
 * recompute the mean of the list in linear time. They can instead use
 * this function to compute the new mean by providing the current mean,
 * the number of elements in the list that produced it and the value to remove.
 *
 * @since 3.0.0
 * @param {number} mean current mean
 * @param {number} n number of items in the list
 * @param {number} value the value to remove
 * @returns {number} the new mean
 *
 * @example
 * subtractFromMean(20.5, 6, 53); // => 14
 */
function subtractFromMean(mean, n, value) {
    return (mean * n - value) / (n - 1);
}

/**
 * The Root Mean Square (RMS) is
 * a mean function used as a measure of the magnitude of a set
 * of numbers, regardless of their sign.
 * This is the square root of the mean of the squares of the
 * input numbers.
 * This runs in `O(n)`, linear time, with respect to the length of the array.
 *
 * @param {Array<number>} x a sample of one or more data points
 * @returns {number} root mean square
 * @throws {Error} if x is empty
 * @example
 * rootMeanSquare([-1, 1, -1, 1]); // => 1
 */
function rootMeanSquare(x) {
    if (x.length === 0) {
        throw new Error("rootMeanSquare requires at least one data point");
    }

    var sumOfSquares = 0;
    for (var i = 0; i < x.length; i++) {
        sumOfSquares += Math.pow(x[i], 2);
    }

    return Math.sqrt(sumOfSquares / x.length);
}

/**
 * The`coefficient of variation`_ is the ratio of the standard deviation to the mean.
 * .._`coefficient of variation`: https://en.wikipedia.org/wiki/Coefficient_of_variation
 *
 *
 * @param {Array} x input
 * @returns {number} coefficient of variation
 * @example
 * coefficientOfVariation([1, 2, 3, 4]).toFixed(3); // => 0.516
 * coefficientOfVariation([1, 2, 3, 4, 5]).toFixed(3); // => 0.527
 * coefficientOfVariation([-1, 0, 1, 2, 3, 4]).toFixed(3); // => 1.247
 */
function coefficientOfVariation(x) {
    return sampleStandardDeviation(x) / mean(x);
}

/**
 * This is to compute [a one-sample t-test](https://en.wikipedia.org/wiki/Student%27s_t-test#One-sample_t-test), comparing the mean
 * of a sample to a known value, x.
 *
 * in this case, we're trying to determine whether the
 * population mean is equal to the value that we know, which is `x`
 * here. Usually the results here are used to look up a
 * [p-value](http://en.wikipedia.org/wiki/P-value), which, for
 * a certain level of significance, will let you determine that the
 * null hypothesis can or cannot be rejected.
 *
 * @param {Array<number>} x sample of one or more numbers
 * @param {number} expectedValue expected value of the population mean
 * @returns {number} value
 * @example
 * tTest([1, 2, 3, 4, 5, 6], 3.385).toFixed(2); // => '0.16'
 */
function tTest(x, expectedValue) {
    // The mean of the sample
    var sampleMean = mean(x);

    // The standard deviation of the sample
    var sd = standardDeviation(x);

    // Square root the length of the sample
    var rootN = Math.sqrt(x.length);

    // returning the t value
    return (sampleMean - expectedValue) / (sd / rootN);
}

/**
 * This is to compute [two sample t-test](http://en.wikipedia.org/wiki/Student's_t-test).
 * Tests whether "mean(X)-mean(Y) = difference", (
 * in the most common case, we often have `difference == 0` to test if two samples
 * are likely to be taken from populations with the same mean value) with
 * no prior knowledge on standard deviations of both samples
 * other than the fact that they have the same standard deviation.
 *
 * Usually the results here are used to look up a
 * [p-value](http://en.wikipedia.org/wiki/P-value), which, for
 * a certain level of significance, will let you determine that the
 * null hypothesis can or cannot be rejected.
 *
 * `diff` can be omitted if it equals 0.
 *
 * [This is used to reject](https://en.wikipedia.org/wiki/Exclusion_of_the_null_hypothesis)
 * a null hypothesis that the two populations that have been sampled into
 * `sampleX` and `sampleY` are equal to each other.
 *
 * @param {Array<number>} sampleX a sample as an array of numbers
 * @param {Array<number>} sampleY a sample as an array of numbers
 * @param {number} [difference=0]
 * @returns {number|null} test result
 *
 * @example
 * tTestTwoSample([1, 2, 3, 4], [3, 4, 5, 6], 0); // => -2.1908902300206643
 */
function tTestTwoSample(sampleX, sampleY, difference) {
    var n = sampleX.length;
    var m = sampleY.length;

    // If either sample doesn't actually have any values, we can't
    // compute this at all, so we return `null`.
    if (!n || !m) {
        return null;
    }

    // default difference (mu) is zero
    if (!difference) {
        difference = 0;
    }

    var meanX = mean(sampleX);
    var meanY = mean(sampleY);
    var sampleVarianceX = sampleVariance(sampleX);
    var sampleVarianceY = sampleVariance(sampleY);

    if (
        typeof meanX === "number" &&
        typeof meanY === "number" &&
        typeof sampleVarianceX === "number" &&
        typeof sampleVarianceY === "number"
    ) {
        var weightedVariance =
            ((n - 1) * sampleVarianceX + (m - 1) * sampleVarianceY) /
            (n + m - 2);

        return (
            (meanX - meanY - difference) /
            Math.sqrt(weightedVariance * (1 / n + 1 / m))
        );
    }
}

/**
 * This function calculates the Wilcoxon rank sum statistic for the first sample
 * with respect to the second. The Wilcoxon rank sum test is a non-parametric
 * alternative to the t-test which is equivalent to the
 * [Mann-Whitney U test](https://en.wikipedia.org/wiki/Mann%E2%80%93Whitney_U_test).
 * The statistic is calculated by pooling all the observations together, ranking them,
 * and then summing the ranks associated with one of the samples. If this rank sum is
 * sufficiently large or small we reject the hypothesis that the two samples come
 * from the same distribution in favor of the alternative that one is shifted with
 * respect to the other.
 *
 * @param {Array<number>} sampleX a sample as an array of numbers
 * @param {Array<number>} sampleY a sample as an array of numbers
 * @returns {number} rank sum for sampleX
 *
 * @example
 * wilcoxonRankSum([1, 4, 8], [9, 12, 15]); // => 6
 */
function wilcoxonRankSum(sampleX, sampleY) {
    if (!sampleX.length || !sampleY.length) {
        throw new Error("Neither sample can be empty");
    }

    var pooledSamples = sampleX
        .map(function (x) { return ({ label: "x", value: x }); })
        .concat(sampleY.map(function (y) { return ({ label: "y", value: y }); }))
        .sort(function (a, b) { return a.value - b.value; });

    for (var rank = 0; rank < pooledSamples.length; rank++) {
        pooledSamples[rank].rank = rank;
    }

    var tiedRanks = [pooledSamples[0].rank];
    for (var i = 1; i < pooledSamples.length; i++) {
        if (pooledSamples[i].value === pooledSamples[i - 1].value) {
            tiedRanks.push(pooledSamples[i].rank);
            if (i === pooledSamples.length - 1) {
                replaceRanksInPlace(pooledSamples, tiedRanks);
            }
        } else if (tiedRanks.length > 1) {
            replaceRanksInPlace(pooledSamples, tiedRanks);
        } else {
            tiedRanks = [pooledSamples[i].rank];
        }
    }

    function replaceRanksInPlace(pooledSamples, tiedRanks) {
        var average = (tiedRanks[0] + tiedRanks[tiedRanks.length - 1]) / 2;
        for (var i = 0; i < tiedRanks.length; i++) {
            pooledSamples[tiedRanks[i]].rank = average;
        }
    }

    var rankSum = 0;

    for (var i$1 = 0; i$1 < pooledSamples.length; i$1++) {
        var sample = pooledSamples[i$1];
        if (sample.label === "x") {
            rankSum += sample.rank + 1;
        }
    }

    return rankSum;
}

/**
 * [Bayesian Classifier](http://en.wikipedia.org/wiki/Naive_Bayes_classifier)
 *
 * This is a naïve bayesian classifier that takes
 * singly-nested objects.
 *
 * @class
 * @example
 * var bayes = new BayesianClassifier();
 * bayes.train({
 *   species: 'Cat'
 * }, 'animal');
 * var result = bayes.score({
 *   species: 'Cat'
 * })
 * // result
 * // {
 * //   animal: 1
 * // }
 */
var BayesianClassifier = function BayesianClassifier() {
    // The number of items that are currently
    // classified in the model
    this.totalCount = 0;
    // Every item classified in the model
    this.data = {};
};

/**
 * Train the classifier with a new item, which has a single
 * dimension of Javascript literal keys and values.
 *
 * @param {Object} item an object with singly-deep properties
 * @param {string} category the category this item belongs to
 * @return {undefined} adds the item to the classifier
 */
BayesianClassifier.prototype.train = function train (item, category) {
    // If the data object doesn't have any values
    // for this category, create a new object for it.
    if (!this.data[category]) {
        this.data[category] = {};
    }

    // Iterate through each key in the item.
    for (var k in item) {
        var v = item[k];
        // Initialize the nested object `data[category][k][item[k]]`
        // with an object of keys that equal 0.
        if (this.data[category][k] === undefined) {
            this.data[category][k] = {};
        }
        if (this.data[category][k][v] === undefined) {
            this.data[category][k][v] = 0;
        }

        // And increment the key for this key/value combination.
        this.data[category][k][v]++;
    }

    // Increment the number of items classified
    this.totalCount++;
};

/**
 * Generate a score of how well this item matches all
 * possible categories based on its attributes
 *
 * @param {Object} item an item in the same format as with train
 * @returns {Object} of probabilities that this item belongs to a
 * given category.
 */
BayesianClassifier.prototype.score = function score (item) {
    // Initialize an empty array of odds per category.
    var odds = {};
    var category;
    // Iterate through each key in the item,
    // then iterate through each category that has been used
    // in previous calls to `.train()`
    for (var k in item) {
        var v = item[k];
        for (category in this.data) {
            // Create an empty object for storing key - value combinations
            // for this category.
            odds[category] = {};

            // If this item doesn't even have a property, it counts for nothing,
            // but if it does have the property that we're looking for from
            // the item to categorize, it counts based on how popular it is
            // versus the whole population.
            if (this.data[category][k]) {
                odds[category][k + "_" + v] =
                    (this.data[category][k][v] || 0) / this.totalCount;
            } else {
                odds[category][k + "_" + v] = 0;
            }
        }
    }

    // Set up a new object that will contain sums of these odds by category
    var oddsSums = {};

    for (category in odds) {
        // Tally all of the odds for each category-combination pair -
        // the non-existence of a category does not add anything to the
        // score.
        oddsSums[category] = 0;
        for (var combination in odds[category]) {
            oddsSums[category] += odds[category][combination];
        }
    }

    return oddsSums;
};

/**
 * This is a single-layer [Perceptron Classifier](http://en.wikipedia.org/wiki/Perceptron) that takes
 * arrays of numbers and predicts whether they should be classified
 * as either 0 or 1 (negative or positive examples).
 * @class
 * @example
 * // Create the model
 * var p = new PerceptronModel();
 * // Train the model with input with a diagonal boundary.
 * for (var i = 0; i < 5; i++) {
 *     p.train([1, 1], 1);
 *     p.train([0, 1], 0);
 *     p.train([1, 0], 0);
 *     p.train([0, 0], 0);
 * }
 * p.predict([0, 0]); // 0
 * p.predict([0, 1]); // 0
 * p.predict([1, 0]); // 0
 * p.predict([1, 1]); // 1
 */
var PerceptronModel = function PerceptronModel() {
    // The weights, or coefficients of the model;
    // weights are only populated when training with data.
    this.weights = [];
    // The bias term, or intercept; it is also a weight but
    // it's stored separately for convenience as it is always
    // multiplied by one.
    this.bias = 0;
};
/**
 * **Predict**: Use an array of features with the weight array and bias
 * to predict whether an example is labeled 0 or 1.
 *
 * @param {Array<number>} features an array of features as numbers
 * @returns {number} 1 if the score is over 0, otherwise 0
 */
PerceptronModel.prototype.predict = function predict (features) {
    // Only predict if previously trained
    // on the same size feature array(s).
    if (features.length !== this.weights.length) {
        return null;
    }

    // Calculate the sum of features times weights,
    // with the bias added (implicitly times one).
    var score = 0;
    for (var i = 0; i < this.weights.length; i++) {
        score += this.weights[i] * features[i];
    }
    score += this.bias;

    // Classify as 1 if the score is over 0, otherwise 0.
    if (score > 0) {
        return 1;
    } else {
        return 0;
    }
};

/**
 * **Train** the classifier with a new example, which is
 * a numeric array of features and a 0 or 1 label.
 *
 * @param {Array<number>} features an array of features as numbers
 * @param {number} label either 0 or 1
 * @returns {PerceptronModel} this
 */
PerceptronModel.prototype.train = function train (features, label) {
    // Require that only labels of 0 or 1 are considered.
    if (label !== 0 && label !== 1) {
        return null;
    }
    // The length of the feature array determines
    // the length of the weight array.
    // The perceptron will continue learning as long as
    // it keeps seeing feature arrays of the same length.
    // When it sees a new data shape, it initializes.
    if (features.length !== this.weights.length) {
        this.weights = features;
        this.bias = 1;
    }
    // Make a prediction based on current weights.
    var prediction = this.predict(features);
    // Update the weights if the prediction is wrong.
    if (typeof prediction === "number" && prediction !== label) {
        var gradient = label - prediction;
        for (var i = 0; i < this.weights.length; i++) {
            this.weights[i] += gradient * features[i];
        }
        this.bias += gradient;
    }
    return this;
};

/**
 * We use `ε`, epsilon, as a stopping criterion when we want to iterate
 * until we're "close enough". Epsilon is a very small number: for
 * simple statistics, that number is **0.0001**
 *
 * This is used in calculations like the binomialDistribution, in which
 * the process of finding a value is [iterative](https://en.wikipedia.org/wiki/Iterative_method):
 * it progresses until it is close enough.
 *
 * Below is an example of using epsilon in [gradient descent](https://en.wikipedia.org/wiki/Gradient_descent),
 * where we're trying to find a local minimum of a function's derivative,
 * given by the `fDerivative` method.
 *
 * @example
 * // From calculation, we expect that the local minimum occurs at x=9/4
 * var x_old = 0;
 * // The algorithm starts at x=6
 * var x_new = 6;
 * var stepSize = 0.01;
 *
 * function fDerivative(x) {
 *   return 4 * Math.pow(x, 3) - 9 * Math.pow(x, 2);
 * }
 *
 * // The loop runs until the difference between the previous
 * // value and the current value is smaller than epsilon - a rough
 * // meaure of 'close enough'
 * while (Math.abs(x_new - x_old) > ss.epsilon) {
 *   x_old = x_new;
 *   x_new = x_old - stepSize * fDerivative(x_old);
 * }
 *
 * console.log('Local minimum occurs at', x_new);
 */
var epsilon = 0.0001;

/**
 * A [Factorial](https://en.wikipedia.org/wiki/Factorial), usually written n!, is the product of all positive
 * integers less than or equal to n. Often factorial is implemented
 * recursively, but this iterative approach is significantly faster
 * and simpler.
 *
 * @param {number} n input, must be an integer number 1 or greater
 * @returns {number} factorial: n!
 * @throws {Error} if n is less than 0 or not an integer
 * @example
 * factorial(5); // => 120
 */
function factorial(n) {
    // factorial is mathematically undefined for negative numbers
    if (n < 0) {
        throw new Error("factorial requires a non-negative value");
    }

    if (Math.floor(n) !== n) {
        throw new Error("factorial requires an integer input");
    }

    // typically you'll expand the factorial function going down, like
    // 5! = 5 * 4 * 3 * 2 * 1. This is going in the opposite direction,
    // counting from 2 up to the number in question, and since anything
    // multiplied by 1 is itself, the loop only needs to start at 2.
    var accumulator = 1;
    for (var i = 2; i <= n; i++) {
        // for each number up to and including the number `n`, multiply
        // the accumulator my that number.
        accumulator *= i;
    }
    return accumulator;
}

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
        var seriesCoefficient =
            Math.pow(n / Math.E, n) * Math.sqrt(2 * Math.PI * (n + 1 / 6));

        var seriesDenom = n + 1 / 4;

        var seriesExpansion =
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

// Define series coefficients
var COEFFICIENTS = [
    0.99999999999999709182, 57.156235665862923517, -59.597960355475491248,
    14.136097974741747174, -0.49191381609762019978, 0.33994649984811888699e-4,
    0.46523628927048575665e-4, -0.98374475304879564677e-4,
    0.15808870322491248884e-3, -0.21026444172410488319e-3,
    0.2174396181152126432e-3, -0.16431810653676389022e-3,
    0.84418223983852743293e-4, -0.2619083840158140867e-4,
    0.36899182659531622704e-5
];

var g = 607 / 128;
var LOGSQRT2PI = Math.log(Math.sqrt(2 * Math.PI));

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
    var a = COEFFICIENTS[0];

    for (var i = 1; i < 15; i++) {
        a += COEFFICIENTS[i] / (n + i);
    }

    var tmp = g + 0.5 + n;

    // Return natural logarithm of gamma(n)
    return LOGSQRT2PI + Math.log(a) - tmp + (n + 0.5) * Math.log(tmp);
}

/**
 * The [Bernoulli distribution](http://en.wikipedia.org/wiki/Bernoulli_distribution)
 * is the probability discrete
 * distribution of a random variable which takes value 1 with success
 * probability `p` and value 0 with failure
 * probability `q` = 1 - `p`. It can be used, for example, to represent the
 * toss of a coin, where "1" is defined to mean "heads" and "0" is defined
 * to mean "tails" (or vice versa). It is
 * a special case of a Binomial Distribution
 * where `n` = 1.
 *
 * @param {number} p input value, between 0 and 1 inclusive
 * @returns {number[]} values of bernoulli distribution at this point
 * @throws {Error} if p is outside 0 and 1
 * @example
 * bernoulliDistribution(0.3); // => [0.7, 0.3]
 */
function bernoulliDistribution(p) /*: number[] */ {
    // Check that `p` is a valid probability (0 ≤ p ≤ 1)
    if (p < 0 || p > 1) {
        throw new Error(
            "bernoulliDistribution requires probability to be between 0 and 1 inclusive"
        );
    }

    return [1 - p, p];
}

/**
 * The [Binomial Distribution](http://en.wikipedia.org/wiki/Binomial_distribution) is the discrete probability
 * distribution of the number of successes in a sequence of n independent yes/no experiments, each of which yields
 * success with probability `probability`. Such a success/failure experiment is also called a Bernoulli experiment or
 * Bernoulli trial; when trials = 1, the Binomial Distribution is a Bernoulli Distribution.
 *
 * @param {number} trials number of trials to simulate
 * @param {number} probability
 * @returns {number[]} output
 */
function binomialDistribution(trials, probability) /*: ?number[] */ {
    // Check that `p` is a valid probability (0 ≤ p ≤ 1),
    // that `n` is an integer, strictly positive.
    if (probability < 0 || probability > 1 || trials <= 0 || trials % 1 !== 0) {
        return undefined;
    }

    // We initialize `x`, the random variable, and `accumulator`, an accumulator
    // for the cumulative distribution function to 0. `distribution_functions`
    // is the object we'll return with the `probability_of_x` and the
    // `cumulativeProbability_of_x`, as well as the calculated mean &
    // variance. We iterate until the `cumulativeProbability_of_x` is
    // within `epsilon` of 1.0.
    var x = 0;
    var cumulativeProbability = 0;
    var cells = [];
    var binomialCoefficient = 1;

    // This algorithm iterates through each potential outcome,
    // until the `cumulativeProbability` is very close to 1, at
    // which point we've defined the vast majority of outcomes
    do {
        // a [probability mass function](https://en.wikipedia.org/wiki/Probability_mass_function)
        cells[x] =
            binomialCoefficient *
            Math.pow(probability, x) *
            Math.pow(1 - probability, trials - x);
        cumulativeProbability += cells[x];
        x++;
        binomialCoefficient = (binomialCoefficient * (trials - x + 1)) / x;
        // when the cumulativeProbability is nearly 1, we've calculated
        // the useful range of this distribution
    } while (cumulativeProbability < 1 - epsilon);

    return cells;
}

/**
 * The [Poisson Distribution](http://en.wikipedia.org/wiki/Poisson_distribution)
 * is a discrete probability distribution that expresses the probability
 * of a given number of events occurring in a fixed interval of time
 * and/or space if these events occur with a known average rate and
 * independently of the time since the last event.
 *
 * The Poisson Distribution is characterized by the strictly positive
 * mean arrival or occurrence rate, `λ`.
 *
 * @param {number} lambda location poisson distribution
 * @returns {number[]} values of poisson distribution at that point
 */
function poissonDistribution(lambda) /*: ?number[] */ {
    // Check that lambda is strictly positive
    if (lambda <= 0) {
        return undefined;
    }

    // our current place in the distribution
    var x = 0;
    // and we keep track of the current cumulative probability, in
    // order to know when to stop calculating chances.
    var cumulativeProbability = 0;
    // the calculated cells to be returned
    var cells = [];
    var factorialX = 1;

    // This algorithm iterates through each potential outcome,
    // until the `cumulativeProbability` is very close to 1, at
    // which point we've defined the vast majority of outcomes
    do {
        // a [probability mass function](https://en.wikipedia.org/wiki/Probability_mass_function)
        cells[x] = (Math.exp(-lambda) * Math.pow(lambda, x)) / factorialX;
        cumulativeProbability += cells[x];
        x++;
        factorialX *= x;
        // when the cumulativeProbability is nearly 1, we've calculated
        // the useful range of this distribution
    } while (cumulativeProbability < 1 - epsilon);

    return cells;
}

/**
 * **Percentage Points of the χ2 (Chi-Squared) Distribution**
 *
 * The [χ2 (Chi-Squared) Distribution](http://en.wikipedia.org/wiki/Chi-squared_distribution) is used in the common
 * chi-squared tests for goodness of fit of an observed distribution to a theoretical one, the independence of two
 * criteria of classification of qualitative data, and in confidence interval estimation for a population standard
 * deviation of a normal distribution from a sample standard deviation.
 *
 * Values from Appendix 1, Table III of William W. Hines & Douglas C. Montgomery, "Probability and Statistics in
 * Engineering and Management Science", Wiley (1980).
 */
var chiSquaredDistributionTable = {
    1: {
        0.995: 0,
        0.99: 0,
        0.975: 0,
        0.95: 0,
        0.9: 0.02,
        0.5: 0.45,
        0.1: 2.71,
        0.05: 3.84,
        0.025: 5.02,
        0.01: 6.63,
        0.005: 7.88
    },
    2: {
        0.995: 0.01,
        0.99: 0.02,
        0.975: 0.05,
        0.95: 0.1,
        0.9: 0.21,
        0.5: 1.39,
        0.1: 4.61,
        0.05: 5.99,
        0.025: 7.38,
        0.01: 9.21,
        0.005: 10.6
    },
    3: {
        0.995: 0.07,
        0.99: 0.11,
        0.975: 0.22,
        0.95: 0.35,
        0.9: 0.58,
        0.5: 2.37,
        0.1: 6.25,
        0.05: 7.81,
        0.025: 9.35,
        0.01: 11.34,
        0.005: 12.84
    },
    4: {
        0.995: 0.21,
        0.99: 0.3,
        0.975: 0.48,
        0.95: 0.71,
        0.9: 1.06,
        0.5: 3.36,
        0.1: 7.78,
        0.05: 9.49,
        0.025: 11.14,
        0.01: 13.28,
        0.005: 14.86
    },
    5: {
        0.995: 0.41,
        0.99: 0.55,
        0.975: 0.83,
        0.95: 1.15,
        0.9: 1.61,
        0.5: 4.35,
        0.1: 9.24,
        0.05: 11.07,
        0.025: 12.83,
        0.01: 15.09,
        0.005: 16.75
    },
    6: {
        0.995: 0.68,
        0.99: 0.87,
        0.975: 1.24,
        0.95: 1.64,
        0.9: 2.2,
        0.5: 5.35,
        0.1: 10.65,
        0.05: 12.59,
        0.025: 14.45,
        0.01: 16.81,
        0.005: 18.55
    },
    7: {
        0.995: 0.99,
        0.99: 1.25,
        0.975: 1.69,
        0.95: 2.17,
        0.9: 2.83,
        0.5: 6.35,
        0.1: 12.02,
        0.05: 14.07,
        0.025: 16.01,
        0.01: 18.48,
        0.005: 20.28
    },
    8: {
        0.995: 1.34,
        0.99: 1.65,
        0.975: 2.18,
        0.95: 2.73,
        0.9: 3.49,
        0.5: 7.34,
        0.1: 13.36,
        0.05: 15.51,
        0.025: 17.53,
        0.01: 20.09,
        0.005: 21.96
    },
    9: {
        0.995: 1.73,
        0.99: 2.09,
        0.975: 2.7,
        0.95: 3.33,
        0.9: 4.17,
        0.5: 8.34,
        0.1: 14.68,
        0.05: 16.92,
        0.025: 19.02,
        0.01: 21.67,
        0.005: 23.59
    },
    10: {
        0.995: 2.16,
        0.99: 2.56,
        0.975: 3.25,
        0.95: 3.94,
        0.9: 4.87,
        0.5: 9.34,
        0.1: 15.99,
        0.05: 18.31,
        0.025: 20.48,
        0.01: 23.21,
        0.005: 25.19
    },
    11: {
        0.995: 2.6,
        0.99: 3.05,
        0.975: 3.82,
        0.95: 4.57,
        0.9: 5.58,
        0.5: 10.34,
        0.1: 17.28,
        0.05: 19.68,
        0.025: 21.92,
        0.01: 24.72,
        0.005: 26.76
    },
    12: {
        0.995: 3.07,
        0.99: 3.57,
        0.975: 4.4,
        0.95: 5.23,
        0.9: 6.3,
        0.5: 11.34,
        0.1: 18.55,
        0.05: 21.03,
        0.025: 23.34,
        0.01: 26.22,
        0.005: 28.3
    },
    13: {
        0.995: 3.57,
        0.99: 4.11,
        0.975: 5.01,
        0.95: 5.89,
        0.9: 7.04,
        0.5: 12.34,
        0.1: 19.81,
        0.05: 22.36,
        0.025: 24.74,
        0.01: 27.69,
        0.005: 29.82
    },
    14: {
        0.995: 4.07,
        0.99: 4.66,
        0.975: 5.63,
        0.95: 6.57,
        0.9: 7.79,
        0.5: 13.34,
        0.1: 21.06,
        0.05: 23.68,
        0.025: 26.12,
        0.01: 29.14,
        0.005: 31.32
    },
    15: {
        0.995: 4.6,
        0.99: 5.23,
        0.975: 6.27,
        0.95: 7.26,
        0.9: 8.55,
        0.5: 14.34,
        0.1: 22.31,
        0.05: 25,
        0.025: 27.49,
        0.01: 30.58,
        0.005: 32.8
    },
    16: {
        0.995: 5.14,
        0.99: 5.81,
        0.975: 6.91,
        0.95: 7.96,
        0.9: 9.31,
        0.5: 15.34,
        0.1: 23.54,
        0.05: 26.3,
        0.025: 28.85,
        0.01: 32,
        0.005: 34.27
    },
    17: {
        0.995: 5.7,
        0.99: 6.41,
        0.975: 7.56,
        0.95: 8.67,
        0.9: 10.09,
        0.5: 16.34,
        0.1: 24.77,
        0.05: 27.59,
        0.025: 30.19,
        0.01: 33.41,
        0.005: 35.72
    },
    18: {
        0.995: 6.26,
        0.99: 7.01,
        0.975: 8.23,
        0.95: 9.39,
        0.9: 10.87,
        0.5: 17.34,
        0.1: 25.99,
        0.05: 28.87,
        0.025: 31.53,
        0.01: 34.81,
        0.005: 37.16
    },
    19: {
        0.995: 6.84,
        0.99: 7.63,
        0.975: 8.91,
        0.95: 10.12,
        0.9: 11.65,
        0.5: 18.34,
        0.1: 27.2,
        0.05: 30.14,
        0.025: 32.85,
        0.01: 36.19,
        0.005: 38.58
    },
    20: {
        0.995: 7.43,
        0.99: 8.26,
        0.975: 9.59,
        0.95: 10.85,
        0.9: 12.44,
        0.5: 19.34,
        0.1: 28.41,
        0.05: 31.41,
        0.025: 34.17,
        0.01: 37.57,
        0.005: 40
    },
    21: {
        0.995: 8.03,
        0.99: 8.9,
        0.975: 10.28,
        0.95: 11.59,
        0.9: 13.24,
        0.5: 20.34,
        0.1: 29.62,
        0.05: 32.67,
        0.025: 35.48,
        0.01: 38.93,
        0.005: 41.4
    },
    22: {
        0.995: 8.64,
        0.99: 9.54,
        0.975: 10.98,
        0.95: 12.34,
        0.9: 14.04,
        0.5: 21.34,
        0.1: 30.81,
        0.05: 33.92,
        0.025: 36.78,
        0.01: 40.29,
        0.005: 42.8
    },
    23: {
        0.995: 9.26,
        0.99: 10.2,
        0.975: 11.69,
        0.95: 13.09,
        0.9: 14.85,
        0.5: 22.34,
        0.1: 32.01,
        0.05: 35.17,
        0.025: 38.08,
        0.01: 41.64,
        0.005: 44.18
    },
    24: {
        0.995: 9.89,
        0.99: 10.86,
        0.975: 12.4,
        0.95: 13.85,
        0.9: 15.66,
        0.5: 23.34,
        0.1: 33.2,
        0.05: 36.42,
        0.025: 39.36,
        0.01: 42.98,
        0.005: 45.56
    },
    25: {
        0.995: 10.52,
        0.99: 11.52,
        0.975: 13.12,
        0.95: 14.61,
        0.9: 16.47,
        0.5: 24.34,
        0.1: 34.28,
        0.05: 37.65,
        0.025: 40.65,
        0.01: 44.31,
        0.005: 46.93
    },
    26: {
        0.995: 11.16,
        0.99: 12.2,
        0.975: 13.84,
        0.95: 15.38,
        0.9: 17.29,
        0.5: 25.34,
        0.1: 35.56,
        0.05: 38.89,
        0.025: 41.92,
        0.01: 45.64,
        0.005: 48.29
    },
    27: {
        0.995: 11.81,
        0.99: 12.88,
        0.975: 14.57,
        0.95: 16.15,
        0.9: 18.11,
        0.5: 26.34,
        0.1: 36.74,
        0.05: 40.11,
        0.025: 43.19,
        0.01: 46.96,
        0.005: 49.65
    },
    28: {
        0.995: 12.46,
        0.99: 13.57,
        0.975: 15.31,
        0.95: 16.93,
        0.9: 18.94,
        0.5: 27.34,
        0.1: 37.92,
        0.05: 41.34,
        0.025: 44.46,
        0.01: 48.28,
        0.005: 50.99
    },
    29: {
        0.995: 13.12,
        0.99: 14.26,
        0.975: 16.05,
        0.95: 17.71,
        0.9: 19.77,
        0.5: 28.34,
        0.1: 39.09,
        0.05: 42.56,
        0.025: 45.72,
        0.01: 49.59,
        0.005: 52.34
    },
    30: {
        0.995: 13.79,
        0.99: 14.95,
        0.975: 16.79,
        0.95: 18.49,
        0.9: 20.6,
        0.5: 29.34,
        0.1: 40.26,
        0.05: 43.77,
        0.025: 46.98,
        0.01: 50.89,
        0.005: 53.67
    },
    40: {
        0.995: 20.71,
        0.99: 22.16,
        0.975: 24.43,
        0.95: 26.51,
        0.9: 29.05,
        0.5: 39.34,
        0.1: 51.81,
        0.05: 55.76,
        0.025: 59.34,
        0.01: 63.69,
        0.005: 66.77
    },
    50: {
        0.995: 27.99,
        0.99: 29.71,
        0.975: 32.36,
        0.95: 34.76,
        0.9: 37.69,
        0.5: 49.33,
        0.1: 63.17,
        0.05: 67.5,
        0.025: 71.42,
        0.01: 76.15,
        0.005: 79.49
    },
    60: {
        0.995: 35.53,
        0.99: 37.48,
        0.975: 40.48,
        0.95: 43.19,
        0.9: 46.46,
        0.5: 59.33,
        0.1: 74.4,
        0.05: 79.08,
        0.025: 83.3,
        0.01: 88.38,
        0.005: 91.95
    },
    70: {
        0.995: 43.28,
        0.99: 45.44,
        0.975: 48.76,
        0.95: 51.74,
        0.9: 55.33,
        0.5: 69.33,
        0.1: 85.53,
        0.05: 90.53,
        0.025: 95.02,
        0.01: 100.42,
        0.005: 104.22
    },
    80: {
        0.995: 51.17,
        0.99: 53.54,
        0.975: 57.15,
        0.95: 60.39,
        0.9: 64.28,
        0.5: 79.33,
        0.1: 96.58,
        0.05: 101.88,
        0.025: 106.63,
        0.01: 112.33,
        0.005: 116.32
    },
    90: {
        0.995: 59.2,
        0.99: 61.75,
        0.975: 65.65,
        0.95: 69.13,
        0.9: 73.29,
        0.5: 89.33,
        0.1: 107.57,
        0.05: 113.14,
        0.025: 118.14,
        0.01: 124.12,
        0.005: 128.3
    },
    100: {
        0.995: 67.33,
        0.99: 70.06,
        0.975: 74.22,
        0.95: 77.93,
        0.9: 82.36,
        0.5: 99.33,
        0.1: 118.5,
        0.05: 124.34,
        0.025: 129.56,
        0.01: 135.81,
        0.005: 140.17
    }
};

/**
 * The [χ2 (Chi-Squared) Goodness-of-Fit Test](http://en.wikipedia.org/wiki/Goodness_of_fit#Pearson.27s_chi-squared_test)
 * uses a measure of goodness of fit which is the sum of differences between observed and expected outcome frequencies
 * (that is, counts of observations), each squared and divided by the number of observations expected given the
 * hypothesized distribution. The resulting χ2 statistic, `chiSquared`, can be compared to the chi-squared distribution
 * to determine the goodness of fit. In order to determine the degrees of freedom of the chi-squared distribution, one
 * takes the total number of observed frequencies and subtracts the number of estimated parameters. The test statistic
 * follows, approximately, a chi-square distribution with (k − c) degrees of freedom where `k` is the number of non-empty
 * cells and `c` is the number of estimated parameters for the distribution.
 *
 * @param {Array<number>} data
 * @param {Function} distributionType a function that returns a point in a distribution:
 * for instance, binomial, bernoulli, or poisson
 * @param {number} significance
 * @returns {number} chi squared goodness of fit
 * @example
 * // Data from Poisson goodness-of-fit example 10-19 in William W. Hines & Douglas C. Montgomery,
 * // "Probability and Statistics in Engineering and Management Science", Wiley (1980).
 * var data1019 = [
 *     0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 *     0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
 *     1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
 *     2, 2, 2, 2, 2, 2, 2, 2, 2,
 *     3, 3, 3, 3
 * ];
 * ss.chiSquaredGoodnessOfFit(data1019, ss.poissonDistribution, 0.05); //= false
 */
function chiSquaredGoodnessOfFit(data, distributionType, significance) {
    // Estimate from the sample data, a weighted mean.
    var inputMean = mean(data);
    // Calculated value of the χ2 statistic.
    var chiSquared = 0;
    // Number of hypothesized distribution parameters estimated, expected to be supplied in the distribution test.
    // Lose one degree of freedom for estimating `lambda` from the sample data.
    var c = 1;
    // The hypothesized distribution.
    // Generate the hypothesized distribution.
    var hypothesizedDistribution = distributionType(inputMean);
    var observedFrequencies = [];
    var expectedFrequencies = [];

    // Create an array holding a histogram from the sample data, of
    // the form `{ value: numberOfOcurrences }`
    for (var i = 0; i < data.length; i++) {
        if (observedFrequencies[data[i]] === undefined) {
            observedFrequencies[data[i]] = 0;
        }
        observedFrequencies[data[i]]++;
    }

    // The histogram we created might be sparse - there might be gaps
    // between values. So we iterate through the histogram, making
    // sure that instead of undefined, gaps have 0 values.
    for (var i$1 = 0; i$1 < observedFrequencies.length; i$1++) {
        if (observedFrequencies[i$1] === undefined) {
            observedFrequencies[i$1] = 0;
        }
    }

    // Create an array holding a histogram of expected data given the
    // sample size and hypothesized distribution.
    for (var k in hypothesizedDistribution) {
        if (k in observedFrequencies) {
            expectedFrequencies[+k] = hypothesizedDistribution[k] * data.length;
        }
    }

    // Working backward through the expected frequencies, collapse classes
    // if less than three observations are expected for a class.
    // This transformation is applied to the observed frequencies as well.
    for (var k$1 = expectedFrequencies.length - 1; k$1 >= 0; k$1--) {
        if (expectedFrequencies[k$1] < 3) {
            expectedFrequencies[k$1 - 1] += expectedFrequencies[k$1];
            expectedFrequencies.pop();

            observedFrequencies[k$1 - 1] += observedFrequencies[k$1];
            observedFrequencies.pop();
        }
    }

    // Iterate through the squared differences between observed & expected
    // frequencies, accumulating the `chiSquared` statistic.
    for (var k$2 = 0; k$2 < observedFrequencies.length; k$2++) {
        chiSquared +=
            Math.pow(observedFrequencies[k$2] - expectedFrequencies[k$2], 2) /
            expectedFrequencies[k$2];
    }

    // Calculate degrees of freedom for this test and look it up in the
    // `chiSquaredDistributionTable` in order to
    // accept or reject the goodness-of-fit of the hypothesized distribution.
    // Degrees of freedom, calculated as (number of class intervals -
    // number of hypothesized distribution parameters estimated - 1)
    var degreesOfFreedom = observedFrequencies.length - c - 1;
    return (
        chiSquaredDistributionTable[degreesOfFreedom][significance] < chiSquared
    );
}

var SQRT_2PI$1 = Math.sqrt(2 * Math.PI);

/**
 * [Well-known kernels](https://en.wikipedia.org/wiki/Kernel_(statistics)#Kernel_functions_in_common_use)
 * @private
 */
var kernels = {
    /**
     * The gaussian kernel.
     * @private
     */
    gaussian: function (u) {
        return Math.exp(-0.5 * u * u) / SQRT_2PI$1;
    }
};

/**
 * Well known bandwidth selection methods
 * @private
 */
var bandwidthMethods = {
    /**
     * The ["normal reference distribution"
     * rule-of-thumb](https://stat.ethz.ch/R-manual/R-devel/library/MASS/html/bandwidth.nrd.html),
     * a commonly used version of [Silverman's
     * rule-of-thumb](https://en.wikipedia.org/wiki/Kernel_density_estimation#A_rule-of-thumb_bandwidth_estimator).
     * @private
     */
    nrd: function (x) {
        var s = sampleStandardDeviation(x);
        var iqr = interquartileRange(x);
        if (typeof iqr === "number") {
            s = Math.min(s, iqr / 1.34);
        }
        return 1.06 * s * Math.pow(x.length, -0.2);
    }
};

/**
 * [Kernel density estimation](https://en.wikipedia.org/wiki/Kernel_density_estimation)
 * is a useful tool for, among other things, estimating the shape of the
 * underlying probability distribution from a sample.
 *
 * @name kernelDensityEstimation
 * @param X sample values
 * @param kernel The kernel function to use. If a function is provided, it should return non-negative values and integrate to 1. Defaults to 'gaussian'.
 * @param bandwidthMethod The "bandwidth selection" method to use, or a fixed bandwidth value. Defaults to "nrd", the commonly-used ["normal reference distribution" rule-of-thumb](https://stat.ethz.ch/R-manual/R-devel/library/MASS/html/bandwidth.nrd.html).
 * @returns {Function} An estimated [probability density function](https://en.wikipedia.org/wiki/Probability_density_function) for the given sample. The returned function runs in `O(X.length)`.
 */
function kernelDensityEstimation(X, kernel, bandwidthMethod) {
    var kernelFn;
    if (kernel === undefined) {
        kernelFn = kernels.gaussian;
    } else if (typeof kernel === "string") {
        if (!kernels[kernel]) {
            throw new Error('Unknown kernel "' + kernel + '"');
        }
        kernelFn = kernels[kernel];
    } else {
        kernelFn = kernel;
    }

    var bandwidth;
    if (typeof bandwidthMethod === "undefined") {
        bandwidth = bandwidthMethods.nrd(X);
    } else if (typeof bandwidthMethod === "string") {
        if (!bandwidthMethods[bandwidthMethod]) {
            throw new Error(
                'Unknown bandwidth method "' + bandwidthMethod + '"'
            );
        }
        bandwidth = bandwidthMethods[bandwidthMethod](X);
    } else {
        bandwidth = bandwidthMethod;
    }

    return function (x) {
        var i = 0;
        var sum = 0;
        for (i = 0; i < X.length; i++) {
            sum += kernelFn((x - X[i]) / bandwidth);
        }
        return sum / bandwidth / X.length;
    };
}

/**
 * The [Z-Score, or Standard Score](http://en.wikipedia.org/wiki/Standard_score).
 *
 * The standard score is the number of standard deviations an observation
 * or datum is above or below the mean. Thus, a positive standard score
 * represents a datum above the mean, while a negative standard score
 * represents a datum below the mean. It is a dimensionless quantity
 * obtained by subtracting the population mean from an individual raw
 * score and then dividing the difference by the population standard
 * deviation.
 *
 * The z-score is only defined if one knows the population parameters;
 * if one only has a sample set, then the analogous computation with
 * sample mean and sample standard deviation yields the
 * Student's t-statistic.
 *
 * @param {number} x
 * @param {number} mean
 * @param {number} standardDeviation
 * @return {number} z score
 * @example
 * zScore(78, 80, 5); // => -0.4
 */
function zScore(x, mean, standardDeviation) {
    return (x - mean) / standardDeviation;
}

var SQRT_2PI = Math.sqrt(2 * Math.PI);

function cumulativeDistribution(z) {
    var sum = z;
    var tmp = z;

    // 15 iterations are enough for 4-digit precision
    for (var i = 1; i < 15; i++) {
        tmp *= (z * z) / (2 * i + 1);
        sum += tmp;
    }
    return (
        Math.round((0.5 + (sum / SQRT_2PI) * Math.exp((-z * z) / 2)) * 1e4) /
        1e4
    );
}

/**
 * A standard normal table, also called the unit normal table or Z table,
 * is a mathematical table for the values of Φ (phi), which are the values of
 * the [cumulative distribution function](https://en.wikipedia.org/wiki/Normal_distribution#Cumulative_distribution_function)
 * of the normal distribution. It is used to find the probability that a
 * statistic is observed below, above, or between values on the standard
 * normal distribution, and by extension, any normal distribution.
 */
var standardNormalTable = [];

for (var z = 0; z <= 3.09; z += 0.01) {
    standardNormalTable.push(cumulativeDistribution(z));
}

/**
 * **[Cumulative Standard Normal Probability](http://en.wikipedia.org/wiki/Standard_normal_table)**
 *
 * Since probability tables cannot be
 * printed for every normal distribution, as there are an infinite variety
 * of normal distributions, it is common practice to convert a normal to a
 * standard normal and then use the standard normal table to find probabilities.
 *
 * You can use `.5 + .5 * errorFunction(x / Math.sqrt(2))` to calculate the probability
 * instead of looking it up in a table.
 *
 * @param {number} z
 * @returns {number} cumulative standard normal probability
 */
function cumulativeStdNormalProbability(z) {
    // Calculate the position of this value.
    var absZ = Math.abs(z);
    // Each row begins with a different
    // significant digit: 0.5, 0.6, 0.7, and so on. Each value in the table
    // corresponds to a range of 0.01 in the input values, so the value is
    // multiplied by 100.
    var index = Math.min(
        Math.round(absZ * 100),
        standardNormalTable.length - 1
    );

    // The index we calculate must be in the table as a positive value,
    // but we still pay attention to whether the input is positive
    // or negative, and flip the output value as a last step.
    if (z >= 0) {
        return standardNormalTable[index];
    } else {
        // due to floating-point arithmetic, values in the table with
        // 4 significant figures can nevertheless end up as repeating
        // fractions when they're computed here.
        return Math.round((1 - standardNormalTable[index]) * 1e4) / 1e4;
    }
}

/**
 * **[Logistic Cumulative Distribution Function](https://en.wikipedia.org/wiki/Logistic_distribution)**
 *
 * @param {number} x
 * @returns {number} cumulative standard logistic probability
 */
function cumulativeStdLogisticProbability(x) {
    return 1 / (Math.exp(-x) + 1);
}

/**
 * **[Gaussian error function](http://en.wikipedia.org/wiki/Error_function)**
 *
 * The `errorFunction(x/(sd * Math.sqrt(2)))` is the probability that a value in a
 * normal distribution with standard deviation sd is within x of the mean.
 *
 * This function returns a numerical approximation to the exact value.
 * It uses Horner's method to evaluate the polynomial of τ (tau).
 *
 * @param {number} x input
 * @return {number} error estimation
 * @example
 * errorFunction(1).toFixed(2); // => '0.84'
 */
function errorFunction(x) {
    var t = 1 / (1 + 0.5 * Math.abs(x));
    var tau =
        t *
        Math.exp(
            -x * x +
                ((((((((0.17087277 * t - 0.82215223) * t + 1.48851587) * t -
                    1.13520398) *
                    t +
                    0.27886807) *
                    t -
                    0.18628806) *
                    t +
                    0.09678418) *
                    t +
                    0.37409196) *
                    t +
                    1.00002368) *
                    t -
                1.26551223
        );
    if (x >= 0) {
        return 1 - tau;
    } else {
        return tau - 1;
    }
}

/**
 * The Inverse [Gaussian error function](http://en.wikipedia.org/wiki/Error_function)
 * returns a numerical approximation to the value that would have caused
 * `errorFunction()` to return x.
 *
 * @param {number} x value of error function
 * @returns {number} estimated inverted value
 */
function inverseErrorFunction(x) {
    var a = (8 * (Math.PI - 3)) / (3 * Math.PI * (4 - Math.PI));

    var inv = Math.sqrt(
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

/**
 * The [Probit](http://en.wikipedia.org/wiki/Probit)
 * is the inverse of cumulativeStdNormalProbability(),
 * and is also known as the normal quantile function.
 *
 * It returns the number of standard deviations from the mean
 * where the p'th quantile of values can be found in a normal distribution.
 * So, for example, probit(0.5 + 0.6827/2) ≈ 1 because 68.27% of values are
 * normally found within 1 standard deviation above or below the mean.
 *
 * @param {number} p
 * @returns {number} probit
 */
function probit(p) {
    if (p === 0) {
        p = epsilon;
    } else if (p >= 1) {
        p = 1 - epsilon;
    }
    return Math.sqrt(2) * inverseErrorFunction(2 * p - 1);
}

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

/**
 * Conducts a [permutation test](https://en.wikipedia.org/wiki/Resampling_(statistics)#Permutation_tests)
 * to determine if two data sets are *significantly* different from each other, using
 * the difference of means between the groups as the test statistic.
 * The function allows for the following hypotheses:
 * - two_tail = Null hypothesis: the two distributions are equal.
 * - greater = Null hypothesis: observations from sampleX tend to be smaller than those from sampleY.
 * - less = Null hypothesis: observations from sampleX tend to be greater than those from sampleY.
 * [Learn more about one-tail vs two-tail tests.](https://en.wikipedia.org/wiki/One-_and_two-tailed_tests)
 *
 * @param {Array<number>} sampleX first dataset (e.g. treatment data)
 * @param {Array<number>} sampleY second dataset (e.g. control data)
 * @param {string} alternative alternative hypothesis, either 'two_sided' (default), 'greater', or 'less'
 * @param {number} k number of values in permutation distribution.
 * @param {Function} [randomSource=Math.random] an optional entropy source
 * @returns {number} p-value The probability of observing the difference between groups (as or more extreme than what we did), assuming the null hypothesis.
 *
 * @example
 * var control = [2, 5, 3, 6, 7, 2, 5];
 * var treatment = [20, 5, 13, 12, 7, 2, 2];
 * permutationTest(control, treatment); // ~0.1324
 */
function permutationTest(sampleX, sampleY, alternative, k, randomSource) {
    // Set default arguments
    if (k === undefined) {
        k = 10000;
    }
    if (alternative === undefined) {
        alternative = "two_side";
    }
    if (
        alternative !== "two_side" &&
        alternative !== "greater" &&
        alternative !== "less"
    ) {
        throw new Error(
            "`alternative` must be either 'two_side', 'greater', or 'less'."
        );
    }

    // get means for each sample
    var meanX = mean(sampleX);
    var meanY = mean(sampleY);

    // calculate initial test statistic. This will be our point of comparison with
    // the generated test statistics.
    var testStatistic = meanX - meanY;

    // create test-statistic distribution
    var testStatDsn = new Array(k);

    // combine datsets so we can easily shuffle later
    var allData = sampleX.concat(sampleY);
    var midIndex = Math.floor(allData.length / 2);

    for (var i = 0; i < k; i++) {
        // 1. shuffle data assignments
        shuffleInPlace(allData, randomSource);
        var permLeft = allData.slice(0, midIndex);
        var permRight = allData.slice(midIndex, allData.length);

        // 2.re-calculate test statistic
        var permTestStatistic = mean(permLeft) - mean(permRight);

        // 3. store test statistic to build test statistic distribution
        testStatDsn[i] = permTestStatistic;
    }

    // Calculate p-value depending on alternative
    // For this test, we calculate the percentage of 'extreme' test statistics (subject to our hypothesis)
    // more info on permutation test p-value calculations: https://onlinecourses.science.psu.edu/stat464/node/35
    var numExtremeTStats = 0;
    if (alternative === "two_side") {
        for (var i$1 = 0; i$1 <= k; i$1++) {
            if (Math.abs(testStatDsn[i$1]) >= Math.abs(testStatistic)) {
                numExtremeTStats += 1;
            }
        }
    } else if (alternative === "greater") {
        for (var i$2 = 0; i$2 <= k; i$2++) {
            if (testStatDsn[i$2] >= testStatistic) {
                numExtremeTStats += 1;
            }
        }
    } else {
        // alternative === 'less'
        for (var i$3 = 0; i$3 <= k; i$3++) {
            /* c8 ignore start */
            if (testStatDsn[i$3] <= testStatistic) {
                numExtremeTStats += 1;
            }
            /* c8 ignore end */
        }
    }

    return numExtremeTStats / k;
}

/**
 * [Sign](https://en.wikipedia.org/wiki/Sign_function) is a function
 * that extracts the sign of a real number
 *
 * @param {number} x input value
 * @returns {number} sign value either 1, 0 or -1
 * @throws {TypeError} if the input argument x is not a number
 * @private
 *
 * @example
 * sign(2); // => 1
 */
function sign(x) {
    if (typeof x === "number") {
        if (x < 0) {
            return -1;
        } else if (x === 0) {
            return 0;
        } else {
            return 1;
        }
    } else {
        throw new TypeError("not a number");
    }
}

/**
 * [Bisection method](https://en.wikipedia.org/wiki/Bisection_method) is a root-finding
 * method that repeatedly bisects an interval to find the root.
 *
 * This function returns a numerical approximation to the exact value.
 *
 * @param {Function} func input function
 * @param {number} start - start of interval
 * @param {number} end - end of interval
 * @param {number} maxIterations - the maximum number of iterations
 * @param {number} errorTolerance - the error tolerance
 * @returns {number} estimated root value
 * @throws {TypeError} Argument func must be a function
 *
 * @example
 * bisect(Math.cos,0,4,100,0.003); // => 1.572265625
 */
function bisect(func, start, end, maxIterations, errorTolerance) {
    if (typeof func !== "function")
        { throw new TypeError("func must be a function"); }

    for (var i = 0; i < maxIterations; i++) {
        var output = (start + end) / 2;

        if (
            func(output) === 0 ||
            Math.abs((end - start) / 2) < errorTolerance
        ) {
            return output;
        }

        if (sign(func(output)) === sign(func(start))) {
            start = output;
        } else {
            end = output;
        }
    }

    throw new Error("maximum number of iterations exceeded");
}

/**
 * Calculate Euclidean distance between two points.
 * @param {Array<number>} left First N-dimensional point.
 * @param {Array<number>} right Second N-dimensional point.
 * @returns {number} Distance.
 */
function euclideanDistance(left, right) {
    var sum = 0;
    for (var i = 0; i < left.length; i++) {
        var diff = left[i] - right[i];
        sum += diff * diff;
    }
    return Math.sqrt(sum);
}

/**
 * @typedef {Object} kMeansReturn
 * @property {Array<number>} labels The labels.
 * @property {Array<Array<number>>} centroids The cluster centroids.
 */

/**
 * Perform k-means clustering.
 *
 * @param {Array<Array<number>>} points N-dimensional coordinates of points to be clustered.
 * @param {number} numCluster How many clusters to create.
 * @param {Function} randomSource An optional entropy source that generates uniform values in [0, 1).
 * @return {kMeansReturn} Labels (same length as data) and centroids (same length as numCluster).
 * @throws {Error} If any centroids wind up friendless (i.e., without associated points).
 *
 * @example
 * kMeansCluster([[0.0, 0.5], [1.0, 0.5]], 2); // => {labels: [0, 1], centroids: [[0.0, 0.5], [1.0 0.5]]}
 */
function kMeansCluster(points, numCluster, randomSource) {
    if ( randomSource === void 0 ) randomSource = Math.random;

    var oldCentroids = null;
    var newCentroids = sample(points, numCluster, randomSource);
    var labels = null;
    var change = Number.MAX_VALUE;
    while (change !== 0) {
        labels = labelPoints(points, newCentroids);
        oldCentroids = newCentroids;
        newCentroids = calculateCentroids(points, labels, numCluster);
        change = calculateChange(newCentroids, oldCentroids);
    }
    return {
        labels: labels,
        centroids: newCentroids
    };
}

/**
 * Label each point according to which centroid it is closest to.
 *
 * @private
 * @param {Array<Array<number>>} points Array of XY coordinates.
 * @param {Array<Array<number>>} centroids Current centroids.
 * @return {Array<number>} Group labels.
 */
function labelPoints(points, centroids) {
    return points.map(function (p) {
        var minDist = Number.MAX_VALUE;
        var label = -1;
        for (var i = 0; i < centroids.length; i++) {
            var dist = euclideanDistance(p, centroids[i]);
            if (dist < minDist) {
                minDist = dist;
                label = i;
            }
        }
        return label;
    });
}

/**
 * Calculate centroids for points given labels.
 *
 * @private
 * @param {Array<Array<number>>} points Array of XY coordinates.
 * @param {Array<number>} labels Which groups points belong to.
 * @param {number} numCluster Number of clusters being created.
 * @return {Array<Array<number>>} Centroid for each group.
 * @throws {Error} If any centroids wind up friendless (i.e., without associated points).
 */
function calculateCentroids(points, labels, numCluster) {
    // Initialize accumulators.
    var dimension = points[0].length;
    var centroids = makeMatrix(numCluster, dimension);
    var counts = Array(numCluster).fill(0);

    // Add points to centroids' accumulators and count points per centroid.
    var numPoints = points.length;
    for (var i = 0; i < numPoints; i++) {
        var point = points[i];
        var label = labels[i];
        var current = centroids[label];
        for (var j = 0; j < dimension; j++) {
            current[j] += point[j];
        }
        counts[label] += 1;
    }

    // Rescale centroids, checking for any that have no points.
    for (var i$1 = 0; i$1 < numCluster; i$1++) {
        if (counts[i$1] === 0) {
            throw new Error(("Centroid " + i$1 + " has no friends"));
        }
        var centroid = centroids[i$1];
        for (var j$1 = 0; j$1 < dimension; j$1++) {
            centroid[j$1] /= counts[i$1];
        }
    }

    return centroids;
}

/**
 * Calculate the difference between old centroids and new centroids.
 *
 * @private
 * @param {Array<Array<number>>} left One list of centroids.
 * @param {Array<Array<number>>} right Another list of centroids.
 * @return {number} Distance between centroids.
 */
function calculateChange(left, right) {
    var total = 0;
    for (var i = 0; i < left.length; i++) {
        total += euclideanDistance(left[i], right[i]);
    }
    return total;
}

/**
 * Calculate the [silhouette values](https://en.wikipedia.org/wiki/Silhouette_(clustering))
 * for clustered data.
 *
 * @param {Array<Array<number>>} points N-dimensional coordinates of points.
 * @param {Array<number>} labels Labels of points. This must be the same length as `points`,
 * and values must lie in [0..G-1], where G is the number of groups.
 * @return {Array<number>} The silhouette value for each point.
 *
 * @example
 * silhouette([[0.25], [0.75]], [0, 0]); // => [1.0, 1.0]
 */
function silhouette(points, labels) {
    if (points.length !== labels.length) {
        throw new Error("must have exactly as many labels as points");
    }
    var groupings = createGroups(labels);
    var distances = calculateAllDistances(points);
    var result = [];
    for (var i = 0; i < points.length; i++) {
        var s = 0;
        if (groupings[labels[i]].length > 1) {
            var a = meanDistanceFromPointToGroup(
                i,
                groupings[labels[i]],
                distances
            );
            var b = meanDistanceToNearestGroup(
                i,
                labels,
                groupings,
                distances
            );
            s = (b - a) / Math.max(a, b);
        }
        result.push(s);
    }
    return result;
}

/**
 * Create a lookup table mapping group IDs to point IDs.
 *
 * @private
 * @param {Array<number>} labels Labels of points. This must be the same length as `points`,
 * and values must lie in [0..G-1], where G is the number of groups.
 * @return {Array<Array<number>>} An array of length G, each of whose entries is an array
 * containing the indices of the points in that group.
 */
function createGroups(labels) {
    var numGroups = 1 + max(labels);
    var result = Array(numGroups);
    for (var i = 0; i < labels.length; i++) {
        var label = labels[i];
        if (result[label] === undefined) {
            result[label] = [];
        }
        result[label].push(i);
    }
    return result;
}

/**
 * Create a lookup table of all inter-point distances.
 *
 * @private
 * @param {Array<Array<number>>} points N-dimensional coordinates of points.
 * @return {Array<Array<number>>} A symmetric square array of inter-point distances
 * (zero on the diagonal).
 */
function calculateAllDistances(points) {
    var numPoints = points.length;
    var result = makeMatrix(numPoints, numPoints);
    for (var i = 0; i < numPoints; i++) {
        for (var j = 0; j < i; j++) {
            result[i][j] = euclideanDistance(points[i], points[j]);
            result[j][i] = result[i][j];
        }
    }
    return result;
}

/**
 * Calculate the mean distance between this point and all the points in the
 * nearest group (as determined by which point in another group is closest).
 *
 * @private
 * @param {number} which The index of this point.
 * @param {Array<number>} labels Labels of points.
 * @param {Array<Array<number>>} groupings An array whose entries are arrays
 * containing the indices of the points in that group.
 * @param {Array<Array<number>>} distances A symmetric square array of inter-point
 * distances.
 * @return {number} The mean distance from this point to others in the nearest
 * group.
 */
function meanDistanceToNearestGroup(which, labels, groupings, distances) {
    var label = labels[which];
    var result = Number.MAX_VALUE;
    for (var i = 0; i < groupings.length; i++) {
        if (i !== label) {
            var d = meanDistanceFromPointToGroup(
                which,
                groupings[i],
                distances
            );
            if (d < result) {
                result = d;
            }
        }
    }
    return result;
}

/**
 * Calculate the mean distance between a point and all the points in a group
 * (possibly its own).
 *
 * @private
 * @param {number} which The index of this point.
 * @param {Array<number>} group The indices of all the points in the group in
 * question.
 * @param {Array<Array<number>>} distances A symmetric square array of inter-point
 * distances.
 * @return {number} The mean distance from this point to others in the
 * specified group.
 */
function meanDistanceFromPointToGroup(which, group, distances) {
    var total = 0;
    for (var i = 0; i < group.length; i++) {
        total += distances[which][group[i]];
    }
    return total / group.length;
}

/**
 * Calculate the [silhouette metric](https://en.wikipedia.org/wiki/Silhouette_(clustering))
 * for a set of N-dimensional points arranged in groups. The metric is the largest
 * individual silhouette value for the data.
 *
 * @param {Array<Array<number>>} points N-dimensional coordinates of points.
 * @param {Array<number>} labels Labels of points. This must be the same length as `points`,
 * and values must lie in [0..G-1], where G is the number of groups.
 * @return {number} The silhouette metric for the groupings.
 *
 * @example
 * silhouetteMetric([[0.25], [0.75]], [0, 0]); // => 1.0
 */
function silhouetteMetric(points, labels) {
    var values = silhouette(points, labels);
    return max(values);
}

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

/**
 * Approximate equality.
 *
 * @param {number} actual The value to be tested.
 * @param {number} expected The reference value.
 * @param {number} tolerance The acceptable relative difference.
 * @return {boolean} Whether numbers are within tolerance.
 */
function approxEqual(actual, expected, tolerance) {
    if ( tolerance === void 0 ) tolerance = epsilon;

    return relativeError(actual, expected) <= tolerance;
}

export { BayesianClassifier, PerceptronModel, addToMean, approxEqual, mean as average, meanSimple as averageSimple, BayesianClassifier as bayesian, bernoulliDistribution, binomialDistribution, bisect, chiSquaredDistributionTable, chiSquaredGoodnessOfFit, chunk, ckmeans, coefficientOfVariation, combinations, combinationsReplacement, combineMeans, combineVariances, cumulativeStdLogisticProbability, cumulativeStdNormalProbability, epsilon, equalIntervalBreaks, errorFunction as erf, errorFunction, extent, extentSorted, factorial, gamma, gammaln, geometricMean, harmonicMean, interquartileRange, inverseErrorFunction, interquartileRange as iqr, jenks, kMeansCluster, kernelDensityEstimation as kde, kernelDensityEstimation, linearRegression, linearRegressionLine, logAverage, logit, medianAbsoluteDeviation as mad, max, maxSorted, mean, meanSimple, median, medianAbsoluteDeviation, medianSorted, min, minSorted, mode, modeFast, modeSorted, numericSort, PerceptronModel as perceptron, permutationTest, permutationsHeap, poissonDistribution, probit, product, quantile, quantileRank, quantileRankSorted, quantileSorted, quickselect, rSquared, relativeError, rootMeanSquare as rms, rootMeanSquare, sample, sampleCorrelation, sampleCovariance, sampleKurtosis, sampleRankCorrelation, sampleSkewness, sampleStandardDeviation, sampleVariance, sampleWithReplacement, shuffle, shuffleInPlace, sign, silhouette, silhouetteMetric, standardDeviation, standardNormalTable, subtractFromMean, sum, sumNthPowerDeviations, sumSimple, tTest, tTestTwoSample, uniqueCountSorted, variance, wilcoxonRankSum, zScore };
//# sourceMappingURL=simple-statistics.mjs.map
