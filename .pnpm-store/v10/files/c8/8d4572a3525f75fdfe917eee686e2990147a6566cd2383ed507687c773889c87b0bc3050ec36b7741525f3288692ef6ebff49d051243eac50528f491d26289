"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class EvaluationTracker {
    constructor() {
        this.evaluations = new Map();
        this.scored = false;
    }
    evaluationCount() {
        return this.evaluations.size;
    }
    hasScores() {
        return this.scored;
    }
    evaluated(item) {
        return this.evaluations.has(item.uuid);
    }
    evaluationScore(item) {
        return this.evaluations.get(item.uuid);
    }
    trackEvaluation(item, score = undefined) {
        if (typeof score !== 'undefined') {
            this.scored = true;
        }
        this.evaluations.set(item.uuid, score);
    }
}
exports.default = EvaluationTracker;
