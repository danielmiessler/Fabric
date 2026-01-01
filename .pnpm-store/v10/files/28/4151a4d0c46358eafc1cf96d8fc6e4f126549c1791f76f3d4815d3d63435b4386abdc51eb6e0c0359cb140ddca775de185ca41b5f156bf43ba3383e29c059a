# eslint-compat-utils

***This package is still in the experimental stage.***

Provides an API for ESLint custom rules that is compatible with the latest ESLint even when using older ESLint.

## Installation

```bash
npm install eslint-compat-utils
```

## Usage

```js
const { getSourceCode } = require("eslint-compat-utils");
module.exports = {
  meta: { /* ... */ },
  create(context) {
    const sourceCode = getSourceCode(context)
    return {
      "Program"(node) {
        const scope = sourceCode.getScope(node);
      },
    };
  },
}
```

### API

#### `getSourceCode(context)`

Returns an extended instance of `context.sourceCode` or the result of `context.getSourceCode()`. Extended instances can use new APIs such as `getScope(node)` even with old ESLint.

#### `getCwd(context)`

Gets the value of `context.cwd`, but for older ESLint it returns the result of `context.getCwd()`.
Versions older than v6.6.0 return a value from the result of `process.cwd()`.

#### `getFilename(context)`

Gets the value of `context.filename`, but for older ESLint it returns the result of `context.getFilename()`.

#### `getPhysicalFilename(context)`

Gets the value of `context.physicalFilename`, but for older ESLint it returns the result of `context.getPhysicalFilename()`.
Versions older than v7.28.0 return a value guessed from the result of `context.getFilename()`, but it may be incorrect.
