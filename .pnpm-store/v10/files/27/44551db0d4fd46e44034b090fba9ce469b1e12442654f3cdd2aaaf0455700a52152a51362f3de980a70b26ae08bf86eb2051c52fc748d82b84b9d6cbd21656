import Item from '../Item';
import EvaluationTracker from './EvaluationTracker';
import ChangeTracker from './ChangeTracker';
import ItemGroup from './ItemGroup';
import ItemMerger from './ItemMerger';
export default interface Page {
    index: number;
    itemGroups: ItemGroup[];
}
export declare function asPages(evaluationTracker: EvaluationTracker, changeTracker: ChangeTracker, schema: string[], items: Item[], itemMerger?: ItemMerger): Page[];
