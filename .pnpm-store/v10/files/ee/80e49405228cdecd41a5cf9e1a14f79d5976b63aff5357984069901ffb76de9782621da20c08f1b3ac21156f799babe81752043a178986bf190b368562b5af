import type ItemMerger from './debug/ItemMerger';

interface Debug {
  /**
   * If set to true, all items will be shown as relevant.
   */
  readonly showAll?: boolean;
  /**
   * If this is set, the debug UI will group items and display a merged item.
   */
  readonly itemMerger?: ItemMerger;
}

export default interface TransformDescriptor {
  readonly requireColumns: string[];
  readonly consumesGlobels: string[];
  readonly producesGlobels: string[];
  /**
   * If this is set, the debug UI will group items and display a merged item.
   */
  readonly debug?: Debug;
}

const defaults: TransformDescriptor = {
  requireColumns: [],
  consumesGlobels: [],
  producesGlobels: [],
};

export function toDescriptor(partial: Partial<TransformDescriptor>): TransformDescriptor {
  return {
    ...defaults,
    ...partial,
  };
}
