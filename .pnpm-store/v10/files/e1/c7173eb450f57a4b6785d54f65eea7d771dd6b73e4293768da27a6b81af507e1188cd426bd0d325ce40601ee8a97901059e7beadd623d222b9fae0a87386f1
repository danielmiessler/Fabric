import type Item from '..//Item';

export default class ItemGroup {
  top: Item;
  elements: Item[];

  constructor(top: Item, items: Item[] = []) {
    this.top = top;
    this.elements = items;
  }

  hasMany(): boolean {
    return this.elements.length > 0;
  }

  unpacked(): Item[] {
    if (this.elements.length > 0) {
      return this.elements;
    }
    return [this.top];
  }
}
