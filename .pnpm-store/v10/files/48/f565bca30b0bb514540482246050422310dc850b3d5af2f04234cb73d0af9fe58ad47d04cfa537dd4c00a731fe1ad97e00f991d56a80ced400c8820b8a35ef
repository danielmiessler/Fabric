/**
 * Multi-stage progress. Progress is expressed in a number between 0 and 1.
 */
export default class Progress {
    stages: string[];
    stageDetails: string[];
    stageProgress: number[];
    stageWeights: number[];
    constructor(stages: string[], weights?: number[]);
    isComplete(stageIndex: number): boolean;
    isProgressing(stageIndex: number): boolean;
    totalProgress(): number;
}
