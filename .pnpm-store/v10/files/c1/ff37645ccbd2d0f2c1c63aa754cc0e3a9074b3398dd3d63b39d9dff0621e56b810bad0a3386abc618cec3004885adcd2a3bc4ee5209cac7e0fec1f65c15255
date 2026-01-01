<p align="center">
  <img src="https://ds300.github.io/patch-package/patch-package.svg" width="80%" alt="patch-package" />
</p>

`patch-package` lets app authors instantly make and keep fixes to npm
dependencies. It's a vital band-aid for those of us living on the bleeding edge.

```sh
# fix a bug in one of your dependencies
vim node_modules/some-package/brokenFile.js

# run patch-package to create a .patch file
npx patch-package some-package

# commit the patch file to share the fix with your team
git add patches/some-package+3.14.15.patch
git commit -m "fix brokenFile.js in some-package"
```

Patches created by `patch-package` are automatically and gracefully applied when
you use `npm`(>=5) or `yarn`.

No more waiting around for pull requests to be merged and published. No more
forking repos just to fix that one tiny thing preventing your app from working.

## Set-up

In package.json

```diff
 "scripts": {
+  "postinstall": "patch-package"
 }
```

Then

### npm

    npm i patch-package

You can use `--save-dev` if you don't need to run npm in production, e.g. if
you're making a web frontend.

### yarn v1

    yarn add patch-package postinstall-postinstall

You can use `--dev` if you don't need to run yarn in production, e.g. if you're
making a web frontend.

To understand why yarn needs the `postinstall-postinstall` package see:
[Why use postinstall-postinstall](#why-use-postinstall-postinstall-with-yarn)

### yarn workspaces

Same as for yarn ☝️ Note that if you want to patch un-hoisted packages you'll
need to repeat the setup process for the child package. Also make sure you're in
the child package directory when you run `patch-package` to generate the patch
files.

### yarn v2+

yarn 2+ have native support for patching dependencies via
[`yarn patch`](https://yarnpkg.com/cli/patch). You do not need to use
patch-package on these projects.

### pnpm

pnpm has native support for patching dependencies via
[`pnpm patch`](https://pnpm.io/cli/patch). You do not need to use patch-package
on these projects.

### Heroku

For `patch-package` to work on Heroku applications, you must specify
[`NPM_CONFIG_PRODUCTION=false` or `YARN_PRODUCTION=false`](https://devcenter.heroku.com/articles/nodejs-support#package-installation).
See [this issue](https://github.com/ds300/patch-package/issues/130) for more
details.

### Docker and CI

- If having errors about working directory ("cannot run in wd [...]") when
  building in Docker, you might need to adjust configuration in `.npmrc`. See
  [#185](https://github.com/ds300/patch-package/issues/185).
- In your `Dockerfile`, remember to copy over the patch files _before_ running
  `[npm|yarn] install`
- If you cache `node_modules` rather than running `yarn install` every time,
  make sure that the `patches` dir is included in your cache key somehow.
  Otherwise if you update a patch then the change may not be reflected on
  subsequent CI runs.

### CircleCI

Create a hash of your patches before loading/saving your cache. If using a Linux
machine, run `md5sum patches/* > patches.hash`. If running on a macOS machine,
use `md5 patches/* > patches.hash`

```yaml
- run:
    name: patch-package hash
    command: md5sum patches/* > patches.hash
```

Then, update your hash key to include a checksum of that file:

```yaml
- restore_cache:
    key:
      app-node_modules-v1-{{ checksum "yarn.lock" }}-{{ checksum "patches.hash"
      }}
```

As well as the save_cache

```yaml
- save_cache:
    key:
      app-node_modules-v1-{{ checksum "yarn.lock" }}-{{ checksum "patches.hash"
      }}
    paths:
      - ./node_modules
```

## Usage

### Making patches

First make changes to the files of a particular package in your node_modules
folder, then run

    yarn patch-package package-name

or use npx (included with `npm > 5.2`)

    npx patch-package package-name

where `package-name` matches the name of the package you made changes to.

If this is the first time you've used `patch-package`, it will create a folder
called `patches` in the root dir of your app. Inside will be a file called
`package-name+0.44.0.patch` or something, which is a diff between normal old
`package-name` and your fixed version. Commit this to share the fix with your
team.

#### Options

- `--create-issue`

  For packages whose source is hosted on GitHub this option opens a web browser
  with a draft issue based on your diff.

- `--use-yarn`

  By default, patch-package checks whether you use npm or yarn based on which
  lockfile you have. If you have both, it uses npm by default. Set this option
  to override that default and always use yarn.

- `--exclude <regexp>`

  Ignore paths matching the regexp when creating patch files. Paths are relative
  to the root dir of the package to be patched.

  Default value: `package\\.json$`

- `--include <regexp>`

  Only consider paths matching the regexp when creating patch files. Paths are
  relative to the root dir of the package to be patched.

  Default value: `.*`

- `--case-sensitive-path-filtering`

  Make regexps used in --include or --exclude filters case-sensitive.

- `--patch-dir`

  Specify the name for the directory in which to put the patch files.

#### Nested packages

If you are trying to patch a package at, e.g.
`node_modules/package/node_modules/another-package` you can just put a `/`
between the package names:

    npx patch-package package/another-package

It works with scoped packages too

    npx patch-package @my/package/@my/other-package

### Updating patches

Use exactly the same process as for making patches in the first place, i.e. make
more changes, run patch-package, commit the changes to the patch file.

### Applying patches

Run `patch-package` without arguments to apply all patches in your project.

#### Options

- `--error-on-fail`

  Forces patch-package to exit with code 1 after failing.

  When running locally patch-package always exits with 0 by default. This
  happens even after failing to apply patches because otherwise yarn.lock and
  package.json might get out of sync with node_modules, which can be very
  confusing.

  `--error-on-fail` is **switched on** by default on CI.

  See https://github.com/ds300/patch-package/issues/86 for background.

- `--reverse`

  Un-applies all patches.

  Note that this will fail if the patched files have changed since being
  patched. In that case, you'll probably need to re-install `node_modules`.

  This option was added to help people using CircleCI avoid
  [an issue around caching and patch file updates](https://github.com/ds300/patch-package/issues/37)
  but might be useful in other contexts too.

- `--patch-dir`

  Specify the name for the directory in which the patch files are located

#### Notes

To apply patches individually, you may use `git`:

    git apply --ignore-whitespace patches/package-name+0.44.2.patch

or `patch` in unixy environments:

    patch -p1 -i patches/package-name+0.44.2.patch

### Dev-only patches

If you deploy your package to production (e.g. your package is a server) then
any patched `devDependencies` will not be present when patch-package runs in
production. It will happily ignore those patch files if the package to be
patched is listed directly in the `devDependencies` of your package.json. If
it's a transitive dependency patch-package can't detect that it is safe to
ignore and will throw an error. To fix this, mark patches for transitive dev
dependencies as dev-only by renaming from, e.g.

    package-name+0.44.0.patch

to

    package-name+0.44.0.dev.patch

This will allow those patch files to be safely ignored when
`NODE_ENV=production`.

### Creating multiple patches for the same package

_💡 This is an advanced feature and is not recommended unless you really, really
need it._

Let's say you have a patch for react-native called

- `patches/react-native+0.72.0.patch`

If you want to add another patch file to `react-native`, you can use the
`--append` flag while supplying a name for the patch.

Just make you changes inside `node_modules/react-native` then run e.g.

    npx patch-package react-native --append 'fix-touchable-opacity'

This will create a new patch file while renaming the old patch file so that you
now have:

- `patches/react-native+0.72.0+001+initial.patch`
- `patches/react-native+0.72.0+002+fix-touchable-opacity.patch`

The patches are ordered in a sequence, so that they can build on each other if
necessary. **Think of these as commits in a git history**.

#### Updating a sequenced patch file

If the patch file is the last one in the sequence, you can just make your
changes inside e.g. `node_modules/react-native` and then run

    npx patch-package react-native

This will update the last patch file in the sequence.

If the patch file is not the last one in the sequence **you need to use the
`--rebase` feature** to un-apply the succeeding patch files first.

Using the example above, let's say you want to update the `001+initial` patch
but leave the other patch alone. You can run

    npx patch-package react-native --rebase patches/react-native+0.72.0+001+initial.patch

This will undo the `002-fix-touchable-opacity` patch file. You can then make
your changes and run

    npx patch-package react-native

to finish the rebase by updating the `001+initial` patch file and re-apply the
`002-fix-touchable-opacity` patch file, leaving you with all patches applied and
up-to-date.

#### Inserting a new patch file in the middle of an existing sequence

Using the above example, let's say you want to insert a new patch file between
the `001+initial` and `002+fix-touchable-opacity` patch files. You can run

    npx patch-package react-native --rebase patches/react-native+0.72.0+001+initial.patch

This will undo the `002-fix-touchable-opacity` patch file. You can then make any
changes you want to insert in a new patch file and run

    npx patch-package react-native --append 'fix-console-warnings'

This will create a new patch file while renaming any successive patches to
maintain the sequence order, leaving you with

- `patches/react-native+0.72.0+001+initial.patch`
- `patches/react-native+0.72.0+002+fix-console-warnings.patch`
- `patches/react-native+0.72.0+003+fix-touchable-opacity.patch`

To insert a new patch file at the start of the sequence, you can run

    npx patch-package react-native --rebase 0

Which will un-apply all patch files in the sequence. Then follow the process
above to create a new patch file numbered `001`.

#### Deleting a sequenced patch file

To delete a sequenced patch file, just delete it, then remove and reinstall your
`node_modules` folder.

If you deleted one of the patch files other than the last one, you don't need to
update the sequence numbers in the successive patch file names, but you might
want to do so to keep things tidy.

#### Partially applying a broken patch file

Normally patch application is atomic per patch file. i.e. if a patch file
contains an error anywhere then none of the changes in the patch file will be
applied and saved to disk.

This can be problematic if you have a patch with many changes and you want to
keep some of them and update others.

In this case you can use the `--partial` option. Patch-package will apply as
many of the changes as it can and then leave it to you to fix the rest.

Any errors encountered will be written to a file `./patch-package-errors.log` to
help you keep track of what needs fixing.

## Benefits of patching over forking

- Sometimes forks need extra build steps, e.g. with react-native for Android.
  Forget that noise.
- Get told in big red letters when the dependency changed and you need to check
  that your fix is still valid.
- Keep your patches colocated with the code that depends on them.
- Patches can be reviewed as part of your normal review process, forks probably
  can't

## When to fork instead

- The change is too consequential to be developed in situ.
- The change would be useful to other people as-is.
- You can afford to make a proper PR to upstream.

## Isn't this dangerous?

Nope. The technique is quite robust. Here are some things to keep in mind
though:

- It's easy to forget to run `yarn` or `npm` when switching between branches
  that do and don't have patch files.
- Long lived patches can be costly to maintain if they affect an area of code
  that is updated regularly and you want to update the package regularly too.
- Big semantic changes can be hard to review. Keep them small and obvious or add
  plenty of comments.
- Changes can also impact the behaviour of other untouched packages. It's
  normally obvious when this will happen, and often desired, but be careful
  nonetheless.

## Why use postinstall-postinstall with Yarn?

Most times when you do a `yarn`, `yarn add`, `yarn remove`, or `yarn install`
(which is the same as just `yarn`) Yarn will completely replace the contents of
your node_modules with freshly unpackaged modules. patch-package uses the
`postinstall` hook to modify these fresh modules, so that they behave well
according to your will.

Yarn only runs the `postinstall` hook after `yarn` and `yarn add`, but not after
`yarn remove`. The `postinstall-postinstall` package is used to make sure your
`postinstall` hook gets executed even after a `yarn remove`.

## License

MIT

[![Empowered by Futurice's open source sponsorship program](https://img.shields.io/badge/sponsor-chilicorn-ff69b4.svg)](http://futurice.com/blog/sponsoring-free-time-open-source-activities?utm_source=github&utm_medium=spice&utm_campaign=patch-package)
