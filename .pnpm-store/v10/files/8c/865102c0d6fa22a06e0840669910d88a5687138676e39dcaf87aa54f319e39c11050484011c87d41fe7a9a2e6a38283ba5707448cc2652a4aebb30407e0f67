'use strict';

const eslint = require('eslint');
const semver = require('semver');
const convertConfig = require('./shared/eslint-compat-utils.503aeaae.cjs');
const getUnsupported = require('./shared/eslint-compat-utils.cb53cf36.cjs');
require('module');

function _interopNamespaceCompat(e) {
  if (e && typeof e === 'object' && 'default' in e) return e;
  const n = Object.create(null);
  if (e) {
    for (const k in e) {
      n[k] = e[k];
    }
  }
  n.default = e;
  return n;
}

const eslint__namespace = /*#__PURE__*/_interopNamespaceCompat(eslint);
const semver__namespace = /*#__PURE__*/_interopNamespaceCompat(semver);

var __defProp = Object.defineProperty;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __publicField = (obj, key, value) => {
  __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
  return value;
};
let cacheESLint, cacheLegacyESLint;
function getESLint() {
  return cacheESLint != null ? cacheESLint : cacheESLint = getESLintInternal();
  function getESLintInternal() {
    if (semver__namespace.gte(eslint__namespace.Linter.version, "9.0.0-0")) {
      return eslint__namespace.ESLint;
    }
    return getUnsupported.getUnsupported().FlatESLint || (eslint__namespace.ESLint ? getESLintClassFromLegacyESLint(eslint__namespace.ESLint) : getESLintClassFromLegacyESLint(getLegacyESLintClassFromCLIEngine()));
  }
}
function getLegacyESLint() {
  return cacheLegacyESLint != null ? cacheLegacyESLint : cacheLegacyESLint = getLegacyESLintInternal();
  function getLegacyESLintInternal() {
    return getUnsupported.getUnsupported().LegacyESLint || eslint__namespace.ESLint || getLegacyESLintClassFromCLIEngine();
  }
}
function getESLintClassFromLegacyESLint(legacyESLintClass) {
  return class ESLintFromLegacyESLint extends legacyESLintClass {
    static get version() {
      return legacyESLintClass.version;
    }
    constructor(options) {
      super(adjustOptions(options));
    }
  };
  function adjustOptions(options) {
    const {
      baseConfig: originalBaseConfig,
      overrideConfig: originalOverrideConfig,
      overrideConfigFile,
      ...newOptions
    } = options || {};
    if (originalBaseConfig) {
      const [baseConfig, plugins] = convertConfig$1(originalBaseConfig);
      newOptions.baseConfig = baseConfig;
      if (plugins) {
        newOptions.plugins = plugins;
      }
    }
    if (originalOverrideConfig) {
      const [overrideConfig, plugins] = convertConfig$1(originalOverrideConfig);
      newOptions.overrideConfig = overrideConfig;
      if (plugins) {
        newOptions.plugins = plugins;
      }
    }
    if (overrideConfigFile) {
      if (overrideConfigFile === true) {
        newOptions.useEslintrc = false;
      } else {
        newOptions.overrideConfigFile = overrideConfigFile;
      }
    }
    return newOptions;
  }
  function convertConfig$1(config) {
    const pluginDefs = {};
    const newConfigs = [];
    for (const configItem of Array.isArray(config) ? config : [config]) {
      const { plugins, ...otherConfig } = configItem;
      if (typeof otherConfig.processor !== "string")
        delete otherConfig.processor;
      const newConfig = {
        files: ["**/*.*", "*.*", "**/*", "*"],
        ...convertConfig.convertConfigToRc(otherConfig)
      };
      if (plugins) {
        newConfig.plugins = Object.keys(plugins);
      }
      Object.assign(pluginDefs, plugins);
      newConfigs.push(newConfig);
    }
    return [{ overrides: newConfigs }, pluginDefs];
  }
}
function getLegacyESLintClassFromCLIEngine() {
  const CLIEngine = eslint__namespace.CLIEngine;
  class LegacyESLintFromCLIEngine {
    constructor(options) {
      __publicField(this, "engine");
      const {
        overrideConfig: {
          plugins,
          globals,
          rules,
          overrides,
          ...overrideConfig
        } = {
          plugins: [],
          globals: {},
          rules: {},
          overrides: []
        },
        fix,
        reportUnusedDisableDirectives,
        plugins: pluginsMap,
        ...otherOptions
      } = options || {};
      const cliEngineOptions = {
        baseConfig: {
          ...overrides ? {
            overrides
          } : {}
        },
        fix: Boolean(fix),
        reportUnusedDisableDirectives: reportUnusedDisableDirectives ? reportUnusedDisableDirectives !== "off" : void 0,
        ...otherOptions,
        globals: globals ? Object.keys(globals).filter((n) => globals[n]) : void 0,
        plugins: plugins || [],
        rules: rules ? Object.fromEntries(
          Object.entries(rules).flatMap(
            ([ruleId, opt]) => opt ? [[ruleId, opt]] : []
          )
        ) : void 0,
        ...overrideConfig
      };
      this.engine = new CLIEngine(cliEngineOptions);
      for (const [name, plugin] of Object.entries(pluginsMap || {})) {
        this.engine.addPlugin(name, plugin);
      }
    }
    static get version() {
      return CLIEngine.version;
    }
    // eslint-disable-next-line @typescript-eslint/require-await -- ignore
    async lintText(...params) {
      var _a;
      const result = this.engine.executeOnText(params[0], (_a = params[1]) == null ? void 0 : _a.filePath);
      return result.results;
    }
    // eslint-disable-next-line @typescript-eslint/require-await -- ignore
    async lintFiles(...params) {
      const result = this.engine.executeOnFiles(
        Array.isArray(params[0]) ? params[0] : [params[0]]
      );
      return result.results;
    }
    // eslint-disable-next-line @typescript-eslint/require-await -- ignore
    static async outputFixes(...params) {
      return CLIEngine.outputFixes({
        results: params[0]
      });
    }
  }
  return LegacyESLintFromCLIEngine;
}

exports.getESLint = getESLint;
exports.getLegacyESLint = getLegacyESLint;
