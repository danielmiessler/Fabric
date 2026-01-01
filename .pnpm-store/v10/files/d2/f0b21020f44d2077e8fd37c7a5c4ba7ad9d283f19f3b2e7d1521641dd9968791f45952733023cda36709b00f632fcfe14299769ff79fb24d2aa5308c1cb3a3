import { flatMap, flatten, groupBy } from 'src/support/functional';

test('flatMap', async () => {
  expect(flatMap([], (e) => e)).toEqual([]);
  expect(flatMap([[1, 2], [3], [4, 5, 6]], (e) => e)).toEqual([1, 2, 3, 4, 5, 6]);
  expect(flatMap([{ x: [1, 2] }, { x: [3] }, { x: [4, 5, 6] }], (e) => e.x)).toEqual([1, 2, 3, 4, 5, 6]);
});

test('flatten', async () => {
  expect(flatten([])).toEqual([]);
  expect(flatten([[1, 2], [3], [4, 5, 6]])).toEqual([1, 2, 3, 4, 5, 6]);
});

test('groupBy', async () => {
  expect(groupBy([], (e) => e)).toEqual([]);
  expect(groupBy([1, 2, 1, 3, 2, 4, 4], (e) => e)).toEqual([[1, 1], [2, 2], [3], [4, 4]]);
  expect(
    groupBy(
      [
        { k: 'a', v: 1 },
        { k: 'a', v: 2 },
        { k: 'b', v: 3 },
      ],
      (e) => e.k,
    ),
  ).toEqual([
    [
      { k: 'a', v: 1 },
      { k: 'a', v: 2 },
    ],
    [{ k: 'b', v: 3 }],
  ]);
});
