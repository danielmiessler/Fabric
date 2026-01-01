export interface Stat {
	file: string;
	pkg?: string;
	start: number;
	end: number;
}

export interface StatCollection {
	name: string;
	options: CollectionOptions;

	start: (file: string) => () => void;
	stats: Stat[];
	packageStats?: PackageStats[];
	collectionStart: number;
	duration?: number;
	finish: () => Promise<void> | void;
	finished: boolean;
}

export interface PackageStats {
	pkg: string;
	files: number;
	duration: number;
}

export interface CollectionOptions {
	logInProgress: (collection: StatCollection, now: number) => boolean;
	logResult: (collection: StatCollection) => boolean;
}
