/**
 * Generate a slug.
 *
 * Does not track previously generated slugs: repeated calls with the same value
 * will result in the exact same slug.
 * Use the `GithubSlugger` class to get unique slugs.
 *
 * @param  {string} value
 *   String of text to slugify
 * @param  {boolean} [maintainCase=false]
 *   Keep the current case, otherwise make all lowercase
 * @return {string}
 *   A unique slug string
 */
export function slug(value: string, maintainCase?: boolean | undefined): string;
/**
 * Slugger.
 */
export default class BananaSlug {
    /** @type {Record<string, number>} */
    occurrences: Record<string, number>;
    /**
     * Generate a unique slug.
    *
    * Tracks previously generated slugs: repeated calls with the same value
    * will result in different slugs.
    * Use the `slug` function to get same slugs.
     *
     * @param  {string} value
     *   String of text to slugify
     * @param  {boolean} [maintainCase=false]
     *   Keep the current case, otherwise make all lowercase
     * @return {string}
     *   A unique slug string
     */
    slug(value: string, maintainCase?: boolean | undefined): string;
    /**
     * Reset - Forget all previous slugs
     *
     * @return void
     */
    reset(): void;
}
