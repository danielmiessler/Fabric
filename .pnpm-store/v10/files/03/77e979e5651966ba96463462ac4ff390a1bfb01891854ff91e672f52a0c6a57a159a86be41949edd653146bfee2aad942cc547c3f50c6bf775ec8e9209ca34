/**
 * Multi-stage progress. Progress is expressed in a number between 0 and 1.
 */
export default class Progress {
  stages: string[];
  stageDetails: string[];
  stageProgress: number[];
  stageWeights: number[];

  constructor(stages: string[], weights: number[] = []) {
    this.stages = stages;
    this.stageDetails = new Array(stages.length);
    this.stageProgress = new Array(stages.length).fill(0);
    if (weights.length === 0) {
      this.stageWeights = new Array(stages.length).fill(1 / stages.length);
    } else {
      if (weights.length !== stages.length)
        throw new Error(
          `Provided only ${weights.length} weights but expected ${stages.length} for ${stages.length} stages`,
        );
      const weightsSummed = weights.reduce((sum, weight) => +(sum + weight).toFixed(12), 0);
      if (weightsSummed !== 1)
        throw new Error(`Weights [${weights.join(', ')}] should sum up to 1, but did to ${weightsSummed}`);
      this.stageWeights = weights;
    }
  }

  isComplete(stageIndex: number) {
    return this.stageProgress[stageIndex] === 1;
  }

  isProgressing(stageIndex: number) {
    const previousComplete = stageIndex === 0 || this.isComplete(stageIndex - 1);
    return previousComplete && this.stageProgress[stageIndex] < 1;
  }

  totalProgress() {
    const stageCount = this.stages.length;
    const stageProgressSummed = this.stageProgress.reduce(
      (sum, stageProgress, index) => sum + stageProgress * this.stageWeights[index] * this.stages.length,
      0,
    );
    return stageProgressSummed / stageCount;
  }
}
