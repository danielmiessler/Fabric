import euclideanDistance from "./euclidean_distance.js";
import makeMatrix from "./make_matrix.js";
import max from "./max.js";

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
    const groupings = createGroups(labels);
    const distances = calculateAllDistances(points);
    const result = [];
    for (let i = 0; i < points.length; i++) {
        let s = 0;
        if (groupings[labels[i]].length > 1) {
            const a = meanDistanceFromPointToGroup(
                i,
                groupings[labels[i]],
                distances
            );
            const b = meanDistanceToNearestGroup(
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
    const numGroups = 1 + max(labels);
    const result = Array(numGroups);
    for (let i = 0; i < labels.length; i++) {
        const label = labels[i];
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
    const numPoints = points.length;
    const result = makeMatrix(numPoints, numPoints);
    for (let i = 0; i < numPoints; i++) {
        for (let j = 0; j < i; j++) {
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
    const label = labels[which];
    let result = Number.MAX_VALUE;
    for (let i = 0; i < groupings.length; i++) {
        if (i !== label) {
            const d = meanDistanceFromPointToGroup(
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
    let total = 0;
    for (let i = 0; i < group.length; i++) {
        total += distances[which][group[i]];
    }
    return total / group.length;
}

export default silhouette;
