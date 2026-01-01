/*
 * Pull Breaks Values for Jenks
 *
 * the second part of the jenks recipe: take the calculated matrices
 * and derive an array of n breaks.
 *
 * @private
 */
function jenksBreaks(data, lowerClassLimits, nClasses) {
    let k = data.length;
    const kclass = [];
    let countNum = nClasses;

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

export default jenksBreaks;
