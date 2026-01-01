# sade [![Build Status](https://travis-ci.org/lukeed/sade.svg?branch=master)](https://travis-ci.org/lukeed/sade)

> Smooth (CLI) Operator 🎶

Sade is a small but powerful tool for building command-line interface (CLI) applications for Node.js that are fast, responsive, and helpful!

It enables default commands, git-like subcommands, option flags with aliases, default option values with type-casting, required-vs-optional argument handling, command validation, and automated help text generation!

Your app's UX will be as smooth as butter... just like [Sade's voice](https://www.youtube.com/watch?v=4TYv2PhG89A). 😉


## Install

```
$ npm install --save sade
```


## Usage

***Input:***

```js
#!/usr/bin/env node

const sade = require('sade');

const prog = sade('my-cli');

prog
  .version('1.0.5')
  .option('--global, -g', 'An example global flag')
  .option('-c, --config', 'Provide path to custom config', 'foo.config.js');

prog
  .command('build <src> <dest>')
  .describe('Build the source directory. Expects an `index.js` entry file.')
  .option('-o, --output', 'Change the name of the output file', 'bundle.js')
  .example('build src build --global --config my-conf.js')
  .example('build app public -o main.js')
  .action((src, dest, opts) => {
    console.log(`> building from ${src} to ${dest}`);
    console.log('> these are extra opts', opts);
  });

prog.parse(process.argv);
```

***Output:***

```a
$ my-cli --help

  Usage
    $ my-cli <command> [options]

  Available Commands
    build    Build the source directory.

  For more info, run any command with the `--help` flag
    $ my-cli build --help

  Options
    -v, --version    Displays current version
    -g, --global     An example global flag
    -c, --config     Provide path to custom config  (default foo.config.js)
    -h, --help       Displays this message


$ my-cli build --help

  Description
    Build the source directory.
    Expects an `index.js` entry file.

  Usage
    $ my-cli build <src> [options]

  Options
    -o, --output    Change the name of the output file  (default bundle.js)
    -g, --global    An example global flag
    -c, --config    Provide path to custom config  (default foo.config.js)
    -h, --help      Displays this message

  Examples
    $ my-cli build src build --global --config my-conf.js
    $ my-cli build app public -o main.js
```

## Tips

- **Define your global/program-wide version, options, description, and/or examples first.**<br>
  _Once you define a Command, you can't access the global-scope again._

- **Define all commands & options in the order that you want them to appear.**<br>
  _Sade will not mutate or sort your CLI for you. Global options print before local options._

- **Required arguments without values will error & exit**<br>
  _An `Insufficient arguments!` error will be displayed along with a help prompt._

- **Don't worry about manually displaying help~!**<br>
  _Your help text is displayed automatically... including command-specific help text!_

- **Automatic default/basic patterns**<br>
  _Usage text will always append `[options]` & `--help` and `--version` are done for you._

- **Only define what you want to display!**<br>
  _Help text sections (example, options, etc) will only display if you provide values._


## Subcommands

Subcommands are defined & parsed like any other command! When defining their [`usage`](#usage-1), everything up until the first argument (`[foo]` or `<foo>`) is interpreted as the command string.

They should be defined in the order that you want them to appear in your general `--help` output.

Lastly, it is _not_ necessary to define the subcommand's "base" as an additional command. However, if you choose to do so, it's recommended that you define it first for better visibility.

```js
const prog = sade('git');

// Not necessary for subcommands to work, but it's here anyway!
prog
  .command('remote')
  .describe('Manage set of tracked repositories')
  .action(opts => {
    console.log('~> Print current remotes...');
  });

prog
  .command('remote add <name> <url>', 'Demo...')
  .action((name, url, opts) => {
    console.log(`~> Adding a new remote (${name}) to ${url}`);
  });

prog
  .command('remote rename <old> <new>', 'Demo...')
  .action((old, nxt, opts) => {
    console.log(`~> Renaming from ${old} to ${nxt}~!`);
  });
```


## Single Command Mode

In certain circumstances, you may only need `sade` for a single-command CLI application.

> **Note:** Until `v1.6.0`, this made for an awkward pairing.

To enable this, you may make use of the [`isSingle`](#issingle) argument. Doing so allows you to pass the program's entire [`usage` text](#usage-1) into the `name` argument.

With "Single Command Mode" enabled, your entire binary operates as one command. This means that any [`prog.command`](#progcommandusage-desc-opts) calls are disallowed & will instead throw an Error. Of course, you may still define a program version, a description, an example or two, and declare options. You are customizing the program's attributes as a whole.<sup>*</sup>

> <sup>*</sup> This is true for multi-command applications, too, up until your first `prog.command()` call!

***Example***

Let's reconstruct [`sirv-cli`](https://github.com/lukeed/sirv), which is a single-command application that (optionally) accepts a directory from which to serve files. It also offers a slew of option flags:

```js
sade('sirv [dir]', true)
  .version('1.0.0')
  .describe('Run a static file server')
  .example('public -qeim 31536000')
  .example('--port 8080 --etag')
  .example('my-app --dev')
  .option('-D, --dev', 'Enable "dev" mode')
  .option('-e, --etag', 'Enable "Etag" header')
  // There are a lot...
  .option('-H, --host', 'Hostname to bind', 'localhost')
  .option('-p, --port', 'Port to bind', 5000)
  .action((dir, opts) => {
    // Program handler
  })
  .parse(process.argv);
```

When `sirv --help` is run, the generated help text is trimmed, fully aware that there's only one command in this program:

```
  Description
    Run a static file server

  Usage
    $ sirv [dir] [options]

  Options
    -D, --dev        Enable "dev" mode
    -e, --etag       Enable "Etag" header
    -H, --host       Hostname to bind  (default localhost)
    -p, --port       Port to bind  (default 5000)
    -v, --version    Displays current version
    -h, --help       Displays this message

  Examples
    $ sirv public -qeim 31536000
    $ sirv --port 8080 --etag
    $ sirv my-app --dev
```

## Command Aliases

Command aliases are alternative names (aliases) for a command. They are often used as shortcuts or as typo relief!

The aliased names do not appear in the general help text.<br>
Instead, they only appear within the Command-specific help text under an "Aliases" section.

***Limitations***

* You cannot assign aliases while in [Single Command Mode](#single-command-mode)
* You cannot call [`prog.alias()`](#progaliasnames) before defining any Commands (via `prog.commmand()`)
* You, the developer, must keep track of which aliases have already been used and/or exist as Command names

***Example***

Let's reconstruct the `npm install` command as a Sade program:

```js
sade('npm')
  // ...
  .command('install [package]', 'Install a package', {
    alias: ['i', 'add', 'isntall']
  })
  .option('-P, --save-prod', 'Package will appear in your dependencies.')
  .option('-D, --save-dev', 'Package will appear in your devDependencies.')
  .option('-O, --save-optional', 'Package will appear in your optionalDependencies')
  .option('-E, --save-exact', 'Save exact versions instead of using a semver range operator')
  // ...
```

When we run `npm --help` we'll see this general help text:

```
  Usage
    $ npm <command> [options]

  Available Commands
    install    Install a package

  For more info, run any command with the `--help` flag
    $ npm install --help

  Options
    -v, --version    Displays current version
    -h, --help       Displays this message
```

When we run `npm install --help` &mdash; ***or*** the help flag with any of `install`'s aliases &mdash; we'll see this command-specific help text:

```
  Description
    Install a package

  Usage
    $ npm install [package] [options]

  Aliases
    $ npm i
    $ npm add
    $ npm isntall

  Options
    -P, --save-prod        Package will appear in your dependencies.
    -D, --save-dev         Package will appear in your devDependencies.
    -O, --save-optional    Package will appear in your optionalDependencies
    -E, --save-exact       Save exact versions instead of using a semver range operator
    -h, --help             Displays this message
```



## API

### sade(name, isSingle)
Returns: `Program`

Returns your chainable Sade instance, aka your `Program`.

#### name
Type: `String`<br>
Required: `true`

The name of your `Program` / binary application.

#### isSingle
Type: `Boolean`<br>
Default: `name.includes(' ');`

If your `Program` is meant to have ***only one command***.<br>
When `true`, this simplifies your generated `--help` output such that:

* the "root-level help" is your _only_ help text
* the "root-level help" does not display an `Available Commands` section
* the "root-level help" does not inject `$ name <command>` into the `Usage` section
* the "root-level help" does not display `For more info, run any command with the `--help` flag` text

You may customize the `Usage` of your command by modifying the `name` argument directly.<br>
Please read [Single Command Mode](#single-command-mode) for an example and more information.

> **Important:** Whenever `name` includes a custom usage, then `isSingle` is automatically assumed and enforced!

### prog.command(usage, desc, opts)

Create a new Command for your Program. This changes the current state of your Program.

All configuration methods (`prog.describe`, `prog.action`, etc) will apply to this Command until another Command has been created!

#### usage

Type: `String`

The usage pattern for your current Command. This will be included in the general or command-specific `--help` output.

_Required_ arguments are wrapped with `<` and `>` characters; for example, `<foo>` and `<bar>`.

_Optional_ arguments are wrapped with `[` and `]` characters; for example, `[foo]` and `[bar]`.

All arguments are ***positionally important***, which means they are passed to your current Command's [`handler`](#handler) function in the order that they were defined.

When optional arguments are defined but don't receive a value, their positionally-equivalent function parameter will be `undefined`.

> **Important:** You **must** define & expect required arguments _before_ optional arguments!

```js
sade('foo')

  .command('greet <adjective> <noun>')
  .action((adjective, noun, opts) => {
    console.log(`Hello, ${adjective} ${noun}!`);
  })

  .command('drive <vehicle> [color] [speed]')
  .action((vehicle, color, speed, opts) => {
    let arr = ['Driving my'];
    arr.push(color ? `${color} ${vehicle}` : vehicle);
    speed && arr.push(`at ${speed}`);
    opts.yolo && arr.push('...YOLO!!');
    let str = arr.join(' ');
    console.log(str);
  });
```

```sh
$ foo greet beautiful person
# //=> Hello, beautiful person!

$ foo drive car
# //=> Driving my car

$ foo drive car red
# //=> Driving my red card

$ foo drive car blue 100mph --yolo
# //=> Driving my blue car at 100mph ...YOLO!!
```


#### desc

Type: `String`<br>
Default: `''`

The Command's description. The value is passed directly to [`prog.describe`](#progdescribetext).

#### opts

Type: `Object`<br>
Default: `{}`

##### opts.alias
Type: `String|Array`

Optionally define one or more aliases for the current Command.<br>
When declared, the `opts.alias` value is passed _directly_ to the [`prog.alias`](#progaliasnames) method.

```js
// Program A is equivalent to Program B
// ---

const A = sade('bin')
  .command('build', 'My build command', { alias: 'b' })
  .command('watch', 'My watch command', { alias: ['w', 'dev'] });

const B = sade('bin')
  .command('build', 'My build command').alias('b')
  .command('watch', 'My watch command').alias('w', 'dev');
```


##### opts.default

Type: `Boolean`

Manually set/force the current Command to be the Program's default command. This ensures that the current Command will run if no command was specified.

> **Important:** If you run your Program without a Command _and_ without specifying a default command, your Program will exit with a `No command specified` error.

```js
const prog = sade('greet');

prog.command('hello');
//=> only runs if :: `$ greet hello`

// $ greet
//=> error: No command specified.

prog.command('howdy', '', { default:true });
//=> runs as `$ greet` OR `$ greet howdy`

// $ greet
//=> runs 'howdy' handler

// $ greet foobar
//=> error: Invalid command
```


### prog.describe(text)

Add a description to the current Command.

#### text

Type: `String|Array`

The description text for the current Command. This will be included in the general or command-specific `--help` output.

Internally, your description will be separated into an `Array` of sentences.

For general `--help` output, ***only*** the first sentence will be displayed. However, **all sentences** will be printed for command-specific `--help` text.

> **Note:** Pass an `Array` if you don't want internal assumptions. However, the first item is _always_ displayed in general help, so it's recommended to keep it short.


### prog.alias(...names)

Define one or more aliases for the current Command.

> **Important:** An error will be thrown if:<br>1) the program is in [Single Command Mode](#single-command-mode); or<br>2) `prog.alias` is called before any `prog.command`.

#### names

Type: `String`

The list of alternative names (aliases) for the current Command.<br>
For example, you may want to define shortcuts and/or common typos for the Command's full name.

> **Important:** Sade _does not_ check if the incoming `names` are already in use by other Commands or their aliases.<br>During conflicts, the Command with the same `name` is given priority, otherwise the first Command (according to Program order) with `name` as an alias is chosen.

The `prog.alias()` is append-only, so calling it multiple times within a Command context will _keep_ all aliases, including those initially passed via [`opts.alias`](#optsdefault).

```js
sade('bin')
  .command('hello <name>', 'Greet someone by their name', {
    alias: ['hey', 'yo']
  })
  .alias('hi', 'howdy')
  .alias('hola', 'oi');
//=> hello aliases: hey, yo, hi, howdy, hola, oi
```


### prog.action(handler)

Attach a callback to the current Command.

#### handler

Type: `Function`

The function to run when the current Command is executed.

Its parameters are based (positionally) on your Command's [`usage`](#usage-1) definition.

All options, flags, and extra/unknown values are included as the last parameter.

> **Note:** Optional arguments are also passed as parameters & may be `undefined`!

```js
sade('foo')
  .command('cp <src> <dest>')
  .option('-f, --force', 'Overwrite without confirmation')
  .option('-c, --clone-dir', 'Copy files to additional directory')
  .option('-v, --verbose', 'Enable verbose output')
  .action((src, dest, opts) => {
    console.log(`Copying files from ${src} --> ${dest}`);
    opts.c && console.log(`ALSO copying files from ${src} --> ${opts['clone-dir']}`);
    console.log('My options:', opts);
  })

// $ foo cp original my-copy -v
//=> Copying files from original --> my-copy
//=> My options: { _:[], v:true, verbose:true }

// $ foo cp original my-copy --clone-dir my-backup
//=> Copying files from original --> my-copy
//=> ALSO copying files from original --> my-backup
//=> My options: { _:[], c:'my-backup', 'clone-dir':'my-backup' }
```


### prog.example(str)

Add an example for the current Command.

#### str

Type: `String`

The example string to add. This will be included in the general or command-specific `--help` output.

> **Note:** Your example's `str` will be prefixed with your Program's [`name`](#sadename).


### prog.option(flags, desc, value)

Add an Option to the current Command.

#### flags

Type: `String`

The Option's flags, which may optionally include an alias.

You may use a comma (`,`) or a space (` `) to separate the flags.

> **Note:** The short & long flags can be declared in any order. However, the alias will always be displayed first.

> **Important:** If using hyphenated flag names, they will be accessible **as declared** within your [`action()`](#progactionhandler) handler!

```js
prog.option('--global'); // no alias
prog.option('-g, --global'); // alias first, comma
prog.option('--global -g'); // alias last, space
// etc...
```

#### desc

Type: `String`

The description for the Option.

#### value

Type: `String`

The **default** value for the Option.

Flags and aliases, if parsed, are `true` by default. See [`mri`](https://github.com/lukeed/mri#minimist) for more info.

> **Note:** You probably only want to define a default `value` if you're expecting a `String` or `Number` value type.

If you _do_ pass a `String` or `Number` value type, your flag value will be casted to the same type. See [`mri#options.default`](https://github.com/lukeed/mri#optionsdefault) for info~!


### prog.version(str)

The `--version` and `-v` flags will automatically output the Program version.

#### str

Type: `String`<br>
Default: `0.0.0`

The new version number for your Program.

> **Note:** Your Program `version` is `0.0.0` until you change it.

### prog.parse(arr, opts)

Parse a set of CLI arguments.

#### arr

Type: `Array`

Your Program's `process.argv` input.

> **Important:** Do not `.slice(2)`! Doing so will break parsing~!

#### opts

Type: `Object`<br>
Default: `{}`

Additional `process.argv` parsing config. See [`mri`'s options](https://github.com/lukeed/mri#mriargs-options) for details.

> **Important:** These values _override_ any internal values!

```js
prog
  .command('hello')
  .option('-f, --force', 'My flag');
//=> currently has alias pair: f <--> force

prog.parse(process.argv, {
  alias: {
    f: ['foo', 'fizz']
  },
  default: {
    abc: 123
  }
});
//=> ADDS alias pair: f <--> foo
//=> REMOVES alias pair: f <--> force
//=> ADDS alias pair: f <--> fizz
//=> ADDS default: abc -> 123 (number)
```

#### opts.unknown

Type: `Function`<br>
Default: `undefined`

Callback to run when an unspecified option flag has been found. This is [passed directly to `mri`](https://github.com/lukeed/mri#optionsunknown).

Your handler will receive the unknown flag (string) as its only argument.<br>
You may return a string, which will be used as a custom error message. Otherwise, a default message is displayed.

```js
sade('sirv')
  .command('start [dir]')
  .parse(process.argv, {
    unknown: arg => `Custom error message: ${arg}`
  });

/*
$ sirv start --foobar

  ERROR
    Custom error message: --foobar

  Run `$ sirv --help` for more info.
*/
```

#### opts.lazy

Type: `Boolean`<br>
Default: `false`

If true, Sade will not immediately execute the `action` handler. Instead, `parse()` will return an object of `{ name, args, handler }` shape, wherein the `name` is the command name, `args` is all arguments that _would be_ passed to the action handler, and `handler` is the function itself.

From this, you may choose when to run the `handler` function. You also have the option to further modify the `args` for any reason, if needed.

```js
let { name, args, handler } = prog.parse(process.argv, { lazy:true });
console.log('> Received command: ', name);

// later on...
handler.apply(null, args);
```

### prog.help(cmd)

Manually display the help text for a given command. If no command name is provided, the general/global help is printed.

Your general and command-specific help text is automatically attached to the `--help` and `-h` flags.

> **Note:** You don't have to call this directly! It's automatically run when you `bin --help`

#### cmd
Type: `String`<br>
Default: `null`

The name of the command for which to display help. Otherwise displays the general help.


## License

MIT © [Luke Edwards](https://lukeed.com)
