import Item from '../Item';
import PageViewport from '../parse/PageViewport';
import EvaluationTracker from '../debug/EvaluationTracker';
import GlobalDefinition from '../GlobalDefinition';
import Globals from '../Globals';
export default class TransformContext {
    fontMap: Map<string, object>;
    pageViewports: PageViewport[];
    private globals;
    private evaluations;
    pageCount: number;
    constructor(fontMap: Map<string, object>, pageViewports: PageViewport[], globals: Globals, evaluations?: EvaluationTracker);
    trackEvaluation(item: Item, score?: any): void;
    globalIsDefined<T>(definition: GlobalDefinition<T>): boolean;
    getGlobal<T>(definition: GlobalDefinition<T>): T;
    getGlobalOptionally<T>(definition: GlobalDefinition<T>): T | undefined;
}
