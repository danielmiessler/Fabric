# yarn-lockfile

Parse and/or write `yarn.lock` files

## Usage Examples

```js
const fs = require('fs');
const lockfile = require('@yarnpkg/lockfile');
// or (es6)
import fs from 'fs';
import * as lockfile from '@yarnpkg/lockfile';

let file = fs.readFileSync('yarn.lock', 'utf8');
let json = lockfile.parse(file);

console.log(json);

let fileAgain = lockfile.stringify(json);

console.log(fileAgain);
```
