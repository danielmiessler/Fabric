import TransformDescriptor from '../TransformDescriptor';
import AnnotatedColumn from './AnnotatedColumn';
import Item from '../Item';
import Page from './Page';
import ChangeIndex from './ChangeIndex';
import EvaluationIndex from './EvaluationIndex';
import Globals from '../Globals';
export default class StageResult {
    globals: Globals;
    descriptor: TransformDescriptor;
    schema: AnnotatedColumn[];
    pages: Page[];
    evaluations: EvaluationIndex;
    changes: ChangeIndex;
    messages: string[];
    constructor(globals: Globals, descriptor: TransformDescriptor, schema: AnnotatedColumn[], pages: Page[], evaluations: EvaluationIndex, changes: ChangeIndex, messages: string[]);
    itemsUnpacked(): Item[];
    itemsCleanedAndUnpacked(): Item[];
    selectPages(relevantChangesOnly: boolean, groupItems: boolean): Page[];
    pagesWithUnpackedItems(): Page[];
}
export declare function initialStage(inputSchema: string[], inputItems: Item[]): StageResult;
