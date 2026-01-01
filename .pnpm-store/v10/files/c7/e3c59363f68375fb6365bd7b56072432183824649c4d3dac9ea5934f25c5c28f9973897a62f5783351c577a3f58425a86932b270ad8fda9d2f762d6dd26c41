import sampleCorrelation from "./sample_correlation.js";

/**
 * The [rank correlation](https://en.wikipedia.org/wiki/Rank_correlation) is
 * a measure of the strength of monotonic relationship between two arrays
 *
 * @param {Array<number>} x first input
 * @param {Array<number>} y second input
 * @returns {number} sample rank correlation
 */
function sampleRankCorrelation(x, y) {
    const xIndexes = x
        .map((value, index) => [value, index])
        .sort((a, b) => a[0] - b[0])
        .map((pair) => pair[1]);
    const yIndexes = y
        .map((value, index) => [value, index])
        .sort((a, b) => a[0] - b[0])
        .map((pair) => pair[1]);

    // At this step, we have an array of indexes
    // that map from sorted numbers to their original indexes. We reverse
    // that so that it is an array of the sorted destination index.
    const xRanks = Array(xIndexes.length);
    const yRanks = Array(xIndexes.length);
    for (let i = 0; i < xIndexes.length; i++) {
        xRanks[xIndexes[i]] = i;
        yRanks[yIndexes[i]] = i;
    }

    return sampleCorrelation(xRanks, yRanks);
}

export default sampleRankCorrelation;
