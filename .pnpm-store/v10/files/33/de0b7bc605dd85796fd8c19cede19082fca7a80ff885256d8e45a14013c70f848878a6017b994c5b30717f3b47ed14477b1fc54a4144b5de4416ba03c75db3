import * as eslint from 'eslint';
import * as semver from 'semver';
import { createRequire } from 'module';

function safeRequire(name) {
  try {
    return createRequire(`${process.cwd()}/__placeholder__.js`)(name);
  } catch {
    return void 0;
  }
}
function safeRequireResolve(name) {
  try {
    return createRequire(`${process.cwd()}/__placeholder__.js`).resolve(name);
  } catch {
    return name;
  }
}

const builtInGlobals = /* @__PURE__ */ new Map([
  [
    3,
    Object.entries({
      Array: false,
      Boolean: false,
      constructor: false,
      Date: false,
      decodeURI: false,
      decodeURIComponent: false,
      encodeURI: false,
      encodeURIComponent: false,
      Error: false,
      escape: false,
      eval: false,
      EvalError: false,
      Function: false,
      hasOwnProperty: false,
      Infinity: false,
      isFinite: false,
      isNaN: false,
      isPrototypeOf: false,
      Math: false,
      NaN: false,
      Number: false,
      Object: false,
      parseFloat: false,
      parseInt: false,
      propertyIsEnumerable: false,
      RangeError: false,
      ReferenceError: false,
      RegExp: false,
      String: false,
      SyntaxError: false,
      toLocaleString: false,
      toString: false,
      TypeError: false,
      undefined: false,
      unescape: false,
      URIError: false,
      valueOf: false
    })
  ],
  [
    5,
    Object.entries({
      JSON: false
    })
  ],
  [
    2015,
    Object.entries({
      ArrayBuffer: false,
      DataView: false,
      Float32Array: false,
      Float64Array: false,
      Int16Array: false,
      Int32Array: false,
      Int8Array: false,
      Intl: false,
      Map: false,
      Promise: false,
      Proxy: false,
      Reflect: false,
      Set: false,
      Symbol: false,
      Uint16Array: false,
      Uint32Array: false,
      Uint8Array: false,
      Uint8ClampedArray: false,
      WeakMap: false,
      WeakSet: false
    })
  ],
  [
    2017,
    Object.entries({
      Atomics: false,
      SharedArrayBuffer: false
    })
  ],
  [
    2020,
    Object.entries({
      BigInt: false,
      BigInt64Array: false,
      BigUint64Array: false,
      globalThis: false
    })
  ],
  [
    2021,
    Object.entries({
      AggregateError: false,
      FinalizationRegistry: false,
      WeakRef: false
    })
  ]
]);
function convertConfigToRc(config, linter) {
  var _a, _b;
  if (Array.isArray(config)) {
    throw new Error("Array config is not supported.");
  }
  const {
    languageOptions: originalLanguageOptions,
    plugins,
    ...newConfig
  } = config;
  if (originalLanguageOptions) {
    const {
      parser,
      globals,
      parserOptions,
      ecmaVersion,
      sourceType,
      ...languageOptions
    } = originalLanguageOptions;
    newConfig.parserOptions = {
      ...!ecmaVersion || ecmaVersion === "latest" ? { ecmaVersion: getLatestEcmaVersion() } : { ecmaVersion },
      ...sourceType ? { sourceType } : { sourceType: "module" },
      ...languageOptions,
      ...parserOptions,
      ...newConfig.parserOptions
    };
    const resolvedEcmaVersion = newConfig.parserOptions.ecmaVersion;
    newConfig.globals = {
      ...Object.fromEntries(
        [...builtInGlobals.entries()].flatMap(
          ([version, editionGlobals]) => resolvedEcmaVersion < version ? [] : editionGlobals
        )
      ),
      ...newConfig.globals
    };
    if (globals) {
      newConfig.globals = {
        ...globals,
        ...newConfig.globals
      };
    }
    if (parser && !newConfig.parser) {
      const parserName = getParserName(parser);
      newConfig.parser = parserName;
      (_a = linter == null ? void 0 : linter.defineParser) == null ? void 0 : _a.call(linter, parserName, parser);
    }
  }
  if (plugins) {
    for (const [pluginName, plugin] of Object.entries(plugins)) {
      for (const [ruleName, rule] of Object.entries(plugin.rules || {})) {
        (_b = linter == null ? void 0 : linter.defineRule) == null ? void 0 : _b.call(linter, `${pluginName}/${ruleName}`, rule);
      }
    }
  }
  newConfig.env = {
    es6: true,
    ...newConfig.env
  };
  return newConfig;
}
function getParserName(parser) {
  var _a;
  const name = ((_a = parser.meta) == null ? void 0 : _a.name) || parser.name;
  if (name === "typescript-eslint/parser") {
    return safeRequireResolve("@typescript-eslint/parser");
  } else if (name == null && parser === safeRequire("@typescript-eslint/parser"))
    return safeRequireResolve("@typescript-eslint/parser");
  return safeRequireResolve(name);
}
function getLatestEcmaVersion() {
  const eslintVersion = eslint.Linter.version;
  return semver.gte(eslintVersion, "8.0.0") ? "latest" : semver.gte(eslintVersion, "7.8.0") ? 2021 : semver.gte(eslintVersion, "6.2.0") ? 2020 : semver.gte(eslintVersion, "5.0.0") ? 2019 : 2018;
}

export { convertConfigToRc as c };
