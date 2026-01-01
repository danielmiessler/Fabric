import Item from './Item';
import ItemTransformer from './transformer/ItemTransformer';
import StageResult from './debug/StageResult';
import PageViewport from './parse/PageViewport';
export default class Debugger {
    fontMap: Map<string, object>;
    private pageViewports;
    pageCount: number;
    private transformers;
    private stageResultCache;
    stageNames: string[];
    stageDescriptions: string[];
    constructor(fontMap: Map<string, object>, pageViewports: PageViewport[], pageCount: number, inputSchema: string[], inputItems: Item[], transformers: ItemTransformer[]);
    stageResult(stageIndex: number): StageResult;
}
