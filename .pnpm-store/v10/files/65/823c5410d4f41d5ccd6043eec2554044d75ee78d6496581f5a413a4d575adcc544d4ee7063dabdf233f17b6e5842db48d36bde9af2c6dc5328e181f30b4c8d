import Item from 'src/Item';
import Page, { asPages } from 'src/debug/Page';
import ItemGroup from 'src/debug/ItemGroup';
import ItemMerger from 'src/debug/ItemMerger';
import { items } from 'test/testItems';
import EvaluationTracker from 'src/debug/EvaluationTracker';
import ChangeTracker from 'src/debug/ChangeTracker';

test('empty', async () => {
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  expect(asPages(evaluationTracker, changeTracker, [], [])).toEqual([]);
});

test('no merger', async () => {
  const pageItems = [
    items(0, [{ id: 1, line: 1 }]),
    items(1, [
      { id: 2, line: 1 },
      { id: 3, line: 1 },
      { id: 4, line: 1 },
    ]),
    items(2, [{ id: 5, line: 1 }]),
  ];
  const flattenedItems = new Array<Item>().concat(...pageItems);
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const pages = asPages(evaluationTracker, changeTracker, ['id', 'line'], flattenedItems);
  expect(pages).toEqual([
    { index: 0, itemGroups: pageItems[0].map((item) => new ItemGroup(item)) },
    { index: 1, itemGroups: pageItems[1].map((item) => new ItemGroup(item)) },
    { index: 2, itemGroups: pageItems[2].map((item) => new ItemGroup(item)) },
  ] as Page[]);
  expect(changeTracker.changeCount()).toEqual(0);
});

test('merger', async () => {
  const pageItems = [
    items(0, [{ id: 1, line: 1 }]),
    items(1, [
      { id: 2, line: 1 },
      { id: 3, line: 1 },
      { id: 4, line: 2 },
    ]),
    items(2, [{ id: 5, line: 1 }]),
  ];
  const flattenedItems = new Array<Item>().concat(...pageItems);
  const merger: ItemMerger = { groupKey: 'line', merge: (items) => items[0] };
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const pages = asPages(evaluationTracker, changeTracker, ['id', 'line'], flattenedItems, merger);

  expect(pages).toEqual([
    { index: 0, itemGroups: pageItems[0].map((item) => new ItemGroup(item)) },
    {
      index: 1,
      itemGroups: [
        new ItemGroup(
          merger.merge(evaluationTracker, changeTracker, ['id', 'line'], pageItems[1].slice(0, 2)),
          pageItems[1].slice(0, 2),
        ),
        new ItemGroup(pageItems[1][2]),
      ],
    },
    { index: 2, itemGroups: pageItems[2].map((item) => new ItemGroup(item)) },
  ] as Page[]);
});
