"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineWrapperListener = defineWrapperListener;
exports.getProxyNode = getProxyNode;
exports.buildProxyListener = buildProxyListener;
exports.getCoreRule = getCoreRule;
const eslint_1 = require("eslint");
/**
 * Define the wrapped core rule.
 */
function defineWrapperListener(coreRule, context, proxyOptions) {
    const listener = coreRule.create(context);
    const svelteListener = proxyOptions.createListenerProxy?.(listener) ?? listener;
    return svelteListener;
}
/**
 * Get the proxy node
 */
function getProxyNode(node, properties) {
    const cache = {};
    return new Proxy(node, {
        get(_t, key) {
            if (key in cache) {
                return cache[key];
            }
            if (key in properties) {
                return (cache[key] = properties[key]);
            }
            return node[key];
        }
    });
}
/**
 * Build the proxy rule listener
 */
function buildProxyListener(base, convertNode) {
    const listeners = {};
    for (const [key, listener] of Object.entries(base)) {
        listeners[key] = function (...args) {
            listener.call(this, ...args.map((arg) => {
                if (typeof arg === 'object' &&
                    'type' in arg &&
                    typeof arg.type === 'string' &&
                    'range' in arg) {
                    return convertNode(arg);
                }
                return arg;
            }));
        };
    }
    return listeners;
}
let ruleMap = null;
/**
 * Get the core rule implementation from the rule name
 */
function getCoreRule(ruleName) {
    try {
        const map = ruleMap
            ? ruleMap
            : (ruleMap = new eslint_1.Linter().getRules());
        return map.get(ruleName);
    }
    catch {
        // getRules() is no longer available in flat config.
    }
    // eslint-disable-next-line @typescript-eslint/no-require-imports , @typescript-eslint/no-var-requires -- Ignore
    const { builtinRules } = require('eslint/use-at-your-own-risk');
    ruleMap = builtinRules;
    return builtinRules.get(ruleName) || null;
}
