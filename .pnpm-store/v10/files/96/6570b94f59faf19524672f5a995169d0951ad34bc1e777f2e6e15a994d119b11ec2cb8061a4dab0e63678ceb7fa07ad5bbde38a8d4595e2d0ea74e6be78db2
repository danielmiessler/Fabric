import { assert } from './assert';

export type HeadlineType = 'H1' | 'H2' | 'H3' | 'H4' | 'H5' | 'H6';

export type ListType = 'LIST' | 'NUMBERED_LIST';

export type TextType = HeadlineType | 'TOC' | 'FOOTNOTES' | 'CODE' | ListType | 'PARAGRAPH';

function types(...types: TextType[]) {
  return types;
}

export function toBlockType(type: TextType): TextType {
  if (type === 'NUMBERED_LIST') {
    return 'LIST';
  }
  return type;
}

export function isHeadline(type: TextType) {
  return types('H1', 'H2', 'H3', 'H4', 'H5', 'H6').includes(type);
}

export function toHeadlineType(number: 1 | 2 | 3 | 4 | 5 | 6 | number): HeadlineType {
  assert(number > 0 && number < 7, `Expected headline level between 1 and 6 but was ${number}`);
  return `H${number.toString()}` as HeadlineType;
}

export function headlineLevel(type: TextType) {
  const level = parseInt(type[1]);
  assert(level > 0 && level < 7, `Expected headline level between 1 and 6 but was ${level} (from '${type}')`);
  return level;
}

export function mergeToBlock(type: TextType) {
  return types('FOOTNOTES', 'CODE', 'LIST', 'NUMBERED_LIST').includes(type);
}

export function mergeFollowingNonTypedItems(type: TextType) {
  return types('FOOTNOTES').includes(type);
}

export function mergeFollowingNonTypedItemsWithSmallDistance(type: TextType) {
  return types('LIST', 'NUMBERED_LIST').includes(type);
}

// Discard token types like bold for certain text types
export function discardTokenTypes(blockTypes: TextType[]) {
  if (blockTypes.includes('CODE')) {
    return true;
  }
  return false;
}
