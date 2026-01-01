/**
 * Calculate Euclidean distance between two points.
 * @param {Array<number>} left First N-dimensional point.
 * @param {Array<number>} right Second N-dimensional point.
 * @returns {number} Distance.
 */
function euclideanDistance(left, right) {
    let sum = 0;
    for (let i = 0; i < left.length; i++) {
        const diff = left[i] - right[i];
        sum += diff * diff;
    }
    return Math.sqrt(sum);
}

export default euclideanDistance;
