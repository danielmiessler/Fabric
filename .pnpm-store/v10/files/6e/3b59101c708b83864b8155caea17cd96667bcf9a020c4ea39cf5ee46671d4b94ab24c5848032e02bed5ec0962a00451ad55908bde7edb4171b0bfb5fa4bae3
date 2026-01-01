import Item from '../Item';
import EvaluationIndex from './EvaluationIndex';

export default class EvaluationTracker implements EvaluationIndex {
  private evaluations: Map<string, unknown> = new Map();
  private scored = false;

  evaluationCount() {
    return this.evaluations.size;
  }

  hasScores(): boolean {
    return this.scored;
  }

  evaluated(item: Item): boolean {
    return this.evaluations.has(item.uuid);
  }

  evaluationScore(item: Item) {
    return this.evaluations.get(item.uuid);
  }

  trackEvaluation(item: Item, score: unknown = undefined) {
    if (typeof score !== 'undefined') {
      this.scored = true;
    }
    this.evaluations.set(item.uuid, score);
  }
}
