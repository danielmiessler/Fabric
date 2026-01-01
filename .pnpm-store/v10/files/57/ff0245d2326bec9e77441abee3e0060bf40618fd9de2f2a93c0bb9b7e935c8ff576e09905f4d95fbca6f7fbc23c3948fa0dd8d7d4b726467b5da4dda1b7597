// src/index.ts
import { PurgeCSS } from "purgecss";

// src/extractors/regex.ts
var REGEX_SPECIAL = /[\\^$.*+?()[\]{}|]/g;
var REGEX_HAS_SPECIAL = RegExp(REGEX_SPECIAL.source);
function toSource(source) {
  source = Array.isArray(source) ? source : [source];
  source = source.map((item) => item instanceof RegExp ? item.source : item);
  return source.join("");
}
function pattern(source) {
  return new RegExp(toSource(source), "g");
}
function any(sources) {
  return `(?:${sources.map(toSource).join("|")})`;
}
function optional(source) {
  return `(?:${toSource(source)})?`;
}

// src/extractors/default-extractor.ts
function defaultExtractor() {
  let patterns = Array.from(buildRegExps());
  return (content) => {
    let results = [];
    for (let pattern2 of patterns) {
      results = [...results, ...content.match(pattern2) ?? []];
    }
    return results.filter((v) => v !== void 0).map(clipAtBalancedParens);
  };
}
function* buildRegExps() {
  let separator = ":";
  let prefix = "";
  let utility = any([
    /\[[^\s:'"`]+:[^\s\[\]]+\]/,
    /\[[^\s:'"`]+:[^\s]+?\[[^\s]+\][^\s]+?\]/,
    pattern([
      /-?(?:\w+)/,
      optional(
        any([
          pattern([
            /-(?:\w+-)*\[[^\s:]+\]/,
            /(?![{([]])/,
            /(?:\/[^\s'"`\\><$]*)?/
          ]),
          pattern([
            /-(?:\w+-)*\[[^\s]+\]/,
            /(?![{([]])/,
            /(?:\/[^\s'"`\\$]*)?/
          ]),
          /[-\/][^\s'"`\\$={><]*/
        ])
      )
    ])
  ]);
  let variantPatterns = [
    any([
      pattern([/@\[[^\s"'`]+\](\/[^\s"'`]+)?/, separator]),
      pattern([/([^\s"'`\[\\]+-)?\[[^\s"'`]+\]/, separator]),
      pattern([/[^\s"'`\[\\]+/, separator])
    ]),
    any([
      pattern([/([^\s"'`\[\\]+-)?\[[^\s`]+\]/, separator]),
      pattern([/[^\s`\[\\]+/, separator])
    ])
  ];
  for (const variantPattern of variantPatterns) {
    yield pattern([
      "((?=((",
      variantPattern,
      ")+))\\2)?",
      /!?/,
      prefix,
      utility
    ]);
  }
  yield /[^<>"'`\s.(){}[\]#=%$]*[^<>"'`\s.(){}[\]#=%:$]/g;
}
var SPECIALS = /([\[\]'"`])([^\[\]'"`])?/g;
var ALLOWED_CLASS_CHARACTERS = /[^"'`\s<>\]]+/;
function clipAtBalancedParens(input) {
  if (!input.includes("-[")) {
    return input;
  }
  let depth = 0;
  let openStringTypes = [];
  let matches = input.matchAll(SPECIALS);
  const matched = Array.from(matches).flatMap((match) => {
    const [, ...groups] = match;
    return groups.map(
      (group, idx) => Object.assign([], match, {
        index: match.index + idx,
        0: group
      })
    );
  });
  for (let match of matched) {
    let char = match[0];
    let inStringType = openStringTypes.at(-1);
    if (char === inStringType) {
      openStringTypes.pop();
    } else if (char === "'" || char === '"' || char === "`") {
      openStringTypes.push(char);
    }
    if (inStringType) {
      continue;
    } else if (char === "[") {
      depth++;
      continue;
    } else if (char === "]") {
      depth--;
      continue;
    }
    if (depth < 0) {
      return input.substring(0, match.index - 1);
    }
    if (depth === 0 && !ALLOWED_CLASS_CHARACTERS.test(char)) {
      return input.substring(0, match.index);
    }
  }
  return input;
}

// src/index.ts
import { walk } from "estree-walker";
import path from "path";
var EXT_CSS = /\.(css)$/;
var MAX_STRING_LITERAL_LENGTH = 5e4;
function purgeCss(purgeOptions) {
  var _a;
  let viteConfig;
  const selectors = /* @__PURE__ */ new Set();
  const standard = [
    "*",
    "html",
    "body",
    /aria-current/,
    /^\:[-a-z]+$/,
    ...((_a = purgeOptions == null ? void 0 : purgeOptions.safelist) == null ? void 0 : _a.standard) ?? []
  ];
  const extractor = (purgeOptions == null ? void 0 : purgeOptions.defaultExtractor) ?? defaultExtractor();
  const moduleIds = /* @__PURE__ */ new Set();
  return {
    name: "vite-plugin-tailwind-purgecss",
    apply: "build",
    enforce: "post",
    load(id) {
      if (EXT_CSS.test(id))
        return;
      moduleIds.add(id);
    },
    configResolved(config) {
      viteConfig = config;
    },
    async generateBundle(options, bundle) {
      var _a2;
      const assets = {};
      for (const id of moduleIds) {
        const info = this.getModuleInfo(id);
        if ((info == null ? void 0 : info.isIncluded) !== true || info.code === null)
          continue;
        const ast = this.parse(info.code);
        walk(ast, {
          enter(node, parent, key, index) {
            if (node.type === "Literal" && typeof node.value === "string") {
              node.value.split(/\s+/).forEach((word) => {
                if (word.length < MAX_STRING_LITERAL_LENGTH) {
                  extractor(word).forEach((selector) => selectors.add(selector));
                } else
                  selectors.add(word);
              });
            }
            if (node.type === "Identifier") {
              selectors.add(node.name);
            }
            if (node.type === "TemplateElement") {
              const value = node.value.cooked ?? node.value.raw;
              value.split(/\s+/).forEach((word) => {
                if (word.length < MAX_STRING_LITERAL_LENGTH) {
                  extractor(word).forEach((selector) => selectors.add(selector));
                } else
                  selectors.add(word);
              });
            }
          }
        });
      }
      for (const [fileName, chunkOrAsset] of Object.entries(bundle)) {
        if (chunkOrAsset.type === "asset" && EXT_CSS.test(fileName)) {
          assets[fileName] = chunkOrAsset;
        }
      }
      for (const selector of selectors) {
        standard.push(selector);
      }
      const htmlGlob = viteConfig.root.split(path.sep).join("/") + "/**/*.html";
      for (const [fileName, asset] of Object.entries(assets)) {
        const purgeCSSResult = await new PurgeCSS().purge({
          ...purgeOptions,
          content: [htmlGlob, ...(purgeOptions == null ? void 0 : purgeOptions.content) ?? []],
          css: [{ raw: asset.source.trim(), name: fileName }],
          safelist: {
            ...purgeOptions == null ? void 0 : purgeOptions.safelist,
            standard,
            greedy: [/svelte-/, /data-theme/, ...((_a2 = purgeOptions == null ? void 0 : purgeOptions.safelist) == null ? void 0 : _a2.greedy) ?? []]
          }
        });
        if (purgeCSSResult[0]) {
          delete bundle[asset.fileName];
          this.emitFile({
            ...asset,
            type: "asset",
            source: purgeCSSResult[0].css
          });
        }
      }
    }
  };
}
var src_default = purgeCss;
export {
  src_default as default,
  purgeCss
};
//# sourceMappingURL=index.js.map