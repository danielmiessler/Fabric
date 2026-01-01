import Item from 'src/Item';

export function items(page: number, data: object[]): Item[] {
  return data.map((data) => new Item(page, data));
}

export function item(page: number, data: object): Item {
  return new Item(page, data);
}

export function idItems(page: number, ids: string[]): Item[] {
  return ids.map((id) => new Item(page, { id }));
}

export function idItem(page: number, id: string): Item {
  return new Item(page, { id });
}

export function realisticItems(page: number, data: object[]): Item[] {
  return data.map(
    (data, idx) =>
      new Item(page, {
        line: idx,
        x: idx,
        y: idx,
        str: `text ${idx}`,
        fontName: 'g_d0_f2',
        dir: 'ltr',
        width: 10,
        height: 10,
        ...data,
      }),
  );
}
