# Check your code with svelte-check

Provides CLI diagnostics checks for:

-   Unused CSS
-   Svelte A11y hints
-   JavaScript/TypeScript compiler errors

Requires Node 16 or later.

### Usage:

#### Local / in your project

Installation:

`npm i svelte-check --save-dev`

Package.json:

```json
{
    // ...
    "scripts": {
        "svelte-check": "svelte-check"
        // ...
    },
    // ...
    "devDependencies": {
        "svelte-check": "..."
        // ...
    }
}
```

Usage:

`npm run svelte-check`

#### Global (not recommended)

Installation:

`npm i svelte-check svelte -g`

Usage:

1. Go to folder where to start checking
2. `svelte-check`

### Args:

| Flag                                                            | Description                                                                                                                                                                                                                                                                                                                                                                                 |
| --------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--workspace <path>`                                            | Path to your workspace. All subdirectories except node_modules and those listed in `--ignore` are checked                                                                                                                                                                                                                                                                                   |
| `--output <human\|human-verbose\|machine\|machine-verbose>`     |
| `--watch`                                                       | Will not exit after one pass but keep watching files for changes and rerun diagnostics                                                                                                                                                                                                                                                                                                      |
| `--preserveWatchOutput`                                         | Do not clear the screen in watch mode                                                                                                                                                                                                                                                                                                                                                       |
| `--tsconfig <path>`                                             | Pass a path to a tsconfig or jsconfig file. The path can be relative to the workspace path or absolute. Doing this means that only files matched by the files/include/exclude pattern of the config file are diagnosed. It also means that errors from TypeScript and JavaScript files are reported. If not given, will do an upwards traversal looking for the next jsconfig/tsconfig.json |
| `--no-tsconfig`                                                 | Use this if you only want to check the Svelte files found in the current directory and below and ignore any JS/TS files (they will not be type-checked)                                                                                                                                                                                                                                     |
| `--ignore <path1,path2>`                                        | Only has an effect when used in conjunction with `--no-tsconfig`. Files/folders to ignore - relative to workspace root, comma-separated, inside quotes. Example: `--ignore "dist,build"`. When used in conjunction with `--tsconfig`, this will only have effect on the files watched, not on the files that are diagnosed, which is then determined by the `tsconfig.json`                 |
| `--fail-on-warnings`                                            | Will also exit with error code when there are warnings                                                                                                                                                                                                                                                                                                                                      |
| `--compiler-warnings <code1:error\|ignore,code2:error\|ignore>` | A list of Svelte compiler warning codes. Each entry defines whether that warning should be ignored or treated as an error. Warnings are comma-separated, between warning code and error level is a colon; all inside quotes. Example: `--compiler-warnings "css-unused-selector:ignore,unused-export-let:error"`                                                                            |
| `--diagnostic-sources <js,svelte,css>`                          | A list of diagnostic sources which should run diagnostics on your code. Possible values are `js` (includes TS), `svelte`, `css`. Comma-separated, inside quotes. By default all are active. Example: `--diagnostic-sources "js,svelte"`                                                                                                                                                     |
| `--threshold <error\|warning>`                                  | Filters the diagnostics to display. `error` will output only errors while `warning` will output warnings and errors.                                                                                                                                                                                                                                                                        |

### FAQ

#### Why is there no option to only check specific files (for example only staged files)?

`svelte-check` needs to know the whole project to do valid checks. Imagine you alter a component property `export let foo` to `export let bar`, but you don't update any of the component usages. They all have errors now but you would not catch them if you only run checks on changed files.

### More docs, preprocessor setup and troubleshooting

[See here](/docs/README.md).

### Machine-Readable Output

Setting the `--output` to `machine` or `machine-verbose` will format output in a way that is easier to read
by machines, e.g. inside CI pipelines, for code quality checks, etc.

Each row corresponds to a new record. Rows are made up of columns that are separated by a
single space character. The first column of every row contains a timestamp in milliseconds
which can be used for monitoring purposes. The second column gives us the "row type", based
on which the number and types of subsequent columns may differ.

The first row is of type `START` and contains the workspace folder (wrapped in quotes).

###### Example:

```
1590680325583 START "/home/user/language-tools/packages/language-server/test/plugins/typescript/testfiles"
```

Any number of `ERROR` or `WARNING` records may follow. Their structure is identical and depends on the output argoument.

If the argument is `machine` it will tell us the filename, the starting line and column numbers, and the error message. The filename is relative to the workspace directory. The filename and the message are both wrapped in quotes.

###### Example:

```
1590680326283 ERROR "codeactions.svelte" 1:16 "Cannot find module 'blubb' or its corresponding type declarations."
1590680326778 WARNING "imported-file.svelte" 0:37 "Component has unused export property 'prop'. If it is for external reference only, please consider using `export const prop`"
```

If the argument is `machine-verbose` it will tell us the filename, the starting line and column numbers, the ending line and column numbers, the error message, the code of diagnostic, the human-friendly description of the code and the human-friendly source of the diagnostic (eg. svelte/typescript). The filename is relative to the workspace directory. Each diagnostic is represented as an [ndjson](https://en.wikipedia.org/wiki/JSON_streaming#Newline-Delimited_JSON) line prefixed by the timestamp of the log.

###### Example:

```
1590680326283 {"type":"ERROR","fn":"codeaction.svelte","start":{"line":1,"character":16},"end":{"line":1,"character":23},"message":"Cannot find module 'blubb' or its corresponding type declarations.","code":2307,"source":"js"}
1590680326778 {"type":"WARNING","filename":"imported-file.svelte","start":{"line":0,"character":37},"end":{"line":0,"character":51},"message":"Component has unused export property 'prop'. If it is for external reference only, please consider using `export
const prop`","code":"unused-export-let","source":"svelte"}
```

The output concludes with a `COMPLETED` message that summarizes total numbers of files, errors and warnings that were encountered during the check.

###### Example:

```
1590680326807 COMPLETED 20 FILES 21 ERRORS 1 WARNINGS 3 FILES_WITH_PROBLEMS
```

If the application experiences a runtime error, this error will appear as a `FAILURE` record.

###### Example:

```
1590680328921 FAILURE "Connection closed"
```

### Credits

-   Vue's [VTI](https://github.com/vuejs/vetur/tree/master/vti) which laid the foundation for `svelte-check`
