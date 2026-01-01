import euclideanDistance from "./euclidean_distance.js";
import makeMatrix from "./make_matrix.js";
import sample from "./sample.js";

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
function kMeansCluster(points, numCluster, randomSource = Math.random) {
    let oldCentroids = null;
    let newCentroids = sample(points, numCluster, randomSource);
    let labels = null;
    let change = Number.MAX_VALUE;
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
    return points.map((p) => {
        let minDist = Number.MAX_VALUE;
        let label = -1;
        for (let i = 0; i < centroids.length; i++) {
            const dist = euclideanDistance(p, centroids[i]);
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
    const dimension = points[0].length;
    const centroids = makeMatrix(numCluster, dimension);
    const counts = Array(numCluster).fill(0);

    // Add points to centroids' accumulators and count points per centroid.
    const numPoints = points.length;
    for (let i = 0; i < numPoints; i++) {
        const point = points[i];
        const label = labels[i];
        const current = centroids[label];
        for (let j = 0; j < dimension; j++) {
            current[j] += point[j];
        }
        counts[label] += 1;
    }

    // Rescale centroids, checking for any that have no points.
    for (let i = 0; i < numCluster; i++) {
        if (counts[i] === 0) {
            throw new Error(`Centroid ${i} has no friends`);
        }
        const centroid = centroids[i];
        for (let j = 0; j < dimension; j++) {
            centroid[j] /= counts[i];
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
    let total = 0;
    for (let i = 0; i < left.length; i++) {
        total += euclideanDistance(left[i], right[i]);
    }
    return total;
}

export default kMeansCluster;
