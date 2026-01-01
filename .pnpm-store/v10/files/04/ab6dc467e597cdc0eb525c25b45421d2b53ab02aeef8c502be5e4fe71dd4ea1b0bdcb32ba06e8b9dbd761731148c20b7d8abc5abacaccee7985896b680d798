import { Parser } from 'acorn';

export function tsPlugin(options?: {
	dts?: boolean;
	/** Whether to use JSX. Defaults to false */
	jsx?:
		| boolean
		| {
				allowNamespaces?: boolean;
				allowNamespacedObjects?: boolean;
		  };
}): (BaseParser: typeof Parser) => typeof Parser;
