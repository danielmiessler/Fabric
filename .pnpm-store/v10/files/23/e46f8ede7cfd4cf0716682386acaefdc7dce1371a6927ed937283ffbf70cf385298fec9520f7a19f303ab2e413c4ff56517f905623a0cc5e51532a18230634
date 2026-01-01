export function flatMap<T, S>(array: T[], func: (element: T, idx: number) => S[]): S[] {
  return array.reduce((result, entry, idx) => result.concat(func(entry, idx)), [] as S[]);
}

export function groupBy<T, S>(array: T[], groupKey: (element: T) => S): T[][] {
  const groupMap = array.reduce((map: Map<S, T[]>, element: T) => {
    const key = groupKey(element);
    const elementsInGroup = map.get(key);
    if (elementsInGroup) {
      elementsInGroup.push(element);
    } else {
      map.set(key, [element]);
    }
    return map;
  }, new Map());
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  return Array.from(groupMap, ([_, value]) => value);
}

export function flatten<T>(array: T[][]): T[] {
  return flatMap(array, (e) => e);
}

export function arraysEqual<T>(a: T[], b: T[]): boolean {
  if (a === b) return true;
  if (a == null || b == null) return false;
  if (a.length !== b.length) return false;

  for (let i = 0; i < a.length; ++i) {
    if (a[i] !== b[i]) return false;
  }
  return true;
}
