import GlobalValue from './GlobalValue';

export default class GlobalDefinition<T> {
  constructor(public key: string) {}

  value(value: T) {
    return new GlobalValue(this, value);
  }

  overrideValue(value: T) {
    return new GlobalValue(this, value, true);
  }
}
