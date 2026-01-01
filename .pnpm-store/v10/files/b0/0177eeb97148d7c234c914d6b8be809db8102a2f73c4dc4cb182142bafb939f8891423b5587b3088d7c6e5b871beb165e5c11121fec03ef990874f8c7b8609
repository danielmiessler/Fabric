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
    const lowerClassLimits = [];
    const varianceCombinations = [];
    // loop counters
    let i;
    let j;
    // the variance, as computed at each step in the calculation
    let variance = 0;

    // Initialize and fill each matrix with zeroes
    for (i = 0; i < data.length + 1; i++) {
        const tmp1 = [];
        const tmp2 = [];
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

    for (let l = 2; l < data.length + 1; l++) {
        // `SZ` originally. this is the sum of the values seen thus
        // far when calculating variance.
        let sum = 0;
        // `ZSQ` originally. the sum of squares of values seen
        // thus far
        let sumSquares = 0;
        // `WT` originally. This is the number of
        let w = 0;
        // `IV` originally
        let i4 = 0;

        // in several instances, you could say `Math.pow(x, 2)`
        // instead of `x * x`, but this is slower in some browsers
        // introduces an unnecessary concept.
        for (let m = 1; m < l + 1; m++) {
            // `III` originally
            const lowerClassLimit = l - m + 1;
            const val = data[lowerClassLimit - 1];

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

export default jenksMatrices;
