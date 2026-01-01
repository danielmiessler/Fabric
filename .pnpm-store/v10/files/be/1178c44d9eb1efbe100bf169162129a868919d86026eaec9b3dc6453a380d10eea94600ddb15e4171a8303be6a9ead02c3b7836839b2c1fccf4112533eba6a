import Item from '../Item';
import PageViewport from '../parse/PageViewport';
import EvaluationTracker from '../debug/EvaluationTracker';
import GlobalDefinition from '../GlobalDefinition';
import Globals from '../Globals';

export default class TransformContext {
  pageCount: number;

  constructor(
    public fontMap: Map<string, object>,
    public pageViewports: PageViewport[],
    private globals: Globals,
    private evaluations = new EvaluationTracker(),
  ) {
    this.pageCount = pageViewports.length;
  }

  trackEvaluation(item: Item, score: any = undefined) {
    this.evaluations.trackEvaluation(item, score);
  }

  globalIsDefined<T>(definition: GlobalDefinition<T>): boolean {
    return this.globals.isDefined(definition);
  }

  getGlobal<T>(definition: GlobalDefinition<T>): T {
    return this.globals.get(definition);
  }

  getGlobalOptionally<T>(definition: GlobalDefinition<T>): T | undefined {
    return this.globals.getOptional(definition);
  }
}
