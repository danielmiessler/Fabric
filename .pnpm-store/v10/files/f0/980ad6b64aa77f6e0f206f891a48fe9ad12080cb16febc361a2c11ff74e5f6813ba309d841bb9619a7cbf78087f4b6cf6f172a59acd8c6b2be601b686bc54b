import TransformDescriptor, { toDescriptor } from '../TransformDescriptor';
import { PARSE_SCHEMA } from '../PdfParser';
import AnnotatedColumn from './AnnotatedColumn';
import Item from '../Item';
import Page, { asPages } from './Page';
import ChangeIndex from './ChangeIndex';
import ChangeTracker from './ChangeTracker';
import ItemGroup from './ItemGroup';
import EvaluationIndex from './EvaluationIndex';
import EvaluationTracker from './EvaluationTracker';
import Globals from '../Globals';

export default class StageResult {
  constructor(
    public globals: Globals,
    public descriptor: TransformDescriptor,
    public schema: AnnotatedColumn[],
    public pages: Page[],
    public evaluations: EvaluationIndex,
    public changes: ChangeIndex,
    public messages: string[],
  ) {}

  itemsUnpacked(): Item[] {
    return this.pages.reduce((items: Item[], page: Page) => {
      page.itemGroups.forEach((itemGroup) => itemGroup.unpacked().forEach((item) => items.push(item)));
      return items;
    }, []);
  }

  itemsCleanedAndUnpacked(): Item[] {
    return this.pages.reduce((items: Item[], page: Page) => {
      page.itemGroups.forEach((itemGroup) =>
        itemGroup
          .unpacked()
          .filter((item) => !this.changes.isRemoved(item))
          .forEach((item) => items.push(item)),
      );
      return items;
    }, []);
  }

  selectPages(relevantChangesOnly: boolean, groupItems: boolean): Page[] {
    let result: Page[];

    // Ungroup pages
    if (!groupItems && this.descriptor?.debug?.itemMerger) {
      result = this.pagesWithUnpackedItems();
    } else {
      result = this.pages;
    }

    // Filter out item (groups) with no changes
    if (relevantChangesOnly && !this.descriptor.debug?.showAll === true) {
      result = result.map(
        (page) =>
          ({
            ...page,
            itemGroups: page.itemGroups.filter(
              (itemGroup) => this.evaluations.evaluated(itemGroup.top) || this.changes.hasChanged(itemGroup.top),
            ),
          } as Page),
      );
    }

    return result;
  }

  pagesWithUnpackedItems(): Page[] {
    return this.pages.map(
      (page) =>
        ({
          ...page,
          itemGroups: new Array<ItemGroup>().concat(
            ...page.itemGroups.map((itemGroup) => itemGroup.unpacked().map((item) => new ItemGroup(item))),
          ),
        } as Page),
    );
  }
}

export function initialStage(inputSchema: string[], inputItems: Item[]): StageResult {
  const schema = inputSchema.map((column) => ({ name: column }));
  const evaluations = new EvaluationTracker();
  const changes = new ChangeTracker();
  const pages = asPages(evaluations, changes, PARSE_SCHEMA, inputItems);
  const messages = [
    `Parsed ${inputItems.length === 0 ? 0 : inputItems[inputItems.length - 1].page + 1} pages with ${
      inputItems.length
    } items`,
  ];
  return new StageResult(
    new Globals(),
    toDescriptor({ debug: { showAll: true } }),
    schema,
    pages,
    evaluations,
    changes,
    messages,
  );
}
