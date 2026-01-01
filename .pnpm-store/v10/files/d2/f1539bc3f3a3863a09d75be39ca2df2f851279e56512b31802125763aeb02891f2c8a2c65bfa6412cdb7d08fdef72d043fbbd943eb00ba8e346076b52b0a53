export function numbersAreAsencding(numbers: number[]): boolean {
  return numbers.every((num, idx) => (idx > 0 ? num >= numbers[idx - 1] : true));
}

export function numbersAreConsecutive(numbers: number[]): boolean {
  return numbers.every((num, idx) => (idx > 0 ? num === numbers[idx - 1] + 1 : true));
}
