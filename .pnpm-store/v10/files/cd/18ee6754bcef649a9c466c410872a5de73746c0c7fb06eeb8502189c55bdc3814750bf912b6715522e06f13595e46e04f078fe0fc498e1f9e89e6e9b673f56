import PageFactorFinder from 'src/support/PageFactorFinder';

interface Container {
  index: number;
  numbers: number[];
}

const extractor = (container: Container) => container;

test('distraction free - straight', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [1] },
    { index: 1, numbers: [2] },
    { index: 2, numbers: [3] },
    { index: 3, numbers: [4] },
    { index: 4, numbers: [5] },
    { index: 5, numbers: [6] },
  ];
  expect(finder.find(containers, extractor)).toEqual(1);
});

test('distraction free - accept gap in numbers', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [1] },
    { index: 1, numbers: [2] },
    { index: 2, numbers: [3] },
    { index: 3, numbers: [4] },
    { index: 4, numbers: [] },
    { index: 5, numbers: [6] },
  ];

  expect(finder.find(containers, extractor)).toEqual(1);
});

test('distraction free - accept gap in pages', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [1] },
    { index: 1, numbers: [2] },
    { index: 2, numbers: [3] },
    { index: 3, numbers: [4] },
    { index: 5, numbers: [6] },
  ];

  expect(finder.find(containers, extractor)).toEqual(1);
});

test('distraction free - defered', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [2006] },
    { index: 1, numbers: [] },
    { index: 2, numbers: [1, 1] },
    { index: 3, numbers: [2] },
    { index: 4, numbers: [3] },
    { index: 5, numbers: [4] },
    { index: 6, numbers: [5] },
    { index: 7, numbers: [6] },
    { index: 8, numbers: [7] },
    { index: 9, numbers: [8] },
  ];

  expect(finder.find(containers, extractor)).toEqual(-1);
});

test('distraction loaden - straight', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [1, -3453] },
    { index: 1, numbers: [2, 355] },
    { index: 2, numbers: [3, 950, 4] },
    { index: 3, numbers: [4, 534, 5] },
    { index: 4, numbers: [5, 6] },
    { index: 5, numbers: [6, 35335] },
  ];

  expect(finder.find(containers, extractor)).toEqual(1);
});

test('distraction loaden - defered', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [2006] },
    { index: 1, numbers: [5] },
    { index: 2, numbers: [1, 7678] },
    { index: 3, numbers: [2, 2] },
    { index: 4, numbers: [3, 4] },
    { index: 5, numbers: [4, 5, 65, 8] },
    { index: 6, numbers: [5, 9] },
    { index: 7, numbers: [6] },
    { index: 8, numbers: [27, 7, 19] },
    { index: 9, numbers: [-4, 2016, 8] },
  ];

  expect(finder.find(containers, extractor)).toEqual(-1);
});

test('many numbers but no meaningful match', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [3] },
    { index: 1, numbers: [7] },
    { index: 2, numbers: [4] },
    { index: 3, numbers: [6, 5] },
    { index: 4, numbers: [13, 9] },
    { index: 5, numbers: [8, 7] },
    { index: 6, numbers: [11] },
    { index: 7, numbers: [1] },
  ];

  expect(finder.find(containers, extractor)).toBeUndefined();
});

test('many numbers but no match', () => {
  const finder = new PageFactorFinder();
  const containers: Container[] = [
    { index: 0, numbers: [22] },
    { index: 1, numbers: [7] },
    { index: 2, numbers: [14] },
    { index: 3, numbers: [1, 5] },
    { index: 4, numbers: [13, 9] },
    { index: 5, numbers: [8, 787] },
    { index: 6, numbers: [12] },
    { index: 7, numbers: [1] },
  ];

  expect(finder.find(containers, extractor)).toBeUndefined();
});
