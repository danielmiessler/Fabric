import { assert } from '../assert';

export const TAB_CHAR_CODE = 9;
export const WHITESPACE_CHAR_CODE = 32;
export const MIN_DIGIT_CHAR_CODE = 48;
export const MAX_DIGIT_CHAR_CODE = 57;
export const PERIOD_CHAR_CODES = [46, 190];
export const DASHS_CHAR_CODES = [45, 189, 8211];

export function isDigit(charCode: number): boolean {
  return charCode >= MIN_DIGIT_CHAR_CODE && charCode <= MAX_DIGIT_CHAR_CODE;
}

export function toCharcodes(text: string): number[] {
  const codes: number[] = [];
  for (let index = 0; index < text.length; index++) {
    codes.push(text.charCodeAt(index));
  }
  return codes;
}

export function filterOutDigits(text: string): string {
  return String.fromCharCode(...toCharcodes(text).filter((code) => !isDigit(code)));
}

export function filterOutWhitespaces(text: string): string {
  return filterOut(text, [TAB_CHAR_CODE, WHITESPACE_CHAR_CODE]);
}

export function filterOut(text: string, codes: number[]): string {
  return String.fromCharCode(...toCharcodes(text).filter((code) => !codes.includes(code)));
}

export function extractNumbers(text: string): number[] {
  return (text.match(/\d+/g) || []).map(Number);
}

export function extractEndingNumber(text: string): number | undefined {
  const match = text.match(/\d+$/g);
  if (match) {
    assert(match.length == 1, `Expected only one match, but got ${match}`);
    return Number(match[0]);
  }
  return undefined;
}

export function isListItemCharacter(string: string) {
  if (string.length > 1) {
    return false;
  }
  const char = string.charAt(0);
  return char === '-' || char === '•' || char === '–';
}

export function isListItem(value: string) {
  return /^[\s]*[-•–][\s].*$/g.test(value);
}

export function isNumberedListItem(value: string) {
  return /^[\s]*\d*\.(?:\s|$)/g.test(value);
}

export function isNumber(value: string) {
  for (let i = 0; i < value.length; i++) {
    const charCode = value.charCodeAt(i);
    if (!isDigit(charCode)) {
      return false;
    }
  }
  return true;
}
