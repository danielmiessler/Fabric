# locate-character

Get the line and column number of a particular character in a string.

## Installation

`npm install locate-character`, or get it from [unpkg.com/locate-character](https://unpkg.com/locate-character).

## Usage

To search for a particular character, using the index or a search string, use `locate`:

```js
import { locate } from 'locate-character';

const sample = `
A flea and a fly in a flue
Were imprisoned, so what could they do?
Said the fly, "let us flee!"
"Let us fly!" said the flea.
So they flew through a flaw in the flue.
`.trim();

// Using a character index
const index = sample.indexOf('fly');
locate(sample, index);
// -> { line: 0, column: 13, character: 13 }

// Using the string itself
locate(sample, 'fly');
// -> { line: 0, column: 13, character: 13 }

// Using the string with a start index
locate(sample, 'fly', { startIndex: 14 });
// -> { line: 2, column: 9, character: 76 }
```

If you will be searching the same string repeatedly, it's much faster if you use `getLocator`:

```js
import { getLocator } from 'locate-character';

const locate = getLocator(sample);

let location = locate(13);
// -> { line: 0, column: 13, character: 13 }

location = locate('fly', { startIndex: location.character + 1 });
// -> { line: 2, column: 9, character: 76 }

location = locate('fly', { startIndex: location.character + 1 });
// -> { line: 3, column: 8, character: 104 }
```

In some situations (for example, dealing with sourcemaps), you need one-based line numbers:

```js
getLocator(sample, { offsetLine: 1 });
locate(sample, { offsetLine: 1 });
```

There's also an `offsetColumn` option which is less useful in real-world situations.

## License

MIT
