import PageMapping from 'src/PageMapping';

test('1-to-1', async () => {
  const mapping = new PageMapping(1, false);
  expect(mapping.pageFactor).toEqual(1);
  expect([...Array(3).keys()].map((i) => mapping.pageLabel(i))).toEqual(['1', '2', '3']);
});

test('one starter page', async () => {
  const mapping = new PageMapping(0, true);
  expect([...Array(5).keys()].map((i) => mapping.pageLabel(i))).toEqual(['I', '1', '2', '3', '4']);
});

test('three starter pages', async () => {
  const mapping = new PageMapping(-2, true);
  expect([...Array(5).keys()].map((i) => mapping.pageLabel(i))).toEqual(['I', 'II', 'III', '1', '2']);
});
