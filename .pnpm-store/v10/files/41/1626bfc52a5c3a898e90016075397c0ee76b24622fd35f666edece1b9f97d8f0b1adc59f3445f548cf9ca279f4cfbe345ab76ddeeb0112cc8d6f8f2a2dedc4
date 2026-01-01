import { regex } from './regex.js'

const own = Object.hasOwnProperty

/**
 * Slugger.
 */
export default class BananaSlug {
  /**
   * Create a new slug class.
   */
  constructor () {
    /** @type {Record<string, number>} */
    // eslint-disable-next-line no-unused-expressions
    this.occurrences

    this.reset()
  }

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
  slug (value, maintainCase) {
    const self = this
    let result = slug(value, maintainCase === true)
    const originalSlug = result

    while (own.call(self.occurrences, result)) {
      self.occurrences[originalSlug]++
      result = originalSlug + '-' + self.occurrences[originalSlug]
    }

    self.occurrences[result] = 0

    return result
  }

  /**
   * Reset - Forget all previous slugs
   *
   * @return void
   */
  reset () {
    this.occurrences = Object.create(null)
  }
}

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
export function slug (value, maintainCase) {
  if (typeof value !== 'string') return ''
  if (!maintainCase) value = value.toLowerCase()
  return value.replace(regex, '').replace(/ /g, '-')
}
