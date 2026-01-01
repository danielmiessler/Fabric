import Item from '../Item';
import { groupByElement, groupByPage } from '../support/groupingUtils';
import EvaluationTracker from './EvaluationTracker';
import ChangeTracker from './ChangeTracker';
import ItemGroup from './ItemGroup';
import ItemMerger from './ItemMerger';

export default interface Page {
  index: number;
  itemGroups: ItemGroup[];
}

export function asPages(
  evaluationTracker: EvaluationTracker,
  changeTracker: ChangeTracker,
  schema: string[],
  items: Item[],
  itemMerger?: ItemMerger,
): Page[] {
  return groupByPage(items).map((pageItems: Item[]) => {
    let itemGroups: ItemGroup[];
    if (itemMerger) {
      itemGroups = groupByElement(pageItems, itemMerger.groupKey).map((groupItems) => {
        if (groupItems.length > 1) {
          const top = itemMerger.merge(evaluationTracker, changeTracker, schema, groupItems);
          return new ItemGroup(top, groupItems);
        } else {
          return new ItemGroup(groupItems[0]);
        }
      });
    } else {
      itemGroups = pageItems.map((item) => new ItemGroup(item));
    }
    return { index: pageItems[0].page, itemGroups } as Page;
  });
}
