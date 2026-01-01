# mri [![CI](https://github.com/lukeed/mri/workflows/CI/badge.svg?branch=master&event=push)](https://github.com/lukeed/mri/actions)

> Quickly scan for CLI flags and arguments

This is a [fast](#benchmarks) and lightweight alternative to [`minimist`](https://github.com/substack/minimist) and [`yargs-parser`](https://github.com/yargs/yargs-parser).

It only exists because I find that I usually don't need most of what `minimist` and `yargs-parser` have to offer. However, `mri` is similar _enough_ that it might function as a "drop-in replacement" for you, too!

See [Comparisons](#comparisons) for more info.

## Install

```sh
$ npm install --save mri
```

## Usage

```sh
$ demo-cli --foo --bar=baz -mtv -- hello world
```

```js
const mri = require('mri');

const argv = process.argv.slice(2);

mri(argv);
//=> { _: ['hello', 'world'], foo:true, bar:'baz', m:true, t:true, v:true }

mri(argv, { boolean:['bar'] });
//=> { _: ['baz', 'hello', 'world'], foo:true, bar:true, m:true, t:true, v:true }

mri(argv, {
  alias: {
    b: 'bar',
    foo: ['f', 'fuz']
  }
});
//=> { _: ['hello', 'world'], foo:true, f:true, fuz:true, b:'baz', bar:'baz', m:true, t:true, v:true }
```

## API

### mri(args, options)
Return: `Object`

#### args
Type: `Array`<br>
Default: `[]`

An array of arguments to parse. For CLI usage, send `process.argv.slice(2)`. See [`process.argv`](https://nodejs.org/docs/latest/api/process.html#process_process_argv) for info.

#### options.alias
Type: `Object`<br>
Default: `{}`

An object of keys whose values are `String`s or `Array<String>` of aliases. These will be added to the parsed output with matching values.

#### options.boolean
Type: `Array|String`<br>
Default: `[]`

A single key (or array of keys) that should be parsed as `Boolean`s.

#### options.default
Type: `Object`<br>
Default: `{}`

An `key:value` object of defaults. If a default is provided for a key, its type (`typeof`) will be used to cast parsed arguments.

```js
mri(['--foo', 'bar']);
//=> { _:[], foo:'bar' }

mri(['--foo', 'bar'], {
  default: { foo:true, baz:'hello', bat:42 }
});
//=> { _:['bar'], foo:true, baz:'hello', bat:42 }
```

> **Note:** Because `--foo` has a default of `true`, its output is cast to a Boolean. This means that `foo=true`, making `'bar'` an extra argument (`_` key).

#### options.string
Type: `Array|String`<br>
Default: `[]`

A single key (or array of keys) that should be parsed as `String`s.

#### options.unknown
Type: `Function`<br>
Default: `undefined`

Callback that is run when a parsed flag has not been defined as a known key or alias. Its only parameter is the unknown flag itself; eg `--foobar` or `-f`.

Once an unknown flag is encountered, parsing will terminate, regardless of your return value.

> **Note:** `mri` _only_ checks for unknown flags if `options.unknown` **and** `options.alias` are populated. Otherwise, everything will be accepted.


## Comparisons

#### minimist

- `mri` is 5x faster (see [benchmarks](#benchmarks))
- Numerical values are cast as `Number`s when possible
  - A key (and its aliases) will always honor `opts.boolean` or `opts.string`
- Short flag groups are treated as `Boolean`s by default:
    ```js
    minimist(['-abc', 'hello']);
    //=> { _:[], a:'', b:'', c:'hello' }

    mri(['-abc', 'hello']);
    //=> { _:[], a:true, b:true, c:'hello' }
    ```
- The `opts.unknown` behaves differently:
  - Unlike `minimist`, `mri` will not continue continue parsing after encountering an unknown flag
- Missing `options`:
  - `opts.stopEarly`
  - `opts['--']`
- Ignores newlines (`\n`) within args (see [test](https://github.com/substack/minimist/blob/master/test/parse.js#L69-L80))
- Ignores slashBreaks within args (see [test](https://github.com/substack/minimist/blob/master/test/parse.js#L147-L157))
- Ignores dot-nested flags (see [test](https://github.com/substack/minimist/blob/master/test/parse.js#L180-L197))

#### yargs-parser

- `mri` is 40x faster (see [benchmarks](#benchmarks))
- Numerical values are cast as `Number`s when possible
  - A key (and its aliases) will always honor `opts.boolean` or `opts.string`
- Missing `options`:
  - `opts.array`
  - `opts.config`
  - `opts.coerce`
  - `opts.count`
  - `opts.envPrefix`
  - `opts.narg`
  - `opts.normalize`
  - `opts.configuration`
  - `opts.number`
  - `opts['--']`
- Missing [`parser.detailed()`](https://github.com/yargs/yargs-parser#requireyargs-parserdetailedargs-opts) method
- No [additional configuration](https://github.com/yargs/yargs-parser#configuration) object
- Added [`options.unknown`](#optionsunknown) feature


## Benchmarks

> Running Node.js v10.13.0

```
Load Times:
  nopt          3.179ms
  yargs-parser  2.137ms
  minimist      0.746ms
  mri           0.517ms

Benchmark:
  minimist      x    328,747 ops/sec ±1.09% (89 runs sampled)
  mri           x  1,622,801 ops/sec ±0.94% (92 runs sampled)
  nopt          x    888,223 ops/sec ±0.22% (92 runs sampled)
  yargs-parser  x     30,538 ops/sec ±0.81% (91 runs sampled)
```

## License

MIT © [Luke Edwards](https://lukeed.com)
