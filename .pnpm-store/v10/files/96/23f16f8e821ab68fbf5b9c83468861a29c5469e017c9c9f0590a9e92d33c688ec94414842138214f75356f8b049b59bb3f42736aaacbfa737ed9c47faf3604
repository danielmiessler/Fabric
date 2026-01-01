# string-similarity

Finds degree of similarity between two strings, based on [Dice's Coefficient](http://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient), which is mostly better than [Levenshtein distance](http://en.wikipedia.org/wiki/Levenshtein_distance).

## Table of Contents

- [string-similarity](#string-similarity)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
    - [For Node.js](#for-nodejs)
    - [For browser apps](#for-browser-apps)
  - [API](#api)
    - [compareTwoStrings(string1, string2)](#comparetwostringsstring1-string2)
      - [Arguments](#arguments)
      - [Returns](#returns)
      - [Examples](#examples)
    - [findBestMatch(mainString, targetStrings)](#findbestmatchmainstring-targetstrings)
      - [Arguments](#arguments-1)
      - [Returns](#returns-1)
      - [Examples](#examples-1)
  - [Release Notes](#release-notes)
    - [2.0.0](#200)
    - [3.0.0](#300)
    - [3.0.1](#301)
    - [4.0.1](#401)
    - [4.0.2](#402)
    - [4.0.3](#403)

## Usage

#### For Node.js

Install using:

```shell
npm install string-similarity --save
```

In your code:

```javascript
var stringSimilarity = require("string-similarity");

var similarity = stringSimilarity.compareTwoStrings("healed", "sealed");

var matches = stringSimilarity.findBestMatch("healed", [
  "edward",
  "sealed",
  "theatre",
]);
```

#### For browser apps

Include `<script src="//unpkg.com/string-similarity/umd/string-similarity.min.js"></script>` to get the latest version.

Or `<script src="//unpkg.com/string-similarity@4.0.1/umd/string-similarity.min.js"></script>` to get a specific version (4.0.1) in this case.

This exposes a global variable called `stringSimilarity` which you can start using.

```
<script>
  stringSimilarity.compareTwoStrings('what!', 'who?');
</script>
```

(The package is exposed as UMD, so you can consume it as such)

## API

The package contains two methods:

### compareTwoStrings(string1, string2)

Returns a fraction between 0 and 1, which indicates the degree of similarity between the two strings. 0 indicates completely different strings, 1 indicates identical strings. The comparison is case-sensitive.

##### Arguments

1. string1 (string): The first string
2. string2 (string): The second string

Order does not make a difference.

##### Returns

(number): A fraction from 0 to 1, both inclusive. Higher number indicates more similarity.

##### Examples

```javascript
stringSimilarity.compareTwoStrings("healed", "sealed");
// → 0.8

stringSimilarity.compareTwoStrings(
  "Olive-green table for sale, in extremely good condition.",
  "For sale: table in very good  condition, olive green in colour."
);
// → 0.6060606060606061

stringSimilarity.compareTwoStrings(
  "Olive-green table for sale, in extremely good condition.",
  "For sale: green Subaru Impreza, 210,000 miles"
);
// → 0.2558139534883721

stringSimilarity.compareTwoStrings(
  "Olive-green table for sale, in extremely good condition.",
  "Wanted: mountain bike with at least 21 gears."
);
// → 0.1411764705882353
```

### findBestMatch(mainString, targetStrings)

Compares `mainString` against each string in `targetStrings`.

##### Arguments

1. mainString (string): The string to match each target string against.
2. targetStrings (Array): Each string in this array will be matched against the main string.

##### Returns

(Object): An object with a `ratings` property, which gives a similarity rating for each target string, a `bestMatch` property, which specifies which target string was most similar to the main string, and a `bestMatchIndex` property, which specifies the index of the bestMatch in the targetStrings array.

##### Examples

```javascript
stringSimilarity.findBestMatch('Olive-green table for sale, in extremely good condition.', [
  'For sale: green Subaru Impreza, 210,000 miles',
  'For sale: table in very good condition, olive green in colour.',
  'Wanted: mountain bike with at least 21 gears.'
]);
// →
{ ratings:
   [ { target: 'For sale: green Subaru Impreza, 210,000 miles',
       rating: 0.2558139534883721 },
     { target: 'For sale: table in very good condition, olive green in colour.',
       rating: 0.6060606060606061 },
     { target: 'Wanted: mountain bike with at least 21 gears.',
       rating: 0.1411764705882353 } ],
  bestMatch:
   { target: 'For sale: table in very good condition, olive green in colour.',
     rating: 0.6060606060606061 },
  bestMatchIndex: 1
}
```

## Release Notes

### 2.0.0

- Removed production dependencies
- Updated to ES6 (this breaks backward-compatibility for pre-ES6 apps)

### 3.0.0

- Performance improvement for `compareTwoStrings(..)`: now O(n) instead of O(n^2)
- The algorithm has been tweaked slightly to disregard spaces and word boundaries. This will change the rating values slightly but not enough to make a significant difference
- Adding a `bestMatchIndex` to the results for `findBestMatch(..)` to point to the best match in the supplied `targetStrings` array

### 3.0.1

- Refactoring: removed unused functions; used `substring` instead of `substr`
- Updated dependencies

### 4.0.1

- Distributing as an UMD build to be used in browsers.

### 4.0.2

- Update dependencies to latest versions.

### 4.0.3

- Make compatible with IE and ES5. Also, update deps. (see [PR56](https://github.com/aceakash/string-similarity/pull/56))

### 4.0.4

- Simplify some conditional statements. Also, update deps. (see [PR50](https://github.com/aceakash/string-similarity/pull/50))

![Build status](https://codeship.com/projects/2aa453d0-0959-0134-8a76-4abcb29fe9b4/status?branch=master)
[![Known Vulnerabilities](https://snyk.io/test/github/aceakash/string-similarity/badge.svg)](https://snyk.io/test/github/aceakash/string-similarity)
