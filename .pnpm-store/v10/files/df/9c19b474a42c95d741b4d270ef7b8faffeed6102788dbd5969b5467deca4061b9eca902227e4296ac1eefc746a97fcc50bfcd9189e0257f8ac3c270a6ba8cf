'use strict';

const eslint = require('eslint');
const semver = require('semver');
const convertConfig = require('./shared/eslint-compat-utils.503aeaae.cjs');
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

let cacheLinter;
function getLinter() {
  return cacheLinter != null ? cacheLinter : cacheLinter = getLinterInternal();
  function getLinterInternal() {
    if (semver__namespace.gte(eslint__namespace.Linter.version, "9.0.0-0")) {
      return eslint__namespace.Linter;
    }
    return getLinterClassFromLegacyLinter();
  }
}
function getLinterClassFromLegacyLinter() {
  return class LinterFromLegacyLinter extends eslint__namespace.Linter {
    static get version() {
      return eslint__namespace.Linter.version;
    }
    verify(code, config, option) {
      const { processor, ...otherConfig } = config || {};
      const newConfig = convertConfig.convertConfigToRc(otherConfig, this);
      const newOption = convertOption.convertOptionToLegacy(processor, option, config || {});
      return super.verify(code, newConfig, newOption);
    }
  };
}

exports.getLinter = getLinter;
