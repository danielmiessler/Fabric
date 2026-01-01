import Item from '../Item';

export default interface ChangeIndex {
  /**
   * Return the number of tracked changes.
   */
  changeCount(): number;

  /**
   * Returns the change if for the given item if there is any
   * @param item
   */
  change(item: Item): Change | undefined;

  /**
   * Returs true if the given item has changed
   * @param item
   */
  hasChanged(item: Item): boolean;

  /**
   * Returns true if there is a change and it's category is ChangeCategory.PLUS
   * @param item
   */
  isPlusChange(item: Item): boolean;

  /**
   * Returns true if there is a change and it's category is ChangeCategory.NEUTRAL
   * @param item
   */
  isNeutralChange(item: Item): boolean;

  /**
   * Returns true if there is a change and it's category is ChangeCategory.MINUS
   * @param item
   */
  isMinusChange(item: Item): boolean;

  /**
   * Returns true if the item was removed.
   * @param item
   */
  isRemoved(item: Item): boolean;
}

export abstract class Change {
  constructor(public category: ChangeCategory) {}
}

// This is merely for coloring different kind of changes
export enum ChangeCategory {
  PLUS = 'PLUS',
  MINUS = 'MINUS',
  NEUTRAL = 'NEUTRAL',
}

export class Addition extends Change {
  constructor() {
    super(ChangeCategory.PLUS);
  }
}

export class Removal extends Change {
  constructor() {
    super(ChangeCategory.MINUS);
  }
}

export class ContentChange extends Change {
  constructor() {
    super(ChangeCategory.NEUTRAL);
  }
}

export enum Direction {
  UP = 'UP',
  DOWN = 'DOWN',
}

export class PositionChange extends Change {
  constructor(public direction: Direction, public amount: number) {
    super(direction === Direction.UP ? ChangeCategory.PLUS : ChangeCategory.MINUS);
  }
}
