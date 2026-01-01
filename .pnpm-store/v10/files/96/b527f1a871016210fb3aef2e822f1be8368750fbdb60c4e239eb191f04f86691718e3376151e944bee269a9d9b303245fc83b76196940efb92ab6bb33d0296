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

    const pooledSamples = sampleX
        .map((x) => ({ label: "x", value: x }))
        .concat(sampleY.map((y) => ({ label: "y", value: y })))
        .sort((a, b) => a.value - b.value);

    for (let rank = 0; rank < pooledSamples.length; rank++) {
        pooledSamples[rank].rank = rank;
    }

    let tiedRanks = [pooledSamples[0].rank];
    for (let i = 1; i < pooledSamples.length; i++) {
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
        const average = (tiedRanks[0] + tiedRanks[tiedRanks.length - 1]) / 2;
        for (let i = 0; i < tiedRanks.length; i++) {
            pooledSamples[tiedRanks[i]].rank = average;
        }
    }

    let rankSum = 0;

    for (let i = 0; i < pooledSamples.length; i++) {
        const sample = pooledSamples[i];
        if (sample.label === "x") {
            rankSum += sample.rank + 1;
        }
    }

    return rankSum;
}

export default wilcoxonRankSum;
