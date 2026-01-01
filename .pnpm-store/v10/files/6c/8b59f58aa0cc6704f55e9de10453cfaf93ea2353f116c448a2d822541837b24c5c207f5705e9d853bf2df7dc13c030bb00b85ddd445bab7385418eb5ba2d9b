import CompactLines from 'src/transformer/CompactLines';
import { emptyContext } from './testContext';
import { items } from '../testItems';

const transformer = new CompactLines();

test('Transform - 2 column pdf (Smart Immunity)', async () => {
  const results = transformer.transform(
    emptyContext(),
    items(0, [
      {
        x: 54,
        y: 91.52,
        str: 'most, get what feels more like a flu. And',
        height: 12,
      },
      {
        x: 324,
        y: 710.52,
        str: 'some, it turns out, aren’t even aware that',
        height: 12,
      },
    ]),
  );
  expect(results.items.map((item) => item.data['line'])).toEqual([0, 1]);
});

test('Transform - raised characters (example.pdf)', async () => {
  const results = transformer.transform(
    emptyContext(),
    items(0, [
      {
        x: 240,
        y: 585,
        str: 'Dies ist eine Test-PDF',
        height: 11,
      },
      {
        x: 352.69,
        y: 585,
        str: '.',
        height: 11,
      },
      {
        x: 348,
        y: 588,
        str: '1',
        height: 7.33,
      },
      { x: 208, y: 572, str: 'Für’s Testen des', height: 11 },
    ]),
  );
  expect(results.items.map((item) => item.data['line'])).toEqual([0, 0, 0, 1]);
});

test('Transform - lowered charactes (dict.pdf)', async () => {
  const results = transformer.transform(
    emptyContext(),
    items(0, [
      { str: 'Let', x: 100.35, y: 625.05, height: 11.96 },
      { str: 'D', x: 122.38, y: 625.05, height: 11.96 },
      { str: '(', x: 100.35, y: 610.61, height: 11.96 },
      { str: 'v', x: 104.9, y: 610.61, height: 11.96 },
      { str: '0', x: 110.57, y: 608.82, height: 7.97 },
      { str: ', a', x: 115.29, y: 610.61, height: 11.96 },
      { str: 'all are different,', x: 100.35, y: 596.16, height: 11.96 },
    ]),
  );
  expect(results.items.map((item) => item.data['line'])).toEqual([0, 0, 1, 1, 1, 1, 2]);
});
