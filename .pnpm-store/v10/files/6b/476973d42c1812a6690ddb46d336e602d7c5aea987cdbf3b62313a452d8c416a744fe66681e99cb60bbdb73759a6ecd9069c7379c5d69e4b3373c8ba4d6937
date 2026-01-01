import ChangeTracker from 'src/debug/ChangeTracker';
import { detectChanges } from 'src/debug/detectChanges';
import { PositionChange, Direction, Addition, Removal } from 'src/debug/ChangeIndex';
import { idItem, idItems } from 'test/testItems';

test('No changes', async () => {
  const items = idItems(0, ['A', 'B', 'C']);

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, items, items);
  expect(tracker.changeCount()).toEqual(0);
  expect(debugItems).toEqual(items);
});

test('Positional changes', async () => {
  const itemsIn = idItems(0, ['A', 'D', 'C', 'B']);
  const itemsOut = [itemsIn[0], itemsIn[3], itemsIn[2], itemsIn[1]];

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(tracker.changeCount()).toEqual(2);
  expect(tracker.change(itemsIn[1])).toEqual(new PositionChange(Direction.DOWN, 2));
  expect(tracker.change(itemsIn[3])).toEqual(new PositionChange(Direction.UP, 2));
  expect(tracker.change(itemsOut[1])).toEqual(new PositionChange(Direction.UP, 2));
  expect(tracker.change(itemsOut[3])).toEqual(new PositionChange(Direction.DOWN, 2));
  expect(debugItems).toEqual(itemsOut);
});

test('Additions', async () => {
  const itemsIn = idItems(0, ['A', 'B', 'C']);
  const itemsOut = [itemsIn[0], idItem(0, 'A2'), itemsIn[1], itemsIn[2], idItem(0, 'C2')];

  const tracker = new ChangeTracker();

  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(tracker.changeCount()).toEqual(2);
  expect(tracker.change(itemsOut[1])).toEqual(new Addition());
  expect(tracker.change(itemsOut[4])).toEqual(new Addition());
  expect(debugItems).toEqual(itemsOut);
});

test('Removals', async () => {
  const itemsIn = idItems(0, ['A', 'B', 'C', 'D']);
  const itemsOut = [itemsIn[1], itemsIn[3]];

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(tracker.changeCount()).toEqual(2);
  expect(tracker.change(debugItems[0])).toEqual(new Removal());
  expect(tracker.change(debugItems[2])).toEqual(new Removal());
  expect(debugItems).toEqual(itemsIn);
});

test('Total Replacement', async () => {
  const itemsIn = idItems(0, ['A', 'B', 'C']);
  const itemsOut = idItems(0, ['D', 'E', 'F']);

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(tracker.changeCount()).toEqual(6);
  expect(tracker.change(debugItems[0])).toEqual(new Removal());
  expect(tracker.change(debugItems[1])).toEqual(new Removal());
  expect(tracker.change(debugItems[2])).toEqual(new Removal());
  expect(tracker.change(debugItems[3])).toEqual(new Addition());
  expect(tracker.change(debugItems[4])).toEqual(new Addition());
  expect(tracker.change(debugItems[5])).toEqual(new Addition());
  expect(debugItems).toEqual([...itemsIn, ...itemsOut]);
});

test('Mixed1', async () => {
  const itemsIn = idItems(0, ['A', 'C', 'D']);
  const itemsOut = [idItem(0, 'B'), itemsIn[2], itemsIn[1]];

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(tracker.changeCount()).toEqual(4);
  expect(tracker.change(debugItems[0])).toEqual(new Removal());
  expect(tracker.change(debugItems[1])).toEqual(new Addition());
  expect(tracker.change(debugItems[2])).toEqual(new PositionChange(Direction.UP, 1));
  expect(tracker.change(debugItems[3])).toEqual(new PositionChange(Direction.DOWN, 1));
  expect(debugItems).toEqual([itemsIn[0], itemsOut[0], itemsIn[2], itemsIn[1]]);
});

test('Mixed2', async () => {
  const itemsIn = idItems(0, ['A', 'B', 'C', 'D', 'E', 'X', 'Y']);
  const itemsOut = [itemsIn[0], itemsIn[5], itemsIn[6], idItem(0, 'Z'), itemsIn[3], itemsIn[4]];

  const tracker = new ChangeTracker();
  const debugItems = detectChanges(tracker, itemsIn, itemsOut);
  expect(debugItems.map((item) => item.data['id'])).toEqual(['A', 'B', 'C', 'X', 'Y', 'Z', 'D', 'E']);
  expect(tracker.changeCount()).toEqual(7);
  expect(tracker.change(debugItems[0])).toBeUndefined();
  expect(tracker.change(debugItems[1])).toEqual(new Removal());
  expect(tracker.change(debugItems[2])).toEqual(new Removal());
  expect(tracker.change(debugItems[3])).toEqual(new PositionChange(Direction.UP, 2));
  expect(tracker.change(debugItems[4])).toEqual(new PositionChange(Direction.UP, 2));
  expect(tracker.change(debugItems[5])).toEqual(new Addition());
  expect(tracker.change(debugItems[6])).toEqual(new PositionChange(Direction.DOWN, 2));
  expect(tracker.change(debugItems[7])).toEqual(new PositionChange(Direction.DOWN, 2));
});
