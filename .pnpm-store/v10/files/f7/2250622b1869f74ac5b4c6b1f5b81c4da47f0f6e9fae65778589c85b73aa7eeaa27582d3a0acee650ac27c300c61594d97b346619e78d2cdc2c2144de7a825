import StageResult from 'src/debug/StageResult';
import { toDescriptor } from 'src/TransformDescriptor';
import EvaluationTracker from 'src/debug/EvaluationTracker';
import ChangeTracker from 'src/debug/ChangeTracker';
import AnnotatedColumn from 'src/debug/AnnotatedColumn';
import Page, { asPages } from 'src/debug/Page';
import { items } from '../testItems';
import LineItemMerger from 'src/debug/LineItemMerger';
import Globals from 'src/Globals';

test('itemsUnpacked', async () => {
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const itemMerger = new LineItemMerger(false);
  const descriptor = toDescriptor({ debug: { itemMerger } });
  const schema: AnnotatedColumn[] = [{ name: 'A' }];
  const flatItems = [
    ...items(0, [
      { idx: 0, line: 1 },
      { idx: 1, line: 1 },
      { idx: 2, line: 2 },
    ]),
    ...items(1, [{ idx: 3, line: 1 }]),
    ...items(2, [
      { idx: 4, line: 1 },
      { idx: 5, line: 1 },
    ]),
  ];
  const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
  const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

  expect(result.itemsUnpacked().map((item) => item.data['idx'])).toEqual([0, 1, 2, 3, 4, 5]);
  expect(result.itemsCleanedAndUnpacked().map((item) => item.data['idx'])).toEqual([0, 1, 2, 3, 4, 5]);
});

test('itemsCleanedAndUnpacked', async () => {
  const evaluationTracker = new EvaluationTracker();
  const changeTracker = new ChangeTracker();
  const itemMerger = new LineItemMerger(false);
  const descriptor = toDescriptor({ debug: { itemMerger } });
  const schema: AnnotatedColumn[] = [{ name: 'A' }];
  const flatItems = [
    ...items(0, [
      { idx: 0, line: 1 },
      { idx: 1, line: 1 },
      { idx: 2, line: 2 },
    ]),
    ...items(1, [{ idx: 3, line: 1 }]),
    ...items(2, [
      { idx: 4, line: 1 },
      { idx: 5, line: 1 },
    ]),
  ];
  const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
  changeTracker.trackRemoval(flatItems[1]);
  changeTracker.trackRemoval(flatItems[4]);
  const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

  expect(result.itemsUnpacked().map((item) => item.data['idx'])).toEqual([0, 1, 2, 3, 4, 5]);
  expect(result.itemsCleanedAndUnpacked().map((item) => item.data['idx'])).toEqual([0, 2, 3, 5]);
});

describe('select pages', () => {
  function groupElements(page: Page, elementName: string) {
    return page.itemGroups.map((group) => group.unpacked().map((item) => item.data['idx']));
  }

  test('Evaluation+Changes', async () => {
    const evaluationTracker = new EvaluationTracker();
    const changeTracker = new ChangeTracker();
    const itemMerger = new LineItemMerger(false);
    const descriptor = toDescriptor({ debug: { itemMerger } });
    const schema: AnnotatedColumn[] = [{ name: 'A' }];
    const flatItems = items(0, [
      { idx: 0, line: 1 }, // nada
      { idx: 1, line: 2 }, // eval
      { idx: 2, line: 3 }, // eval + change
      { idx: 3, line: 4 }, // eval
      { idx: 4, line: 4 }, // change
      { idx: 5, line: 4 },
    ]);
    evaluationTracker.trackEvaluation(flatItems[1]);
    evaluationTracker.trackEvaluation(flatItems[2]);
    evaluationTracker.trackEvaluation(flatItems[3]);
    changeTracker.trackAddition(flatItems[2]);
    changeTracker.trackAddition(flatItems[4]);
    const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
    const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

    const allGrouped = result.selectPages(false, true);
    expect(allGrouped.map((page) => page.index)).toEqual([0]);
    expect(groupElements(allGrouped[0], 'idx')).toEqual([[0], [1], [2], [3, 4, 5]]);

    const relevantGrouped = result.selectPages(true, true);
    expect(relevantGrouped.map((page) => page.index)).toEqual([0]);
    expect(groupElements(relevantGrouped[0], 'idx')).toEqual([[1], [2], [3, 4, 5]]);

    const relevantUnpacked = result.selectPages(true, false);
    expect(relevantUnpacked.map((page) => page.index)).toEqual([0]);
    expect(groupElements(relevantUnpacked[0], 'idx')).toEqual([[1], [2], [3], [4]]);

    const allUnpacked = result.selectPages(false, false);
    expect(allUnpacked.map((page) => page.index)).toEqual([0]);
    expect(groupElements(allUnpacked[0], 'idx')).toEqual([[0], [1], [2], [3], [4], [5]]);
  });

  test('Changes on group level', async () => {
    const evaluationTracker = new EvaluationTracker();
    const changeTracker = new ChangeTracker();
    const itemMerger = new LineItemMerger(true);
    const descriptor = toDescriptor({ debug: { itemMerger } });
    const schema: AnnotatedColumn[] = [{ name: 'A' }];
    const flatItems = [
      ...items(0, [
        { idx: 0, line: 1 },
        { idx: 1, line: 1 },
        { idx: 2, line: 2 },
      ]),
      ...items(1, [{ idx: 3, line: 1 }]),
      ...items(2, [
        { idx: 4, line: 1 },
        { idx: 5, line: 1 },
      ]),
    ];
    const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
    const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

    const allGrouped = result.selectPages(false, true);
    expect(allGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(allGrouped[0], 'idx')).toEqual([[0, 1], [2]]);
    expect(groupElements(allGrouped[1], 'idx')).toEqual([[3]]);
    expect(groupElements(allGrouped[2], 'idx')).toEqual([[4, 5]]);

    const relevantGrouped = result.selectPages(true, true);
    expect(relevantGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantGrouped[0], 'idx')).toEqual([[0, 1]]);
    expect(groupElements(relevantGrouped[1], 'idx')).toEqual([]);
    expect(groupElements(relevantGrouped[2], 'idx')).toEqual([[4, 5]]);

    const relevantUnpacked = result.selectPages(true, false);
    expect(relevantUnpacked.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantUnpacked[0], 'idx')).toEqual([]);
    expect(groupElements(relevantUnpacked[1], 'idx')).toEqual([]);
    expect(groupElements(relevantUnpacked[2], 'idx')).toEqual([]);

    const allUnpacked = result.selectPages(false, false);
    expect(allUnpacked.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(allUnpacked[0], 'idx')).toEqual([[0], [1], [2]]);
    expect(groupElements(allUnpacked[1], 'idx')).toEqual([[3]]);
    expect(groupElements(allUnpacked[2], 'idx')).toEqual([[4], [5]]);
  });

  test('Changes on element level', async () => {
    const evaluationTracker = new EvaluationTracker();
    const changeTracker = new ChangeTracker();
    const itemMerger = new LineItemMerger(false);
    const descriptor = toDescriptor({ debug: { itemMerger } });
    const schema: AnnotatedColumn[] = [{ name: 'A' }];
    const flatItems = [
      ...items(0, [
        { idx: 0, line: 1 },
        { idx: 1, line: 1 },
        { idx: 2, line: 2 },
      ]),
      ...items(1, [{ idx: 3, line: 1 }]),
      ...items(2, [
        { idx: 4, line: 1 },
        { idx: 5, line: 1 },
        { idx: 6, line: 2 },
      ]),
    ];
    changeTracker.trackAddition(flatItems[3]);
    changeTracker.trackAddition(flatItems[5]);
    const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
    const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

    const allGrouped = result.selectPages(false, true);
    expect(allGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(allGrouped[0], 'idx')).toEqual([[0, 1], [2]]);
    expect(groupElements(allGrouped[1], 'idx')).toEqual([[3]]);
    expect(groupElements(allGrouped[2], 'idx')).toEqual([[4, 5], [6]]);

    const relevantGrouped = result.selectPages(true, true);
    expect(relevantGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantGrouped[0], 'idx')).toEqual([]);
    expect(groupElements(relevantGrouped[1], 'idx')).toEqual([[3]]);
    expect(groupElements(relevantGrouped[2], 'idx')).toEqual([[4, 5]]);

    const relevantUnpacked = result.selectPages(true, false);
    expect(relevantUnpacked.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantUnpacked[0], 'idx')).toEqual([]);
    expect(groupElements(relevantUnpacked[1], 'idx')).toEqual([[3]]);
    expect(groupElements(relevantUnpacked[2], 'idx')).toEqual([[5]]);

    const allUnpacked = result.selectPages(false, false);
    expect(allUnpacked.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(allUnpacked[0], 'idx')).toEqual([[0], [1], [2]]);
    expect(groupElements(allUnpacked[1], 'idx')).toEqual([[3]]);
    expect(groupElements(allUnpacked[2], 'idx')).toEqual([[4], [5], [6]]);
  });

  test('showAll - grouped - merger', async () => {
    const evaluationTracker = new EvaluationTracker();
    const changeTracker = new ChangeTracker();
    const itemMerger = new LineItemMerger(false);
    const descriptor = toDescriptor({ debug: { itemMerger, showAll: true } });
    const schema: AnnotatedColumn[] = [{ name: 'A' }];
    const flatItems = [
      ...items(0, [
        { idx: 0, line: 1 },
        { idx: 1, line: 1 },
        { idx: 2, line: 2 },
      ]),
      ...items(1, [{ idx: 3, line: 1 }]),
      ...items(2, [
        { idx: 4, line: 1 },
        { idx: 5, line: 1 },
      ]),
    ];
    const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems, itemMerger);
    const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

    const relevantGrouped = result.selectPages(true, true);
    expect(relevantGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantGrouped[0], 'idx')).toEqual([[0, 1], [2]]);
    expect(groupElements(relevantGrouped[1], 'idx')).toEqual([[3]]);
    expect(groupElements(relevantGrouped[2], 'idx')).toEqual([[4, 5]]);
  });

  test('showAll - grouped - no merger', async () => {
    const evaluationTracker = new EvaluationTracker();
    const changeTracker = new ChangeTracker();
    const descriptor = toDescriptor({ debug: { showAll: true } });
    const schema: AnnotatedColumn[] = [{ name: 'A' }];
    const flatItems = [
      ...items(0, [{ idx: 0 }, { idx: 1 }, { idx: 2 }]),
      ...items(1, [{ idx: 3 }]),
      ...items(2, [{ idx: 4 }, { idx: 5 }]),
    ];
    const pages = asPages(evaluationTracker, changeTracker, ['idx', 'line'], flatItems);
    const result = new StageResult(new Globals(), descriptor, schema, pages, evaluationTracker, changeTracker, []);

    const relevantGrouped = result.selectPages(true, true);
    expect(relevantGrouped.map((page) => page.index)).toEqual([0, 1, 2]);
    expect(groupElements(relevantGrouped[0], 'idx')).toEqual([[0], [1], [2]]);
    expect(groupElements(relevantGrouped[1], 'idx')).toEqual([[3]]);
    expect(groupElements(relevantGrouped[2], 'idx')).toEqual([[4], [5]]);
  });
});
