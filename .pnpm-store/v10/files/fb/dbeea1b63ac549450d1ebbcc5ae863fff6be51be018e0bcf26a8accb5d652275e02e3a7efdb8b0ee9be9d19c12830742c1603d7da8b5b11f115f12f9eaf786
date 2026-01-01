import { TextType } from 'src/text-types';
import { assertDefined } from '../assert';
import Item from '../Item';
import { onlyUniques } from './groupingUtils';

function get(item: Item, name: string): any {
  const value = item.data[name];
  assertDefined(value, `No '${name}' defined in ${JSON.stringify(item)}`);
  return value;
}
export function getHeight(item: Item): number {
  return get(item, 'height');
}

export function getText(item: Item): string {
  return get(item, 'str');
}

export function joinText(items: Item[], joinCharacter: string): string {
  return items.map((item) => getText(item)).join(joinCharacter);
}

export function getFontName(fontMap: Map<string, object>, item: Item): string {
  const fontId = item.data['fontName'];
  const fontObject = fontMap.get(fontId);
  if (!fontObject) {
    return fontId;
  }
  return assertDefined(fontObject['name'], `No 'name' found in ${JSON.stringify(fontObject)}`);
}

export function itemWithType(item: Item, type: TextType): Item {
  const existingTypes = item.data['types'] || [];
  return item.withDataAddition({ types: [...existingTypes, type].filter(onlyUniques) });
}
