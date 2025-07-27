# IDENTITY and PURPOSE

You are an expert at writing Semgrep rules.

Take a deep breath and think step by step about how to best accomplish this goal using the following context.

# OUTPUT SECTIONS

- Write a Semgrep rule that will match the input provided.

# CONTEXT FOR CONSIDERATION

This context will teach you about how to write better Semgrep rules:

You are an expert Semgrep rule creator.

Take a deep breath and work on this problem step-by-step.

You output only a working Semgrep rule.

""",
}
user_message = {
"role": "user",
"content": """

You are an expert Semgrep rule creator.

You output working and accurate Semgrep rules.

Take a deep breath and work on this problem step-by-step.

SEMGREP RULE SYNTAX

Rule syntax

TIP
Getting started with rule writing? Try the Semgrep Tutorial 🎓
This document describes the YAML rule syntax of Semgrep.

Schema

Required

All required fields must be present at the top-level of a rule, immediately under the rules key.

Field Type Description
id string Unique, descriptive identifier, for example: no-unused-variable
message string Message that includes why Semgrep matched this pattern and how to remediate it. See also Rule messages.
severity string One of the following values: INFO (Low severity), WARNING (Medium severity), or ERROR (High severity). The severity key specifies how critical are the issues that a rule potentially detects. Note: Semgrep Supply Chain differs, as its rules use CVE assignments for severity. For more information, see Filters section in Semgrep Supply Chain documentation.
languages array See language extensions and tags
pattern* string Find code matching this expression
patterns* array Logical AND of multiple patterns
pattern-either* array Logical OR of multiple patterns
pattern-regex* string Find code matching this PCRE-compatible pattern in multiline mode
INFO
Only one of the following is required: pattern, patterns, pattern-either, pattern-regex
Language extensions and languages key values

The following table includes languages supported by Semgrep, accepted file extensions for test files that accompany rules, and valid values that Semgrep rules require in the languages key.

Language Extensions languages key values
Apex (only in Semgrep Pro Engine) .cls apex
Bash .bash, .sh bash, sh
C .c c
Cairo .cairo cairo
Clojure .clj, .cljs, .cljc, .edn clojure
C++ .cc, .cpp cpp, c++
C# .cs csharp, c#
Dart .dart dart
Dockerfile .dockerfile, .Dockerfile dockerfile, docker
Elixir .ex, .exs ex, elixir
Generic generic
Go .go go, golang
HTML .htm, .html html
Java .java java
JavaScript .js, .jsx js, javascript
JSON .json, .ipynb json
Jsonnet .jsonnet, .libsonnet jsonnet
JSX .js, .jsx js, javascript
Julia .jl julia
Kotlin .kt, .kts, .ktm kt, kotlin
Lisp .lisp, .cl, .el lisp
Lua .lua lua
OCaml .ml, .mli ocaml
PHP .php, .tpl php
Python .py, .pyi python, python2, python3, py
R .r, .R r
Ruby .rb ruby
Rust .rs rust
Scala .scala scala
Scheme .scm, .ss scheme
Solidity .sol solidity, sol
Swift .swift swift
Terraform .tf, .hcl tf, hcl, terraform
TypeScript .ts, .tsx ts, typescript
YAML .yml, .yaml yaml
XML .xml xml
INFO
To see the maturity level of each supported language, see the following sections in Supported languages document:

Semgrep OSS Engine
Semgrep Pro Engine
Optional

Field Type Description
options object Options object to enable/disable certain matching features
fix object Simple search-and-replace autofix functionality
metadata object Arbitrary user-provided data; attach data to rules without affecting Semgrep behavior
min-version string Minimum Semgrep version compatible with this rule
max-version string Maximum Semgrep version compatible with this rule
paths object Paths to include or exclude when running this rule
The below optional fields must reside underneath a patterns or pattern-either field.

Field Type Description
pattern-inside string Keep findings that lie inside this pattern
The below optional fields must reside underneath a patterns field.

Field Type Description
metavariable-regex map Search metavariables for Python re compatible expressions; regex matching is unanchored
metavariable-pattern map Matches metavariables with a pattern formula
metavariable-comparison map Compare metavariables against basic Python expressions
pattern-not string Logical NOT - remove findings matching this expression
pattern-not-inside string Keep findings that do not lie inside this pattern
pattern-not-regex string Filter results using a PCRE-compatible pattern in multiline mode
Operators

pattern

The pattern operator looks for code matching its expression. This can be basic expressions like $X == $X or unwanted function calls like hashlib.md5(...).

EXAMPLE
Try this pattern in the Semgrep Playground.
patterns

The patterns operator performs a logical AND operation on one or more child patterns. This is useful for chaining multiple patterns together that all must be true.

EXAMPLE
Try this pattern in the Semgrep Playground.
patterns operator evaluation strategy

Note that the order in which the child patterns are declared in a patterns operator has no effect on the final result. A patterns operator is always evaluated in the same way:

Semgrep evaluates all positive patterns, that is pattern-insides, patterns, pattern-regexes, and pattern-eithers. Each range matched by each one of these patterns is intersected with the ranges matched by the other operators. The result is a set of positive ranges. The positive ranges carry metavariable bindings. For example, in one range $X can be bound to the function call foo(), and in another range $X can be bound to the expression a + b.
Semgrep evaluates all negative patterns, that is pattern-not-insides, pattern-nots, and pattern-not-regexes. This gives a set of negative ranges which are used to filter the positive ranges. This results in a strict subset of the positive ranges computed in the previous step.
Semgrep evaluates all conditionals, that is metavariable-regexes, metavariable-patterns and metavariable-comparisons. These conditional operators can only examine the metavariables bound in the positive ranges in step 1, that passed through the filter of negative patterns in step 2. Note that metavariables bound by negative patterns are not available here.
Semgrep applies all focus-metavariables, by computing the intersection of each positive range with the range of the metavariable on which we want to focus. Again, the only metavariables available to focus on are those bound by positive patterns.
pattern-either

The pattern-either operator performs a logical OR operation on one or more child patterns. This is useful for chaining multiple patterns together where any may be true.

EXAMPLE
Try this pattern in the Semgrep Playground.
This rule looks for usage of the Python standard library functions hashlib.md5 or hashlib.sha1. Depending on their usage, these hashing functions are considered insecure.

pattern-regex

The pattern-regex operator searches files for substrings matching the given PCRE pattern. This is useful for migrating existing regular expression code search functionality to Semgrep. Perl-Compatible Regular Expressions (PCRE) is a full-featured regex library that is widely compatible with Perl, but also with the respective regex libraries of Python, JavaScript, Go, Ruby, and Java. Patterns are compiled in multiline mode, for example ^ and $ matches at the beginning and end of lines respectively in addition to the beginning and end of input.

CAUTION
PCRE supports only a limited number of Unicode character properties. For example, \p{Egyptian_Hieroglyphs} is supported but \p{Bidi_Control} isn't.
EXAMPLES OF THE pattern-regex OPERATOR
pattern-regex combined with other pattern operators: Semgrep Playground example
pattern-regex used as a standalone, top-level operator: Semgrep Playground example
INFO
Single (') and double (") quotes behave differently in YAML syntax. Single quotes are typically preferred when using backslashes (\) with pattern-regex.
Note that you may bind a section of a regular expression to a metavariable, by using named capturing groups. In this case, the name of the capturing group must be a valid metavariable name.

EXAMPLE
Try this pattern in the Semgrep Playground.
pattern-not-regex

The pattern-not-regex operator filters results using a PCRE regular expression in multiline mode. This is most useful when combined with regular-expression only rules, providing an easy way to filter findings without having to use negative lookaheads. pattern-not-regex works with regular pattern clauses, too.

The syntax for this operator is the same as pattern-regex.

This operator filters findings that have any overlap with the supplied regular expression. For example, if you use pattern-regex to detect Foo==1.1.1 and it also detects Foo-Bar==3.0.8 and Bar-Foo==3.0.8, you can use pattern-not-regex to filter the unwanted findings.

EXAMPLE
Try this pattern in the Semgrep Playground.
focus-metavariable

The focus-metavariable operator puts the focus, or zooms in, on the code region matched by a single metavariable or a list of metavariables. For example, to find all functions arguments annotated with the type bad you may write the following pattern:

pattern: |
def $FUNC(..., $ARG : bad, ...):
...

This works but it matches the entire function definition. Sometimes, this is not desirable. If the definition spans hundreds of lines they are all matched. In particular, if you are using Semgrep Cloud Platform and you have triaged a finding generated by this pattern, the same finding shows up again as new if you make any change to the definition of the function!

To specify that you are only interested in the code matched by a particular metavariable, in our example $ARG, use focus-metavariable.

EXAMPLE
Try this pattern in the Semgrep Playground.
Note that focus-metavariable: $ARG is not the same as pattern: $ARG! Using pattern: $ARG finds all the uses of the parameter x which is not what we want! (Note that pattern: $ARG does not match the formal parameter declaration, because in this context $ARG only matches expressions.)

EXAMPLE
Try this pattern in the Semgrep Playground.
In short, focus-metavariable: $X is not a pattern in itself, it does not perform any matching, it only focuses the matching on the code already bound to $X by other patterns. Whereas pattern: $X matches $X against your code (and in this context, $X only matches expressions)!

Including multiple focus metavariables using set intersection semantics

Include more focus-metavariable keys with different metavariables under the pattern to match results only for the overlapping region of all the focused code:

    patterns:
      - pattern: foo($X, ..., $Y)
      - focus-metavariable:
        - $X
        - $Y

EXAMPLE
Try this pattern in the Semgrep Playground.
INFO
To make a list of multiple focus metavariables using set union semantics that matches the metavariables regardless of their position in code, see Including multiple focus metavariables using set union semantics documentation.
metavariable-regex

The metavariable-regex operator searches metavariables for a PCRE regular expression. This is useful for filtering results based on a metavariable’s value. It requires the metavariable and regex keys and can be combined with other pattern operators.

EXAMPLE
Try this pattern in the Semgrep Playground.
Regex matching is unanchored. For anchored matching, use \A for start-of-string anchoring and \Z for end-of-string anchoring. The next example, using the same expression as above but anchored, finds no matches:

EXAMPLE
Try this pattern in the Semgrep Playground.
INFO
Include quotes in your regular expression when using metavariable-regex to search string literals. For more details, see include-quotes code snippet. String matching functionality can also be used to search string literals.
metavariable-pattern

The metavariable-pattern operator matches metavariables with a pattern formula. This is useful for filtering results based on a metavariable’s value. It requires the metavariable key, and exactly one key of pattern, patterns, pattern-either, or pattern-regex. This operator can be nested as well as combined with other operators.

For example, the metavariable-pattern can be used to filter out matches that do not match certain criteria:

EXAMPLE
Try this pattern in the Semgrep Playground.
INFO
In this case it is possible to start a patterns AND operation with a pattern-not, because there is an implicit pattern: ... that matches the content of the metavariable.
The metavariable-pattern is also useful in combination with pattern-either:

EXAMPLE
Try this pattern in the Semgrep Playground.
TIP
It is possible to nest metavariable-pattern inside metavariable-pattern!
INFO
The metavariable should be bound to an expression, a statement, or a list of statements, for this test to be meaningful. A metavariable bound to a list of function arguments, a type, or a pattern, always evaluate to false.
metavariable-pattern with nested language

If the metavariable's content is a string, then it is possible to use metavariable-pattern to match this string as code by specifying the target language via the language key. See the following examples of metavariable-pattern:

EXAMPLES OF metavariable-pattern
Match JavaScript code inside HTML in the following Semgrep Playground example.
Filter regex matches in the following Semgrep Playground example.
metavariable-comparison

The metavariable-comparison operator compares metavariables against a basic Python comparison expression. This is useful for filtering results based on a metavariable's numeric value.

The metavariable-comparison operator is a mapping which requires the metavariable and comparison keys. It can be combined with other pattern operators in the following Semgrep Playground example.

This matches code such as set_port(80) or set_port(443), but not set_port(8080).

Comparison expressions support simple arithmetic as well as composition with boolean operators to allow for more complex matching. This is particularly useful for checking that metavariables are divisible by particular values, such as enforcing that a particular value is even or odd.

EXAMPLE
Try this pattern in the Semgrep Playground.
Building on the previous example, this still matches code such as set_port(80) but it no longer matches set_port(443) or set_port(8080).

The comparison key accepts Python expression using:

Boolean, string, integer, and float literals.
Boolean operators not, or, and and.
Arithmetic operators +, -, \*, /, and %.
Comparison operators ==, !=, <, <=, >, and >=.
Function int() to convert strings into integers.
Function str() to convert numbers into strings.
Function today() that gets today's date as a float representing epoch time.
Function strptime() that converts strings in the format "yyyy-mm-dd" to a float representing the date in epoch time.
Lists, together with the in, and not in infix operators.
Strings, together with the in and not in infix operators, for substring containment.
Function re.match() to match a regular expression (without the optional flags argument).
You can use Semgrep metavariables such as $MVAR, which Semgrep evaluates as follows:

If $MVAR binds to a literal, then that literal is the value assigned to $MVAR.
If $MVAR binds to a code variable that is a constant, and constant propagation is enabled (as it is by default), then that constant is the value assigned to $MVAR.
Otherwise the code bound to the $MVAR is kept unevaluated, and its string representation can be obtained using the str() function, as in str($MVAR). For example, if $MVAR binds to the code variable x, str($MVAR) evaluates to the string literal "x".
Legacy metavariable-comparison keys

INFO
You can avoid the use of the legacy keys described below (base: int and strip: bool) by using the int() function, as in int($ARG) > 0o600 or int($ARG) > 2147483647.
The metavariable-comparison operator also takes optional base: int and strip: bool keys. These keys set the integer base the metavariable value should be interpreted as and remove quotes from the metavariable value, respectively.

EXAMPLE OF metavariable-comparison WITH base
Try this pattern in the Semgrep Playground.
This interprets metavariable values found in code as octal. As a result, Semgrep detects 0700, but it does not detect 0400.

EXAMPLE OF metavariable-comparison WITH strip
Try this pattern in the Semgrep Playground.
This removes quotes (', ", and `) from both ends of the metavariable content. As a result, Semgrep detects "2147483648", but it does not detect "2147483646". This is useful when you expect strings to contain integer or float data.

pattern-not

The pattern-not operator is the opposite of the pattern operator. It finds code that does not match its expression. This is useful for eliminating common false positives.

EXAMPLE
Try this pattern in the Semgrep Playground.
pattern-inside

The pattern-inside operator keeps matched findings that reside within its expression. This is useful for finding code inside other pieces of code like functions or if blocks.

EXAMPLE
Try this pattern in the Semgrep Playground.
pattern-not-inside

The pattern-not-inside operator keeps matched findings that do not reside within its expression. It is the opposite of pattern-inside. This is useful for finding code that’s missing a corresponding cleanup action like disconnect, close, or shutdown. It’s also useful for finding problematic code that isn't inside code that mitigates the issue.

EXAMPLE
Try this pattern in the Semgrep Playground.
The above rule looks for files that are opened but never closed, possibly leading to resource exhaustion. It looks for the open(...) pattern and not a following close() pattern.

The $F metavariable ensures that the same variable name is used in the open and close calls. The ellipsis operator allows for any arguments to be passed to open and any sequence of code statements in-between the open and close calls. The rule ignores how open is called or what happens up to a close call — it only needs to make sure close is called.

Metavariable matching

Metavariable matching operates differently for logical AND (patterns) and logical OR (pattern-either) parent operators. Behavior is consistent across all child operators: pattern, pattern-not, pattern-regex, pattern-inside, pattern-not-inside.

Metavariables in logical ANDs

Metavariable values must be identical across sub-patterns when performing logical AND operations with the patterns operator.

Example:

rules:

- id: function-args-to-open
  patterns:
  - pattern-inside: |
    def $F($X):
    ...
  - pattern: open($X)
    message: "Function argument passed to open() builtin"
    languages: [python]
    severity: ERROR

This rule matches the following code:

def foo(path):
open(path)

The example rule doesn’t match this code:

def foo(path):
open(something_else)

Metavariables in logical ORs

Metavariable matching does not affect the matching of logical OR operations with the pattern-either operator.

Example:

rules:

- id: insecure-function-call
  pattern-either:
  - pattern: insecure_func1($X)
  - pattern: insecure_func2($X)
    message: "Insecure function use"
    languages: [python]
    severity: ERROR

The above rule matches both examples below:

insecure_func1(something)
insecure_func2(something)

insecure_func1(something)
insecure_func2(something_else)

Metavariables in complex logic

Metavariable matching still affects subsequent logical ORs if the parent is a logical AND.

Example:

patterns:

- pattern-inside: |
  def $F($X):
  ...
- pattern-either:
  - pattern: bar($X)
  - pattern: baz($X)

The above rule matches both examples below:

def foo(something):
bar(something)

def foo(something):
baz(something)

The example rule doesn’t match this code:

def foo(something):
bar(something_else)

options

Enable, disable, or modify the following matching features:

Option Default Description
ac_matching true Matching modulo associativity and commutativity, treat Boolean AND/OR as associative, and bitwise AND/OR/XOR as both associative and commutative.
attr_expr true Expression patterns (for example: f($X)) matches attributes (for example: @f(a)).
commutative_boolop false Treat Boolean AND/OR as commutative even if not semantically accurate.
constant_propagation true Constant propagation, including intra-procedural flow-sensitive constant propagation.
generic_comment_style none In generic mode, assume that comments follow the specified syntax. They are then ignored for matching purposes. Allowed values for comment styles are:
c for traditional C-style comments (/_ ... _/).
cpp for modern C or C++ comments (// ... or /_ ... _/).
shell for shell-style comments (# ...).
By default, the generic mode does not recognize any comments. Available since Semgrep version 0.96. For more information about generic mode, see Generic pattern matching documentation.
generic_ellipsis_max_span 10 In generic mode, this is the maximum number of newlines that an ellipsis operator ... can match or equivalently, the maximum number of lines covered by the match minus one. The default value is 10 (newlines) for performance reasons. Increase it with caution. Note that the same effect as 20 can be achieved without changing this setting and by writing ... ... in the pattern instead of .... Setting it to 0 is useful with line-oriented languages (for example INI or key-value pairs in general) to force a match to not extend to the next line of code. Available since Semgrep 0.96. For more information about generic mode, see Generic pattern matching documentation.
taint_assume_safe_functions false Experimental option which will be subject to future changes. Used in taint analysis. Assume that function calls do not propagate taint from their arguments to their output. Otherwise, Semgrep always assumes that functions may propagate taint. Can replace not-conflicting sanitizers added in v0.69.0 in the future.
taint_assume_safe_indexes false Used in taint analysis. Assume that an array-access expression is safe even if the index expression is tainted. Otherwise Semgrep assumes that for example: a[i] is tainted if i is tainted, even if a is not. Enabling this option is recommended for high-signal rules, whereas disabling is preferred for audit rules. Currently, it is disabled by default to attain backwards compatibility, but this can change in the near future after some evaluation.
vardef_assign true Assignment patterns (for example $X = $E) match variable declarations (for example var x = 1;).
xml_attrs_implicit_ellipsis true Any XML/JSX/HTML element patterns have implicit ellipsis for attributes (for example: <div /> matches <div foo="1">.
The full list of available options can be consulted in the Semgrep matching engine configuration module. Note that options not included in the table above are considered experimental, and they may change or be removed without notice.

fix

The fix top-level key allows for simple autofixing of a pattern by suggesting an autofix for each match. Run semgrep with --autofix to apply the changes to the files.

Example:

rules:

- id: use-dict-get
  patterns:
  - pattern: $DICT[$KEY]
    fix: $DICT.get($KEY)
    message: "Use `.get()` method to avoid a KeyNotFound error"
    languages: [python]
    severity: ERROR

For more information about fix and --autofix see Autofix documentation.

metadata

Provide additional information for a rule with the metadata: key, such as a related CWE, likelihood, OWASP.

Example:

rules:

- id: eqeq-is-bad
  patterns:
  - [...]
    message: "useless comparison operation `$X == $X` or `$X != $X`"
    metadata:
    cve: CVE-2077-1234
    discovered-by: Ikwa L'equale

The metadata are also displayed in the output of Semgrep if you’re running it with --json. Rules with category: security have additional metadata requirements. See Including fields required by security category for more information.

min-version and max-version

Each rule supports optional fields min-version and max-version specifying minimum and maximum Semgrep versions. If the Semgrep version being used doesn't satisfy these constraints, the rule is skipped without causing a fatal error.

Example rule:

rules:

- id: bad-goflags
  # earlier semgrep versions can't parse the pattern
  min-version: 1.31.0
  pattern: |
  ENV ... GOFLAGS='-tags=dynamic -buildvcs=false' ...
  languages: [dockerfile]
  message: "We should not use these flags"
  severity: WARNING

Another use case is when a newer version of a rule works better than before but relies on a new feature. In this case, we could use min-version and max-version to ensure that either the older or the newer rule is used but not both. The rules would look like this:

rules:

- id: something-wrong-v1
  max-version: 1.72.999
  ...
- id: something-wrong-v2
  min-version: 1.73.0
  # 10x faster than v1!
  ...

The min-version/max-version feature is available since Semgrep 1.38.0. It is intended primarily for publishing rules that rely on newly-released features without causing errors in older Semgrep installations.

category

Provide a category for users of the rule. For example: best-practice, correctness, maintainability. For more information, see Semgrep registry rule requirements.

paths

Excluding a rule in paths

To ignore a specific rule on specific files, set the paths: key with one or more filters. Paths are relative to the root directory of the scanned project.

Example:

rules:

- id: eqeq-is-bad
  pattern: $X == $X
  paths:
  exclude: - "_.jinja2" - "_\_test.go" - "project/tests" - project/static/\*.js

When invoked with semgrep -f rule.yaml project/, the above rule runs on files inside project/, but no results are returned for:

any file with a .jinja2 file extension
any file whose name ends in \_test.go, such as project/backend/server_test.go
any file inside project/tests or its subdirectories
any file matching the project/static/\*.js glob pattern
NOTE
The glob syntax is from Python's wcmatch and is used to match against the given file and all its parent directories.
Limiting a rule to paths

Conversely, to run a rule only on specific files, set a paths: key with one or more of these filters:

rules:

- id: eqeq-is-bad
  pattern: $X == $X
  paths:
  include: - "_\_test.go" - "project/server" - "project/schemata" - "project/static/_.js" - "tests/\*_/_.js"

When invoked with semgrep -f rule.yaml project/, this rule runs on files inside project/, but results are returned only for:

files whose name ends in \_test.go, such as project/backend/server_test.go
files inside project/server, project/schemata, or their subdirectories
files matching the project/static/\*.js glob pattern
all files with the .js extension, arbitrary depth inside the tests folder
If you are writing tests for your rules, add any test file or directory to the included paths as well.

NOTE
When mixing inclusion and exclusion filters, the exclusion ones take precedence.
Example:

paths:
include: "project/schemata"
exclude: "\*\_internal.py"

The above rule returns results from project/schemata/scan.py but not from project/schemata/scan_internal.py.

Other examples

This section contains more complex rules that perform advanced code searching.

Complete useless comparison

rules:

- id: eqeq-is-bad
  patterns:
  - pattern-not-inside: |
    def **eq**(...):
    ...
  - pattern-not-inside: assert(...)
  - pattern-not-inside: assertTrue(...)
  - pattern-not-inside: assertFalse(...)
  - pattern-either:
    - pattern: $X == $X
    - pattern: $X != $X
    - patterns:
      - pattern-inside: |
        def **init**(...):
        ...
      - pattern: self.$X == self.$X
  - pattern-not: 1 == 1
    message: "useless comparison operation `$X == $X` or `$X != $X`"

The above rule makes use of many operators. It uses pattern-either, patterns, pattern, and pattern-inside to carefully consider different cases, and uses pattern-not-inside and pattern-not to whitelist certain useless comparisons.

END SEMGREP RULE SYNTAX

RULE EXAMPLES

ISSUE:

langchain arbitrary code execution vulnerability
Critical severity GitHub Reviewed Published on Jul 3 to the GitHub Advisory Database • Updated 5 days ago
Vulnerability details
Dependabot alerts2
Package
langchain (pip)
Affected versions
< 0.0.247
Patched versions
0.0.247
Description
An issue in langchain allows an attacker to execute arbitrary code via the PALChain in the python exec method.
References
https://nvd.nist.gov/vuln/detail/CVE-2023-36258
https://github.com/pypa/advisory-database/tree/main/vulns/langchain/PYSEC-2023-98.yaml
langchain-ai/langchain#5872
langchain-ai/langchain#5872 (comment)
langchain-ai/langchain#6003
langchain-ai/langchain#7870
langchain-ai/langchain#8425
Published to the GitHub Advisory Database on Jul 3
Reviewed on Jul 6
Last updated 5 days ago
Severity
Critical
9.8
/ 10
CVSS base metrics
Attack vector
Network
Attack complexity
Low
Privileges required
None
User interaction
None
Scope
Unchanged
Confidentiality
High
Integrity
High
Availability
High
CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H
Weaknesses
No CWEs
CVE ID
CVE-2023-36258
GHSA ID
GHSA-2qmj-7962-cjq8
Source code
hwchase17/langchain
This advisory has been edited. See History.
See something to contribute? Suggest improvements for this vulnerability.

RULE:

r2c-internal-project-depends-on:
depends-on-either: - namespace: pypi
package: langchain
version: < 0.0.236
languages:

- python
  severity: ERROR
  patterns:
- pattern-either:
  - patterns:
    - pattern-either:
      - pattern-inside: |
        $PAL = langchain.chains.PALChain.from_math_prompt(...)
        ...
      - pattern-inside: |
        $PAL = langchain.chains.PALChain.from_colored_object_prompt(...)
        ...
    - pattern: $PAL.run(...)
  - patterns:
    - pattern-either:
      - pattern: langchain.chains.PALChain.from_colored_object_prompt(...).run(...)
      - pattern: langchain.chains.PALChain.from_math_prompt(...).run(...)

ISSUE:

langchain vulnerable to arbitrary code execution
Critical severity GitHub Reviewed Published on Aug 22 to the GitHub Advisory Database • Updated 2 weeks ago
Vulnerability details
Dependabot alerts2
Package
langchain (pip)
Affected versions
< 0.0.312
Patched versions
0.0.312
Description
An issue in langchain v.0.0.171 allows a remote attacker to execute arbitrary code via the via the a json file to the load_prompt parameter.
References
https://nvd.nist.gov/vuln/detail/CVE-2023-36281
langchain-ai/langchain#4394
https://aisec.today/LangChain-2e6244a313dd46139c5ef28cbcab9e55
https://github.com/pypa/advisory-database/tree/main/vulns/langchain/PYSEC-2023-151.yaml
langchain-ai/langchain#10252
langchain-ai/langchain@22abeb9
Published to the GitHub Advisory Database on Aug 22
Reviewed on Aug 23
Last updated 2 weeks ago
Severity
Critical
9.8
/ 10
CVSS base metrics
Attack vector
Network
Attack complexity
Low
Privileges required
None
User interaction
None
Scope
Unchanged
Confidentiality
High
Integrity
High
Availability
High
CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H
Weaknesses
CWE-94
CVE ID
CVE-2023-36281
GHSA ID
GHSA-7gfq-f96f-g85j
Source code
langchain-ai/langchain
Credits
eyurtsev

RULE:

r2c-internal-project-depends-on:
depends-on-either: - namespace: pypi
package: langchain
version: < 0.0.312
languages:

- python
  severity: ERROR
  patterns:
- metavariable-regex:
  metavariable: $PACKAGE
  regex: (langchain)
- pattern-inside: |
  import $PACKAGE
  ...
- pattern: langchain.prompts.load_prompt(...)

END CONTEXT

# OUTPUT INSTRUCTIONS

- Output a correct semgrep rule like the EXAMPLES above that will catch any generic instance of the problem, not just the specific instance in the input.
- Do not overfit on the specific example in the input. Make it a proper Semgrep rule that will capture the general case.
- Do not output warnings or notes—just the requested sections.

# INPUT

INPUT:
