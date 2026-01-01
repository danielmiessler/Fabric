import type { AST } from 'svelte-eslint-parser';
import type { Linter } from 'eslint';
type Define = {
    loc: AST.Position;
};
export declare class CommentDirectives {
    private readonly reportUnusedDisableDirectives;
    private readonly ruleId;
    private readonly lineDisableDirectives;
    private readonly blockDirectives;
    constructor(options: {
        reportUnusedDisableDirectives?: boolean;
        ruleId: string;
    });
    filterMessages(messages: Linter.LintMessage[]): Linter.LintMessage[];
    disableLine(line: number, rule: string | ((ruleId: string) => boolean), define: Define): void;
    disableBlock(pos: AST.Position, rule: string | ((ruleId: string) => boolean), define: Define): void;
    enableBlock(pos: AST.Position, rule: string | ((ruleId: string) => boolean), define: Define): void;
    private block;
}
export {};
