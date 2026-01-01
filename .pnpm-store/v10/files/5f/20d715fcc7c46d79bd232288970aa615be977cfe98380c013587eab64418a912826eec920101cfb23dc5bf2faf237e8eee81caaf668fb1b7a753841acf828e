import type { Linter } from 'eslint';
declare module 'eslint' {
    namespace Linter {
        interface RulesRecord extends RuleOptions {
        }
    }
}
export interface RuleOptions {
    /**
     * disallow conditionals where the type is always truthy or always falsy
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/@typescript-eslint/no-unnecessary-condition/
     * @deprecated
     */
    'svelte/@typescript-eslint/no-unnecessary-condition'?: Linter.RuleEntry<SvelteTypescriptEslintNoUnnecessaryCondition>;
    /**
     * disallows the use of languages other than those specified in the configuration for the lang attribute of `<script>` and `<style>` blocks.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/block-lang/
     */
    'svelte/block-lang'?: Linter.RuleEntry<SvelteBlockLang>;
    /**
     * disallow usage of button without an explicit type attribute
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/button-has-type/
     */
    'svelte/button-has-type'?: Linter.RuleEntry<SvelteButtonHasType>;
    /**
     * support comment-directives in HTML template
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/comment-directive/
     */
    'svelte/comment-directive'?: Linter.RuleEntry<SvelteCommentDirective>;
    /**
     * derived store should use same variable names between values and callback
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/derived-has-same-inputs-outputs/
     */
    'svelte/derived-has-same-inputs-outputs'?: Linter.RuleEntry<[]>;
    /**
     * require slot type declaration using the `$$Slots` interface
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/experimental-require-slot-types/
     */
    'svelte/experimental-require-slot-types'?: Linter.RuleEntry<[]>;
    /**
     * require the strictEvents attribute on `<script>` tags
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/experimental-require-strict-events/
     */
    'svelte/experimental-require-strict-events'?: Linter.RuleEntry<[]>;
    /**
     * enforce the location of first attribute
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/first-attribute-linebreak/
     */
    'svelte/first-attribute-linebreak'?: Linter.RuleEntry<SvelteFirstAttributeLinebreak>;
    /**
     * Require or disallow a line break before tag's closing brackets
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/html-closing-bracket-new-line/
     */
    'svelte/html-closing-bracket-new-line'?: Linter.RuleEntry<SvelteHtmlClosingBracketNewLine>;
    /**
     * require or disallow a space before tag's closing brackets
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/html-closing-bracket-spacing/
     */
    'svelte/html-closing-bracket-spacing'?: Linter.RuleEntry<SvelteHtmlClosingBracketSpacing>;
    /**
     * enforce quotes style of HTML attributes
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/html-quotes/
     */
    'svelte/html-quotes'?: Linter.RuleEntry<SvelteHtmlQuotes>;
    /**
     * enforce self-closing style
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/html-self-closing/
     */
    'svelte/html-self-closing'?: Linter.RuleEntry<SvelteHtmlSelfClosing>;
    /**
     * enforce consistent indentation
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/indent/
     */
    'svelte/indent'?: Linter.RuleEntry<SvelteIndent>;
    /**
     * Svelte runtime prevents calling the same reactive statement twice in a microtask. But between different microtask, it doesn't prevent.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/infinite-reactive-loop/
     */
    'svelte/infinite-reactive-loop'?: Linter.RuleEntry<[]>;
    /**
     * enforce the maximum number of attributes per line
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/max-attributes-per-line/
     */
    'svelte/max-attributes-per-line'?: Linter.RuleEntry<SvelteMaxAttributesPerLine>;
    /**
     * enforce unified spacing in mustache
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/mustache-spacing/
     */
    'svelte/mustache-spacing'?: Linter.RuleEntry<SvelteMustacheSpacing>;
    /**
     * disallow the use of `{@debug}`
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-at-debug-tags/
     */
    'svelte/no-at-debug-tags'?: Linter.RuleEntry<[]>;
    /**
     * disallow use of `{@html}` to prevent XSS attack
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-at-html-tags/
     */
    'svelte/no-at-html-tags'?: Linter.RuleEntry<[]>;
    /**
     * disallow DOM manipulating
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dom-manipulating/
     */
    'svelte/no-dom-manipulating'?: Linter.RuleEntry<[]>;
    /**
     * disallow duplicate conditions in `{#if}` / `{:else if}` chains
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dupe-else-if-blocks/
     */
    'svelte/no-dupe-else-if-blocks'?: Linter.RuleEntry<[]>;
    /**
     * disallow duplicate `on:` directives
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dupe-on-directives/
     */
    'svelte/no-dupe-on-directives'?: Linter.RuleEntry<[]>;
    /**
     * disallow duplicate style properties
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dupe-style-properties/
     */
    'svelte/no-dupe-style-properties'?: Linter.RuleEntry<[]>;
    /**
     * disallow duplicate `use:` directives
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dupe-use-directives/
     */
    'svelte/no-dupe-use-directives'?: Linter.RuleEntry<[]>;
    /**
     * disallow dynamic slot name
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-dynamic-slot-name/
     */
    'svelte/no-dynamic-slot-name'?: Linter.RuleEntry<[]>;
    /**
     * disallow exporting load functions in `*.svelte` module in SvelteKit page components.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-export-load-in-svelte-module-in-kit-pages/
     */
    'svelte/no-export-load-in-svelte-module-in-kit-pages'?: Linter.RuleEntry<[]>;
    /**
     * disallow wrapping single reactive statements in curly braces
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-extra-reactive-curlies/
     */
    'svelte/no-extra-reactive-curlies'?: Linter.RuleEntry<[]>;
    /**
     * disallow using goto() without the base path
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-goto-without-base/
     */
    'svelte/no-goto-without-base'?: Linter.RuleEntry<[]>;
    /**
     * disallow ignoring the unsubscribe method returned by the `subscribe()` on Svelte stores.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-ignored-unsubscribe/
     */
    'svelte/no-ignored-unsubscribe'?: Linter.RuleEntry<[]>;
    /**
     * disallow reactive statements that don't reference reactive values.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-immutable-reactive-statements/
     */
    'svelte/no-immutable-reactive-statements'?: Linter.RuleEntry<[]>;
    /**
     * disallow attributes and directives that produce inline styles
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-inline-styles/
     */
    'svelte/no-inline-styles'?: Linter.RuleEntry<SvelteNoInlineStyles>;
    /**
     * disallow variable or `function` declarations in nested blocks
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-inner-declarations/
     */
    'svelte/no-inner-declarations'?: Linter.RuleEntry<SvelteNoInnerDeclarations>;
    /**
     * Warns against the use of `$inspect` directive
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-inspect/
     */
    'svelte/no-inspect'?: Linter.RuleEntry<[]>;
    /**
     * disallow use of not function in event handler
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-not-function-handler/
     */
    'svelte/no-not-function-handler'?: Linter.RuleEntry<[]>;
    /**
     * disallow objects in text mustache interpolation
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-object-in-text-mustaches/
     */
    'svelte/no-object-in-text-mustaches'?: Linter.RuleEntry<[]>;
    /**
     * it's not necessary to define functions in reactive statements
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-reactive-functions/
     */
    'svelte/no-reactive-functions'?: Linter.RuleEntry<[]>;
    /**
     * don't assign literal values in reactive statements
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-reactive-literals/
     */
    'svelte/no-reactive-literals'?: Linter.RuleEntry<[]>;
    /**
     * disallow reassigning reactive values
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-reactive-reassign/
     */
    'svelte/no-reactive-reassign'?: Linter.RuleEntry<SvelteNoReactiveReassign>;
    /**
     * disallow specific HTML elements
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-restricted-html-elements/
     */
    'svelte/no-restricted-html-elements'?: Linter.RuleEntry<SvelteNoRestrictedHtmlElements>;
    /**
     * disallow shorthand style properties that override related longhand properties
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-shorthand-style-property-overrides/
     */
    'svelte/no-shorthand-style-property-overrides'?: Linter.RuleEntry<[]>;
    /**
     * disallow spaces around equal signs in attribute
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-spaces-around-equal-signs-in-attribute/
     */
    'svelte/no-spaces-around-equal-signs-in-attribute'?: Linter.RuleEntry<[]>;
    /**
     * disallow using async/await inside svelte stores because it causes issues with the auto-unsubscribing features
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-store-async/
     */
    'svelte/no-store-async'?: Linter.RuleEntry<[]>;
    /**
     * svelte/internal will be removed in Svelte 6.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-svelte-internal/
     */
    'svelte/no-svelte-internal'?: Linter.RuleEntry<[]>;
    /**
     * disallow `target="_blank"` attribute without `rel="noopener noreferrer"`
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-target-blank/
     */
    'svelte/no-target-blank'?: Linter.RuleEntry<SvelteNoTargetBlank>;
    /**
     * disallow trailing whitespace at the end of lines
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-trailing-spaces/
     */
    'svelte/no-trailing-spaces'?: Linter.RuleEntry<SvelteNoTrailingSpaces>;
    /**
     * disallow unknown `style:property`
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-unknown-style-directive-property/
     */
    'svelte/no-unknown-style-directive-property'?: Linter.RuleEntry<SvelteNoUnknownStyleDirectiveProperty>;
    /**
     * disallow the use of a class in the template without a corresponding style
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-unused-class-name/
     */
    'svelte/no-unused-class-name'?: Linter.RuleEntry<SvelteNoUnusedClassName>;
    /**
     * disallow unused svelte-ignore comments
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-unused-svelte-ignore/
     */
    'svelte/no-unused-svelte-ignore'?: Linter.RuleEntry<[]>;
    /**
     * disallow unnecessary mustache interpolations
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/no-useless-mustaches/
     */
    'svelte/no-useless-mustaches'?: Linter.RuleEntry<SvelteNoUselessMustaches>;
    /**
     * require class directives instead of ternary expressions
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/prefer-class-directive/
     */
    'svelte/prefer-class-directive'?: Linter.RuleEntry<SveltePreferClassDirective>;
    /**
     * destructure values from object stores for better change tracking & fewer redraws
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/prefer-destructured-store-props/
     */
    'svelte/prefer-destructured-store-props'?: Linter.RuleEntry<[]>;
    /**
     * require style directives instead of style attribute
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/prefer-style-directive/
     */
    'svelte/prefer-style-directive'?: Linter.RuleEntry<[]>;
    /**
     * require keyed `{#each}` block
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-each-key/
     */
    'svelte/require-each-key'?: Linter.RuleEntry<[]>;
    /**
     * require type parameters for `createEventDispatcher`
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-event-dispatcher-types/
     */
    'svelte/require-event-dispatcher-types'?: Linter.RuleEntry<[]>;
    /**
     * require style attributes that can be optimized
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-optimized-style-attribute/
     */
    'svelte/require-optimized-style-attribute'?: Linter.RuleEntry<[]>;
    /**
     * store callbacks must use `set` param
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-store-callbacks-use-set-param/
     */
    'svelte/require-store-callbacks-use-set-param'?: Linter.RuleEntry<[]>;
    /**
     * disallow to use of the store itself as an operand. Need to use $ prefix or get function.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-store-reactive-access/
     */
    'svelte/require-store-reactive-access'?: Linter.RuleEntry<[]>;
    /**
     * require initial value in store
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/require-stores-init/
     */
    'svelte/require-stores-init'?: Linter.RuleEntry<[]>;
    /**
     * enforce use of shorthand syntax in attribute
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/shorthand-attribute/
     */
    'svelte/shorthand-attribute'?: Linter.RuleEntry<SvelteShorthandAttribute>;
    /**
     * enforce use of shorthand syntax in directives
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/shorthand-directive/
     */
    'svelte/shorthand-directive'?: Linter.RuleEntry<SvelteShorthandDirective>;
    /**
     * enforce order of attributes
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/sort-attributes/
     */
    'svelte/sort-attributes'?: Linter.RuleEntry<SvelteSortAttributes>;
    /**
     * enforce consistent spacing after the `<!--` and before the `-->` in a HTML comment
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/spaced-html-comment/
     */
    'svelte/spaced-html-comment'?: Linter.RuleEntry<SvelteSpacedHtmlComment>;
    /**
     * system rule for working this plugin
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/system/
     */
    'svelte/system'?: Linter.RuleEntry<[]>;
    /**
     * disallow warnings when compiling.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/valid-compile/
     */
    'svelte/valid-compile'?: Linter.RuleEntry<SvelteValidCompile>;
    /**
     * enforce keys to use variables defined in the `{#each}` block
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/valid-each-key/
     */
    'svelte/valid-each-key'?: Linter.RuleEntry<[]>;
    /**
     * disallow props other than data or errors in SvelteKit page components.
     * @see https://sveltejs.github.io/eslint-plugin-svelte/rules/valid-prop-names-in-kit-pages/
     */
    'svelte/valid-prop-names-in-kit-pages'?: Linter.RuleEntry<[]>;
}
type SvelteTypescriptEslintNoUnnecessaryCondition = [] | [
    {
        allowConstantLoopConditions?: boolean;
        allowRuleToRunWithoutStrictNullChecksIKnowWhatIAmDoing?: boolean;
    }
];
type SvelteBlockLang = [] | [
    {
        enforceScriptPresent?: boolean;
        enforceStylePresent?: boolean;
        script?: ((string | null) | [(string | null), ...((string | null))[]]);
        style?: ((string | null) | [(string | null), ...((string | null))[]]);
    }
];
type SvelteButtonHasType = [] | [
    {
        button?: boolean;
        submit?: boolean;
        reset?: boolean;
    }
];
type SvelteCommentDirective = [] | [
    {
        reportUnusedDisableDirectives?: boolean;
    }
];
type SvelteFirstAttributeLinebreak = [] | [
    {
        multiline?: ("below" | "beside");
        singleline?: ("below" | "beside");
    }
];
type SvelteHtmlClosingBracketNewLine = [] | [
    {
        singleline?: ("always" | "never");
        multiline?: ("always" | "never");
        selfClosingTag?: {
            singleline?: ("always" | "never");
            multiline?: ("always" | "never");
        };
    }
];
type SvelteHtmlClosingBracketSpacing = [] | [
    {
        startTag?: ("always" | "never" | "ignore");
        endTag?: ("always" | "never" | "ignore");
        selfClosingTag?: ("always" | "never" | "ignore");
    }
];
type SvelteHtmlQuotes = [] | [
    {
        prefer?: ("double" | "single");
        dynamic?: {
            quoted?: boolean;
            avoidInvalidUnquotedInHTML?: boolean;
        };
    }
];
type SvelteHtmlSelfClosing = [] | [
    ({
        void?: ("never" | "always" | "ignore");
        normal?: ("never" | "always" | "ignore");
        foreign?: ("never" | "always" | "ignore");
        component?: ("never" | "always" | "ignore");
        svelte?: ("never" | "always" | "ignore");
    } | ("all" | "html" | "none"))
];
type SvelteIndent = [] | [
    {
        indent?: (number | "tab");
        indentScript?: boolean;
        switchCase?: number;
        alignAttributesVertically?: boolean;
        ignoredNodes?: (string & {
            [k: string]: unknown | undefined;
        } & {
            [k: string]: unknown | undefined;
        })[];
    }
];
type SvelteMaxAttributesPerLine = [] | [
    {
        multiline?: number;
        singleline?: number;
    }
];
type SvelteMustacheSpacing = [] | [
    {
        textExpressions?: ("never" | "always");
        attributesAndProps?: ("never" | "always");
        directiveExpressions?: ("never" | "always");
        tags?: {
            openingBrace?: ("never" | "always");
            closingBrace?: ("never" | "always" | "always-after-expression");
        };
    }
];
type SvelteNoInlineStyles = [] | [
    {
        allowTransitions?: boolean;
    }
];
type SvelteNoInnerDeclarations = [] | [("functions" | "both")] | [
    ("functions" | "both"),
    {
        blockScopedFunctions?: ("allow" | "disallow");
    }
];
type SvelteNoReactiveReassign = [] | [
    {
        props?: boolean;
    }
];
type SvelteNoRestrictedHtmlElements = [
    (string | {
        elements?: [string, ...(string)[]];
        message?: string;
    }),
    ...((string | {
        elements?: [string, ...(string)[]];
        message?: string;
    }))[]
];
type SvelteNoTargetBlank = [] | [
    {
        allowReferrer?: boolean;
        enforceDynamicLinks?: ("always" | "never");
    }
];
type SvelteNoTrailingSpaces = [] | [
    {
        skipBlankLines?: boolean;
        ignoreComments?: boolean;
    }
];
type SvelteNoUnknownStyleDirectiveProperty = [] | [
    {
        ignoreProperties?: [string, ...(string)[]];
        ignorePrefixed?: boolean;
    }
];
type SvelteNoUnusedClassName = [] | [
    {
        allowedClassNames?: string[];
    }
];
type SvelteNoUselessMustaches = [] | [
    {
        ignoreIncludesComment?: boolean;
        ignoreStringEscape?: boolean;
    }
];
type SveltePreferClassDirective = [] | [
    {
        prefer?: ("always" | "empty");
    }
];
type SvelteShorthandAttribute = [] | [
    {
        prefer?: ("always" | "never");
    }
];
type SvelteShorthandDirective = [] | [
    {
        prefer?: ("always" | "never");
    }
];
type SvelteSortAttributes = [] | [
    {
        order?: (string | [string, ...(string)[]] | {
            match: (string | [string, ...(string)[]]);
            sort: ("alphabetical" | "ignore");
        })[];
        alphabetical?: boolean;
    }
];
type SvelteSpacedHtmlComment = [] | [("always" | "never")];
type SvelteValidCompile = [] | [
    {
        ignoreWarnings?: boolean;
    }
];
export {};
