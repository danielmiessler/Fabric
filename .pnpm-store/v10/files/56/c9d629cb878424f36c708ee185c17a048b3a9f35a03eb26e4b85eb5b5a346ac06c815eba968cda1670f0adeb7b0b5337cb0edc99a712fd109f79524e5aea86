import {
  filterOutDigits,
  filterOutWhitespaces,
  extractNumbers,
  extractEndingNumber,
  isListItem,
  isNumberedListItem,
} from 'src/support/stringFunctions';

test('filterOutDigits', async () => {
  expect(filterOutDigits('')).toEqual('');
  expect(filterOutDigits('a b c')).toEqual('a b c');
  expect(filterOutDigits('a1b 2c 3')).toEqual('ab c ');
});

test('filterOutWhitespaces', async () => {
  expect(filterOutWhitespaces('')).toEqual('');
  expect(filterOutWhitespaces('a b c')).toEqual('abc');
  expect(filterOutWhitespaces('ab  c ')).toEqual('abc');
});

test('extractNumbers', async () => {
  expect(extractNumbers('')).toEqual([]);
  expect(extractNumbers('a b c')).toEqual([]);
  expect(extractNumbers('a1b 2c 3')).toEqual([1, 2, 3]);
  expect(extractNumbers('a12 21 304')).toEqual([12, 21, 304]);
});

test('extractEndingNumbers', async () => {
  expect(extractEndingNumber('')).toBeUndefined();
  expect(extractEndingNumber('a b c')).toBeUndefined();
  expect(extractEndingNumber('a1b 2c 3')).toEqual(3);
  expect(extractEndingNumber('a12 21 304')).toEqual(304);
  expect(extractEndingNumber('abc ... 304')).toEqual(304);
});

describe('functions: isListItem', () => {
  it('Match', () => {
    expect(isListItem('- my text')).toEqual(true);
    expect(isListItem('- my text -')).toEqual(true);
    expect(isListItem(' - my text')).toEqual(true);
    expect(isListItem('  - my text')).toEqual(true);
    expect(isListItem(' -  my text')).toEqual(true);

    expect(isListItem('• my text')).toEqual(true);
    expect(isListItem(' • my text')).toEqual(true);
    expect(isListItem('  • my text')).toEqual(true);

    expect(isListItem('– my text')).toEqual(true);
    expect(isListItem(' – my text')).toEqual(true);
  });

  it('No Match', () => {
    expect(isListItem('my text')).toEqual(false);
    expect(isListItem('-my text')).toEqual(false);
    expect(isListItem('•my text')).toEqual(false);
    expect(isListItem(' -my text')).toEqual(false);
  });
});

describe('functions: isNumberedListItem', () => {
  it('Match', () => {
    expect(isNumberedListItem('1.')).toEqual(true);
    expect(isNumberedListItem('1. my text')).toEqual(true);
    expect(isNumberedListItem('2. my text')).toEqual(true);
    expect(isNumberedListItem('23. my text')).toEqual(true);
    expect(isNumberedListItem('23.   my text')).toEqual(true);
    expect(isNumberedListItem(' 23.   my text')).toEqual(true);
    expect(isNumberedListItem('  23.   my text')).toEqual(true);
  });

  it('No Match', () => {
    expect(isNumberedListItem('1two')).toEqual(false);
    expect(isNumberedListItem('1 two')).toEqual(false);
    expect(isNumberedListItem('1.two')).toEqual(false);
  });
});
