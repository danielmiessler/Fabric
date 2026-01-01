/**
 * [Bayesian Classifier](http://en.wikipedia.org/wiki/Naive_Bayes_classifier)
 *
 * This is a naïve bayesian classifier that takes
 * singly-nested objects.
 *
 * @class
 * @example
 * var bayes = new BayesianClassifier();
 * bayes.train({
 *   species: 'Cat'
 * }, 'animal');
 * var result = bayes.score({
 *   species: 'Cat'
 * })
 * // result
 * // {
 * //   animal: 1
 * // }
 */
class BayesianClassifier {
    /*:: totalCount: number */
    /*:: data: Object */
    constructor() {
        // The number of items that are currently
        // classified in the model
        this.totalCount = 0;
        // Every item classified in the model
        this.data = {};
    }

    /**
     * Train the classifier with a new item, which has a single
     * dimension of Javascript literal keys and values.
     *
     * @param {Object} item an object with singly-deep properties
     * @param {string} category the category this item belongs to
     * @return {undefined} adds the item to the classifier
     */
    train(item, category) {
        // If the data object doesn't have any values
        // for this category, create a new object for it.
        if (!this.data[category]) {
            this.data[category] = {};
        }

        // Iterate through each key in the item.
        for (const k in item) {
            const v = item[k];
            // Initialize the nested object `data[category][k][item[k]]`
            // with an object of keys that equal 0.
            if (this.data[category][k] === undefined) {
                this.data[category][k] = {};
            }
            if (this.data[category][k][v] === undefined) {
                this.data[category][k][v] = 0;
            }

            // And increment the key for this key/value combination.
            this.data[category][k][v]++;
        }

        // Increment the number of items classified
        this.totalCount++;
    }

    /**
     * Generate a score of how well this item matches all
     * possible categories based on its attributes
     *
     * @param {Object} item an item in the same format as with train
     * @returns {Object} of probabilities that this item belongs to a
     * given category.
     */
    score(item) {
        // Initialize an empty array of odds per category.
        const odds = {};
        let category;
        // Iterate through each key in the item,
        // then iterate through each category that has been used
        // in previous calls to `.train()`
        for (const k in item) {
            const v = item[k];
            for (category in this.data) {
                // Create an empty object for storing key - value combinations
                // for this category.
                odds[category] = {};

                // If this item doesn't even have a property, it counts for nothing,
                // but if it does have the property that we're looking for from
                // the item to categorize, it counts based on how popular it is
                // versus the whole population.
                if (this.data[category][k]) {
                    odds[category][k + "_" + v] =
                        (this.data[category][k][v] || 0) / this.totalCount;
                } else {
                    odds[category][k + "_" + v] = 0;
                }
            }
        }

        // Set up a new object that will contain sums of these odds by category
        const oddsSums = {};

        for (category in odds) {
            // Tally all of the odds for each category-combination pair -
            // the non-existence of a category does not add anything to the
            // score.
            oddsSums[category] = 0;
            for (const combination in odds[category]) {
                oddsSums[category] += odds[category][combination];
            }
        }

        return oddsSums;
    }
}

export default BayesianClassifier;
