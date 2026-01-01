/**
 * @typedef {(error?: Error | null | undefined, ...output: Array<any>) => void} Callback
 *   Callback.
 *
 * @typedef {(...input: Array<any>) => any} Middleware
 *   Ware.
 *
 * @typedef Pipeline
 *   Pipeline.
 * @property {Run} run
 *   Run the pipeline.
 * @property {Use} use
 *   Add middleware.
 *
 * @typedef {(...input: Array<any>) => void} Run
 *   Call all middleware.
 *
 *   Calls `done` on completion with either an error or the output of the
 *   last middleware.
 *
 *   > 👉 **Note**: as the length of input defines whether async functions get a
 *   > `next` function,
 *   > it’s recommended to keep `input` at one value normally.

 *
 * @typedef {(fn: Middleware) => Pipeline} Use
 *   Add middleware.
 */
/**
 * Create new middleware.
 *
 * @returns {Pipeline}
 *   Pipeline.
 */
export function trough(): Pipeline;
/**
 * Wrap `middleware` into a uniform interface.
 *
 * You can pass all input to the resulting function.
 * `callback` is then called with the output of `middleware`.
 *
 * If `middleware` accepts more arguments than the later given in input,
 * an extra `done` function is passed to it after that input,
 * which must be called by `middleware`.
 *
 * The first value in `input` is the main input value.
 * All other input values are the rest input values.
 * The values given to `callback` are the input values,
 * merged with every non-nullish output value.
 *
 * * if `middleware` throws an error,
 *   returns a promise that is rejected,
 *   or calls the given `done` function with an error,
 *   `callback` is called with that error
 * * if `middleware` returns a value or returns a promise that is resolved,
 *   that value is the main output value
 * * if `middleware` calls `done`,
 *   all non-nullish values except for the first one (the error) overwrite the
 *   output values
 *
 * @param {Middleware} middleware
 *   Function to wrap.
 * @param {Callback} callback
 *   Callback called with the output of `middleware`.
 * @returns {Run}
 *   Wrapped middleware.
 */
export function wrap(middleware: Middleware, callback: Callback): Run;
/**
 * Callback.
 */
export type Callback = (error?: Error | null | undefined, ...output: Array<any>) => void;
/**
 * Ware.
 */
export type Middleware = (...input: Array<any>) => any;
/**
 * Pipeline.
 */
export type Pipeline = {
    /**
     *   Run the pipeline.
     */
    run: Run;
    /**
     *   Add middleware.
     */
    use: Use;
};
/**
 * Call all middleware.
 *
 * Calls `done` on completion with either an error or the output of the
 * last middleware.
 *
 * > 👉 **Note**: as the length of input defines whether async functions get a
 * > `next` function,
 * > it’s recommended to keep `input` at one value normally.
 */
export type Run = (...input: Array<any>) => void;
/**
 * Add middleware.
 */
export type Use = (fn: Middleware) => Pipeline;
//# sourceMappingURL=index.d.ts.map