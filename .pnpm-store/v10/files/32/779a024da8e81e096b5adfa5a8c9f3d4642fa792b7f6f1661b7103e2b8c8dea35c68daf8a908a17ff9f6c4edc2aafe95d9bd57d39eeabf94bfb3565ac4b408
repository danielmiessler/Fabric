// Note: types exposed from `index.d.ts`
import rehypeParse from 'rehype-parse'
import rehypeStringify from 'rehype-stringify'
import {unified} from 'unified'

/**
 * Create a new unified processor that already uses `rehype-parse` and
 * `rehype-stringify`.
 */
export const rehype = unified().use(rehypeParse).use(rehypeStringify).freeze()
