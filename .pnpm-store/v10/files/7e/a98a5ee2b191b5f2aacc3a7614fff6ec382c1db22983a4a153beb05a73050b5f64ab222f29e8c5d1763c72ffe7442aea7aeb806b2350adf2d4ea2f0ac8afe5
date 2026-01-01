import LineItemMerger from 'src/debug/LineItemMerger';
import EvaluationTracker from 'src/debug/EvaluationTracker';
import ChangeTracker from 'src/debug/ChangeTracker';
import Item from 'src/Item';
import { items, realisticItems } from '../testItems';
import { Addition, ContentChange } from 'src/debug/ChangeIndex';

const defaultSchema = ['line', 'x', 'y', 'str', 'fontNmae', 'dir', 'width', 'height'];

test('Basics', async () => {
  const itemMerger = new LineItemMerger();
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  expect(itemMerger.groupKey).toEqual('line');
  const mergedItem = itemMerger.merge(
    evaluationTracker,
    changeTracker,
    defaultSchema,
    items(0, [
      {
        line: 2,
        x: 240,
        y: 585,
        str: 'Dies ist eine Test-PDF',
        fontName: 'g_d0_f2',
        dir: 'ltr',
        width: 108.62,
        height: 11,
      },
      {
        line: 2,
        x: 352.69,
        y: 585,
        str: '.',
        fontName: 'g_d0_f2',
        dir: 'ltr',
        width: 3.06,
        height: 11,
      },
      {
        line: 2,
        x: 348,
        y: 588,
        str: '1',
        fontName: 'g_d0_f2',
        dir: 'ltr',
        width: 4.08,
        height: 7.33,
      },
    ]),
  );
  expect(mergedItem.withoutUuid()).toEqual(
    new Item(0, {
      line: 2,
      x: 240,
      y: 585,
      str: 'Dies ist eine Test-PDF . 1',
      fontName: ['g_d0_f2'],
      dir: ['ltr'],
      width: 115.76,
      height: 11,
    }).withoutUuid(),
  );
});

test('Track all lines as changes', async () => {
  const itemMerger = new LineItemMerger(true);
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const mergedItem = itemMerger.merge(
    evaluationTracker,
    changeTracker,
    defaultSchema,
    realisticItems(0, [{ line: 1 }, { line: 1 }]),
  );
  expect(changeTracker.change(mergedItem)).toEqual(new Addition());
});

test('Mark lines containing evaluated items as evaluated', async () => {
  const itemMerger = new LineItemMerger();
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const items1 = realisticItems(0, [{ line: 1 }, { line: 1 }]);
  const items2 = realisticItems(0, [{ line: 2 }, { line: 2 }]);
  evaluationTracker.trackEvaluation(items1[1]);
  const mergedItem1 = itemMerger.merge(evaluationTracker, changeTracker, defaultSchema, items1);
  const mergedItem2 = itemMerger.merge(evaluationTracker, changeTracker, defaultSchema, items2);
  expect(evaluationTracker.evaluated(mergedItem1)).toBeTruthy();
  expect(evaluationTracker.evaluated(mergedItem2)).toBeFalsy();
});

test('Mark lines containing changed items as changed', async () => {
  const itemMerger = new LineItemMerger();
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const items1 = realisticItems(0, [{ line: 1 }, { line: 1 }]);
  const items2 = realisticItems(0, [{ line: 2 }, { line: 2 }]);
  changeTracker.trackPositionalChange(items1[1], 1, 0);
  const mergedItem1 = itemMerger.merge(evaluationTracker, changeTracker, defaultSchema, items1);
  const mergedItem2 = itemMerger.merge(evaluationTracker, changeTracker, defaultSchema, items2);
  expect(changeTracker.change(mergedItem1)).toEqual(new ContentChange());
  expect(changeTracker.change(mergedItem2)).toEqual(undefined);
});
