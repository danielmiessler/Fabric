import max from "./max.js";
import min from "./min.js";

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

    const theMin = min(x);
    const theMax = max(x);

    // the first break will always be the minimum value
    // in the xset
    const breaks = [theMin];

    // The size of each break is the full range of the x
    // divided by the number of classes requested
    const breakSize = (theMax - theMin) / nClasses;

    // In the case of nClasses = 1, this loop won't run
    // and the returned breaks will be [min, max]
    for (let i = 1; i < nClasses; i++) {
        breaks.push(breaks[0] + breakSize * i);
    }

    // the last break will always be the
    // maximum.
    breaks.push(theMax);

    return breaks;
}

export default equalIntervalBreaks;
