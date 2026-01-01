declare module 'code-red' {
	export function b(strings: TemplateStringsArray, ...values: any[]): Node[];

	export function x(strings: TemplateStringsArray, ...values: any[]): Expression & {
		start: number;
		end: number;
	};

	export function p(strings: TemplateStringsArray, ...values: any[]): (Property | SpreadElement) & {
		start: number;
		end: number;
	};
	export function parse(source: string, opts: any): any;
	export function parseExpressionAt(source: string, index: number, opts: any): any;
	export type Expression = import('estree').Expression;
	export type Node = import('estree').Node;
	export type ObjectExpression = import('estree').ObjectExpression;
	export type Property = import('estree').Property;
	export type SpreadElement = import('estree').SpreadElement;
	export type CommentWithLocation = CommentWithLocation_1;
	/**
	 * @returns // TODO
	 */
	export function print(node: Node_1, opts?: PrintOptions): {
		code: string;
		map: any;
	};
	type Node_1 = import('estree').Node;
	type PrintOptions = {
		file?: string;
		sourceMapSource?: string;
		sourceMapContent?: string;
		sourceMapEncodeMappings?: boolean;
		getName?: (name: string) => string;
	};
	type Comment = import('estree').Comment;
	type CommentWithLocation_1 = Comment & {
		start: number;
		end: number;
		has_trailing_newline?: boolean;
	};
}

//# sourceMappingURL=index.d.ts.map