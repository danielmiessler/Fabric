import { numbersAreAsencding, numbersAreConsecutive } from 'src/support/numberFunctions';

test('numbersAreAsencding', async () => {
  expect(numbersAreAsencding([])).toBeTruthy();
  expect(numbersAreAsencding([0])).toBeTruthy();
  expect(numbersAreAsencding([0, 0])).toBeTruthy();
  expect(numbersAreAsencding([0, 1])).toBeTruthy();
  expect(numbersAreAsencding([0, 1, 1, 2, 34])).toBeTruthy();
  expect(numbersAreAsencding([412, 545, 53636])).toBeTruthy();

  expect(numbersAreAsencding([1, 0])).toBeFalsy();
  expect(numbersAreAsencding([0, 1, 0])).toBeFalsy();
});

test('numbersAreConsecutive', async () => {
  expect(numbersAreConsecutive([])).toBeTruthy();
  expect(numbersAreConsecutive([0])).toBeTruthy();
  expect(numbersAreConsecutive([0, 1])).toBeTruthy();
  expect(numbersAreConsecutive([0, 1, 2])).toBeTruthy();
  expect(numbersAreConsecutive([1, 2, 3])).toBeTruthy();
  expect(numbersAreConsecutive([401, 402, 403])).toBeTruthy();

  expect(numbersAreConsecutive([0, 0])).toBeFalsy();
  expect(numbersAreConsecutive([1, 0])).toBeFalsy();
  expect(numbersAreConsecutive([0, 1, 2, 4])).toBeFalsy();
});
