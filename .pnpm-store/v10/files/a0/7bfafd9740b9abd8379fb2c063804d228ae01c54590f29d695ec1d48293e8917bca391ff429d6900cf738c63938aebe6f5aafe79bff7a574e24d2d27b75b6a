interface TOCCrawlerArgs {
    /** Set generate mode to automatically set heading IDs. */
    mode?: 'generate' | undefined;
    /** Provide query list of elements. Defaults h2-h6. */
    queryElements?: string;
    scrollTarget?: string;
    /** Reload the action when this key value changes. */
    key?: unknown;
    prefix?: string;
    suffix?: string;
}
export declare function tocCrawler(node: HTMLElement, args?: TOCCrawlerArgs): {
    update(newArgs: TOCCrawlerArgs): void;
    destroy(): void;
};
export {};
