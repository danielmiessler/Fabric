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
export declare abstract class Change {
    category: ChangeCategory;
    constructor(category: ChangeCategory);
}
export declare enum ChangeCategory {
    PLUS = "PLUS",
    MINUS = "MINUS",
    NEUTRAL = "NEUTRAL"
}
export declare class Addition extends Change {
    constructor();
}
export declare class Removal extends Change {
    constructor();
}
export declare class ContentChange extends Change {
    constructor();
}
export declare enum Direction {
    UP = "UP",
    DOWN = "DOWN"
}
export declare class PositionChange extends Change {
    direction: Direction;
    amount: number;
    constructor(direction: Direction, amount: number);
}
