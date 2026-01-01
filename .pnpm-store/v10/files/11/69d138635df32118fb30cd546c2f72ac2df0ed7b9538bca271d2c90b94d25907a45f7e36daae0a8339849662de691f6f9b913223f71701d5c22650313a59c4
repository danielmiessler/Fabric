'use strict';

const path = require('path');
const fs = require('fs');

function applyPolyfills(object, polyfill) {
  return new Proxy(object, {
    get(_target, p) {
      var _a;
      return (_a = object[p]) != null ? _a : polyfill[p];
    }
  });
}

function getParent(node) {
  return node.parent;
}

const cache = /* @__PURE__ */ new WeakMap();
function getSourceCode(context) {
  const original = context.sourceCode || context.getSourceCode();
  const cached = cache.get(original);
  if (cached) {
    return cached;
  }
  const sourceCode = applyPolyfills(original, {
    getScope(node) {
      const inner = node.type !== "Program";
      for (let n = node; n; n = getParent(n)) {
        const scope = original.scopeManager.acquire(n, inner);
        if (scope) {
          if (scope.type === "function-expression-name") {
            return scope.childScopes[0];
          }
          return scope;
        }
      }
      return original.scopeManager.scopes[0];
    },
    markVariableAsUsed(name, refNode = original.ast) {
      const currentScope = sourceCode.getScope(refNode);
      if (currentScope === context.getScope()) {
        return context.markVariableAsUsed(name);
      }
      let initialScope = currentScope;
      if (currentScope.type === "global" && currentScope.childScopes.length > 0 && currentScope.childScopes[0].block === original.ast) {
        initialScope = currentScope.childScopes[0];
      }
      for (let scope = initialScope; scope; scope = scope.upper) {
        const variable = scope.variables.find(
          (scopeVar) => scopeVar.name === name
        );
        if (variable) {
          variable.eslintUsed = true;
          return true;
        }
      }
      return false;
    },
    getAncestors(node) {
      const result = [];
      for (let ancestor = getParent(node); ancestor; ancestor = ancestor.parent) {
        result.unshift(ancestor);
      }
      return result;
    },
    getDeclaredVariables(node) {
      return original.scopeManager.getDeclaredVariables(node);
    },
    isSpaceBetween(first, second) {
      if (first.range[0] <= second.range[1] && second.range[0] <= first.range[1]) {
        return false;
      }
      const [startingNodeOrToken, endingNodeOrToken] = first.range[1] <= second.range[0] ? [first, second] : [second, first];
      const tokens = sourceCode.getTokensBetween(first, second, {
        includeComments: true
      });
      let startIndex = startingNodeOrToken.range[1];
      for (const token of tokens) {
        if (startIndex !== token.range[0]) {
          return true;
        }
        startIndex = token.range[1];
      }
      return startIndex !== endingNodeOrToken.range[0];
    }
  });
  cache.set(original, sourceCode);
  return sourceCode;
}

function getCwd(context) {
  var _a, _b, _c;
  return (_c = (_b = context.cwd) != null ? _b : (_a = context.getCwd) == null ? void 0 : _a.call(context)) != null ? _c : (
    // getCwd is added in v6.6.0
    process.cwd()
  );
}

function getFilename(context) {
  var _a;
  return (_a = context.filename) != null ? _a : context.getFilename();
}

function getPhysicalFilename(context) {
  var _a, _b;
  const physicalFilename = (_b = context.physicalFilename) != null ? _b : (_a = context.getPhysicalFilename) == null ? void 0 : _a.call(context);
  if (physicalFilename != null) {
    return physicalFilename;
  }
  const filename = getFilename(context);
  let target = filename;
  while (/^\d+_/u.test(path.basename(target)) && !fs.existsSync(target)) {
    const next = path.dirname(target);
    if (next === target || !path.extname(next)) {
      break;
    }
    target = next;
  }
  return target;
}

exports.getCwd = getCwd;
exports.getFilename = getFilename;
exports.getPhysicalFilename = getPhysicalFilename;
exports.getSourceCode = getSourceCode;
