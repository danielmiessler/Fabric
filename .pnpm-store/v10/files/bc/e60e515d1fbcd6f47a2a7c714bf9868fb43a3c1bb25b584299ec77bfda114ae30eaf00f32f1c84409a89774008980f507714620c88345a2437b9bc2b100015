const globalName = '___SVELTE_HMR_HOT_API'
const globalAdapterName = '___SVELTE_HMR_HOT_API_PROXY_ADAPTER'

const defaultHotOptions = {
  // preserve all local state
  preserveLocalState: false,

  // escape hatchs from preservation of local state
  //
  // disable preservation of state for this component
  noPreserveStateKey: ['@hmr:reset', '@!hmr'],
  // enable preservation of state for all variables in this component
  preserveAllLocalStateKey: '@hmr:keep-all',
  // enable preservation of state for a given variable (must be inline or
  // above the target variable or variables; can be repeated)
  preserveLocalStateKey: '@hmr:keep',

  // don't reload on fatal error
  noReload: false,
  // try to recover after runtime errors during component init
  // defaults to false because some runtime errors are fatal and require a full reload
  optimistic: false,
  // auto accept modules of components that have named exports (i.e. exports
  // from context="module")
  acceptNamedExports: true,
  // auto accept modules of components have accessors (either accessors compile
  // option, or <svelte:option accessors={true} />) -- this means that if you
  // set accessors compile option globally, you must also set this option to
  // true, or no component will be hot reloaded (but there are a lot of edge
  // cases that HMR can't support correctly with accessors)
  acceptAccessors: true,
  // only inject CSS instead of recreating components when only CSS changes
  injectCss: false,
  // to mitigate FOUC between dispose (remove stylesheet) and accept
  cssEjectDelay: 100,

  // Svelte Native mode
  native: false,
  // name of the adapter import binding
  importAdapterName: globalAdapterName,

  // disable runtime error overlay
  noOverlay: false,

  // set to true on systems that supports them, to avoid useless caught /
  // managed (but still...) exceptions, by using Svelte's current_component
  // instead of get_current_component
  allowLiveBinding: false,

  // use acceptExports instead of accept, to fix support of context:module
  partialAccept: false,
}

const defaultHotApi = 'hot-api-esm.js'

const json = JSON.stringify

const posixify = file => file.replace(/[/\\]/g, '/')

const renderApplyHmr = ({
  id,
  cssId,
  nonCssHash,
  hotOptions: {
    injectCss,
    partialAccept,
    _accept = partialAccept ? "acceptExports(['default'])" : 'accept()',
  },
  hotOptionsJson,
  hotApiImport,
  adapterImport,
  importAdapterName,
  meta,
  acceptable,
  preserveLocalState,
  emitCss,
  imports = [
    `import * as ${globalName} from ${JSON.stringify(hotApiImport)}`,
    `import { adapter as ${importAdapterName} } from ${JSON.stringify(
      adapterImport
    )}`,
  ],
}) =>
  // this silly formatting keeps all original characters in their position,
  // thus saving us from having to provide a sourcemap
  //
  // NOTE the `if (false) accept()` line is for tools that wants to see the
  // accept call in the actual accepted module to enable HMR (Vite and, I
  // believe, Snowpack 3)
  //
  `${imports.join(';')};${`
    if (${meta} && ${meta}.hot) {
      ${acceptable ? `if (false) ${meta}.hot.${_accept};` : ''};
      $2 = ${globalName}.applyHmr({
        m: ${meta},
        id: ${json(id)},
        hotOptions: ${hotOptionsJson},
        Component: $2,
        ProxyAdapter: ${importAdapterName},
        acceptable: ${json(acceptable)},
        preserveLocalState: ${json(preserveLocalState)},
        ${cssId ? `cssId: ${json(cssId)},` : ''}
        ${nonCssHash ? `nonCssHash: ${json(nonCssHash)},` : ''}
        emitCss: ${json(emitCss)},
      });
    }
  `
    .split('\n')
    .map(x => x.trim())
    .filter(Boolean)
    .join(' ')}
export default $2;
${
  // NOTE when doing CSS only voodoo, we have to inject the stylesheet as soon
  // as the component is loaded because Svelte normally do that when a component
  // is instantiated, but we might already have instances in the large when a
  // component is loaded with HMR
  cssId && injectCss && !emitCss
    ? `
if (typeof add_css !== 'undefined' && !document.getElementById(${json(
        cssId
      )})) add_css();`
    : ``
}
`

// replace from last occurrence of the splitter
const replaceLast = (string, splitter, regex, replacer) => {
  const lastIndex = string.lastIndexOf(splitter)
  return (
    string.slice(0, lastIndex) +
    string.slice(lastIndex).replace(regex, replacer)
  )
}

// https://github.com/darkskyapp/string-hash/blob/master/index.js
// (via https://github.com/sveltejs/svelte/blob/91d758e35b2b2154512ddd11e6b6d9d65708a99e/src/compiler/compile/utils/hash.ts#L2)
const stringHashcode = str => {
  let hash = 5381
  let i = str.length
  while (i--) hash = ((hash << 5) - hash) ^ str.charCodeAt(i)
  return (hash >>> 0).toString(36)
}

const normalizeJsCode = (code, cssHash) =>
  code
    // ignore css hashes in the code (that have changed, necessarily)
    .replace(new RegExp('\\b' + cssHash + '\\b', 'g'), '')
    // Svelte now adds locations in dev mode, code locations can change when
    // CSS change, but we're unaffected (not real behaviour changes)
    .replace(/\badd_location\s*\([^)]*\)\s*;?/g, '')

const parseCssId = (code, compileOptions = {}, parseHash, originalCode) => {
  // the regex matching is very pretty conservative 'cause I don't want to
  // match something else by error... I'm probably make it more lax if I have
  // to fix it 3 times in a single week ¯\_(ツ)_/¯
  let match = /^function add_css\(\) \{[\s\S]*?^}/m.exec(code)

  if (!match) {
    // guard: injectCss is off, no need to compute hashes
    if (!parseHash) return {}
    // guard: compile.css is true, so we should have found the add_css function,
    //        something unexpected is unraveling here, fall back to caution
    if (compileOptions.css) return {}
    // trying to get CSS id the same way as Svelte does it
    match = /<style[^>]*>([\s\S]*)<\/\s*style\s*>/.exec(originalCode)
    const cssHash = match && match[1] ? stringHashcode(match[1]) : null
    if (!cssHash) return {}
    return {
      cssId: `svelte-${cssHash}-style`,
      nonCssHash: stringHashcode(normalizeJsCode(code, cssHash)),
    }
  }

  const codeExceptCSS =
    code.slice(0, match.index) + code.slice(match.index + match[0].length)

  match = /\bstyle\.id\s*=\s*(['"])([^'"]*)\1/.exec(match[0])
  const cssId = match ? match[2] : null

  if (!parseHash || !cssId) return { cssId }

  const cssHash = cssId.split('-')[1]
  const nonCssHash = stringHashcode(normalizeJsCode(codeExceptCSS, cssHash))

  return { cssId, nonCssHash }
}

const isNamedExport = v => v.export_name && v.module

const isProp = v => v.export_name && !v.module

const resolvePackageImport = (name = 'svelte-hmr', version = '') => {
  const versionSpec = version ? '@' + version : ''
  return target => `${name}${versionSpec}/${target}`
}

const createMakeHot = ({ resolveAbsoluteImport, pkg = {} }) => ({
  // NOTE hotOptions can be customized by end user through plugin options, while
  // options passed to this function can only customized by the plugin implementer
  //
  // meta can be 'import.meta' or 'module'
  // const createMakeHot = (hotApi = defaultHotApi, options) => {
  walk,
  meta = 'import.meta',

  hotApi: customHotApi = '',
  adapter: customAdapter = '',
  // use absolute file paths to import runtime deps of svelte-hmr
  // - pnpm needs absolute paths because svelte-hmr, being a transitive dep via
  //   the bundler plugin, is not directly importable from user's project
  //   (see https://github.com/rixo/svelte-hmr/issues/11)
  // - Snowpack source=remote needs a version number, otherwise it will try to
  //   load import, resulting in a possible version mismatch between the code
  //   transform and the runtime
  //   (see: https://github.com/rixo/svelte-hmr/issues/27#issuecomment-800142127)
  absoluteImports = true,
  versionNonAbsoluteImports = true,

  hotOptions: hotOptionsArg = {},
}) => {
  const hotOptions = { ...defaultHotOptions, ...hotOptionsArg }

  const hotOptionsJson = JSON.stringify(hotOptions)

  // NOTE Native adapter cannot be required in code (as opposed to this
  // generated code) because it requires modules from NativeScript's code that
  // are not resolvable for non-native users (and those missing modules would
  // prevent webpack from building).
  //
  // careful with relative paths
  // (see https://github.com/rixo/svelte-hmr/issues/11)
  const defaultAdapter = hotOptions.native
    ? 'svelte-native/proxy-adapter-native.js'
    : 'proxy-adapter-dom.js'

  const resolveImport = absoluteImports
    ? resolveAbsoluteImport
    : resolvePackageImport(pkg.name, versionNonAbsoluteImports && pkg.version)

  const adapterImport = posixify(
    customAdapter || resolveImport('runtime/' + defaultAdapter)
  )

  const hotApiImport = posixify(
    customHotApi || resolveImport('runtime/' + defaultHotApi)
  )

  const resolvePreserveLocalStateKey = ({
    preserveLocalStateKey,
    compiled,
  }) => {
    const containsKey = comments =>
      comments &&
      comments.some(({ value }) => value.includes(preserveLocalStateKey))

    const variables = new Set()

    const addReference = node => {
      if (!node.name) {
        // eslint-disable-next-line no-console
        console.warn('Incorrect identifier for preserveLocalStateKey')
      }
      variables.add(node.name)
    }

    const processNodes = targets => targets.forEach(processNode)

    const processNode = node => {
      switch (node.type) {
        case 'Identifier':
          variables.add(node.name)
          return true
        case 'UpdateExpression':
          addReference(node.argument)
          return true
        case 'VariableDeclarator':
          addReference(node.id)
          return true
        case 'AssignmentExpression':
          processNode(node.left)
          return true
        case 'ExpressionStatement':
          processNode(node.expression)
          return true

        case 'VariableDeclaration':
          processNodes(node.declarations)
          return true
        case 'SequenceExpression': // ++, --
          processNodes(node.expressions)
          return true
      }
      return false
    }

    const stack = []

    if (compiled.ast.instance) {
      walk(compiled.ast.instance, {
        leave() {
          stack.shift()
        },
        enter(node) {
          stack.unshift(node)
          if (
            containsKey(node.leadingComments) ||
            containsKey(node.trailingComments)
          ) {
            stack
              .slice(0, 3)
              .reverse()
              .some(processNode)
          }
        },
      })
    }

    return [...variables]
  }

  const resolvePreserveLocalState = ({
    hotOptions,
    originalCode,
    compiled,
  }) => {
    const {
      preserveLocalState,
      noPreserveStateKey,
      preserveLocalStateKey,
      preserveAllLocalStateKey,
    } = hotOptions
    if (originalCode) {
      const hasKey = key => {
        const test = k => originalCode.indexOf(k) !== -1
        return Array.isArray(key) ? key.some(test) : test(key)
      }
      // noPreserveStateKey
      if (noPreserveStateKey && hasKey(noPreserveStateKey)) {
        return false
      }
      // preserveAllLocalStateKey
      if (preserveAllLocalStateKey && hasKey(preserveAllLocalStateKey)) {
        return true
      }
      // preserveLocalStateKey
      if (preserveLocalStateKey && hasKey(preserveLocalStateKey)) {
        // returns an array of variable names to preserve
        return resolvePreserveLocalStateKey({ preserveLocalStateKey, compiled })
      }
    }
    // preserveLocalState
    if (preserveLocalState) {
      return true
    }
    return false
  }

  const hasAccessorsOption = compiled => {
    if (!compiled.ast || !compiled.ast.html) return
    let accessors = false
    walk(compiled.ast.html, {
      enter(node) {
        if (accessors) return
        if (node.type !== 'Options') return
        if (!node.attributes) return
        accessors = node.attributes.some(
          ({ name, value }) => name === 'accessors' && value
        )
      },
    })
    return accessors
  }

  const isAcceptable = (hotOptions, compileOptions, compiled) => {
    if (!compiled || !compileOptions) {
      // this should never happen, since it's the bundler plugins that control
      // what version of svelte-hmr they embark, and they should be kept up to
      // date with our API
      //
      // eslint-disable-next-line no-console
      console.warn(
        'WARNING Your bundler plugin is outdated for this version of svelte-hmr'
      )
      return true
    }

    const { vars } = compiled

    // if the module has named exports (in context="module"), then we can't
    // auto accept the component, since all the consumers need to be aware of
    // the change (e.g. rerender with the new exports value)
    if (!hotOptions.acceptNamedExports && vars.some(isNamedExport)) {
      return false
    }

    // ...same for accessors: if accessible props change, then we need to
    // rerender/rerun all the consumers to reflect the change
    //
    // NOTE we can still accept components with no props, since they won't
    // have accessors... this is actually all we can safely accept in this case
    //
    if (
      !hotOptions.acceptAccessors &&
      // we test is we have props first, because searching for the
      // <svelte:options /> tag in the AST is probably the most expensive here
      vars.some(isProp) &&
      (compileOptions.accessors || hasAccessorsOption(compiled))
    ) {
      return false
    }

    return true
  }

  const parseMakeHotArgs = args => {
    // case: named args (object)
    if (args.length === 1) return args[0]
    // case: legacy (positional)
    const [
      id,
      compiledCode,
      hotOptions,
      compiled,
      originalCode,
      compileOptions,
    ] = args
    return {
      id,
      compiledCode,
      hotOptions,
      compiled,
      originalCode,
      compileOptions,
    }
  }

  function makeHot(...args) {
    const {
      id,
      compiledCode,
      compiled,
      originalCode,
      compileOptions,
    } = parseMakeHotArgs(args)

    const { importAdapterName, injectCss } = hotOptions

    const emitCss =
      compileOptions &&
      (compileOptions.css === false || compileOptions.css === 'external')

    const preserveLocalState = resolvePreserveLocalState({
      hotOptions,
      originalCode,
      compiled,
    })

    const replacement = renderApplyHmr({
      id,
      // adds cssId & nonCssHash
      ...((injectCss || !emitCss) &&
        parseCssId(compiledCode, compileOptions, injectCss, originalCode)),
      hotOptions,
      hotOptionsJson,
      preserveLocalState,
      hotApiImport,
      adapterImport,
      importAdapterName,
      meta,
      acceptable: isAcceptable(hotOptions, compileOptions, compiled),
      // CSS is handled outside of Svelte: don't tamper with it!
      emitCss,
    })

    // NOTE export default can appear in strings in use code
    // see: https://github.com/rixo/svelte-hmr/issues/34
    return replaceLast(
      compiledCode,
      'export default',
      /(\n?export default ([^;]*);)/,
      (match, $1, $2) => replacement.replace(/\$2/g, () => $2)
    )
  }

  // rollup-plugin-svelte-hot needs hotApi path (for tests)
  // makeHot.hotApi = hotApi
  Object.defineProperty(makeHot, 'hotApi', {
    get() {
      // TODO makeHot.hotApi has been lost in 0.12 => 0.13... still needed?
      debugger // eslint-disable-line no-debugger
      return undefined
    },
  })

  return makeHot
}

module.exports = createMakeHot
