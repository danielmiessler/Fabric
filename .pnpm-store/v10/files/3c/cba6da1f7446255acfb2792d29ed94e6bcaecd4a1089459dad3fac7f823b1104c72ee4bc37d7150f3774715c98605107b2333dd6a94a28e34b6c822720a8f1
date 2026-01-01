import * as eslint from 'eslint';
import * as semver from 'semver';
import { c as convertConfigToRc } from './shared/eslint-compat-utils.1a5060cf.mjs';
import { g as getUnsupported } from './shared/eslint-compat-utils.3ecba7ac.mjs';
import { c as convertOptionToLegacy } from './shared/eslint-compat-utils.cb6790c2.mjs';
import 'module';

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
    if (semver.gte(eslint.Linter.version, "9.0.0-0")) {
      cachePrefix = "rule-to-test/";
      return eslint.RuleTester;
    }
    const flatRuleTester = getUnsupported().FlatRuleTester;
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
  return class RuleTesterForV8 extends eslint.RuleTester {
    constructor(options) {
      var _a;
      const defineRules = [];
      const { processor, ...others } = options;
      super(
        convertConfigToRc(others, {
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
    const converted = convertConfigToRc(otherConfig);
    if (!processor) {
      return converted;
    }
    return {
      ...converted,
      filename: convertOptionToLegacy(processor, config.filename, config)
    };
  }
}

export { getRuleIdPrefix, getRuleTester };
