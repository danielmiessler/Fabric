/**
 * The code in this file is copied from https://github.com/lukeed/clsx and modified to suit the needs of tailwind-merge better.
 *
 * Specifically:
 * - Runtime code from https://github.com/lukeed/clsx/blob/v1.2.1/src/index.js
 * - TypeScript types from https://github.com/lukeed/clsx/blob/v1.2.1/clsx.d.ts
 *
 * Original code has MIT license: Copyright (c) Luke Edwards <luke.edwards05@gmail.com> (lukeed.com)
 */
type ClassNameValue = ClassNameArray | string | null | undefined | 0 | 0n | false;
type ClassNameArray = ClassNameValue[];
declare function twJoin(...classLists: ClassNameValue[]): string;

/**
 * Type the tailwind-merge configuration adheres to.
 */
interface Config<ClassGroupIds extends string, ThemeGroupIds extends string> extends ConfigStaticPart, ConfigGroupsPart<ClassGroupIds, ThemeGroupIds> {
}
/**
 * The static part of the tailwind-merge configuration. When merging multiple configurations, the properties of this interface are always overridden.
 */
interface ConfigStaticPart {
    /**
     * Integer indicating size of LRU cache used for memoizing results.
     * - Cache might be up to twice as big as `cacheSize`
     * - No cache is used for values <= 0
     */
    cacheSize: number;
    /**
     * Prefix added to Tailwind-generated classes
     * @see https://tailwindcss.com/docs/configuration#prefix
     */
    prefix?: string;
    /**
     * Custom separator for modifiers in Tailwind classes
     * @see https://tailwindcss.com/docs/configuration#separator
     */
    separator: string;
    /**
     * Allows to customize parsing of individual classes passed to `twMerge`.
     * All classes passed to `twMerge` outside of cache hits are passed to this function before it is determined whether the class is a valid Tailwind CSS class.
     *
     * This is an experimental feature and may introduce breaking changes in any minor version update.
     */
    experimentalParseClassName?(param: ExperimentalParseClassNameParam): ExperimentalParsedClassName;
}
/**
 * Type of param passed to the `experimentalParseClassName` function.
 *
 * This is an experimental feature and may introduce breaking changes in any minor version update.
 */
interface ExperimentalParseClassNameParam {
    className: string;
    parseClassName(className: string): ExperimentalParsedClassName;
}
/**
 * Type of the result returned by the `experimentalParseClassName` function.
 *
 * This is an experimental feature and may introduce breaking changes in any minor version update.
 */
interface ExperimentalParsedClassName {
    /**
     * Modifiers of the class in the order they appear in the class.
     *
     * @example ['hover', 'dark'] // for `hover:dark:bg-gray-100`
     */
    modifiers: string[];
    /**
     * Whether the class has an `!important` modifier.
     *
     * @example true // for `hover:dark:!bg-gray-100`
     */
    hasImportantModifier: boolean;
    /**
     * Base class without preceding modifiers.
     *
     * @example 'bg-gray-100' // for `hover:dark:bg-gray-100`
     */
    baseClassName: string;
    /**
     * Index position of a possible postfix modifier in the class.
     * If the class has no postfix modifier, this is `undefined`.
     *
     * This property is prefixed with "maybe" because tailwind-merge does not know whether something is a postfix modifier or part of the base class since it's possible to configure Tailwind CSS classes which include a `/` in the base class name.
     *
     * If a `maybePostfixModifierPosition` is present, tailwind-merge first tries to match the `baseClassName` without the possible postfix modifier to a class group. If tht fails, it tries again with the possible postfix modifier.
     *
     * @example 11 // for `bg-gray-100/50`
     */
    maybePostfixModifierPosition: number | undefined;
}
/**
 * The dynamic part of the tailwind-merge configuration. When merging multiple configurations, the user can choose to either override or extend the properties of this interface.
 */
interface ConfigGroupsPart<ClassGroupIds extends string, ThemeGroupIds extends string> {
    /**
     * Theme scales used in classGroups.
     * The keys are the same as in the Tailwind config but the values are sometimes defined more broadly.
     */
    theme: NoInfer<ThemeObject<ThemeGroupIds>>;
    /**
     * Object with groups of classes.
     * @example
     * {
     *     // Creates group of classes `group`, `of` and `classes`
     *     'group-id': ['group', 'of', 'classes'],
     *     // Creates group of classes `look-at-me-other` and `look-at-me-group`.
     *     'other-group': [{ 'look-at-me': ['other', 'group']}]
     * }
     */
    classGroups: NoInfer<Record<ClassGroupIds, ClassGroup<ThemeGroupIds>>>;
    /**
     * Conflicting classes across groups.
     * The key is ID of class group which creates conflict, values are IDs of class groups which receive a conflict.
     * A class group ID is the key of a class group in classGroups object.
     * @example { gap: ['gap-x', 'gap-y'] }
     */
    conflictingClassGroups: NoInfer<Partial<Record<ClassGroupIds, readonly ClassGroupIds[]>>>;
    /**
     * Postfix modifiers conflicting with other class groups.
     * A class group ID is the key of a class group in classGroups object.
     * @example { 'font-size': ['leading'] }
     */
    conflictingClassGroupModifiers: NoInfer<Partial<Record<ClassGroupIds, readonly ClassGroupIds[]>>>;
}
/**
 * Type of the configuration object that can be passed to `extendTailwindMerge`.
 */
interface ConfigExtension<ClassGroupIds extends string, ThemeGroupIds extends string> extends Partial<ConfigStaticPart> {
    override?: PartialPartial<ConfigGroupsPart<ClassGroupIds, ThemeGroupIds>>;
    extend?: PartialPartial<ConfigGroupsPart<ClassGroupIds, ThemeGroupIds>>;
}
type PartialPartial<T> = {
    [P in keyof T]?: Partial<T[P]>;
};
type ThemeObject<ThemeGroupIds extends string> = Record<ThemeGroupIds, ClassGroup<ThemeGroupIds>>;
type ClassGroup<ThemeGroupIds extends string> = readonly ClassDefinition<ThemeGroupIds>[];
type ClassDefinition<ThemeGroupIds extends string> = string | ClassValidator | ThemeGetter | ClassObject<ThemeGroupIds>;
type ClassValidator = (classPart: string) => boolean;
interface ThemeGetter {
    (theme: ThemeObject<AnyThemeGroupIds>): ClassGroup<AnyClassGroupIds>;
    isThemeGetter: true;
}
type ClassObject<ThemeGroupIds extends string> = Record<string, readonly ClassDefinition<ThemeGroupIds>[]>;
/**
 * Hack from https://stackoverflow.com/questions/56687668/a-way-to-disable-type-argument-inference-in-generics/56688073#56688073
 *
 * Could be replaced with NoInfer utility type from TypeScript (https://www.typescriptlang.org/docs/handbook/utility-types.html#noinfertype), but that is only supported in TypeScript 5.4 or higher, so I should wait some time before using it.
 */
type NoInfer<T> = [T][T extends any ? 0 : never];
/**
 * Theme group IDs included in the default configuration of tailwind-merge.
 *
 * If you want to use a scale that is not supported in the `ThemeObject` type,
 * consider using `classGroups` instead of `theme`.
 *
 * @see https://github.com/dcastil/tailwind-merge/blob/main/docs/configuration.md#theme
 *      (the list of supported keys may vary between `tailwind-merge` versions)
 */
type DefaultThemeGroupIds = 'blur' | 'borderColor' | 'borderRadius' | 'borderSpacing' | 'borderWidth' | 'brightness' | 'colors' | 'contrast' | 'gap' | 'gradientColorStopPositions' | 'gradientColorStops' | 'grayscale' | 'hueRotate' | 'inset' | 'invert' | 'margin' | 'opacity' | 'padding' | 'saturate' | 'scale' | 'sepia' | 'skew' | 'space' | 'spacing' | 'translate';
/**
 * Class group IDs included in the default configuration of tailwind-merge.
 */
type DefaultClassGroupIds = 'accent' | 'align-content' | 'align-items' | 'align-self' | 'animate' | 'appearance' | 'aspect' | 'auto-cols' | 'auto-rows' | 'backdrop-blur' | 'backdrop-brightness' | 'backdrop-contrast' | 'backdrop-filter' | 'backdrop-grayscale' | 'backdrop-hue-rotate' | 'backdrop-invert' | 'backdrop-opacity' | 'backdrop-saturate' | 'backdrop-sepia' | 'basis' | 'bg-attachment' | 'bg-blend' | 'bg-clip' | 'bg-color' | 'bg-image' | 'bg-opacity' | 'bg-origin' | 'bg-position' | 'bg-repeat' | 'bg-size' | 'blur' | 'border-collapse' | 'border-color-b' | 'border-color-e' | 'border-color-l' | 'border-color-r' | 'border-color-s' | 'border-color-t' | 'border-color-x' | 'border-color-y' | 'border-color' | 'border-opacity' | 'border-spacing-x' | 'border-spacing-y' | 'border-spacing' | 'border-style' | 'border-w-b' | 'border-w-e' | 'border-w-l' | 'border-w-r' | 'border-w-s' | 'border-w-t' | 'border-w-x' | 'border-w-y' | 'border-w' | 'bottom' | 'box-decoration' | 'box' | 'break-after' | 'break-before' | 'break-inside' | 'break' | 'brightness' | 'caption' | 'caret-color' | 'clear' | 'col-end' | 'col-start-end' | 'col-start' | 'columns' | 'container' | 'content' | 'contrast' | 'cursor' | 'delay' | 'display' | 'divide-color' | 'divide-opacity' | 'divide-style' | 'divide-x-reverse' | 'divide-x' | 'divide-y-reverse' | 'divide-y' | 'drop-shadow' | 'duration' | 'ease' | 'end' | 'fill' | 'filter' | 'flex-direction' | 'flex-wrap' | 'flex' | 'float' | 'font-family' | 'font-size' | 'font-smoothing' | 'font-style' | 'font-weight' | 'forced-color-adjust' | 'fvn-figure' | 'fvn-fraction' | 'fvn-normal' | 'fvn-ordinal' | 'fvn-slashed-zero' | 'fvn-spacing' | 'gap-x' | 'gap-y' | 'gap' | 'gradient-from-pos' | 'gradient-from' | 'gradient-to-pos' | 'gradient-to' | 'gradient-via-pos' | 'gradient-via' | 'grayscale' | 'grid-cols' | 'grid-flow' | 'grid-rows' | 'grow' | 'h' | 'hue-rotate' | 'hyphens' | 'indent' | 'inset-x' | 'inset-y' | 'inset' | 'invert' | 'isolation' | 'justify-content' | 'justify-items' | 'justify-self' | 'leading' | 'left' | 'line-clamp' | 'list-image' | 'list-style-position' | 'list-style-type' | 'm' | 'max-h' | 'max-w' | 'mb' | 'me' | 'min-h' | 'min-w' | 'mix-blend' | 'ml' | 'mr' | 'ms' | 'mt' | 'mx' | 'my' | 'object-fit' | 'object-position' | 'opacity' | 'order' | 'outline-color' | 'outline-offset' | 'outline-style' | 'outline-w' | 'overflow-x' | 'overflow-y' | 'overflow' | 'overscroll-x' | 'overscroll-y' | 'overscroll' | 'p' | 'pb' | 'pe' | 'pl' | 'place-content' | 'place-items' | 'place-self' | 'placeholder-color' | 'placeholder-opacity' | 'pointer-events' | 'position' | 'pr' | 'ps' | 'pt' | 'px' | 'py' | 'resize' | 'right' | 'ring-color' | 'ring-offset-color' | 'ring-offset-w' | 'ring-opacity' | 'ring-w-inset' | 'ring-w' | 'rotate' | 'rounded-b' | 'rounded-bl' | 'rounded-br' | 'rounded-e' | 'rounded-ee' | 'rounded-es' | 'rounded-l' | 'rounded-r' | 'rounded-s' | 'rounded-se' | 'rounded-ss' | 'rounded-t' | 'rounded-tl' | 'rounded-tr' | 'rounded' | 'row-end' | 'row-start-end' | 'row-start' | 'saturate' | 'scale-x' | 'scale-y' | 'scale' | 'scroll-behavior' | 'scroll-m' | 'scroll-mb' | 'scroll-me' | 'scroll-ml' | 'scroll-mr' | 'scroll-ms' | 'scroll-mt' | 'scroll-mx' | 'scroll-my' | 'scroll-p' | 'scroll-pb' | 'scroll-pe' | 'scroll-pl' | 'scroll-pr' | 'scroll-ps' | 'scroll-pt' | 'scroll-px' | 'scroll-py' | 'select' | 'sepia' | 'shadow-color' | 'shadow' | 'shrink' | 'size' | 'skew-x' | 'skew-y' | 'snap-align' | 'snap-stop' | 'snap-strictness' | 'snap-type' | 'space-x-reverse' | 'space-x' | 'space-y-reverse' | 'space-y' | 'sr' | 'start' | 'stroke-w' | 'stroke' | 'table-layout' | 'text-alignment' | 'text-color' | 'text-decoration-color' | 'text-decoration-style' | 'text-decoration-thickness' | 'text-decoration' | 'text-opacity' | 'text-overflow' | 'text-transform' | 'text-wrap' | 'top' | 'touch-pz' | 'touch-x' | 'touch-y' | 'touch' | 'tracking' | 'transform-origin' | 'transform' | 'transition' | 'translate-x' | 'translate-y' | 'underline-offset' | 'vertical-align' | 'visibility' | 'w' | 'whitespace' | 'will-change' | 'z';
type AnyClassGroupIds = string;
type AnyThemeGroupIds = string;
/**
 * type of the tailwind-merge configuration that allows for any possible configuration.
 */
type AnyConfig = Config<AnyClassGroupIds, AnyThemeGroupIds>;

type CreateConfigFirst = () => AnyConfig;
type CreateConfigSubsequent$1 = (config: AnyConfig) => AnyConfig;
type TailwindMerge = (...classLists: ClassNameValue[]) => string;
declare function createTailwindMerge(createConfigFirst: CreateConfigFirst, ...createConfigRest: CreateConfigSubsequent$1[]): TailwindMerge;

declare const getDefaultConfig: () => {
    readonly cacheSize: 500;
    readonly separator: ":";
    readonly theme: {
        readonly colors: readonly [() => boolean];
        readonly spacing: readonly [(value: string) => boolean, (value: string) => boolean];
        readonly blur: readonly ["none", "", (value: string) => boolean, (value: string) => boolean];
        readonly brightness: ((value: string) => boolean)[];
        readonly borderColor: readonly [ThemeGetter];
        readonly borderRadius: readonly ["none", "", "full", (value: string) => boolean, (value: string) => boolean];
        readonly borderSpacing: readonly [(value: string) => boolean, ThemeGetter];
        readonly borderWidth: readonly ["", (value: string) => boolean, (value: string) => boolean];
        readonly contrast: ((value: string) => boolean)[];
        readonly grayscale: readonly ["", "0", (value: string) => boolean];
        readonly hueRotate: ((value: string) => boolean)[];
        readonly invert: readonly ["", "0", (value: string) => boolean];
        readonly gap: readonly [(value: string) => boolean, ThemeGetter];
        readonly gradientColorStops: readonly [ThemeGetter];
        readonly gradientColorStopPositions: readonly [(value: string) => boolean, (value: string) => boolean];
        readonly inset: readonly ["auto", (value: string) => boolean, ThemeGetter];
        readonly margin: readonly ["auto", (value: string) => boolean, ThemeGetter];
        readonly opacity: ((value: string) => boolean)[];
        readonly padding: readonly [(value: string) => boolean, ThemeGetter];
        readonly saturate: ((value: string) => boolean)[];
        readonly scale: ((value: string) => boolean)[];
        readonly sepia: readonly ["", "0", (value: string) => boolean];
        readonly skew: ((value: string) => boolean)[];
        readonly space: readonly [(value: string) => boolean, ThemeGetter];
        readonly translate: readonly [(value: string) => boolean, ThemeGetter];
    };
    readonly classGroups: {
        /**
         * Aspect Ratio
         * @see https://tailwindcss.com/docs/aspect-ratio
         */
        readonly aspect: readonly [{
            readonly aspect: readonly ["auto", "square", "video", (value: string) => boolean];
        }];
        /**
         * Container
         * @see https://tailwindcss.com/docs/container
         */
        readonly container: readonly ["container"];
        /**
         * Columns
         * @see https://tailwindcss.com/docs/columns
         */
        readonly columns: readonly [{
            readonly columns: readonly [(value: string) => boolean];
        }];
        /**
         * Break After
         * @see https://tailwindcss.com/docs/break-after
         */
        readonly 'break-after': readonly [{
            readonly 'break-after': readonly ["auto", "avoid", "all", "avoid-page", "page", "left", "right", "column"];
        }];
        /**
         * Break Before
         * @see https://tailwindcss.com/docs/break-before
         */
        readonly 'break-before': readonly [{
            readonly 'break-before': readonly ["auto", "avoid", "all", "avoid-page", "page", "left", "right", "column"];
        }];
        /**
         * Break Inside
         * @see https://tailwindcss.com/docs/break-inside
         */
        readonly 'break-inside': readonly [{
            readonly 'break-inside': readonly ["auto", "avoid", "avoid-page", "avoid-column"];
        }];
        /**
         * Box Decoration Break
         * @see https://tailwindcss.com/docs/box-decoration-break
         */
        readonly 'box-decoration': readonly [{
            readonly 'box-decoration': readonly ["slice", "clone"];
        }];
        /**
         * Box Sizing
         * @see https://tailwindcss.com/docs/box-sizing
         */
        readonly box: readonly [{
            readonly box: readonly ["border", "content"];
        }];
        /**
         * Display
         * @see https://tailwindcss.com/docs/display
         */
        readonly display: readonly ["block", "inline-block", "inline", "flex", "inline-flex", "table", "inline-table", "table-caption", "table-cell", "table-column", "table-column-group", "table-footer-group", "table-header-group", "table-row-group", "table-row", "flow-root", "grid", "inline-grid", "contents", "list-item", "hidden"];
        /**
         * Floats
         * @see https://tailwindcss.com/docs/float
         */
        readonly float: readonly [{
            readonly float: readonly ["right", "left", "none", "start", "end"];
        }];
        /**
         * Clear
         * @see https://tailwindcss.com/docs/clear
         */
        readonly clear: readonly [{
            readonly clear: readonly ["left", "right", "both", "none", "start", "end"];
        }];
        /**
         * Isolation
         * @see https://tailwindcss.com/docs/isolation
         */
        readonly isolation: readonly ["isolate", "isolation-auto"];
        /**
         * Object Fit
         * @see https://tailwindcss.com/docs/object-fit
         */
        readonly 'object-fit': readonly [{
            readonly object: readonly ["contain", "cover", "fill", "none", "scale-down"];
        }];
        /**
         * Object Position
         * @see https://tailwindcss.com/docs/object-position
         */
        readonly 'object-position': readonly [{
            readonly object: readonly ["bottom", "center", "left", "left-bottom", "left-top", "right", "right-bottom", "right-top", "top", (value: string) => boolean];
        }];
        /**
         * Overflow
         * @see https://tailwindcss.com/docs/overflow
         */
        readonly overflow: readonly [{
            readonly overflow: readonly ["auto", "hidden", "clip", "visible", "scroll"];
        }];
        /**
         * Overflow X
         * @see https://tailwindcss.com/docs/overflow
         */
        readonly 'overflow-x': readonly [{
            readonly 'overflow-x': readonly ["auto", "hidden", "clip", "visible", "scroll"];
        }];
        /**
         * Overflow Y
         * @see https://tailwindcss.com/docs/overflow
         */
        readonly 'overflow-y': readonly [{
            readonly 'overflow-y': readonly ["auto", "hidden", "clip", "visible", "scroll"];
        }];
        /**
         * Overscroll Behavior
         * @see https://tailwindcss.com/docs/overscroll-behavior
         */
        readonly overscroll: readonly [{
            readonly overscroll: readonly ["auto", "contain", "none"];
        }];
        /**
         * Overscroll Behavior X
         * @see https://tailwindcss.com/docs/overscroll-behavior
         */
        readonly 'overscroll-x': readonly [{
            readonly 'overscroll-x': readonly ["auto", "contain", "none"];
        }];
        /**
         * Overscroll Behavior Y
         * @see https://tailwindcss.com/docs/overscroll-behavior
         */
        readonly 'overscroll-y': readonly [{
            readonly 'overscroll-y': readonly ["auto", "contain", "none"];
        }];
        /**
         * Position
         * @see https://tailwindcss.com/docs/position
         */
        readonly position: readonly ["static", "fixed", "absolute", "relative", "sticky"];
        /**
         * Top / Right / Bottom / Left
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly inset: readonly [{
            readonly inset: readonly [ThemeGetter];
        }];
        /**
         * Right / Left
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly 'inset-x': readonly [{
            readonly 'inset-x': readonly [ThemeGetter];
        }];
        /**
         * Top / Bottom
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly 'inset-y': readonly [{
            readonly 'inset-y': readonly [ThemeGetter];
        }];
        /**
         * Start
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly start: readonly [{
            readonly start: readonly [ThemeGetter];
        }];
        /**
         * End
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly end: readonly [{
            readonly end: readonly [ThemeGetter];
        }];
        /**
         * Top
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly top: readonly [{
            readonly top: readonly [ThemeGetter];
        }];
        /**
         * Right
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly right: readonly [{
            readonly right: readonly [ThemeGetter];
        }];
        /**
         * Bottom
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly bottom: readonly [{
            readonly bottom: readonly [ThemeGetter];
        }];
        /**
         * Left
         * @see https://tailwindcss.com/docs/top-right-bottom-left
         */
        readonly left: readonly [{
            readonly left: readonly [ThemeGetter];
        }];
        /**
         * Visibility
         * @see https://tailwindcss.com/docs/visibility
         */
        readonly visibility: readonly ["visible", "invisible", "collapse"];
        /**
         * Z-Index
         * @see https://tailwindcss.com/docs/z-index
         */
        readonly z: readonly [{
            readonly z: readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Flex Basis
         * @see https://tailwindcss.com/docs/flex-basis
         */
        readonly basis: readonly [{
            readonly basis: readonly ["auto", (value: string) => boolean, ThemeGetter];
        }];
        /**
         * Flex Direction
         * @see https://tailwindcss.com/docs/flex-direction
         */
        readonly 'flex-direction': readonly [{
            readonly flex: readonly ["row", "row-reverse", "col", "col-reverse"];
        }];
        /**
         * Flex Wrap
         * @see https://tailwindcss.com/docs/flex-wrap
         */
        readonly 'flex-wrap': readonly [{
            readonly flex: readonly ["wrap", "wrap-reverse", "nowrap"];
        }];
        /**
         * Flex
         * @see https://tailwindcss.com/docs/flex
         */
        readonly flex: readonly [{
            readonly flex: readonly ["1", "auto", "initial", "none", (value: string) => boolean];
        }];
        /**
         * Flex Grow
         * @see https://tailwindcss.com/docs/flex-grow
         */
        readonly grow: readonly [{
            readonly grow: readonly ["", "0", (value: string) => boolean];
        }];
        /**
         * Flex Shrink
         * @see https://tailwindcss.com/docs/flex-shrink
         */
        readonly shrink: readonly [{
            readonly shrink: readonly ["", "0", (value: string) => boolean];
        }];
        /**
         * Order
         * @see https://tailwindcss.com/docs/order
         */
        readonly order: readonly [{
            readonly order: readonly ["first", "last", "none", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grid Template Columns
         * @see https://tailwindcss.com/docs/grid-template-columns
         */
        readonly 'grid-cols': readonly [{
            readonly 'grid-cols': readonly [() => boolean];
        }];
        /**
         * Grid Column Start / End
         * @see https://tailwindcss.com/docs/grid-column
         */
        readonly 'col-start-end': readonly [{
            readonly col: readonly ["auto", {
                readonly span: readonly ["full", (value: string) => boolean, (value: string) => boolean];
            }, (value: string) => boolean];
        }];
        /**
         * Grid Column Start
         * @see https://tailwindcss.com/docs/grid-column
         */
        readonly 'col-start': readonly [{
            readonly 'col-start': readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grid Column End
         * @see https://tailwindcss.com/docs/grid-column
         */
        readonly 'col-end': readonly [{
            readonly 'col-end': readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grid Template Rows
         * @see https://tailwindcss.com/docs/grid-template-rows
         */
        readonly 'grid-rows': readonly [{
            readonly 'grid-rows': readonly [() => boolean];
        }];
        /**
         * Grid Row Start / End
         * @see https://tailwindcss.com/docs/grid-row
         */
        readonly 'row-start-end': readonly [{
            readonly row: readonly ["auto", {
                readonly span: readonly [(value: string) => boolean, (value: string) => boolean];
            }, (value: string) => boolean];
        }];
        /**
         * Grid Row Start
         * @see https://tailwindcss.com/docs/grid-row
         */
        readonly 'row-start': readonly [{
            readonly 'row-start': readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grid Row End
         * @see https://tailwindcss.com/docs/grid-row
         */
        readonly 'row-end': readonly [{
            readonly 'row-end': readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grid Auto Flow
         * @see https://tailwindcss.com/docs/grid-auto-flow
         */
        readonly 'grid-flow': readonly [{
            readonly 'grid-flow': readonly ["row", "col", "dense", "row-dense", "col-dense"];
        }];
        /**
         * Grid Auto Columns
         * @see https://tailwindcss.com/docs/grid-auto-columns
         */
        readonly 'auto-cols': readonly [{
            readonly 'auto-cols': readonly ["auto", "min", "max", "fr", (value: string) => boolean];
        }];
        /**
         * Grid Auto Rows
         * @see https://tailwindcss.com/docs/grid-auto-rows
         */
        readonly 'auto-rows': readonly [{
            readonly 'auto-rows': readonly ["auto", "min", "max", "fr", (value: string) => boolean];
        }];
        /**
         * Gap
         * @see https://tailwindcss.com/docs/gap
         */
        readonly gap: readonly [{
            readonly gap: readonly [ThemeGetter];
        }];
        /**
         * Gap X
         * @see https://tailwindcss.com/docs/gap
         */
        readonly 'gap-x': readonly [{
            readonly 'gap-x': readonly [ThemeGetter];
        }];
        /**
         * Gap Y
         * @see https://tailwindcss.com/docs/gap
         */
        readonly 'gap-y': readonly [{
            readonly 'gap-y': readonly [ThemeGetter];
        }];
        /**
         * Justify Content
         * @see https://tailwindcss.com/docs/justify-content
         */
        readonly 'justify-content': readonly [{
            readonly justify: readonly ["normal", "start", "end", "center", "between", "around", "evenly", "stretch"];
        }];
        /**
         * Justify Items
         * @see https://tailwindcss.com/docs/justify-items
         */
        readonly 'justify-items': readonly [{
            readonly 'justify-items': readonly ["start", "end", "center", "stretch"];
        }];
        /**
         * Justify Self
         * @see https://tailwindcss.com/docs/justify-self
         */
        readonly 'justify-self': readonly [{
            readonly 'justify-self': readonly ["auto", "start", "end", "center", "stretch"];
        }];
        /**
         * Align Content
         * @see https://tailwindcss.com/docs/align-content
         */
        readonly 'align-content': readonly [{
            readonly content: readonly ["normal", "start", "end", "center", "between", "around", "evenly", "stretch", "baseline"];
        }];
        /**
         * Align Items
         * @see https://tailwindcss.com/docs/align-items
         */
        readonly 'align-items': readonly [{
            readonly items: readonly ["start", "end", "center", "baseline", "stretch"];
        }];
        /**
         * Align Self
         * @see https://tailwindcss.com/docs/align-self
         */
        readonly 'align-self': readonly [{
            readonly self: readonly ["auto", "start", "end", "center", "stretch", "baseline"];
        }];
        /**
         * Place Content
         * @see https://tailwindcss.com/docs/place-content
         */
        readonly 'place-content': readonly [{
            readonly 'place-content': readonly ["start", "end", "center", "between", "around", "evenly", "stretch", "baseline"];
        }];
        /**
         * Place Items
         * @see https://tailwindcss.com/docs/place-items
         */
        readonly 'place-items': readonly [{
            readonly 'place-items': readonly ["start", "end", "center", "baseline", "stretch"];
        }];
        /**
         * Place Self
         * @see https://tailwindcss.com/docs/place-self
         */
        readonly 'place-self': readonly [{
            readonly 'place-self': readonly ["auto", "start", "end", "center", "stretch"];
        }];
        /**
         * Padding
         * @see https://tailwindcss.com/docs/padding
         */
        readonly p: readonly [{
            readonly p: readonly [ThemeGetter];
        }];
        /**
         * Padding X
         * @see https://tailwindcss.com/docs/padding
         */
        readonly px: readonly [{
            readonly px: readonly [ThemeGetter];
        }];
        /**
         * Padding Y
         * @see https://tailwindcss.com/docs/padding
         */
        readonly py: readonly [{
            readonly py: readonly [ThemeGetter];
        }];
        /**
         * Padding Start
         * @see https://tailwindcss.com/docs/padding
         */
        readonly ps: readonly [{
            readonly ps: readonly [ThemeGetter];
        }];
        /**
         * Padding End
         * @see https://tailwindcss.com/docs/padding
         */
        readonly pe: readonly [{
            readonly pe: readonly [ThemeGetter];
        }];
        /**
         * Padding Top
         * @see https://tailwindcss.com/docs/padding
         */
        readonly pt: readonly [{
            readonly pt: readonly [ThemeGetter];
        }];
        /**
         * Padding Right
         * @see https://tailwindcss.com/docs/padding
         */
        readonly pr: readonly [{
            readonly pr: readonly [ThemeGetter];
        }];
        /**
         * Padding Bottom
         * @see https://tailwindcss.com/docs/padding
         */
        readonly pb: readonly [{
            readonly pb: readonly [ThemeGetter];
        }];
        /**
         * Padding Left
         * @see https://tailwindcss.com/docs/padding
         */
        readonly pl: readonly [{
            readonly pl: readonly [ThemeGetter];
        }];
        /**
         * Margin
         * @see https://tailwindcss.com/docs/margin
         */
        readonly m: readonly [{
            readonly m: readonly [ThemeGetter];
        }];
        /**
         * Margin X
         * @see https://tailwindcss.com/docs/margin
         */
        readonly mx: readonly [{
            readonly mx: readonly [ThemeGetter];
        }];
        /**
         * Margin Y
         * @see https://tailwindcss.com/docs/margin
         */
        readonly my: readonly [{
            readonly my: readonly [ThemeGetter];
        }];
        /**
         * Margin Start
         * @see https://tailwindcss.com/docs/margin
         */
        readonly ms: readonly [{
            readonly ms: readonly [ThemeGetter];
        }];
        /**
         * Margin End
         * @see https://tailwindcss.com/docs/margin
         */
        readonly me: readonly [{
            readonly me: readonly [ThemeGetter];
        }];
        /**
         * Margin Top
         * @see https://tailwindcss.com/docs/margin
         */
        readonly mt: readonly [{
            readonly mt: readonly [ThemeGetter];
        }];
        /**
         * Margin Right
         * @see https://tailwindcss.com/docs/margin
         */
        readonly mr: readonly [{
            readonly mr: readonly [ThemeGetter];
        }];
        /**
         * Margin Bottom
         * @see https://tailwindcss.com/docs/margin
         */
        readonly mb: readonly [{
            readonly mb: readonly [ThemeGetter];
        }];
        /**
         * Margin Left
         * @see https://tailwindcss.com/docs/margin
         */
        readonly ml: readonly [{
            readonly ml: readonly [ThemeGetter];
        }];
        /**
         * Space Between X
         * @see https://tailwindcss.com/docs/space
         */
        readonly 'space-x': readonly [{
            readonly 'space-x': readonly [ThemeGetter];
        }];
        /**
         * Space Between X Reverse
         * @see https://tailwindcss.com/docs/space
         */
        readonly 'space-x-reverse': readonly ["space-x-reverse"];
        /**
         * Space Between Y
         * @see https://tailwindcss.com/docs/space
         */
        readonly 'space-y': readonly [{
            readonly 'space-y': readonly [ThemeGetter];
        }];
        /**
         * Space Between Y Reverse
         * @see https://tailwindcss.com/docs/space
         */
        readonly 'space-y-reverse': readonly ["space-y-reverse"];
        /**
         * Width
         * @see https://tailwindcss.com/docs/width
         */
        readonly w: readonly [{
            readonly w: readonly ["auto", "min", "max", "fit", "svw", "lvw", "dvw", (value: string) => boolean, ThemeGetter];
        }];
        /**
         * Min-Width
         * @see https://tailwindcss.com/docs/min-width
         */
        readonly 'min-w': readonly [{
            readonly 'min-w': readonly [(value: string) => boolean, ThemeGetter, "min", "max", "fit"];
        }];
        /**
         * Max-Width
         * @see https://tailwindcss.com/docs/max-width
         */
        readonly 'max-w': readonly [{
            readonly 'max-w': readonly [(value: string) => boolean, ThemeGetter, "none", "full", "min", "max", "fit", "prose", {
                readonly screen: readonly [(value: string) => boolean];
            }, (value: string) => boolean];
        }];
        /**
         * Height
         * @see https://tailwindcss.com/docs/height
         */
        readonly h: readonly [{
            readonly h: readonly [(value: string) => boolean, ThemeGetter, "auto", "min", "max", "fit", "svh", "lvh", "dvh"];
        }];
        /**
         * Min-Height
         * @see https://tailwindcss.com/docs/min-height
         */
        readonly 'min-h': readonly [{
            readonly 'min-h': readonly [(value: string) => boolean, ThemeGetter, "min", "max", "fit", "svh", "lvh", "dvh"];
        }];
        /**
         * Max-Height
         * @see https://tailwindcss.com/docs/max-height
         */
        readonly 'max-h': readonly [{
            readonly 'max-h': readonly [(value: string) => boolean, ThemeGetter, "min", "max", "fit", "svh", "lvh", "dvh"];
        }];
        /**
         * Size
         * @see https://tailwindcss.com/docs/size
         */
        readonly size: readonly [{
            readonly size: readonly [(value: string) => boolean, ThemeGetter, "auto", "min", "max", "fit"];
        }];
        /**
         * Font Size
         * @see https://tailwindcss.com/docs/font-size
         */
        readonly 'font-size': readonly [{
            readonly text: readonly ["base", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Font Smoothing
         * @see https://tailwindcss.com/docs/font-smoothing
         */
        readonly 'font-smoothing': readonly ["antialiased", "subpixel-antialiased"];
        /**
         * Font Style
         * @see https://tailwindcss.com/docs/font-style
         */
        readonly 'font-style': readonly ["italic", "not-italic"];
        /**
         * Font Weight
         * @see https://tailwindcss.com/docs/font-weight
         */
        readonly 'font-weight': readonly [{
            readonly font: readonly ["thin", "extralight", "light", "normal", "medium", "semibold", "bold", "extrabold", "black", (value: string) => boolean];
        }];
        /**
         * Font Family
         * @see https://tailwindcss.com/docs/font-family
         */
        readonly 'font-family': readonly [{
            readonly font: readonly [() => boolean];
        }];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-normal': readonly ["normal-nums"];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-ordinal': readonly ["ordinal"];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-slashed-zero': readonly ["slashed-zero"];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-figure': readonly ["lining-nums", "oldstyle-nums"];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-spacing': readonly ["proportional-nums", "tabular-nums"];
        /**
         * Font Variant Numeric
         * @see https://tailwindcss.com/docs/font-variant-numeric
         */
        readonly 'fvn-fraction': readonly ["diagonal-fractions", "stacked-fractions"];
        /**
         * Letter Spacing
         * @see https://tailwindcss.com/docs/letter-spacing
         */
        readonly tracking: readonly [{
            readonly tracking: readonly ["tighter", "tight", "normal", "wide", "wider", "widest", (value: string) => boolean];
        }];
        /**
         * Line Clamp
         * @see https://tailwindcss.com/docs/line-clamp
         */
        readonly 'line-clamp': readonly [{
            readonly 'line-clamp': readonly ["none", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Line Height
         * @see https://tailwindcss.com/docs/line-height
         */
        readonly leading: readonly [{
            readonly leading: readonly ["none", "tight", "snug", "normal", "relaxed", "loose", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * List Style Image
         * @see https://tailwindcss.com/docs/list-style-image
         */
        readonly 'list-image': readonly [{
            readonly 'list-image': readonly ["none", (value: string) => boolean];
        }];
        /**
         * List Style Type
         * @see https://tailwindcss.com/docs/list-style-type
         */
        readonly 'list-style-type': readonly [{
            readonly list: readonly ["none", "disc", "decimal", (value: string) => boolean];
        }];
        /**
         * List Style Position
         * @see https://tailwindcss.com/docs/list-style-position
         */
        readonly 'list-style-position': readonly [{
            readonly list: readonly ["inside", "outside"];
        }];
        /**
         * Placeholder Color
         * @deprecated since Tailwind CSS v3.0.0
         * @see https://tailwindcss.com/docs/placeholder-color
         */
        readonly 'placeholder-color': readonly [{
            readonly placeholder: readonly [ThemeGetter];
        }];
        /**
         * Placeholder Opacity
         * @see https://tailwindcss.com/docs/placeholder-opacity
         */
        readonly 'placeholder-opacity': readonly [{
            readonly 'placeholder-opacity': readonly [ThemeGetter];
        }];
        /**
         * Text Alignment
         * @see https://tailwindcss.com/docs/text-align
         */
        readonly 'text-alignment': readonly [{
            readonly text: readonly ["left", "center", "right", "justify", "start", "end"];
        }];
        /**
         * Text Color
         * @see https://tailwindcss.com/docs/text-color
         */
        readonly 'text-color': readonly [{
            readonly text: readonly [ThemeGetter];
        }];
        /**
         * Text Opacity
         * @see https://tailwindcss.com/docs/text-opacity
         */
        readonly 'text-opacity': readonly [{
            readonly 'text-opacity': readonly [ThemeGetter];
        }];
        /**
         * Text Decoration
         * @see https://tailwindcss.com/docs/text-decoration
         */
        readonly 'text-decoration': readonly ["underline", "overline", "line-through", "no-underline"];
        /**
         * Text Decoration Style
         * @see https://tailwindcss.com/docs/text-decoration-style
         */
        readonly 'text-decoration-style': readonly [{
            readonly decoration: readonly ["solid", "dashed", "dotted", "double", "none", "wavy"];
        }];
        /**
         * Text Decoration Thickness
         * @see https://tailwindcss.com/docs/text-decoration-thickness
         */
        readonly 'text-decoration-thickness': readonly [{
            readonly decoration: readonly ["auto", "from-font", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Text Underline Offset
         * @see https://tailwindcss.com/docs/text-underline-offset
         */
        readonly 'underline-offset': readonly [{
            readonly 'underline-offset': readonly ["auto", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Text Decoration Color
         * @see https://tailwindcss.com/docs/text-decoration-color
         */
        readonly 'text-decoration-color': readonly [{
            readonly decoration: readonly [ThemeGetter];
        }];
        /**
         * Text Transform
         * @see https://tailwindcss.com/docs/text-transform
         */
        readonly 'text-transform': readonly ["uppercase", "lowercase", "capitalize", "normal-case"];
        /**
         * Text Overflow
         * @see https://tailwindcss.com/docs/text-overflow
         */
        readonly 'text-overflow': readonly ["truncate", "text-ellipsis", "text-clip"];
        /**
         * Text Wrap
         * @see https://tailwindcss.com/docs/text-wrap
         */
        readonly 'text-wrap': readonly [{
            readonly text: readonly ["wrap", "nowrap", "balance", "pretty"];
        }];
        /**
         * Text Indent
         * @see https://tailwindcss.com/docs/text-indent
         */
        readonly indent: readonly [{
            readonly indent: readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Vertical Alignment
         * @see https://tailwindcss.com/docs/vertical-align
         */
        readonly 'vertical-align': readonly [{
            readonly align: readonly ["baseline", "top", "middle", "bottom", "text-top", "text-bottom", "sub", "super", (value: string) => boolean];
        }];
        /**
         * Whitespace
         * @see https://tailwindcss.com/docs/whitespace
         */
        readonly whitespace: readonly [{
            readonly whitespace: readonly ["normal", "nowrap", "pre", "pre-line", "pre-wrap", "break-spaces"];
        }];
        /**
         * Word Break
         * @see https://tailwindcss.com/docs/word-break
         */
        readonly break: readonly [{
            readonly break: readonly ["normal", "words", "all", "keep"];
        }];
        /**
         * Hyphens
         * @see https://tailwindcss.com/docs/hyphens
         */
        readonly hyphens: readonly [{
            readonly hyphens: readonly ["none", "manual", "auto"];
        }];
        /**
         * Content
         * @see https://tailwindcss.com/docs/content
         */
        readonly content: readonly [{
            readonly content: readonly ["none", (value: string) => boolean];
        }];
        /**
         * Background Attachment
         * @see https://tailwindcss.com/docs/background-attachment
         */
        readonly 'bg-attachment': readonly [{
            readonly bg: readonly ["fixed", "local", "scroll"];
        }];
        /**
         * Background Clip
         * @see https://tailwindcss.com/docs/background-clip
         */
        readonly 'bg-clip': readonly [{
            readonly 'bg-clip': readonly ["border", "padding", "content", "text"];
        }];
        /**
         * Background Opacity
         * @deprecated since Tailwind CSS v3.0.0
         * @see https://tailwindcss.com/docs/background-opacity
         */
        readonly 'bg-opacity': readonly [{
            readonly 'bg-opacity': readonly [ThemeGetter];
        }];
        /**
         * Background Origin
         * @see https://tailwindcss.com/docs/background-origin
         */
        readonly 'bg-origin': readonly [{
            readonly 'bg-origin': readonly ["border", "padding", "content"];
        }];
        /**
         * Background Position
         * @see https://tailwindcss.com/docs/background-position
         */
        readonly 'bg-position': readonly [{
            readonly bg: readonly ["bottom", "center", "left", "left-bottom", "left-top", "right", "right-bottom", "right-top", "top", (value: string) => boolean];
        }];
        /**
         * Background Repeat
         * @see https://tailwindcss.com/docs/background-repeat
         */
        readonly 'bg-repeat': readonly [{
            readonly bg: readonly ["no-repeat", {
                readonly repeat: readonly ["", "x", "y", "round", "space"];
            }];
        }];
        /**
         * Background Size
         * @see https://tailwindcss.com/docs/background-size
         */
        readonly 'bg-size': readonly [{
            readonly bg: readonly ["auto", "cover", "contain", (value: string) => boolean];
        }];
        /**
         * Background Image
         * @see https://tailwindcss.com/docs/background-image
         */
        readonly 'bg-image': readonly [{
            readonly bg: readonly ["none", {
                readonly 'gradient-to': readonly ["t", "tr", "r", "br", "b", "bl", "l", "tl"];
            }, (value: string) => boolean];
        }];
        /**
         * Background Color
         * @see https://tailwindcss.com/docs/background-color
         */
        readonly 'bg-color': readonly [{
            readonly bg: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops From Position
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-from-pos': readonly [{
            readonly from: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops Via Position
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-via-pos': readonly [{
            readonly via: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops To Position
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-to-pos': readonly [{
            readonly to: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops From
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-from': readonly [{
            readonly from: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops Via
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-via': readonly [{
            readonly via: readonly [ThemeGetter];
        }];
        /**
         * Gradient Color Stops To
         * @see https://tailwindcss.com/docs/gradient-color-stops
         */
        readonly 'gradient-to': readonly [{
            readonly to: readonly [ThemeGetter];
        }];
        /**
         * Border Radius
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly rounded: readonly [{
            readonly rounded: readonly [ThemeGetter];
        }];
        /**
         * Border Radius Start
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-s': readonly [{
            readonly 'rounded-s': readonly [ThemeGetter];
        }];
        /**
         * Border Radius End
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-e': readonly [{
            readonly 'rounded-e': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Top
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-t': readonly [{
            readonly 'rounded-t': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Right
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-r': readonly [{
            readonly 'rounded-r': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Bottom
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-b': readonly [{
            readonly 'rounded-b': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Left
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-l': readonly [{
            readonly 'rounded-l': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Start Start
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-ss': readonly [{
            readonly 'rounded-ss': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Start End
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-se': readonly [{
            readonly 'rounded-se': readonly [ThemeGetter];
        }];
        /**
         * Border Radius End End
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-ee': readonly [{
            readonly 'rounded-ee': readonly [ThemeGetter];
        }];
        /**
         * Border Radius End Start
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-es': readonly [{
            readonly 'rounded-es': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Top Left
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-tl': readonly [{
            readonly 'rounded-tl': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Top Right
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-tr': readonly [{
            readonly 'rounded-tr': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Bottom Right
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-br': readonly [{
            readonly 'rounded-br': readonly [ThemeGetter];
        }];
        /**
         * Border Radius Bottom Left
         * @see https://tailwindcss.com/docs/border-radius
         */
        readonly 'rounded-bl': readonly [{
            readonly 'rounded-bl': readonly [ThemeGetter];
        }];
        /**
         * Border Width
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w': readonly [{
            readonly border: readonly [ThemeGetter];
        }];
        /**
         * Border Width X
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-x': readonly [{
            readonly 'border-x': readonly [ThemeGetter];
        }];
        /**
         * Border Width Y
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-y': readonly [{
            readonly 'border-y': readonly [ThemeGetter];
        }];
        /**
         * Border Width Start
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-s': readonly [{
            readonly 'border-s': readonly [ThemeGetter];
        }];
        /**
         * Border Width End
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-e': readonly [{
            readonly 'border-e': readonly [ThemeGetter];
        }];
        /**
         * Border Width Top
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-t': readonly [{
            readonly 'border-t': readonly [ThemeGetter];
        }];
        /**
         * Border Width Right
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-r': readonly [{
            readonly 'border-r': readonly [ThemeGetter];
        }];
        /**
         * Border Width Bottom
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-b': readonly [{
            readonly 'border-b': readonly [ThemeGetter];
        }];
        /**
         * Border Width Left
         * @see https://tailwindcss.com/docs/border-width
         */
        readonly 'border-w-l': readonly [{
            readonly 'border-l': readonly [ThemeGetter];
        }];
        /**
         * Border Opacity
         * @see https://tailwindcss.com/docs/border-opacity
         */
        readonly 'border-opacity': readonly [{
            readonly 'border-opacity': readonly [ThemeGetter];
        }];
        /**
         * Border Style
         * @see https://tailwindcss.com/docs/border-style
         */
        readonly 'border-style': readonly [{
            readonly border: readonly ["solid", "dashed", "dotted", "double", "none", "hidden"];
        }];
        /**
         * Divide Width X
         * @see https://tailwindcss.com/docs/divide-width
         */
        readonly 'divide-x': readonly [{
            readonly 'divide-x': readonly [ThemeGetter];
        }];
        /**
         * Divide Width X Reverse
         * @see https://tailwindcss.com/docs/divide-width
         */
        readonly 'divide-x-reverse': readonly ["divide-x-reverse"];
        /**
         * Divide Width Y
         * @see https://tailwindcss.com/docs/divide-width
         */
        readonly 'divide-y': readonly [{
            readonly 'divide-y': readonly [ThemeGetter];
        }];
        /**
         * Divide Width Y Reverse
         * @see https://tailwindcss.com/docs/divide-width
         */
        readonly 'divide-y-reverse': readonly ["divide-y-reverse"];
        /**
         * Divide Opacity
         * @see https://tailwindcss.com/docs/divide-opacity
         */
        readonly 'divide-opacity': readonly [{
            readonly 'divide-opacity': readonly [ThemeGetter];
        }];
        /**
         * Divide Style
         * @see https://tailwindcss.com/docs/divide-style
         */
        readonly 'divide-style': readonly [{
            readonly divide: readonly ["solid", "dashed", "dotted", "double", "none"];
        }];
        /**
         * Border Color
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color': readonly [{
            readonly border: readonly [ThemeGetter];
        }];
        /**
         * Border Color X
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-x': readonly [{
            readonly 'border-x': readonly [ThemeGetter];
        }];
        /**
         * Border Color Y
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-y': readonly [{
            readonly 'border-y': readonly [ThemeGetter];
        }];
        /**
         * Border Color S
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-s': readonly [{
            readonly 'border-s': readonly [ThemeGetter];
        }];
        /**
         * Border Color E
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-e': readonly [{
            readonly 'border-e': readonly [ThemeGetter];
        }];
        /**
         * Border Color Top
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-t': readonly [{
            readonly 'border-t': readonly [ThemeGetter];
        }];
        /**
         * Border Color Right
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-r': readonly [{
            readonly 'border-r': readonly [ThemeGetter];
        }];
        /**
         * Border Color Bottom
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-b': readonly [{
            readonly 'border-b': readonly [ThemeGetter];
        }];
        /**
         * Border Color Left
         * @see https://tailwindcss.com/docs/border-color
         */
        readonly 'border-color-l': readonly [{
            readonly 'border-l': readonly [ThemeGetter];
        }];
        /**
         * Divide Color
         * @see https://tailwindcss.com/docs/divide-color
         */
        readonly 'divide-color': readonly [{
            readonly divide: readonly [ThemeGetter];
        }];
        /**
         * Outline Style
         * @see https://tailwindcss.com/docs/outline-style
         */
        readonly 'outline-style': readonly [{
            readonly outline: readonly ["", "solid", "dashed", "dotted", "double", "none"];
        }];
        /**
         * Outline Offset
         * @see https://tailwindcss.com/docs/outline-offset
         */
        readonly 'outline-offset': readonly [{
            readonly 'outline-offset': readonly [(value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Outline Width
         * @see https://tailwindcss.com/docs/outline-width
         */
        readonly 'outline-w': readonly [{
            readonly outline: readonly [(value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Outline Color
         * @see https://tailwindcss.com/docs/outline-color
         */
        readonly 'outline-color': readonly [{
            readonly outline: readonly [ThemeGetter];
        }];
        /**
         * Ring Width
         * @see https://tailwindcss.com/docs/ring-width
         */
        readonly 'ring-w': readonly [{
            readonly ring: readonly ["", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Ring Width Inset
         * @see https://tailwindcss.com/docs/ring-width
         */
        readonly 'ring-w-inset': readonly ["ring-inset"];
        /**
         * Ring Color
         * @see https://tailwindcss.com/docs/ring-color
         */
        readonly 'ring-color': readonly [{
            readonly ring: readonly [ThemeGetter];
        }];
        /**
         * Ring Opacity
         * @see https://tailwindcss.com/docs/ring-opacity
         */
        readonly 'ring-opacity': readonly [{
            readonly 'ring-opacity': readonly [ThemeGetter];
        }];
        /**
         * Ring Offset Width
         * @see https://tailwindcss.com/docs/ring-offset-width
         */
        readonly 'ring-offset-w': readonly [{
            readonly 'ring-offset': readonly [(value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Ring Offset Color
         * @see https://tailwindcss.com/docs/ring-offset-color
         */
        readonly 'ring-offset-color': readonly [{
            readonly 'ring-offset': readonly [ThemeGetter];
        }];
        /**
         * Box Shadow
         * @see https://tailwindcss.com/docs/box-shadow
         */
        readonly shadow: readonly [{
            readonly shadow: readonly ["", "inner", "none", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Box Shadow Color
         * @see https://tailwindcss.com/docs/box-shadow-color
         */
        readonly 'shadow-color': readonly [{
            readonly shadow: readonly [() => boolean];
        }];
        /**
         * Opacity
         * @see https://tailwindcss.com/docs/opacity
         */
        readonly opacity: readonly [{
            readonly opacity: readonly [ThemeGetter];
        }];
        /**
         * Mix Blend Mode
         * @see https://tailwindcss.com/docs/mix-blend-mode
         */
        readonly 'mix-blend': readonly [{
            readonly 'mix-blend': readonly ["normal", "multiply", "screen", "overlay", "darken", "lighten", "color-dodge", "color-burn", "hard-light", "soft-light", "difference", "exclusion", "hue", "saturation", "color", "luminosity", "plus-lighter", "plus-darker"];
        }];
        /**
         * Background Blend Mode
         * @see https://tailwindcss.com/docs/background-blend-mode
         */
        readonly 'bg-blend': readonly [{
            readonly 'bg-blend': readonly ["normal", "multiply", "screen", "overlay", "darken", "lighten", "color-dodge", "color-burn", "hard-light", "soft-light", "difference", "exclusion", "hue", "saturation", "color", "luminosity"];
        }];
        /**
         * Filter
         * @deprecated since Tailwind CSS v3.0.0
         * @see https://tailwindcss.com/docs/filter
         */
        readonly filter: readonly [{
            readonly filter: readonly ["", "none"];
        }];
        /**
         * Blur
         * @see https://tailwindcss.com/docs/blur
         */
        readonly blur: readonly [{
            readonly blur: readonly [ThemeGetter];
        }];
        /**
         * Brightness
         * @see https://tailwindcss.com/docs/brightness
         */
        readonly brightness: readonly [{
            readonly brightness: readonly [ThemeGetter];
        }];
        /**
         * Contrast
         * @see https://tailwindcss.com/docs/contrast
         */
        readonly contrast: readonly [{
            readonly contrast: readonly [ThemeGetter];
        }];
        /**
         * Drop Shadow
         * @see https://tailwindcss.com/docs/drop-shadow
         */
        readonly 'drop-shadow': readonly [{
            readonly 'drop-shadow': readonly ["", "none", (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Grayscale
         * @see https://tailwindcss.com/docs/grayscale
         */
        readonly grayscale: readonly [{
            readonly grayscale: readonly [ThemeGetter];
        }];
        /**
         * Hue Rotate
         * @see https://tailwindcss.com/docs/hue-rotate
         */
        readonly 'hue-rotate': readonly [{
            readonly 'hue-rotate': readonly [ThemeGetter];
        }];
        /**
         * Invert
         * @see https://tailwindcss.com/docs/invert
         */
        readonly invert: readonly [{
            readonly invert: readonly [ThemeGetter];
        }];
        /**
         * Saturate
         * @see https://tailwindcss.com/docs/saturate
         */
        readonly saturate: readonly [{
            readonly saturate: readonly [ThemeGetter];
        }];
        /**
         * Sepia
         * @see https://tailwindcss.com/docs/sepia
         */
        readonly sepia: readonly [{
            readonly sepia: readonly [ThemeGetter];
        }];
        /**
         * Backdrop Filter
         * @deprecated since Tailwind CSS v3.0.0
         * @see https://tailwindcss.com/docs/backdrop-filter
         */
        readonly 'backdrop-filter': readonly [{
            readonly 'backdrop-filter': readonly ["", "none"];
        }];
        /**
         * Backdrop Blur
         * @see https://tailwindcss.com/docs/backdrop-blur
         */
        readonly 'backdrop-blur': readonly [{
            readonly 'backdrop-blur': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Brightness
         * @see https://tailwindcss.com/docs/backdrop-brightness
         */
        readonly 'backdrop-brightness': readonly [{
            readonly 'backdrop-brightness': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Contrast
         * @see https://tailwindcss.com/docs/backdrop-contrast
         */
        readonly 'backdrop-contrast': readonly [{
            readonly 'backdrop-contrast': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Grayscale
         * @see https://tailwindcss.com/docs/backdrop-grayscale
         */
        readonly 'backdrop-grayscale': readonly [{
            readonly 'backdrop-grayscale': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Hue Rotate
         * @see https://tailwindcss.com/docs/backdrop-hue-rotate
         */
        readonly 'backdrop-hue-rotate': readonly [{
            readonly 'backdrop-hue-rotate': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Invert
         * @see https://tailwindcss.com/docs/backdrop-invert
         */
        readonly 'backdrop-invert': readonly [{
            readonly 'backdrop-invert': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Opacity
         * @see https://tailwindcss.com/docs/backdrop-opacity
         */
        readonly 'backdrop-opacity': readonly [{
            readonly 'backdrop-opacity': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Saturate
         * @see https://tailwindcss.com/docs/backdrop-saturate
         */
        readonly 'backdrop-saturate': readonly [{
            readonly 'backdrop-saturate': readonly [ThemeGetter];
        }];
        /**
         * Backdrop Sepia
         * @see https://tailwindcss.com/docs/backdrop-sepia
         */
        readonly 'backdrop-sepia': readonly [{
            readonly 'backdrop-sepia': readonly [ThemeGetter];
        }];
        /**
         * Border Collapse
         * @see https://tailwindcss.com/docs/border-collapse
         */
        readonly 'border-collapse': readonly [{
            readonly border: readonly ["collapse", "separate"];
        }];
        /**
         * Border Spacing
         * @see https://tailwindcss.com/docs/border-spacing
         */
        readonly 'border-spacing': readonly [{
            readonly 'border-spacing': readonly [ThemeGetter];
        }];
        /**
         * Border Spacing X
         * @see https://tailwindcss.com/docs/border-spacing
         */
        readonly 'border-spacing-x': readonly [{
            readonly 'border-spacing-x': readonly [ThemeGetter];
        }];
        /**
         * Border Spacing Y
         * @see https://tailwindcss.com/docs/border-spacing
         */
        readonly 'border-spacing-y': readonly [{
            readonly 'border-spacing-y': readonly [ThemeGetter];
        }];
        /**
         * Table Layout
         * @see https://tailwindcss.com/docs/table-layout
         */
        readonly 'table-layout': readonly [{
            readonly table: readonly ["auto", "fixed"];
        }];
        /**
         * Caption Side
         * @see https://tailwindcss.com/docs/caption-side
         */
        readonly caption: readonly [{
            readonly caption: readonly ["top", "bottom"];
        }];
        /**
         * Tranisition Property
         * @see https://tailwindcss.com/docs/transition-property
         */
        readonly transition: readonly [{
            readonly transition: readonly ["none", "all", "", "colors", "opacity", "shadow", "transform", (value: string) => boolean];
        }];
        /**
         * Transition Duration
         * @see https://tailwindcss.com/docs/transition-duration
         */
        readonly duration: readonly [{
            readonly duration: ((value: string) => boolean)[];
        }];
        /**
         * Transition Timing Function
         * @see https://tailwindcss.com/docs/transition-timing-function
         */
        readonly ease: readonly [{
            readonly ease: readonly ["linear", "in", "out", "in-out", (value: string) => boolean];
        }];
        /**
         * Transition Delay
         * @see https://tailwindcss.com/docs/transition-delay
         */
        readonly delay: readonly [{
            readonly delay: ((value: string) => boolean)[];
        }];
        /**
         * Animation
         * @see https://tailwindcss.com/docs/animation
         */
        readonly animate: readonly [{
            readonly animate: readonly ["none", "spin", "ping", "pulse", "bounce", (value: string) => boolean];
        }];
        /**
         * Transform
         * @see https://tailwindcss.com/docs/transform
         */
        readonly transform: readonly [{
            readonly transform: readonly ["", "gpu", "none"];
        }];
        /**
         * Scale
         * @see https://tailwindcss.com/docs/scale
         */
        readonly scale: readonly [{
            readonly scale: readonly [ThemeGetter];
        }];
        /**
         * Scale X
         * @see https://tailwindcss.com/docs/scale
         */
        readonly 'scale-x': readonly [{
            readonly 'scale-x': readonly [ThemeGetter];
        }];
        /**
         * Scale Y
         * @see https://tailwindcss.com/docs/scale
         */
        readonly 'scale-y': readonly [{
            readonly 'scale-y': readonly [ThemeGetter];
        }];
        /**
         * Rotate
         * @see https://tailwindcss.com/docs/rotate
         */
        readonly rotate: readonly [{
            readonly rotate: readonly [(value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Translate X
         * @see https://tailwindcss.com/docs/translate
         */
        readonly 'translate-x': readonly [{
            readonly 'translate-x': readonly [ThemeGetter];
        }];
        /**
         * Translate Y
         * @see https://tailwindcss.com/docs/translate
         */
        readonly 'translate-y': readonly [{
            readonly 'translate-y': readonly [ThemeGetter];
        }];
        /**
         * Skew X
         * @see https://tailwindcss.com/docs/skew
         */
        readonly 'skew-x': readonly [{
            readonly 'skew-x': readonly [ThemeGetter];
        }];
        /**
         * Skew Y
         * @see https://tailwindcss.com/docs/skew
         */
        readonly 'skew-y': readonly [{
            readonly 'skew-y': readonly [ThemeGetter];
        }];
        /**
         * Transform Origin
         * @see https://tailwindcss.com/docs/transform-origin
         */
        readonly 'transform-origin': readonly [{
            readonly origin: readonly ["center", "top", "top-right", "right", "bottom-right", "bottom", "bottom-left", "left", "top-left", (value: string) => boolean];
        }];
        /**
         * Accent Color
         * @see https://tailwindcss.com/docs/accent-color
         */
        readonly accent: readonly [{
            readonly accent: readonly ["auto", ThemeGetter];
        }];
        /**
         * Appearance
         * @see https://tailwindcss.com/docs/appearance
         */
        readonly appearance: readonly [{
            readonly appearance: readonly ["none", "auto"];
        }];
        /**
         * Cursor
         * @see https://tailwindcss.com/docs/cursor
         */
        readonly cursor: readonly [{
            readonly cursor: readonly ["auto", "default", "pointer", "wait", "text", "move", "help", "not-allowed", "none", "context-menu", "progress", "cell", "crosshair", "vertical-text", "alias", "copy", "no-drop", "grab", "grabbing", "all-scroll", "col-resize", "row-resize", "n-resize", "e-resize", "s-resize", "w-resize", "ne-resize", "nw-resize", "se-resize", "sw-resize", "ew-resize", "ns-resize", "nesw-resize", "nwse-resize", "zoom-in", "zoom-out", (value: string) => boolean];
        }];
        /**
         * Caret Color
         * @see https://tailwindcss.com/docs/just-in-time-mode#caret-color-utilities
         */
        readonly 'caret-color': readonly [{
            readonly caret: readonly [ThemeGetter];
        }];
        /**
         * Pointer Events
         * @see https://tailwindcss.com/docs/pointer-events
         */
        readonly 'pointer-events': readonly [{
            readonly 'pointer-events': readonly ["none", "auto"];
        }];
        /**
         * Resize
         * @see https://tailwindcss.com/docs/resize
         */
        readonly resize: readonly [{
            readonly resize: readonly ["none", "y", "x", ""];
        }];
        /**
         * Scroll Behavior
         * @see https://tailwindcss.com/docs/scroll-behavior
         */
        readonly 'scroll-behavior': readonly [{
            readonly scroll: readonly ["auto", "smooth"];
        }];
        /**
         * Scroll Margin
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-m': readonly [{
            readonly 'scroll-m': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin X
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-mx': readonly [{
            readonly 'scroll-mx': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Y
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-my': readonly [{
            readonly 'scroll-my': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Start
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-ms': readonly [{
            readonly 'scroll-ms': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin End
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-me': readonly [{
            readonly 'scroll-me': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Top
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-mt': readonly [{
            readonly 'scroll-mt': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Right
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-mr': readonly [{
            readonly 'scroll-mr': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Bottom
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-mb': readonly [{
            readonly 'scroll-mb': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Margin Left
         * @see https://tailwindcss.com/docs/scroll-margin
         */
        readonly 'scroll-ml': readonly [{
            readonly 'scroll-ml': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-p': readonly [{
            readonly 'scroll-p': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding X
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-px': readonly [{
            readonly 'scroll-px': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Y
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-py': readonly [{
            readonly 'scroll-py': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Start
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-ps': readonly [{
            readonly 'scroll-ps': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding End
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-pe': readonly [{
            readonly 'scroll-pe': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Top
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-pt': readonly [{
            readonly 'scroll-pt': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Right
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-pr': readonly [{
            readonly 'scroll-pr': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Bottom
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-pb': readonly [{
            readonly 'scroll-pb': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Padding Left
         * @see https://tailwindcss.com/docs/scroll-padding
         */
        readonly 'scroll-pl': readonly [{
            readonly 'scroll-pl': readonly [(value: string) => boolean, ThemeGetter];
        }];
        /**
         * Scroll Snap Align
         * @see https://tailwindcss.com/docs/scroll-snap-align
         */
        readonly 'snap-align': readonly [{
            readonly snap: readonly ["start", "end", "center", "align-none"];
        }];
        /**
         * Scroll Snap Stop
         * @see https://tailwindcss.com/docs/scroll-snap-stop
         */
        readonly 'snap-stop': readonly [{
            readonly snap: readonly ["normal", "always"];
        }];
        /**
         * Scroll Snap Type
         * @see https://tailwindcss.com/docs/scroll-snap-type
         */
        readonly 'snap-type': readonly [{
            readonly snap: readonly ["none", "x", "y", "both"];
        }];
        /**
         * Scroll Snap Type Strictness
         * @see https://tailwindcss.com/docs/scroll-snap-type
         */
        readonly 'snap-strictness': readonly [{
            readonly snap: readonly ["mandatory", "proximity"];
        }];
        /**
         * Touch Action
         * @see https://tailwindcss.com/docs/touch-action
         */
        readonly touch: readonly [{
            readonly touch: readonly ["auto", "none", "manipulation"];
        }];
        /**
         * Touch Action X
         * @see https://tailwindcss.com/docs/touch-action
         */
        readonly 'touch-x': readonly [{
            readonly 'touch-pan': readonly ["x", "left", "right"];
        }];
        /**
         * Touch Action Y
         * @see https://tailwindcss.com/docs/touch-action
         */
        readonly 'touch-y': readonly [{
            readonly 'touch-pan': readonly ["y", "up", "down"];
        }];
        /**
         * Touch Action Pinch Zoom
         * @see https://tailwindcss.com/docs/touch-action
         */
        readonly 'touch-pz': readonly ["touch-pinch-zoom"];
        /**
         * User Select
         * @see https://tailwindcss.com/docs/user-select
         */
        readonly select: readonly [{
            readonly select: readonly ["none", "text", "all", "auto"];
        }];
        /**
         * Will Change
         * @see https://tailwindcss.com/docs/will-change
         */
        readonly 'will-change': readonly [{
            readonly 'will-change': readonly ["auto", "scroll", "contents", "transform", (value: string) => boolean];
        }];
        /**
         * Fill
         * @see https://tailwindcss.com/docs/fill
         */
        readonly fill: readonly [{
            readonly fill: readonly [ThemeGetter, "none"];
        }];
        /**
         * Stroke Width
         * @see https://tailwindcss.com/docs/stroke-width
         */
        readonly 'stroke-w': readonly [{
            readonly stroke: readonly [(value: string) => boolean, (value: string) => boolean, (value: string) => boolean];
        }];
        /**
         * Stroke
         * @see https://tailwindcss.com/docs/stroke
         */
        readonly stroke: readonly [{
            readonly stroke: readonly [ThemeGetter, "none"];
        }];
        /**
         * Screen Readers
         * @see https://tailwindcss.com/docs/screen-readers
         */
        readonly sr: readonly ["sr-only", "not-sr-only"];
        /**
         * Forced Color Adjust
         * @see https://tailwindcss.com/docs/forced-color-adjust
         */
        readonly 'forced-color-adjust': readonly [{
            readonly 'forced-color-adjust': readonly ["auto", "none"];
        }];
    };
    readonly conflictingClassGroups: {
        readonly overflow: readonly ["overflow-x", "overflow-y"];
        readonly overscroll: readonly ["overscroll-x", "overscroll-y"];
        readonly inset: readonly ["inset-x", "inset-y", "start", "end", "top", "right", "bottom", "left"];
        readonly 'inset-x': readonly ["right", "left"];
        readonly 'inset-y': readonly ["top", "bottom"];
        readonly flex: readonly ["basis", "grow", "shrink"];
        readonly gap: readonly ["gap-x", "gap-y"];
        readonly p: readonly ["px", "py", "ps", "pe", "pt", "pr", "pb", "pl"];
        readonly px: readonly ["pr", "pl"];
        readonly py: readonly ["pt", "pb"];
        readonly m: readonly ["mx", "my", "ms", "me", "mt", "mr", "mb", "ml"];
        readonly mx: readonly ["mr", "ml"];
        readonly my: readonly ["mt", "mb"];
        readonly size: readonly ["w", "h"];
        readonly 'font-size': readonly ["leading"];
        readonly 'fvn-normal': readonly ["fvn-ordinal", "fvn-slashed-zero", "fvn-figure", "fvn-spacing", "fvn-fraction"];
        readonly 'fvn-ordinal': readonly ["fvn-normal"];
        readonly 'fvn-slashed-zero': readonly ["fvn-normal"];
        readonly 'fvn-figure': readonly ["fvn-normal"];
        readonly 'fvn-spacing': readonly ["fvn-normal"];
        readonly 'fvn-fraction': readonly ["fvn-normal"];
        readonly 'line-clamp': readonly ["display", "overflow"];
        readonly rounded: readonly ["rounded-s", "rounded-e", "rounded-t", "rounded-r", "rounded-b", "rounded-l", "rounded-ss", "rounded-se", "rounded-ee", "rounded-es", "rounded-tl", "rounded-tr", "rounded-br", "rounded-bl"];
        readonly 'rounded-s': readonly ["rounded-ss", "rounded-es"];
        readonly 'rounded-e': readonly ["rounded-se", "rounded-ee"];
        readonly 'rounded-t': readonly ["rounded-tl", "rounded-tr"];
        readonly 'rounded-r': readonly ["rounded-tr", "rounded-br"];
        readonly 'rounded-b': readonly ["rounded-br", "rounded-bl"];
        readonly 'rounded-l': readonly ["rounded-tl", "rounded-bl"];
        readonly 'border-spacing': readonly ["border-spacing-x", "border-spacing-y"];
        readonly 'border-w': readonly ["border-w-s", "border-w-e", "border-w-t", "border-w-r", "border-w-b", "border-w-l"];
        readonly 'border-w-x': readonly ["border-w-r", "border-w-l"];
        readonly 'border-w-y': readonly ["border-w-t", "border-w-b"];
        readonly 'border-color': readonly ["border-color-s", "border-color-e", "border-color-t", "border-color-r", "border-color-b", "border-color-l"];
        readonly 'border-color-x': readonly ["border-color-r", "border-color-l"];
        readonly 'border-color-y': readonly ["border-color-t", "border-color-b"];
        readonly 'scroll-m': readonly ["scroll-mx", "scroll-my", "scroll-ms", "scroll-me", "scroll-mt", "scroll-mr", "scroll-mb", "scroll-ml"];
        readonly 'scroll-mx': readonly ["scroll-mr", "scroll-ml"];
        readonly 'scroll-my': readonly ["scroll-mt", "scroll-mb"];
        readonly 'scroll-p': readonly ["scroll-px", "scroll-py", "scroll-ps", "scroll-pe", "scroll-pt", "scroll-pr", "scroll-pb", "scroll-pl"];
        readonly 'scroll-px': readonly ["scroll-pr", "scroll-pl"];
        readonly 'scroll-py': readonly ["scroll-pt", "scroll-pb"];
        readonly touch: readonly ["touch-x", "touch-y", "touch-pz"];
        readonly 'touch-x': readonly ["touch"];
        readonly 'touch-y': readonly ["touch"];
        readonly 'touch-pz': readonly ["touch"];
    };
    readonly conflictingClassGroupModifiers: {
        readonly 'font-size': readonly ["leading"];
    };
};

type CreateConfigSubsequent = (config: AnyConfig) => AnyConfig;
declare const extendTailwindMerge: <AdditionalClassGroupIds extends string = never, AdditionalThemeGroupIds extends string = never>(configExtension: ConfigExtension<DefaultClassGroupIds | AdditionalClassGroupIds, DefaultThemeGroupIds | AdditionalThemeGroupIds> | CreateConfigSubsequent, ...createConfig: CreateConfigSubsequent[]) => (...classLists: ClassNameValue[]) => string;

declare const fromTheme: <AdditionalThemeGroupIds extends string = never, DefaultThemeGroupIdsInner extends string = DefaultThemeGroupIds>(key: NoInfer<DefaultThemeGroupIdsInner | AdditionalThemeGroupIds>) => ThemeGetter;

/**
 * @param baseConfig Config where other config will be merged into. This object will be mutated.
 * @param configExtension Partial config to merge into the `baseConfig`.
 */
declare const mergeConfigs: <ClassGroupIds extends string, ThemeGroupIds extends string = never>(baseConfig: AnyConfig, { cacheSize, prefix, separator, experimentalParseClassName, extend, override, }: ConfigExtension<ClassGroupIds, ThemeGroupIds>) => AnyConfig;

declare const twMerge: (...classLists: ClassNameValue[]) => string;

declare const isLength: (value: string) => boolean;
declare const isArbitraryLength: (value: string) => boolean;
declare const isNumber: (value: string) => boolean;
declare const isArbitraryNumber: (value: string) => boolean;
declare const isInteger: (value: string) => boolean;
declare const isPercent: (value: string) => boolean;
declare const isArbitraryValue: (value: string) => boolean;
declare const isTshirtSize: (value: string) => boolean;
declare const isArbitrarySize: (value: string) => boolean;
declare const isArbitraryPosition: (value: string) => boolean;
declare const isArbitraryImage: (value: string) => boolean;
declare const isArbitraryShadow: (value: string) => boolean;
declare const isAny: () => boolean;

declare const validators_d_isAny: typeof isAny;
declare const validators_d_isArbitraryImage: typeof isArbitraryImage;
declare const validators_d_isArbitraryLength: typeof isArbitraryLength;
declare const validators_d_isArbitraryNumber: typeof isArbitraryNumber;
declare const validators_d_isArbitraryPosition: typeof isArbitraryPosition;
declare const validators_d_isArbitraryShadow: typeof isArbitraryShadow;
declare const validators_d_isArbitrarySize: typeof isArbitrarySize;
declare const validators_d_isArbitraryValue: typeof isArbitraryValue;
declare const validators_d_isInteger: typeof isInteger;
declare const validators_d_isLength: typeof isLength;
declare const validators_d_isNumber: typeof isNumber;
declare const validators_d_isPercent: typeof isPercent;
declare const validators_d_isTshirtSize: typeof isTshirtSize;
declare namespace validators_d {
  export { validators_d_isAny as isAny, validators_d_isArbitraryImage as isArbitraryImage, validators_d_isArbitraryLength as isArbitraryLength, validators_d_isArbitraryNumber as isArbitraryNumber, validators_d_isArbitraryPosition as isArbitraryPosition, validators_d_isArbitraryShadow as isArbitraryShadow, validators_d_isArbitrarySize as isArbitrarySize, validators_d_isArbitraryValue as isArbitraryValue, validators_d_isInteger as isInteger, validators_d_isLength as isLength, validators_d_isNumber as isNumber, validators_d_isPercent as isPercent, validators_d_isTshirtSize as isTshirtSize };
}

export { type ClassNameValue, type ClassValidator, type Config, type ConfigExtension, type DefaultClassGroupIds, type DefaultThemeGroupIds, type ExperimentalParseClassNameParam, type ExperimentalParsedClassName, createTailwindMerge, extendTailwindMerge, fromTheme, getDefaultConfig, mergeConfigs, twJoin, twMerge, validators_d as validators };
