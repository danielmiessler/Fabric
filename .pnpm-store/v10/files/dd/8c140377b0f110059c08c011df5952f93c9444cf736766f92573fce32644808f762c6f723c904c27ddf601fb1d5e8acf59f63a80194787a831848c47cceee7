import { TextType, isHeadline, headlineLevel, toHeadlineType } from 'src/text-types';

describe('headline', () => {
  test('isHeadline', async () => {
    expect(isHeadline('H0' as TextType)).toBe(false);
    expect(isHeadline('H1')).toBe(true);
    expect(isHeadline('H2')).toBe(true);
    expect(isHeadline('H3')).toBe(true);
    expect(isHeadline('H4')).toBe(true);
    expect(isHeadline('H5')).toBe(true);
    expect(isHeadline('H6')).toBe(true);
    expect(isHeadline('H7' as TextType)).toBe(false);
  });

  test('headlineLevel', async () => {
    expect(headlineLevel('H1')).toBe(1);
    expect(headlineLevel('H2')).toBe(2);
    expect(headlineLevel('H3')).toBe(3);
    expect(headlineLevel('H4')).toBe(4);
    expect(headlineLevel('H5')).toBe(5);
    expect(headlineLevel('H6')).toBe(6);
    expect(() => headlineLevel('H0' as TextType)).toThrow(
      Error("Expected headline level between 1 and 6 but was 0 (from 'H0')"),
    );
    expect(() => headlineLevel('H7' as TextType)).toThrow(
      Error("Expected headline level between 1 and 6 but was 7 (from 'H7')"),
    );
  });

  test('toHeadlineType', async () => {
    expect(toHeadlineType(1)).toEqual('H1');
    expect(toHeadlineType(2)).toEqual('H2');
    expect(toHeadlineType(3)).toEqual('H3');
    expect(toHeadlineType(4)).toEqual('H4');
    expect(toHeadlineType(5)).toEqual('H5');
    expect(toHeadlineType(6)).toEqual('H6');
    expect(() => toHeadlineType(1 - 1)).toThrow(Error('Expected headline level between 1 and 6 but was 0'));
    expect(() => toHeadlineType(6 + 1)).toThrow(Error('Expected headline level between 1 and 6 but was 7'));
  });
});
