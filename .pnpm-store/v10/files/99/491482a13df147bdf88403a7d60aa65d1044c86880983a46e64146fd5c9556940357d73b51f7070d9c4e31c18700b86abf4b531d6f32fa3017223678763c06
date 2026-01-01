import Item from './Item';
import ItemTransformer from './transformer/ItemTransformer';
import TransformContext from './transformer/TransformContext';
import StageResult, { initialStage } from './debug/StageResult';
import ColumnAnnotation from './debug/ColumnAnnotation';
import type AnnotatedColumn from './debug/AnnotatedColumn';
import { detectChanges } from './debug/detectChanges';
import { asPages } from './debug/Page';
import EvaluationTracker from './debug/EvaluationTracker';
import ChangeTracker from './debug/ChangeTracker';
import PageViewport from './parse/PageViewport';
import Globals from './Globals';

// TODO only cache the parse ?
export default class Debugger {
  private transformers: ItemTransformer[];
  private stageResultCache: StageResult[];
  stageNames: string[];
  stageDescriptions: string[];

  constructor(
    public fontMap: Map<string, object>,
    private pageViewports: PageViewport[],
    public pageCount: number,
    inputSchema: string[],
    inputItems: Item[],
    transformers: ItemTransformer[],
  ) {
    this.transformers = transformers;
    this.stageNames = ['Parse Result', ...transformers.map((t) => t.name)];
    this.stageDescriptions = ['Initial items as parsed by PDFjs', ...transformers.map((t) => t.description)];
    this.stageResultCache = [initialStage(inputSchema, inputItems)];
  }

  stageResult(stageIndex: number): StageResult {
    for (let idx = 0; idx < stageIndex + 1; idx++) {
      if (!this.stageResultCache[idx]) {
        const evaluations = new EvaluationTracker();
        const transformer = this.transformers[idx - 1];
        const previousStageResult: StageResult = this.stageResultCache[idx - 1];
        const context = new TransformContext(
          this.fontMap,
          this.pageViewports,
          previousStageResult.globals,
          evaluations,
        );
        const previousItems = previousStageResult.itemsCleanedAndUnpacked();
        const inputSchema = toSimpleSchema(previousStageResult);
        const outputSchema = transformer.schemaTransformer(inputSchema);
        const itemResult = transformer.transform(context, [...previousItems]);
        const globals = new Globals(previousStageResult.globals).withValues(itemResult.globals);

        const changes = new ChangeTracker();
        const items = detectChanges(changes, previousItems, itemResult.items);
        const pages = asPages(evaluations, changes, outputSchema, items, transformer.descriptor.debug?.itemMerger);
        const messages = itemResult.messages;
        if (changes.changeCount() > 0 && messages.length === 0) {
          messages.unshift(`Detected ${changes.changeCount()} changes`);
        }

        this.stageResultCache.push(
          new StageResult(
            globals,
            transformer.descriptor,
            toAnnotatedSchema(inputSchema, outputSchema),
            pages,
            evaluations,
            changes,
            messages,
          ),
        );
      }
    }
    return this.stageResultCache[stageIndex];
  }
}

function toSimpleSchema(stageResult: StageResult): string[] {
  return stageResult.schema
    .filter((column) => !column.annotation || column.annotation !== ColumnAnnotation.REMOVED)
    .map((column) => column.name);
}

function toAnnotatedSchema(inputSchema: string[], outputSchema: string[]): AnnotatedColumn[] {
  const annotatedSchema: AnnotatedColumn[] = [];
  let out_idx = 0;
  for (let in_idx = 0; in_idx < inputSchema.length; in_idx++) {
    const nextInputColumn = inputSchema[in_idx];
    const indexInOut = outputSchema.indexOf(nextInputColumn);
    if (indexInOut === -1) {
      annotatedSchema.push({ name: nextInputColumn, annotation: ColumnAnnotation.REMOVED });
    } else if (indexInOut > out_idx) {
      while (out_idx < indexInOut) {
        annotatedSchema.push({ name: outputSchema[out_idx], annotation: ColumnAnnotation.ADDED });
        out_idx++;
      }
      annotatedSchema.push({ name: nextInputColumn });
      out_idx++;
    } else {
      annotatedSchema.push({ name: nextInputColumn });
      out_idx++;
    }
  }
  for (let index = out_idx; index < outputSchema.length; index++) {
    annotatedSchema.push({ name: outputSchema[index], annotation: ColumnAnnotation.ADDED });
  }
  return annotatedSchema;
}
