import Item from '../Item';

export function flatMap<T, S>(array: T[], func: (entry: T, idx: number) => S[]): S[] {
  return array.reduce((result, entry, idx) => result.concat(func(entry, idx)), [] as S[]);
}

export function onlyUniques<T>(value: T, index: number, self: T[]) {
  return self.indexOf(value) === index;
}

export function ascending(a: number, b: number): number {
  return a - b;
}

export function descending(a: number, b: number): number {
  return b - a;
}

export function min(array: number[], prevMin?: number): number {
  return array.reduce((prev, curr) => {
    if (prev === undefined || curr < prev) {
      return curr;
    }
    return prev;
  }, prevMin);
}

export function max(array: number[], prevMin?: number): number {
  return array.reduce((prev, curr) => {
    if (prev === undefined || curr > prev) {
      return curr;
    }
    return prev;
  }, prevMin);
}

export function count<T>(array: T[], find: (entry: T) => boolean): number {
  return array.reduce((count, entry) => (find(entry) ? count + 1 : count), 0);
}

export function median(values: number[]) {
  if (values.length === 0) return 0;

  values.sort(function (a, b) {
    return a - b;
  });

  const half = Math.floor(values.length / 2);

  if (values.length % 2) return values[half];

  return (values[half - 1] + values[half]) / 2.0;
}

type KeyExtractor<O> = (item: Item) => O;
type PageItemTransformer = (page: number, items: Item[]) => Item[];
type LineItemTransformer = (page: number, line: number, items: Item[]) => Item[];

function groupBy<O>(items: Item[], extractKey: KeyExtractor<O>): Item[][] {
  return items.reduce((pageItems: Item[][], item: Item) => {
    const lastPageItems = pageItems[pageItems.length - 1];
    if (!lastPageItems || extractKey(item) !== extractKey(lastPageItems[0])) {
      pageItems.push([item]);
    } else {
      lastPageItems.push(item);
    }
    return pageItems;
  }, []);
}

export function groupByPage(items: Item[]): Item[][] {
  return groupBy(items, (item) => item.page);
}

export function groupByBlock(items: Item[]): Item[][] {
  return groupByElement(items, 'block');
}

export function groupByLine(items: Item[]): Item[][] {
  return groupByElement(items, 'line');
}

export function groupByElement(items: Item[], elementName: string): Item[][] {
  return groupBy(items, (item) => item.data[elementName]);
}

export function transformGroupedByPage(items: Item[], groupedTransformer: PageItemTransformer): Item[] {
  return new Array<Item>().concat(
    ...groupByPage(items).map((pageItems) => groupedTransformer(pageItems[0].page, pageItems)),
  );
}

export function transformGroupedByPageAndLine(items: Item[], groupedTransformer: LineItemTransformer): Item[] {
  const transformedItems: Item[] = [];
  groupByPage(items).forEach((pageItems) => {
    groupByElement(pageItems, 'line').forEach((lineItems) => {
      transformedItems.push(...groupedTransformer(pageItems[0].page, lineItems[0].data['line'], lineItems));
    });
  });
  return transformedItems;
}

export function mostFrequent<T extends number | string>(items: Item[], dataElementKey: string): T | undefined {
  const occurenceMap = items.reduce((map: Map<T, number>, item) => {
    const key = item.data[dataElementKey];
    const occurrence = map.get(key) || 0;
    map.set(key, occurrence + 1);
    return map;
  }, new Map());

  const topElement = [...occurenceMap].reduce(
    (topEntry: [T | undefined, number], entry: [T, number]) => (entry[1] >= topEntry[1] ? entry : topEntry),
    [undefined, 0],
  )[0];

  //TODO optimally we should handle the 50/50 case
  return topElement;
}

export function majorityElement<T>(items: Item[], extract: (item: Item) => T): T | undefined {
  if (items.length == 0) {
    return;
  }
  let maj = 0,
    count = 1;

  for (let i = 1; i < items.length; i++) {
    if (extract(items[i]) === extract(items[maj])) {
      count++;
    } else {
      count--;
    }

    if (count === 0) {
      maj = i;
      count = 1;
    }
  }
  return extract(items[maj]);
}

export function isGreaterWithTolerance(num1: number, num2: number, tolerance = 0.01) {
  return num1 - num2 > tolerance;
}
