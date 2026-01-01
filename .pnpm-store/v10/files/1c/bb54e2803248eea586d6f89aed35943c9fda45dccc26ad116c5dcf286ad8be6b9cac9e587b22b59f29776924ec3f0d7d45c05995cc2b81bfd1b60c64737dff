import type {Value} from 'vfile'
import type {CompileResults} from './lib/index.js'

export type {
  // `CompileResultMap` is typed and exposed below.
  CompileResults,
  Compiler,
  // `Data` is typed and exposed below.
  Parser,
  Pluggable,
  PluggableList,
  Plugin,
  PluginTuple,
  Preset,
  ProcessCallback,
  Processor,
  RunCallback,
  // `Settings` is typed and exposed below.
  TransformCallback,
  Transformer
} from './lib/index.js'

export {unified} from './lib/index.js'

// See: <https://github.com/sindresorhus/type-fest/blob/main/source/empty-object.d.ts>
declare const emptyObjectSymbol: unique symbol

/**
 * Interface of known results from compilers.
 *
 * Normally, compilers result in text ({@linkcode Value} of `vfile`).
 * When you compile to something else, such as a React node (as in,
 * `rehype-react`), you can augment this interface to include that type.
 *
 * ```ts
 * import type {ReactNode} from 'somewhere'
 *
 * declare module 'unified' {
 *   interface CompileResultMap {
 *     // Register a new result (value is used, key should match it).
 *     ReactNode: ReactNode
 *   }
 * }
 *
 * export {} // You may not need this, but it makes sure the file is a module.
 * ```
 *
 * Use {@linkcode CompileResults} to access the values.
 */
export interface CompileResultMap {
  // Note: if `Value` from `VFile` is changed, this should too.
  Uint8Array: Uint8Array
  string: string
}

/**
 * Interface of known data that can be supported by all plugins.
 *
 * Typically, options can be given to a specific plugin, but sometimes it makes
 * sense to have information shared with several plugins.
 * For example, a list of HTML elements that are self-closing, which is needed
 * during all phases.
 *
 * To type this, do something like:
 *
 * ```ts
 * declare module 'unified' {
 *   interface Data {
 *     htmlVoidElements?: Array<string> | undefined
 *   }
 * }
 *
 * export {} // You may not need this, but it makes sure the file is a module.
 * ```
 */
export interface Data {
  settings?: Settings | undefined
}

/**
 * Interface of known extra options, that can be supported by parser and
 * compilers.
 *
 * This exists so that users can use packages such as `remark`, which configure
 * both parsers and compilers (in this case `remark-parse` and
 * `remark-stringify`), and still provide options for them.
 *
 * When you make parsers or compilers, that could be packaged up together,
 * you should support `this.data('settings')` as input and merge it with
 * explicitly passed `options`.
 * Then, to type it, using `remark-stringify` as an example, do something like:
 *
 * ```ts
 * declare module 'unified' {
 *   interface Settings {
 *     bullet: '*' | '+' | '-'
 *     // …
 *   }
 * }
 *
 * export {} // You may not need this, but it makes sure the file is a module.
 * ```
 */
export interface Settings {
  [emptyObjectSymbol]?: never
}
