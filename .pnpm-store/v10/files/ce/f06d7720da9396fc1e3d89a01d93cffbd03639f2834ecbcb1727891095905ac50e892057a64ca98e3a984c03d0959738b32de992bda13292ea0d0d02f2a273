import Item from 'src/Item';
import {
  groupByPage,
  groupByLine,
  groupByElement,
  transformGroupedByPage,
  transformGroupedByPageAndLine,
  mostFrequent,
  flatMap,
  onlyUniques,
  count,
  min,
  max,
} from 'src/support/groupingUtils';
import { items } from 'test/testItems';

test('count', async () => {
  expect(count([], () => true)).toEqual(0);
  expect(count([1, 2, 3, 4, 5, 6], (e) => e % 2 == 0)).toEqual(3);
  expect(count(['A', 'B', 'c'], (e) => e === e.toUpperCase())).toEqual(2);
});

test('flatMap', async () => {
  expect(flatMap([], (e) => e)).toEqual([]);
  expect(flatMap([[1, 2], [3], [4, 5, 6]], (e) => e)).toEqual([1, 2, 3, 4, 5, 6]);
  expect(flatMap([{ x: [1, 2] }, { x: [3] }, { x: [4, 5, 6] }], (e) => e.x)).toEqual([1, 2, 3, 4, 5, 6]);
});

test('onlyUniques', async () => {
  expect([].filter(onlyUniques)).toEqual([]);
  expect([1, 2, 3].filter(onlyUniques)).toEqual([1, 2, 3]);
  expect([1, 2, 3, 3, 2, 1].filter(onlyUniques)).toEqual([1, 2, 3]);
});

test('min', async () => {
  expect(min([])).toBeUndefined();
  expect(min([1])).toEqual(1);
  expect(min([1, 2])).toEqual(1);
  expect(min([2, 1])).toEqual(1);
  expect(min([1, 4, -3, 5])).toEqual(-3);

  expect(min([2, 3], 1)).toEqual(1);
  expect(min([2, 3], 4)).toEqual(2);
});

test('max', async () => {
  expect(max([])).toBeUndefined();
  expect(max([1])).toEqual(1);
  expect(max([1, 2])).toEqual(2);
  expect(max([2, 1])).toEqual(2);
  expect(max([1, 4, -3, 5])).toEqual(5);

  expect(max([2, 3], 1)).toEqual(3);
  expect(max([2, 3], 4)).toEqual(4);
});

describe('groupByPage', () => {
  test('empty', async () => {
    expect(groupByPage([])).toEqual([]);
  });

  test('group', async () => {
    const pageItems = [
      [new Item(0, { id: 1 })],
      [new Item(1, { id: 2 }), new Item(1, { id: 3 })],
      [new Item(2, { id: 4 })],
    ];
    const flattenedItems = new Array<Item>().concat(...pageItems);
    const transformedItems = groupByPage(flattenedItems);
    expect(transformedItems).toEqual(pageItems);
  });
});

describe('groupByLine', () => {
  test('empty', async () => {
    expect(groupByLine([])).toEqual([]);
  });

  test('group', async () => {
    const groupedItems = [
      [new Item(0, { line: 1, id: 1 })],
      [new Item(0, { line: 2, id: 2 }), new Item(0, { line: 2, id: 3 })],
      [new Item(0, { line: 3, id: 4 })],
    ];
    const flattenedItems = new Array<Item>().concat(...groupedItems);
    const transformedItems = groupByLine(flattenedItems);
    expect(transformedItems).toEqual(groupedItems);
  });
});

describe('groupByElement', () => {
  test('empty', async () => {
    expect(groupByElement([], 'line')).toEqual([]);
  });

  test('group', async () => {
    const groupedItems = [
      [new Item(0, { line: 1, id: 1 })],
      [new Item(0, { line: 2, id: 2 }), new Item(0, { line: 2, id: 3 })],
      [new Item(0, { line: 3, id: 4 })],
    ];
    const flattenedItems = new Array<Item>().concat(...groupedItems);
    const transformedItems = groupByElement(flattenedItems, 'line');
    expect(transformedItems).toEqual(groupedItems);
  });
});

describe('transformGroupedByPage', () => {
  test('empty', async () => {
    const transformedItems = transformGroupedByPage([], () => fail("shoudln't be called"));
    expect(transformedItems).toEqual([]);
  });

  test('pipe through', async () => {
    const pageItems = [
      [new Item(0, { id: 1 })],
      [new Item(1, { id: 2 }), new Item(1, { id: 3 })],
      [new Item(2, { id: 4 })],
    ];
    const flattenedItems = new Array<Item>().concat(...pageItems);
    const transformedItems = transformGroupedByPage(flattenedItems, (page, items) => {
      expect(items).toEqual(pageItems[page]);
      return items;
    });
    expect(transformedItems).toEqual(flattenedItems);
  });

  test('change', async () => {
    const input = [new Item(0, { v: 0 }), new Item(1, { v: 0 })];
    const transformedItems = transformGroupedByPage(input, (_, items) => {
      return [items[0].withData({ v: 1 })];
    });
    expect(transformedItems).toEqual(input.map((item) => item.withData({ v: 1 })));
  });
});

describe('transformGroupedByPageAndLine', () => {
  test('empty', async () => {
    const transformedItems = transformGroupedByPageAndLine([], () => fail("shoudln't be called"));
    expect(transformedItems).toEqual([]);
  });

  test('transform', async () => {
    const pageItems = [
      items(0, [{ line: 1, id: 1 }]),
      items(1, [
        { line: 1, id: 2 },
        { line: 1, id: 3 },
        { line: 2, id: 4 },
      ]),
      items(2, [{ line: 1, id: 5 }]),
    ];
    const flattenedItems = new Array<Item>().concat(...pageItems);
    const transformedItems = transformGroupedByPageAndLine(flattenedItems, (page, line, items) => {
      return [new Item(0, { group: `${page}/${line}:${items.length}` })];
    });
    expect(transformedItems.map((item) => item.data['group'])).toEqual(['0/1:1', '1/1:2', '1/2:1', '2/1:1']);
  });
});

describe('mostFrequent', () => {
  test('empty', async () => {
    const majorityItem = mostFrequent([], 'x');
    expect(majorityItem).toBeUndefined();
  });

  test('clear winner1', async () => {
    const allItems = items(0, [{ x: 1 }, { x: 1 }, { x: 2 }, { x: 3 }]);
    const majorityItem = mostFrequent(allItems, 'x');
    expect(majorityItem).toEqual(1);
  });

  test('clear winner2', async () => {
    const allItems = items(0, [{ x: 1 }, { x: 2 }, { x: 3 }, { x: 2 }]);
    const majorityItem = mostFrequent(allItems, 'x');
    expect(majorityItem).toEqual(2);
  });

  test('clear winner3 string', async () => {
    const allItems = items(0, [{ x: 'A' }, { x: 'B' }, { x: 'C' }, { x: 'C' }]);
    const majorityItem = mostFrequent(allItems, 'x');
    expect(majorityItem).toEqual('C');
  });

  test('clear 50/50', async () => {
    const allItems = items(0, [{ x: 1 }, { x: 2 }, { x: 1 }, { x: 2 }]);
    const majorityItem = mostFrequent(allItems, 'x');
    expect(majorityItem).toEqual(2);
  });
});
