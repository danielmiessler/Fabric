/**
 * This is a single-layer [Perceptron Classifier](http://en.wikipedia.org/wiki/Perceptron) that takes
 * arrays of numbers and predicts whether they should be classified
 * as either 0 or 1 (negative or positive examples).
 * @class
 * @example
 * // Create the model
 * var p = new PerceptronModel();
 * // Train the model with input with a diagonal boundary.
 * for (var i = 0; i < 5; i++) {
 *     p.train([1, 1], 1);
 *     p.train([0, 1], 0);
 *     p.train([1, 0], 0);
 *     p.train([0, 0], 0);
 * }
 * p.predict([0, 0]); // 0
 * p.predict([0, 1]); // 0
 * p.predict([1, 0]); // 0
 * p.predict([1, 1]); // 1
 */
class PerceptronModel {
    /*:: bias: number */
    /*:: weights: Array<number> */
    constructor() {
        // The weights, or coefficients of the model;
        // weights are only populated when training with data.
        this.weights = [];
        // The bias term, or intercept; it is also a weight but
        // it's stored separately for convenience as it is always
        // multiplied by one.
        this.bias = 0;
    }
    /**
     * **Predict**: Use an array of features with the weight array and bias
     * to predict whether an example is labeled 0 or 1.
     *
     * @param {Array<number>} features an array of features as numbers
     * @returns {number} 1 if the score is over 0, otherwise 0
     */
    predict(features) {
        // Only predict if previously trained
        // on the same size feature array(s).
        if (features.length !== this.weights.length) {
            return null;
        }

        // Calculate the sum of features times weights,
        // with the bias added (implicitly times one).
        let score = 0;
        for (let i = 0; i < this.weights.length; i++) {
            score += this.weights[i] * features[i];
        }
        score += this.bias;

        // Classify as 1 if the score is over 0, otherwise 0.
        if (score > 0) {
            return 1;
        } else {
            return 0;
        }
    }

    /**
     * **Train** the classifier with a new example, which is
     * a numeric array of features and a 0 or 1 label.
     *
     * @param {Array<number>} features an array of features as numbers
     * @param {number} label either 0 or 1
     * @returns {PerceptronModel} this
     */
    train(features, label) {
        // Require that only labels of 0 or 1 are considered.
        if (label !== 0 && label !== 1) {
            return null;
        }
        // The length of the feature array determines
        // the length of the weight array.
        // The perceptron will continue learning as long as
        // it keeps seeing feature arrays of the same length.
        // When it sees a new data shape, it initializes.
        if (features.length !== this.weights.length) {
            this.weights = features;
            this.bias = 1;
        }
        // Make a prediction based on current weights.
        const prediction = this.predict(features);
        // Update the weights if the prediction is wrong.
        if (typeof prediction === "number" && prediction !== label) {
            const gradient = label - prediction;
            for (let i = 0; i < this.weights.length; i++) {
                this.weights[i] += gradient * features[i];
            }
            this.bias += gradient;
        }
        return this;
    }
}

export default PerceptronModel;
