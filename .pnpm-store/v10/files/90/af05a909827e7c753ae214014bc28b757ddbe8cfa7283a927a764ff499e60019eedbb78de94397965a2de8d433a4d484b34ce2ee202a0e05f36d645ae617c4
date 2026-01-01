// Note: types exposed from `automatic-runtime-html.d.ts` because TS has bugs
// when generating types.
import {createAutomaticRuntime} from './create-automatic-runtime.js'
import {h} from './index.js'

// Export `JSX` as a global for TypeScript.
export * from './jsx-automatic.js'

export const {Fragment, jsxDEV, jsxs, jsx} = createAutomaticRuntime(h)
