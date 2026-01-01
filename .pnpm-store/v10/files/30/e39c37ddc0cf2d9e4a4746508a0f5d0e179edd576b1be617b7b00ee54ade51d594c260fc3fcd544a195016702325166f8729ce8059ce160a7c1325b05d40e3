"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
/**
 * Splits the given node by the given logical operator.
 * @param operator Logical operator `||` or `&&`.
 * @param node The node to split.
 * @returns Array of conditions that makes the node when joined by the operator.
 */
function splitByLogicalOperator(operator, node) {
    if (node.type === 'LogicalExpression' && node.operator === operator) {
        return [
            ...splitByLogicalOperator(operator, node.left),
            ...splitByLogicalOperator(operator, node.right)
        ];
    }
    return [node];
}
/**
 * Split with ||.
 */
function splitByOr(node) {
    return splitByLogicalOperator('||', node);
}
/**
 * Split with &&.
 */
function splitByAnd(node) {
    return splitByLogicalOperator('&&', node);
}
/**
 * Build OrOperands
 */
function buildOrOperands(node) {
    const orOperands = splitByOr(node);
    return {
        node,
        operands: orOperands.map((orOperand) => {
            const andOperands = splitByAnd(orOperand);
            return {
                node: orOperand,
                operands: andOperands
            };
        })
    };
}
exports.default = (0, utils_1.createRule)('no-dupe-else-if-blocks', {
    meta: {
        docs: {
            description: 'disallow duplicate conditions in `{#if}` / `{:else if}` chains',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [],
        messages: {
            unexpected: 'This branch can never execute. Its condition is a duplicate or covered by previous conditions in the `{#if}` / `{:else if}` chain.'
        },
        type: 'problem' // "problem",
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        /**
         * Determines whether the two given nodes are considered to be equal. In particular, given that the nodes
         * represent expressions in a boolean context, `||` and `&&` can be considered as commutative operators.
         * @param a First node.
         * @param b Second node.
         * @returns `true` if the nodes are considered to be equal.
         */
        function equal(a, b) {
            if (a.type !== b.type) {
                return false;
            }
            if (a.type === 'LogicalExpression' &&
                b.type === 'LogicalExpression' &&
                (a.operator === '||' || a.operator === '&&') &&
                a.operator === b.operator) {
                return ((equal(a.left, b.left) && equal(a.right, b.right)) ||
                    (equal(a.left, b.right) && equal(a.right, b.left)));
            }
            return (0, ast_utils_1.equalTokens)(a, b, sourceCode);
        }
        /**
         * Determines whether the first given AndOperands is a subset of the second given AndOperands.
         *
         * e.g. A: (a && b), B: (a && b && c): B is a subset of A.
         *
         * @param operandsA The AndOperands to compare from.
         * @param operandsB The AndOperands to compare against.
         * @returns `true` if the `andOperandsA` is a subset of the `andOperandsB`.
         */
        function isSubset(operandsA, operandsB) {
            return operandsA.operands.every((operandA) => operandsB.operands.some((operandB) => equal(operandA, operandB)));
        }
        /** Iterate SvelteIfBlock nodes */
        function* iterateIfElseIf(node) {
            let target = node;
            while (target.parent.type === 'SvelteElseBlock' &&
                target.parent.children.includes(target) &&
                target.parent.parent.type === 'SvelteIfBlock') {
                yield target.parent.parent;
                target = target.parent.parent;
            }
        }
        return {
            SvelteIfBlock(node) {
                const test = node.expression;
                const conditionsToCheck = test.type === 'LogicalExpression' && test.operator === '&&'
                    ? [...splitByAnd(test), test]
                    : [test];
                const listToCheck = conditionsToCheck.map(buildOrOperands);
                for (const currentIdBlock of iterateIfElseIf(node)) {
                    if (currentIdBlock.expression) {
                        const currentOrOperands = buildOrOperands(currentIdBlock.expression);
                        for (const condition of listToCheck) {
                            const operands = (condition.operands = condition.operands.filter((orOperand) => {
                                return !currentOrOperands.operands.some((currentOrOperand) => isSubset(currentOrOperand, orOperand));
                            }));
                            if (!operands.length) {
                                context.report({
                                    node: condition.node,
                                    messageId: 'unexpected'
                                });
                                return;
                            }
                        }
                    }
                }
            }
        };
    }
});
