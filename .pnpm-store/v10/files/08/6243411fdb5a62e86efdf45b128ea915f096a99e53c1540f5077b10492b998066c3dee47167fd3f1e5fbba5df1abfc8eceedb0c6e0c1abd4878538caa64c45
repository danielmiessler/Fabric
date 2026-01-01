"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.logPatchSequenceError = exports.makePatch = void 0;
const chalk_1 = __importDefault(require("chalk"));
const console_1 = __importDefault(require("console"));
const fs_1 = require("fs");
const fs_extra_1 = require("fs-extra");
const tmp_1 = require("tmp");
const zlib_1 = require("zlib");
const applyPatches_1 = require("./applyPatches");
const createIssue_1 = require("./createIssue");
const filterFiles_1 = require("./filterFiles");
const getPackageResolution_1 = require("./getPackageResolution");
const getPackageVersion_1 = require("./getPackageVersion");
const hash_1 = require("./hash");
const PackageDetails_1 = require("./PackageDetails");
const parse_1 = require("./patch/parse");
const patchFs_1 = require("./patchFs");
const path_1 = require("./path");
const resolveRelativeFileDependencies_1 = require("./resolveRelativeFileDependencies");
const spawnSafe_1 = require("./spawnSafe");
const stateFile_1 = require("./stateFile");
function printNoPackageFoundError(packageName, packageJsonPath) {
    console_1.default.log(`No such package ${packageName}

  File not found: ${packageJsonPath}`);
}
function makePatch({ packagePathSpecifier, appPath, packageManager, includePaths, excludePaths, patchDir, createIssue, mode, }) {
    var _a, _b, _c, _d, _e;
    const packageDetails = PackageDetails_1.getPatchDetailsFromCliString(packagePathSpecifier);
    if (!packageDetails) {
        console_1.default.log("No such package", packagePathSpecifier);
        return;
    }
    const state = stateFile_1.getPatchApplicationState(packageDetails);
    const isRebasing = (_a = state === null || state === void 0 ? void 0 : state.isRebasing) !== null && _a !== void 0 ? _a : false;
    // If we are rebasing and no patches have been applied, --append is the only valid option because
    // there are no previous patches to overwrite/update
    if (isRebasing &&
        (state === null || state === void 0 ? void 0 : state.patches.filter((p) => p.didApply).length) === 0 &&
        mode.type === "overwrite_last") {
        mode = { type: "append", name: "initial" };
    }
    if (isRebasing && state) {
        stateFile_1.verifyAppliedPatches({ appPath, patchDir, state });
    }
    if (mode.type === "overwrite_last" &&
        isRebasing &&
        (state === null || state === void 0 ? void 0 : state.patches.length) === 0) {
        mode = { type: "append", name: "initial" };
    }
    const existingPatches = patchFs_1.getGroupedPatches(patchDir).pathSpecifierToPatchFiles[packageDetails.pathSpecifier] || [];
    // apply all existing patches if appending
    // otherwise apply all but the last
    const previouslyAppliedPatches = state === null || state === void 0 ? void 0 : state.patches.filter((p) => p.didApply);
    const patchesToApplyBeforeDiffing = isRebasing
        ? mode.type === "append"
            ? existingPatches.slice(0, previouslyAppliedPatches.length)
            : state.patches[state.patches.length - 1].didApply
                ? existingPatches.slice(0, previouslyAppliedPatches.length - 1)
                : existingPatches.slice(0, previouslyAppliedPatches.length)
        : mode.type === "append"
            ? existingPatches
            : existingPatches.slice(0, -1);
    if (createIssue && mode.type === "append") {
        console_1.default.log("--create-issue is not compatible with --append.");
        process.exit(1);
    }
    if (createIssue && isRebasing) {
        console_1.default.log("--create-issue is not compatible with rebasing.");
        process.exit(1);
    }
    const numPatchesAfterCreate = mode.type === "append" || existingPatches.length === 0
        ? existingPatches.length + 1
        : existingPatches.length;
    const vcs = createIssue_1.getPackageVCSDetails(packageDetails);
    const canCreateIssue = !isRebasing &&
        createIssue_1.shouldRecommendIssue(vcs) &&
        numPatchesAfterCreate === 1 &&
        mode.type !== "append";
    const appPackageJson = require(path_1.join(appPath, "package.json"));
    const packagePath = path_1.join(appPath, packageDetails.path);
    const packageJsonPath = path_1.join(packagePath, "package.json");
    if (!fs_extra_1.existsSync(packageJsonPath)) {
        printNoPackageFoundError(packagePathSpecifier, packageJsonPath);
        process.exit(1);
    }
    const tmpRepo = tmp_1.dirSync({ unsafeCleanup: true });
    const tmpRepoPackagePath = path_1.join(tmpRepo.name, packageDetails.path);
    const tmpRepoNpmRoot = tmpRepoPackagePath.slice(0, -`/node_modules/${packageDetails.name}`.length);
    const tmpRepoPackageJsonPath = path_1.join(tmpRepoNpmRoot, "package.json");
    try {
        const patchesDir = path_1.resolve(path_1.join(appPath, patchDir));
        console_1.default.info(chalk_1.default.grey("•"), "Creating temporary folder");
        // make a blank package.json
        fs_extra_1.mkdirpSync(tmpRepoNpmRoot);
        fs_extra_1.writeFileSync(tmpRepoPackageJsonPath, JSON.stringify({
            dependencies: {
                [packageDetails.name]: getPackageResolution_1.getPackageResolution({
                    packageDetails,
                    packageManager,
                    appPath,
                }),
            },
            resolutions: resolveRelativeFileDependencies_1.resolveRelativeFileDependencies(appPath, appPackageJson.resolutions || {}),
        }));
        const packageVersion = getPackageVersion_1.getPackageVersion(path_1.join(path_1.resolve(packageDetails.path), "package.json"));
        [".npmrc", ".yarnrc", ".yarn"].forEach((rcFile) => {
            const rcPath = path_1.join(appPath, rcFile);
            if (fs_extra_1.existsSync(rcPath)) {
                fs_extra_1.copySync(rcPath, path_1.join(tmpRepo.name, rcFile), { dereference: true });
            }
        });
        if (packageManager === "yarn") {
            console_1.default.info(chalk_1.default.grey("•"), `Installing ${packageDetails.name}@${packageVersion} with yarn`);
            try {
                // try first without ignoring scripts in case they are required
                // this works in 99.99% of cases
                spawnSafe_1.spawnSafeSync(`yarn`, ["install", "--ignore-engines"], {
                    cwd: tmpRepoNpmRoot,
                    logStdErrOnError: false,
                });
            }
            catch (e) {
                // try again while ignoring scripts in case the script depends on
                // an implicit context which we haven't reproduced
                spawnSafe_1.spawnSafeSync(`yarn`, ["install", "--ignore-engines", "--ignore-scripts"], {
                    cwd: tmpRepoNpmRoot,
                });
            }
        }
        else {
            console_1.default.info(chalk_1.default.grey("•"), `Installing ${packageDetails.name}@${packageVersion} with npm`);
            try {
                // try first without ignoring scripts in case they are required
                // this works in 99.99% of cases
                spawnSafe_1.spawnSafeSync(`npm`, ["i", "--force"], {
                    cwd: tmpRepoNpmRoot,
                    logStdErrOnError: false,
                    stdio: "ignore",
                });
            }
            catch (e) {
                // try again while ignoring scripts in case the script depends on
                // an implicit context which we haven't reproduced
                spawnSafe_1.spawnSafeSync(`npm`, ["i", "--ignore-scripts", "--force"], {
                    cwd: tmpRepoNpmRoot,
                    stdio: "ignore",
                });
            }
        }
        const git = (...args) => spawnSafe_1.spawnSafeSync("git", args, {
            cwd: tmpRepo.name,
            env: Object.assign(Object.assign({}, process.env), { HOME: tmpRepo.name }),
            maxBuffer: 1024 * 1024 * 100,
        });
        // remove nested node_modules just to be safe
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, "node_modules"));
        // remove .git just to be safe
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, ".git"));
        // remove patch-package state file
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, stateFile_1.STATE_FILE_NAME));
        // commit the package
        console_1.default.info(chalk_1.default.grey("•"), "Diffing your files with clean files");
        fs_extra_1.writeFileSync(path_1.join(tmpRepo.name, ".gitignore"), "!/node_modules\n\n");
        git("init");
        git("config", "--local", "user.name", "patch-package");
        git("config", "--local", "user.email", "patch@pack.age");
        // remove ignored files first
        filterFiles_1.removeIgnoredFiles(tmpRepoPackagePath, includePaths, excludePaths);
        for (const patchDetails of patchesToApplyBeforeDiffing) {
            if (!applyPatches_1.applyPatch({
                patchDetails,
                patchDir,
                patchFilePath: path_1.join(appPath, patchDir, patchDetails.patchFilename),
                reverse: false,
                cwd: tmpRepo.name,
                bestEffort: false,
            })) {
                // TODO: add better error message once --rebase is implemented
                console_1.default.log(`Failed to apply patch ${patchDetails.patchFilename} to ${packageDetails.pathSpecifier}`);
                process.exit(1);
            }
        }
        git("add", "-f", packageDetails.path);
        git("commit", "--allow-empty", "-m", "init");
        // replace package with user's version
        fs_extra_1.removeSync(tmpRepoPackagePath);
        // pnpm installs packages as symlinks, copySync would copy only the symlink
        fs_extra_1.copySync(fs_extra_1.realpathSync(packagePath), tmpRepoPackagePath);
        // remove nested node_modules just to be safe
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, "node_modules"));
        // remove .git just to be safe
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, ".git"));
        // remove patch-package state file
        fs_extra_1.removeSync(path_1.join(tmpRepoPackagePath, stateFile_1.STATE_FILE_NAME));
        // also remove ignored files like before
        filterFiles_1.removeIgnoredFiles(tmpRepoPackagePath, includePaths, excludePaths);
        // stage all files
        git("add", "-f", packageDetails.path);
        // get diff of changes
        const diffResult = git("diff", "--cached", "--no-color", "--ignore-space-at-eol", "--no-ext-diff", "--src-prefix=a/", "--dst-prefix=b/");
        if (diffResult.stdout.length === 0) {
            console_1.default.log(`⁉️  Not creating patch file for package '${packagePathSpecifier}'`);
            console_1.default.log(`⁉️  There don't appear to be any changes.`);
            if (isRebasing && mode.type === "overwrite_last") {
                console_1.default.log("\n💡 To remove a patch file, delete it and then reinstall node_modules from scratch.");
            }
            process.exit(1);
            return;
        }
        try {
            parse_1.parsePatchFile(diffResult.stdout.toString());
        }
        catch (e) {
            if (e.message.includes("Unexpected file mode string: 120000")) {
                console_1.default.log(`
⛔️ ${chalk_1.default.red.bold("ERROR")}

  Your changes involve creating symlinks. patch-package does not yet support
  symlinks.
  
  ️Please use ${chalk_1.default.bold("--include")} and/or ${chalk_1.default.bold("--exclude")} to narrow the scope of your patch if
  this was unintentional.
`);
            }
            else {
                const outPath = "./patch-package-error.json.gz";
                fs_extra_1.writeFileSync(outPath, zlib_1.gzipSync(JSON.stringify({
                    error: { message: e.message, stack: e.stack },
                    patch: diffResult.stdout.toString(),
                })));
                console_1.default.log(`
⛔️ ${chalk_1.default.red.bold("ERROR")}
        
  patch-package was unable to read the patch-file made by git. This should not
  happen.
  
  A diagnostic file was written to
  
    ${outPath}
  
  Please attach it to a github issue
  
    https://github.com/ds300/patch-package/issues/new?title=New+patch+parse+failed&body=Please+attach+the+diagnostic+file+by+dragging+it+into+here+🙏
  
  Note that this diagnostic file will contain code from the package you were
  attempting to patch.

`);
            }
            process.exit(1);
            return;
        }
        // maybe delete existing
        if (mode.type === "append" && !isRebasing && existingPatches.length === 1) {
            // if we are appending to an existing patch that doesn't have a sequence number let's rename it
            const prevPatch = existingPatches[0];
            if (prevPatch.sequenceNumber === undefined) {
                const newFileName = createPatchFileName({
                    packageDetails,
                    packageVersion,
                    sequenceNumber: 1,
                    sequenceName: (_b = prevPatch.sequenceName) !== null && _b !== void 0 ? _b : "initial",
                });
                const oldPath = path_1.join(appPath, patchDir, prevPatch.patchFilename);
                const newPath = path_1.join(appPath, patchDir, newFileName);
                fs_1.renameSync(oldPath, newPath);
                prevPatch.sequenceNumber = 1;
                prevPatch.patchFilename = newFileName;
                prevPatch.sequenceName = (_c = prevPatch.sequenceName) !== null && _c !== void 0 ? _c : "initial";
            }
        }
        const lastPatch = existingPatches[state ? state.patches.length - 1 : existingPatches.length - 1];
        const sequenceName = mode.type === "append" ? mode.name : lastPatch === null || lastPatch === void 0 ? void 0 : lastPatch.sequenceName;
        const sequenceNumber = mode.type === "append"
            ? ((_d = lastPatch === null || lastPatch === void 0 ? void 0 : lastPatch.sequenceNumber) !== null && _d !== void 0 ? _d : 0) + 1
            : lastPatch === null || lastPatch === void 0 ? void 0 : lastPatch.sequenceNumber;
        const patchFileName = createPatchFileName({
            packageDetails,
            packageVersion,
            sequenceName,
            sequenceNumber,
        });
        const patchPath = path_1.join(patchesDir, patchFileName);
        if (!fs_extra_1.existsSync(path_1.dirname(patchPath))) {
            // scoped package
            fs_extra_1.mkdirSync(path_1.dirname(patchPath));
        }
        // if we are inserting a new patch into a sequence we most likely need to update the sequence numbers
        if (isRebasing && mode.type === "append") {
            const patchesToNudge = existingPatches.slice(state.patches.length);
            if (sequenceNumber === undefined) {
                throw new Error("sequenceNumber is undefined while rebasing");
            }
            if (((_e = patchesToNudge[0]) === null || _e === void 0 ? void 0 : _e.sequenceNumber) !== undefined &&
                patchesToNudge[0].sequenceNumber <= sequenceNumber) {
                let next = sequenceNumber + 1;
                for (const p of patchesToNudge) {
                    const newName = createPatchFileName({
                        packageDetails,
                        packageVersion,
                        sequenceName: p.sequenceName,
                        sequenceNumber: next++,
                    });
                    console_1.default.log("Renaming", chalk_1.default.bold(p.patchFilename), "to", chalk_1.default.bold(newName));
                    const oldPath = path_1.join(appPath, patchDir, p.patchFilename);
                    const newPath = path_1.join(appPath, patchDir, newName);
                    fs_1.renameSync(oldPath, newPath);
                }
            }
        }
        fs_extra_1.writeFileSync(patchPath, diffResult.stdout);
        console_1.default.log(`${chalk_1.default.green("✔")} Created file ${path_1.join(patchDir, patchFileName)}\n`);
        const prevState = patchesToApplyBeforeDiffing.map((p) => ({
            patchFilename: p.patchFilename,
            didApply: true,
            patchContentHash: hash_1.hashFile(path_1.join(appPath, patchDir, p.patchFilename)),
        }));
        const nextState = [
            ...prevState,
            {
                patchFilename: patchFileName,
                didApply: true,
                patchContentHash: hash_1.hashFile(patchPath),
            },
        ];
        // if any patches come after this one we just made, we should reapply them
        let didFailWhileFinishingRebase = false;
        if (isRebasing) {
            const currentPatches = patchFs_1.getGroupedPatches(path_1.join(appPath, patchDir))
                .pathSpecifierToPatchFiles[packageDetails.pathSpecifier];
            const previouslyUnappliedPatches = currentPatches.slice(nextState.length);
            if (previouslyUnappliedPatches.length) {
                console_1.default.log(`Fast forwarding...`);
                for (const patch of previouslyUnappliedPatches) {
                    const patchFilePath = path_1.join(appPath, patchDir, patch.patchFilename);
                    if (!applyPatches_1.applyPatch({
                        patchDetails: patch,
                        patchDir,
                        patchFilePath,
                        reverse: false,
                        cwd: process.cwd(),
                        bestEffort: false,
                    })) {
                        didFailWhileFinishingRebase = true;
                        logPatchSequenceError({ patchDetails: patch });
                        nextState.push({
                            patchFilename: patch.patchFilename,
                            didApply: false,
                            patchContentHash: hash_1.hashFile(patchFilePath),
                        });
                        break;
                    }
                    else {
                        console_1.default.log(`  ${chalk_1.default.green("✔")} ${patch.patchFilename}`);
                        nextState.push({
                            patchFilename: patch.patchFilename,
                            didApply: true,
                            patchContentHash: hash_1.hashFile(patchFilePath),
                        });
                    }
                }
            }
        }
        if (isRebasing || numPatchesAfterCreate > 1) {
            stateFile_1.savePatchApplicationState({
                packageDetails,
                patches: nextState,
                isRebasing: didFailWhileFinishingRebase,
            });
        }
        else {
            stateFile_1.clearPatchApplicationState(packageDetails);
        }
        if (canCreateIssue) {
            if (createIssue) {
                createIssue_1.openIssueCreationLink({
                    packageDetails,
                    patchFileContents: diffResult.stdout.toString(),
                    packageVersion,
                    patchPath,
                });
            }
            else {
                createIssue_1.maybePrintIssueCreationPrompt(vcs, packageDetails, packageManager);
            }
        }
    }
    catch (e) {
        console_1.default.log(e);
        throw e;
    }
    finally {
        tmpRepo.removeCallback();
    }
}
exports.makePatch = makePatch;
function createPatchFileName({ packageDetails, packageVersion, sequenceNumber, sequenceName, }) {
    const packageNames = packageDetails.packageNames
        .map((name) => name.replace(/\//g, "+"))
        .join("++");
    const nameAndVersion = `${packageNames}+${packageVersion}`;
    const num = sequenceNumber === undefined
        ? ""
        : `+${sequenceNumber.toString().padStart(3, "0")}`;
    const name = !sequenceName ? "" : `+${sequenceName}`;
    return `${nameAndVersion}${num}${name}.patch`;
}
function logPatchSequenceError({ patchDetails, }) {
    console_1.default.log(`
${chalk_1.default.red.bold("⛔ ERROR")}

Failed to apply patch file ${chalk_1.default.bold(patchDetails.patchFilename)}.

If this patch file is no longer useful, delete it and run

  ${chalk_1.default.bold(`patch-package`)}

To partially apply the patch (if possible) and output a log of errors to fix, run

  ${chalk_1.default.bold(`patch-package --partial`)}

After which you should make any required changes inside ${patchDetails.path}, and finally run

  ${chalk_1.default.bold(`patch-package ${patchDetails.pathSpecifier}`)}

to update the patch file.
`);
}
exports.logPatchSequenceError = logPatchSequenceError;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibWFrZVBhdGNoLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL21ha2VQYXRjaC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7Ozs7QUFBQSxrREFBeUI7QUFDekIsc0RBQTZCO0FBQzdCLDJCQUErQjtBQUMvQix1Q0FRaUI7QUFDakIsNkJBQTZCO0FBQzdCLCtCQUErQjtBQUMvQixpREFBMkM7QUFDM0MsK0NBS3NCO0FBRXRCLCtDQUFrRDtBQUNsRCxpRUFBNkQ7QUFDN0QsMkRBQXVEO0FBQ3ZELGlDQUFpQztBQUNqQyxxREFJeUI7QUFDekIseUNBQThDO0FBQzlDLHVDQUE2QztBQUM3QyxpQ0FBK0M7QUFDL0MsdUZBQW1GO0FBQ25GLDJDQUEyQztBQUMzQywyQ0FPb0I7QUFFcEIsU0FBUyx3QkFBd0IsQ0FDL0IsV0FBbUIsRUFDbkIsZUFBdUI7SUFFdkIsaUJBQU8sQ0FBQyxHQUFHLENBQ1QsbUJBQW1CLFdBQVc7O29CQUVkLGVBQWUsRUFBRSxDQUNsQyxDQUFBO0FBQ0gsQ0FBQztBQUVELFNBQWdCLFNBQVMsQ0FBQyxFQUN4QixvQkFBb0IsRUFDcEIsT0FBTyxFQUNQLGNBQWMsRUFDZCxZQUFZLEVBQ1osWUFBWSxFQUNaLFFBQVEsRUFDUixXQUFXLEVBQ1gsSUFBSSxHQVVMOztJQUNDLE1BQU0sY0FBYyxHQUFHLDZDQUE0QixDQUFDLG9CQUFvQixDQUFDLENBQUE7SUFFekUsSUFBSSxDQUFDLGNBQWMsRUFBRTtRQUNuQixpQkFBTyxDQUFDLEdBQUcsQ0FBQyxpQkFBaUIsRUFBRSxvQkFBb0IsQ0FBQyxDQUFBO1FBQ3BELE9BQU07S0FDUDtJQUVELE1BQU0sS0FBSyxHQUFHLG9DQUF3QixDQUFDLGNBQWMsQ0FBQyxDQUFBO0lBQ3RELE1BQU0sVUFBVSxHQUFHLE1BQUEsS0FBSyxhQUFMLEtBQUssdUJBQUwsS0FBSyxDQUFFLFVBQVUsbUNBQUksS0FBSyxDQUFBO0lBRTdDLGlHQUFpRztJQUNqRyxvREFBb0Q7SUFDcEQsSUFDRSxV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QXlCLENBQUM7OzBEQUdyQyxZQUFZLENBQUMsSUFDZjs7SUFFRSxlQUFLLENBQUMsSUFBSSxDQUFDLGlCQUFpQixZQUFZLENBQUMsYUFBYSxFQUFFLENBQUM7OztDQUc1RCxDQUFDLENBQUE7QUFDRixDQUFDO0FBMUJELHNEQTBCQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBjaGFsayBmcm9tIFwiY2hhbGtcIlxuaW1wb3J0IGNvbnNvbGUgZnJvbSBcImNvbnNvbGVcIlxuaW1wb3J0IHsgcmVuYW1lU3luYyB9IGZyb20gXCJmc1wiXG5pbXBvcnQge1xuICBjb3B5U3luYyxcbiAgZXhpc3RzU3luYyxcbiAgbWtkaXJwU3luYyxcbiAgbWtkaXJTeW5jLFxuICByZWFscGF0aFN5bmMsXG4gIHJlbW92ZVN5bmMsXG4gIHdyaXRlRmlsZVN5bmMsXG59IGZyb20gXCJmcy1leHRyYVwiXG5pbXBvcnQgeyBkaXJTeW5jIH0gZnJvbSBcInRtcFwiXG5pbXBvcnQgeyBnemlwU3luYyB9IGZyb20gXCJ6bGliXCJcbmltcG9ydCB7IGFwcGx5UGF0Y2ggfSBmcm9tIFwiLi9hcHBseVBhdGNoZXNcIlxuaW1wb3J0IHtcbiAgZ2V0UGFja2FnZVZDU0RldGFpbHMsXG4gIG1heWJlUHJpbnRJc3N1ZUNyZWF0aW9uUHJvbXB0LFxuICBvcGVuSXNzdWVDcmVhdGlvbkxpbmssXG4gIHNob3VsZFJlY29tbWVuZElzc3VlLFxufSBmcm9tIFwiLi9jcmVhdGVJc3N1ZVwiXG5pbXBvcnQgeyBQYWNrYWdlTWFuYWdlciB9IGZyb20gXCIuL2RldGVjdFBhY2thZ2VNYW5hZ2VyXCJcbmltcG9ydCB7IHJlbW92ZUlnbm9yZWRGaWxlcyB9IGZyb20gXCIuL2ZpbHRlckZpbGVzXCJcbmltcG9ydCB7IGdldFBhY2thZ2VSZXNvbHV0aW9uIH0gZnJvbSBcIi4vZ2V0UGFja2FnZVJlc29sdXRpb25cIlxuaW1wb3J0IHsgZ2V0UGFja2FnZVZlcnNpb24gfSBmcm9tIFwiLi9nZXRQYWNrYWdlVmVyc2lvblwiXG5pbXBvcnQgeyBoYXNoRmlsZSB9IGZyb20gXCIuL2hhc2hcIlxuaW1wb3J0IHtcbiAgZ2V0UGF0Y2hEZXRhaWxzRnJvbUNsaVN0cmluZyxcbiAgUGFja2FnZURldGFpbHMsXG4gIFBhdGNoZWRQYWNrYWdlRGV0YWlscyxcbn0gZnJvbSBcIi4vUGFja2FnZURldGFpbHNcIlxuaW1wb3J0IHsgcGFyc2VQYXRjaEZpbGUgfSBmcm9tIFwiLi9wYXRjaC9wYXJzZVwiXG5pbXBvcnQgeyBnZXRHcm91cGVkUGF0Y2hlcyB9IGZyb20gXCIuL3BhdGNoRnNcIlxuaW1wb3J0IHsgZGlybmFtZSwgam9pbiwgcmVzb2x2ZSB9IGZyb20gXCIuL3BhdGhcIlxuaW1wb3J0IHsgcmVzb2x2ZVJlbGF0aXZlRmlsZURlcGVuZGVuY2llcyB9IGZyb20gXCIuL3Jlc29sdmVSZWxhdGl2ZUZpbGVEZXBlbmRlbmNpZXNcIlxuaW1wb3J0IHsgc3Bhd25TYWZlU3luYyB9IGZyb20gXCIuL3NwYXduU2FmZVwiXG5pbXBvcnQge1xuICBjbGVhclBhdGNoQXBwbGljYXRpb25TdGF0ZSxcbiAgZ2V0UGF0Y2hBcHBsaWNhdGlvblN0YXRlLFxuICBQYXRjaFN0YXRlLFxuICBzYXZlUGF0Y2hBcHBsaWNhdGlvblN0YXRlLFxuICBTVEFURV9GSUxFX05BTUUsXG4gIHZlcmlmeUFwcGxpZWRQYXRjaGVzLFxufSBmcm9tIFwiLi9zdGF0ZUZpbGVcIlxuXG5mdW5jdGlvbiBwcmludE5vUGFja2FnZUZvdW5kRXJyb3IoXG4gIHBhY2thZ2VOYW1lOiBzdHJpbmcsXG4gIHBhY2thZ2VKc29uUGF0aDogc3RyaW5nLFxuKSB7XG4gIGNvbnNvbGUubG9nKFxuICAgIGBObyBzdWNoIHBhY2thZ2UgJHtwYWNrYWdlTmFtZX1cblxuICBGaWxlIG5vdCBmb3VuZDogJHtwYWNrYWdlSnNvblBhdGh9YCxcbiAgKVxufVxuXG5leHBvcnQgZnVuY3Rpb24gbWFrZVBhdGNoKHtcbiAgcGFja2FnZVBhdGhTcGVjaWZpZXIsXG4gIGFwcFBhdGgsXG4gIHBhY2thZ2VNYW5hZ2VyLFxuICBpbmNsdWRlUGF0aHMsXG4gIGV4Y2x1ZGVQYXRocyxcbiAgcGF0Y2hEaXIsXG4gIGNyZWF0ZUlzc3VlLFxuICBtb2RlLFxufToge1xuICBwYWNrYWdlUGF0aFNwZWNpZmllcjogc3RyaW5nXG4gIGFwcFBhdGg6IHN0cmluZ1xuICBwYWNrYWdlTWFuYWdlcjogUGFja2FnZU1hbmFnZXJcbiAgaW5jbHVkZVBhdGhzOiBSZWdFeHBcbiAgZXhjbHVkZVBhdGhzOiBSZWdFeHBcbiAgcGF0Y2hEaXI6IHN0cmluZ1xuICBjcmVhdGVJc3N1ZTogYm9vbGVhblxuICBtb2RlOiB7IHR5cGU6IFwib3ZlcndyaXRlX2xhc3RcIiB9IHwgeyB0eXBlOiBcImFwcGVuZFwiOyBuYW1lPzogc3RyaW5nIH1cbn0pIHtcbiAgY29uc3QgcGFja2FnZURldGFpbHMgPSBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nKHBhY2thZ2VQYXRoU3BlY2lmaWVyKVxuXG4gIGlmICghcGFja2FnZURldGFpbHMpIHtcbiAgICBjb25zb2xlLmxvZyhcIk5vIHN1Y2ggcGFja2FnZVwiLCBwYWNrYWdlUGF0aFNwZWNpZmllcilcbiAgICByZXR1cm5cbiAgfVxuXG4gIGNvbnN0IHN0YXRlID0gZ2V0UGF0Y2hBcHBsaWNhdGlvblN0YXRlKHBhY2thZ2VEZXRhaWxzKVxuICBjb25zdCBpc1JlYmFzaW5nID0gc3RhdGU/LmlzUmViYXNpbmcgPz8gZmFsc2VcblxuICAvLyBJZiB3ZSBhcmUgcmViYXNpbmcgYW5kIG5vIHBhdGNoZXMgaGF2ZSBiZWVuIGFwcGxpZWQsIC0tYXBwZW5kIGlzIHRoZSBvbmx5IHZhbGlkIG9wdGlvbiBiZWNhdXNlXG4gIC8vIHRoZXJlIGFyZSBubyBwcmV2aW91cyBwYXRjaGVzIHRvIG92ZXJ3cml0ZS91cGRhdGVcbiAgaWYgKFxuICAgIGlzUmViYXNpbmcgJiZcbiAgICBzdGF0ZT8ucGF0Y2hlcy5maWx0ZXIoKHApID0+IHAuZGlkQXBwbHkpLmxlbmd0aCA9PT0gMCAmJlxuICAgIG1vZGUudHlwZSA9PT0gXCJvdmVyd3JpdGVfbGFzdFwiXG4gICkge1xuICAgIG1vZGUgPSB7IHR5cGU6IFwiYXBwZW5kXCIsIG5hbWU6IFwiaW5pdGlhbFwiIH1cbiAgfVxuXG4gIGlmIChpc1JlYmFzaW5nICYmIHN0YXRlKSB7XG4gICAgdmVyaWZ5QXBwbGllZFBhdGNoZXMoeyBhcHBQYXRoLCBwYXRjaERpciwgc3RhdGUgfSlcbiAgfVxuXG4gIGlmIChcbiAgICBtb2RlLnR5cGUgPT09IFwib3ZlcndyaXRlX2xhc3RcIiAmJlxuICAgIGlzUmViYXNpbmcgJiZcbiAgICBzdGF0ZT8ucGF0Y2hlcy5sZW5ndGggPT09IDBcbiAgKSB7XG4gICAgbW9kZSA9IHsgdHlwZTogXCJhcHBlbmRcIiwgbmFtZTogXCJpbml0aWFsXCIgfVxuICB9XG5cbiAgY29uc3QgZXhpc3RpbmdQYXRjaGVzID1cbiAgICBnZXRHcm91cGVkUGF0Y2hlcyhwYXRjaERpcikucGF0aFNwZWNpZmllclRvUGF0Y2hGaWxlc1tcbiAgICAgIHBhY2thZ2VEZXRhaWxzLnBhdGhTcGVjaWZpZXJcbiAgICBdIHx8IFtdXG5cbiAgLy8gYXBwbHkgYWxsIGV4aXN0aW5nIHBhdGNoZXMgaWYgYXBwZW5kaW5nXG4gIC8vIG90aGVyd2lzZSBhcHBseSBhbGwgYnV0IHRoZSBsYXN0XG4gIGNvbnN0IHByZXZpb3VzbHlBcHBsaWVkUGF0Y2hlcyA9IHN0YXRlPy5wYXRjaGVzLmZpbHRlcigocCkgPT4gcC5kaWRBcHBseSlcbiAgY29uc3QgcGF0Y2hlc1RvQXBwbHlCZWZvcmVEaWZmaW5nOiBQYXRjaGVkUGFja2FnZURldGFpbHNbXSA9IGlzUmViYXNpbmdcbiAgICA/IG1vZGUudHlwZSA9PT0gXCJhcHBlbmRcIlxuICAgICAgPyBleGlzdGluZ1BhdGNoZXMuc2xpY2UoMCwgcHJldmlvdXNseUFwcGxpZWRQYXRjaGVzIS5sZW5ndGgpXG4gICAgICA6IHN0YXRlIS5wYXRjaGVzW3N0YXRlIS5wYXRjaGVzLmxlbmd0aCAtIDFdLmRpZEFwcGx5XG4gICAgICA/IGV4aXN0aW5nUGF0Y2hlcy5zbGljZSgwLCBwcmV2aW91c2x5QXBwbGllZFBhdGNoZXMhLmxlbmd0aCAtIDEpXG4gICAgICA6IGV4aXN0aW5nUGF0Y2hlcy5zbGljZSgwLCBwcmV2aW91c2x5QXBwbGllZFBhdGNoZXMhLmxlbmd0aClcbiAgICA6IG1vZGUudHlwZSA9PT0gXCJhcHBlbmRcIlxuICAgID8gZXhpc3RpbmdQYXRjaGVzXG4gICAgOiBleGlzdGluZ1BhdGNoZXMuc2xpY2UoMCwgLTEpXG5cbiAgaWYgKGNyZWF0ZUlzc3VlICYmIG1vZGUudHlwZSA9PT0gXCJhcHBlbmRcIikge1xuICAgIGNvbnNvbGUubG9nKFwiLS1jcmVhdGUtaXNzdWUgaXMgbm90IGNvbXBhdGlibGUgd2l0aCAtLWFwcGVuZC5cIilcbiAgICBwcm9jZXNzLmV4aXQoMSlcbiAgfVxuXG4gIGlmIChjcmVhdGVJc3N1ZSAmJiBpc1JlYmFzaW5nKSB7XG4gICAgY29uc29sZS5sb2coXCItLWNyZWF0ZS1pc3N1ZSBpcyBub3QgY29tcGF0aWJsZSB3aXRoIHJlYmFzaW5nLlwiKVxuICAgIHByb2Nlc3MuZXhpdCgxKVxuICB9XG5cbiAgY29uc3QgbnVtUGF0Y2hlc0FmdGVyQ3JlYXRlID1cbiAgICBtb2RlLnR5cGUgPT09IFwiYXBwZW5kXCIgfHwgZXhpc3RpbmdQYXRjaGVzLmxlbmd0aCA9PT0gMFxuICAgICAgPyBleGlzdGluZ1BhdGNoZXMubGVuZ3RoICsgMVxuICAgICAgOiBleGlzdGluZ1BhdGNoZXMubGVuZ3RoXG4gIGNvbnN0IHZjcyA9IGdldFBhY2thZ2VWQ1NEZXRhaWxzKHBhY2thZ2VEZXRhaWxzKVxuICBjb25zdCBjYW5DcmVhdGVJc3N1ZSA9XG4gICAgIWlzUmViYXNpbmcgJiZcbiAgICBzaG91bGRSZWNvbW1lbmRJc3N1ZSh2Y3MpICYmXG4gICAgbnVtUGF0Y2hlc0FmdGVyQ3JlYXRlID09PSAxICYmXG4gICAgbW9kZS50eXBlICE9PSBcImFwcGVuZFwiXG5cbiAgY29uc3QgYXBwUGFja2FnZUpzb24gPSByZXF1aXJlKGpvaW4oYXBwUGF0aCwgXCJwYWNrYWdlLmpzb25cIikpXG4gIGNvbnN0IHBhY2thZ2VQYXRoID0gam9pbihhcHBQYXRoLCBwYWNrYWdlRGV0YWlscy5wYXRoKVxuICBjb25zdCBwYWNrYWdlSnNvblBhdGggPSBqb2luKHBhY2thZ2VQYXRoLCBcInBhY2thZ2UuanNvblwiKVxuXG4gIGlmICghZXhpc3RzU3luYyhwYWNrYWdlSnNvblBhdGgpKSB7XG4gICAgcHJpbnROb1BhY2thZ2VGb3VuZEVycm9yKHBhY2thZ2VQYXRoU3BlY2lmaWVyLCBwYWNrYWdlSnNvblBhdGgpXG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cblxuICBjb25zdCB0bXBSZXBvID0gZGlyU3luYyh7IHVuc2FmZUNsZWFudXA6IHRydWUgfSlcbiAgY29uc3QgdG1wUmVwb1BhY2thZ2VQYXRoID0gam9pbih0bXBSZXBvLm5hbWUsIHBhY2thZ2VEZXRhaWxzLnBhdGgpXG4gIGNvbnN0IHRtcFJlcG9OcG1Sb290ID0gdG1wUmVwb1BhY2thZ2VQYXRoLnNsaWNlKFxuICAgIDAsXG4gICAgLWAvbm9kZV9tb2R1bGVzLyR7cGFja2FnZURldGFpbHMubmFtZX1gLmxlbmd0aCxcbiAgKVxuXG4gIGNvbnN0IHRtcFJlcG9QYWNrYWdlSnNvblBhdGggPSBqb2luKHRtcFJlcG9OcG1Sb290LCBcInBhY2thZ2UuanNvblwiKVxuXG4gIHRyeSB7XG4gICAgY29uc3QgcGF0Y2hlc0RpciA9IHJlc29sdmUoam9pbihhcHBQYXRoLCBwYXRjaERpcikpXG5cbiAgICBjb25zb2xlLmluZm8oY2hhbGsuZ3JleShcIuKAolwiKSwgXCJDcmVhdGluZyB0ZW1wb3JhcnkgZm9sZGVyXCIpXG5cbiAgICAvLyBtYWtlIGEgYmxhbmsgcGFja2FnZS5qc29uXG4gICAgbWtkaXJwU3luYyh0bXBSZXBvTnBtUm9vdClcbiAgICB3cml0ZUZpbGVTeW5jKFxuICAgICAgdG1wUmVwb1BhY2thZ2VKc29uUGF0aCxcbiAgICAgIEpTT04uc3RyaW5naWZ5KHtcbiAgICAgICAgZGVwZW5kZW5jaWVzOiB7XG4gICAgICAgICAgW3BhY2thZ2VEZXRhaWxzLm5hbWVdOiBnZXRQYWNrYWdlUmVzb2x1dGlvbih7XG4gICAgICAgICAgICBwYWNrYWdlRGV0YWlscyxcbiAgICAgICAgICAgIHBhY2thZ2VNYW5hZ2VyLFxuICAgICAgICAgICAgYXBwUGF0aCxcbiAgICAgICAgICB9KSxcbiAgICAgICAgfSxcbiAgICAgICAgcmVzb2x1dGlvbnM6IHJlc29sdmVSZWxhdGl2ZUZpbGVEZXBlbmRlbmNpZXMoXG4gICAgICAgICAgYXBwUGF0aCxcbiAgICAgICAgICBhcHBQYWNrYWdlSnNvbi5yZXNvbHV0aW9ucyB8fCB7fSxcbiAgICAgICAgKSxcbiAgICAgIH0pLFxuICAgIClcblxuICAgIGNvbnN0IHBhY2thZ2VWZXJzaW9uID0gZ2V0UGFja2FnZVZlcnNpb24oXG4gICAgICBqb2luKHJlc29sdmUocGFja2FnZURldGFpbHMucGF0aCksIFwicGFja2FnZS5qc29uXCIpLFxuICAgIClcblxuICAgIC8vIGNvcHkgLm5wbXJjLy55YXJucmMgaW4gY2FzZSBwYWNrYWdlcyBhcmUgaG9zdGVkIGluIHByaXZhdGUgcmVnaXN0cnlcbiAgICAvLyBjb3B5IC55YXJuIGRpcmVjdG9yeSBhcyB3ZWxsIHRvIGVuc3VyZSBpbnN0YWxsYXRpb25zIHdvcmsgaW4geWFybiAyXG4gICAgLy8gdHNsaW50OmRpc2FibGUtbmV4dC1saW5lOmFsaWduXG4gICAgO1tcIi5ucG1yY1wiLCBcIi55YXJucmNcIiwgXCIueWFyblwiXS5mb3JFYWNoKChyY0ZpbGUpID0+IHtcbiAgICAgIGNvbnN0IHJjUGF0aCA9IGpvaW4oYXBwUGF0aCwgcmNGaWxlKVxuICAgICAgaWYgKGV4aXN0c1N5bmMocmNQYXRoKSkge1xuICAgICAgICBjb3B5U3luYyhyY1BhdGgsIGpvaW4odG1wUmVwby5uYW1lLCByY0ZpbGUpLCB7IGRlcmVmZXJlbmNlOiB0cnVlIH0pXG4gICAgICB9XG4gICAgfSlcblxuICAgIGlmIChwYWNrYWdlTWFuYWdlciA9PT0gXCJ5YXJuXCIpIHtcbiAgICAgIGNvbnNvbGUuaW5mbyhcbiAgICAgICAgY2hhbGsuZ3JleShcIuKAolwiKSxcbiAgICAgICAgYEluc3RhbGxpbmcgJHtwYWNrYWdlRGV0YWlscy5uYW1lfUAke3BhY2thZ2VWZXJzaW9ufSB3aXRoIHlhcm5gLFxuICAgICAgKVxuICAgICAgdHJ5IHtcbiAgICAgICAgLy8gdHJ5IGZpcnN0IHdpdGhvdXQgaWdub3Jpbmcgc2NyaXB0cyBpbiBjYXNlIHRoZXkgYXJlIHJlcXVpcmVkXG4gICAgICAgIC8vIHRoaXMgd29ya3MgaW4gOTkuOTklIG9mIGNhc2VzXG4gICAgICAgIHNwYXduU2FmZVN5bmMoYHlhcm5gLCBbXCJpbnN0YWxsXCIsIFwiLS1pZ25vcmUtZW5naW5lc1wiXSwge1xuICAgICAgICAgIGN3ZDogdG1wUmVwb05wbVJvb3QsXG4gICAgICAgICAgbG9nU3RkRXJyT25FcnJvcjogZmFsc2UsXG4gICAgICAgIH0pXG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIC8vIHRyeSBhZ2FpbiB3aGlsZSBpZ25vcmluZyBzY3JpcHRzIGluIGNhc2UgdGhlIHNjcmlwdCBkZXBlbmRzIG9uXG4gICAgICAgIC8vIGFuIGltcGxpY2l0IGNvbnRleHQgd2hpY2ggd2UgaGF2ZW4ndCByZXByb2R1Y2VkXG4gICAgICAgIHNwYXduU2FmZVN5bmMoXG4gICAgICAgICAgYHlhcm5gLFxuICAgICAgICAgIFtcImluc3RhbGxcIiwgXCItLWlnbm9yZS1lbmdpbmVzXCIsIFwiLS1pZ25vcmUtc2NyaXB0c1wiXSxcbiAgICAgICAgICB7XG4gICAgICAgICAgICBjd2Q6IHRtcFJlcG9OcG1Sb290LFxuICAgICAgICAgIH0sXG4gICAgICAgIClcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgY29uc29sZS5pbmZvKFxuICAgICAgICBjaGFsay5ncmV5KFwi4oCiXCIpLFxuICAgICAgICBgSW5zdGFsbGluZyAke3BhY2thZ2VEZXRhaWxzLm5hbWV9QCR7cGFja2FnZVZlcnNpb259IHdpdGggbnBtYCxcbiAgICAgIClcbiAgICAgIHRyeSB7XG4gICAgICAgIC8vIHRyeSBmaXJzdCB3aXRob3V0IGlnbm9yaW5nIHNjcmlwdHMgaW4gY2FzZSB0aGV5IGFyZSByZXF1aXJlZFxuICAgICAgICAvLyB0aGlzIHdvcmtzIGluIDk5Ljk5JSBvZiBjYXNlc1xuICAgICAgICBzcGF3blNhZmVTeW5jKGBucG1gLCBbXCJpXCIsIFwiLS1mb3JjZVwiXSwge1xuICAgICAgICAgIGN3ZDogdG1wUmVwb05wbVJvb3QsXG4gICAgICAgICAgbG9nU3RkRXJyT25FcnJvcjogZmFsc2UsXG4gICAgICAgICAgc3RkaW86IFwiaWdub3JlXCIsXG4gICAgICAgIH0pXG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIC8vIHRyeSBhZ2FpbiB3aGlsZSBpZ25vcmluZyBzY3JpcHRzIGluIGNhc2UgdGhlIHNjcmlwdCBkZXBlbmRzIG9uXG4gICAgICAgIC8vIGFuIGltcGxpY2l0IGNvbnRleHQgd2hpY2ggd2UgaGF2ZW4ndCByZXByb2R1Y2VkXG4gICAgICAgIHNwYXduU2FmZVN5bmMoYG5wbWAsIFtcImlcIiwgXCItLWlnbm9yZS1zY3JpcHRzXCIsIFwiLS1mb3JjZVwiXSwge1xuICAgICAgICAgIGN3ZDogdG1wUmVwb05wbVJvb3QsXG4gICAgICAgICAgc3RkaW86IFwiaWdub3JlXCIsXG4gICAgICAgIH0pXG4gICAgICB9XG4gICAgfVxuXG4gICAgY29uc3QgZ2l0ID0gKC4uLmFyZ3M6IHN0cmluZ1tdKSA9PlxuICAgICAgc3Bhd25TYWZlU3luYyhcImdpdFwiLCBhcmdzLCB7XG4gICAgICAgIGN3ZDogdG1wUmVwby5uYW1lLFxuICAgICAgICBlbnY6IHsgLi4ucHJvY2Vzcy5lbnYsIEhPTUU6IHRtcFJlcG8ubmFtZSB9LFxuICAgICAgICBtYXhCdWZmZXI6IDEwMjQgKiAxMDI0ICogMTAwLFxuICAgICAgfSlcblxuICAgIC8vIHJlbW92ZSBuZXN0ZWQgbm9kZV9tb2R1bGVzIGp1c3QgdG8gYmUgc2FmZVxuICAgIHJlbW92ZVN5bmMoam9pbih0bXBSZXBvUGFja2FnZVBhdGgsIFwibm9kZV9tb2R1bGVzXCIpKVxuICAgIC8vIHJlbW92ZSAuZ2l0IGp1c3QgdG8gYmUgc2FmZVxuICAgIHJlbW92ZVN5bmMoam9pbih0bXBSZXBvUGFja2FnZVBhdGgsIFwiLmdpdFwiKSlcbiAgICAvLyByZW1vdmUgcGF0Y2gtcGFja2FnZSBzdGF0ZSBmaWxlXG4gICAgcmVtb3ZlU3luYyhqb2luKHRtcFJlcG9QYWNrYWdlUGF0aCwgU1RBVEVfRklMRV9OQU1FKSlcblxuICAgIC8vIGNvbW1pdCB0aGUgcGFja2FnZVxuICAgIGNvbnNvbGUuaW5mbyhjaGFsay5ncmV5KFwi4oCiXCIpLCBcIkRpZmZpbmcgeW91ciBmaWxlcyB3aXRoIGNsZWFuIGZpbGVzXCIpXG4gICAgd3JpdGVGaWxlU3luYyhqb2luKHRtcFJlcG8ubmFtZSwgXCIuZ2l0aWdub3JlXCIpLCBcIiEvbm9kZV9tb2R1bGVzXFxuXFxuXCIpXG4gICAgZ2l0KFwiaW5pdFwiKVxuICAgIGdpdChcImNvbmZpZ1wiLCBcIi0tbG9jYWxcIiwgXCJ1c2VyLm5hbWVcIiwgXCJwYXRjaC1wYWNrYWdlXCIpXG4gICAgZ2l0KFwiY29uZmlnXCIsIFwiLS1sb2NhbFwiLCBcInVzZXIuZW1haWxcIiwgXCJwYXRjaEBwYWNrLmFnZVwiKVxuXG4gICAgLy8gcmVtb3ZlIGlnbm9yZWQgZmlsZXMgZmlyc3RcbiAgICByZW1vdmVJZ25vcmVkRmlsZXModG1wUmVwb1BhY2thZ2VQYXRoLCBpbmNsdWRlUGF0aHMsIGV4Y2x1ZGVQYXRocylcblxuICAgIGZvciAoY29uc3QgcGF0Y2hEZXRhaWxzIG9mIHBhdGNoZXNUb0FwcGx5QmVmb3JlRGlmZmluZykge1xuICAgICAgaWYgKFxuICAgICAgICAhYXBwbHlQYXRjaCh7XG4gICAgICAgICAgcGF0Y2hEZXRhaWxzLFxuICAgICAgICAgIHBhdGNoRGlyLFxuICAgICAgICAgIHBhdGNoRmlsZVBhdGg6IGpvaW4oYXBwUGF0aCwgcGF0Y2hEaXIsIHBhdGNoRGV0YWlscy5wYXRjaEZpbGVuYW1lKSxcbiAgICAgICAgICByZXZlcnNlOiBmYWxzZSxcbiAgICAgICAgICBjd2Q6IHRtcFJlcG8ubmFtZSxcbiAgICAgICAgICBiZXN0RWZmb3J0OiBmYWxzZSxcbiAgICAgICAgfSlcbiAgICAgICkge1xuICAgICAgICAvLyBUT0RPOiBhZGQgYmV0dGVyIGVycm9yIG1lc3NhZ2Ugb25jZSAtLXJlYmFzZSBpcyBpbXBsZW1lbnRlZFxuICAgICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgICBgRmFpbGVkIHRvIGFwcGx5IHBhdGNoICR7cGF0Y2hEZXRhaWxzLnBhdGNoRmlsZW5hbWV9IHRvICR7cGFja2FnZURldGFpbHMucGF0aFNwZWNpZmllcn1gLFxuICAgICAgICApXG4gICAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgICAgfVxuICAgIH1cbiAgICBnaXQoXCJhZGRcIiwgXCItZlwiLCBwYWNrYWdlRGV0YWlscy5wYXRoKVxuICAgIGdpdChcImNvbW1pdFwiLCBcIi0tYWxsb3ctZW1wdHlcIiwgXCItbVwiLCBcImluaXRcIilcblxuICAgIC8vIHJlcGxhY2UgcGFja2FnZSB3aXRoIHVzZXIncyB2ZXJzaW9uXG4gICAgcmVtb3ZlU3luYyh0bXBSZXBvUGFja2FnZVBhdGgpXG5cbiAgICAvLyBwbnBtIGluc3RhbGxzIHBhY2thZ2VzIGFzIHN5bWxpbmtzLCBjb3B5U3luYyB3b3VsZCBjb3B5IG9ubHkgdGhlIHN5bWxpbmtcbiAgICBjb3B5U3luYyhyZWFscGF0aFN5bmMocGFja2FnZVBhdGgpLCB0bXBSZXBvUGFja2FnZVBhdGgpXG5cbiAgICAvLyByZW1vdmUgbmVzdGVkIG5vZGVfbW9kdWxlcyBqdXN0IHRvIGJlIHNhZmVcbiAgICByZW1vdmVTeW5jKGpvaW4odG1wUmVwb1BhY2thZ2VQYXRoLCBcIm5vZGVfbW9kdWxlc1wiKSlcbiAgICAvLyByZW1vdmUgLmdpdCBqdXN0IHRvIGJlIHNhZmVcbiAgICByZW1vdmVTeW5jKGpvaW4odG1wUmVwb1BhY2thZ2VQYXRoLCBcIi5naXRcIikpXG4gICAgLy8gcmVtb3ZlIHBhdGNoLXBhY2thZ2Ugc3RhdGUgZmlsZVxuICAgIHJlbW92ZVN5bmMoam9pbih0bXBSZXBvUGFja2FnZVBhdGgsIFNUQVRFX0ZJTEVfTkFNRSkpXG5cbiAgICAvLyBhbHNvIHJlbW92ZSBpZ25vcmVkIGZpbGVzIGxpa2UgYmVmb3JlXG4gICAgcmVtb3ZlSWdub3JlZEZpbGVzKHRtcFJlcG9QYWNrYWdlUGF0aCwgaW5jbHVkZVBhdGhzLCBleGNsdWRlUGF0aHMpXG5cbiAgICAvLyBzdGFnZSBhbGwgZmlsZXNcbiAgICBnaXQoXCJhZGRcIiwgXCItZlwiLCBwYWNrYWdlRGV0YWlscy5wYXRoKVxuXG4gICAgLy8gZ2V0IGRpZmYgb2YgY2hhbmdlc1xuICAgIGNvbnN0IGRpZmZSZXN1bHQgPSBnaXQoXG4gICAgICBcImRpZmZcIixcbiAgICAgIFwiLS1jYWNoZWRcIixcbiAgICAgIFwiLS1uby1jb2xvclwiLFxuICAgICAgXCItLWlnbm9yZS1zcGFjZS1hdC1lb2xcIixcbiAgICAgIFwiLS1uby1leHQtZGlmZlwiLFxuICAgICAgXCItLXNyYy1wcmVmaXg9YS9cIixcbiAgICAgIFwiLS1kc3QtcHJlZml4PWIvXCIsXG4gICAgKVxuXG4gICAgaWYgKGRpZmZSZXN1bHQuc3Rkb3V0Lmxlbmd0aCA9PT0gMCkge1xuICAgICAgY29uc29sZS5sb2coXG4gICAgICAgIGDigYnvuI8gIE5vdCBjcmVhdGluZyBwYXRjaCBmaWxlIGZvciBwYWNrYWdlICcke3BhY2thZ2VQYXRoU3BlY2lmaWVyfSdgLFxuICAgICAgKVxuICAgICAgY29uc29sZS5sb2coYOKBie+4jyAgVGhlcmUgZG9uJ3QgYXBwZWFyIHRvIGJlIGFueSBjaGFuZ2VzLmApXG4gICAgICBpZiAoaXNSZWJhc2luZyAmJiBtb2RlLnR5cGUgPT09IFwib3ZlcndyaXRlX2xhc3RcIikge1xuICAgICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgICBcIlxcbvCfkqEgVG8gcmVtb3ZlIGEgcGF0Y2ggZmlsZSwgZGVsZXRlIGl0IGFuZCB0aGVuIHJlaW5zdGFsbCBub2RlX21vZHVsZXMgZnJvbSBzY3JhdGNoLlwiLFxuICAgICAgICApXG4gICAgICB9XG4gICAgICBwcm9jZXNzLmV4aXQoMSlcbiAgICAgIHJldHVyblxuICAgIH1cblxuICAgIHRyeSB7XG4gICAgICBwYXJzZVBhdGNoRmlsZShkaWZmUmVzdWx0LnN0ZG91dC50b1N0cmluZygpKVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGlmIChcbiAgICAgICAgKGUgYXMgRXJyb3IpLm1lc3NhZ2UuaW5jbHVkZXMoXCJVbmV4cGVjdGVkIGZpbGUgbW9kZSBzdHJpbmc6IDEyMDAwMFwiKVxuICAgICAgKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKGBcbuKblO+4jyAke2NoYWxrLnJlZC5ib2xkKFwiRVJST1JcIil9XG5cbiAgWW91ciBjaGFuZ2VzIGludm9sdmUgY3JlYXRpbmcgc3ltbGlua3MuIHBhdGNoLXBhY2thZ2UgZG9lcyBub3QgeWV0IHN1cHBvcnRcbiAgc3ltbGlua3MuXG4gIFxuICDvuI9QbGVhc2UgdXNlICR7Y2hhbGsuYm9sZChcIi0taW5jbHVkZVwiKX0gYW5kL29yICR7Y2hhbGsuYm9sZChcbiAgICAgICAgICBcIi0tZXhjbHVkZVwiLFxuICAgICAgICApfSB0byBuYXJyb3cgdGhlIHNjb3BlIG9mIHlvdXIgcGF0Y2ggaWZcbiAgdGhpcyB3YXMgdW5pbnRlbnRpb25hbC5cbmApXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBjb25zdCBvdXRQYXRoID0gXCIuL3BhdGNoLXBhY2thZ2UtZXJyb3IuanNvbi5nelwiXG4gICAgICAgIHdyaXRlRmlsZVN5bmMoXG4gICAgICAgICAgb3V0UGF0aCxcbiAgICAgICAgICBnemlwU3luYyhcbiAgICAgICAgICAgIEpTT04uc3RyaW5naWZ5KHtcbiAgICAgICAgICAgICAgZXJyb3I6IHsgbWVzc2FnZTogZS5tZXNzYWdlLCBzdGFjazogZS5zdGFjayB9LFxuICAgICAgICAgICAgICBwYXRjaDogZGlmZlJlc3VsdC5zdGRvdXQudG9TdHJpbmcoKSxcbiAgICAgICAgICAgIH0pLFxuICAgICAgICAgICksXG4gICAgICAgIClcbiAgICAgICAgY29uc29sZS5sb2coYFxu4puU77iPICR7Y2hhbGsucmVkLmJvbGQoXCJFUlJPUlwiKX1cbiAgICAgICAgXG4gIHBhdGNoLXBhY2thZ2Ugd2FzIHVuYWJsZSB0byByZWFkIHRoZSBwYXRjaC1maWxlIG1hZGUgYnkgZ2l0LiBUaGlzIHNob3VsZCBub3RcbiAgaGFwcGVuLlxuICBcbiAgQSBkaWFnbm9zdGljIGZpbGUgd2FzIHdyaXR0ZW4gdG9cbiAgXG4gICAgJHtvdXRQYXRofVxuICBcbiAgUGxlYXNlIGF0dGFjaCBpdCB0byBhIGdpdGh1YiBpc3N1ZVxuICBcbiAgICBodHRwczovL2dpdGh1Yi5jb20vZHMzMDAvcGF0Y2gtcGFja2FnZS9pc3N1ZXMvbmV3P3RpdGxlPU5ldytwYXRjaCtwYXJzZStmYWlsZWQmYm9keT1QbGVhc2UrYXR0YWNoK3RoZStkaWFnbm9zdGljK2ZpbGUrYnkrZHJhZ2dpbmcraXQraW50bytoZXJlK/CfmY9cbiAgXG4gIE5vdGUgdGhhdCB0aGlzIGRpYWdub3N0aWMgZmlsZSB3aWxsIGNvbnRhaW4gY29kZSBmcm9tIHRoZSBwYWNrYWdlIHlvdSB3ZXJlXG4gIGF0dGVtcHRpbmcgdG8gcGF0Y2guXG5cbmApXG4gICAgICB9XG4gICAgICBwcm9jZXNzLmV4aXQoMSlcbiAgICAgIHJldHVyblxuICAgIH1cblxuICAgIC8vIG1heWJlIGRlbGV0ZSBleGlzdGluZ1xuICAgIGlmIChtb2RlLnR5cGUgPT09IFwiYXBwZW5kXCIgJiYgIWlzUmViYXNpbmcgJiYgZXhpc3RpbmdQYXRjaGVzLmxlbmd0aCA9PT0gMSkge1xuICAgICAgLy8gaWYgd2UgYXJlIGFwcGVuZGluZyB0byBhbiBleGlzdGluZyBwYXRjaCB0aGF0IGRvZXNuJ3QgaGF2ZSBhIHNlcXVlbmNlIG51bWJlciBsZXQncyByZW5hbWUgaXRcbiAgICAgIGNvbnN0IHByZXZQYXRjaCA9IGV4aXN0aW5nUGF0Y2hlc1swXVxuICAgICAgaWYgKHByZXZQYXRjaC5zZXF1ZW5jZU51bWJlciA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGNvbnN0IG5ld0ZpbGVOYW1lID0gY3JlYXRlUGF0Y2hGaWxlTmFtZSh7XG4gICAgICAgICAgcGFja2FnZURldGFpbHMsXG4gICAgICAgICAgcGFja2FnZVZlcnNpb24sXG4gICAgICAgICAgc2VxdWVuY2VOdW1iZXI6IDEsXG4gICAgICAgICAgc2VxdWVuY2VOYW1lOiBwcmV2UGF0Y2guc2VxdWVuY2VOYW1lID8/IFwiaW5pdGlhbFwiLFxuICAgICAgICB9KVxuICAgICAgICBjb25zdCBvbGRQYXRoID0gam9pbihhcHBQYXRoLCBwYXRjaERpciwgcHJldlBhdGNoLnBhdGNoRmlsZW5hbWUpXG4gICAgICAgIGNvbnN0IG5ld1BhdGggPSBqb2luKGFwcFBhdGgsIHBhdGNoRGlyLCBuZXdGaWxlTmFtZSlcbiAgICAgICAgcmVuYW1lU3luYyhvbGRQYXRoLCBuZXdQYXRoKVxuICAgICAgICBwcmV2UGF0Y2guc2VxdWVuY2VOdW1iZXIgPSAxXG4gICAgICAgIHByZXZQYXRjaC5wYXRjaEZpbGVuYW1lID0gbmV3RmlsZU5hbWVcbiAgICAgICAgcHJldlBhdGNoLnNlcXVlbmNlTmFtZSA9IHByZXZQYXRjaC5zZXF1ZW5jZU5hbWUgPz8gXCJpbml0aWFsXCJcbiAgICAgIH1cbiAgICB9XG5cbiAgICBjb25zdCBsYXN0UGF0Y2ggPSBleGlzdGluZ1BhdGNoZXNbXG4gICAgICBzdGF0ZSA/IHN0YXRlLnBhdGNoZXMubGVuZ3RoIC0gMSA6IGV4aXN0aW5nUGF0Y2hlcy5sZW5ndGggLSAxXG4gICAgXSBhcyBQYXRjaGVkUGFja2FnZURldGFpbHMgfCB1bmRlZmluZWRcbiAgICBjb25zdCBzZXF1ZW5jZU5hbWUgPVxuICAgICAgbW9kZS50eXBlID09PSBcImFwcGVuZFwiID8gbW9kZS5uYW1lIDogbGFzdFBhdGNoPy5zZXF1ZW5jZU5hbWVcbiAgICBjb25zdCBzZXF1ZW5jZU51bWJlciA9XG4gICAgICBtb2RlLnR5cGUgPT09IFwiYXBwZW5kXCJcbiAgICAgICAgPyAobGFzdFBhdGNoPy5zZXF1ZW5jZU51bWJlciA/PyAwKSArIDFcbiAgICAgICAgOiBsYXN0UGF0Y2g/LnNlcXVlbmNlTnVtYmVyXG5cbiAgICBjb25zdCBwYXRjaEZpbGVOYW1lID0gY3JlYXRlUGF0Y2hGaWxlTmFtZSh7XG4gICAgICBwYWNrYWdlRGV0YWlscyxcbiAgICAgIHBhY2thZ2VWZXJzaW9uLFxuICAgICAgc2VxdWVuY2VOYW1lLFxuICAgICAgc2VxdWVuY2VOdW1iZXIsXG4gICAgfSlcblxuICAgIGNvbnN0IHBhdGNoUGF0aDogc3RyaW5nID0gam9pbihwYXRjaGVzRGlyLCBwYXRjaEZpbGVOYW1lKVxuICAgIGlmICghZXhpc3RzU3luYyhkaXJuYW1lKHBhdGNoUGF0aCkpKSB7XG4gICAgICAvLyBzY29wZWQgcGFja2FnZVxuICAgICAgbWtkaXJTeW5jKGRpcm5hbWUocGF0Y2hQYXRoKSlcbiAgICB9XG5cbiAgICAvLyBpZiB3ZSBhcmUgaW5zZXJ0aW5nIGEgbmV3IHBhdGNoIGludG8gYSBzZXF1ZW5jZSB3ZSBtb3N0IGxpa2VseSBuZWVkIHRvIHVwZGF0ZSB0aGUgc2VxdWVuY2UgbnVtYmVyc1xuICAgIGlmIChpc1JlYmFzaW5nICYmIG1vZGUudHlwZSA9PT0gXCJhcHBlbmRcIikge1xuICAgICAgY29uc3QgcGF0Y2hlc1RvTnVkZ2UgPSBleGlzdGluZ1BhdGNoZXMuc2xpY2Uoc3RhdGUhLnBhdGNoZXMubGVuZ3RoKVxuICAgICAgaWYgKHNlcXVlbmNlTnVtYmVyID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgdGhyb3cgbmV3IEVycm9yKFwic2VxdWVuY2VOdW1iZXIgaXMgdW5kZWZpbmVkIHdoaWxlIHJlYmFzaW5nXCIpXG4gICAgICB9XG4gICAgICBpZiAoXG4gICAgICAgIHBhdGNoZXNUb051ZGdlWzBdPy5zZXF1ZW5jZU51bWJlciAhPT0gdW5kZWZpbmVkICYmXG4gICAgICAgIHBhdGNoZXNUb051ZGdlWzBdLnNlcXVlbmNlTnVtYmVyIDw9IHNlcXVlbmNlTnVtYmVyXG4gICAgICApIHtcbiAgICAgICAgbGV0IG5leHQgPSBzZXF1ZW5jZU51bWJlciArIDFcbiAgICAgICAgZm9yIChjb25zdCBwIG9mIHBhdGNoZXNUb051ZGdlKSB7XG4gICAgICAgICAgY29uc3QgbmV3TmFtZSA9IGNyZWF0ZVBhdGNoRmlsZU5hbWUoe1xuICAgICAgICAgICAgcGFja2FnZURldGFpbHMsXG4gICAgICAgICAgICBwYWNrYWdlVmVyc2lvbixcbiAgICAgICAgICAgIHNlcXVlbmNlTmFtZTogcC5zZXF1ZW5jZU5hbWUsXG4gICAgICAgICAgICBzZXF1ZW5jZU51bWJlcjogbmV4dCsrLFxuICAgICAgICAgIH0pXG4gICAgICAgICAgY29uc29sZS5sb2coXG4gICAgICAgICAgICBcIlJlbmFtaW5nXCIsXG4gICAgICAgICAgICBjaGFsay5ib2xkKHAucGF0Y2hGaWxlbmFtZSksXG4gICAgICAgICAgICBcInRvXCIsXG4gICAgICAgICAgICBjaGFsay5ib2xkKG5ld05hbWUpLFxuICAgICAgICAgIClcbiAgICAgICAgICBjb25zdCBvbGRQYXRoID0gam9pbihhcHBQYXRoLCBwYXRjaERpciwgcC5wYXRjaEZpbGVuYW1lKVxuICAgICAgICAgIGNvbnN0IG5ld1BhdGggPSBqb2luKGFwcFBhdGgsIHBhdGNoRGlyLCBuZXdOYW1lKVxuICAgICAgICAgIHJlbmFtZVN5bmMob2xkUGF0aCwgbmV3UGF0aClcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHdyaXRlRmlsZVN5bmMocGF0Y2hQYXRoLCBkaWZmUmVzdWx0LnN0ZG91dClcbiAgICBjb25zb2xlLmxvZyhcbiAgICAgIGAke2NoYWxrLmdyZWVuKFwi4pyUXCIpfSBDcmVhdGVkIGZpbGUgJHtqb2luKHBhdGNoRGlyLCBwYXRjaEZpbGVOYW1lKX1cXG5gLFxuICAgIClcblxuICAgIGNvbnN0IHByZXZTdGF0ZTogUGF0Y2hTdGF0ZVtdID0gcGF0Y2hlc1RvQXBwbHlCZWZvcmVEaWZmaW5nLm1hcChcbiAgICAgIChwKTogUGF0Y2hTdGF0ZSA9PiAoe1xuICAgICAgICBwYXRjaEZpbGVuYW1lOiBwLnBhdGNoRmlsZW5hbWUsXG4gICAgICAgIGRpZEFwcGx5OiB0cnVlLFxuICAgICAgICBwYXRjaENvbnRlbnRIYXNoOiBoYXNoRmlsZShqb2luKGFwcFBhdGgsIHBhdGNoRGlyLCBwLnBhdGNoRmlsZW5hbWUpKSxcbiAgICAgIH0pLFxuICAgIClcbiAgICBjb25zdCBuZXh0U3RhdGU6IFBhdGNoU3RhdGVbXSA9IFtcbiAgICAgIC4uLnByZXZTdGF0ZSxcbiAgICAgIHtcbiAgICAgICAgcGF0Y2hGaWxlbmFtZTogcGF0Y2hGaWxlTmFtZSxcbiAgICAgICAgZGlkQXBwbHk6IHRydWUsXG4gICAgICAgIHBhdGNoQ29udGVudEhhc2g6IGhhc2hGaWxlKHBhdGNoUGF0aCksXG4gICAgICB9LFxuICAgIF1cblxuICAgIC8vIGlmIGFueSBwYXRjaGVzIGNvbWUgYWZ0ZXIgdGhpcyBvbmUgd2UganVzdCBtYWRlLCB3ZSBzaG91bGQgcmVhcHBseSB0aGVtXG4gICAgbGV0IGRpZEZhaWxXaGlsZUZpbmlzaGluZ1JlYmFzZSA9IGZhbHNlXG4gICAgaWYgKGlzUmViYXNpbmcpIHtcbiAgICAgIGNvbnN0IGN1cnJlbnRQYXRjaGVzID0gZ2V0R3JvdXBlZFBhdGNoZXMoam9pbihhcHBQYXRoLCBwYXRjaERpcikpXG4gICAgICAgIC5wYXRoU3BlY2lmaWVyVG9QYXRjaEZpbGVzW3BhY2thZ2VEZXRhaWxzLnBhdGhTcGVjaWZpZXJdXG5cbiAgICAgIGNvbnN0IHByZXZpb3VzbHlVbmFwcGxpZWRQYXRjaGVzID0gY3VycmVudFBhdGNoZXMuc2xpY2UobmV4dFN0YXRlLmxlbmd0aClcbiAgICAgIGlmIChwcmV2aW91c2x5VW5hcHBsaWVkUGF0Y2hlcy5sZW5ndGgpIHtcbiAgICAgICAgY29uc29sZS5sb2coYEZhc3QgZm9yd2FyZGluZy4uLmApXG4gICAgICAgIGZvciAoY29uc3QgcGF0Y2ggb2YgcHJldmlvdXNseVVuYXBwbGllZFBhdGNoZXMpIHtcbiAgICAgICAgICBjb25zdCBwYXRjaEZpbGVQYXRoID0gam9pbihhcHBQYXRoLCBwYXRjaERpciwgcGF0Y2gucGF0Y2hGaWxlbmFtZSlcbiAgICAgICAgICBpZiAoXG4gICAgICAgICAgICAhYXBwbHlQYXRjaCh7XG4gICAgICAgICAgICAgIHBhdGNoRGV0YWlsczogcGF0Y2gsXG4gICAgICAgICAgICAgIHBhdGNoRGlyLFxuICAgICAgICAgICAgICBwYXRjaEZpbGVQYXRoLFxuICAgICAgICAgICAgICByZXZlcnNlOiBmYWxzZSxcbiAgICAgICAgICAgICAgY3dkOiBwcm9jZXNzLmN3ZCgpLFxuICAgICAgICAgICAgICBiZXN0RWZmb3J0OiBmYWxzZSxcbiAgICAgICAgICAgIH0pXG4gICAgICAgICAgKSB7XG4gICAgICAgICAgICBkaWRGYWlsV2hpbGVGaW5pc2hpbmdSZWJhc2UgPSB0cnVlXG4gICAgICAgICAgICBsb2dQYXRjaFNlcXVlbmNlRXJyb3IoeyBwYXRjaERldGFpbHM6IHBhdGNoIH0pXG4gICAgICAgICAgICBuZXh0U3RhdGUucHVzaCh7XG4gICAgICAgICAgICAgIHBhdGNoRmlsZW5hbWU6IHBhdGNoLnBhdGNoRmlsZW5hbWUsXG4gICAgICAgICAgICAgIGRpZEFwcGx5OiBmYWxzZSxcbiAgICAgICAgICAgICAgcGF0Y2hDb250ZW50SGFzaDogaGFzaEZpbGUocGF0Y2hGaWxlUGF0aCksXG4gICAgICAgICAgICB9KVxuICAgICAgICAgICAgYnJlYWtcbiAgICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgICAgY29uc29sZS5sb2coYCAgJHtjaGFsay5ncmVlbihcIuKclFwiKX0gJHtwYXRjaC5wYXRjaEZpbGVuYW1lfWApXG4gICAgICAgICAgICBuZXh0U3RhdGUucHVzaCh7XG4gICAgICAgICAgICAgIHBhdGNoRmlsZW5hbWU6IHBhdGNoLnBhdGNoRmlsZW5hbWUsXG4gICAgICAgICAgICAgIGRpZEFwcGx5OiB0cnVlLFxuICAgICAgICAgICAgICBwYXRjaENvbnRlbnRIYXNoOiBoYXNoRmlsZShwYXRjaEZpbGVQYXRoKSxcbiAgICAgICAgICAgIH0pXG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgaWYgKGlzUmViYXNpbmcgfHwgbnVtUGF0Y2hlc0FmdGVyQ3JlYXRlID4gMSkge1xuICAgICAgc2F2ZVBhdGNoQXBwbGljYXRpb25TdGF0ZSh7XG4gICAgICAgIHBhY2thZ2VEZXRhaWxzLFxuICAgICAgICBwYXRjaGVzOiBuZXh0U3RhdGUsXG4gICAgICAgIGlzUmViYXNpbmc6IGRpZEZhaWxXaGlsZUZpbmlzaGluZ1JlYmFzZSxcbiAgICAgIH0pXG4gICAgfSBlbHNlIHtcbiAgICAgIGNsZWFyUGF0Y2hBcHBsaWNhdGlvblN0YXRlKHBhY2thZ2VEZXRhaWxzKVxuICAgIH1cblxuICAgIGlmIChjYW5DcmVhdGVJc3N1ZSkge1xuICAgICAgaWYgKGNyZWF0ZUlzc3VlKSB7XG4gICAgICAgIG9wZW5Jc3N1ZUNyZWF0aW9uTGluayh7XG4gICAgICAgICAgcGFja2FnZURldGFpbHMsXG4gICAgICAgICAgcGF0Y2hGaWxlQ29udGVudHM6IGRpZmZSZXN1bHQuc3Rkb3V0LnRvU3RyaW5nKCksXG4gICAgICAgICAgcGFja2FnZVZlcnNpb24sXG4gICAgICAgICAgcGF0Y2hQYXRoLFxuICAgICAgICB9KVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgbWF5YmVQcmludElzc3VlQ3JlYXRpb25Qcm9tcHQodmNzLCBwYWNrYWdlRGV0YWlscywgcGFja2FnZU1hbmFnZXIpXG4gICAgICB9XG4gICAgfVxuICB9IGNhdGNoIChlKSB7XG4gICAgY29uc29sZS5sb2coZSlcbiAgICB0aHJvdyBlXG4gIH0gZmluYWxseSB7XG4gICAgdG1wUmVwby5yZW1vdmVDYWxsYmFjaygpXG4gIH1cbn1cblxuZnVuY3Rpb24gY3JlYXRlUGF0Y2hGaWxlTmFtZSh7XG4gIHBhY2thZ2VEZXRhaWxzLFxuICBwYWNrYWdlVmVyc2lvbixcbiAgc2VxdWVuY2VOdW1iZXIsXG4gIHNlcXVlbmNlTmFtZSxcbn06IHtcbiAgcGFja2FnZURldGFpbHM6IFBhY2thZ2VEZXRhaWxzXG4gIHBhY2thZ2VWZXJzaW9uOiBzdHJpbmdcbiAgc2VxdWVuY2VOdW1iZXI/OiBudW1iZXJcbiAgc2VxdWVuY2VOYW1lPzogc3RyaW5nXG59KSB7XG4gIGNvbnN0IHBhY2thZ2VOYW1lcyA9IHBhY2thZ2VEZXRhaWxzLnBhY2thZ2VOYW1lc1xuICAgIC5tYXAoKG5hbWUpID0+IG5hbWUucmVwbGFjZSgvXFwvL2csIFwiK1wiKSlcbiAgICAuam9pbihcIisrXCIpXG5cbiAgY29uc3QgbmFtZUFuZFZlcnNpb24gPSBgJHtwYWNrYWdlTmFtZXN9KyR7cGFja2FnZVZlcnNpb259YFxuICBjb25zdCBudW0gPVxuICAgIHNlcXVlbmNlTnVtYmVyID09PSB1bmRlZmluZWRcbiAgICAgID8gXCJcIlxuICAgICAgOiBgKyR7c2VxdWVuY2VOdW1iZXIudG9TdHJpbmcoKS5wYWRTdGFydCgzLCBcIjBcIil9YFxuICBjb25zdCBuYW1lID0gIXNlcXVlbmNlTmFtZSA/IFwiXCIgOiBgKyR7c2VxdWVuY2VOYW1lfWBcblxuICByZXR1cm4gYCR7bmFtZUFuZFZlcnNpb259JHtudW19JHtuYW1lfS5wYXRjaGBcbn1cblxuZXhwb3J0IGZ1bmN0aW9uIGxvZ1BhdGNoU2VxdWVuY2VFcnJvcih7XG4gIHBhdGNoRGV0YWlscyxcbn06IHtcbiAgcGF0Y2hEZXRhaWxzOiBQYXRjaGVkUGFja2FnZURldGFpbHNcbn0pIHtcbiAgY29uc29sZS5sb2coYFxuJHtjaGFsay5yZWQuYm9sZChcIuKblCBFUlJPUlwiKX1cblxuRmFpbGVkIHRvIGFwcGx5IHBhdGNoIGZpbGUgJHtjaGFsay5ib2xkKHBhdGNoRGV0YWlscy5wYXRjaEZpbGVuYW1lKX0uXG5cbklmIHRoaXMgcGF0Y2ggZmlsZSBpcyBubyBsb25nZXIgdXNlZnVsLCBkZWxldGUgaXQgYW5kIHJ1blxuXG4gICR7Y2hhbGsuYm9sZChgcGF0Y2gtcGFja2FnZWApfVxuXG5UbyBwYXJ0aWFsbHkgYXBwbHkgdGhlIHBhdGNoIChpZiBwb3NzaWJsZSkgYW5kIG91dHB1dCBhIGxvZyBvZiBlcnJvcnMgdG8gZml4LCBydW5cblxuICAke2NoYWxrLmJvbGQoYHBhdGNoLXBhY2thZ2UgLS1wYXJ0aWFsYCl9XG5cbkFmdGVyIHdoaWNoIHlvdSBzaG91bGQgbWFrZSBhbnkgcmVxdWlyZWQgY2hhbmdlcyBpbnNpZGUgJHtcbiAgICBwYXRjaERldGFpbHMucGF0aFxuICB9LCBhbmQgZmluYWxseSBydW5cblxuICAke2NoYWxrLmJvbGQoYHBhdGNoLXBhY2thZ2UgJHtwYXRjaERldGFpbHMucGF0aFNwZWNpZmllcn1gKX1cblxudG8gdXBkYXRlIHRoZSBwYXRjaCBmaWxlLlxuYClcbn1cbiJdfQ==