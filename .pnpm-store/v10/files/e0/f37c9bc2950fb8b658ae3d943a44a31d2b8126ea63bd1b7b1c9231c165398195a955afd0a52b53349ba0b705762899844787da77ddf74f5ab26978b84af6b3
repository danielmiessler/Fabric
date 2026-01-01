import ItemMerger from './ItemMerger';
import Item from '../Item';
import EvaluationTracker from './EvaluationTracker';
import ChangeTracker from './ChangeTracker';
export default class LineItemMerger extends ItemMerger {
    private trackAsNew;
    constructor(trackAsNew?: boolean);
    merge(evaluationTracker: EvaluationTracker, changeTracker: ChangeTracker, schema: string[], items: Item[]): Item;
}
