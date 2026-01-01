import mean from "./mean.js";
import shuffleInPlace from "./shuffle_in_place.js";

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
    const meanX = mean(sampleX);
    const meanY = mean(sampleY);

    // calculate initial test statistic. This will be our point of comparison with
    // the generated test statistics.
    const testStatistic = meanX - meanY;

    // create test-statistic distribution
    const testStatDsn = new Array(k);

    // combine datsets so we can easily shuffle later
    const allData = sampleX.concat(sampleY);
    const midIndex = Math.floor(allData.length / 2);

    for (let i = 0; i < k; i++) {
        // 1. shuffle data assignments
        shuffleInPlace(allData, randomSource);
        const permLeft = allData.slice(0, midIndex);
        const permRight = allData.slice(midIndex, allData.length);

        // 2.re-calculate test statistic
        const permTestStatistic = mean(permLeft) - mean(permRight);

        // 3. store test statistic to build test statistic distribution
        testStatDsn[i] = permTestStatistic;
    }

    // Calculate p-value depending on alternative
    // For this test, we calculate the percentage of 'extreme' test statistics (subject to our hypothesis)
    // more info on permutation test p-value calculations: https://onlinecourses.science.psu.edu/stat464/node/35
    let numExtremeTStats = 0;
    if (alternative === "two_side") {
        for (let i = 0; i <= k; i++) {
            if (Math.abs(testStatDsn[i]) >= Math.abs(testStatistic)) {
                numExtremeTStats += 1;
            }
        }
    } else if (alternative === "greater") {
        for (let i = 0; i <= k; i++) {
            if (testStatDsn[i] >= testStatistic) {
                numExtremeTStats += 1;
            }
        }
    } else {
        // alternative === 'less'
        for (let i = 0; i <= k; i++) {
            /* c8 ignore start */
            if (testStatDsn[i] <= testStatistic) {
                numExtremeTStats += 1;
            }
            /* c8 ignore end */
        }
    }

    return numExtremeTStats / k;
}

export default permutationTest;
