export function assert(condition: boolean, message: string) {
  if (!condition) {
    throw new Error(message || 'Assertion failed');
  }
}

export function assertNot(condition: boolean, message: string) {
  if (condition) {
    throw new Error(message || 'Assertion failed');
  }
}

export function assertDefined<T>(value: T | undefined, message: string): T {
  if (value === null || typeof value === 'undefined') {
    throw new Error(message || 'Assertion failed');
  }
  return value;
}
