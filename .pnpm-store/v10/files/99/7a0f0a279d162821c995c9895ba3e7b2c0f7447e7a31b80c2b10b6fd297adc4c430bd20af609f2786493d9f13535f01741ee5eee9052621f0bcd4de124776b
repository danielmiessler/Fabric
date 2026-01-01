import max from "./max.js";
import silhouette from "./silhouette.js";

/**
 * Calculate the [silhouette metric](https://en.wikipedia.org/wiki/Silhouette_(clustering))
 * for a set of N-dimensional points arranged in groups. The metric is the largest
 * individual silhouette value for the data.
 *
 * @param {Array<Array<number>>} points N-dimensional coordinates of points.
 * @param {Array<number>} labels Labels of points. This must be the same length as `points`,
 * and values must lie in [0..G-1], where G is the number of groups.
 * @return {number} The silhouette metric for the groupings.
 *
 * @example
 * silhouetteMetric([[0.25], [0.75]], [0, 0]); // => 1.0
 */
function silhouetteMetric(points, labels) {
    const values = silhouette(points, labels);
    return max(values);
}

export default silhouetteMetric;
