'use strict';

const eslint = require('eslint');
const semver = require('semver');
const convertConfig = require('./shared/eslint-compat-utils.503aeaae.cjs');
const getUnsupported = require('./shared/eslint-compat-utils.cb53cf36.cjs');
const convertOption = require('./shared/eslint-compat-utils.808b5669.cjs');
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
let cacheRuleTester;
let cachePrefix = "";
function getRuleTester() {
  return cacheRuleTester != null ? cacheRuleTester : cacheRuleTester = getRuleTesterInternal();
  function getRuleTesterInternal() {
    if (semver__namespace.gte(eslint__namespace.Linter.version, "9.0.0-0")) {
      cachePrefix = "rule-to-test/";
      return eslint__namespace.RuleTester;
    }
    const flatRuleTester = getUnsupported.getUnsupported().FlatRuleTester;
    if (flatRuleTester) {
      cachePrefix = "rule-to-test/";
      return patchForV8FlatRuleTester(flatRuleTester);
    }
    return getRuleTesterClassFromLegacyRuleTester();
  }
}
function getRuleIdPrefix() {
  getRuleTester();
  return cachePrefix;
}
function patchForV8FlatRuleTester(flatRuleTester) {
  return class RuleTesterWithPatch extends flatRuleTester {
    constructor(options) {
      super(patchConfig(options));
    }
  };
  function patchConfig(config) {
    return {
      files: ["**/*.*"],
      ...config
    };
  }
}
function getRuleTesterClassFromLegacyRuleTester() {
  return class RuleTesterForV8 extends eslint__namespace.RuleTester {
    constructor(options) {
      var _a;
      const defineRules = [];
      const { processor, ...others } = options;
      super(
        convertConfig.convertConfigToRc(others, {
          defineRule(...args) {
            defineRules.push(args);
          }
        })
      );
      __publicField(this, "defaultProcessor");
      for (const args of defineRules) {
        (_a = this.linter) == null ? void 0 : _a.defineRule(...args);
      }
      this.defaultProcessor = processor;
    }
    run(name, rule, tests) {
      super.run(name, rule, {
        valid: (tests.valid || []).map(
          (test) => typeof test === "string" ? test : convert(test, this.defaultProcessor)
        ),
        invalid: (tests.invalid || []).map(
          (test) => convert(test, this.defaultProcessor)
        )
      });
    }
  };
  function convert(config, defaultProcessor) {
    const { processor: configProcessor, ...otherConfig } = config;
    const processor = configProcessor || defaultProcessor;
    const converted = convertConfig.convertConfigToRc(otherConfig);
    if (!processor) {
      return converted;
    }
    return {
      ...converted,
      filename: convertOption.convertOptionToLegacy(processor, config.filename, config)
    };
  }
}

exports.getRuleIdPrefix = getRuleIdPrefix;
exports.getRuleTester = getRuleTester;
