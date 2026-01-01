declare module 'locate-character' {
	export function getLocator(source: string, options?: Options | undefined): (search: string | number, index?: number | undefined) => Location | undefined;

	export function locate(source: string, search: string | number, options?: Options | undefined): Location | undefined;
	export type Location = Location_1;
	interface Options {
		offsetLine?: number;
		offsetColumn?: number;
		startIndex?: number;
	}

	interface Location_1 {
		line: number;
		column: number;
		character: number;
	}
}

//# sourceMappingURL=index.d.ts.map