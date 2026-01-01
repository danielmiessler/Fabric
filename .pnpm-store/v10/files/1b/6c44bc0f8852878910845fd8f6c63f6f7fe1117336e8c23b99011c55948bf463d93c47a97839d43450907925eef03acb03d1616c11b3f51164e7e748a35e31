import ItemMerger from './ItemMerger';
import Item from '../Item';
import EvaluationTracker from './EvaluationTracker';
import ChangeTracker from './ChangeTracker';
import { flatten } from '../support/functional';
import { onlyUniques } from '../support/groupingUtils';

export default class LineItemMerger extends ItemMerger {
  constructor(private trackAsNew = false) {
    super('line');
  }

  merge(evaluationTracker: EvaluationTracker, changeTracker: ChangeTracker, schema: string[], items: Item[]): Item {
    const page = items[0].page;
    const block = items[0].data['block'];
    const line = items[0].data['line'];
    const str = items.map((item) => item.data['str']).join(' ');
    const x = Math.min(...items.map((item) => item.data['x']));
    const y = Math.min(...items.map((item) => item.data['y']));
    const width = items.reduce((sum, item) => sum + item.data['width'], 0);
    const height = Math.max(...items.map((item) => item.data['height']));
    const fontNames = [...new Set(items.map((item) => item.data['fontName']))];
    const directions = [...new Set(items.map((item) => item.data['dir']))];

    const newItem = new Item(page, {
      str,
      block,
      line,
      x,
      y,
      width,
      height,
      fontName: fontNames,
      dir: directions,
    });

    if (schema.includes('types')) {
      const types = flatten(items.map((item) => item.data['types'] || [])).filter(onlyUniques);
      if (types.length > 0) {
        newItem.data['types'] = types;
      }
    }

    const evaluatedItem = items.find((item) => evaluationTracker.evaluated(item));
    if (evaluatedItem) evaluationTracker.trackEvaluation(newItem, evaluationTracker.evaluationScore(evaluatedItem));

    if (this.trackAsNew) {
      changeTracker.trackAddition(newItem);
    } else if (items.every((item) => changeTracker.isRemoved(item))) {
      changeTracker.trackRemoval(newItem);
    } else if (items.find((item) => changeTracker.hasChanged(item))) {
      changeTracker.trackContentChange(newItem);
    }
    return newItem;
  }
}
