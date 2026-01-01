import * as eslint from 'eslint';
import * as semver from 'semver';
import { c as convertConfigToRc } from './shared/eslint-compat-utils.1a5060cf.mjs';
import { c as convertOptionToLegacy } from './shared/eslint-compat-utils.cb6790c2.mjs';
import 'module';

let cacheLinter;
function getLinter() {
  return cacheLinter != null ? cacheLinter : cacheLinter = getLinterInternal();
  function getLinterInternal() {
    if (semver.gte(eslint.Linter.version, "9.0.0-0")) {
      return eslint.Linter;
    }
    return getLinterClassFromLegacyLinter();
  }
}
function getLinterClassFromLegacyLinter() {
  return class LinterFromLegacyLinter extends eslint.Linter {
    static get version() {
      return eslint.Linter.version;
    }
    verify(code, config, option) {
      const { processor, ...otherConfig } = config || {};
      const newConfig = convertConfigToRc(otherConfig, this);
      const newOption = convertOptionToLegacy(processor, option, config || {});
      return super.verify(code, newConfig, newOption);
    }
  };
}

export { getLinter };
