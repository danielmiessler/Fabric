import type EvaluationTracker from './EvaluationTracker';
import type ChangeTracker from './ChangeTracker';
import type Item from '../Item';
/**
 * Groups individual items and merges them to a kind of top level summary item.
 */
export default abstract class ItemMerger {
    groupKey: string;
    constructor(groupKey: string);
    abstract merge(evaluationTracker: EvaluationTracker, changeTracker: ChangeTracker, schema: string[], items: Item[]): Item;
}
