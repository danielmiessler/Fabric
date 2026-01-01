# Changelog

## 8.0.1

- Update `tmp` to fix security alert. #578

## 8.0.0

- Add support for multiple patch files for a single package. #474

### Breaking Changes

- Removed support for some very old legacy filename format stuff (pre-2019).

## 7.0.2

- Bump semver again (contribution from @rsanchez in #477)

## 7.0.1

- Bump semver (contribution from @stianjensen in #466)

### Breaking changes

- Bump `yaml` to fix security issue. Required bumping minimum node version from
  8 to 14. (contribution from @mayank99 in #463)

### Other changes

- Bump `cross-spawn` (contribution from @stianjensen in #457)
- Replace `is-ci` with `ci-info` (contribution from @paescuj in #446)
- Make version number parsing more robust (contribution from @MHekert in #361)

## 6.5.1

Special thanks to @orta again for putting this release together.

- Improve error message for missing lockfile entry (contribution from @Andarist
  and @farskid in #439)
- Fix creating patches for github dependencies (contribution from @dijonkitchen
  in #444)
- Improve yarn 2 support by copying .yarn dir when creating patch (contribution
  from @bdwain in #438)
- Upgrade `fs-extra` (contribution from @stianjensen in #441)
- Add support for nmp lockfile v2 (contribution from @anas10 in #434)

## 6.5.0

Special thanks to @orta for stepping in to get this release out.

- Don't show issue creation tips for DefinitelyTyped (contribution from @orta in
  #331)
- Add --error-on-warn flag (contribution from @mvargeson in #345)
- Add excplit git diff path prefixes (contribution from @tye-singwa in #395)
- Add yarn 2 compatibility (contribution from @maman in #363)
- Update license (contribution from @Serjobas in #422)
- Fix patch file parse but (contribution from @bill2004158 in #301)
- Add unit tests for getPackageVersion (contribution from @MHekert in #359)
- Deref symlink on copying .npmrc and .yarnrc in makePatch (contribution from
  @jimmyltsinn in 402)
- Fix package.json exclusion (contribution from @milahu in #390)
- Upgrade minimist (contribution from @tapanprakasht in #412)
- Upgrade chalk (contribution from @stianjensen in #389)
- Support .yarnrc (contribution from @NMinhNguyen in #222)

## 6.4.7

- Use npm i --force when creating patches

## 6.4.6

- Update find-yarn-workspace-root (contribution from @sarimarton in #282)
- Fix rare npm crash when creating patches (contribution from @kf6kjg in #271)

## 6.4.5

- Increase max pipe buffer size when creating patch. (contribution from
  @nomi9995 in #287)

## 6.4.4

- Refine --create-issue template

## 6.4.3

- Ensure real path is used when copying package contents. Potentially adds pnpm
  support (contribution from @milahu in #252)

## 6.4.2

- Add package version to --create-issue template

## 6.4.1

- Add backlink to patch-package repo to collect public usage of --create-issue

## 6.4.0

- Add --create-issue option

## 6.3.1

- Fix another createPatch bug involving .git folder hygiene (contribution from
  @gomain in #258)

## 6.3.0

- Add --error-on-fail cli option
- Collate errors and warnings to avoid exiting early (with help from
  @akwodkiewicz in #217)
- Fix .dev.patch suffix bug (with help from @pdcastro in #224)
- Support build metadata in package.json version strings (with help from
  @snowystinger in #213)
- Change 'no patch files found' message from red to blue (contribution from
  @dmhalejr in #211)
- Fix environment variable propogation for spawned tasks (contribution from
  @chpio in #223)
- Add --patch-dir option to readme (contribution from @mikehardy in #225)
- Fix createPatch bug involving .git folder hygiene (contribution from @haroenv
  in #231)

## 6.2.2

- Take config from .yarnrc when making patches (contribution from @NMinhNguyen
  in #222)

## 6.2.1

- Avoid infinite loop when invoked without package.json
- Fall back to version in package-lock (contribution from @bschlenk in #206)
- Add tips about Docker and CI to README (contribution from @harriha in #208)
- Remove update notifier (contribution from @christianbundy in #196)
- Add github actions checks (contribution from @asadm in #186 and @DanielRuf in
  #188, #187)
- Fix url in README (contribution from @DanielRuf in #184)
- Upgrade node in CI to 12 (contribution from @DanielRuf in #183)

## 6.2.0

- Add support for dev-only patches (#159 again)

## 6.1.4

- Use --ignore-scripts when making patches if it fails without (#151)

## 6.1.3

- Fail when patches are found for uninstalled packages (#159)
- Support private registries declared in .npmrc (Contribution from @cherniavskii
  in #152)

## 6.1.2

- Explicitly handle failure edge case where symlinks are modified/created (#118)

## 6.1.1

- Fix npm edge case of package installed from url (#134)

## 6.1.0

- Add support for yarn workspaces
- Degrade 'file mode change not required' error to warning

## 6.0.7

- Don't try to detect package manager on patch application (#133)

## 6.0.6

- Better error messages for patch parse failures (#131)

## 6.0.5

- Use lockfiles to get package resolutions.

## 6.0.4

- Fix patch parsing issue affecting windows (#124)

## 6.0.3

- Allow relative file paths for --patch-dir (Contribution from @lots0logs in
  #119)
- Fix version string handling (Contribution from @teppeis in #122)
- Add support for custom resolutions field in app's package.json (#125)
- Fix property access bug (#123)
- Move @types/is-ci from dependencies to devDependencies (#121)

## 6.0.2

- Revert failure exit code when no patches are found.

## 6.0.1

- Document --patch-dir option in readme

## 6.0.0

### Highlights!

- No longer dependent on Git to apply patches
- Patch files are created much much much faster 🏃🏽‍♀️💨
- Fixed lots of small bugs
- Explicit support for nested packages

For full details see the prerelease notes.

## 6.0.0-18

- Allow fuzzy patch hunk application.
- Minor UI tweaks
- Ignore global git config to prevent issues like #109 & #115
- Add --ignore-engines to yarn invocation

## 6.0.0-17

- Fix the removal of old patch files when creating new ones.

## 6.0.0-16

- Fail postinstall only on CI to prevent weird upgrade issues locally (see #86)
- Fail if no patches are present

## 6.0.0-15

- Handle mode changes
- Backwards-compatible patch file parsing

## 6.0.0-14

- Handle renaming files properly

## 6.0.0-13

- Handle large diffs by not calling .toString on stdout buffer
- Git usage fixes

## 6.0.0-12

- Support explicit nested package patching
- Improve performance of patch creation

## 6.0.0-11

- Handle crlf line breaks in patch parser (Contribution from @NMinhNguyen)

## 6.0.0-10

- Add --patches-dir option (Contribution from @davidpett)

## 6.0.0-9

- Fix patch application bug when creating new files (Contribution from
  @stmarkidis)

## 6.0.0-8

- Improve diffing speed (Contribution from @KevinVlaanderen)

## 6.0.0-7

- Use --no-ext-diff option when generating diffs (Contribution from @janv)

## 6.0.0-6

- Make include/exclude regexes applied relative to the package root. Fixes #54

## 6.0.0-5

- Fix preventing scripts from running when making patch file

## 6.0.0-4

- Don't delete package.json during patch creation

## 6.0.0-3

- Bugfixes for patch application
- Prevent scripts from running when making patch file

## 6.0.0-2

- Bugfixes for patch application

## 6.0.0-1

- Bugfixes for patch application

## 6.0.0-0

- Reimplement most of patch application in TypeScript

## 5.1.1

- Fix idempotency regression. See #39

## 5.1.0

- Add `--reverse` option for patch application. See #37

## 5.0.0

- Remove yarn patching code
- Recommend postinstall-prepare in README for yarn compatibility

## 4.0.0

- Ignore all package.json files by default
- Exit with appropriate error when git is not available

## 3.6.1

- Fix bug where patch-package was complaining about failing when it had, in
  fact, succeeded. See #31

## 3.6.0

- Remove git headers from patch files to prevent git from thinking files are
  part of the index

## 3.5.3

- Change the way patch files are re-written when the project root dir is not the
  same as the git root dir.
- Remove redundant windows warning about whitespace

## 3.5.2

- Update README

## 3.5.1

- Ignore end-of-line whitespace when creating and applying patch files

## 3.5.0

- Add support for filtering particular paths

## 3.4.6

- Ensure use of unix-style paths everywhere

## 3.4.5

- Fix description of a hunk in error message

## 3.4.4

- Don't use `git apply` with `--unsafe-paths` since it is useless.

## 3.4.3

- Fix typo in error message

## 3.4.2

- Revert previous fix and suggest editing .gitattributes as an alternative

## 3.4.1

- Fix CRLF handling on Windows

## 3.4.0

- Add npm shrinkwrap support

## 3.3.6

- Use posix paths and line separators even on windows, for git's sake.

## 3.3.5

- Resolve paths in patch files for situations where the git root is not the same
  as the app root.

## 3.3.4

- Pass --unsafe-paths option to `git apply` to let it work on arbitrary file
  paths (i.e. files which are not in a git repo or files which are outside of
  the working directory)

## 3.3.3

- Fix bug introduced in 3.3.2
- Add progress reporting during patch making

## 3.3.2

Windows fixes:

- Use `cross-spawn` for spawning child processes
- Use `git apply` for applying patches, rather than `patch`

Contribution by [@ashmind](https://github.com/ashmind)

## 3.3.1

- Use `fs-extra` to copy files instead of the `cp` shell command, which doesn't
  work on Windows

## 3.3.0

- Use `+` instead of `:` in patch file names because `:` is illegal on Windows.

## 3.2.1

- Make update-notifier message show local install

## 3.2.0

- Add update-notifier to notify users of patch-package updates

## 3.1.0

- Add support for scoped packages

## 3.0.0

- Add support for npm5
- Make yarn patching a cli option, off by default

## 2.1.1

Improve error messages

## 2.1.0

Reduce yarn error to a warning

## 2.0.0

- Require yarn as a peer dependency
- Remove support for NPM

Moving fast and breaking things. It turns out yarn doesn't run the prepare hook
after removing a package, so we use patch-package to patch a local version of
yarn. I'm not proud of this. Probably wouldn't have released this in the first
place if I had known that yarn didn't have all the right hooks. Oh well. Now I
have a reason to contribute to Yarn I guess.

## 1.2.1

- Fix patch creation logic around nested node_modules

## 1.2.0

- Enable picking up changes in nested node_modules folders.
- Enable adding new files, not just patching existing files.

## 1.1.1

- Fix bug that made exit code 1 regardless of success or failure.

## 1.1.0

- Make applying patches work
- Add nice colorful log messages.

## 1.0.0

Initial broken release
