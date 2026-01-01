import type * as mri from 'mri';

type Arrayable<T> = T | T[];

declare function sade(usage: string, isSingle?: boolean): sade.Sade;

declare namespace sade {
	export type Handler = (...args: any[]) => any;
	export type Value = number | string | boolean | null;

	export interface LazyOutput {
		name: string;
		handler: Handler;
		args: string[];
	}

	export interface Sade {
		command(usage: string, description?: string, options?: {
			alias?: Arrayable<string>;
			default?: boolean;
		}): Sade;

		option(flag: string, description?: string, value?: Value): Sade;
		action(handler: Handler): Sade;
		describe(text: Arrayable<string>): Sade;
		alias(...names: string[]): Sade;
		example(usage: string): Sade;

		parse(arr: string[], opts: { lazy: true } & mri.Options): LazyOutput;
		parse(arr: string[], opts?: { lazy?: boolean } & mri.Options): void;

		version(value: string): Sade;
		help(cmd?: string): void;
	}
}

export = sade;
