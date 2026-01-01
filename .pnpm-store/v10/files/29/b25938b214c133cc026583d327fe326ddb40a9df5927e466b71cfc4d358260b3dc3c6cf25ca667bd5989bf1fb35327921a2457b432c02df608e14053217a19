"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/**
 * Multi-stage progress. Progress is expressed in a number between 0 and 1.
 */
class Progress {
    constructor(stages, weights = []) {
        this.stages = stages;
        this.stageDetails = new Array(stages.length);
        this.stageProgress = new Array(stages.length).fill(0);
        if (weights.length === 0) {
            this.stageWeights = new Array(stages.length).fill(1 / stages.length);
        }
        else {
            if (weights.length !== stages.length)
                throw new Error(`Provided only ${weights.length} weights but expected ${stages.length} for ${stages.length} stages`);
            const weightsSummed = weights.reduce((sum, weight) => +(sum + weight).toFixed(12), 0);
            if (weightsSummed !== 1)
                throw new Error(`Weights [${weights.join(', ')}] should sum up to 1, but did to ${weightsSummed}`);
            this.stageWeights = weights;
        }
    }
    isComplete(stageIndex) {
        return this.stageProgress[stageIndex] === 1;
    }
    isProgressing(stageIndex) {
        const previousComplete = stageIndex === 0 || this.isComplete(stageIndex - 1);
        return previousComplete && this.stageProgress[stageIndex] < 1;
    }
    totalProgress() {
        const stageCount = this.stages.length;
        const stageProgressSummed = this.stageProgress.reduce((sum, stageProgress, index) => sum + stageProgress * this.stageWeights[index] * this.stages.length, 0);
        return stageProgressSummed / stageCount;
    }
}
exports.default = Progress;
