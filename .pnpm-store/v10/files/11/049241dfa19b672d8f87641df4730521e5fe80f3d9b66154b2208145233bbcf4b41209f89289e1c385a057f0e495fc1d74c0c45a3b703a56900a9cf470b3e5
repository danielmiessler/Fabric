import GlobalDefinition from 'src/GlobalDefinition';
import Globals from 'src/Globals';

const MyGlobalString = new GlobalDefinition<string>('myGlobalString');
const MyGlobalNumber = new GlobalDefinition<number>('myGlobalNumber');

test('not set', async () => {
  const globals = new Globals();
  globals.set(MyGlobalString, '23');
  expect(globals.isDefined(MyGlobalNumber)).toBeFalsy();
  expect(globals.getOptional(MyGlobalNumber)).toBeUndefined();
  expect(() => globals.get(MyGlobalNumber)).toThrow(
    `No global with key '${MyGlobalNumber.key}' registered. Only [${MyGlobalString.key}]`,
  );
});

test('set', async () => {
  const globals = new Globals();
  globals.set(MyGlobalNumber, 24);

  expect(globals.isDefined(MyGlobalNumber)).toBeTruthy();
  expect(globals.get(MyGlobalNumber)).toEqual(24);
  expect(globals.getOptional(MyGlobalNumber)).toEqual(24);
  expect(globals.keys()).toEqual([MyGlobalNumber.key]);
});

test('set, already exists', async () => {
  const globals = new Globals();
  globals.set(MyGlobalNumber, 24);
  expect(() => globals.set(MyGlobalNumber, 25)).toThrow("Global with key 'myGlobalNumber' already registered.");
});

test('override', async () => {
  const globals = new Globals();
  globals.set(MyGlobalNumber, 24);
  globals.override(MyGlobalNumber, 25);

  expect(globals.isDefined(MyGlobalNumber)).toBeTruthy();
  expect(globals.get(MyGlobalNumber)).toEqual(25);
});

test('inheritence', async () => {
  const globals1 = new Globals();
  globals1.set(MyGlobalNumber, 24);
  const globals2 = new Globals(globals1);
  globals2.set(MyGlobalString, 'myKey');

  expect(globals2.keys()).toEqual([MyGlobalNumber.key, MyGlobalString.key]);
  expect(globals2.isDefined(MyGlobalNumber)).toBeTruthy();
  expect(globals2.isDefined(MyGlobalString)).toBeTruthy();
  expect(globals2.get(MyGlobalNumber)).toEqual(24);
  expect(globals2.get(MyGlobalString)).toEqual('myKey');
});
