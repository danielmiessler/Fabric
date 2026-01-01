import jenksBreaks from "./jenks_breaks.js";
import jenksMatrices from "./jenks_matrices.js";

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
    const matrices = jenksMatrices(data, nClasses);
    // we only need lower class limits here
    const lowerClassLimits = matrices.lowerClassLimits;

    // extract nClasses out of the computed matrices
    return jenksBreaks(data, lowerClassLimits, nClasses);
}

export default jenks;
