import { ascending } from './groupingUtils';
import { medianSorted } from 'simple-statistics';

export default class Range {
  constructor(public min: number, public median: number, public max: number) {}

  description() {
    return `${this.median}(${this.min}/${this.max})`;
  }
}

export function fromMiddle(middle: number, maxDeviation: number) {
  return new Range(middle - maxDeviation, middle, middle + maxDeviation);
}

export function deriveFromUnsorted(numbers: number[]) {
  numbers.sort(ascending);
  const median = medianSorted(numbers);
  const min = numbers[0];
  const max = numbers[numbers.length - 1];
  return new Range(min, median, max);
}
