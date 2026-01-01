import ChangeIndex, { Change } from './ChangeIndex';
import Item from '../Item';
export default class ChangeTracker implements ChangeIndex {
    private changes;
    private addChange;
    changeCount(): number;
    trackAddition(item: Item): void;
    trackRemoval(item: Item): void;
    trackPositionalChange(item: Item, oldPosition: number, newPosition: number): void;
    trackContentChange(item: Item): void;
    change(item: Item): Change | undefined;
    hasChanged(item: Item): boolean;
    isPlusChange(item: Item): boolean;
    isNeutralChange(item: Item): boolean;
    isMinusChange(item: Item): boolean;
    isRemoved(item: Item): boolean;
}
