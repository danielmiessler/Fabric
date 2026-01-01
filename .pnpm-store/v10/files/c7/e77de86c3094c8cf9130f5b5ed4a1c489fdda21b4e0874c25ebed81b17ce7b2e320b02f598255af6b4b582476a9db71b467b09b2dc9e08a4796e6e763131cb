# find-yarn-workspace-root

![CI](https://github.com/square/find-yarn-workspace-root/workflows/CI/badge.svg)

Algorithm for finding the root of a yarn workspace, extracted from yarnpkg.com

## Installation

```bash
yarn add find-yarn-workspace-root
```

## Usage

```js
const findWorkspaceRoot = require('find-yarn-workspace-root');

const workspaceRoot = findWorkspaceRoot(__dirname); // Absolute path or null
```

## Contributing

Contributions are welcome! Just clone this repository and install the dependencies:

```bash
git clone https://github.com/square/find-yarn-workspace-root.git
cd find-yarn-workspace-root
yarn
```

Note that you'll need `node` and `yarn` installed. Next, verify the tests all pass:

```bash
yarn test
```

Then create a branch for your bugfix/feature, make changes and update the tests, and submit a pull request. Please do not change the version in `package.json` when submitting a pull request. We determine the next version automatically based on the commits since the last release.

> *NOTE*: Commit messages follow the [Angular commit message guidelines](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commits).

## Releases

To trigger a release, bump the version using `standard-version`:

1. To update the CHANGELOG, bump the version in `package.json`, and create a git tag, run:

   ```bash
   $(yarn bin)/standard-version
   ```

2. `git push && git push --tags`

3. `npm publish`

---
Copyright 2017 Square, Inc.
