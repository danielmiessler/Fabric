declare module 'devalue' {
	/**
	 * Turn a value into the JavaScript that creates an equivalent value
	 * 
	 */
	export function uneval(value: any, replacer?: (value: any) => string | void): string;
	/**
	 * Turn a value into a JSON string that can be parsed with `devalue.parse`
	 * 
	 */
	export function stringify(value: any, reducers?: Record<string, (value: any) => any>): string;
	/**
	 * Revive a value serialized with `devalue.stringify`
	 * 
	 */
	export function parse(serialized: string, revivers?: Record<string, (value: any) => any>): any;
	/**
	 * Revive a value flattened with `devalue.stringify`
	 * 
	 */
	export function unflatten(parsed: number | any[], revivers?: Record<string, (value: any) => any>): any;

	export {};
}

//# sourceMappingURL=index.d.ts.map