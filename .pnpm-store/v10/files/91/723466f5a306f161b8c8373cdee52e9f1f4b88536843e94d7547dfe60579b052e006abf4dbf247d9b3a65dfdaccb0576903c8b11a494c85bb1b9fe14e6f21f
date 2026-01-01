// src/index.ts
import * as acornNamespace from "acorn";

// src/tokenType.ts
import { TokenType, keywordTypes, tokTypes, TokContext } from "acorn";
var startsExpr = true;
function kwLike(_name, options = {}) {
  return new TokenType("name", options);
}
var acornTypeScriptMap = /* @__PURE__ */ new WeakMap();
var keywordTypeValues = Object.values(keywordTypes);
function generateAcornTypeScript(_acorn) {
  const acorn = _acorn.Parser.acorn || _acorn;
  let acornTypeScript = acornTypeScriptMap.get(acorn);
  if (!acornTypeScript) {
    let tokenIsLiteralPropertyName = function(token) {
      return token === tokTypes.name || token === tokTypes.string || token === tokTypes.num || keywordTypeValues.includes(token) || tsKwTokenTypeValues.includes(token);
    }, tokenIsKeywordOrIdentifier = function(token) {
      return token === tokTypes.name || keywordTypeValues.includes(token) || tsKwTokenTypeValues.includes(token);
    }, tokenIsIdentifier = function(token) {
      return token === tokTypes.name || tsKwTokenTypeValues.includes(token);
    }, tokenIsTSDeclarationStart = function(token) {
      return token === tsKwTokenType.abstract || token === tsKwTokenType.declare || token === tsKwTokenType.enum || token === tsKwTokenType.module || token === tsKwTokenType.namespace || token === tsKwTokenType.interface || token === tsKwTokenType.type;
    }, tokenIsTSTypeOperator = function(token) {
      return token === tsKwTokenType.keyof || token === tsKwTokenType.readonly || token === tsKwTokenType.unique;
    }, tokenIsTemplate = function(token) {
      return token === tokTypes.invalidTemplate;
    };
    const tsKwTokenType = generateTsKwTokenType();
    const tsKwTokenTypeValues = Object.values(tsKwTokenType);
    const tsTokenType = generateTsTokenType();
    const tsTokenContext = generateTsTokenContext();
    const tsKeywordsRegExp = new RegExp(`^(?:${Object.keys(tsKwTokenType).join("|")})$`);
    tsTokenType.jsxTagStart.updateContext = function() {
      this.context.push(tsTokenContext.tc_expr);
      this.context.push(tsTokenContext.tc_oTag);
      this.exprAllowed = false;
    };
    tsTokenType.jsxTagEnd.updateContext = function(prevType) {
      let out = this.context.pop();
      if (out === tsTokenContext.tc_oTag && prevType === tokTypes.slash || out === tsTokenContext.tc_cTag) {
        this.context.pop();
        this.exprAllowed = this.curContext() === tsTokenContext.tc_expr;
      } else {
        this.exprAllowed = true;
      }
    };
    acornTypeScript = {
      tokTypes: {
        ...tsKwTokenType,
        ...tsTokenType
      },
      tokContexts: {
        ...tsTokenContext
      },
      keywordsRegExp: tsKeywordsRegExp,
      tokenIsLiteralPropertyName,
      tokenIsKeywordOrIdentifier,
      tokenIsIdentifier,
      tokenIsTSDeclarationStart,
      tokenIsTSTypeOperator,
      tokenIsTemplate
    };
  }
  return acornTypeScript;
}
function generateTsTokenContext() {
  return {
    tc_oTag: new TokContext("<tag", false, false),
    tc_cTag: new TokContext("</tag", false, false),
    tc_expr: new TokContext("<tag>...</tag>", true, true)
  };
}
function generateTsTokenType() {
  return {
    // @ts-expect-error
    at: new TokenType("@"),
    // @ts-expect-error
    jsxName: new TokenType("jsxName"),
    // @ts-expect-error
    jsxText: new TokenType("jsxText", { beforeExpr: true }),
    // @ts-expect-error
    jsxTagStart: new TokenType("jsxTagStart", { startsExpr: true }),
    // @ts-expect-error
    jsxTagEnd: new TokenType("jsxTagEnd")
  };
}
function generateTsKwTokenType() {
  return {
    assert: kwLike("assert", { startsExpr }),
    asserts: kwLike("asserts", { startsExpr }),
    global: kwLike("global", { startsExpr }),
    keyof: kwLike("keyof", { startsExpr }),
    readonly: kwLike("readonly", { startsExpr }),
    unique: kwLike("unique", { startsExpr }),
    abstract: kwLike("abstract", { startsExpr }),
    declare: kwLike("declare", { startsExpr }),
    enum: kwLike("enum", { startsExpr }),
    module: kwLike("module", { startsExpr }),
    namespace: kwLike("namespace", { startsExpr }),
    interface: kwLike("interface", { startsExpr }),
    type: kwLike("type", { startsExpr })
  };
}

// src/scopeflags.ts
var TS_SCOPE_OTHER = 512;
var TS_SCOPE_TS_MODULE = 1024;
var BIND_KIND_VALUE = 1;
var BIND_KIND_TYPE = 2;
var BIND_SCOPE_VAR = 4;
var BIND_SCOPE_LEXICAL = 8;
var BIND_SCOPE_FUNCTION = 16;
var BIND_FLAGS_NONE = 64;
var BIND_FLAGS_CLASS = 128;
var BIND_FLAGS_TS_ENUM = 256;
var BIND_FLAGS_TS_CONST_ENUM = 512;
var BIND_FLAGS_TS_EXPORT_ONLY = 1024;
var BIND_CLASS = BIND_KIND_VALUE | BIND_KIND_TYPE | BIND_SCOPE_LEXICAL | BIND_FLAGS_CLASS;
var BIND_LEXICAL = BIND_KIND_VALUE | 0 | BIND_SCOPE_LEXICAL | 0;
var BIND_VAR = BIND_KIND_VALUE | 0 | BIND_SCOPE_VAR | 0;
var BIND_FUNCTION = BIND_KIND_VALUE | 0 | BIND_SCOPE_FUNCTION | 0;
var BIND_TS_INTERFACE = 0 | BIND_KIND_TYPE | 0 | BIND_FLAGS_CLASS;
var BIND_TS_TYPE = 0 | BIND_KIND_TYPE | 0 | 0;
var BIND_TS_ENUM = BIND_KIND_VALUE | BIND_KIND_TYPE | BIND_SCOPE_LEXICAL | BIND_FLAGS_TS_ENUM;
var BIND_TS_AMBIENT = 0 | 0 | 0 | BIND_FLAGS_TS_EXPORT_ONLY;
var BIND_NONE = 0 | 0 | 0 | BIND_FLAGS_NONE;
var BIND_OUTSIDE = BIND_KIND_VALUE | 0 | 0 | BIND_FLAGS_NONE;
var BIND_TS_CONST_ENUM = BIND_TS_ENUM | BIND_FLAGS_TS_CONST_ENUM;
var BIND_TS_NAMESPACE = 0 | 0 | 0 | BIND_FLAGS_TS_EXPORT_ONLY;
var CLASS_ELEMENT_FLAG_STATIC = 4;
var CLASS_ELEMENT_KIND_GETTER = 2;
var CLASS_ELEMENT_KIND_SETTER = 1;
var CLASS_ELEMENT_KIND_ACCESSOR = CLASS_ELEMENT_KIND_GETTER | CLASS_ELEMENT_KIND_SETTER;
var CLASS_ELEMENT_STATIC_GETTER = CLASS_ELEMENT_KIND_GETTER | CLASS_ELEMENT_FLAG_STATIC;
var CLASS_ELEMENT_STATIC_SETTER = CLASS_ELEMENT_KIND_SETTER | CLASS_ELEMENT_FLAG_STATIC;

// src/whitespace.ts
var skipWhiteSpaceInLine = /(?:[^\S\n\r\u2028\u2029]|\/\/.*|\/\*.*?\*\/)*/y;
var skipWhiteSpaceToLineBreak = new RegExp(
  // Unfortunately JS doesn't support Perl's atomic /(?>pattern)/ or
  // possessive quantifiers, so we use a trick to prevent backtracking
  // when the look-ahead for line terminator fails.
  "(?=(" + // Capture the whitespace and comments that should be skipped inside
  // a look-ahead assertion, and then re-match the group as a unit.
  skipWhiteSpaceInLine.source + "))\\1" + // Look-ahead for either line terminator, start of multi-line comment,
  // or end of string.
  /(?=[\n\r\u2028\u2029]|\/\*(?!.*?\*\/)|$)/.source,
  "y"
  // sticky
);

// src/parseutil.ts
var DestructuringErrors = class {
  constructor() {
    this.shorthandAssign = this.trailingComma = this.parenthesizedAssign = this.parenthesizedBind = this.doubleProto = -1;
  }
};
function isPrivateNameConflicted(privateNameMap, element) {
  const name = element.key.name;
  const curr = privateNameMap[name];
  let next = "true";
  if (element.type === "MethodDefinition" && (element.kind === "get" || element.kind === "set")) {
    next = (element.static ? "s" : "i") + element.kind;
  }
  if (curr === "iget" && next === "iset" || curr === "iset" && next === "iget" || curr === "sget" && next === "sset" || curr === "sset" && next === "sget") {
    privateNameMap[name] = "true";
    return false;
  } else if (!curr) {
    privateNameMap[name] = next;
    return false;
  } else {
    return true;
  }
}
function checkKeyName(node, name) {
  const { computed, key } = node;
  return !computed && (key.type === "Identifier" && key.name === name || key.type === "Literal" && key.value === name);
}

// src/error.ts
var TypeScriptError = {
  AbstractMethodHasImplementation: ({ methodName }) => `Method '${methodName}' cannot have an implementation because it is marked abstract.`,
  AbstractPropertyHasInitializer: ({ propertyName }) => `Property '${propertyName}' cannot have an initializer because it is marked abstract.`,
  AccesorCannotDeclareThisParameter: "'get' and 'set' accessors cannot declare 'this' parameters.",
  AccesorCannotHaveTypeParameters: "An accessor cannot have type parameters.",
  CannotFindName: ({ name }) => `Cannot find name '${name}'.`,
  ClassMethodHasDeclare: "Class methods cannot have the 'declare' modifier.",
  ClassMethodHasReadonly: "Class methods cannot have the 'readonly' modifier.",
  ConstInitiailizerMustBeStringOrNumericLiteralOrLiteralEnumReference: "A 'const' initializer in an ambient context must be a string or numeric literal or literal enum reference.",
  ConstructorHasTypeParameters: "Type parameters cannot appear on a constructor declaration.",
  DeclareAccessor: ({ kind }) => `'declare' is not allowed in ${kind}ters.`,
  DeclareClassFieldHasInitializer: "Initializers are not allowed in ambient contexts.",
  DeclareFunctionHasImplementation: "An implementation cannot be declared in ambient contexts.",
  DuplicateAccessibilityModifier: (
    // `Accessibility modifier already seen: ${modifier}` would be more helpful.
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    () => `Accessibility modifier already seen.`
  ),
  DuplicateModifier: ({ modifier }) => `Duplicate modifier: '${modifier}'.`,
  // `token` matches the terminology used by typescript:
  // https://github.com/microsoft/TypeScript/blob/main/src/compiler/types.ts#L2915
  EmptyHeritageClauseType: ({ token }) => `'${token}' list cannot be empty.`,
  EmptyTypeArguments: "Type argument list cannot be empty.",
  EmptyTypeParameters: "Type parameter list cannot be empty.",
  ExpectedAmbientAfterExportDeclare: "'export declare' must be followed by an ambient declaration.",
  ImportAliasHasImportType: "An import alias can not use 'import type'.",
  IncompatibleModifiers: ({ modifiers }) => `'${modifiers[0]}' modifier cannot be used with '${modifiers[1]}' modifier.`,
  IndexSignatureHasAbstract: "Index signatures cannot have the 'abstract' modifier.",
  IndexSignatureHasAccessibility: ({ modifier }) => `Index signatures cannot have an accessibility modifier ('${modifier}').`,
  IndexSignatureHasDeclare: "Index signatures cannot have the 'declare' modifier.",
  IndexSignatureHasOverride: "'override' modifier cannot appear on an index signature.",
  IndexSignatureHasStatic: "Index signatures cannot have the 'static' modifier.",
  InitializerNotAllowedInAmbientContext: "Initializers are not allowed in ambient contexts.",
  InvalidModifierOnTypeMember: ({ modifier }) => `'${modifier}' modifier cannot appear on a type member.`,
  InvalidModifierOnTypeParameter: ({ modifier }) => `'${modifier}' modifier cannot appear on a type parameter.`,
  InvalidModifierOnTypeParameterPositions: ({ modifier }) => `'${modifier}' modifier can only appear on a type parameter of a class, interface or type alias.`,
  InvalidModifiersOrder: ({ orderedModifiers }) => `'${orderedModifiers[0]}' modifier must precede '${orderedModifiers[1]}' modifier.`,
  InvalidPropertyAccessAfterInstantiationExpression: "Invalid property access after an instantiation expression. You can either wrap the instantiation expression in parentheses, or delete the type arguments.",
  InvalidTupleMemberLabel: "Tuple members must be labeled with a simple identifier.",
  MissingInterfaceName: "'interface' declarations must be followed by an identifier.",
  NonAbstractClassHasAbstractMethod: "Abstract methods can only appear within an abstract class.",
  NonClassMethodPropertyHasAbstractModifer: "'abstract' modifier can only appear on a class, method, or property declaration.",
  OptionalTypeBeforeRequired: "A required element cannot follow an optional element.",
  OverrideNotInSubClass: "This member cannot have an 'override' modifier because its containing class does not extend another class.",
  PatternIsOptional: "A binding pattern parameter cannot be optional in an implementation signature.",
  PrivateElementHasAbstract: "Private elements cannot have the 'abstract' modifier.",
  PrivateElementHasAccessibility: ({ modifier }) => `Private elements cannot have an accessibility modifier ('${modifier}').`,
  PrivateMethodsHasAccessibility: ({ modifier }) => `Private methods cannot have an accessibility modifier ('${modifier}').`,
  ReadonlyForMethodSignature: "'readonly' modifier can only appear on a property declaration or index signature.",
  ReservedArrowTypeParam: "This syntax is reserved in files with the .mts or .cts extension. Add a trailing comma, as in `<T,>() => ...`.",
  ReservedTypeAssertion: "This syntax is reserved in files with the .mts or .cts extension. Use an `as` expression instead.",
  SetAccesorCannotHaveOptionalParameter: "A 'set' accessor cannot have an optional parameter.",
  SetAccesorCannotHaveRestParameter: "A 'set' accessor cannot have rest parameter.",
  SetAccesorCannotHaveReturnType: "A 'set' accessor cannot have a return type annotation.",
  SingleTypeParameterWithoutTrailingComma: ({ typeParameterName }) => `Single type parameter ${typeParameterName} should have a trailing comma. Example usage: <${typeParameterName},>.`,
  StaticBlockCannotHaveModifier: "Static class blocks cannot have any modifier.",
  TypeAnnotationAfterAssign: "Type annotations must come before default assignments, e.g. instead of `age = 25: number` use `age: number = 25`.",
  TypeImportCannotSpecifyDefaultAndNamed: "A type-only import can specify a default import or named bindings, but not both.",
  TypeModifierIsUsedInTypeExports: "The 'type' modifier cannot be used on a named export when 'export type' is used on its export statement.",
  TypeModifierIsUsedInTypeImports: "The 'type' modifier cannot be used on a named import when 'import type' is used on its import statement.",
  UnexpectedParameterModifier: "A parameter property is only allowed in a constructor implementation.",
  UnexpectedReadonly: "'readonly' type modifier is only permitted on array and tuple literal types.",
  GenericsEndWithComma: `Trailing comma is not allowed at the end of generics.`,
  UnexpectedTypeAnnotation: "Did not expect a type annotation here.",
  UnexpectedTypeCastInParameter: "Unexpected type cast in parameter position.",
  UnsupportedImportTypeArgument: "Argument in a type import must be a string literal.",
  UnsupportedParameterPropertyKind: "A parameter property may not be declared using a binding pattern.",
  UnsupportedSignatureParameterKind: ({ type }) => `Name in a signature must be an Identifier, ObjectPattern or ArrayPattern, instead got ${type}.`,
  LetInLexicalBinding: "'let' is not allowed to be used as a name in 'let' or 'const' declarations."
};
var DecoratorsError = {
  UnexpectedLeadingDecorator: "Leading decorators must be attached to a class declaration.",
  DecoratorConstructor: "Decorators can't be used with a constructor. Did you mean '@dec class { ... }'?",
  TrailingDecorator: "Decorators must be attached to a class element.",
  SpreadElementDecorator: `Decorators can't be used with SpreadElement`
};

// src/extentions/decorators.ts
function generateParseDecorators(Parse, acornTypeScript, acorn) {
  const { tokTypes: tt } = acorn;
  const { tokTypes: tokTypes2 } = acornTypeScript;
  return class ParseDecorators extends Parse {
    takeDecorators(node) {
      const decorators = this.decoratorStack[this.decoratorStack.length - 1];
      if (decorators.length) {
        node.decorators = decorators;
        this.resetStartLocationFromNode(node, decorators[0]);
        this.decoratorStack[this.decoratorStack.length - 1] = [];
      }
    }
    parseDecorators(allowExport) {
      const currentContextDecorators = this.decoratorStack[this.decoratorStack.length - 1];
      while (this.match(tokTypes2.at)) {
        const decorator = this.parseDecorator();
        currentContextDecorators.push(decorator);
      }
      if (this.match(tt._export)) {
        if (!allowExport) {
          this.unexpected();
        }
      } else if (!this.canHaveLeadingDecorator()) {
        this.raise(this.start, DecoratorsError.UnexpectedLeadingDecorator);
      }
    }
    parseDecorator() {
      const node = this.startNode();
      this.next();
      this.decoratorStack.push([]);
      const startPos = this.start;
      const startLoc = this.startLoc;
      let expr;
      if (this.match(tt.parenL)) {
        const startPos2 = this.start;
        const startLoc2 = this.startLoc;
        this.next();
        expr = this.parseExpression();
        this.expect(tt.parenR);
        if (this.options.preserveParens) {
          let par = this.startNodeAt(startPos2, startLoc2);
          par.expression = expr;
          expr = this.finishNode(par, "ParenthesizedExpression");
        }
      } else {
        expr = this.parseIdent(false);
        while (this.eat(tt.dot)) {
          const node2 = this.startNodeAt(startPos, startLoc);
          node2.object = expr;
          node2.property = this.parseIdent(true);
          node2.computed = false;
          expr = this.finishNode(node2, "MemberExpression");
        }
      }
      node.expression = this.parseMaybeDecoratorArguments(expr);
      this.decoratorStack.pop();
      return this.finishNode(node, "Decorator");
    }
    parseMaybeDecoratorArguments(expr) {
      if (this.eat(tt.parenL)) {
        const node = this.startNodeAtNode(expr);
        node.callee = expr;
        node.arguments = this.parseExprList(tt.parenR, false);
        return this.finishNode(node, "CallExpression");
      }
      return expr;
    }
  };
}

// src/extentions/jsx/xhtml.ts
var xhtml_default = {
  quot: '"',
  amp: "&",
  apos: "'",
  lt: "<",
  gt: ">",
  nbsp: "\xA0",
  iexcl: "\xA1",
  cent: "\xA2",
  pound: "\xA3",
  curren: "\xA4",
  yen: "\xA5",
  brvbar: "\xA6",
  sect: "\xA7",
  uml: "\xA8",
  copy: "\xA9",
  ordf: "\xAA",
  laquo: "\xAB",
  not: "\xAC",
  shy: "\xAD",
  reg: "\xAE",
  macr: "\xAF",
  deg: "\xB0",
  plusmn: "\xB1",
  sup2: "\xB2",
  sup3: "\xB3",
  acute: "\xB4",
  micro: "\xB5",
  para: "\xB6",
  middot: "\xB7",
  cedil: "\xB8",
  sup1: "\xB9",
  ordm: "\xBA",
  raquo: "\xBB",
  frac14: "\xBC",
  frac12: "\xBD",
  frac34: "\xBE",
  iquest: "\xBF",
  Agrave: "\xC0",
  Aacute: "\xC1",
  Acirc: "\xC2",
  Atilde: "\xC3",
  Auml: "\xC4",
  Aring: "\xC5",
  AElig: "\xC6",
  Ccedil: "\xC7",
  Egrave: "\xC8",
  Eacute: "\xC9",
  Ecirc: "\xCA",
  Euml: "\xCB",
  Igrave: "\xCC",
  Iacute: "\xCD",
  Icirc: "\xCE",
  Iuml: "\xCF",
  ETH: "\xD0",
  Ntilde: "\xD1",
  Ograve: "\xD2",
  Oacute: "\xD3",
  Ocirc: "\xD4",
  Otilde: "\xD5",
  Ouml: "\xD6",
  times: "\xD7",
  Oslash: "\xD8",
  Ugrave: "\xD9",
  Uacute: "\xDA",
  Ucirc: "\xDB",
  Uuml: "\xDC",
  Yacute: "\xDD",
  THORN: "\xDE",
  szlig: "\xDF",
  agrave: "\xE0",
  aacute: "\xE1",
  acirc: "\xE2",
  atilde: "\xE3",
  auml: "\xE4",
  aring: "\xE5",
  aelig: "\xE6",
  ccedil: "\xE7",
  egrave: "\xE8",
  eacute: "\xE9",
  ecirc: "\xEA",
  euml: "\xEB",
  igrave: "\xEC",
  iacute: "\xED",
  icirc: "\xEE",
  iuml: "\xEF",
  eth: "\xF0",
  ntilde: "\xF1",
  ograve: "\xF2",
  oacute: "\xF3",
  ocirc: "\xF4",
  otilde: "\xF5",
  ouml: "\xF6",
  divide: "\xF7",
  oslash: "\xF8",
  ugrave: "\xF9",
  uacute: "\xFA",
  ucirc: "\xFB",
  uuml: "\xFC",
  yacute: "\xFD",
  thorn: "\xFE",
  yuml: "\xFF",
  OElig: "\u0152",
  oelig: "\u0153",
  Scaron: "\u0160",
  scaron: "\u0161",
  Yuml: "\u0178",
  fnof: "\u0192",
  circ: "\u02C6",
  tilde: "\u02DC",
  Alpha: "\u0391",
  Beta: "\u0392",
  Gamma: "\u0393",
  Delta: "\u0394",
  Epsilon: "\u0395",
  Zeta: "\u0396",
  Eta: "\u0397",
  Theta: "\u0398",
  Iota: "\u0399",
  Kappa: "\u039A",
  Lambda: "\u039B",
  Mu: "\u039C",
  Nu: "\u039D",
  Xi: "\u039E",
  Omicron: "\u039F",
  Pi: "\u03A0",
  Rho: "\u03A1",
  Sigma: "\u03A3",
  Tau: "\u03A4",
  Upsilon: "\u03A5",
  Phi: "\u03A6",
  Chi: "\u03A7",
  Psi: "\u03A8",
  Omega: "\u03A9",
  alpha: "\u03B1",
  beta: "\u03B2",
  gamma: "\u03B3",
  delta: "\u03B4",
  epsilon: "\u03B5",
  zeta: "\u03B6",
  eta: "\u03B7",
  theta: "\u03B8",
  iota: "\u03B9",
  kappa: "\u03BA",
  lambda: "\u03BB",
  mu: "\u03BC",
  nu: "\u03BD",
  xi: "\u03BE",
  omicron: "\u03BF",
  pi: "\u03C0",
  rho: "\u03C1",
  sigmaf: "\u03C2",
  sigma: "\u03C3",
  tau: "\u03C4",
  upsilon: "\u03C5",
  phi: "\u03C6",
  chi: "\u03C7",
  psi: "\u03C8",
  omega: "\u03C9",
  thetasym: "\u03D1",
  upsih: "\u03D2",
  piv: "\u03D6",
  ensp: "\u2002",
  emsp: "\u2003",
  thinsp: "\u2009",
  zwnj: "\u200C",
  zwj: "\u200D",
  lrm: "\u200E",
  rlm: "\u200F",
  ndash: "\u2013",
  mdash: "\u2014",
  lsquo: "\u2018",
  rsquo: "\u2019",
  sbquo: "\u201A",
  ldquo: "\u201C",
  rdquo: "\u201D",
  bdquo: "\u201E",
  dagger: "\u2020",
  Dagger: "\u2021",
  bull: "\u2022",
  hellip: "\u2026",
  permil: "\u2030",
  prime: "\u2032",
  Prime: "\u2033",
  lsaquo: "\u2039",
  rsaquo: "\u203A",
  oline: "\u203E",
  frasl: "\u2044",
  euro: "\u20AC",
  image: "\u2111",
  weierp: "\u2118",
  real: "\u211C",
  trade: "\u2122",
  alefsym: "\u2135",
  larr: "\u2190",
  uarr: "\u2191",
  rarr: "\u2192",
  darr: "\u2193",
  harr: "\u2194",
  crarr: "\u21B5",
  lArr: "\u21D0",
  uArr: "\u21D1",
  rArr: "\u21D2",
  dArr: "\u21D3",
  hArr: "\u21D4",
  forall: "\u2200",
  part: "\u2202",
  exist: "\u2203",
  empty: "\u2205",
  nabla: "\u2207",
  isin: "\u2208",
  notin: "\u2209",
  ni: "\u220B",
  prod: "\u220F",
  sum: "\u2211",
  minus: "\u2212",
  lowast: "\u2217",
  radic: "\u221A",
  prop: "\u221D",
  infin: "\u221E",
  ang: "\u2220",
  and: "\u2227",
  or: "\u2228",
  cap: "\u2229",
  cup: "\u222A",
  int: "\u222B",
  there4: "\u2234",
  sim: "\u223C",
  cong: "\u2245",
  asymp: "\u2248",
  ne: "\u2260",
  equiv: "\u2261",
  le: "\u2264",
  ge: "\u2265",
  sub: "\u2282",
  sup: "\u2283",
  nsub: "\u2284",
  sube: "\u2286",
  supe: "\u2287",
  oplus: "\u2295",
  otimes: "\u2297",
  perp: "\u22A5",
  sdot: "\u22C5",
  lceil: "\u2308",
  rceil: "\u2309",
  lfloor: "\u230A",
  rfloor: "\u230B",
  lang: "\u2329",
  rang: "\u232A",
  loz: "\u25CA",
  spades: "\u2660",
  clubs: "\u2663",
  hearts: "\u2665",
  diams: "\u2666"
};

// src/extentions/jsx/index.ts
var hexNumber = /^[\da-fA-F]+$/;
var decimalNumber = /^\d+$/;
function getQualifiedJSXName(object) {
  if (!object) return object;
  if (object.type === "JSXIdentifier") return object.name;
  if (object.type === "JSXNamespacedName") return object.namespace.name + ":" + object.name.name;
  if (object.type === "JSXMemberExpression")
    return getQualifiedJSXName(object.object) + "." + getQualifiedJSXName(object.property);
}
function generateJsxParser(acorn, acornTypeScript, Parser, jsxOptions) {
  const tt = acorn.tokTypes;
  const tok = acornTypeScript.tokTypes;
  const isNewLine = acorn.isNewLine;
  const isIdentifierChar = acorn.isIdentifierChar;
  const options = Object.assign(
    {
      allowNamespaces: true,
      allowNamespacedObjects: true
    },
    jsxOptions || {}
  );
  return class JsxParser extends Parser {
    // Reads inline JSX contents token.
    jsx_readToken() {
      let out = "", chunkStart = this.pos;
      for (; ; ) {
        if (this.pos >= this.input.length) this.raise(this.start, "Unterminated JSX contents");
        let ch = this.input.charCodeAt(this.pos);
        switch (ch) {
          case 60:
          // '<'
          case 123:
            if (this.pos === this.start) {
              if (ch === 60 && this.exprAllowed) {
                ++this.pos;
                return this.finishToken(tok.jsxTagStart);
              }
              return this.getTokenFromCode(ch);
            }
            out += this.input.slice(chunkStart, this.pos);
            return this.finishToken(tok.jsxText, out);
          case 38:
            out += this.input.slice(chunkStart, this.pos);
            out += this.jsx_readEntity();
            chunkStart = this.pos;
            break;
          case 62:
          // '>'
          case 125:
            this.raise(
              this.pos,
              "Unexpected token `" + this.input[this.pos] + "`. Did you mean `" + (ch === 62 ? "&gt;" : "&rbrace;") + '` or `{"' + this.input[this.pos] + '"}`?'
            );
          default:
            if (isNewLine(ch)) {
              out += this.input.slice(chunkStart, this.pos);
              out += this.jsx_readNewLine(true);
              chunkStart = this.pos;
            } else {
              ++this.pos;
            }
        }
      }
    }
    jsx_readNewLine(normalizeCRLF) {
      let ch = this.input.charCodeAt(this.pos);
      let out;
      ++this.pos;
      if (ch === 13 && this.input.charCodeAt(this.pos) === 10) {
        ++this.pos;
        out = normalizeCRLF ? "\n" : "\r\n";
      } else {
        out = String.fromCharCode(ch);
      }
      if (this.options.locations) {
        ++this.curLine;
        this.lineStart = this.pos;
      }
      return out;
    }
    jsx_readString(quote) {
      let out = "", chunkStart = ++this.pos;
      for (; ; ) {
        if (this.pos >= this.input.length) this.raise(this.start, "Unterminated string constant");
        let ch = this.input.charCodeAt(this.pos);
        if (ch === quote) break;
        if (ch === 38) {
          out += this.input.slice(chunkStart, this.pos);
          out += this.jsx_readEntity();
          chunkStart = this.pos;
        } else if (isNewLine(ch)) {
          out += this.input.slice(chunkStart, this.pos);
          out += this.jsx_readNewLine(false);
          chunkStart = this.pos;
        } else {
          ++this.pos;
        }
      }
      out += this.input.slice(chunkStart, this.pos++);
      return this.finishToken(tt.string, out);
    }
    jsx_readEntity() {
      let str = "", count = 0, entity;
      let ch = this.input[this.pos];
      if (ch !== "&") this.raise(this.pos, "Entity must start with an ampersand");
      let startPos = ++this.pos;
      while (this.pos < this.input.length && count++ < 10) {
        ch = this.input[this.pos++];
        if (ch === ";") {
          if (str[0] === "#") {
            if (str[1] === "x") {
              str = str.substr(2);
              if (hexNumber.test(str)) entity = String.fromCharCode(parseInt(str, 16));
            } else {
              str = str.substr(1);
              if (decimalNumber.test(str)) entity = String.fromCharCode(parseInt(str, 10));
            }
          } else {
            entity = xhtml_default[str];
          }
          break;
        }
        str += ch;
      }
      if (!entity) {
        this.pos = startPos;
        return "&";
      }
      return entity;
    }
    // Read a JSX identifier (valid tag or attribute name).
    //
    // Optimized version since JSX identifiers can't contain
    // escape characters and so can be read as single slice.
    // Also assumes that first character was already checked
    // by isIdentifierStart in readToken.
    jsx_readWord() {
      let ch, start = this.pos;
      do {
        ch = this.input.charCodeAt(++this.pos);
      } while (isIdentifierChar(ch) || ch === 45);
      return this.finishToken(tok.jsxName, this.input.slice(start, this.pos));
    }
    // Parse next token as JSX identifier
    jsx_parseIdentifier() {
      let node = this.startNode();
      if (this.type === tok.jsxName) node.name = this.value;
      else if (this.type.keyword) node.name = this.type.keyword;
      else this.unexpected();
      this.next();
      return this.finishNode(node, "JSXIdentifier");
    }
    // Parse namespaced identifier.
    jsx_parseNamespacedName() {
      let startPos = this.start, startLoc = this.startLoc;
      let name = this.jsx_parseIdentifier();
      if (!options.allowNamespaces || !this.eat(tt.colon)) return name;
      var node = this.startNodeAt(startPos, startLoc);
      node.namespace = name;
      node.name = this.jsx_parseIdentifier();
      return this.finishNode(node, "JSXNamespacedName");
    }
    // Parses element name in any form - namespaced, member
    // or single identifier.
    jsx_parseElementName() {
      if (this.type === tok.jsxTagEnd) return "";
      let startPos = this.start, startLoc = this.startLoc;
      let node = this.jsx_parseNamespacedName();
      if (this.type === tt.dot && node.type === "JSXNamespacedName" && !options.allowNamespacedObjects) {
        this.unexpected();
      }
      while (this.eat(tt.dot)) {
        let newNode = this.startNodeAt(startPos, startLoc);
        newNode.object = node;
        newNode.property = this.jsx_parseIdentifier();
        node = this.finishNode(newNode, "JSXMemberExpression");
      }
      return node;
    }
    // Parses any type of JSX attribute value.
    jsx_parseAttributeValue() {
      switch (this.type) {
        case tt.braceL:
          let node = this.jsx_parseExpressionContainer();
          if (node.expression.type === "JSXEmptyExpression")
            this.raise(node.start, "JSX attributes must only be assigned a non-empty expression");
          return node;
        case tok.jsxTagStart:
        case tt.string:
          return this.parseExprAtom();
        default:
          this.raise(this.start, "JSX value should be either an expression or a quoted JSX text");
      }
    }
    // JSXEmptyExpression is unique type since it doesn't actually parse anything,
    // and so it should start at the end of last read token (left brace) and finish
    // at the beginning of the next one (right brace).
    jsx_parseEmptyExpression() {
      let node = this.startNodeAt(this.lastTokEnd, this.lastTokEndLoc);
      return this.finishNodeAt(node, "JSXEmptyExpression", this.start, this.startLoc);
    }
    // Parses JSX expression enclosed into curly brackets.
    jsx_parseExpressionContainer() {
      let node = this.startNode();
      this.next();
      node.expression = this.type === tt.braceR ? this.jsx_parseEmptyExpression() : this.parseExpression();
      this.expect(tt.braceR);
      return this.finishNode(node, "JSXExpressionContainer");
    }
    // Parses following JSX attribute name-value pair.
    jsx_parseAttribute() {
      let node = this.startNode();
      if (this.eat(tt.braceL)) {
        this.expect(tt.ellipsis);
        node.argument = this.parseMaybeAssign();
        this.expect(tt.braceR);
        return this.finishNode(node, "JSXSpreadAttribute");
      }
      node.name = this.jsx_parseNamespacedName();
      node.value = this.eat(tt.eq) ? this.jsx_parseAttributeValue() : null;
      return this.finishNode(node, "JSXAttribute");
    }
    // Parses JSX opening tag starting after '<'.
    jsx_parseOpeningElementAt(startPos, startLoc) {
      let node = this.startNodeAt(startPos, startLoc);
      node.attributes = [];
      let nodeName = this.jsx_parseElementName();
      if (nodeName) node.name = nodeName;
      while (this.type !== tt.slash && this.type !== tok.jsxTagEnd)
        node.attributes.push(this.jsx_parseAttribute());
      node.selfClosing = this.eat(tt.slash);
      this.expect(tok.jsxTagEnd);
      return this.finishNode(node, nodeName ? "JSXOpeningElement" : "JSXOpeningFragment");
    }
    // Parses JSX closing tag starting after '</'.
    jsx_parseClosingElementAt(startPos, startLoc) {
      let node = this.startNodeAt(startPos, startLoc);
      let nodeName = this.jsx_parseElementName();
      if (nodeName) node.name = nodeName;
      this.expect(tok.jsxTagEnd);
      return this.finishNode(node, nodeName ? "JSXClosingElement" : "JSXClosingFragment");
    }
    // Parses entire JSX element, including it's opening tag
    // (starting after '<'), attributes, contents and closing tag.
    jsx_parseElementAt(startPos, startLoc) {
      let node = this.startNodeAt(startPos, startLoc);
      let children = [];
      let openingElement = this.jsx_parseOpeningElementAt(startPos, startLoc);
      let closingElement = null;
      if (!openingElement.selfClosing) {
        contents: for (; ; ) {
          switch (this.type) {
            case tok.jsxTagStart:
              startPos = this.start;
              startLoc = this.startLoc;
              this.next();
              if (this.eat(tt.slash)) {
                closingElement = this.jsx_parseClosingElementAt(startPos, startLoc);
                break contents;
              }
              children.push(this.jsx_parseElementAt(startPos, startLoc));
              break;
            case tok.jsxText:
              children.push(this.parseExprAtom());
              break;
            case tt.braceL:
              children.push(this.jsx_parseExpressionContainer());
              break;
            default:
              this.unexpected();
          }
        }
        if (getQualifiedJSXName(closingElement.name) !== getQualifiedJSXName(openingElement.name)) {
          this.raise(
            closingElement.start,
            "Expected corresponding JSX closing tag for <" + getQualifiedJSXName(openingElement.name) + ">"
          );
        }
      }
      let fragmentOrElement = openingElement.name ? "Element" : "Fragment";
      node["opening" + fragmentOrElement] = openingElement;
      node["closing" + fragmentOrElement] = closingElement;
      node.children = children;
      if (this.type === tt.relational && this.value === "<") {
        this.raise(this.start, "Adjacent JSX elements must be wrapped in an enclosing tag");
      }
      return this.finishNode(node, "JSX" + fragmentOrElement);
    }
    // Parse JSX text
    jsx_parseText() {
      let node = this.parseLiteral(this.value);
      node.type = "JSXText";
      return node;
    }
    // Parses entire JSX element from current position.
    jsx_parseElement() {
      let startPos = this.start, startLoc = this.startLoc;
      this.next();
      return this.jsx_parseElementAt(startPos, startLoc);
    }
  };
}

// src/extentions/import-assertions.ts
function generateParseImportAssertions(Parse, acornTypeScript, acorn) {
  const { tokTypes: tokTypes2 } = acornTypeScript;
  const { tokTypes: tt } = acorn;
  return class ImportAttributes extends Parse {
    parseMaybeImportAttributes(node) {
      if (this.type === tt._with || this.type === tokTypes2.assert) {
        this.next();
        const attributes = this.parseImportAttributes();
        if (attributes) {
          node.attributes = attributes;
        }
      }
    }
    parseImportAttributes() {
      this.expect(tt.braceL);
      const attrs = this.parseWithEntries();
      this.expect(tt.braceR);
      return attrs;
    }
    parseWithEntries() {
      const attrs = [];
      const attrNames = /* @__PURE__ */ new Set();
      do {
        if (this.type === tt.braceR) {
          break;
        }
        const node = this.startNode();
        let withionKeyNode;
        if (this.type === tt.string) {
          withionKeyNode = this.parseLiteral(this.value);
        } else {
          withionKeyNode = this.parseIdent(true);
        }
        this.next();
        node.key = withionKeyNode;
        if (attrNames.has(node.key.name)) {
          this.raise(this.pos, "Duplicated key in attributes");
        }
        attrNames.add(node.key.name);
        if (this.type !== tt.string) {
          this.raise(this.pos, "Only string is supported as an attribute value");
        }
        node.value = this.parseLiteral(this.value);
        attrs.push(this.finishNode(node, "ImportAttribute"));
      } while (this.eat(tt.comma));
      return attrs;
    }
  };
}

// src/index.ts
var skipWhiteSpace = /(?:\s|\/\/.*|\/\*[^]*?\*\/)*/g;
function assert(x) {
  if (!x) {
    throw new Error("Assert fail");
  }
}
function tsIsClassAccessor(modifier) {
  return modifier === "accessor";
}
function tsIsVarianceAnnotations(modifier) {
  return modifier === "in" || modifier === "out";
}
var FUNC_STATEMENT = 1;
var FUNC_HANGING_STATEMENT = 2;
var FUNC_NULLABLE_ID = 4;
var acornScope = {
  SCOPE_TOP: 1,
  SCOPE_FUNCTION: 2,
  SCOPE_ASYNC: 4,
  SCOPE_GENERATOR: 8,
  SCOPE_ARROW: 16,
  SCOPE_SIMPLE_CATCH: 32,
  SCOPE_SUPER: 64,
  SCOPE_DIRECT_SUPER: 128,
  SCOPE_CLASS_STATIC_BLOCK: 256,
  SCOPE_VAR: 256,
  BIND_NONE: 0,
  // Not a binding
  BIND_VAR: 1,
  // Var-style binding
  BIND_LEXICAL: 2,
  // Let- or const-style binding
  BIND_FUNCTION: 3,
  // Function declaration
  BIND_SIMPLE_CATCH: 4,
  // Simple (identifier pattern) catch binding
  BIND_OUTSIDE: 5,
  // Special case for function names as bound inside the
  BIND_TS_TYPE: 6,
  BIND_TS_INTERFACE: 7,
  BIND_TS_NAMESPACE: 8,
  BIND_FLAGS_TS_EXPORT_ONLY: 1024,
  BIND_FLAGS_TS_IMPORT: 4096,
  BIND_FLAGS_TS_ENUM: 256,
  BIND_FLAGS_TS_CONST_ENUM: 512,
  BIND_FLAGS_CLASS: 128
  // function
};
function functionFlags(async, generator) {
  return acornScope.SCOPE_FUNCTION | (async ? acornScope.SCOPE_ASYNC : 0) | (generator ? acornScope.SCOPE_GENERATOR : 0);
}
function isPossiblyLiteralEnum(expression) {
  if (expression.type !== "MemberExpression") return false;
  const { computed, property } = expression;
  if (computed && (property.type !== "TemplateLiteral" || property.expressions.length > 0)) {
    return false;
  }
  return isUncomputedMemberExpressionChain(expression.object);
}
function isUncomputedMemberExpressionChain(expression) {
  if (expression.type === "Identifier") return true;
  if (expression.type !== "MemberExpression") return false;
  if (expression.computed) return false;
  return isUncomputedMemberExpressionChain(expression.object);
}
function tsIsAccessModifier(modifier) {
  return modifier === "private" || modifier === "public" || modifier === "protected";
}
function tokenCanStartExpression(token) {
  return Boolean(token.startsExpr);
}
function nonNull(x) {
  if (x == null) {
    throw new Error(`Unexpected ${x} value.`);
  }
  return x;
}
function keywordTypeFromName(value) {
  switch (value) {
    case "any":
      return "TSAnyKeyword";
    case "boolean":
      return "TSBooleanKeyword";
    case "bigint":
      return "TSBigIntKeyword";
    case "never":
      return "TSNeverKeyword";
    case "number":
      return "TSNumberKeyword";
    case "object":
      return "TSObjectKeyword";
    case "string":
      return "TSStringKeyword";
    case "symbol":
      return "TSSymbolKeyword";
    case "undefined":
      return "TSUndefinedKeyword";
    case "unknown":
      return "TSUnknownKeyword";
    default:
      return void 0;
  }
}
function tsPlugin(options) {
  const { dts = false } = options || {};
  const disallowAmbiguousJSXLike = !!options?.jsx;
  return function(Parser) {
    const _acorn = Parser.acorn || acornNamespace;
    const acornTypeScript = generateAcornTypeScript(_acorn);
    const tt = _acorn.tokTypes;
    const keywordTypes2 = _acorn.keywordTypes;
    const isIdentifierStart = _acorn.isIdentifierStart;
    const lineBreak = _acorn.lineBreak;
    const isNewLine = _acorn.isNewLine;
    const tokContexts = _acorn.tokContexts;
    const isIdentifierChar = _acorn.isIdentifierChar;
    const {
      tokTypes: tokTypes2,
      tokContexts: tsTokContexts,
      keywordsRegExp,
      tokenIsLiteralPropertyName,
      tokenIsTemplate,
      tokenIsTSDeclarationStart,
      tokenIsIdentifier,
      tokenIsKeywordOrIdentifier,
      tokenIsTSTypeOperator
    } = acornTypeScript;
    function nextLineBreak(code, from, end = code.length) {
      for (let i = from; i < end; i++) {
        let next = code.charCodeAt(i);
        if (isNewLine(next))
          return i < end - 1 && next === 13 && code.charCodeAt(i + 1) === 10 ? i + 2 : i + 1;
      }
      return -1;
    }
    Parser = generateParseDecorators(Parser, acornTypeScript, _acorn);
    if (options?.jsx) {
      Parser = generateJsxParser(
        _acorn,
        acornTypeScript,
        Parser,
        typeof options.jsx === "boolean" ? {} : options.jsx
      );
    }
    Parser = generateParseImportAssertions(Parser, acornTypeScript, _acorn);
    class TypeScriptParser extends Parser {
      constructor(options2, input, startPos) {
        super(options2, input, startPos);
        this.preValue = null;
        this.preToken = null;
        this.isLookahead = false;
        this.isAmbientContext = false;
        this.inAbstractClass = false;
        this.inType = false;
        this.inDisallowConditionalTypesContext = false;
        this.maybeInArrowParameters = false;
        this.shouldParseArrowReturnType = void 0;
        this.shouldParseAsyncArrowReturnType = void 0;
        this.decoratorStack = [[]];
        this.importsStack = [[]];
        /**
         * we will only parse one import node or export node at same time.
         * default kind is undefined
         * */
        this.importOrExportOuterKind = void 0;
        this.tsParseConstModifier = (node) => {
          this.tsParseModifiers({
            modified: node,
            allowedModifiers: ["const"],
            // for better error recovery
            disallowedModifiers: ["in", "out"],
            errorTemplate: TypeScriptError.InvalidModifierOnTypeParameterPositions
          });
        };
        this.ecmaVersion = this.options.ecmaVersion;
      }
      // support in Class static
      static get acornTypeScript() {
        return acornTypeScript;
      }
      // support in runtime, get acornTypeScript be this
      get acornTypeScript() {
        return acornTypeScript;
      }
      getTokenFromCodeInType(code) {
        if (code === 62) {
          return this.finishOp(tt.relational, 1);
        }
        if (code === 60) {
          return this.finishOp(tt.relational, 1);
        }
        return super.getTokenFromCode(code);
      }
      readToken(code) {
        if (!this.inType) {
          let context = this.curContext();
          if (context === tsTokContexts.tc_expr) return this.jsx_readToken();
          if (context === tsTokContexts.tc_oTag || context === tsTokContexts.tc_cTag) {
            if (isIdentifierStart(code)) return this.jsx_readWord();
            if (code == 62) {
              ++this.pos;
              return this.finishToken(tokTypes2.jsxTagEnd);
            }
            if ((code === 34 || code === 39) && context == tsTokContexts.tc_oTag)
              return this.jsx_readString(code);
          }
          if (code === 60 && this.exprAllowed && this.input.charCodeAt(this.pos + 1) !== 33) {
            ++this.pos;
            if (options?.jsx) {
              return this.finishToken(tokTypes2.jsxTagStart);
            } else {
              return this.finishToken(tt.relational, "<");
            }
          }
        }
        return super.readToken(code);
      }
      getTokenFromCode(code) {
        if (this.inType) {
          return this.getTokenFromCodeInType(code);
        }
        if (code === 64) {
          ++this.pos;
          return this.finishToken(tokTypes2.at);
        }
        return super.getTokenFromCode(code);
      }
      isAbstractClass() {
        return this.ts_isContextual(tokTypes2.abstract) && this.lookahead().type === tt._class;
      }
      finishNode(node, type) {
        if (node.type !== "" && node.end !== 0) {
          return node;
        }
        return super.finishNode(node, type);
      }
      // tryParse will clone parser state.
      // It is expensive and should be used with cautions
      tryParse(fn, oldState = this.cloneCurLookaheadState()) {
        const abortSignal = { node: null };
        try {
          const node = fn((node2 = null) => {
            abortSignal.node = node2;
            throw abortSignal;
          });
          return {
            node,
            error: null,
            thrown: false,
            aborted: false,
            failState: null
          };
        } catch (error) {
          const failState = this.getCurLookaheadState();
          this.setLookaheadState(oldState);
          if (error instanceof SyntaxError) {
            return {
              node: null,
              error,
              thrown: true,
              aborted: false,
              failState
            };
          }
          if (error === abortSignal) {
            return {
              node: abortSignal.node,
              error: null,
              thrown: false,
              aborted: true,
              failState
            };
          }
          throw error;
        }
      }
      setOptionalParametersError(refExpressionErrors, resultError) {
        refExpressionErrors.optionalParametersLoc = resultError?.loc ?? this.startLoc;
      }
      // used after we have finished parsing types
      reScan_lt_gt() {
        if (this.type === tt.relational) {
          this.pos -= 1;
          this.readToken_lt_gt(this.fullCharCodeAtPos());
        }
      }
      reScan_lt() {
        const { type } = this;
        if (type === tt.bitShift) {
          this.pos -= 2;
          this.finishOp(tt.relational, 1);
          return tt.relational;
        }
        return type;
      }
      resetEndLocation(node, endPos = this.lastTokEnd, endLoc = this.lastTokEndLoc) {
        node.end = endPos;
        node.loc.end = endLoc;
        if (this.options.ranges) node.range[1] = endPos;
      }
      startNodeAtNode(type) {
        return super.startNodeAt(type.start, type.loc.start);
      }
      nextTokenStart() {
        return this.nextTokenStartSince(this.pos);
      }
      tsHasSomeModifiers(member, modifiers) {
        return modifiers.some((modifier) => {
          if (tsIsAccessModifier(modifier)) {
            return member.accessibility === modifier;
          }
          return !!member[modifier];
        });
      }
      tsIsStartOfStaticBlocks() {
        return this.isContextual("static") && this.lookaheadCharCode() === 123;
      }
      tsCheckForInvalidTypeCasts(items) {
        items.forEach((node) => {
          if (node?.type === "TSTypeCastExpression") {
            this.raise(node.typeAnnotation.start, TypeScriptError.UnexpectedTypeAnnotation);
          }
        });
      }
      atPossibleAsyncArrow(base) {
        return base.type === "Identifier" && base.name === "async" && this.lastTokEnd === base.end && !this.canInsertSemicolon() && base.end - base.start === 5 && base.start === this.potentialArrowAt;
      }
      tsIsIdentifier() {
        return tokenIsIdentifier(this.type);
      }
      tsTryParseTypeOrTypePredicateAnnotation() {
        return this.match(tt.colon) ? this.tsParseTypeOrTypePredicateAnnotation(tt.colon) : void 0;
      }
      tsTryParseGenericAsyncArrowFunction(startPos, startLoc, forInit) {
        if (!this.tsMatchLeftRelational()) {
          return void 0;
        }
        const oldMaybeInArrowParameters = this.maybeInArrowParameters;
        this.maybeInArrowParameters = true;
        const res = this.tsTryParseAndCatch(() => {
          const node = this.startNodeAt(startPos, startLoc);
          node.typeParameters = this.tsParseTypeParameters(this.tsParseConstModifier);
          super.parseFunctionParams(node);
          node.returnType = this.tsTryParseTypeOrTypePredicateAnnotation();
          this.expect(tt.arrow);
          return node;
        });
        this.maybeInArrowParameters = oldMaybeInArrowParameters;
        if (!res) {
          return void 0;
        }
        return super.parseArrowExpression(
          res,
          /* params are already set */
          null,
          /* async */
          true,
          /* forInit */
          forInit
        );
      }
      // Used when parsing type arguments from ES productions, where the first token
      // has been created without state.inType. Thus we need to rescan the lt token.
      tsParseTypeArgumentsInExpression() {
        if (this.reScan_lt() !== tt.relational) {
          return void 0;
        }
        return this.tsParseTypeArguments();
      }
      tsInNoContext(cb) {
        const oldContext = this.context;
        this.context = [oldContext[0]];
        try {
          return cb();
        } finally {
          this.context = oldContext;
        }
      }
      tsTryParseTypeAnnotation() {
        return this.match(tt.colon) ? this.tsParseTypeAnnotation() : void 0;
      }
      isUnparsedContextual(nameStart, name) {
        const nameEnd = nameStart + name.length;
        if (this.input.slice(nameStart, nameEnd) === name) {
          const nextCh = this.input.charCodeAt(nameEnd);
          return !(isIdentifierChar(nextCh) || // check if `nextCh is between 0xd800 - 0xdbff,
          // if `nextCh` is NaN, `NaN & 0xfc00` is 0, the function
          // returns true
          (nextCh & 64512) === 55296);
        }
        return false;
      }
      isAbstractConstructorSignature() {
        return this.ts_isContextual(tokTypes2.abstract) && this.lookahead().type === tt._new;
      }
      nextTokenStartSince(pos) {
        skipWhiteSpace.lastIndex = pos;
        return skipWhiteSpace.test(this.input) ? skipWhiteSpace.lastIndex : pos;
      }
      lookaheadCharCode() {
        return this.input.charCodeAt(this.nextTokenStart());
      }
      compareLookaheadState(state, state2) {
        for (const key of Object.keys(state)) {
          if (state[key] !== state2[key]) return false;
        }
        return true;
      }
      createLookaheadState() {
        this.value = null;
        this.context = [this.curContext()];
      }
      getCurLookaheadState() {
        return {
          endLoc: this.endLoc,
          lastTokEnd: this.lastTokEnd,
          lastTokStart: this.lastTokStart,
          lastTokStartLoc: this.lastTokStartLoc,
          pos: this.pos,
          value: this.value,
          type: this.type,
          start: this.start,
          end: this.end,
          context: this.context,
          startLoc: this.startLoc,
          lastTokEndLoc: this.lastTokEndLoc,
          curLine: this.curLine,
          lineStart: this.lineStart,
          curPosition: this.curPosition,
          containsEsc: this.containsEsc
        };
      }
      cloneCurLookaheadState() {
        return {
          pos: this.pos,
          value: this.value,
          type: this.type,
          start: this.start,
          end: this.end,
          context: this.context && this.context.slice(),
          startLoc: this.startLoc,
          lastTokEndLoc: this.lastTokEndLoc,
          endLoc: this.endLoc,
          lastTokEnd: this.lastTokEnd,
          lastTokStart: this.lastTokStart,
          lastTokStartLoc: this.lastTokStartLoc,
          curLine: this.curLine,
          lineStart: this.lineStart,
          curPosition: this.curPosition,
          containsEsc: this.containsEsc
        };
      }
      setLookaheadState(state) {
        this.pos = state.pos;
        this.value = state.value;
        this.endLoc = state.endLoc;
        this.lastTokEnd = state.lastTokEnd;
        this.lastTokStart = state.lastTokStart;
        this.lastTokStartLoc = state.lastTokStartLoc;
        this.type = state.type;
        this.start = state.start;
        this.end = state.end;
        this.context = state.context;
        this.startLoc = state.startLoc;
        this.lastTokEndLoc = state.lastTokEndLoc;
        this.curLine = state.curLine;
        this.lineStart = state.lineStart;
        this.curPosition = state.curPosition;
        this.containsEsc = state.containsEsc;
      }
      // Utilities
      tsLookAhead(f) {
        const state = this.getCurLookaheadState();
        const res = f();
        this.setLookaheadState(state);
        return res;
      }
      lookahead(number) {
        const oldState = this.getCurLookaheadState();
        this.createLookaheadState();
        this.isLookahead = true;
        if (number !== void 0) {
          for (let i = 0; i < number; i++) {
            this.nextToken();
          }
        } else {
          this.nextToken();
        }
        this.isLookahead = false;
        const curState = this.getCurLookaheadState();
        this.setLookaheadState(oldState);
        return curState;
      }
      readWord() {
        let word = this.readWord1();
        let type = tt.name;
        if (this.keywords.test(word)) {
          type = keywordTypes2[word];
        } else if (new RegExp(keywordsRegExp).test(word)) {
          type = tokTypes2[word];
        }
        return this.finishToken(type, word);
      }
      skipBlockComment() {
        let startLoc;
        if (!this.isLookahead) startLoc = this.options.onComment && this.curPosition();
        let start = this.pos, end = this.input.indexOf("*/", this.pos += 2);
        if (end === -1) this.raise(this.pos - 2, "Unterminated comment");
        this.pos = end + 2;
        if (this.options.locations) {
          for (let nextBreak, pos = start; (nextBreak = nextLineBreak(this.input, pos, this.pos)) > -1; ) {
            ++this.curLine;
            pos = this.lineStart = nextBreak;
          }
        }
        if (this.isLookahead) return;
        if (this.options.onComment) {
          this.options.onComment(
            true,
            this.input.slice(start + 2, end),
            start,
            this.pos,
            startLoc,
            this.curPosition()
          );
        }
      }
      skipLineComment(startSkip) {
        let start = this.pos;
        let startLoc;
        if (!this.isLookahead) startLoc = this.options.onComment && this.curPosition();
        let ch = this.input.charCodeAt(this.pos += startSkip);
        while (this.pos < this.input.length && !isNewLine(ch)) {
          ch = this.input.charCodeAt(++this.pos);
        }
        if (this.isLookahead) return;
        if (this.options.onComment)
          this.options.onComment(
            false,
            this.input.slice(start + startSkip, this.pos),
            start,
            this.pos,
            startLoc,
            this.curPosition()
          );
      }
      finishToken(type, val) {
        this.preValue = this.value;
        this.preToken = this.type;
        this.end = this.pos;
        if (this.options.locations) this.endLoc = this.curPosition();
        let prevType = this.type;
        this.type = type;
        this.value = val;
        if (!this.isLookahead) {
          this.updateContext(prevType);
        }
      }
      resetStartLocation(node, start, startLoc) {
        node.start = start;
        node.loc.start = startLoc;
        if (this.options.ranges) node.range[0] = start;
      }
      isLineTerminator() {
        return this.eat(tt.semi) || super.canInsertSemicolon();
      }
      hasFollowingLineBreak() {
        skipWhiteSpaceToLineBreak.lastIndex = this.end;
        return skipWhiteSpaceToLineBreak.test(this.input);
      }
      addExtra(node, key, value, enumerable = true) {
        if (!node) return;
        const extra = node.extra = node.extra || {};
        if (enumerable) {
          extra[key] = value;
        } else {
          Object.defineProperty(extra, key, { enumerable, value });
        }
      }
      /**
       * Test if current token is a literal property name
       * https://tc39.es/ecma262/#prod-LiteralPropertyName
       * LiteralPropertyName:
       *   IdentifierName
       *   StringLiteral
       *   NumericLiteral
       *   BigIntLiteral
       */
      isLiteralPropertyName() {
        return tokenIsLiteralPropertyName(this.type);
      }
      hasPrecedingLineBreak() {
        return lineBreak.test(this.input.slice(this.lastTokEnd, this.start));
      }
      createIdentifier(node, name) {
        node.name = name;
        return this.finishNode(node, "Identifier");
      }
      /**
       * Reset the start location of node to the start location of locationNode
       */
      resetStartLocationFromNode(node, locationNode) {
        this.resetStartLocation(node, locationNode.start, locationNode.loc.start);
      }
      // This is used in flow and typescript plugin
      // Determine whether a parameter is a this param
      isThisParam(param) {
        return param.type === "Identifier" && param.name === "this";
      }
      isLookaheadContextual(name) {
        const next = this.nextTokenStart();
        return this.isUnparsedContextual(next, name);
      }
      /**
       * ts type isContextual
       * @param {TokenType} type
       * @param {TokenType} token
       * @returns {boolean}
       * */
      ts_type_isContextual(type, token) {
        return type === token && !this.containsEsc;
      }
      /**
       * ts isContextual
       * @param {TokenType} token
       * @returns {boolean}
       * */
      ts_isContextual(token) {
        return this.type === token && !this.containsEsc;
      }
      ts_isContextualWithState(state, token) {
        return state.type === token && !state.containsEsc;
      }
      isContextualWithState(keyword, state) {
        return state.type === tt.name && state.value === keyword && !state.containsEsc;
      }
      tsIsStartOfMappedType() {
        this.next();
        if (this.eat(tt.plusMin)) {
          return this.ts_isContextual(tokTypes2.readonly);
        }
        if (this.ts_isContextual(tokTypes2.readonly)) {
          this.next();
        }
        if (!this.match(tt.bracketL)) {
          return false;
        }
        this.next();
        if (!this.tsIsIdentifier()) {
          return false;
        }
        this.next();
        return this.match(tt._in);
      }
      tsInDisallowConditionalTypesContext(cb) {
        const oldInDisallowConditionalTypesContext = this.inDisallowConditionalTypesContext;
        this.inDisallowConditionalTypesContext = true;
        try {
          return cb();
        } finally {
          this.inDisallowConditionalTypesContext = oldInDisallowConditionalTypesContext;
        }
      }
      tsTryParseType() {
        return this.tsEatThenParseType(tt.colon);
      }
      /**
       * Whether current token matches given type
       *
       * @param {TokenType} type
       * @returns {boolean}
       * @memberof Tokenizer
       */
      match(type) {
        return this.type === type;
      }
      matchJsx(type) {
        return this.type === acornTypeScript.tokTypes[type];
      }
      ts_eatWithState(type, nextCount, state) {
        const targetType = state.type;
        if (type === targetType) {
          for (let i = 0; i < nextCount; i++) {
            this.next();
          }
          return true;
        } else {
          return false;
        }
      }
      ts_eatContextualWithState(name, nextCount, state) {
        if (keywordsRegExp.test(name)) {
          if (this.ts_isContextualWithState(state, tokTypes2[name])) {
            for (let i = 0; i < nextCount; i++) {
              this.next();
            }
            return true;
          }
          return false;
        } else {
          if (!this.isContextualWithState(name, state)) return false;
          for (let i = 0; i < nextCount; i++) {
            this.next();
          }
          return true;
        }
      }
      canHaveLeadingDecorator() {
        return this.match(tt._class);
      }
      eatContextual(name) {
        if (keywordsRegExp.test(name)) {
          if (this.ts_isContextual(tokTypes2[name])) {
            this.next();
            return true;
          }
          return false;
        } else {
          return super.eatContextual(name);
        }
      }
      tsIsExternalModuleReference() {
        return this.isContextual("require") && this.lookaheadCharCode() === 40;
      }
      tsParseExternalModuleReference() {
        const node = this.startNode();
        this.expectContextual("require");
        this.expect(tt.parenL);
        if (!this.match(tt.string)) {
          this.unexpected();
        }
        node.expression = this.parseExprAtom();
        this.expect(tt.parenR);
        return this.finishNode(node, "TSExternalModuleReference");
      }
      tsParseEntityName(allowReservedWords = true) {
        let entity = this.parseIdent(allowReservedWords);
        while (this.eat(tt.dot)) {
          const node = this.startNodeAtNode(entity);
          node.left = entity;
          node.right = this.parseIdent(allowReservedWords);
          entity = this.finishNode(node, "TSQualifiedName");
        }
        return entity;
      }
      tsParseEnumMember() {
        const node = this.startNode();
        node.id = this.match(tt.string) ? this.parseLiteral(this.value) : this.parseIdent(
          /* liberal */
          true
        );
        if (this.eat(tt.eq)) {
          node.initializer = this.parseMaybeAssign();
        }
        return this.finishNode(node, "TSEnumMember");
      }
      tsParseEnumDeclaration(node, properties = {}) {
        if (properties.const) node.const = true;
        if (properties.declare) node.declare = true;
        this.expectContextual("enum");
        node.id = this.parseIdent();
        this.checkLValSimple(node.id);
        this.expect(tt.braceL);
        node.members = this.tsParseDelimitedList("EnumMembers", this.tsParseEnumMember.bind(this));
        this.expect(tt.braceR);
        return this.finishNode(node, "TSEnumDeclaration");
      }
      tsParseModuleBlock() {
        const node = this.startNode();
        this.enterScope(TS_SCOPE_OTHER);
        this.expect(tt.braceL);
        node.body = [];
        while (this.type !== tt.braceR) {
          let stmt = this.parseStatement(null, true);
          node.body.push(stmt);
        }
        this.next();
        super.exitScope();
        return this.finishNode(node, "TSModuleBlock");
      }
      tsParseAmbientExternalModuleDeclaration(node) {
        if (this.ts_isContextual(tokTypes2.global)) {
          node.global = true;
          node.id = this.parseIdent();
        } else if (this.match(tt.string)) {
          node.id = this.parseLiteral(this.value);
        } else {
          this.unexpected();
        }
        if (this.match(tt.braceL)) {
          this.enterScope(TS_SCOPE_TS_MODULE);
          node.body = this.tsParseModuleBlock();
          super.exitScope();
        } else {
          super.semicolon();
        }
        return this.finishNode(node, "TSModuleDeclaration");
      }
      tsTryParseDeclare(nany) {
        if (this.isLineTerminator()) {
          return;
        }
        let starttype = this.type;
        let kind;
        if (this.isContextual("let")) {
          starttype = tt._var;
          kind = "let";
        }
        return this.tsInAmbientContext(() => {
          if (starttype === tt._function) {
            nany.declare = true;
            return this.parseFunctionStatement(
              nany,
              /* async */
              false,
              /* declarationPosition */
              true
            );
          }
          if (starttype === tt._class) {
            nany.declare = true;
            return this.parseClass(nany, true);
          }
          if (starttype === tokTypes2.enum) {
            return this.tsParseEnumDeclaration(nany, { declare: true });
          }
          if (starttype === tokTypes2.global) {
            return this.tsParseAmbientExternalModuleDeclaration(nany);
          }
          if (starttype === tt._const || starttype === tt._var) {
            if (!this.match(tt._const) || !this.isLookaheadContextual("enum")) {
              nany.declare = true;
              return this.parseVarStatement(nany, kind || this.value, true);
            }
            this.expect(tt._const);
            return this.tsParseEnumDeclaration(nany, {
              const: true,
              declare: true
            });
          }
          if (starttype === tokTypes2.interface) {
            const result = this.tsParseInterfaceDeclaration(nany, {
              declare: true
            });
            if (result) return result;
          }
          if (tokenIsIdentifier(starttype)) {
            return this.tsParseDeclaration(
              nany,
              this.value,
              /* next */
              true
            );
          }
        });
      }
      tsIsListTerminator(kind) {
        switch (kind) {
          case "EnumMembers":
          case "TypeMembers":
            return this.match(tt.braceR);
          case "HeritageClauseElement":
            return this.match(tt.braceL);
          case "TupleElementTypes":
            return this.match(tt.bracketR);
          case "TypeParametersOrArguments":
            return this.tsMatchRightRelational();
        }
      }
      /**
       * If !expectSuccess, returns undefined instead of failing to parse.
       * If expectSuccess, parseElement should always return a defined value.
       */
      tsParseDelimitedListWorker(kind, parseElement, expectSuccess, refTrailingCommaPos) {
        const result = [];
        let trailingCommaPos = -1;
        for (; ; ) {
          if (this.tsIsListTerminator(kind)) {
            break;
          }
          trailingCommaPos = -1;
          const element = parseElement();
          if (element == null) {
            return void 0;
          }
          result.push(element);
          if (this.eat(tt.comma)) {
            trailingCommaPos = this.lastTokStart;
            continue;
          }
          if (this.tsIsListTerminator(kind)) {
            break;
          }
          if (expectSuccess) {
            this.expect(tt.comma);
          }
          return void 0;
        }
        if (refTrailingCommaPos) {
          refTrailingCommaPos.value = trailingCommaPos;
        }
        return result;
      }
      tsParseDelimitedList(kind, parseElement, refTrailingCommaPos) {
        return nonNull(
          this.tsParseDelimitedListWorker(
            kind,
            parseElement,
            /* expectSuccess */
            true,
            refTrailingCommaPos
          )
        );
      }
      tsParseBracketedList(kind, parseElement, bracket, skipFirstToken, refTrailingCommaPos) {
        if (!skipFirstToken) {
          if (bracket) {
            this.expect(tt.bracketL);
          } else {
            this.expect(tt.relational);
          }
        }
        const result = this.tsParseDelimitedList(kind, parseElement, refTrailingCommaPos);
        if (bracket) {
          this.expect(tt.bracketR);
        } else {
          this.expect(tt.relational);
        }
        return result;
      }
      tsParseTypeParameterName() {
        const typeName = this.parseIdent();
        return typeName.name;
      }
      tsEatThenParseType(token) {
        return !this.match(token) ? void 0 : this.tsNextThenParseType();
      }
      tsExpectThenParseType(token) {
        return this.tsDoThenParseType(() => this.expect(token));
      }
      tsNextThenParseType() {
        return this.tsDoThenParseType(() => this.next());
      }
      tsDoThenParseType(cb) {
        return this.tsInType(() => {
          cb();
          return this.tsParseType();
        });
      }
      tsSkipParameterStart() {
        if (tokenIsIdentifier(this.type) || this.match(tt._this)) {
          this.next();
          return true;
        }
        if (this.match(tt.braceL)) {
          try {
            this.parseObj(true);
            return true;
          } catch {
            return false;
          }
        }
        if (this.match(tt.bracketL)) {
          this.next();
          try {
            this.parseBindingList(tt.bracketR, true, true);
            return true;
          } catch {
            return false;
          }
        }
        return false;
      }
      tsIsUnambiguouslyStartOfFunctionType() {
        this.next();
        if (this.match(tt.parenR) || this.match(tt.ellipsis)) {
          return true;
        }
        if (this.tsSkipParameterStart()) {
          if (this.match(tt.colon) || this.match(tt.comma) || this.match(tt.question) || this.match(tt.eq)) {
            return true;
          }
          if (this.match(tt.parenR)) {
            this.next();
            if (this.match(tt.arrow)) {
              return true;
            }
          }
        }
        return false;
      }
      tsIsStartOfFunctionType() {
        if (this.tsMatchLeftRelational()) {
          return true;
        }
        return this.match(tt.parenL) && this.tsLookAhead(this.tsIsUnambiguouslyStartOfFunctionType.bind(this));
      }
      tsInAllowConditionalTypesContext(cb) {
        const oldInDisallowConditionalTypesContext = this.inDisallowConditionalTypesContext;
        this.inDisallowConditionalTypesContext = false;
        try {
          return cb();
        } finally {
          this.inDisallowConditionalTypesContext = oldInDisallowConditionalTypesContext;
        }
      }
      tsParseBindingListForSignature() {
        return super.parseBindingList(tt.parenR, true, true).map((pattern) => {
          if (pattern.type !== "Identifier" && pattern.type !== "RestElement" && pattern.type !== "ObjectPattern" && pattern.type !== "ArrayPattern") {
            this.raise(
              pattern.start,
              TypeScriptError.UnsupportedSignatureParameterKind({ type: pattern.type })
            );
          }
          return pattern;
        });
      }
      tsParseTypePredicateAsserts() {
        if (this.type !== tokTypes2.asserts) {
          return false;
        }
        const containsEsc = this.containsEsc;
        this.next();
        if (!tokenIsIdentifier(this.type) && !this.match(tt._this)) {
          return false;
        }
        if (containsEsc) {
          this.raise(this.lastTokStart, "Escape sequence in keyword asserts");
        }
        return true;
      }
      tsParseThisTypeNode() {
        const node = this.startNode();
        this.next();
        return this.finishNode(node, "TSThisType");
      }
      tsParseTypeAnnotation(eatColon = true, t = this.startNode()) {
        this.tsInType(() => {
          if (eatColon) this.expect(tt.colon);
          t.typeAnnotation = this.tsParseType();
        });
        return this.finishNode(t, "TSTypeAnnotation");
      }
      tsParseThisTypePredicate(lhs) {
        this.next();
        const node = this.startNodeAtNode(lhs);
        node.parameterName = lhs;
        node.typeAnnotation = this.tsParseTypeAnnotation(
          /* eatColon */
          false
        );
        node.asserts = false;
        return this.finishNode(node, "TSTypePredicate");
      }
      tsParseThisTypeOrThisTypePredicate() {
        const thisKeyword = this.tsParseThisTypeNode();
        if (this.isContextual("is") && !this.hasPrecedingLineBreak()) {
          return this.tsParseThisTypePredicate(thisKeyword);
        } else {
          return thisKeyword;
        }
      }
      tsParseTypePredicatePrefix() {
        const id = this.parseIdent();
        if (this.isContextual("is") && !this.hasPrecedingLineBreak()) {
          this.next();
          return id;
        }
      }
      tsParseTypeOrTypePredicateAnnotation(returnToken) {
        return this.tsInType(() => {
          const t = this.startNode();
          this.expect(returnToken);
          const node = this.startNode();
          const asserts = !!this.tsTryParse(this.tsParseTypePredicateAsserts.bind(this));
          if (asserts && this.match(tt._this)) {
            let thisTypePredicate = this.tsParseThisTypeOrThisTypePredicate();
            if (thisTypePredicate.type === "TSThisType") {
              node.parameterName = thisTypePredicate;
              node.asserts = true;
              node.typeAnnotation = null;
              thisTypePredicate = this.finishNode(node, "TSTypePredicate");
            } else {
              this.resetStartLocationFromNode(thisTypePredicate, node);
              thisTypePredicate.asserts = true;
            }
            t.typeAnnotation = thisTypePredicate;
            return this.finishNode(t, "TSTypeAnnotation");
          }
          const typePredicateVariable = this.tsIsIdentifier() && this.tsTryParse(this.tsParseTypePredicatePrefix.bind(this));
          if (!typePredicateVariable) {
            if (!asserts) {
              return this.tsParseTypeAnnotation(
                /* eatColon */
                false,
                t
              );
            }
            node.parameterName = this.parseIdent();
            node.asserts = asserts;
            node.typeAnnotation = null;
            t.typeAnnotation = this.finishNode(node, "TSTypePredicate");
            return this.finishNode(t, "TSTypeAnnotation");
          }
          const type = this.tsParseTypeAnnotation(
            /* eatColon */
            false
          );
          node.parameterName = typePredicateVariable;
          node.typeAnnotation = type;
          node.asserts = asserts;
          t.typeAnnotation = this.finishNode(node, "TSTypePredicate");
          return this.finishNode(t, "TSTypeAnnotation");
        });
      }
      // Note: In TypeScript implementation we must provide `yieldContext` and `awaitContext`,
      // but here it's always false, because this is only used for types.
      tsFillSignature(returnToken, signature) {
        const returnTokenRequired = returnToken === tt.arrow;
        const paramsKey = "parameters";
        const returnTypeKey = "typeAnnotation";
        signature.typeParameters = this.tsTryParseTypeParameters();
        this.expect(tt.parenL);
        signature[paramsKey] = this.tsParseBindingListForSignature();
        if (returnTokenRequired) {
          signature[returnTypeKey] = this.tsParseTypeOrTypePredicateAnnotation(returnToken);
        } else if (this.match(returnToken)) {
          signature[returnTypeKey] = this.tsParseTypeOrTypePredicateAnnotation(returnToken);
        }
      }
      tsTryNextParseConstantContext() {
        if (this.lookahead().type !== tt._const) return null;
        this.next();
        const typeReference = this.tsParseTypeReference();
        if (typeReference.typeParameters || typeReference.typeArguments) {
          this.raise(
            typeReference.typeName.start,
            TypeScriptError.CannotFindName({
              name: "const"
            })
          );
        }
        return typeReference;
      }
      tsParseFunctionOrConstructorType(type, abstract) {
        const node = this.startNode();
        if (type === "TSConstructorType") {
          node.abstract = !!abstract;
          if (abstract) this.next();
          this.next();
        }
        this.tsInAllowConditionalTypesContext(() => this.tsFillSignature(tt.arrow, node));
        return this.finishNode(node, type);
      }
      tsParseUnionOrIntersectionType(kind, parseConstituentType, operator) {
        const node = this.startNode();
        const hasLeadingOperator = this.eat(operator);
        const types = [];
        do {
          types.push(parseConstituentType());
        } while (this.eat(operator));
        if (types.length === 1 && !hasLeadingOperator) {
          return types[0];
        }
        node.types = types;
        return this.finishNode(node, kind);
      }
      tsCheckTypeAnnotationForReadOnly(node) {
        switch (node.typeAnnotation.type) {
          case "TSTupleType":
          case "TSArrayType":
            return;
          default:
            this.raise(node.start, TypeScriptError.UnexpectedReadonly);
        }
      }
      tsParseTypeOperator() {
        const node = this.startNode();
        const operator = this.value;
        this.next();
        node.operator = operator;
        node.typeAnnotation = this.tsParseTypeOperatorOrHigher();
        if (operator === "readonly") {
          this.tsCheckTypeAnnotationForReadOnly(node);
        }
        return this.finishNode(node, "TSTypeOperator");
      }
      tsParseConstraintForInferType() {
        if (this.eat(tt._extends)) {
          const constraint = this.tsInDisallowConditionalTypesContext(() => this.tsParseType());
          if (this.inDisallowConditionalTypesContext || !this.match(tt.question)) {
            return constraint;
          }
        }
      }
      tsParseInferType() {
        const node = this.startNode();
        this.expectContextual("infer");
        const typeParameter = this.startNode();
        typeParameter.name = this.tsParseTypeParameterName();
        typeParameter.constraint = this.tsTryParse(() => this.tsParseConstraintForInferType());
        node.typeParameter = this.finishNode(typeParameter, "TSTypeParameter");
        return this.finishNode(node, "TSInferType");
      }
      tsParseLiteralTypeNode() {
        const node = this.startNode();
        node.literal = (() => {
          switch (this.type) {
            case tt.num:
            // we don't need bigint type here
            // case tt.bigint:
            case tt.string:
            case tt._true:
            case tt._false:
              return this.parseExprAtom();
            default:
              this.unexpected();
          }
        })();
        return this.finishNode(node, "TSLiteralType");
      }
      tsParseImportType() {
        const node = this.startNode();
        this.expect(tt._import);
        this.expect(tt.parenL);
        if (!this.match(tt.string)) {
          this.raise(this.start, TypeScriptError.UnsupportedImportTypeArgument);
        }
        node.argument = this.parseExprAtom();
        this.expect(tt.parenR);
        if (this.eat(tt.dot)) {
          node.qualifier = this.tsParseEntityName();
        }
        if (this.tsMatchLeftRelational()) {
          node.typeArguments = this.tsParseTypeArguments();
        }
        return this.finishNode(node, "TSImportType");
      }
      tsParseTypeQuery() {
        const node = this.startNode();
        this.expect(tt._typeof);
        if (this.match(tt._import)) {
          node.exprName = this.tsParseImportType();
        } else {
          node.exprName = this.tsParseEntityName();
        }
        if (!this.hasPrecedingLineBreak() && this.tsMatchLeftRelational()) {
          node.typeArguments = this.tsParseTypeArguments();
        }
        return this.finishNode(node, "TSTypeQuery");
      }
      tsParseMappedTypeParameter() {
        const node = this.startNode();
        node.name = this.tsParseTypeParameterName();
        node.constraint = this.tsExpectThenParseType(tt._in);
        return this.finishNode(node, "TSTypeParameter");
      }
      tsParseMappedType() {
        const node = this.startNode();
        this.expect(tt.braceL);
        if (this.match(tt.plusMin)) {
          node.readonly = this.value;
          this.next();
          this.expectContextual("readonly");
        } else if (this.eatContextual("readonly")) {
          node.readonly = true;
        }
        this.expect(tt.bracketL);
        node.typeParameter = this.tsParseMappedTypeParameter();
        node.nameType = this.eatContextual("as") ? this.tsParseType() : null;
        this.expect(tt.bracketR);
        if (this.match(tt.plusMin)) {
          node.optional = this.value;
          this.next();
          this.expect(tt.question);
        } else if (this.eat(tt.question)) {
          node.optional = true;
        }
        node.typeAnnotation = this.tsTryParseType();
        this.semicolon();
        this.expect(tt.braceR);
        return this.finishNode(node, "TSMappedType");
      }
      tsParseTypeLiteral() {
        const node = this.startNode();
        node.members = this.tsParseObjectTypeMembers();
        return this.finishNode(node, "TSTypeLiteral");
      }
      tsParseTupleElementType() {
        const startLoc = this.startLoc;
        const startPos = this["start"];
        const rest = this.eat(tt.ellipsis);
        let type = this.tsParseType();
        const optional = this.eat(tt.question);
        const labeled = this.eat(tt.colon);
        if (labeled) {
          const labeledNode = this.startNodeAtNode(type);
          labeledNode.optional = optional;
          if (type.type === "TSTypeReference" && !type.typeArguments && type.typeName.type === "Identifier") {
            labeledNode.label = type.typeName;
          } else {
            this.raise(type.start, TypeScriptError.InvalidTupleMemberLabel);
            labeledNode.label = type;
          }
          labeledNode.elementType = this.tsParseType();
          type = this.finishNode(labeledNode, "TSNamedTupleMember");
        } else if (optional) {
          const optionalTypeNode = this.startNodeAtNode(type);
          optionalTypeNode.typeAnnotation = type;
          type = this.finishNode(optionalTypeNode, "TSOptionalType");
        }
        if (rest) {
          const restNode = this.startNodeAt(startPos, startLoc);
          restNode.typeAnnotation = type;
          type = this.finishNode(restNode, "TSRestType");
        }
        return type;
      }
      tsParseTupleType() {
        const node = this.startNode();
        node.elementTypes = this.tsParseBracketedList(
          "TupleElementTypes",
          this.tsParseTupleElementType.bind(this),
          /* bracket */
          true,
          /* skipFirstToken */
          false
        );
        let seenOptionalElement = false;
        node.elementTypes.forEach((elementNode) => {
          const { type } = elementNode;
          if (seenOptionalElement && type !== "TSRestType" && type !== "TSOptionalType" && !(type === "TSNamedTupleMember" && elementNode.optional)) {
            this.raise(elementNode.start, TypeScriptError.OptionalTypeBeforeRequired);
          }
          seenOptionalElement ||= type === "TSNamedTupleMember" && elementNode.optional || type === "TSOptionalType";
          if (type === "TSRestType") {
            elementNode = elementNode.typeAnnotation;
          }
        });
        return this.finishNode(node, "TSTupleType");
      }
      tsParseTemplateLiteralType() {
        const node = this.startNode();
        node.literal = this.parseTemplate({ isTagged: false });
        return this.finishNode(node, "TSLiteralType");
      }
      tsParseTypeReference() {
        const node = this.startNode();
        node.typeName = this.tsParseEntityName();
        if (!this.hasPrecedingLineBreak() && this.tsMatchLeftRelational()) {
          node.typeArguments = this.tsParseTypeArguments();
        }
        return this.finishNode(node, "TSTypeReference");
      }
      tsMatchLeftRelational() {
        return this.match(tt.relational) && this.value === "<";
      }
      tsMatchRightRelational() {
        return this.match(tt.relational) && this.value === ">";
      }
      tsParseParenthesizedType() {
        const node = this.startNode();
        this.expect(tt.parenL);
        node.typeAnnotation = this.tsParseType();
        this.expect(tt.parenR);
        return this.finishNode(node, "TSParenthesizedType");
      }
      tsParseNonArrayType() {
        switch (this.type) {
          case tt.string:
          case tt.num:
          // we don't need bigint type here
          // case tt.bigint:
          case tt._true:
          case tt._false:
            return this.tsParseLiteralTypeNode();
          case tt.plusMin:
            if (this.value === "-") {
              const node = this.startNode();
              const nextToken = this.lookahead();
              if (nextToken.type !== tt.num) {
                this.unexpected();
              }
              node.literal = this.parseMaybeUnary();
              return this.finishNode(node, "TSLiteralType");
            }
            break;
          case tt._this:
            return this.tsParseThisTypeOrThisTypePredicate();
          case tt._typeof:
            return this.tsParseTypeQuery();
          case tt._import:
            return this.tsParseImportType();
          case tt.braceL:
            return this.tsLookAhead(this.tsIsStartOfMappedType.bind(this)) ? this.tsParseMappedType() : this.tsParseTypeLiteral();
          case tt.bracketL:
            return this.tsParseTupleType();
          case tt.parenL:
            return this.tsParseParenthesizedType();
          // parse template string here
          case tt.backQuote:
          case tt.dollarBraceL:
            return this.tsParseTemplateLiteralType();
          default: {
            const { type } = this;
            if (tokenIsIdentifier(type) || type === tt._void || type === tt._null) {
              const nodeType = type === tt._void ? "TSVoidKeyword" : type === tt._null ? "TSNullKeyword" : keywordTypeFromName(this.value);
              if (nodeType !== void 0 && this.lookaheadCharCode() !== 46) {
                const node = this.startNode();
                this.next();
                return this.finishNode(node, nodeType);
              }
              return this.tsParseTypeReference();
            }
          }
        }
        this.unexpected();
      }
      tsParseArrayTypeOrHigher() {
        let type = this.tsParseNonArrayType();
        while (!this.hasPrecedingLineBreak() && this.eat(tt.bracketL)) {
          if (this.match(tt.bracketR)) {
            const node = this.startNodeAtNode(type);
            node.elementType = type;
            this.expect(tt.bracketR);
            type = this.finishNode(node, "TSArrayType");
          } else {
            const node = this.startNodeAtNode(type);
            node.objectType = type;
            node.indexType = this.tsParseType();
            this.expect(tt.bracketR);
            type = this.finishNode(node, "TSIndexedAccessType");
          }
        }
        return type;
      }
      tsParseTypeOperatorOrHigher() {
        const isTypeOperator = tokenIsTSTypeOperator(this.type) && !this.containsEsc;
        return isTypeOperator ? this.tsParseTypeOperator() : this.isContextual("infer") ? this.tsParseInferType() : this.tsInAllowConditionalTypesContext(() => this.tsParseArrayTypeOrHigher());
      }
      tsParseIntersectionTypeOrHigher() {
        return this.tsParseUnionOrIntersectionType(
          "TSIntersectionType",
          this.tsParseTypeOperatorOrHigher.bind(this),
          tt.bitwiseAND
        );
      }
      tsParseUnionTypeOrHigher() {
        return this.tsParseUnionOrIntersectionType(
          "TSUnionType",
          this.tsParseIntersectionTypeOrHigher.bind(this),
          tt.bitwiseOR
        );
      }
      tsParseNonConditionalType() {
        if (this.tsIsStartOfFunctionType()) {
          return this.tsParseFunctionOrConstructorType("TSFunctionType");
        }
        if (this.match(tt._new)) {
          return this.tsParseFunctionOrConstructorType("TSConstructorType");
        } else if (this.isAbstractConstructorSignature()) {
          return this.tsParseFunctionOrConstructorType(
            "TSConstructorType",
            /* abstract */
            true
          );
        }
        return this.tsParseUnionTypeOrHigher();
      }
      /** Be sure to be in a type context before calling this, using `tsInType`. */
      tsParseType() {
        assert(this.inType);
        const type = this.tsParseNonConditionalType();
        if (this.inDisallowConditionalTypesContext || this.hasPrecedingLineBreak() || !this.eat(tt._extends)) {
          return type;
        }
        const node = this.startNodeAtNode(type);
        node.checkType = type;
        node.extendsType = this.tsInDisallowConditionalTypesContext(
          () => this.tsParseNonConditionalType()
        );
        this.expect(tt.question);
        node.trueType = this.tsInAllowConditionalTypesContext(() => this.tsParseType());
        this.expect(tt.colon);
        node.falseType = this.tsInAllowConditionalTypesContext(() => this.tsParseType());
        return this.finishNode(node, "TSConditionalType");
      }
      tsIsUnambiguouslyIndexSignature() {
        this.next();
        if (tokenIsIdentifier(this.type)) {
          this.next();
          return this.match(tt.colon);
        }
        return false;
      }
      /**
       * Runs `cb` in a type context.
       * This should be called one token *before* the first type token,
       * so that the call to `next()` is run in type context.
       */
      tsInType(cb) {
        const oldInType = this.inType;
        this.inType = true;
        try {
          return cb();
        } finally {
          this.inType = oldInType;
        }
      }
      tsTryParseIndexSignature(node) {
        if (!(this.match(tt.bracketL) && this.tsLookAhead(this.tsIsUnambiguouslyIndexSignature.bind(this)))) {
          return void 0;
        }
        this.expect(tt.bracketL);
        const id = this.parseIdent();
        id.typeAnnotation = this.tsParseTypeAnnotation();
        this.resetEndLocation(id);
        this.expect(tt.bracketR);
        node.parameters = [id];
        const type = this.tsTryParseTypeAnnotation();
        if (type) node.typeAnnotation = type;
        this.tsParseTypeMemberSemicolon();
        return this.finishNode(node, "TSIndexSignature");
      }
      // for better error recover
      tsParseNoneModifiers(node) {
        this.tsParseModifiers({
          modified: node,
          allowedModifiers: [],
          disallowedModifiers: ["in", "out"],
          errorTemplate: TypeScriptError.InvalidModifierOnTypeParameterPositions
        });
      }
      tsParseTypeParameter(parseModifiers = this.tsParseNoneModifiers.bind(this)) {
        const node = this.startNode();
        parseModifiers(node);
        node.name = this.tsParseTypeParameterName();
        node.constraint = this.tsEatThenParseType(tt._extends);
        node.default = this.tsEatThenParseType(tt.eq);
        return this.finishNode(node, "TSTypeParameter");
      }
      tsParseTypeParameters(parseModifiers) {
        const node = this.startNode();
        if (this.tsMatchLeftRelational() || this.matchJsx("jsxTagStart")) {
          this.next();
        } else {
          this.unexpected();
        }
        const refTrailingCommaPos = { value: -1 };
        node.params = this.tsParseBracketedList(
          "TypeParametersOrArguments",
          this.tsParseTypeParameter.bind(this, parseModifiers),
          /* bracket */
          false,
          /* skipFirstToken */
          true,
          refTrailingCommaPos
        );
        if (node.params.length === 0) {
          this.raise(this.start, TypeScriptError.EmptyTypeParameters);
        }
        if (refTrailingCommaPos.value !== -1) {
          this.addExtra(node, "trailingComma", refTrailingCommaPos.value);
        }
        return this.finishNode(node, "TSTypeParameterDeclaration");
      }
      tsTryParseTypeParameters(parseModifiers) {
        if (this.tsMatchLeftRelational()) {
          return this.tsParseTypeParameters(parseModifiers);
        }
      }
      tsTryParse(f) {
        const state = this.getCurLookaheadState();
        const result = f();
        if (result !== void 0 && result !== false) {
          return result;
        } else {
          this.setLookaheadState(state);
          return void 0;
        }
      }
      tsTokenCanFollowModifier() {
        return (this.match(tt.bracketL) || this.match(tt.braceL) || this.match(tt.star) || this.match(tt.ellipsis) || this.match(tt.privateId) || this.isLiteralPropertyName()) && !this.hasPrecedingLineBreak();
      }
      tsNextTokenCanFollowModifier() {
        this.next(true);
        return this.tsTokenCanFollowModifier();
      }
      /** Parses a modifier matching one the given modifier names. */
      tsParseModifier(allowedModifiers, stopOnStartOfClassStaticBlock) {
        const modifier = this.value;
        if (allowedModifiers.indexOf(modifier) !== -1 && !this.containsEsc) {
          if (stopOnStartOfClassStaticBlock && this.tsIsStartOfStaticBlocks()) {
            return void 0;
          }
          if (this.tsTryParse(this.tsNextTokenCanFollowModifier.bind(this))) {
            return modifier;
          }
        }
        return void 0;
      }
      tsParseModifiersByMap({
        modified,
        map
      }) {
        for (const key of Object.keys(map)) {
          modified[key] = map[key];
        }
      }
      /** Parses a list of modifiers, in any order.
       *  If you need a specific order, you must call this function multiple times:
       *    this.tsParseModifiers({ modified: node, allowedModifiers: ['public'] });
       *    this.tsParseModifiers({ modified: node, allowedModifiers: ["abstract", "readonly"] });
       */
      tsParseModifiers({
        modified,
        allowedModifiers,
        disallowedModifiers,
        stopOnStartOfClassStaticBlock,
        errorTemplate = TypeScriptError.InvalidModifierOnTypeMember
      }) {
        const modifiedMap = {};
        const enforceOrder = (loc, modifier, before, after) => {
          if (modifier === before && modified[after]) {
            this.raise(
              loc.column,
              TypeScriptError.InvalidModifiersOrder({ orderedModifiers: [before, after] })
            );
          }
        };
        const incompatible = (loc, modifier, mod1, mod2) => {
          if (modified[mod1] && modifier === mod2 || modified[mod2] && modifier === mod1) {
            this.raise(
              loc.column,
              TypeScriptError.IncompatibleModifiers({ modifiers: [mod1, mod2] })
            );
          }
        };
        for (; ; ) {
          const startLoc = this.startLoc;
          const modifier = this.tsParseModifier(
            allowedModifiers.concat(disallowedModifiers ?? []),
            stopOnStartOfClassStaticBlock
          );
          if (!modifier) break;
          if (tsIsAccessModifier(modifier)) {
            if (modified.accessibility) {
              this.raise(this.start, TypeScriptError.DuplicateAccessibilityModifier());
            } else {
              enforceOrder(startLoc, modifier, modifier, "override");
              enforceOrder(startLoc, modifier, modifier, "static");
              enforceOrder(startLoc, modifier, modifier, "readonly");
              enforceOrder(startLoc, modifier, modifier, "accessor");
              modifiedMap.accessibility = modifier;
              modified["accessibility"] = modifier;
            }
          } else if (tsIsVarianceAnnotations(modifier)) {
            if (modified[modifier]) {
              this.raise(this.start, TypeScriptError.DuplicateModifier({ modifier }));
            } else {
              enforceOrder(startLoc, modifier, "in", "out");
              modifiedMap[modifier] = modifier;
              modified[modifier] = true;
            }
          } else if (tsIsClassAccessor(modifier)) {
            if (modified[modifier]) {
              this.raise(this.start, TypeScriptError.DuplicateModifier({ modifier }));
            } else {
              incompatible(startLoc, modifier, "accessor", "readonly");
              incompatible(startLoc, modifier, "accessor", "static");
              incompatible(startLoc, modifier, "accessor", "override");
              modifiedMap[modifier] = modifier;
              modified[modifier] = true;
            }
          } else if (modifier === "const") {
            if (modified[modifier]) {
              this.raise(this.start, TypeScriptError.DuplicateModifier({ modifier }));
            } else {
              modifiedMap[modifier] = modifier;
              modified[modifier] = true;
            }
          } else {
            if (Object.hasOwnProperty.call(modified, modifier)) {
              this.raise(this.start, TypeScriptError.DuplicateModifier({ modifier }));
            } else {
              enforceOrder(startLoc, modifier, "static", "readonly");
              enforceOrder(startLoc, modifier, "static", "override");
              enforceOrder(startLoc, modifier, "override", "readonly");
              enforceOrder(startLoc, modifier, "abstract", "override");
              incompatible(startLoc, modifier, "declare", "override");
              incompatible(startLoc, modifier, "static", "abstract");
              modifiedMap[modifier] = modifier;
              modified[modifier] = true;
            }
          }
          if (disallowedModifiers?.includes(modifier)) {
            this.raise(this.start, errorTemplate);
          }
        }
        return modifiedMap;
      }
      tsParseInOutModifiers(node) {
        this.tsParseModifiers({
          modified: node,
          allowedModifiers: ["in", "out"],
          disallowedModifiers: [
            "public",
            "private",
            "protected",
            "readonly",
            "declare",
            "abstract",
            "override"
          ],
          errorTemplate: TypeScriptError.InvalidModifierOnTypeParameter
        });
      }
      // Handle type assertions
      parseMaybeUnary(refExpressionErrors, sawUnary, incDec, forInit) {
        if (!options?.jsx && this.tsMatchLeftRelational()) {
          return this.tsParseTypeAssertion();
        } else {
          return super.parseMaybeUnary(refExpressionErrors, sawUnary, incDec, forInit);
        }
      }
      tsParseTypeAssertion() {
        if (disallowAmbiguousJSXLike) {
          this.raise(this.start, TypeScriptError.ReservedTypeAssertion);
        }
        const result = this.tryParse(() => {
          const node = this.startNode();
          const _const = this.tsTryNextParseConstantContext();
          node.typeAnnotation = _const || this.tsNextThenParseType();
          this.expect(tt.relational);
          node.expression = this.parseMaybeUnary();
          return this.finishNode(node, "TSTypeAssertion");
        });
        if (result.error) {
          return this.tsParseTypeParameters(this.tsParseConstModifier);
        } else {
          return result.node;
        }
      }
      tsParseTypeArguments() {
        const node = this.startNode();
        node.params = this.tsInType(
          () => (
            // Temporarily remove a JSX parsing context, which makes us scan different tokens.
            this.tsInNoContext(() => {
              this.expect(tt.relational);
              return this.tsParseDelimitedList(
                "TypeParametersOrArguments",
                this.tsParseType.bind(this)
              );
            })
          )
        );
        if (node.params.length === 0) {
          this.raise(this.start, TypeScriptError.EmptyTypeArguments);
        }
        this.exprAllowed = false;
        this.expect(tt.relational);
        return this.finishNode(node, "TSTypeParameterInstantiation");
      }
      tsParseHeritageClause(token) {
        const originalStart = this.start;
        const delimitedList = this.tsParseDelimitedList("HeritageClauseElement", () => {
          const node = this.startNode();
          node.expression = this.tsParseEntityName();
          if (this.tsMatchLeftRelational()) {
            node.typeParameters = this.tsParseTypeArguments();
          }
          return this.finishNode(node, "TSExpressionWithTypeArguments");
        });
        if (!delimitedList.length) {
          this.raise(originalStart, TypeScriptError.EmptyHeritageClauseType({ token }));
        }
        return delimitedList;
      }
      tsParseTypeMemberSemicolon() {
        if (!this.eat(tt.comma) && !this.isLineTerminator()) {
          this.expect(tt.semi);
        }
      }
      tsTryParseAndCatch(f) {
        const result = this.tryParse(
          (abort) => (
            // @ts-expect-error todo(flow->ts)
            f() || abort()
          )
        );
        if (result.aborted || !result.node) return void 0;
        if (result.error) this.setLookaheadState(result.failState);
        return result.node;
      }
      tsParseSignatureMember(kind, node) {
        this.tsFillSignature(tt.colon, node);
        this.tsParseTypeMemberSemicolon();
        return this.finishNode(node, kind);
      }
      tsParsePropertyOrMethodSignature(node, readonly) {
        if (this.eat(tt.question)) node.optional = true;
        const nodeAny = node;
        if (this.match(tt.parenL) || this.tsMatchLeftRelational()) {
          if (readonly) {
            this.raise(node.start, TypeScriptError.ReadonlyForMethodSignature);
          }
          const method = nodeAny;
          if (method.kind && this.tsMatchLeftRelational()) {
            this.raise(this.start, TypeScriptError.AccesorCannotHaveTypeParameters);
          }
          this.tsFillSignature(tt.colon, method);
          this.tsParseTypeMemberSemicolon();
          const paramsKey = "parameters";
          const returnTypeKey = "typeAnnotation";
          if (method.kind === "get") {
            if (method[paramsKey].length > 0) {
              this.raise(this.start, "A 'get' accesor must not have any formal parameters.");
              if (this.isThisParam(method[paramsKey][0])) {
                this.raise(this.start, TypeScriptError.AccesorCannotDeclareThisParameter);
              }
            }
          } else if (method.kind === "set") {
            if (method[paramsKey].length !== 1) {
              this.raise(this.start, "A 'get' accesor must not have any formal parameters.");
            } else {
              const firstParameter = method[paramsKey][0];
              if (this.isThisParam(firstParameter)) {
                this.raise(this.start, TypeScriptError.AccesorCannotDeclareThisParameter);
              }
              if (firstParameter.type === "Identifier" && firstParameter.optional) {
                this.raise(this.start, TypeScriptError.SetAccesorCannotHaveOptionalParameter);
              }
              if (firstParameter.type === "RestElement") {
                this.raise(this.start, TypeScriptError.SetAccesorCannotHaveRestParameter);
              }
            }
            if (method[returnTypeKey]) {
              this.raise(
                method[returnTypeKey].start,
                TypeScriptError.SetAccesorCannotHaveReturnType
              );
            }
          } else {
            method.kind = "method";
          }
          return this.finishNode(method, "TSMethodSignature");
        } else {
          const property = nodeAny;
          if (readonly) property.readonly = true;
          const type = this.tsTryParseTypeAnnotation();
          if (type) property.typeAnnotation = type;
          this.tsParseTypeMemberSemicolon();
          return this.finishNode(property, "TSPropertySignature");
        }
      }
      tsParseTypeMember() {
        const node = this.startNode();
        if (this.match(tt.parenL) || this.tsMatchLeftRelational()) {
          return this.tsParseSignatureMember("TSCallSignatureDeclaration", node);
        }
        if (this.match(tt._new)) {
          const id = this.startNode();
          this.next();
          if (this.match(tt.parenL) || this.tsMatchLeftRelational()) {
            return this.tsParseSignatureMember("TSConstructSignatureDeclaration", node);
          } else {
            node.key = this.createIdentifier(id, "new");
            return this.tsParsePropertyOrMethodSignature(node, false);
          }
        }
        this.tsParseModifiers({
          modified: node,
          allowedModifiers: ["readonly"],
          disallowedModifiers: [
            "declare",
            "abstract",
            "private",
            "protected",
            "public",
            "static",
            "override"
          ]
        });
        const idx = this.tsTryParseIndexSignature(node);
        if (idx) {
          return idx;
        }
        this.parsePropertyName(node);
        if (!node.computed && node.key.type === "Identifier" && (node.key.name === "get" || node.key.name === "set") && this.tsTokenCanFollowModifier()) {
          node.kind = node.key.name;
          this.parsePropertyName(node);
        }
        return this.tsParsePropertyOrMethodSignature(node, !!node.readonly);
      }
      tsParseList(kind, parseElement) {
        const result = [];
        while (!this.tsIsListTerminator(kind)) {
          result.push(parseElement());
        }
        return result;
      }
      tsParseObjectTypeMembers() {
        this.expect(tt.braceL);
        const members = this.tsParseList("TypeMembers", this.tsParseTypeMember.bind(this));
        this.expect(tt.braceR);
        return members;
      }
      tsParseInterfaceDeclaration(node, properties = {}) {
        if (this.hasFollowingLineBreak()) return null;
        this.expectContextual("interface");
        if (properties.declare) node.declare = true;
        if (tokenIsIdentifier(this.type)) {
          node.id = this.parseIdent();
          this.checkLValSimple(node.id, acornScope.BIND_TS_INTERFACE);
        } else {
          node.id = null;
          this.raise(this.start, TypeScriptError.MissingInterfaceName);
        }
        node.typeParameters = this.tsTryParseTypeParameters(this.tsParseInOutModifiers.bind(this));
        if (this.eat(tt._extends)) {
          node.extends = this.tsParseHeritageClause("extends");
        }
        const body = this.startNode();
        body.body = this.tsParseInterfaceBody();
        node.body = this.finishNode(body, "TSInterfaceBody");
        return this.finishNode(node, "TSInterfaceDeclaration");
      }
      /**
       * Parse interface body, ensuring the closing brace is read outside of type context
       * so that decorators following the interface are properly tokenized.
       */
      tsParseInterfaceBody() {
        this.expect(tt.braceL);
        const oldInType = this.inType;
        this.inType = true;
        let members = this.tsParseList("TypeMembers", this.tsParseTypeMember.bind(this));
        this.inType = oldInType;
        this.expect(tt.braceR);
        return members;
      }
      tsParseAbstractDeclaration(node) {
        if (this.match(tt._class)) {
          node.abstract = true;
          return this.parseClass(node, true);
        } else if (this.ts_isContextual(tokTypes2.interface)) {
          if (!this.hasFollowingLineBreak()) {
            node.abstract = true;
            return this.tsParseInterfaceDeclaration(node);
          }
        } else {
          this.unexpected(node.start);
        }
      }
      tsIsDeclarationStart() {
        return tokenIsTSDeclarationStart(this.type);
      }
      tsParseExpressionStatement(node, expr) {
        switch (expr.name) {
          case "declare": {
            const declaration = this.tsTryParseDeclare(node);
            if (declaration) {
              declaration.declare = true;
              return declaration;
            }
            break;
          }
          case "global":
            if (this.match(tt.braceL)) {
              this.enterScope(TS_SCOPE_TS_MODULE);
              const mod = node;
              mod.global = true;
              mod.id = expr;
              mod.body = this.tsParseModuleBlock();
              super.exitScope();
              return this.finishNode(mod, "TSModuleDeclaration");
            }
            break;
          default:
            return this.tsParseDeclaration(
              node,
              expr.name,
              /* next */
              false
            );
        }
      }
      tsParseModuleReference() {
        return this.tsIsExternalModuleReference() ? this.tsParseExternalModuleReference() : this.tsParseEntityName(
          /* allowReservedWords */
          false
        );
      }
      tsIsExportDefaultSpecifier() {
        const { type } = this;
        const isAsync = this.isAsyncFunction();
        const isLet = this.isLet();
        if (tokenIsIdentifier(type)) {
          if (isAsync && !this.containsEsc || isLet) {
            return false;
          }
          if ((type === tokTypes2.type || type === tokTypes2.interface) && !this.containsEsc) {
            const ahead = this.lookahead();
            if (tokenIsIdentifier(ahead.type) && !this.isContextualWithState("from", ahead) || ahead.type === tt.braceL) {
              return false;
            }
          }
        } else if (!this.match(tt._default)) {
          return false;
        }
        const next = this.nextTokenStart();
        const hasFrom = this.isUnparsedContextual(next, "from");
        if (this.input.charCodeAt(next) === 44 || tokenIsIdentifier(this.type) && hasFrom) {
          return true;
        }
        if (this.match(tt._default) && hasFrom) {
          const nextAfterFrom = this.input.charCodeAt(this.nextTokenStartSince(next + 4));
          return nextAfterFrom === 34 || nextAfterFrom === 39;
        }
        return false;
      }
      tsInAmbientContext(cb) {
        const oldIsAmbientContext = this.isAmbientContext;
        this.isAmbientContext = true;
        try {
          return cb();
        } finally {
          this.isAmbientContext = oldIsAmbientContext;
        }
      }
      tsCheckLineTerminator(next) {
        if (next) {
          if (this.hasFollowingLineBreak()) return false;
          this.next();
          return true;
        }
        return !this.isLineTerminator();
      }
      tsParseModuleOrNamespaceDeclaration(node, nested = false) {
        node.id = this.parseIdent();
        if (!nested) {
          this.checkLValSimple(node.id, acornScope.BIND_TS_NAMESPACE);
        }
        if (this.eat(tt.dot)) {
          const inner = this.startNode();
          this.tsParseModuleOrNamespaceDeclaration(inner, true);
          node.body = inner;
        } else {
          this.enterScope(TS_SCOPE_TS_MODULE);
          node.body = this.tsParseModuleBlock();
          super.exitScope();
        }
        return this.finishNode(node, "TSModuleDeclaration");
      }
      checkLValSimple(expr, bindingType = acornScope.BIND_NONE, checkClashes) {
        if (expr.type === "TSNonNullExpression" || expr.type === "TSAsExpression") {
          expr = expr.expression;
        }
        return super.checkLValSimple(expr, bindingType, checkClashes);
      }
      tsParseTypeAliasDeclaration(node) {
        node.id = this.parseIdent();
        this.checkLValSimple(node.id, acornScope.BIND_TS_TYPE);
        node.typeAnnotation = this.tsInType(() => {
          node.typeParameters = this.tsTryParseTypeParameters(
            this.tsParseInOutModifiers.bind(this)
          );
          this.expect(tt.eq);
          if (this.ts_isContextual(tokTypes2.interface) && this.lookahead().type !== tt.dot) {
            const node2 = this.startNode();
            this.next();
            return this.finishNode(node2, "TSIntrinsicKeyword");
          }
          return this.tsParseType();
        });
        this.semicolon();
        return this.finishNode(node, "TSTypeAliasDeclaration");
      }
      // Common to tsTryParseDeclare, tsTryParseExportDeclaration, and tsParseExpressionStatement.
      tsParseDeclaration(node, value, next) {
        switch (value) {
          case "abstract":
            if (this.tsCheckLineTerminator(next) && (this.match(tt._class) || tokenIsIdentifier(this.type))) {
              return this.tsParseAbstractDeclaration(node);
            }
            break;
          case "module":
            if (this.tsCheckLineTerminator(next)) {
              if (this.match(tt.string)) {
                return this.tsParseAmbientExternalModuleDeclaration(node);
              } else if (tokenIsIdentifier(this.type)) {
                return this.tsParseModuleOrNamespaceDeclaration(node);
              }
            }
            break;
          case "namespace":
            if (this.tsCheckLineTerminator(next) && tokenIsIdentifier(this.type)) {
              return this.tsParseModuleOrNamespaceDeclaration(node);
            }
            break;
          case "type":
            if (this.tsCheckLineTerminator(next) && tokenIsIdentifier(this.type)) {
              return this.tsParseTypeAliasDeclaration(node);
            }
            break;
        }
      }
      // Note: this won't b·e called unless the keyword is allowed in
      // `shouldParseExportDeclaration`.
      tsTryParseExportDeclaration() {
        return this.tsParseDeclaration(
          this.startNode(),
          this.value,
          /* next */
          true
        );
      }
      tsParseImportEqualsDeclaration(node, isExport) {
        node.isExport = isExport || false;
        node.id = this.parseIdent();
        this.checkLValSimple(node.id, acornScope.BIND_LEXICAL);
        super.expect(tt.eq);
        const moduleReference = this.tsParseModuleReference();
        if (node.importKind === "type" && moduleReference.type !== "TSExternalModuleReference") {
          this.raise(moduleReference.start, TypeScriptError.ImportAliasHasImportType);
        }
        node.moduleReference = moduleReference;
        super.semicolon();
        return this.finishNode(node, "TSImportEqualsDeclaration");
      }
      isExportDefaultSpecifier() {
        if (this.tsIsDeclarationStart()) return false;
        const { type } = this;
        if (tokenIsIdentifier(type)) {
          if (this.isContextual("async") || this.isContextual("let")) {
            return false;
          }
          if ((type === tokTypes2.type || type === tokTypes2.interface) && !this.containsEsc) {
            const ahead = this.lookahead();
            if (tokenIsIdentifier(ahead.type) && !this.isContextualWithState("from", ahead) || ahead.type === tt.braceL) {
              return false;
            }
          }
        } else if (!this.match(tt._default)) {
          return false;
        }
        const next = this.nextTokenStart();
        const hasFrom = this.isUnparsedContextual(next, "from");
        if (this.input.charCodeAt(next) === 44 || tokenIsIdentifier(this.type) && hasFrom) {
          return true;
        }
        if (this.match(tt._default) && hasFrom) {
          const nextAfterFrom = this.input.charCodeAt(this.nextTokenStartSince(next + 4));
          return nextAfterFrom === 34 || nextAfterFrom === 39;
        }
        return false;
      }
      parseTemplate({ isTagged = false } = {}) {
        let node = this.startNode();
        this.next();
        node.expressions = [];
        let curElt = this.parseTemplateElement({ isTagged });
        node.quasis = [curElt];
        while (!curElt.tail) {
          if (this.type === tt.eof) this.raise(this.pos, "Unterminated template literal");
          this.expect(tt.dollarBraceL);
          node.expressions.push(this.inType ? this.tsParseType() : this.parseExpression());
          this.expect(tt.braceR);
          node.quasis.push(curElt = this.parseTemplateElement({ isTagged }));
        }
        this.next();
        return this.finishNode(node, "TemplateLiteral");
      }
      parseFunction(node, statement, allowExpressionBody, isAsync, forInit) {
        this.initFunction(node);
        if (this.ecmaVersion >= 9 || this.ecmaVersion >= 6 && !isAsync) {
          if (this.type === tt.star && statement & FUNC_HANGING_STATEMENT) {
            this.unexpected();
          }
          node.generator = this.eat(tt.star);
        }
        if (this.ecmaVersion >= 8) {
          node.async = !!isAsync;
        }
        if (statement & FUNC_STATEMENT) {
          node.id = statement & FUNC_NULLABLE_ID && this.type !== tt.name ? null : this.parseIdent();
        }
        let oldYieldPos = this.yieldPos, oldAwaitPos = this.awaitPos, oldAwaitIdentPos = this.awaitIdentPos;
        const oldMaybeInArrowParameters = this.maybeInArrowParameters;
        this.maybeInArrowParameters = false;
        this.yieldPos = 0;
        this.awaitPos = 0;
        this.awaitIdentPos = 0;
        this.enterScope(functionFlags(node.async, node.generator));
        if (!(statement & FUNC_STATEMENT)) {
          node.id = this.type === tt.name ? this.parseIdent() : null;
        }
        this.parseFunctionParams(node);
        const isDeclaration = statement & FUNC_STATEMENT;
        this.parseFunctionBody(node, allowExpressionBody, false, forInit, {
          isFunctionDeclaration: isDeclaration
        });
        this.yieldPos = oldYieldPos;
        this.awaitPos = oldAwaitPos;
        this.awaitIdentPos = oldAwaitIdentPos;
        if (statement & FUNC_STATEMENT && node.id && !(statement & FUNC_HANGING_STATEMENT)) {
          if (node.body) {
            this.checkLValSimple(
              node.id,
              this.strict || node.generator || node.async ? this.treatFunctionsAsVar ? acornScope.BIND_VAR : acornScope.BIND_LEXICAL : acornScope.BIND_FUNCTION
            );
          } else {
            this.checkLValSimple(node.id, acornScope.BIND_NONE);
          }
        }
        this.maybeInArrowParameters = oldMaybeInArrowParameters;
        return this.finishNode(node, isDeclaration ? "FunctionDeclaration" : "FunctionExpression");
      }
      parseFunctionBody(node, isArrowFunction = false, isMethod = false, forInit = false, tsConfig) {
        if (this.match(tt.colon)) {
          node.returnType = this.tsParseTypeOrTypePredicateAnnotation(tt.colon);
        }
        const bodilessType = tsConfig?.isFunctionDeclaration ? "TSDeclareFunction" : tsConfig?.isClassMethod ? "TSDeclareMethod" : void 0;
        if (bodilessType && !this.match(tt.braceL) && this.isLineTerminator()) {
          return this.finishNode(node, bodilessType);
        }
        if (bodilessType === "TSDeclareFunction" && this.isAmbientContext) {
          this.raise(node.start, TypeScriptError.DeclareFunctionHasImplementation);
          if (node.declare) {
            super.parseFunctionBody(node, isArrowFunction, isMethod, false);
            return this.finishNode(node, bodilessType);
          }
        }
        super.parseFunctionBody(node, isArrowFunction, isMethod, forInit);
        return node;
      }
      parseNew() {
        if (this.containsEsc) this.raiseRecoverable(this.start, "Escape sequence in keyword new");
        let node = this.startNode();
        let meta = this.parseIdent(true);
        if (this.ecmaVersion >= 6 && this.eat(tt.dot)) {
          node.meta = meta;
          let containsEsc = this.containsEsc;
          node.property = this.parseIdent(true);
          if (node.property.name !== "target")
            this.raiseRecoverable(
              node.property.start,
              "The only valid meta property for new is 'new.target'"
            );
          if (containsEsc)
            this.raiseRecoverable(node.start, "'new.target' must not contain escaped characters");
          if (!this["allowNewDotTarget"])
            this.raiseRecoverable(
              node.start,
              "'new.target' can only be used in functions and class static block"
            );
          return this.finishNode(node, "MetaProperty");
        }
        let startPos = this.start, startLoc = this.startLoc, isImport = this.type === tt._import;
        node.callee = this.parseSubscripts(this.parseExprAtom(), startPos, startLoc, true, false);
        if (isImport && node.callee.type === "ImportExpression") {
          this.raise(startPos, "Cannot use new with import()");
        }
        const { callee } = node;
        if (callee.type === "TSInstantiationExpression" && !callee.extra?.parenthesized) {
          node.typeArguments = callee.typeArguments;
          node.callee = callee.expression;
        }
        if (this.eat(tt.parenL))
          node.arguments = this.parseExprList(tt.parenR, this.ecmaVersion >= 8, false);
        else node.arguments = [];
        return this.finishNode(node, "NewExpression");
      }
      parseExprOp(left, leftStartPos, leftStartLoc, minPrec, forInit) {
        if (tt._in.binop > minPrec && !this.hasPrecedingLineBreak()) {
          let nodeType;
          if (this.isContextual("as")) {
            nodeType = "TSAsExpression";
          }
          if (this.isContextual("satisfies")) {
            nodeType = "TSSatisfiesExpression";
          }
          if (nodeType) {
            const node = this.startNodeAt(leftStartPos, leftStartLoc);
            node.expression = left;
            const _const = this.tsTryNextParseConstantContext();
            if (_const) {
              node.typeAnnotation = _const;
            } else {
              node.typeAnnotation = this.tsNextThenParseType();
            }
            this.finishNode(node, nodeType);
            this.reScan_lt_gt();
            return this.parseExprOp(node, leftStartPos, leftStartLoc, minPrec, forInit);
          }
        }
        return super.parseExprOp(left, leftStartPos, leftStartLoc, minPrec, forInit);
      }
      parseImportSpecifiers() {
        let nodes = [], first = true;
        if (acornTypeScript.tokenIsIdentifier(this.type)) {
          nodes.push(this.parseImportDefaultSpecifier());
          if (!this.eat(tt.comma)) return nodes;
        }
        if (this.type === tt.star) {
          nodes.push(this.parseImportNamespaceSpecifier());
          return nodes;
        }
        this.expect(tt.braceL);
        while (!this.eat(tt.braceR)) {
          if (!first) {
            this.expect(tt.comma);
            if (this.afterTrailingComma(tt.braceR)) break;
          } else first = false;
          nodes.push(this.parseImportSpecifier());
        }
        return nodes;
      }
      /**
       * @param {Node} node this may be ImportDeclaration |
       * TsImportEqualsDeclaration
       * @returns AnyImport
       * */
      parseImport(node) {
        let enterHead = this.lookahead();
        node.importKind = "value";
        this.importOrExportOuterKind = "value";
        if (tokenIsIdentifier(enterHead.type) || this.match(tt.star) || this.match(tt.braceL)) {
          let ahead = this.lookahead(2);
          if (
            // import type, { a } from "b";
            ahead.type !== tt.comma && // import type from "a";
            !this.isContextualWithState("from", ahead) && // import type = require("a");
            ahead.type !== tt.eq && this.ts_eatContextualWithState("type", 1, enterHead)
          ) {
            this.importOrExportOuterKind = "type";
            node.importKind = "type";
            enterHead = this.lookahead();
            ahead = this.lookahead(2);
          }
          if (tokenIsIdentifier(enterHead.type) && ahead.type === tt.eq) {
            this.next();
            const importNode = this.tsParseImportEqualsDeclaration(node);
            this.importOrExportOuterKind = "value";
            return importNode;
          }
        }
        this.next();
        if (this.type === tt.string) {
          node.specifiers = [];
          node.source = this.parseExprAtom();
        } else {
          node.specifiers = this.parseImportSpecifiers();
          this.expectContextual("from");
          node.source = this.type === tt.string ? this.parseExprAtom() : this.unexpected();
        }
        this.parseMaybeImportAttributes(node);
        this.semicolon();
        this.finishNode(node, "ImportDeclaration");
        this.importOrExportOuterKind = "value";
        if (node.importKind === "type" && node.specifiers.length > 1 && node.specifiers[0].type === "ImportDefaultSpecifier") {
          this.raise(node.start, TypeScriptError.TypeImportCannotSpecifyDefaultAndNamed);
        }
        return node;
      }
      parseExportDefaultDeclaration() {
        if (this.isAbstractClass()) {
          const cls = this.startNode();
          this.next();
          cls.abstract = true;
          return this.parseClass(cls, true);
        }
        if (this.match(tokTypes2.interface)) {
          const result = this.tsParseInterfaceDeclaration(this.startNode());
          if (result) return result;
        }
        return super.parseExportDefaultDeclaration();
      }
      parseExportAllDeclaration(node, exports) {
        if (this.ecmaVersion >= 11) {
          if (this.eatContextual("as")) {
            node.exported = this.parseModuleExportName();
            this.checkExport(exports, node.exported, this.lastTokStart);
          } else {
            node.exported = null;
          }
        }
        this.expectContextual("from");
        if (this.type !== tt.string) this.unexpected();
        node.source = this.parseExprAtom();
        this.parseMaybeImportAttributes(node);
        this.semicolon();
        return this.finishNode(node, "ExportAllDeclaration");
      }
      parseDynamicImport(node) {
        this.next();
        node.source = this.parseMaybeAssign();
        if (this.eat(tt.comma)) {
          const expr = this.parseExpression();
          node.arguments = [expr];
        }
        if (!this.eat(tt.parenR)) {
          const errorPos = this.start;
          if (this.eat(tt.comma) && this.eat(tt.parenR)) {
            this.raiseRecoverable(errorPos, "Trailing comma is not allowed in import()");
          } else {
            this.unexpected(errorPos);
          }
        }
        return this.finishNode(node, "ImportExpression");
      }
      parseExport(node, exports) {
        let enterHead = this.lookahead();
        if (this.ts_eatWithState(tt._import, 2, enterHead)) {
          if (this.ts_isContextual(tokTypes2.type) && this.lookaheadCharCode() !== 61) {
            node.importKind = "type";
            this.importOrExportOuterKind = "type";
            this.next();
          } else {
            node.importKind = "value";
            this.importOrExportOuterKind = "value";
          }
          const exportEqualsNode = this.tsParseImportEqualsDeclaration(
            node,
            /* isExport */
            true
          );
          this.importOrExportOuterKind = void 0;
          return exportEqualsNode;
        } else if (this.ts_eatWithState(tt.eq, 2, enterHead)) {
          const assign = node;
          assign.expression = this.parseExpression();
          this.semicolon();
          this.importOrExportOuterKind = void 0;
          return this.finishNode(assign, "TSExportAssignment");
        } else if (this.ts_eatContextualWithState("as", 2, enterHead)) {
          const decl = node;
          this.expectContextual("namespace");
          decl.id = this.parseIdent();
          this.semicolon();
          this.importOrExportOuterKind = void 0;
          return this.finishNode(decl, "TSNamespaceExportDeclaration");
        } else {
          const lookahead2 = this.lookahead(2).type;
          if (this.ts_isContextualWithState(enterHead, tokTypes2.type) && (lookahead2 === tt.braceL || // export type { ... }
          lookahead2 === tt.star)) {
            this.next();
            this.importOrExportOuterKind = "type";
            node.exportKind = "type";
          } else {
            this.importOrExportOuterKind = "value";
            node.exportKind = "value";
          }
          this.next();
          if (this.eat(tt.star)) {
            return this.parseExportAllDeclaration(node, exports);
          }
          if (this.eat(tt._default)) {
            this.checkExport(exports, "default", this.lastTokStart);
            node.declaration = this.parseExportDefaultDeclaration();
            return this.finishNode(node, "ExportDefaultDeclaration");
          }
          if (this.shouldParseExportStatement()) {
            node.declaration = this.parseExportDeclaration(node);
            if (node.declaration.type === "VariableDeclaration")
              this.checkVariableExport(exports, node.declaration.declarations);
            else this.checkExport(exports, node.declaration.id, node.declaration.id.start);
            node.specifiers = [];
            node.source = null;
          } else {
            node.declaration = null;
            node.specifiers = this.parseExportSpecifiers(exports);
            if (this.eatContextual("from")) {
              if (this.type !== tt.string) this.unexpected();
              node.source = this.parseExprAtom();
              this.parseMaybeImportAttributes(node);
            } else {
              for (let spec of node.specifiers) {
                this.checkUnreserved(spec.local);
                this.checkLocalExport(spec.local);
                if (spec.local.type === "Literal") {
                  this.raise(
                    spec.local.start,
                    "A string literal cannot be used as an exported binding without `from`."
                  );
                }
              }
              node.source = null;
            }
            this.semicolon();
          }
          return this.finishNode(node, "ExportNamedDeclaration");
        }
      }
      checkExport(exports, name, _) {
        if (!exports) {
          return;
        }
        if (typeof name !== "string") {
          name = name.type === "Identifier" ? name.name : name.value;
        }
        exports[name] = true;
      }
      parseMaybeDefault(startPos, startLoc, left) {
        const node = super.parseMaybeDefault(startPos, startLoc, left);
        if (node.type === "AssignmentPattern" && node.typeAnnotation && node.right.start < node.typeAnnotation.start) {
          this.raise(node.typeAnnotation.start, TypeScriptError.TypeAnnotationAfterAssign);
        }
        return node;
      }
      typeCastToParameter(node) {
        node.expression.typeAnnotation = node.typeAnnotation;
        this.resetEndLocation(
          node.expression,
          node.typeAnnotation.end,
          node.typeAnnotation.loc?.end
        );
        return node.expression;
      }
      toAssignableList(exprList, isBinding) {
        if (!exprList) exprList = [];
        for (let i = 0; i < exprList.length; i++) {
          const expr = exprList[i];
          if (expr?.type === "TSTypeCastExpression") {
            exprList[i] = this.typeCastToParameter(expr);
          }
        }
        return super.toAssignableList(exprList, isBinding);
      }
      reportReservedArrowTypeParam(node) {
        if (node.params.length === 1 && !node.extra?.trailingComma && disallowAmbiguousJSXLike) {
          this.raise(node.start, TypeScriptError.ReservedArrowTypeParam);
        }
      }
      parseExprAtom(refDestructuringErrors, forInit, forNew) {
        if (this.type === tokTypes2.jsxText) {
          return this.jsx_parseText();
        } else if (this.type === tokTypes2.jsxTagStart) {
          return this.jsx_parseElement();
        } else if (this.type === tokTypes2.at) {
          this.parseDecorators();
          return this.parseExprAtom();
        } else if (tokenIsIdentifier(this.type)) {
          let canBeArrow = this.potentialArrowAt === this.start;
          let startPos = this.start, startLoc = this.startLoc, containsEsc = this.containsEsc;
          let id = this.parseIdent(false);
          if (this.ecmaVersion >= 8 && !containsEsc && id.name === "async" && !this.canInsertSemicolon() && this.eat(tt._function)) {
            this.overrideContext(tokContexts.f_expr);
            return this.parseFunction(
              this.startNodeAt(startPos, startLoc),
              0,
              false,
              true,
              forInit
            );
          }
          if (canBeArrow && !this.canInsertSemicolon()) {
            if (this.eat(tt.arrow))
              return this.parseArrowExpression(
                this.startNodeAt(startPos, startLoc),
                [id],
                false,
                forInit
              );
            if (this.ecmaVersion >= 8 && id.name === "async" && this.type === tt.name && !containsEsc && (!this.potentialArrowInForAwait || this.value !== "of" || this.containsEsc)) {
              id = this.parseIdent(false);
              if (this.canInsertSemicolon() || !this.eat(tt.arrow)) this.unexpected();
              return this.parseArrowExpression(
                this.startNodeAt(startPos, startLoc),
                [id],
                true,
                forInit
              );
            }
          }
          return id;
        } else {
          return super.parseExprAtom(refDestructuringErrors, forInit, forNew);
        }
      }
      parseExprAtomDefault() {
        if (tokenIsIdentifier(this.type)) {
          const canBeArrow = this["potentialArrowAt"] === this.start;
          const containsEsc = this.containsEsc;
          const id = this.parseIdent();
          if (!containsEsc && id.name === "async" && !this.canInsertSemicolon()) {
            const { type } = this;
            if (type === tt._function) {
              this.next();
              return this.parseFunction(this.startNodeAtNode(id), void 0, true, true);
            } else if (tokenIsIdentifier(type)) {
              if (this.lookaheadCharCode() === 61) {
                const paramId = this.parseIdent(false);
                if (this.canInsertSemicolon() || !this.eat(tt.arrow)) this.unexpected();
                return this.parseArrowExpression(this.startNodeAtNode(id), [paramId], true);
              } else {
                return id;
              }
            }
          }
          if (canBeArrow && this.match(tt.arrow) && !this.canInsertSemicolon()) {
            this.next();
            return this.parseArrowExpression(this.startNodeAtNode(id), [id], false);
          }
          return id;
        } else {
          this.unexpected();
        }
      }
      parseIdentNode() {
        let node = this.startNode();
        if (tokenIsKeywordOrIdentifier(this.type) && // Taken from super-class method
        !((this.type.keyword === "class" || this.type.keyword === "function") && (this.lastTokEnd !== this.lastTokStart + 1 || this.input.charCodeAt(this.lastTokStart) !== 46))) {
          node.name = this.value;
        } else {
          return super.parseIdentNode();
        }
        return node;
      }
      parseVarStatement(node, kind, allowMissingInitializer = false) {
        const { isAmbientContext } = this;
        this.next();
        super.parseVar(node, false, kind, allowMissingInitializer || isAmbientContext);
        this.semicolon();
        const declaration = this.finishNode(node, "VariableDeclaration");
        if (!isAmbientContext) return declaration;
        for (const { id, init } of declaration.declarations) {
          if (!init) continue;
          if (kind !== "const" || !!id.typeAnnotation) {
            this.raise(init.start, TypeScriptError.InitializerNotAllowedInAmbientContext);
          } else if (init.type !== "StringLiteral" && init.type !== "BooleanLiteral" && init.type !== "NumericLiteral" && init.type !== "BigIntLiteral" && (init.type !== "TemplateLiteral" || init.expressions.length > 0) && !isPossiblyLiteralEnum(init)) {
            this.raise(
              init.start,
              TypeScriptError.ConstInitiailizerMustBeStringOrNumericLiteralOrLiteralEnumReference
            );
          }
        }
        return declaration;
      }
      parseStatement(context, topLevel, exports) {
        if (this.match(tokTypes2.at)) {
          this.parseDecorators(true);
        }
        if (this.match(tt._const) && this.isLookaheadContextual("enum")) {
          const node = this.startNode();
          this.expect(tt._const);
          return this.tsParseEnumDeclaration(node, { const: true });
        }
        if (this.ts_isContextual(tokTypes2.enum)) {
          return this.tsParseEnumDeclaration(this.startNode());
        }
        if (this.ts_isContextual(tokTypes2.interface)) {
          const result = this.tsParseInterfaceDeclaration(this.startNode());
          if (result) return result;
        }
        return super.parseStatement(context, topLevel, exports);
      }
      // NOTE: unused function
      parseAccessModifier() {
        return this.tsParseModifier(["public", "protected", "private"]);
      }
      parsePostMemberNameModifiers(methodOrProp) {
        const optional = this.eat(tt.question);
        if (optional) methodOrProp.optional = true;
        if (methodOrProp.readonly && this.match(tt.parenL)) {
          this.raise(methodOrProp.start, TypeScriptError.ClassMethodHasReadonly);
        }
        if (methodOrProp.declare && this.match(tt.parenL)) {
          this.raise(methodOrProp.start, TypeScriptError.ClassMethodHasDeclare);
        }
      }
      // Note: The reason we do this in `parseExpressionStatement` and not `parseStatement`
      // is that e.g. `type()` is valid JS, so we must try parsing that first.
      // If it's really a type, we will parse `type` as the statement, and can correct it here
      // by parsing the rest.
      parseExpressionStatement(node, expr) {
        const decl = expr.type === "Identifier" ? this.tsParseExpressionStatement(node, expr) : void 0;
        return decl || super.parseExpressionStatement(node, expr);
      }
      shouldParseExportStatement() {
        if (this.tsIsDeclarationStart()) return true;
        if (this.match(tokTypes2.at)) {
          return true;
        }
        return super.shouldParseExportStatement();
      }
      parseConditional(expr, startPos, startLoc, forInit, refDestructuringErrors) {
        if (this.eat(tt.question)) {
          let node = this.startNodeAt(startPos, startLoc);
          node.test = expr;
          node.consequent = this.parseMaybeAssign();
          this.expect(tt.colon);
          node.alternate = this.parseMaybeAssign(forInit);
          return this.finishNode(node, "ConditionalExpression");
        }
        return expr;
      }
      parseMaybeConditional(forInit, refDestructuringErrors) {
        let startPos = this.start, startLoc = this.startLoc;
        let expr = this.parseExprOps(forInit, refDestructuringErrors);
        if (this.checkExpressionErrors(refDestructuringErrors)) return expr;
        if (!this.maybeInArrowParameters || !this.match(tt.question)) {
          return this.parseConditional(expr, startPos, startLoc, forInit, refDestructuringErrors);
        }
        const result = this.tryParse(
          () => this.parseConditional(expr, startPos, startLoc, forInit, refDestructuringErrors)
        );
        if (!result.node) {
          if (result.error) {
            this.setOptionalParametersError(refDestructuringErrors, result.error);
          }
          return expr;
        }
        if (result.error) this.setLookaheadState(result.failState);
        return result.node;
      }
      parseParenItem(node) {
        const startPos = this.start;
        const startLoc = this.startLoc;
        node = super.parseParenItem(node);
        if (this.eat(tt.question)) {
          node.optional = true;
          this.resetEndLocation(node);
        }
        if (this.match(tt.colon)) {
          const typeCastNode = this.startNodeAt(startPos, startLoc);
          typeCastNode.expression = node;
          typeCastNode.typeAnnotation = this.tsParseTypeAnnotation();
          return this.finishNode(typeCastNode, "TSTypeCastExpression");
        }
        return node;
      }
      parseExportDeclaration(node) {
        if (!this.isAmbientContext && this.ts_isContextual(tokTypes2.declare)) {
          return this.tsInAmbientContext(() => this.parseExportDeclaration(node));
        }
        const startPos = this.start;
        const startLoc = this.startLoc;
        const isDeclare = this.eatContextual("declare");
        if (isDeclare && (this.ts_isContextual(tokTypes2.declare) || !this.shouldParseExportStatement())) {
          this.raise(this.start, TypeScriptError.ExpectedAmbientAfterExportDeclare);
        }
        const isIdentifier = tokenIsIdentifier(this.type);
        const declaration = isIdentifier && this.tsTryParseExportDeclaration() || this.parseStatement(null);
        if (!declaration) return null;
        if (declaration.type === "TSInterfaceDeclaration" || declaration.type === "TSTypeAliasDeclaration" || isDeclare) {
          node.exportKind = "type";
        }
        if (isDeclare) {
          this.resetStartLocation(declaration, startPos, startLoc);
          declaration.declare = true;
        }
        return declaration;
      }
      parseClassId(node, isStatement) {
        if (!isStatement && this.isContextual("implements")) {
          return;
        }
        super.parseClassId(node, isStatement);
        const typeParameters = this.tsTryParseTypeParameters(this.tsParseInOutModifiers.bind(this));
        if (typeParameters) node.typeParameters = typeParameters;
      }
      parseClassPropertyAnnotation(node) {
        if (!node.optional) {
          if (this.value === "!" && this.eat(tt.prefix)) {
            node.definite = true;
          } else if (this.eat(tt.question)) {
            node.optional = true;
          }
        }
        const type = this.tsTryParseTypeAnnotation();
        if (type) node.typeAnnotation = type;
      }
      parseClassField(field) {
        const isPrivate = field.key.type === "PrivateIdentifier";
        if (isPrivate) {
          if (field.abstract) {
            this.raise(field.start, TypeScriptError.PrivateElementHasAbstract);
          }
          if (field.accessibility) {
            this.raise(
              field.start,
              TypeScriptError.PrivateElementHasAccessibility({
                modifier: field.accessibility
              })
            );
          }
          this.parseClassPropertyAnnotation(field);
        } else {
          this.parseClassPropertyAnnotation(field);
          if (this.isAmbientContext && !(field.readonly && !field.typeAnnotation) && this.match(tt.eq)) {
            this.raise(this.start, TypeScriptError.DeclareClassFieldHasInitializer);
          }
          if (field.abstract && this.match(tt.eq)) {
            const { key } = field;
            this.raise(
              this.start,
              TypeScriptError.AbstractPropertyHasInitializer({
                propertyName: key.type === "Identifier" && !field.computed ? key.name : `[${this.input.slice(key.start, key.end)}]`
              })
            );
          }
        }
        return super.parseClassField(field);
      }
      parseClassMethod(method, isGenerator, isAsync, allowsDirectSuper) {
        const isConstructor = method.kind === "constructor";
        const isPrivate = method.key.type === "PrivateIdentifier";
        const typeParameters = this.tsTryParseTypeParameters(this.tsParseConstModifier);
        if (isPrivate) {
          if (typeParameters) method.typeParameters = typeParameters;
          if (method.accessibility) {
            this.raise(
              method.start,
              TypeScriptError.PrivateMethodsHasAccessibility({
                modifier: method.accessibility
              })
            );
          }
        } else {
          if (typeParameters && isConstructor) {
            this.raise(typeParameters.start, TypeScriptError.ConstructorHasTypeParameters);
          }
        }
        const { declare = false, kind } = method;
        if (declare && (kind === "get" || kind === "set")) {
          this.raise(method.start, TypeScriptError.DeclareAccessor({ kind }));
        }
        if (typeParameters) method.typeParameters = typeParameters;
        const key = method.key;
        if (method.kind === "constructor") {
          if (isGenerator) this.raise(key.start, "Constructor can't be a generator");
          if (isAsync) this.raise(key.start, "Constructor can't be an async method");
        } else if (method.static && checkKeyName(method, "prototype")) {
          this.raise(key.start, "Classes may not have a static property named prototype");
        }
        const value = method.value = this.parseMethod(
          isGenerator,
          isAsync,
          allowsDirectSuper,
          true,
          method
        );
        if (method.kind === "get" && value["params"].length !== 0)
          this.raiseRecoverable(value.start, "getter should have no params");
        if (method.kind === "set" && value["params"].length !== 1)
          this.raiseRecoverable(value.start, "setter should have exactly one param");
        if (method.kind === "set" && value["params"][0].type === "RestElement")
          this.raiseRecoverable(value["params"][0].start, "Setter cannot use rest params");
        return this.finishNode(method, "MethodDefinition");
      }
      isClassMethod() {
        return this.match(tt.relational);
      }
      parseClassElement(constructorAllowsSuper) {
        if (this.eat(tt.semi)) return null;
        let node = this.startNode();
        let keyName = "";
        let isGenerator = false;
        let isAsync = false;
        let kind = "method";
        let isStatic = false;
        const modifiers = [
          "declare",
          "private",
          "public",
          "protected",
          "accessor",
          "override",
          "abstract",
          "readonly",
          "static"
        ];
        const modifierMap = this.tsParseModifiers({
          modified: node,
          allowedModifiers: modifiers,
          disallowedModifiers: ["in", "out"],
          stopOnStartOfClassStaticBlock: true,
          errorTemplate: TypeScriptError.InvalidModifierOnTypeParameterPositions
        });
        isStatic = Boolean(modifierMap.static);
        const callParseClassMemberWithIsStatic = () => {
          if (this.tsIsStartOfStaticBlocks()) {
            this.next();
            this.next();
            if (this.tsHasSomeModifiers(node, modifiers)) {
              this.raise(this.start, TypeScriptError.StaticBlockCannotHaveModifier);
            }
            if (this.ecmaVersion >= 13) {
              super.parseClassStaticBlock(node);
              return node;
            }
          } else {
            const idx = this.tsTryParseIndexSignature(node);
            if (idx) {
              if (node.abstract) {
                this.raise(node.start, TypeScriptError.IndexSignatureHasAbstract);
              }
              if (node.accessibility) {
                this.raise(
                  node.start,
                  TypeScriptError.IndexSignatureHasAccessibility({
                    modifier: node.accessibility
                  })
                );
              }
              if (node.declare) {
                this.raise(node.start, TypeScriptError.IndexSignatureHasDeclare);
              }
              if (node.override) {
                this.raise(node.start, TypeScriptError.IndexSignatureHasOverride);
              }
              return idx;
            }
            if (!this.inAbstractClass && node.abstract) {
              this.raise(node.start, TypeScriptError.NonAbstractClassHasAbstractMethod);
            }
            if (node.override) {
              if (!constructorAllowsSuper) {
                this.raise(node.start, TypeScriptError.OverrideNotInSubClass);
              }
            }
            node.static = isStatic;
            if (isStatic) {
              if (!(this.isClassElementNameStart() || this.type === tt.star)) {
                keyName = "static";
              }
            }
            if (!keyName && this.ecmaVersion >= 8 && this.eatContextual("async")) {
              if ((this.isClassElementNameStart() || this.type === tt.star) && !this.canInsertSemicolon()) {
                isAsync = true;
              } else {
                keyName = "async";
              }
            }
            if (!keyName && (this.ecmaVersion >= 9 || !isAsync) && this.eat(tt.star)) {
              isGenerator = true;
            }
            if (!keyName && !isAsync && !isGenerator) {
              const lastValue = this.value;
              if (this.eatContextual("get") || this.eatContextual("set")) {
                if (this.isClassElementNameStart()) {
                  kind = lastValue;
                } else {
                  keyName = lastValue;
                }
              }
            }
            if (keyName) {
              node.computed = false;
              node.key = this.startNodeAt(this.lastTokStart, this.lastTokStartLoc);
              node.key.name = keyName;
              this.finishNode(node.key, "Identifier");
            } else {
              this.parseClassElementName(node);
            }
            this.parsePostMemberNameModifiers(node);
            if (this.isClassMethod() || this.ecmaVersion < 13 || this.type === tt.parenL || kind !== "method" || isGenerator || isAsync) {
              const isConstructor = !node.static && checkKeyName(node, "constructor");
              const allowsDirectSuper = isConstructor && constructorAllowsSuper;
              if (isConstructor && kind !== "method")
                this.raise(node.key.start, "Constructor can't have get/set modifier");
              node.kind = isConstructor ? "constructor" : kind;
              this.parseClassMethod(node, isGenerator, isAsync, allowsDirectSuper);
            } else {
              this.parseClassField(node);
            }
            return node;
          }
        };
        if (node.declare) {
          this.tsInAmbientContext(callParseClassMemberWithIsStatic);
        } else {
          callParseClassMemberWithIsStatic();
        }
        return node;
      }
      isClassElementNameStart() {
        if (this.tsIsIdentifier()) {
          return true;
        }
        return super.isClassElementNameStart();
      }
      parseClassSuper(node) {
        super.parseClassSuper(node);
        if (node.superClass && (this.tsMatchLeftRelational() || this.match(tt.bitShift))) {
          node.superTypeParameters = this.tsParseTypeArgumentsInExpression();
        }
        if (this.eatContextual("implements")) {
          node.implements = this.tsParseHeritageClause("implements");
        }
      }
      parseFunctionParams(node) {
        const typeParameters = this.tsTryParseTypeParameters(this.tsParseConstModifier);
        if (typeParameters) node.typeParameters = typeParameters;
        super.parseFunctionParams(node);
      }
      // `let x: number;`
      parseVarId(decl, kind) {
        super.parseVarId(decl, kind);
        if (decl.id.type === "Identifier" && !this.hasPrecedingLineBreak() && this.value === "!" && this.eat(tt.prefix)) {
          decl.definite = true;
        }
        const type = this.tsTryParseTypeAnnotation();
        if (type) {
          decl.id.typeAnnotation = type;
          this.resetEndLocation(decl.id);
        }
      }
      // parse the return type of an async arrow function - let foo = (async (): number => {});
      parseArrowExpression(node, params, isAsync, forInit) {
        if (this.match(tt.colon)) {
          node.returnType = this.tsParseTypeAnnotation();
        }
        let oldYieldPos = this.yieldPos, oldAwaitPos = this.awaitPos, oldAwaitIdentPos = this.awaitIdentPos;
        this.enterScope(functionFlags(isAsync, false) | acornScope.SCOPE_ARROW);
        this.initFunction(node);
        const oldMaybeInArrowParameters = this.maybeInArrowParameters;
        if (this.ecmaVersion >= 8) node.async = !!isAsync;
        this.yieldPos = 0;
        this.awaitPos = 0;
        this.awaitIdentPos = 0;
        this.maybeInArrowParameters = true;
        node.params = this.toAssignableList(params, true);
        this.maybeInArrowParameters = false;
        this.parseFunctionBody(node, true, false, forInit);
        this.yieldPos = oldYieldPos;
        this.awaitPos = oldAwaitPos;
        this.awaitIdentPos = oldAwaitIdentPos;
        this.maybeInArrowParameters = oldMaybeInArrowParameters;
        return this.finishNode(node, "ArrowFunctionExpression");
      }
      parseMaybeAssignOrigin(forInit, refDestructuringErrors, afterLeftParse) {
        if (this.isContextual("yield")) {
          if (this.inGenerator) return this.parseYield(forInit);
          else this.exprAllowed = false;
        }
        let ownDestructuringErrors = false, oldParenAssign = -1, oldTrailingComma = -1, oldDoubleProto = -1;
        if (refDestructuringErrors) {
          oldParenAssign = refDestructuringErrors.parenthesizedAssign;
          oldTrailingComma = refDestructuringErrors.trailingComma;
          oldDoubleProto = refDestructuringErrors.doubleProto;
          refDestructuringErrors.parenthesizedAssign = refDestructuringErrors.trailingComma = -1;
        } else {
          refDestructuringErrors = new DestructuringErrors();
          ownDestructuringErrors = true;
        }
        let startPos = this.start, startLoc = this.startLoc;
        if (this.type === tt.parenL || tokenIsIdentifier(this.type)) {
          this.potentialArrowAt = this.start;
          this.potentialArrowInForAwait = forInit === "await";
        }
        let left = this.parseMaybeConditional(forInit, refDestructuringErrors);
        if (afterLeftParse) left = afterLeftParse.call(this, left, startPos, startLoc);
        if (this.type.isAssign) {
          let node = this.startNodeAt(startPos, startLoc);
          node.operator = this.value;
          if (this.type === tt.eq) left = this.toAssignable(left, true, refDestructuringErrors);
          if (!ownDestructuringErrors) {
            refDestructuringErrors.parenthesizedAssign = refDestructuringErrors.trailingComma = refDestructuringErrors.doubleProto = -1;
          }
          if (refDestructuringErrors.shorthandAssign >= left.start)
            refDestructuringErrors.shorthandAssign = -1;
          if (!this.maybeInArrowParameters) {
            if (this.type === tt.eq) this.checkLValPattern(left);
            else this.checkLValSimple(left);
          }
          node.left = left;
          this.next();
          node.right = this.parseMaybeAssign(forInit);
          if (oldDoubleProto > -1) refDestructuringErrors.doubleProto = oldDoubleProto;
          return this.finishNode(node, "AssignmentExpression");
        } else {
          if (ownDestructuringErrors) this.checkExpressionErrors(refDestructuringErrors, true);
        }
        if (oldParenAssign > -1) refDestructuringErrors.parenthesizedAssign = oldParenAssign;
        if (oldTrailingComma > -1) refDestructuringErrors.trailingComma = oldTrailingComma;
        return left;
      }
      parseMaybeAssign(forInit, refExpressionErrors, afterLeftParse) {
        let state;
        let jsx;
        let typeCast;
        if (options?.jsx && (this.matchJsx("jsxTagStart") || this.tsMatchLeftRelational())) {
          state = this.cloneCurLookaheadState();
          jsx = this.tryParse(
            () => this.parseMaybeAssignOrigin(forInit, refExpressionErrors, afterLeftParse),
            state
          );
          if (!jsx.error) return jsx.node;
          const context = this.context;
          const currentContext = context[context.length - 1];
          const lastCurrentContext = context[context.length - 2];
          if (currentContext === acornTypeScript.tokContexts.tc_oTag && lastCurrentContext === acornTypeScript.tokContexts.tc_expr) {
            context.pop();
            context.pop();
          } else if (currentContext === acornTypeScript.tokContexts.tc_oTag || currentContext === acornTypeScript.tokContexts.tc_expr) {
            context.pop();
          }
        }
        if (!jsx?.error && !this.tsMatchLeftRelational()) {
          return this.parseMaybeAssignOrigin(forInit, refExpressionErrors, afterLeftParse);
        }
        if (!state || this.compareLookaheadState(state, this.getCurLookaheadState())) {
          state = this.cloneCurLookaheadState();
        }
        let typeParameters;
        const arrow = this.tryParse((abort) => {
          typeParameters = this.tsParseTypeParameters(this.tsParseConstModifier);
          const expr = this.parseMaybeAssignOrigin(forInit, refExpressionErrors, afterLeftParse);
          if (expr.type !== "ArrowFunctionExpression" || expr.extra?.parenthesized) {
            abort();
          }
          if (typeParameters?.params.length !== 0) {
            this.resetStartLocationFromNode(expr, typeParameters);
          }
          expr.typeParameters = typeParameters;
          return expr;
        }, state);
        if (!arrow.error && !arrow.aborted) {
          if (typeParameters) this.reportReservedArrowTypeParam(typeParameters);
          return arrow.node;
        }
        if (!jsx) {
          assert(true);
          typeCast = this.tryParse(
            () => this.parseMaybeAssignOrigin(forInit, refExpressionErrors, afterLeftParse),
            state
          );
          if (!typeCast.error) return typeCast.node;
        }
        if (jsx?.node) {
          this.setLookaheadState(jsx.failState);
          return jsx.node;
        }
        if (arrow.node) {
          this.setLookaheadState(arrow.failState);
          if (typeParameters) this.reportReservedArrowTypeParam(typeParameters);
          return arrow.node;
        }
        if (typeCast?.node) {
          this.setLookaheadState(typeCast.failState);
          return typeCast.node;
        }
        if (jsx?.thrown) throw jsx.error;
        if (arrow.thrown) throw arrow.error;
        if (typeCast?.thrown) throw typeCast.error;
        throw jsx?.error || arrow.error || typeCast?.error;
      }
      parseAssignableListItem(allowModifiers) {
        const decorators = [];
        while (this.match(tokTypes2.at)) {
          decorators.push(this.parseDecorator());
        }
        const startPos = this.start;
        const startLoc = this.startLoc;
        let accessibility;
        let readonly = false;
        let override = false;
        if (allowModifiers !== void 0) {
          const modified = {};
          this.tsParseModifiers({
            modified,
            allowedModifiers: ["public", "private", "protected", "override", "readonly"]
          });
          accessibility = modified.accessibility;
          override = modified.override;
          readonly = modified.readonly;
          if (allowModifiers === false && (accessibility || readonly || override)) {
            this.raise(startLoc.column, TypeScriptError.UnexpectedParameterModifier);
          }
        }
        const left = this.parseMaybeDefault(startPos, startLoc);
        this.parseBindingListItem(left);
        const elt = this.parseMaybeDefault(left["start"], left["loc"], left);
        if (decorators.length) {
          elt.decorators = decorators;
        }
        if (accessibility || readonly || override) {
          const pp = this.startNodeAt(startPos, startLoc);
          if (accessibility) pp.accessibility = accessibility;
          if (readonly) pp.readonly = readonly;
          if (override) pp.override = override;
          if (elt.type !== "Identifier" && elt.type !== "AssignmentPattern") {
            this.raise(pp.start, TypeScriptError.UnsupportedParameterPropertyKind);
          }
          pp.parameter = elt;
          return this.finishNode(pp, "TSParameterProperty");
        }
        return elt;
      }
      // AssignmentPattern
      checkLValInnerPattern(expr, bindingType = acornScope.BIND_NONE, checkClashes) {
        switch (expr.type) {
          case "TSParameterProperty":
            this.checkLValInnerPattern(expr.parameter, bindingType, checkClashes);
            break;
          default: {
            super.checkLValInnerPattern(expr, bindingType, checkClashes);
            break;
          }
        }
      }
      // Allow type annotations inside of a parameter list.
      parseBindingListItem(param) {
        if (this.eat(tt.question)) {
          if (param.type !== "Identifier" && !this.isAmbientContext && !this.inType) {
            this.raise(param.start, TypeScriptError.PatternIsOptional);
          }
          param.optional = true;
        }
        const type = this.tsTryParseTypeAnnotation();
        if (type) param.typeAnnotation = type;
        this.resetEndLocation(param);
        return param;
      }
      isAssignable(node, isBinding) {
        switch (node.type) {
          case "TSTypeCastExpression":
            return this.isAssignable(node.expression, isBinding);
          case "TSParameterProperty":
            return true;
          case "Identifier":
          case "ObjectPattern":
          case "ArrayPattern":
          case "AssignmentPattern":
          case "RestElement":
            return true;
          case "ObjectExpression": {
            const last = node.properties.length - 1;
            return node.properties.every((prop, i) => {
              return prop.type !== "ObjectMethod" && (i === last || prop.type !== "SpreadElement") && this.isAssignable(prop);
            });
          }
          case "Property":
          case "ObjectProperty":
            return this.isAssignable(node.value);
          case "SpreadElement":
            return this.isAssignable(node.argument);
          case "ArrayExpression":
            return node.elements.every(
              (element) => element === null || this.isAssignable(element)
            );
          case "AssignmentExpression":
            return node.operator === "=";
          case "ParenthesizedExpression":
            return this.isAssignable(node.expression);
          case "MemberExpression":
          case "OptionalMemberExpression":
            return !isBinding;
          default:
            return false;
        }
      }
      toAssignable(node, isBinding = false, refDestructuringErrors = new DestructuringErrors()) {
        switch (node.type) {
          case "ParenthesizedExpression":
            return this.toAssignableParenthesizedExpression(
              node,
              isBinding,
              refDestructuringErrors
            );
          case "TSAsExpression":
          case "TSSatisfiesExpression":
          case "TSNonNullExpression":
          case "TSTypeAssertion":
            if (isBinding) {
            } else {
              this.raise(node.start, TypeScriptError.UnexpectedTypeCastInParameter);
            }
            return this.toAssignable(node.expression, isBinding, refDestructuringErrors);
          case "MemberExpression":
            break;
          case "AssignmentExpression":
            if (!isBinding && node.left.type === "TSTypeCastExpression") {
              node.left = this.typeCastToParameter(node.left);
            }
            return super.toAssignable(node, isBinding, refDestructuringErrors);
          case "TSTypeCastExpression": {
            return this.typeCastToParameter(node);
          }
          default:
            return super.toAssignable(node, isBinding, refDestructuringErrors);
        }
        return node;
      }
      toAssignableParenthesizedExpression(node, isBinding, refDestructuringErrors) {
        switch (node.expression.type) {
          case "TSAsExpression":
          case "TSSatisfiesExpression":
          case "TSNonNullExpression":
          case "TSTypeAssertion":
          case "ParenthesizedExpression":
            return this.toAssignable(node.expression, isBinding, refDestructuringErrors);
          default:
            return super.toAssignable(node, isBinding, refDestructuringErrors);
        }
      }
      parseBindingAtom() {
        switch (this.type) {
          case tt._this:
            return this.parseIdent(
              /* liberal */
              true
            );
          default:
            return super.parseBindingAtom();
        }
      }
      shouldParseArrow(exprList) {
        let shouldParseArrowRes;
        if (this.match(tt.colon)) {
          shouldParseArrowRes = exprList.every((expr) => this.isAssignable(expr, true));
        } else {
          shouldParseArrowRes = !this.canInsertSemicolon();
        }
        if (shouldParseArrowRes) {
          if (this.match(tt.colon)) {
            const result = this.tryParse((abort) => {
              const returnType = this.tsParseTypeOrTypePredicateAnnotation(tt.colon);
              if (this.canInsertSemicolon() || !this.match(tt.arrow)) abort();
              return returnType;
            });
            if (result.aborted) {
              this.shouldParseArrowReturnType = void 0;
              return false;
            }
            if (!result.thrown) {
              if (result.error) this.setLookaheadState(result.failState);
              this.shouldParseArrowReturnType = result.node;
            }
          }
          if (!this.match(tt.arrow)) {
            this.shouldParseArrowReturnType = void 0;
            return false;
          }
          return true;
        }
        this.shouldParseArrowReturnType = void 0;
        return shouldParseArrowRes;
      }
      parseParenArrowList(startPos, startLoc, exprList, forInit) {
        const node = this.startNodeAt(startPos, startLoc);
        node.returnType = this.shouldParseArrowReturnType;
        this.shouldParseArrowReturnType = void 0;
        return this.parseArrowExpression(node, exprList, false, forInit);
      }
      parseParenAndDistinguishExpression(canBeArrow, forInit) {
        let startPos = this.start, startLoc = this.startLoc, val, allowTrailingComma = this.ecmaVersion >= 8;
        if (this.ecmaVersion >= 6) {
          const oldMaybeInArrowParameters = this.maybeInArrowParameters;
          this.maybeInArrowParameters = true;
          this.next();
          let innerStartPos = this.start, innerStartLoc = this.startLoc;
          let exprList = [], first = true, lastIsComma = false;
          let refDestructuringErrors = new DestructuringErrors(), oldYieldPos = this.yieldPos, oldAwaitPos = this.awaitPos, spreadStart;
          this.yieldPos = 0;
          this.awaitPos = 0;
          while (this.type !== tt.parenR) {
            first ? first = false : this.expect(tt.comma);
            if (allowTrailingComma && this.afterTrailingComma(tt.parenR, true)) {
              lastIsComma = true;
              break;
            } else if (this.type === tt.ellipsis) {
              spreadStart = this.start;
              exprList.push(this.parseParenItem(this.parseRestBinding()));
              if (this.type === tt.comma) {
                this.raise(this.start, "Comma is not permitted after the rest element");
              }
              break;
            } else {
              exprList.push(
                this.parseMaybeAssign(forInit, refDestructuringErrors, this.parseParenItem)
              );
            }
          }
          let innerEndPos = this.lastTokEnd, innerEndLoc = this.lastTokEndLoc;
          this.expect(tt.parenR);
          this.maybeInArrowParameters = oldMaybeInArrowParameters;
          if (canBeArrow && this.shouldParseArrow(exprList) && this.eat(tt.arrow)) {
            this.checkPatternErrors(refDestructuringErrors, false);
            this.checkYieldAwaitInDefaultParams();
            this.yieldPos = oldYieldPos;
            this.awaitPos = oldAwaitPos;
            return this.parseParenArrowList(startPos, startLoc, exprList, forInit);
          }
          if (!exprList.length || lastIsComma) this.unexpected(this.lastTokStart);
          if (spreadStart) this.unexpected(spreadStart);
          this.checkExpressionErrors(refDestructuringErrors, true);
          this.yieldPos = oldYieldPos || this.yieldPos;
          this.awaitPos = oldAwaitPos || this.awaitPos;
          if (exprList.length > 1) {
            val = this.startNodeAt(innerStartPos, innerStartLoc);
            val.expressions = exprList;
            this.finishNodeAt(val, "SequenceExpression", innerEndPos, innerEndLoc);
          } else {
            val = exprList[0];
          }
        } else {
          val = this.parseParenExpression();
        }
        if (this.options.preserveParens) {
          let par = this.startNodeAt(startPos, startLoc);
          par.expression = val;
          return this.finishNode(par, "ParenthesizedExpression");
        } else {
          return val;
        }
      }
      parseTaggedTemplateExpression(base, startPos, startLoc, optionalChainMember) {
        const node = this.startNodeAt(startPos, startLoc);
        node.tag = base;
        node.quasi = this.parseTemplate({ isTagged: true });
        if (optionalChainMember) {
          this.raise(
            startPos,
            "Tagged Template Literals are not allowed in optionalChain."
          );
        }
        return this.finishNode(node, "TaggedTemplateExpression");
      }
      shouldParseAsyncArrow() {
        if (this.match(tt.colon)) {
          const result = this.tryParse((abort) => {
            const returnType = this.tsParseTypeOrTypePredicateAnnotation(tt.colon);
            if (this.canInsertSemicolon() || !this.match(tt.arrow)) abort();
            return returnType;
          });
          if (result.aborted) {
            this.shouldParseAsyncArrowReturnType = void 0;
            return false;
          }
          if (!result.thrown) {
            if (result.error) this.setLookaheadState(result.failState);
            this.shouldParseAsyncArrowReturnType = result.node;
            return !this.canInsertSemicolon() && this.eat(tt.arrow);
          }
        } else {
          return !this.canInsertSemicolon() && this.eat(tt.arrow);
        }
      }
      parseSubscriptAsyncArrow(startPos, startLoc, exprList, forInit) {
        const arrN = this.startNodeAt(startPos, startLoc);
        arrN.returnType = this.shouldParseAsyncArrowReturnType;
        this.shouldParseAsyncArrowReturnType = void 0;
        return this.parseArrowExpression(arrN, exprList, true, forInit);
      }
      parseExprList(close, allowTrailingComma, allowEmpty, refDestructuringErrors) {
        let elts = [], first = true;
        while (!this.eat(close)) {
          if (!first) {
            this.expect(tt.comma);
            if (allowTrailingComma && this.afterTrailingComma(close)) break;
          } else first = false;
          let elt;
          if (allowEmpty && this.type === tt.comma) elt = null;
          else if (this.type === tt.ellipsis) {
            elt = this.parseSpread(refDestructuringErrors);
            if (this.maybeInArrowParameters && this.match(tt.colon)) {
              elt.typeAnnotation = this.tsParseTypeAnnotation();
            }
            if (refDestructuringErrors && this.type === tt.comma && refDestructuringErrors.trailingComma < 0)
              refDestructuringErrors.trailingComma = this.start;
          } else {
            elt = this.parseMaybeAssign(false, refDestructuringErrors, this.parseParenItem);
          }
          elts.push(elt);
        }
        return elts;
      }
      parseSubscript(base, startPos, startLoc, noCalls, maybeAsyncArrow, optionalChained, forInit) {
        let _optionalChained = optionalChained;
        if (!this.hasPrecedingLineBreak() && // NODE: replace bang
        this.value === "!" && this.match(tt.prefix)) {
          this.exprAllowed = false;
          this.next();
          const nonNullExpression = this.startNodeAt(startPos, startLoc);
          nonNullExpression.expression = base;
          base = this.finishNode(nonNullExpression, "TSNonNullExpression");
          return base;
        }
        let isOptionalCall = false;
        if (this.match(tt.questionDot) && this.lookaheadCharCode() === 60) {
          if (noCalls) {
            return base;
          }
          base.optional = true;
          _optionalChained = isOptionalCall = true;
          this.next();
        }
        if (this.tsMatchLeftRelational() || this.match(tt.bitShift)) {
          let missingParenErrorLoc;
          const result = this.tsTryParseAndCatch(() => {
            if (!noCalls && this.atPossibleAsyncArrow(base)) {
              const asyncArrowFn = this.tsTryParseGenericAsyncArrowFunction(
                startPos,
                startLoc,
                forInit
              );
              if (asyncArrowFn) {
                base = asyncArrowFn;
                return base;
              }
            }
            const typeArguments = this.tsParseTypeArgumentsInExpression();
            if (!typeArguments) return base;
            if (isOptionalCall && !this.match(tt.parenL)) {
              missingParenErrorLoc = this.curPosition();
              return base;
            }
            if (tokenIsTemplate(this.type) || this.type === tt.backQuote) {
              const result2 = this.parseTaggedTemplateExpression(
                base,
                startPos,
                startLoc,
                _optionalChained
              );
              result2.typeArguments = typeArguments;
              return result2;
            }
            if (!noCalls && this.eat(tt.parenL)) {
              let refDestructuringErrors = new DestructuringErrors();
              const node2 = this.startNodeAt(startPos, startLoc);
              node2.callee = base;
              node2.arguments = this.parseExprList(
                tt.parenR,
                this.ecmaVersion >= 8,
                false,
                refDestructuringErrors
              );
              this.tsCheckForInvalidTypeCasts(node2.arguments);
              node2.typeArguments = typeArguments;
              if (_optionalChained) {
                node2.optional = isOptionalCall;
              }
              this.checkExpressionErrors(refDestructuringErrors, true);
              base = this.finishNode(node2, "CallExpression");
              return base;
            }
            const tokenType = this.type;
            if (
              // a<b>>c is not (a<b>)>c, but a<(b>>c)
              this.tsMatchRightRelational() || // a<b>>>c is not (a<b>)>>c, but a<(b>>>c)
              tokenType === tt.bitShift || // a<b>c is (a<b)>c
              tokenType !== tt.parenL && tokenCanStartExpression(tokenType) && !this.hasPrecedingLineBreak()
            ) {
              return;
            }
            const node = this.startNodeAt(startPos, startLoc);
            node.expression = base;
            node.typeArguments = typeArguments;
            return this.finishNode(node, "TSInstantiationExpression");
          });
          if (missingParenErrorLoc) {
            this.unexpected(missingParenErrorLoc);
          }
          if (result) {
            if (result.type === "TSInstantiationExpression" && (this.match(tt.dot) || this.match(tt.questionDot) && this.lookaheadCharCode() !== 40)) {
              this.raise(
                this.start,
                TypeScriptError.InvalidPropertyAccessAfterInstantiationExpression
              );
            }
            base = result;
            return base;
          }
        }
        let optionalSupported = this.ecmaVersion >= 11;
        let optional = optionalSupported && this.eat(tt.questionDot);
        if (noCalls && optional)
          this.raise(
            this.lastTokStart,
            "Optional chaining cannot appear in the callee of new expressions"
          );
        let computed = this.eat(tt.bracketL);
        if (computed || optional && this.type !== tt.parenL && this.type !== tt.backQuote || this.eat(tt.dot)) {
          let node = this.startNodeAt(startPos, startLoc);
          node.object = base;
          if (computed) {
            node.property = this.parseExpression();
            this.expect(tt.bracketR);
          } else if (this.type === tt.privateId && base.type !== "Super") {
            node.property = this.parsePrivateIdent();
          } else {
            node.property = this.parseIdent(this.options.allowReserved !== "never");
          }
          node.computed = !!computed;
          if (optionalSupported) {
            node.optional = optional;
          }
          base = this.finishNode(node, "MemberExpression");
        } else if (!noCalls && this.eat(tt.parenL)) {
          const oldMaybeInArrowParameters = this.maybeInArrowParameters;
          this.maybeInArrowParameters = true;
          let refDestructuringErrors = new DestructuringErrors(), oldYieldPos = this.yieldPos, oldAwaitPos = this.awaitPos, oldAwaitIdentPos = this.awaitIdentPos;
          this.yieldPos = 0;
          this.awaitPos = 0;
          this.awaitIdentPos = 0;
          let exprList = this.parseExprList(
            tt.parenR,
            this.ecmaVersion >= 8,
            false,
            refDestructuringErrors
          );
          if (maybeAsyncArrow && !optional && this.shouldParseAsyncArrow()) {
            this.checkPatternErrors(refDestructuringErrors, false);
            this.checkYieldAwaitInDefaultParams();
            if (this.awaitIdentPos > 0)
              this.raise(
                this.awaitIdentPos,
                "Cannot use 'await' as identifier inside an async function"
              );
            this.yieldPos = oldYieldPos;
            this.awaitPos = oldAwaitPos;
            this.awaitIdentPos = oldAwaitIdentPos;
            base = this.parseSubscriptAsyncArrow(startPos, startLoc, exprList, forInit);
          } else {
            this.checkExpressionErrors(refDestructuringErrors, true);
            this.yieldPos = oldYieldPos || this.yieldPos;
            this.awaitPos = oldAwaitPos || this.awaitPos;
            this.awaitIdentPos = oldAwaitIdentPos || this.awaitIdentPos;
            let node = this.startNodeAt(startPos, startLoc);
            node.callee = base;
            node.arguments = exprList;
            if (optionalSupported) {
              node.optional = optional;
            }
            base = this.finishNode(node, "CallExpression");
          }
          this.maybeInArrowParameters = oldMaybeInArrowParameters;
        } else if (this.type === tt.backQuote) {
          if (optional || _optionalChained) {
            this.raise(
              this.start,
              "Optional chaining cannot appear in the tag of tagged template expressions"
            );
          }
          let node = this.startNodeAt(startPos, startLoc);
          node.tag = base;
          node.quasi = this.parseTemplate({ isTagged: true });
          base = this.finishNode(node, "TaggedTemplateExpression");
        }
        return base;
      }
      parseGetterSetter(prop) {
        prop.kind = prop.key.name;
        this.parsePropertyName(prop);
        prop.value = this.parseMethod(false);
        let paramCount = prop.kind === "get" ? 0 : 1;
        const firstParam = prop.value.params[0];
        const hasContextParam = firstParam && this.isThisParam(firstParam);
        paramCount = hasContextParam ? paramCount + 1 : paramCount;
        if (prop.value.params.length !== paramCount) {
          let start = prop.value.start;
          if (prop.kind === "get") this.raiseRecoverable(start, "getter should have no params");
          else this.raiseRecoverable(start, "setter should have exactly one param");
        } else {
          if (prop.kind === "set" && prop.value.params[0].type === "RestElement")
            this.raiseRecoverable(prop.value.params[0].start, "Setter cannot use rest params");
        }
      }
      parseProperty(isPattern, refDestructuringErrors) {
        if (!isPattern) {
          let decorators = [];
          if (this.match(tokTypes2.at)) {
            while (this.match(tokTypes2.at)) {
              decorators.push(this.parseDecorator());
            }
          }
          const property = super.parseProperty(isPattern, refDestructuringErrors);
          if (property.type === "SpreadElement") {
            if (decorators.length)
              this.raise(property.start, DecoratorsError.SpreadElementDecorator);
          }
          if (decorators.length) {
            property.decorators = decorators;
            decorators = [];
          }
          return property;
        }
        return super.parseProperty(isPattern, refDestructuringErrors);
      }
      parseCatchClauseParam() {
        const param = this.parseBindingAtom();
        let simple = param.type === "Identifier";
        this.enterScope(simple ? acornScope.SCOPE_SIMPLE_CATCH : 0);
        this.checkLValPattern(
          param,
          simple ? acornScope.BIND_SIMPLE_CATCH : acornScope.BIND_LEXICAL
        );
        const type = this.tsTryParseTypeAnnotation();
        if (type) {
          param.typeAnnotation = type;
          this.resetEndLocation(param);
        }
        this.expect(tt.parenR);
        return param;
      }
      parseClass(node, isStatement) {
        const oldInAbstractClass = this.inAbstractClass;
        this.inAbstractClass = !!node.abstract;
        try {
          this.next();
          this.takeDecorators(node);
          const oldStrict = this.strict;
          this.strict = true;
          this.parseClassId(node, isStatement);
          this.parseClassSuper(node);
          const privateNameMap = this.enterClassBody();
          const classBody = this.startNode();
          let hadConstructor = false;
          classBody.body = [];
          let decorators = [];
          this.expect(tt.braceL);
          while (this.type !== tt.braceR) {
            if (this.match(tokTypes2.at)) {
              decorators.push(this.parseDecorator());
              continue;
            }
            const element = this.parseClassElement(node.superClass !== null);
            if (decorators.length) {
              element.decorators = decorators;
              this.resetStartLocationFromNode(element, decorators[0]);
              decorators = [];
            }
            if (element) {
              classBody.body.push(element);
              if (element.type === "MethodDefinition" && element.kind === "constructor" && element.value.type === "FunctionExpression") {
                if (hadConstructor) {
                  this.raiseRecoverable(element.start, "Duplicate constructor in the same class");
                }
                hadConstructor = true;
                if (element.decorators && element.decorators.length > 0) {
                  this.raise(element.start, DecoratorsError.DecoratorConstructor);
                }
              } else if (element.key && element.key.type === "PrivateIdentifier" && isPrivateNameConflicted(privateNameMap, element)) {
                this.raiseRecoverable(
                  element.key.start,
                  `Identifier '#${element.key.name}' has already been declared`
                );
              }
            }
          }
          this.strict = oldStrict;
          this.next();
          if (decorators.length) {
            this.raise(this.start, DecoratorsError.TrailingDecorator);
          }
          node.body = this.finishNode(classBody, "ClassBody");
          this.exitClassBody();
          return this.finishNode(node, isStatement ? "ClassDeclaration" : "ClassExpression");
        } finally {
          this.inAbstractClass = oldInAbstractClass;
        }
      }
      parseClassFunctionParams() {
        const typeParameters = this.tsTryParseTypeParameters();
        let params = this.parseBindingList(tt.parenR, false, this.ecmaVersion >= 8, true);
        if (typeParameters) params.typeParameters = typeParameters;
        return params;
      }
      parseMethod(isGenerator, isAsync, allowDirectSuper, inClassScope, method) {
        let node = this.startNode(), oldYieldPos = this.yieldPos, oldAwaitPos = this.awaitPos, oldAwaitIdentPos = this.awaitIdentPos;
        this.initFunction(node);
        if (this.ecmaVersion >= 6) node.generator = isGenerator;
        if (this.ecmaVersion >= 8) node.async = !!isAsync;
        this.yieldPos = 0;
        this.awaitPos = 0;
        this.awaitIdentPos = 0;
        this.enterScope(
          functionFlags(isAsync, node.generator) | acornScope.SCOPE_SUPER | (allowDirectSuper ? acornScope.SCOPE_DIRECT_SUPER : 0)
        );
        this.expect(tt.parenL);
        node.params = this.parseClassFunctionParams();
        this.checkYieldAwaitInDefaultParams();
        this.parseFunctionBody(node, false, true, false, {
          isClassMethod: inClassScope
        });
        this.yieldPos = oldYieldPos;
        this.awaitPos = oldAwaitPos;
        this.awaitIdentPos = oldAwaitIdentPos;
        if (method && method.abstract) {
          const hasBody = !!node.body;
          if (hasBody) {
            const { key } = method;
            this.raise(
              method.start,
              TypeScriptError.AbstractMethodHasImplementation({
                methodName: key.type === "Identifier" && !method.computed ? key.name : `[${this.input.slice(key.start, key.end)}]`
              })
            );
          }
        }
        return this.finishNode(node, "FunctionExpression");
      }
      static parse(input, options2) {
        if (options2.locations === false) {
          throw new Error(`You have to enable options.locations while using acorn-typescript`);
        } else {
          options2.locations = true;
        }
        const parser = new this(options2, input);
        if (dts) {
          parser.isAmbientContext = true;
        }
        return parser.parse();
      }
      static parseExpressionAt(input, pos, options2) {
        if (options2.locations === false) {
          throw new Error(`You have to enable options.locations while using acorn-typescript`);
        } else {
          options2.locations = true;
        }
        const parser = new this(options2, input, pos);
        if (dts) {
          parser.isAmbientContext = true;
        }
        parser.nextToken();
        return parser.parseExpression();
      }
      parseImportSpecifier() {
        const isMaybeTypeOnly = this.ts_isContextual(tokTypes2.type);
        if (isMaybeTypeOnly) {
          let node = this.startNode();
          node.imported = this.parseModuleExportName();
          this.parseTypeOnlyImportExportSpecifier(
            node,
            /* isImport */
            true,
            this.importOrExportOuterKind === "type"
          );
          return this.finishNode(node, "ImportSpecifier");
        } else {
          const node = super.parseImportSpecifier();
          node.importKind = "value";
          return node;
        }
      }
      parseExportSpecifier(exports) {
        const isMaybeTypeOnly = this.ts_isContextual(tokTypes2.type);
        const isString = this.match(tt.string);
        if (!isString && isMaybeTypeOnly) {
          let node = this.startNode();
          node.local = this.parseModuleExportName();
          this.parseTypeOnlyImportExportSpecifier(
            node,
            /* isImport */
            false,
            this.importOrExportOuterKind === "type"
          );
          this.finishNode(node, "ExportSpecifier");
          this.checkExport(exports, node.exported, node.exported.start);
          return node;
        } else {
          const node = super.parseExportSpecifier(exports);
          node.exportKind = "value";
          return node;
        }
      }
      parseTypeOnlyImportExportSpecifier(node, isImport, isInTypeOnlyImportExport) {
        const leftOfAsKey = isImport ? "imported" : "local";
        const rightOfAsKey = isImport ? "local" : "exported";
        let leftOfAs = node[leftOfAsKey];
        let rightOfAs;
        let hasTypeSpecifier = false;
        let canParseAsKeyword = true;
        const loc = leftOfAs.start;
        if (this.isContextual("as")) {
          const firstAs = this.parseIdent();
          if (this.isContextual("as")) {
            const secondAs = this.parseIdent();
            if (tokenIsKeywordOrIdentifier(this.type)) {
              hasTypeSpecifier = true;
              leftOfAs = firstAs;
              rightOfAs = isImport ? this.parseIdent() : this.parseModuleExportName();
              canParseAsKeyword = false;
            } else {
              rightOfAs = secondAs;
              canParseAsKeyword = false;
            }
          } else if (tokenIsKeywordOrIdentifier(this.type)) {
            canParseAsKeyword = false;
            rightOfAs = isImport ? this.parseIdent() : this.parseModuleExportName();
          } else {
            hasTypeSpecifier = true;
            leftOfAs = firstAs;
          }
        } else if (tokenIsKeywordOrIdentifier(this.type)) {
          hasTypeSpecifier = true;
          if (isImport) {
            leftOfAs = super.parseIdent(true);
            if (!this.isContextual("as")) {
              this.checkUnreserved(leftOfAs);
            }
          } else {
            leftOfAs = this.parseModuleExportName();
          }
        }
        if (hasTypeSpecifier && isInTypeOnlyImportExport) {
          this.raise(
            loc,
            isImport ? TypeScriptError.TypeModifierIsUsedInTypeImports : TypeScriptError.TypeModifierIsUsedInTypeExports
          );
        }
        node[leftOfAsKey] = leftOfAs;
        node[rightOfAsKey] = rightOfAs;
        const kindKey = isImport ? "importKind" : "exportKind";
        node[kindKey] = hasTypeSpecifier ? "type" : "value";
        if (canParseAsKeyword && this.eatContextual("as")) {
          node[rightOfAsKey] = isImport ? this.parseIdent() : this.parseModuleExportName();
        }
        if (!node[rightOfAsKey]) {
          node[rightOfAsKey] = this.copyNode(node[leftOfAsKey]);
        }
        if (isImport) {
          this.checkLValSimple(node[rightOfAsKey], acornScope.BIND_LEXICAL);
        }
      }
      raiseCommonCheck(pos, message, recoverable) {
        switch (message) {
          case "Comma is not permitted after the rest element": {
            if (this.isAmbientContext && this.match(tt.comma) && this.lookaheadCharCode() === 41) {
              this.next();
              return;
            } else {
              return super.raise(pos, message);
            }
          }
        }
        return recoverable ? super.raiseRecoverable(pos, message) : super.raise(pos, message);
      }
      raiseRecoverable(pos, message) {
        return this.raiseCommonCheck(pos, message, true);
      }
      raise(pos, message) {
        return this.raiseCommonCheck(pos, message, true);
      }
      updateContext(prevType) {
        const { type } = this;
        if (type == tt.braceL) {
          var curContext = this.curContext();
          if (curContext == tsTokContexts.tc_oTag) this.context.push(tokContexts.b_expr);
          else if (curContext == tsTokContexts.tc_expr) this.context.push(tokContexts.b_tmpl);
          else super.updateContext(prevType);
          this.exprAllowed = true;
        } else if (type === tt.slash && prevType === tokTypes2.jsxTagStart) {
          this.context.length -= 2;
          this.context.push(tsTokContexts.tc_cTag);
          this.exprAllowed = false;
        } else {
          return super.updateContext(prevType);
        }
      }
      // Parses JSX opening tag starting after '<'.
      jsx_parseOpeningElementAt(startPos, startLoc) {
        let node = this.startNodeAt(startPos, startLoc);
        let nodeName = this.jsx_parseElementName();
        if (nodeName) node.name = nodeName;
        if (this.match(tt.relational) || this.match(tt.bitShift)) {
          const typeArguments = this.tsTryParseAndCatch(
            () => this.tsParseTypeArgumentsInExpression()
          );
          if (typeArguments) node.typeArguments = typeArguments;
        }
        node.attributes = [];
        while (this.type !== tt.slash && this.type !== tokTypes2.jsxTagEnd)
          node.attributes.push(this.jsx_parseAttribute());
        node.selfClosing = this.eat(tt.slash);
        this.expect(tokTypes2.jsxTagEnd);
        return this.finishNode(node, nodeName ? "JSXOpeningElement" : "JSXOpeningFragment");
      }
      enterScope(flags) {
        if (flags === TS_SCOPE_TS_MODULE) {
          this.importsStack.push([]);
        }
        super.enterScope(flags);
        const scope = super.currentScope();
        scope.types = [];
        scope.enums = [];
        scope.constEnums = [];
        scope.classes = [];
        scope.exportOnlyBindings = [];
      }
      exitScope() {
        const scope = super.currentScope();
        if (scope.flags === TS_SCOPE_TS_MODULE) {
          this.importsStack.pop();
        }
        super.exitScope();
      }
      hasImport(name, allowShadow) {
        const len = this.importsStack.length;
        if (this.importsStack[len - 1].indexOf(name) > -1) {
          return true;
        }
        if (!allowShadow && len > 1) {
          for (let i = 0; i < len - 1; i++) {
            if (this.importsStack[i].indexOf(name) > -1) return true;
          }
        }
        return false;
      }
      maybeExportDefined(scope, name) {
        if (this.inModule && scope.flags & acornScope.SCOPE_TOP) {
          this.undefinedExports.delete(name);
        }
      }
      declareName(name, bindingType, pos) {
        if (bindingType & acornScope.BIND_FLAGS_TS_IMPORT) {
          if (this.hasImport(name, true)) {
            this.raise(pos, `Identifier '${name}' has already been declared.`);
          }
          this.importsStack[this.importsStack.length - 1].push(name);
          return;
        }
        const scope = this.currentScope();
        if (bindingType & acornScope.BIND_FLAGS_TS_EXPORT_ONLY) {
          this.maybeExportDefined(scope, name);
          scope.exportOnlyBindings.push(name);
          return;
        }
        if (bindingType === acornScope.BIND_TS_TYPE || bindingType === acornScope.BIND_TS_INTERFACE) {
          if (bindingType === acornScope.BIND_TS_TYPE && scope.types.includes(name)) {
            this.raise(pos, `type '${name}' has already been declared.`);
          }
          scope.types.push(name);
        } else {
          super.declareName(name, bindingType, pos);
        }
        if (bindingType & acornScope.BIND_FLAGS_TS_ENUM) scope.enums.push(name);
        if (bindingType & acornScope.BIND_FLAGS_TS_CONST_ENUM) scope.constEnums.push(name);
        if (bindingType & acornScope.BIND_FLAGS_CLASS) scope.classes.push(name);
      }
      checkLocalExport(id) {
        const { name } = id;
        if (this.hasImport(name)) return;
        const len = this.scopeStack.length;
        for (let i = len - 1; i >= 0; i--) {
          const scope = this.scopeStack[i];
          if (scope.types.indexOf(name) > -1 || scope.exportOnlyBindings.indexOf(name) > -1) return;
        }
        super.checkLocalExport(id);
      }
    }
    return TypeScriptParser;
  };
}
export {
  tsPlugin
};
