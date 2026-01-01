import Debugger from 'src/Debugger';
import Item from 'src/Item';
import ItemTransformer from 'src/transformer/ItemTransformer';
import TransformDescriptor from 'src/TransformDescriptor';
import TransformContext from 'src/transformer/TransformContext';
import LineItemMerger from 'src/debug/LineItemMerger';
import ItemResult from 'src/ItemResult';
import ColumnAnnotation from 'src/debug/ColumnAnnotation';
import AnnotatedColumn from 'src/debug/AnnotatedColumn';
import { items } from './testItems';

class TestTransformer extends ItemTransformer {
  items: Item[];
  constructor(name: string, descriptor: Partial<TransformDescriptor>, outputSchema: string[], items: Item[]) {
    super(name, `Description for ${name}`, descriptor, (_) => outputSchema);
    this.items = items;
  }
  transform(_: TransformContext, __: Item[]): ItemResult {
    return {
      items: this.items,
      messages: [],
    };
  }
}

describe('Transform Items', () => {
  test('Basics', async () => {
    const parsedSchema = ['A', 'B'];
    const parsedItems = items(0, [
      { A: 'a_row1', B: 'b_row1' },
      { A: 'a_row2', B: 'b_row2' },
    ]);

    const trans1Desc = { requireColumns: ['A', 'B'] };
    const trans1Schema = ['C'];
    const trans1Items = parsedItems.map((item) => item.withData({ C: `c=${item.value('A')}+${item.value('B')}` }));

    const transformers = [new TestTransformer('Trans1', trans1Desc, trans1Schema, trans1Items)];
    const debug = new Debugger(new Map(), [], 1, parsedSchema, parsedItems, transformers);

    expect(debug.stageNames).toEqual(['Parse Result', 'Trans1']);
    expect(debug.stageResult(0).schema).toEqual(parsedSchema.map((column) => ({ name: column })));
    expect(debug.stageResult(1).schema).toEqual([
      ...parsedSchema.map((column) => ({ name: column, annotation: ColumnAnnotation.REMOVED })),
      { name: 'C', annotation: ColumnAnnotation.ADDED },
    ]);

    expect(debug.stageResult(0).itemsUnpacked()).toEqual(parsedItems);
    expect(debug.stageResult(1).itemsUnpacked()).toEqual(trans1Items);
  });

  test('Line Merge', async () => {
    const parsedSchema = ['id', 'y'];
    const parsedItems = items(0, [
      { id: 1, y: 1 },
      { id: 2, y: 2 },
      { id: 3, y: 2 },
    ]);

    const trans1Desc = { requireColumns: ['id', 'y'], debug: { itemMerger: new LineItemMerger(true) } };
    const trans1Schema = ['id', 'line'];
    const trans1Items = parsedItems.map((item) => item.withData({ line: item.data['y'] }));

    const transformers = [new TestTransformer('Trans1', trans1Desc, trans1Schema, trans1Items)];
    const debug = new Debugger(new Map(), [], 1, parsedSchema, parsedItems, transformers);

    expect(debug.stageNames).toEqual(['Parse Result', 'Trans1']);
    expect(debug.stageResult(0).schema).toEqual([{ name: 'id' }, { name: 'y' }]);
    expect(debug.stageResult(1).schema).toEqual([
      { name: 'id' },
      { name: 'y', annotation: ColumnAnnotation.REMOVED },
      { name: 'line', annotation: ColumnAnnotation.ADDED },
    ]);

    expect(debug.stageResult(0).itemsUnpacked()).toEqual(parsedItems);
    expect(debug.stageResult(1).itemsUnpacked()).toEqual(trans1Items);

    const lineMergingStage = debug.stageResult(1);
    const { changes, pages } = lineMergingStage;

    //verify item groups
    expect(pages[0].itemGroups.map((itemGroup) => changes.hasChanged(itemGroup.top))).toEqual([false, true]);

    //verify unpacked items
    expect(lineMergingStage.itemsUnpacked().map((item) => changes.hasChanged(item))).toEqual([false, false, false]);
  });
});

test('Change inside of Line', async () => {
  const parsedSchema = ['id', 'line'];
  const parsedItems = items(0, [
    { id: 1, line: 1 },
    { id: 2, line: 1 },
    { id: 3, line: 2 },
    { id: 4, line: 2 },
  ]);

  const trans1Desc = { requireColumns: ['id', 'line'], debug: { itemMerger: new LineItemMerger() } };
  const trans1Schema = parsedSchema;
  const trans1Items = swapElements([...parsedItems], 0, 1);

  const transformers = [new TestTransformer('Trans1', trans1Desc, trans1Schema, trans1Items)];
  const debug = new Debugger(new Map(), [], 1, parsedSchema, parsedItems, transformers);

  expect(debug.stageNames).toEqual(['Parse Result', 'Trans1']);
  expect(debug.stageResult(0).schema).toEqual([{ name: 'id' }, { name: 'line' }]);
  expect(debug.stageResult(1).schema).toEqual([{ name: 'id' }, { name: 'line' }]);
  expect(debug.stageResult(0).itemsUnpacked()).toEqual(parsedItems);
  expect(debug.stageResult(1).itemsUnpacked()).toEqual(trans1Items);

  const { changes, pages } = debug.stageResult(1);

  //verify item groups
  expect(pages[0].itemGroups.map((itemGroup) => changes.hasChanged(itemGroup.top))).toEqual([true, false]);

  //verify unpacked items
  expect(
    debug
      .stageResult(1)
      .itemsUnpacked()
      .map((item) => changes.hasChanged(item)),
  ).toEqual([true, true, false, false]);
});

var swapElements = function (arr: Item[], indexA: number, indexB: number): Item[] {
  var temp = arr[indexA];
  arr[indexA] = arr[indexB];
  arr[indexB] = temp;
  return arr;
};

describe('build schemas', () => {
  const items: Item[] = [];

  function calculateSchema(inputSchema: string[], outputSchema: string[]): AnnotatedColumn[] {
    const transformers = [new TestTransformer('Trans1', {}, outputSchema, items)];
    const debug = new Debugger(new Map(), [], 1, inputSchema, items, transformers);
    return debug.stageResult(1).schema;
  }

  test('Add', async () => {
    const annotatedSchema = calculateSchema(['A', 'B'], ['A', 'B', 'C']);
    expect(annotatedSchema).toEqual([{ name: 'A' }, { name: 'B' }, { name: 'C', annotation: ColumnAnnotation.ADDED }]);
  });

  test('Remove', async () => {
    const annotatedSchema = calculateSchema(['A', 'B'], ['A']);
    expect(annotatedSchema).toEqual([{ name: 'A' }, { name: 'B', annotation: ColumnAnnotation.REMOVED }]);
  });

  test('Replace first', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['X', 'B', 'C']);
    expect(annotatedSchema).toEqual([
      { name: 'A', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'B' },
      { name: 'C' },
    ]);
  });

  test('Replace middle', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['A', 'X', 'C']);
    expect(annotatedSchema).toEqual([
      { name: 'A' },
      { name: 'B', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'C' },
    ]);
  });

  test('Replace last', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['A', 'B', 'X']);
    expect(annotatedSchema).toEqual([
      { name: 'A' },
      { name: 'B' },
      { name: 'C', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
    ]);
  });

  test('Replace first with 2', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['X', 'Y', 'B', 'C']);
    expect(annotatedSchema).toEqual([
      { name: 'A', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'Y', annotation: ColumnAnnotation.ADDED },
      { name: 'B' },
      { name: 'C' },
    ]);
  });

  test('Replace middle with 2', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['A', 'X', 'Y', 'C']);
    expect(annotatedSchema).toEqual([
      { name: 'A' },
      { name: 'B', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'Y', annotation: ColumnAnnotation.ADDED },
      { name: 'C' },
    ]);
  });

  test('Replace last with 2', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['A', 'B', 'X', 'Y']);
    expect(annotatedSchema).toEqual([
      { name: 'A' },
      { name: 'B' },
      { name: 'C', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'Y', annotation: ColumnAnnotation.ADDED },
    ]);
  });

  test('Replace 2 with one', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C', 'D'], ['A', 'X', 'D']);
    expect(annotatedSchema).toEqual([
      { name: 'A' },
      { name: 'B', annotation: ColumnAnnotation.REMOVED },
      { name: 'C', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'D' },
    ]);
  });

  test('Replace all', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C'], ['X']);
    expect(annotatedSchema).toEqual([
      { name: 'A', annotation: ColumnAnnotation.REMOVED },
      { name: 'B', annotation: ColumnAnnotation.REMOVED },
      { name: 'C', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
    ]);
  });

  test('Wild mix all', async () => {
    const annotatedSchema = calculateSchema(['A', 'B', 'C', 'E', 'F', 'G'], ['B', 'X', 'E', 'Y', 'Z', 'G', 'XX']);
    expect(annotatedSchema).toEqual([
      { name: 'A', annotation: ColumnAnnotation.REMOVED },
      { name: 'B' },
      { name: 'C', annotation: ColumnAnnotation.REMOVED },
      { name: 'X', annotation: ColumnAnnotation.ADDED },
      { name: 'E' },
      { name: 'F', annotation: ColumnAnnotation.REMOVED },
      { name: 'Y', annotation: ColumnAnnotation.ADDED },
      { name: 'Z', annotation: ColumnAnnotation.ADDED },
      { name: 'G' },
      { name: 'XX', annotation: ColumnAnnotation.ADDED },
    ]);
  });
});
