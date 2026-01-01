import GlobalDefinition from './GlobalDefinition';
import { assertDefined, assertNot } from './assert';
import GlobalValue from './GlobalValue';

export default class Globals {
  map: Map<string, unknown>;
  constructor(globals?: Globals) {
    this.map = globals ? new Map(globals.map) : new Map();
  }

  keys(): string[] {
    return [...this.map.keys()];
  }

  isDefined<T>(definition: GlobalDefinition<T>): boolean {
    return typeof this.map.get(definition.key) !== 'undefined';
  }

  get<T>(definition: GlobalDefinition<T>): T {
    const element = this.map.get(definition.key) as T;
    assertDefined(
      element,
      `No global with key '${definition.key}' registered. Only [${[...this.map.keys()].join(',')}]`,
    );
    return element;
  }

  getOptional<T>(definition: GlobalDefinition<T>): T | undefined {
    return this.map.get(definition.key) as T;
  }

  set<T>(definition: GlobalDefinition<T>, value: T) {
    assertNot(this.isDefined(definition), `Global with key '${definition.key}' already registered.`);
    this.map.set(definition.key, value);
  }

  override<T>(definition: GlobalDefinition<T>, value: T) {
    this.map.set(definition.key, value);
  }

  withValues(values: GlobalValue<unknown>[] | undefined): Globals {
    values?.forEach((value) => {
      if (value.override) {
        this.override(value.definition, value.value);
      } else {
        this.set(value.definition, value.value);
      }
    });
    return this;
  }
}
