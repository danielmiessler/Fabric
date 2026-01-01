type Dict<T> = Record<string, T>;
type Arrayable<T> = T | T[];
type Default = Dict<any>;

declare function mri<T=Default>(args?: string[], options?: mri.Options): mri.Argv<T>;

declare namespace mri {
	export interface Options {
		boolean?: Arrayable<string>;
		string?: Arrayable<string>;
		alias?: Dict<Arrayable<string>>;
		default?: Dict<any>;
		unknown?(flag: string): void;
	}

	export type Argv<T=Default> = T & {
		_: string[];
	}
}

export = mri;
