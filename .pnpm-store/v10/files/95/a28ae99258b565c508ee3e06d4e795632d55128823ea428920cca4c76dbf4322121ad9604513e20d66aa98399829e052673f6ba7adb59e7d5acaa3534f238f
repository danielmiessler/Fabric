"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.applyPatch = exports.applyPatchesForPackage = exports.applyPatchesForApp = void 0;
const chalk_1 = __importDefault(require("chalk"));
const fs_1 = require("fs");
const fs_extra_1 = require("fs-extra");
const path_1 = require("path");
const semver_1 = __importDefault(require("semver"));
const hash_1 = require("./hash");
const makePatch_1 = require("./makePatch");
const packageIsDevDependency_1 = require("./packageIsDevDependency");
const apply_1 = require("./patch/apply");
const read_1 = require("./patch/read");
const reverse_1 = require("./patch/reverse");
const patchFs_1 = require("./patchFs");
const path_2 = require("./path");
const stateFile_1 = require("./stateFile");
class PatchApplicationError extends Error {
    constructor(msg) {
        super(msg);
    }
}
function getInstalledPackageVersion({ appPath, path, pathSpecifier, isDevOnly, patchFilename, }) {
    const packageDir = path_2.join(appPath, path);
    if (!fs_extra_1.existsSync(packageDir)) {
        if (process.env.NODE_ENV === "production" && isDevOnly) {
            return null;
        }
        let err = `${chalk_1.default.red("Error:")} Patch file found for package ${path_1.posix.basename(pathSpecifier)}` + ` which is not present at ${path_2.relative(".", packageDir)}`;
        if (!isDevOnly && process.env.NODE_ENV === "production") {
            err += `

  If this package is a dev dependency, rename the patch file to
  
    ${chalk_1.default.bold(patchFilename.replace(".patch", ".dev.patch"))}
`;
        }
        throw new PatchApplicationError(err);
    }
    const { version } = require(path_2.join(packageDir, "package.json"));
    // normalize version for `npm ci`
    const result = semver_1.default.valid(version);
    if (result === null) {
        throw new PatchApplicationError(`${chalk_1.default.red("Error:")} Version string '${version}' cannot be parsed from ${path_2.join(packageDir, "package.json")}`);
    }
    return result;
}
function logPatchApplication(patchDetails) {
    const sequenceString = patchDetails.sequenceNumber != null
        ? ` (${patchDetails.sequenceNumber}${patchDetails.sequenceName ? " " + patchDetails.sequenceName : ""})`
        : "";
    console.log(`${chalk_1.default.bold(patchDetails.pathSpecifier)}@${patchDetails.version}${sequenceString} ${chalk_1.default.green("✔")}`);
}
function applyPatchesForApp({ appPath, reverse, patchDir, shouldExitWithError, shouldExitWithWarning, bestEffort, }) {
    const patchesDirectory = path_2.join(appPath, patchDir);
    const groupedPatches = patchFs_1.getGroupedPatches(patchesDirectory);
    if (groupedPatches.numPatchFiles === 0) {
        console.log(chalk_1.default.blueBright("No patch files found"));
        return;
    }
    const errors = [];
    const warnings = [...groupedPatches.warnings];
    for (const patches of Object.values(groupedPatches.pathSpecifierToPatchFiles)) {
        applyPatchesForPackage({
            patches,
            appPath,
            patchDir,
            reverse,
            warnings,
            errors,
            bestEffort,
        });
    }
    for (const warning of warnings) {
        console.log(warning);
    }
    for (const error of errors) {
        console.log(error);
    }
    const problemsSummary = [];
    if (warnings.length) {
        problemsSummary.push(chalk_1.default.yellow(`${warnings.length} warning(s)`));
    }
    if (errors.length) {
        problemsSummary.push(chalk_1.default.red(`${errors.length} error(s)`));
    }
    if (problemsSummary.length) {
        console.log("---");
        console.log("patch-package finished with", problemsSummary.join(", ") + ".");
    }
    if (errors.length && shouldExitWithError) {
        process.exit(1);
    }
    if (warnings.length && shouldExitWithWarning) {
        process.exit(1);
    }
    process.exit(0);
}
exports.applyPatchesForApp = applyPatchesForApp;
function applyPatchesForPackage({ patches, appPath, patchDir, reverse, warnings, errors, bestEffort, }) {
    const pathSpecifier = patches[0].pathSpecifier;
    const state = patches.length > 1 ? stateFile_1.getPatchApplicationState(patches[0]) : null;
    const unappliedPatches = patches.slice(0);
    const appliedPatches = [];
    // if there are multiple patches to apply, we can't rely on the reverse-patch-dry-run behavior to make this operation
    // idempotent, so instead we need to check the state file to see whether we have already applied any of the patches
    // todo: once this is battle tested we might want to use the same approach for single patches as well, but it's not biggie since the dry run thing is fast
    if (unappliedPatches && state) {
        for (let i = 0; i < state.patches.length; i++) {
            const patchThatWasApplied = state.patches[i];
            if (!patchThatWasApplied.didApply) {
                break;
            }
            const patchToApply = unappliedPatches[0];
            const currentPatchHash = hash_1.hashFile(path_2.join(appPath, patchDir, patchToApply.patchFilename));
            if (patchThatWasApplied.patchContentHash === currentPatchHash) {
                // this patch was applied we can skip it
                appliedPatches.push(unappliedPatches.shift());
            }
            else {
                console.log(chalk_1.default.red("Error:"), `The patches for ${chalk_1.default.bold(pathSpecifier)} have changed.`, `You should reinstall your node_modules folder to make sure the package is up to date`);
                process.exit(1);
            }
        }
    }
    if (reverse && state) {
        // if we are reversing the patches we need to make the unappliedPatches array
        // be the reversed version of the appliedPatches array.
        // The applied patches array should then be empty because it is used differently
        // when outputting the state file.
        unappliedPatches.length = 0;
        unappliedPatches.push(...appliedPatches);
        unappliedPatches.reverse();
        appliedPatches.length = 0;
    }
    if (appliedPatches.length) {
        // some patches have already been applied
        appliedPatches.forEach(logPatchApplication);
    }
    if (!unappliedPatches.length) {
        return;
    }
    let failedPatch = null;
    packageLoop: for (const patchDetails of unappliedPatches) {
        try {
            const { name, version, path, isDevOnly, patchFilename } = patchDetails;
            const installedPackageVersion = getInstalledPackageVersion({
                appPath,
                path,
                pathSpecifier,
                isDevOnly: isDevOnly ||
                    // check for direct-dependents in prod
                    (process.env.NODE_ENV === "production" &&
                        packageIsDevDependency_1.packageIsDevDependency({
                            appPath,
                            patchDetails,
                        })),
                patchFilename,
            });
            if (!installedPackageVersion) {
                // it's ok we're in production mode and this is a dev only package
                console.log(`Skipping dev-only ${chalk_1.default.bold(pathSpecifier)}@${version} ${chalk_1.default.blue("✔")}`);
                continue;
            }
            if (applyPatch({
                patchFilePath: path_2.join(appPath, patchDir, patchFilename),
                reverse,
                patchDetails,
                patchDir,
                cwd: process.cwd(),
                bestEffort,
            })) {
                appliedPatches.push(patchDetails);
                // yay patch was applied successfully
                // print warning if version mismatch
                if (installedPackageVersion !== version) {
                    warnings.push(createVersionMismatchWarning({
                        packageName: name,
                        actualVersion: installedPackageVersion,
                        originalVersion: version,
                        pathSpecifier,
                        path,
                    }));
                }
                logPatchApplication(patchDetails);
            }
            else if (patches.length > 1) {
                makePatch_1.logPatchSequenceError({ patchDetails });
                // in case the package has multiple patches, we need to break out of this inner loop
                // because we don't want to apply more patches on top of the broken state
                failedPatch = patchDetails;
                break packageLoop;
            }
            else if (installedPackageVersion === version) {
                // completely failed to apply patch
                // TODO: propagate useful error messages from patch application
                errors.push(createBrokenPatchFileError({
                    packageName: name,
                    patchFilename,
                    pathSpecifier,
                    path,
                }));
                break packageLoop;
            }
            else {
                errors.push(createPatchApplicationFailureError({
                    packageName: name,
                    actualVersion: installedPackageVersion,
                    originalVersion: version,
                    patchFilename,
                    path,
                    pathSpecifier,
                }));
                // in case the package has multiple patches, we need to break out of this inner loop
                // because we don't want to apply more patches on top of the broken state
                break packageLoop;
            }
        }
        catch (error) {
            if (error instanceof PatchApplicationError) {
                errors.push(error.message);
            }
            else {
                errors.push(createUnexpectedError({
                    filename: patchDetails.patchFilename,
                    error: error,
                }));
            }
            // in case the package has multiple patches, we need to break out of this inner loop
            // because we don't want to apply more patches on top of the broken state
            break packageLoop;
        }
    }
    if (patches.length > 1) {
        if (reverse) {
            if (!state) {
                throw new Error("unexpected state: no state file found while reversing");
            }
            // if we removed all the patches that were previously applied we can delete the state file
            if (appliedPatches.length === patches.length) {
                stateFile_1.clearPatchApplicationState(patches[0]);
            }
            else {
                // We failed while reversing patches and some are still in the applied state.
                // We need to update the state file to reflect that.
                // appliedPatches is currently the patches that were successfully reversed, in the order they were reversed
                // So we need to find the index of the last reversed patch in the original patches array
                // and then remove all the patches after that. Sorry for the confusing code.
                const lastReversedPatchIndex = patches.indexOf(appliedPatches[appliedPatches.length - 1]);
                if (lastReversedPatchIndex === -1) {
                    throw new Error("unexpected state: failed to find last reversed patch in original patches array");
                }
                stateFile_1.savePatchApplicationState({
                    packageDetails: patches[0],
                    patches: patches.slice(0, lastReversedPatchIndex).map((patch) => ({
                        didApply: true,
                        patchContentHash: hash_1.hashFile(path_2.join(appPath, patchDir, patch.patchFilename)),
                        patchFilename: patch.patchFilename,
                    })),
                    isRebasing: false,
                });
            }
        }
        else {
            const nextState = appliedPatches.map((patch) => ({
                didApply: true,
                patchContentHash: hash_1.hashFile(path_2.join(appPath, patchDir, patch.patchFilename)),
                patchFilename: patch.patchFilename,
            }));
            if (failedPatch) {
                nextState.push({
                    didApply: false,
                    patchContentHash: hash_1.hashFile(path_2.join(appPath, patchDir, failedPatch.patchFilename)),
                    patchFilename: failedPatch.patchFilename,
                });
            }
            stateFile_1.savePatchApplicationState({
                packageDetails: patches[0],
                patches: nextState,
                isRebasing: !!failedPatch,
            });
        }
        if (failedPatch) {
            process.exit(1);
        }
    }
}
exports.applyPatchesForPackage = applyPatchesForPackage;
function applyPatch({ patchFilePath, reverse, patchDetails, patchDir, cwd, bestEffort, }) {
    const patch = read_1.readPatch({
        patchFilePath,
        patchDetails,
        patchDir,
    });
    const forward = reverse ? reverse_1.reversePatch(patch) : patch;
    try {
        if (!bestEffort) {
            apply_1.executeEffects(forward, { dryRun: true, cwd, bestEffort: false });
        }
        const errors = bestEffort ? [] : undefined;
        apply_1.executeEffects(forward, { dryRun: false, cwd, bestEffort, errors });
        if (errors === null || errors === void 0 ? void 0 : errors.length) {
            console.log("Saving errors to", chalk_1.default.cyan.bold("./patch-package-errors.log"));
            fs_1.writeFileSync("patch-package-errors.log", errors.join("\n\n"));
            process.exit(0);
        }
    }
    catch (e) {
        try {
            const backward = reverse ? patch : reverse_1.reversePatch(patch);
            apply_1.executeEffects(backward, {
                dryRun: true,
                cwd,
                bestEffort: false,
            });
        }
        catch (e) {
            return false;
        }
    }
    return true;
}
exports.applyPatch = applyPatch;
function createVersionMismatchWarning({ packageName, actualVersion, originalVersion, pathSpecifier, path, }) {
    return `
${chalk_1.default.yellow("Warning:")} patch-package detected a patch file version mismatch

  Don't worry! This is probably fine. The patch was still applied
  successfully. Here's the deets:

  Patch file created for

    ${packageName}@${chalk_1.default.bold(originalVersion)}

  applied to

    ${packageName}@${chalk_1.default.bold(actualVersion)}
  
  At path
  
    ${path}

  This warning is just to give you a heads-up. There is a small chance of
  breakage even though the patch was applied successfully. Make sure the package
  still behaves like you expect (you wrote tests, right?) and then run

    ${chalk_1.default.bold(`patch-package ${pathSpecifier}`)}

  to update the version in the patch file name and make this warning go away.
`;
}
function createBrokenPatchFileError({ packageName, patchFilename, path, pathSpecifier, }) {
    return `
${chalk_1.default.red.bold("**ERROR**")} ${chalk_1.default.red(`Failed to apply patch for package ${chalk_1.default.bold(packageName)} at path`)}
  
    ${path}

  This error was caused because patch-package cannot apply the following patch file:

    patches/${patchFilename}

  Try removing node_modules and trying again. If that doesn't work, maybe there was
  an accidental change made to the patch file? Try recreating it by manually
  editing the appropriate files and running:
  
    patch-package ${pathSpecifier}
  
  If that doesn't work, then it's a bug in patch-package, so please submit a bug
  report. Thanks!

    https://github.com/ds300/patch-package/issues
    
`;
}
function createPatchApplicationFailureError({ packageName, actualVersion, originalVersion, patchFilename, path, pathSpecifier, }) {
    return `
${chalk_1.default.red.bold("**ERROR**")} ${chalk_1.default.red(`Failed to apply patch for package ${chalk_1.default.bold(packageName)} at path`)}
  
    ${path}

  This error was caused because ${chalk_1.default.bold(packageName)} has changed since you
  made the patch file for it. This introduced conflicts with your patch,
  just like a merge conflict in Git when separate incompatible changes are
  made to the same piece of code.

  Maybe this means your patch file is no longer necessary, in which case
  hooray! Just delete it!

  Otherwise, you need to generate a new patch file.

  To generate a new one, just repeat the steps you made to generate the first
  one.

  i.e. manually make the appropriate file changes, then run 

    patch-package ${pathSpecifier}

  Info:
    Patch file: patches/${patchFilename}
    Patch was made for version: ${chalk_1.default.green.bold(originalVersion)}
    Installed version: ${chalk_1.default.red.bold(actualVersion)}
`;
}
function createUnexpectedError({ filename, error, }) {
    return `
${chalk_1.default.red.bold("**ERROR**")} ${chalk_1.default.red(`Failed to apply patch file ${chalk_1.default.bold(filename)}`)}
  
${error.stack}

  `;
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYXBwbHlQYXRjaGVzLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL2FwcGx5UGF0Y2hlcy50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7Ozs7QUFBQSxrREFBeUI7QUFDekIsMkJBQWtDO0FBQ2xDLHVDQUFxQztBQUNyQywrQkFBNEI7QUFDNUIsb0RBQTJCO0FBQzNCLGlDQUFpQztBQUNqQywyQ0FBbUQ7QUFFbkQscUVBQWlFO0FBQ2pFLHlDQUE4QztBQUM5Qyx1Q0FBd0M7QUFDeEMsNkNBQThDO0FBQzlDLHVDQUE2QztBQUM3QyxpQ0FBdUM7QUFDdkMsMkNBS29CO0FBRXBCLE1BQU0scUJBQXNCLFNBQVEsS0FBSztJQUN2QyxZQUFZLEdBQVc7UUFDckIsS0FBSyxDQUFDLEdBQUcsQ0FBQyxDQUFBO0lBQ1osQ0FBQztDQUNGO0FBRUQsU0FBUywwQkFBMEIsQ0FBQyxFQUNsQyxPQUFPLEVBQ1AsSUFBSSxFQUNKLGFBQWEsRUFDYixTQUFTLEVBQ1QsYUFBYSxHQU9kO0lBQ0MsTUFBTSxVQUFVLEdBQUcsV0FBSSxDQUFDLE9BQU8sRUFBRSxJQUFJLENBQUMsQ0FBQTtJQUN0QyxJQUFJLENBQUMscUJBQVUsQ0FBQyxVQUFVLENBQUMsRUFBRTtRQUMzQixJQUFJLE9BQU8sQ0FBQyxHQUFHLENBQUMsUUFBUSxLQUFLLFlBQVksSUFBSSxTQUFTLEVBQUU7WUFDdEQsT0FBTyxJQUFJLENBQUE7U0FDWjtRQUVELElBQUksR0FBRyxHQUNMLEdBQUcsZUFBSyxDQUFDLEdBQUcsQ0FBQyxRQUFRLENBQUMsaUNBQWlDLFlBQUssQ0FBQyxRQUFRLENBQ25FLGFBQWEsQ0FDZCxFQUFFLEdBQUcsNEJBQTRCLGVBQVEsQ0FBQyxHQUFHLEVBQUUsVUFBVSxDQUFDLEVBQUUsQ0FBQTtRQUUvRCxJQUFJLENBQUMsU0FBUyxJQUFJLE9BQU8sQ0FBQyxHQUFHLENBQUMsUUFBUSxLQUFLLFlBQVksRUFBRTtZQUN2RCxHQUFHLElBQUk7Ozs7TUFJUCxlQUFLLENB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LFdBQVcsRUFDWCxhQUFhLEVBQ2IsZUFBZSxFQUNmLGFBQWEsRUFDYixJQUFJLEVBQ0osYUFBYSxHQVFkO0lBQ0MsT0FBTztFQUNQLGVBQUssQ0FBQyxHQUFHLENBQUMsSUFBSSxDQUFDLFdBQVcsQ0FBQyxJQUFJLGVBQUssQ0FBQyxHQUFHLENBQ3RDLHFDQUFxQyxlQUFLLENBQUMsSUFBSSxDQUFDLFdBQVcsQ0FBQyxVQUFVLENBQ3ZFOztNQUVHLElBQUk7O2tDQUV3QixlQUFLLENBQUMsSUFBSSxDQUFDLFdBQVcsQ0FBQzs7Ozs7Ozs7Ozs7Ozs7O29CQWVyQyxhQUFhOzs7MEJBR1AsYUFBYTtrQ0FDTCxlQUFLLENBQUMsS0FBSyxDQUFDLElBQUksQ0FBQyxlQUFlLENBQUM7eUJBQzFDLGVBQUssQ0FBQyxHQUFHLENBQUMsSUFBSSxDQUFDLGFBQWEsQ0FBQztDQUNyRCxDQUFBO0FBQ0QsQ0FBQztBQUVELFNBQVMscUJBQXFCLENBQUMsRUFDN0IsUUFBUSxFQUNSLEtBQUssR0FJTjtJQUNDLE9BQU87RUFDUCxlQUFLLENBQUMsR0FBRyxDQUFDLElBQUksQ0FBQyxXQUFXLENBQUMsSUFBSSxlQUFLLENBQUMsR0FBRyxDQUN0Qyw4QkFBOEIsZUFBSyxDQUFDLElBQUksQ0FBQyxRQUFRLENBQUMsRUFBRSxDQUNyRDs7RUFFRCxLQUFLLENBQUMsS0FBSzs7R0FFVixDQUFBO0FBQ0gsQ0FBQyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBjaGFsayBmcm9tIFwiY2hhbGtcIlxuaW1wb3J0IHsgd3JpdGVGaWxlU3luYyB9IGZyb20gXCJmc1wiXG5pbXBvcnQgeyBleGlzdHNTeW5jIH0gZnJvbSBcImZzLWV4dHJhXCJcbmltcG9ydCB7IHBvc2l4IH0gZnJvbSBcInBhdGhcIlxuaW1wb3J0IHNlbXZlciBmcm9tIFwic2VtdmVyXCJcbmltcG9ydCB7IGhhc2hGaWxlIH0gZnJvbSBcIi4vaGFzaFwiXG5pbXBvcnQgeyBsb2dQYXRjaFNlcXVlbmNlRXJyb3IgfSBmcm9tIFwiLi9tYWtlUGF0Y2hcIlxuaW1wb3J0IHsgUGFja2FnZURldGFpbHMsIFBhdGNoZWRQYWNrYWdlRGV0YWlscyB9IGZyb20gXCIuL1BhY2thZ2VEZXRhaWxzXCJcbmltcG9ydCB7IHBhY2thZ2VJc0RldkRlcGVuZGVuY3kgfSBmcm9tIFwiLi9wYWNrYWdlSXNEZXZEZXBlbmRlbmN5XCJcbmltcG9ydCB7IGV4ZWN1dGVFZmZlY3RzIH0gZnJvbSBcIi4vcGF0Y2gvYXBwbHlcIlxuaW1wb3J0IHsgcmVhZFBhdGNoIH0gZnJvbSBcIi4vcGF0Y2gvcmVhZFwiXG5pbXBvcnQgeyByZXZlcnNlUGF0Y2ggfSBmcm9tIFwiLi9wYXRjaC9yZXZlcnNlXCJcbmltcG9ydCB7IGdldEdyb3VwZWRQYXRjaGVzIH0gZnJvbSBcIi4vcGF0Y2hGc1wiXG5pbXBvcnQgeyBqb2luLCByZWxhdGl2ZSB9IGZyb20gXCIuL3BhdGhcIlxuaW1wb3J0IHtcbiAgY2xlYXJQYXRjaEFwcGxpY2F0aW9uU3RhdGUsXG4gIGdldFBhdGNoQXBwbGljYXRpb25TdGF0ZSxcbiAgUGF0Y2hTdGF0ZSxcbiAgc2F2ZVBhdGNoQXBwbGljYXRpb25TdGF0ZSxcbn0gZnJvbSBcIi4vc3RhdGVGaWxlXCJcblxuY2xhc3MgUGF0Y2hBcHBsaWNhdGlvbkVycm9yIGV4dGVuZHMgRXJyb3Ige1xuICBjb25zdHJ1Y3Rvcihtc2c6IHN0cmluZykge1xuICAgIHN1cGVyKG1zZylcbiAgfVxufVxuXG5mdW5jdGlvbiBnZXRJbnN0YWxsZWRQYWNrYWdlVmVyc2lvbih7XG4gIGFwcFBhdGgsXG4gIHBhdGgsXG4gIHBhdGhTcGVjaWZpZXIsXG4gIGlzRGV2T25seSxcbiAgcGF0Y2hGaWxlbmFtZSxcbn06IHtcbiAgYXBwUGF0aDogc3RyaW5nXG4gIHBhdGg6IHN0cmluZ1xuICBwYXRoU3BlY2lmaWVyOiBzdHJpbmdcbiAgaXNEZXZPbmx5OiBib29sZWFuXG4gIHBhdGNoRmlsZW5hbWU6IHN0cmluZ1xufSk6IG51bGwgfCBzdHJpbmcge1xuICBjb25zdCBwYWNrYWdlRGlyID0gam9pbihhcHBQYXRoLCBwYXRoKVxuICBpZiAoIWV4aXN0c1N5bmMocGFja2FnZURpcikpIHtcbiAgICBpZiAocHJvY2Vzcy5lbnYuTk9ERV9FTlYgPT09IFwicHJvZHVjdGlvblwiICYmIGlzRGV2T25seSkge1xuICAgICAgcmV0dXJuIG51bGxcbiAgICB9XG5cbiAgICBsZXQgZXJyID1cbiAgICAgIGAke2NoYWxrLnJlZChcIkVycm9yOlwiKX0gUGF0Y2ggZmlsZSBmb3VuZCBmb3IgcGFja2FnZSAke3Bvc2l4LmJhc2VuYW1lKFxuICAgICAgICBwYXRoU3BlY2lmaWVyLFxuICAgICAgKX1gICsgYCB3aGljaCBpcyBub3QgcHJlc2VudCBhdCAke3JlbGF0aXZlKFwiLlwiLCBwYWNrYWdlRGlyKX1gXG5cbiAgICBpZiAoIWlzRGV2T25seSAmJiBwcm9jZXNzLmVudi5OT0RFX0VOViA9PT0gXCJwcm9kdWN0aW9uXCIpIHtcbiAgICAgIGVyciArPSBgXG5cbiAgSWYgdGhpcyBwYWNrYWdlIGlzIGEgZGV2IGRlcGVuZGVuY3ksIHJlbmFtZSB0aGUgcGF0Y2ggZmlsZSB0b1xuICBcbiAgICAke2NoYWxrLmJvbGQocGF0Y2hGaWxlbmFtZS5yZXBsYWNlKFwiLnBhdGNoXCIsIFwiLmRldi5wYXRjaFwiKSl9XG5gXG4gICAgfVxuICAgIHRocm93IG5ldyBQYXRjaEFwcGxpY2F0aW9uRXJyb3IoZXJyKVxuICB9XG5cbiAgY29uc3QgeyB2ZXJzaW9uIH0gPSByZXF1aXJlKGpvaW4ocGFja2FnZURpciwgXCJwYWNrYWdlLmpzb25cIikpXG4gIC8vIG5vcm1hbGl6ZSB2ZXJzaW9uIGZvciBgbnBtIGNpYFxuICBjb25zdCByZXN1bHQgPSBzZW12ZXIudmFsaWQodmVyc2lvbilcbiAgaWYgKHJlc3VsdCA9PT0gbnVsbCkge1xuICAgIHRocm93IG5ldyBQYXRjaEFwcGxpY2F0aW9uRXJyb3IoXG4gICAgICBgJHtjaGFsay5yZWQoXG4gICAgICAgIFwiRXJyb3I6XCIsXG4gICAgICApfSBWZXJzaW9uIHN0cmluZyAnJHt2ZXJzaW9ufScgY2Fubm90IGJlIHBhcnNlZCBmcm9tICR7am9pbihcbiAgICAgICAgcGFja2FnZURpcixcbiAgICAgICAgXCJwYWNrYWdlLmpzb25cIixcbiAgICAgICl9YCxcbiAgICApXG4gIH1cblxuICByZXR1cm4gcmVzdWx0IGFzIHN0cmluZ1xufVxuXG5mdW5jdGlvbiBsb2dQYXRjaEFwcGxpY2F0aW9uKHBhdGNoRGV0YWlsczogUGF0Y2hlZFBhY2thZ2VEZXRhaWxzKSB7XG4gIGNvbnN0IHNlcXVlbmNlU3RyaW5nID1cbiAgICBwYXRjaERldGFpbHMuc2VxdWVuY2VOdW1iZXIgIT0gbnVsbFxuICAgICAgPyBgICgke3BhdGNoRGV0YWlscy5zZXF1ZW5jZU51bWJlcn0ke1xuICAgICAgICAgIHBhdGNoRGV0YWlscy5zZXF1ZW5jZU5hbWUgPyBcIiBcIiArIHBhdGNoRGV0YWlscy5zZXF1ZW5jZU5hbWUgOiBcIlwiXG4gICAgICAgIH0pYFxuICAgICAgOiBcIlwiXG4gIGNvbnNvbGUubG9nKFxuICAgIGAke2NoYWxrLmJvbGQocGF0Y2hEZXRhaWxzLnBhdGhTcGVjaWZpZXIpfUAke1xuICAgICAgcGF0Y2hEZXRhaWxzLnZlcnNpb25cbiAgICB9JHtzZXF1ZW5jZVN0cmluZ30gJHtjaGFsay5ncmVlbihcIuKclFwiKX1gLFxuICApXG59XG5cbmV4cG9ydCBmdW5jdGlvbiBhcHBseVBhdGNoZXNGb3JBcHAoe1xuICBhcHBQYXRoLFxuICByZXZlcnNlLFxuICBwYXRjaERpcixcbiAgc2hvdWxkRXhpdFdpdGhFcnJvcixcbiAgc2hvdWxkRXhpdFdpdGhXYXJuaW5nLFxuICBiZXN0RWZmb3J0LFxufToge1xuICBhcHBQYXRoOiBzdHJpbmdcbiAgcmV2ZXJzZTogYm9vbGVhblxuICBwYXRjaERpcjogc3RyaW5nXG4gIHNob3VsZEV4aXRXaXRoRXJyb3I6IGJvb2xlYW5cbiAgc2hvdWxkRXhpdFdpdGhXYXJuaW5nOiBib29sZWFuXG4gIGJlc3RFZmZvcnQ6IGJvb2xlYW5cbn0pOiB2b2lkIHtcbiAgY29uc3QgcGF0Y2hlc0RpcmVjdG9yeSA9IGpvaW4oYXBwUGF0aCwgcGF0Y2hEaXIpXG4gIGNvbnN0IGdyb3VwZWRQYXRjaGVzID0gZ2V0R3JvdXBlZFBhdGNoZXMocGF0Y2hlc0RpcmVjdG9yeSlcblxuICBpZiAoZ3JvdXBlZFBhdGNoZXMubnVtUGF0Y2hGaWxlcyA9PT0gMCkge1xuICAgIGNvbnNvbGUubG9nKGNoYWxrLmJsdWVCcmlnaHQoXCJObyBwYXRjaCBmaWxlcyBmb3VuZFwiKSlcbiAgICByZXR1cm5cbiAgfVxuXG4gIGNvbnN0IGVycm9yczogc3RyaW5nW10gPSBbXVxuICBjb25zdCB3YXJuaW5nczogc3RyaW5nW10gPSBbLi4uZ3JvdXBlZFBhdGNoZXMud2FybmluZ3NdXG5cbiAgZm9yIChjb25zdCBwYXRjaGVzIG9mIE9iamVjdC52YWx1ZXMoXG4gICAgZ3JvdXBlZFBhdGNoZXMucGF0aFNwZWNpZmllclRvUGF0Y2hGaWxlcyxcbiAgKSkge1xuICAgIGFwcGx5UGF0Y2hlc0ZvclBhY2thZ2Uoe1xuICAgICAgcGF0Y2hlcyxcbiAgICAgIGFwcFBhdGgsXG4gICAgICBwYXRjaERpcixcbiAgICAgIHJldmVyc2UsXG4gICAgICB3YXJuaW5ncyxcbiAgICAgIGVycm9ycyxcbiAgICAgIGJlc3RFZmZvcnQsXG4gICAgfSlcbiAgfVxuXG4gIGZvciAoY29uc3Qgd2FybmluZyBvZiB3YXJuaW5ncykge1xuICAgIGNvbnNvbGUubG9nKHdhcm5pbmcpXG4gIH1cbiAgZm9yIChjb25zdCBlcnJvciBvZiBlcnJvcnMpIHtcbiAgICBjb25zb2xlLmxvZyhlcnJvcilcbiAgfVxuXG4gIGNvbnN0IHByb2JsZW1zU3VtbWFyeSA9IFtdXG4gIGlmICh3YXJuaW5ncy5sZW5ndGgpIHtcbiAgICBwcm9ibGVtc1N1bW1hcnkucHVzaChjaGFsay55ZWxsb3coYCR7d2FybmluZ3MubGVuZ3RofSB3YXJuaW5nKHMpYCkpXG4gIH1cbiAgaWYgKGVycm9ycy5sZW5ndGgpIHtcbiAgICBwcm9ibGVtc1N1bW1hcnkucHVzaChjaGFsay5yZWQoYCR7ZXJyb3JzLmxlbmd0aH0gZXJyb3IocylgKSlcbiAgfVxuXG4gIGlmIChwcm9ibGVtc1N1bW1hcnkubGVuZ3RoKSB7XG4gICAgY29uc29sZS5sb2coXCItLS1cIilcbiAgICBjb25zb2xlLmxvZyhcInBhdGNoLXBhY2thZ2UgZmluaXNoZWQgd2l0aFwiLCBwcm9ibGVtc1N1bW1hcnkuam9pbihcIiwgXCIpICsgXCIuXCIpXG4gIH1cblxuICBpZiAoZXJyb3JzLmxlbmd0aCAmJiBzaG91bGRFeGl0V2l0aEVycm9yKSB7XG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cblxuICBpZiAod2FybmluZ3MubGVuZ3RoICYmIHNob3VsZEV4aXRXaXRoV2FybmluZykge1xuICAgIHByb2Nlc3MuZXhpdCgxKVxuICB9XG5cbiAgcHJvY2Vzcy5leGl0KDApXG59XG5cbmV4cG9ydCBmdW5jdGlvbiBhcHBseVBhdGNoZXNGb3JQYWNrYWdlKHtcbiAgcGF0Y2hlcyxcbiAgYXBwUGF0aCxcbiAgcGF0Y2hEaXIsXG4gIHJldmVyc2UsXG4gIHdhcm5pbmdzLFxuICBlcnJvcnMsXG4gIGJlc3RFZmZvcnQsXG59OiB7XG4gIHBhdGNoZXM6IFBhdGNoZWRQYWNrYWdlRGV0YWlsc1tdXG4gIGFwcFBhdGg6IHN0cmluZ1xuICBwYXRjaERpcjogc3RyaW5nXG4gIHJldmVyc2U6IGJvb2xlYW5cbiAgd2FybmluZ3M6IHN0cmluZ1tdXG4gIGVycm9yczogc3RyaW5nW11cbiAgYmVzdEVmZm9ydDogYm9vbGVhblxufSkge1xuICBjb25zdCBwYXRoU3BlY2lmaWVyID0gcGF0Y2hlc1swXS5wYXRoU3BlY2lmaWVyXG4gIGNvbnN0IHN0YXRlID0gcGF0Y2hlcy5sZW5ndGggPiAxID8gZ2V0UGF0Y2hBcHBsaWNhdGlvblN0YXRlKHBhdGNoZXNbMF0pIDogbnVsbFxuICBjb25zdCB1bmFwcGxpZWRQYXRjaGVzID0gcGF0Y2hlcy5zbGljZSgwKVxuICBjb25zdCBhcHBsaWVkUGF0Y2hlczogUGF0Y2hlZFBhY2thZ2VEZXRhaWxzW10gPSBbXVxuICAvLyBpZiB0aGVyZSBhcmUgbXVsdGlwbGUgcGF0Y2hlcyB0byBhcHBseSwgd2UgY2FuJ3QgcmVseSBvbiB0aGUgcmV2ZXJzZS1wYXRjaC1kcnktcnVuIGJlaGF2aW9yIHRvIG1ha2UgdGhpcyBvcGVyYXRpb25cbiAgLy8gaWRlbXBvdGVudCwgc28gaW5zdGVhZCB3ZSBuZWVkIHRvIGNoZWNrIHRoZSBzdGF0ZSBmaWxlIHRvIHNlZSB3aGV0aGVyIHdlIGhhdmUgYWxyZWFkeSBhcHBsaWVkIGFueSBvZiB0aGUgcGF0Y2hlc1xuICAvLyB0b2RvOiBvbmNlIHRoaXMgaXMgYmF0dGxlIHRlc3RlZCB3ZSBtaWdodCB3YW50IHRvIHVzZSB0aGUgc2FtZSBhcHByb2FjaCBmb3Igc2luZ2xlIHBhdGNoZXMgYXMgd2VsbCwgYnV0IGl0J3Mgbm90IGJpZ2dpZSBzaW5jZSB0aGUgZHJ5IHJ1biB0aGluZyBpcyBmYXN0XG4gIGlmICh1bmFwcGxpZWRQYXRjaGVzICYmIHN0YXRlKSB7XG4gICAgZm9yIChsZXQgaSA9IDA7IGkgPCBzdGF0ZS5wYXRjaGVzLmxlbmd0aDsgaSsrKSB7XG4gICAgICBjb25zdCBwYXRjaFRoYXRXYXNBcHBsaWVkID0gc3RhdGUucGF0Y2hlc1tpXVxuICAgICAgaWYgKCFwYXRjaFRoYXRXYXNBcHBsaWVkLmRpZEFwcGx5KSB7XG4gICAgICAgIGJyZWFrXG4gICAgICB9XG4gICAgICBjb25zdCBwYXRjaFRvQXBwbHkgPSB1bmFwcGxpZWRQYXRjaGVzWzBdXG4gICAgICBjb25zdCBjdXJyZW50UGF0Y2hIYXNoID0gaGFzaEZpbGUoXG4gICAgICAgIGpvaW4oYXBwUGF0aCwgcGF0Y2hEaXIsIHBhdGNoVG9BcHBseS5wYXRjaEZpbGVuYW1lKSxcbiAgICAgIClcbiAgICAgIGlmIChwYXRjaFRoYXRXYXNBcHBsaWVkLnBhdGNoQ29udGVudEhhc2ggPT09IGN1cnJlbnRQYXRjaEhhc2gpIHtcbiAgICAgICAgLy8gdGhpcyBwYXRjaCB3YXMgYXBwbGllZCB3ZSBjYW4gc2tpcCBpdFxuICAgICAgICBhcHBsaWVkUGF0Y2hlcy5wdXNoKHVuYXBwbGllZFBhdGNoZXMuc2hpZnQoKSEpXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgICBjaGFsay5yZWQoXCJFcnJvcjpcIiksXG4gICAgICAgICAgYFRoZSBwYXRjaGVzIGZvciAke2NoYWxrLmJvbGQocGF0aFNwZWNpZmllcil9IGhhdmUgY2hhbmdlZC5gLFxuICAgICAgICAgIGBZb3Ugc2hvdWxkIHJlaW5zdGFsbCB5b3VyIG5vZGVfbW9kdWxlcyBmb2xkZXIgdG8gbWFrZSBzdXJlIHRoZSBwYWNrYWdlIGlzIHVwIHRvIGRhdGVgLFxuICAgICAgICApXG4gICAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIGlmIChyZXZlcnNlICYmIHN0YXRlKSB7XG4gICAgLy8gaWYgd2UgYXJlIHJldmVyc2luZyB0aGUgcGF0Y2hlcyB3ZSBuZWVkIHRvIG1ha2UgdGhlIHVuYXBwbGllZFBhdGNoZXMgYXJyYXlcbiAgICAvLyBiZSB0aGUgcmV2ZXJzZWQgdmVyc2lvbiBvZiB0aGUgYXBwbGllZFBhdGNoZXMgYXJyYXkuXG4gICAgLy8gVGhlIGFwcGxpZWQgcGF0Y2hlcyBhcnJheSBzaG91bGQgdGhlbiBiZSBlbXB0eSBiZWNhdXNlIGl0IGlzIHVzZWQgZGlmZmVyZW50bHlcbiAgICAvLyB3aGVuIG91dHB1dHRpbmcgdGhlIHN0YXRlIGZpbGUuXG4gICAgdW5hcHBsaWVkUGF0Y2hlcy5sZW5ndGggPSAwXG4gICAgdW5hcHBsaWVkUGF0Y2hlcy5wdXNoKC4uLmFwcGxpZWRQYXRjaGVzKVxuICAgIHVuYXBwbGllZFBhdGNoZXMucmV2ZXJzZSgpXG4gICAgYXBwbGllZFBhdGNoZXMubGVuZ3RoID0gMFxuICB9XG4gIGlmIChhcHBsaWVkUGF0Y2hlcy5sZW5ndGgpIHtcbiAgICAvLyBzb21lIHBhdGNoZXMgaGF2ZSBhbHJlYWR5IGJlZW4gYXBwbGllZFxuICAgIGFwcGxpZWRQYXRjaGVzLmZvckVhY2gobG9nUGF0Y2hBcHBsaWNhdGlvbilcbiAgfVxuICBpZiAoIXVuYXBwbGllZFBhdGNoZXMubGVuZ3RoKSB7XG4gICAgcmV0dXJuXG4gIH1cbiAgbGV0IGZhaWxlZFBhdGNoOiBQYXRjaGVkUGFja2FnZURldGFpbHMgfCBudWxsID0gbnVsbFxuICBwYWNrYWdlTG9vcDogZm9yIChjb25zdCBwYXRjaERldGFpbHMgb2YgdW5hcHBsaWVkUGF0Y2hlcykge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCB7IG5hbWUsIHZlcnNpb24sIHBhdGgsIGlzRGV2T25seSwgcGF0Y2hGaWxlbmFtZSB9ID0gcGF0Y2hEZXRhaWxzXG5cbiAgICAgIGNvbnN0IGluc3RhbGxlZFBhY2thZ2VWZXJzaW9uID0gZ2V0SW5zdGFsbGVkUGFja2FnZVZlcnNpb24oe1xuICAgICAgICBhcHBQYXRoLFxuICAgICAgICBwYXRoLFxuICAgICAgICBwYXRoU3BlY2lmaWVyLFxuICAgICAgICBpc0Rldk9ubHk6XG4gICAgICAgICAgaXNEZXZPbmx5IHx8XG4gICAgICAgICAgLy8gY2hlY2sgZm9yIGRpcmVjdC1kZXBlbmRlbnRzIGluIHByb2RcbiAgICAgICAgICAocHJvY2Vzcy5lbnYuTk9ERV9FTlYgPT09IFwicHJvZHVjdGlvblwiICYmXG4gICAgICAgICAgICBwYWNrYWdlSXNEZXZEZXBlbmRlbmN5KHtcbiAgICAgICAgICAgICAgYXBwUGF0aCxcbiAgICAgICAgICAgICAgcGF0Y2hEZXRhaWxzLFxuICAgICAgICAgICAgfSkpLFxuICAgICAgICBwYXRjaEZpbGVuYW1lLFxuICAgICAgfSlcbiAgICAgIGlmICghaW5zdGFsbGVkUGFja2FnZVZlcnNpb24pIHtcbiAgICAgICAgLy8gaXQncyBvayB3ZSdyZSBpbiBwcm9kdWN0aW9uIG1vZGUgYW5kIHRoaXMgaXMgYSBkZXYgb25seSBwYWNrYWdlXG4gICAgICAgIGNvbnNvbGUubG9nKFxuICAgICAgICAgIGBTa2lwcGluZyBkZXYtb25seSAke2NoYWxrLmJvbGQoXG4gICAgICAgICAgICBwYXRoU3BlY2lmaWVyLFxuICAgICAgICAgICl9QCR7dmVyc2lvbn0gJHtjaGFsay5ibHVlKFwi4pyUXCIpfWAsXG4gICAgICAgIClcbiAgICAgICAgY29udGludWVcbiAgICAgIH1cblxuICAgICAgaWYgKFxuICAgICAgICBhcHBseVBhdGNoKHtcbiAgICAgICAgICBwYXRjaEZpbGVQYXRoOiBqb2luKGFwcFBhdGgsIHBhdGNoRGlyLCBwYXRjaEZpbGVuYW1lKSBhcyBzdHJpbmcsXG4gICAgICAgICAgcmV2ZXJzZSxcbiAgICAgICAgICBwYXRjaERldGFpbHMsXG4gICAgICAgICAgcGF0Y2hEaXIsXG4gICAgICAgICAgY3dkOiBwcm9jZXNzLmN3ZCgpLFxuICAgICAgICAgIGJlc3RFZmZvcnQsXG4gICAgICAgIH0pXG4gICAgICApIHtcbiAgICAgICAgYXBwbGllZFBhdGNoZXMucHVzaChwYXRjaERldGFpbHMpXG4gICAgICAgIC8vIHlheSBwYXRjaCB3YXMgYXBwbGllZCBzdWNjZXNzZnVsbHlcbiAgICAgICAgLy8gcHJpbnQgd2FybmluZyBpZiB2ZXJzaW9uIG1pc21hdGNoXG4gICAgICAgIGlmIChpbnN0YWxsZWRQYWNrYWdlVmVyc2lvbiAhPT0gdmVyc2lvbikge1xuICAgICAgICAgIHdhcm5pbmdzLnB1c2goXG4gICAgICAgICAgICBjcmVhdGVWZXJzaW9uTWlzbWF0Y2hXYXJuaW5nKHtcbiAgICAgICAgICAgICAgcGFja2FnZU5hbWU6IG5hbWUsXG4gICAgICAgICAgICAgIGFjdHVhbFZlcnNpb246IGluc3RhbGxlZFBhY2thZ2VWZXJzaW9uLFxuICAgICAgICAgICAgICBvcmlnaW5hbFZlcnNpb246IHZlcnNpb24sXG4gICAgICAgICAgICAgIHBhdGhTcGVjaWZpZXIsXG4gICAgICAgICAgICAgIHBhdGgsXG4gICAgICAgICAgICB9KSxcbiAgICAgICAgICApXG4gICAgICAgIH1cbiAgICAgICAgbG9nUGF0Y2hBcHBsaWNhdGlvbihwYXRjaERldGFpbHMpXG4gICAgICB9IGVsc2UgaWYgKHBhdGNoZXMubGVuZ3RoID4gMSkge1xuICAgICAgICBsb2dQYXRjaFNlcXVlbmNlRXJyb3IoeyBwYXRjaERldGFpbHMgfSlcbiAgICAgICAgLy8gaW4gY2FzZSB0aGUgcGFja2FnZSBoYXMgbXVsdGlwbGUgcGF0Y2hlcywgd2UgbmVlZCB0byBicmVhayBvdXQgb2YgdGhpcyBpbm5lciBsb29wXG4gICAgICAgIC8vIGJlY2F1c2Ugd2UgZG9uJ3Qgd2FudCB0byBhcHBseSBtb3JlIHBhdGNoZXMgb24gdG9wIG9mIHRoZSBicm9rZW4gc3RhdGVcbiAgICAgICAgZmFpbGVkUGF0Y2ggPSBwYXRjaERldGFpbHNcbiAgICAgICAgYnJlYWsgcGFja2FnZUxvb3BcbiAgICAgIH0gZWxzZSBpZiAoaW5zdGFsbGVkUGFja2FnZVZlcnNpb24gPT09IHZlcnNpb24pIHtcbiAgICAgICAgLy8gY29tcGxldGVseSBmYWlsZWQgdG8gYXBwbHkgcGF0Y2hcbiAgICAgICAgLy8gVE9ETzogcHJvcGFnYXRlIHVzZWZ1bCBlcnJvciBtZXNzYWdlcyBmcm9tIHBhdGNoIGFwcGxpY2F0aW9uXG4gICAgICAgIGVycm9ycy5wdXNoKFxuICAgICAgICAgIGNyZWF0ZUJyb2tlblBhdGNoRmlsZUVycm9yKHtcbiAgICAgICAgICAgIHBhY2thZ2VOYW1lOiBuYW1lLFxuICAgICAgICAgICAgcGF0Y2hGaWxlbmFtZSxcbiAgICAgICAgICAgIHBhdGhTcGVjaWZpZXIsXG4gICAgICAgICAgICBwYXRoLFxuICAgICAgICAgIH0pLFxuICAgICAgICApXG4gICAgICAgIGJyZWFrIHBhY2thZ2VMb29wXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBlcnJvcnMucHVzaChcbiAgICAgICAgICBjcmVhdGVQYXRjaEFwcGxpY2F0aW9uRmFpbHVyZUVycm9yKHtcbiAgICAgICAgICAgIHBhY2thZ2VOYW1lOiBuYW1lLFxuICAgICAgICAgICAgYWN0dWFsVmVyc2lvbjogaW5zdGFsbGVkUGFja2FnZVZlcnNpb24sXG4gICAgICAgICAgICBvcmlnaW5hbFZlcnNpb246IHZlcnNpb24sXG4gICAgICAgICAgICBwYXRjaEZpbGVuYW1lLFxuICAgICAgICAgICAgcGF0aCxcbiAgICAgICAgICAgIHBhdGhTcGVjaWZpZXIsXG4gICAgICAgICAgfSksXG4gICAgICAgIClcbiAgICAgICAgLy8gaW4gY2FzZSB0aGUgcGFja2FnZSBoYXMgbXVsdGlwbGUgcGF0Y2hlcywgd2UgbmVlZCB0byBicmVhayBvdXQgb2YgdGhpcyBpbm5lciBsb29wXG4gICAgICAgIC8vIGJlY2F1c2Ugd2UgZG9uJ3Qgd2FudCB0byBhcHBseSBtb3JlIHBhdGNoZXMgb24gdG9wIG9mIHRoZSBicm9rZW4gc3RhdGVcbiAgICAgICAgYnJlYWsgcGFja2FnZUxvb3BcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnJvcikge1xuICAgICAgaWYgKGVycm9yIGluc3RhbmNlb2YgUGF0Y2hBcHBsaWNhdGlvbkVycm9yKSB7XG4gICAgICAgIGVycm9ycy5wdXNoKGVycm9yLm1lc3NhZ2UpXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBlcnJvcnMucHVzaChcbiAgICAgICAgICBjcmVhdGVVbmV4cGVjdGVkRXJyb3Ioe1xuICAgICAgICAgICAgZmlsZW5hbWU6IHBhdGNoRGV0YWlscy5wYXRjaEZpbGVuYW1lLFxuICAgICAgICAgICAgZXJyb3I6IGVycm9yIGFzIEVycm9yLFxuICAgICAgICAgIH0pLFxuICAgICAgICApXG4gICAgICB9XG4gICAgICAvLyBpbiBjYXNlIHRoZSBwYWNrYWdlIGhhcyBtdWx0aXBsZSBwYXRjaGVzLCB3ZSBuZWVkIHRvIGJyZWFrIG91dCBvZiB0aGlzIGlubmVyIGxvb3BcbiAgICAgIC8vIGJlY2F1c2Ugd2UgZG9uJ3Qgd2FudCB0byBhcHBseSBtb3JlIHBhdGNoZXMgb24gdG9wIG9mIHRoZSBicm9rZW4gc3RhdGVcbiAgICAgIGJyZWFrIHBhY2thZ2VMb29wXG4gICAgfVxuICB9XG5cbiAgaWYgKHBhdGNoZXMubGVuZ3RoID4gMSkge1xuICAgIGlmIChyZXZlcnNlKSB7XG4gICAgICBpZiAoIXN0YXRlKSB7XG4gICAgICAgIHRocm93IG5ldyBFcnJvcihcInVuZXhwZWN0ZWQgc3RhdGU6IG5vIHN0YXRlIGZpbGUgZm91bmQgd2hpbGUgcmV2ZXJzaW5nXCIpXG4gICAgICB9XG4gICAgICAvLyBpZiB3ZSByZW1vdmVkIGFsbCB0aGUgcGF0Y2hlcyB0aGF0IHdlcmUgcHJldmlvdXNseSBhcHBsaWVkIHdlIGNhbiBkZWxldGUgdGhlIHN0YXRlIGZpbGVcbiAgICAgIGlmIChhcHBsaWVkUGF0Y2hlcy5sZW5ndGggPT09IHBhdGNoZXMubGVuZ3RoKSB7XG4gICAgICAgIGNsZWFyUGF0Y2hBcHBsaWNhdGlvblN0YXRlKHBhdGNoZXNbMF0pXG4gICAgICB9IGVsc2Uge1xuICAgICAgICAvLyBXZSBmYWlsZWQgd2hpbGUgcmV2ZXJzaW5nIHBhdGNoZXMgYW5kIHNvbWUgYXJlIHN0aWxsIGluIHRoZSBhcHBsaWVkIHN0YXRlLlxuICAgICAgICAvLyBXZSBuZWVkIHRvIHVwZGF0ZSB0aGUgc3RhdGUgZmlsZSB0byByZWZsZWN0IHRoYXQuXG4gICAgICAgIC8vIGFwcGxpZWRQYXRjaGVzIGlzIGN1cnJlbnRseSB0aGUgcGF0Y2hlcyB0aGF0IHdlcmUgc3VjY2Vzc2Z1bGx5IHJldmVyc2VkLCBpbiB0aGUgb3JkZXIgdGhleSB3ZXJlIHJldmVyc2VkXG4gICAgICAgIC8vIFNvIHdlIG5lZWQgdG8gZmluZCB0aGUgaW5kZXggb2YgdGhlIGxhc3QgcmV2ZXJzZWQgcGF0Y2ggaW4gdGhlIG9yaWdpbmFsIHBhdGNoZXMgYXJyYXlcbiAgICAgICAgLy8gYW5kIHRoZW4gcmVtb3ZlIGFsbCB0aGUgcGF0Y2hlcyBhZnRlciB0aGF0LiBTb3JyeSBmb3IgdGhlIGNvbmZ1c2luZyBjb2RlLlxuICAgICAgICBjb25zdCBsYXN0UmV2ZXJzZWRQYXRjaEluZGV4ID0gcGF0Y2hlcy5pbmRleE9mKFxuICAgICAgICAgIGFwcGxpZWRQYXRjaGVzW2FwcGxpZWRQYXRjaGVzLmxlbmd0aCAtIDFdLFxuICAgICAgICApXG4gICAgICAgIGlmIChsYXN0UmV2ZXJzZWRQYXRjaEluZGV4ID09PSAtMSkge1xuICAgICAgICAgIHRocm93IG5ldyBFcnJvcihcbiAgICAgICAgICAgIFwidW5leHBlY3RlZCBzdGF0ZTogZmFpbGVkIHRvIGZpbmQgbGFzdCByZXZlcnNlZCBwYXRjaCBpbiBvcmlnaW5hbCBwYXRjaGVzIGFycmF5XCIsXG4gICAgICAgICAgKVxuICAgICAgICB9XG5cbiAgICAgICAgc2F2ZVBhdGNoQXBwbGljYXRpb25TdGF0ZSh7XG4gICAgICAgICAgcGFja2FnZURldGFpbHM6IHBhdGNoZXNbMF0sXG4gICAgICAgICAgcGF0Y2hlczogcGF0Y2hlcy5zbGljZSgwLCBsYXN0UmV2ZXJzZWRQYXRjaEluZGV4KS5tYXAoKHBhdGNoKSA9PiAoe1xuICAgICAgICAgICAgZGlkQXBwbHk6IHRydWUsXG4gICAgICAgICAgICBwYXRjaENvbnRlbnRIYXNoOiBoYXNoRmlsZShcbiAgICAgICAgICAgICAgam9pbihhcHBQYXRoLCBwYXRjaERpciwgcGF0Y2gucGF0Y2hGaWxlbmFtZSksXG4gICAgICAgICAgICApLFxuICAgICAgICAgICAgcGF0Y2hGaWxlbmFtZTogcGF0Y2gucGF0Y2hGaWxlbmFtZSxcbiAgICAgICAgICB9KSksXG4gICAgICAgICAgaXNSZWJhc2luZzogZmFsc2UsXG4gICAgICAgIH0pXG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIGNvbnN0IG5leHRTdGF0ZSA9IGFwcGxpZWRQYXRjaGVzLm1hcChcbiAgICAgICAgKHBhdGNoKTogUGF0Y2hTdGF0ZSA9PiAoe1xuICAgICAgICAgIGRpZEFwcGx5OiB0cnVlLFxuICAgICAgICAgIHBhdGNoQ29udGVudEhhc2g6IGhhc2hGaWxlKFxuICAgICAgICAgICAgam9pbihhcHBQYXRoLCBwYXRjaERpciwgcGF0Y2gucGF0Y2hGaWxlbmFtZSksXG4gICAgICAgICAgKSxcbiAgICAgICAgICBwYXRjaEZpbGVuYW1lOiBwYXRjaC5wYXRjaEZpbGVuYW1lLFxuICAgICAgICB9KSxcbiAgICAgIClcblxuICAgICAgaWYgKGZhaWxlZFBhdGNoKSB7XG4gICAgICAgIG5leHRTdGF0ZS5wdXNoKHtcbiAgICAgICAgICBkaWRBcHBseTogZmFsc2UsXG4gICAgICAgICAgcGF0Y2hDb250ZW50SGFzaDogaGFzaEZpbGUoXG4gICAgICAgICAgICBqb2luKGFwcFBhdGgsIHBhdGNoRGlyLCBmYWlsZWRQYXRjaC5wYXRjaEZpbGVuYW1lKSxcbiAgICAgICAgICApLFxuICAgICAgICAgIHBhdGNoRmlsZW5hbWU6IGZhaWxlZFBhdGNoLnBhdGNoRmlsZW5hbWUsXG4gICAgICAgIH0pXG4gICAgICB9XG4gICAgICBzYXZlUGF0Y2hBcHBsaWNhdGlvblN0YXRlKHtcbiAgICAgICAgcGFja2FnZURldGFpbHM6IHBhdGNoZXNbMF0sXG4gICAgICAgIHBhdGNoZXM6IG5leHRTdGF0ZSxcbiAgICAgICAgaXNSZWJhc2luZzogISFmYWlsZWRQYXRjaCxcbiAgICAgIH0pXG4gICAgfVxuICAgIGlmIChmYWlsZWRQYXRjaCkge1xuICAgICAgcHJvY2Vzcy5leGl0KDEpXG4gICAgfVxuICB9XG59XG5cbmV4cG9ydCBmdW5jdGlvbiBhcHBseVBhdGNoKHtcbiAgcGF0Y2hGaWxlUGF0aCxcbiAgcmV2ZXJzZSxcbiAgcGF0Y2hEZXRhaWxzLFxuICBwYXRjaERpcixcbiAgY3dkLFxuICBiZXN0RWZmb3J0LFxufToge1xuICBwYXRjaEZpbGVQYXRoOiBzdHJpbmdcbiAgcmV2ZXJzZTogYm9vbGVhblxuICBwYXRjaERldGFpbHM6IFBhY2thZ2VEZXRhaWxzXG4gIHBhdGNoRGlyOiBzdHJpbmdcbiAgY3dkOiBzdHJpbmdcbiAgYmVzdEVmZm9ydDogYm9vbGVhblxufSk6IGJvb2xlYW4ge1xuICBjb25zdCBwYXRjaCA9IHJlYWRQYXRjaCh7XG4gICAgcGF0Y2hGaWxlUGF0aCxcbiAgICBwYXRjaERldGFpbHMsXG4gICAgcGF0Y2hEaXIsXG4gIH0pXG5cbiAgY29uc3QgZm9yd2FyZCA9IHJldmVyc2UgPyByZXZlcnNlUGF0Y2gocGF0Y2gpIDogcGF0Y2hcbiAgdHJ5IHtcbiAgICBpZiAoIWJlc3RFZmZvcnQpIHtcbiAgICAgIGV4ZWN1dGVFZmZlY3RzKGZvcndhcmQsIHsgZHJ5UnVuOiB0cnVlLCBjd2QsIGJlc3RFZmZvcnQ6IGZhbHNlIH0pXG4gICAgfVxuICAgIGNvbnN0IGVycm9yczogc3RyaW5nW10gfCB1bmRlZmluZWQgPSBiZXN0RWZmb3J0ID8gW10gOiB1bmRlZmluZWRcbiAgICBleGVjdXRlRWZmZWN0cyhmb3J3YXJkLCB7IGRyeVJ1bjogZmFsc2UsIGN3ZCwgYmVzdEVmZm9ydCwgZXJyb3JzIH0pXG4gICAgaWYgKGVycm9ycz8ubGVuZ3RoKSB7XG4gICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgXCJTYXZpbmcgZXJyb3JzIHRvXCIsXG4gICAgICAgIGNoYWxrLmN5YW4uYm9sZChcIi4vcGF0Y2gtcGFja2FnZS1lcnJvcnMubG9nXCIpLFxuICAgICAgKVxuICAgICAgd3JpdGVGaWxlU3luYyhcInBhdGNoLXBhY2thZ2UtZXJyb3JzLmxvZ1wiLCBlcnJvcnMuam9pbihcIlxcblxcblwiKSlcbiAgICAgIHByb2Nlc3MuZXhpdCgwKVxuICAgIH1cbiAgfSBjYXRjaCAoZSkge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCBiYWNrd2FyZCA9IHJldmVyc2UgPyBwYXRjaCA6IHJldmVyc2VQYXRjaChwYXRjaClcbiAgICAgIGV4ZWN1dGVFZmZlY3RzKGJhY2t3YXJkLCB7XG4gICAgICAgIGRyeVJ1bjogdHJ1ZSxcbiAgICAgICAgY3dkLFxuICAgICAgICBiZXN0RWZmb3J0OiBmYWxzZSxcbiAgICAgIH0pXG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgcmV0dXJuIGZhbHNlXG4gICAgfVxuICB9XG5cbiAgcmV0dXJuIHRydWVcbn1cblxuZnVuY3Rpb24gY3JlYXRlVmVyc2lvbk1pc21hdGNoV2FybmluZyh7XG4gIHBhY2thZ2VOYW1lLFxuICBhY3R1YWxWZXJzaW9uLFxuICBvcmlnaW5hbFZlcnNpb24sXG4gIHBhdGhTcGVjaWZpZXIsXG4gIHBhdGgsXG59OiB7XG4gIHBhY2thZ2VOYW1lOiBzdHJpbmdcbiAgYWN0dWFsVmVyc2lvbjogc3RyaW5nXG4gIG9yaWdpbmFsVmVyc2lvbjogc3RyaW5nXG4gIHBhdGhTcGVjaWZpZXI6IHN0cmluZ1xuICBwYXRoOiBzdHJpbmdcbn0pIHtcbiAgcmV0dXJuIGBcbiR7Y2hhbGsueWVsbG93KFwiV2FybmluZzpcIil9IHBhdGNoLXBhY2thZ2UgZGV0ZWN0ZWQgYSBwYXRjaCBmaWxlIHZlcnNpb24gbWlzbWF0Y2hcblxuICBEb24ndCB3b3JyeSEgVGhpcyBpcyBwcm9iYWJseSBmaW5lLiBUaGUgcGF0Y2ggd2FzIHN0aWxsIGFwcGxpZWRcbiAgc3VjY2Vzc2Z1bGx5LiBIZXJlJ3MgdGhlIGRlZXRzOlxuXG4gIFBhdGNoIGZpbGUgY3JlYXRlZCBmb3JcblxuICAgICR7cGFja2FnZU5hbWV9QCR7Y2hhbGsuYm9sZChvcmlnaW5hbFZlcnNpb24pfVxuXG4gIGFwcGxpZWQgdG9cblxuICAgICR7cGFja2FnZU5hbWV9QCR7Y2hhbGsuYm9sZChhY3R1YWxWZXJzaW9uKX1cbiAgXG4gIEF0IHBhdGhcbiAgXG4gICAgJHtwYXRofVxuXG4gIFRoaXMgd2FybmluZyBpcyBqdXN0IHRvIGdpdmUgeW91IGEgaGVhZHMtdXAuIFRoZXJlIGlzIGEgc21hbGwgY2hhbmNlIG9mXG4gIGJyZWFrYWdlIGV2ZW4gdGhvdWdoIHRoZSBwYXRjaCB3YXMgYXBwbGllZCBzdWNjZXNzZnVsbHkuIE1ha2Ugc3VyZSB0aGUgcGFja2FnZVxuICBzdGlsbCBiZWhhdmVzIGxpa2UgeW91IGV4cGVjdCAoeW91IHdyb3RlIHRlc3RzLCByaWdodD8pIGFuZCB0aGVuIHJ1blxuXG4gICAgJHtjaGFsay5ib2xkKGBwYXRjaC1wYWNrYWdlICR7cGF0aFNwZWNpZmllcn1gKX1cblxuICB0byB1cGRhdGUgdGhlIHZlcnNpb24gaW4gdGhlIHBhdGNoIGZpbGUgbmFtZSBhbmQgbWFrZSB0aGlzIHdhcm5pbmcgZ28gYXdheS5cbmBcbn1cblxuZnVuY3Rpb24gY3JlYXRlQnJva2VuUGF0Y2hGaWxlRXJyb3Ioe1xuICBwYWNrYWdlTmFtZSxcbiAgcGF0Y2hGaWxlbmFtZSxcbiAgcGF0aCxcbiAgcGF0aFNwZWNpZmllcixcbn06IHtcbiAgcGFja2FnZU5hbWU6IHN0cmluZ1xuICBwYXRjaEZpbGVuYW1lOiBzdHJpbmdcbiAgcGF0aDogc3RyaW5nXG4gIHBhdGhTcGVjaWZpZXI6IHN0cmluZ1xufSkge1xuICByZXR1cm4gYFxuJHtjaGFsay5yZWQuYm9sZChcIioqRVJST1IqKlwiKX0gJHtjaGFsay5yZWQoXG4gICAgYEZhaWxlZCB0byBhcHBseSBwYXRjaCBmb3IgcGFja2FnZSAke2NoYWxrLmJvbGQocGFja2FnZU5hbWUpfSBhdCBwYXRoYCxcbiAgKX1cbiAgXG4gICAgJHtwYXRofVxuXG4gIFRoaXMgZXJyb3Igd2FzIGNhdXNlZCBiZWNhdXNlIHBhdGNoLXBhY2thZ2UgY2Fubm90IGFwcGx5IHRoZSBmb2xsb3dpbmcgcGF0Y2ggZmlsZTpcblxuICAgIHBhdGNoZXMvJHtwYXRjaEZpbGVuYW1lfVxuXG4gIFRyeSByZW1vdmluZyBub2RlX21vZHVsZXMgYW5kIHRyeWluZyBhZ2Fpbi4gSWYgdGhhdCBkb2Vzbid0IHdvcmssIG1heWJlIHRoZXJlIHdhc1xuICBhbiBhY2NpZGVudGFsIGNoYW5nZSBtYWRlIHRvIHRoZSBwYXRjaCBmaWxlPyBUcnkgcmVjcmVhdGluZyBpdCBieSBtYW51YWxseVxuICBlZGl0aW5nIHRoZSBhcHByb3ByaWF0ZSBmaWxlcyBhbmQgcnVubmluZzpcbiAgXG4gICAgcGF0Y2gtcGFja2FnZSAke3BhdGhTcGVjaWZpZXJ9XG4gIFxuICBJZiB0aGF0IGRvZXNuJ3Qgd29yaywgdGhlbiBpdCdzIGEgYnVnIGluIHBhdGNoLXBhY2thZ2UsIHNvIHBsZWFzZSBzdWJtaXQgYSBidWdcbiAgcmVwb3J0LiBUaGFua3MhXG5cbiAgICBodHRwczovL2dpdGh1Yi5jb20vZHMzMDAvcGF0Y2gtcGFja2FnZS9pc3N1ZXNcbiAgICBcbmBcbn1cblxuZnVuY3Rpb24gY3JlYXRlUGF0Y2hBcHBsaWNhdGlvbkZhaWx1cmVFcnJvcih7XG4gIHBhY2thZ2VOYW1lLFxuICBhY3R1YWxWZXJzaW9uLFxuICBvcmlnaW5hbFZlcnNpb24sXG4gIHBhdGNoRmlsZW5hbWUsXG4gIHBhdGgsXG4gIHBhdGhTcGVjaWZpZXIsXG59OiB7XG4gIHBhY2thZ2VOYW1lOiBzdHJpbmdcbiAgYWN0dWFsVmVyc2lvbjogc3RyaW5nXG4gIG9yaWdpbmFsVmVyc2lvbjogc3RyaW5nXG4gIHBhdGNoRmlsZW5hbWU6IHN0cmluZ1xuICBwYXRoOiBzdHJpbmdcbiAgcGF0aFNwZWNpZmllcjogc3RyaW5nXG59KSB7XG4gIHJldHVybiBgXG4ke2NoYWxrLnJlZC5ib2xkKFwiKipFUlJPUioqXCIpfSAke2NoYWxrLnJlZChcbiAgICBgRmFpbGVkIHRvIGFwcGx5IHBhdGNoIGZvciBwYWNrYWdlICR7Y2hhbGsuYm9sZChwYWNrYWdlTmFtZSl9IGF0IHBhdGhgLFxuICApfVxuICBcbiAgICAke3BhdGh9XG5cbiAgVGhpcyBlcnJvciB3YXMgY2F1c2VkIGJlY2F1c2UgJHtjaGFsay5ib2xkKHBhY2thZ2VOYW1lKX0gaGFzIGNoYW5nZWQgc2luY2UgeW91XG4gIG1hZGUgdGhlIHBhdGNoIGZpbGUgZm9yIGl0LiBUaGlzIGludHJvZHVjZWQgY29uZmxpY3RzIHdpdGggeW91ciBwYXRjaCxcbiAganVzdCBsaWtlIGEgbWVyZ2UgY29uZmxpY3QgaW4gR2l0IHdoZW4gc2VwYXJhdGUgaW5jb21wYXRpYmxlIGNoYW5nZXMgYXJlXG4gIG1hZGUgdG8gdGhlIHNhbWUgcGllY2Ugb2YgY29kZS5cblxuICBNYXliZSB0aGlzIG1lYW5zIHlvdXIgcGF0Y2ggZmlsZSBpcyBubyBsb25nZXIgbmVjZXNzYXJ5LCBpbiB3aGljaCBjYXNlXG4gIGhvb3JheSEgSnVzdCBkZWxldGUgaXQhXG5cbiAgT3RoZXJ3aXNlLCB5b3UgbmVlZCB0byBnZW5lcmF0ZSBhIG5ldyBwYXRjaCBmaWxlLlxuXG4gIFRvIGdlbmVyYXRlIGEgbmV3IG9uZSwganVzdCByZXBlYXQgdGhlIHN0ZXBzIHlvdSBtYWRlIHRvIGdlbmVyYXRlIHRoZSBmaXJzdFxuICBvbmUuXG5cbiAgaS5lLiBtYW51YWxseSBtYWtlIHRoZSBhcHByb3ByaWF0ZSBmaWxlIGNoYW5nZXMsIHRoZW4gcnVuIFxuXG4gICAgcGF0Y2gtcGFja2FnZSAke3BhdGhTcGVjaWZpZXJ9XG5cbiAgSW5mbzpcbiAgICBQYXRjaCBmaWxlOiBwYXRjaGVzLyR7cGF0Y2hGaWxlbmFtZX1cbiAgICBQYXRjaCB3YXMgbWFkZSBmb3IgdmVyc2lvbjogJHtjaGFsay5ncmVlbi5ib2xkKG9yaWdpbmFsVmVyc2lvbil9XG4gICAgSW5zdGFsbGVkIHZlcnNpb246ICR7Y2hhbGsucmVkLmJvbGQoYWN0dWFsVmVyc2lvbil9XG5gXG59XG5cbmZ1bmN0aW9uIGNyZWF0ZVVuZXhwZWN0ZWRFcnJvcih7XG4gIGZpbGVuYW1lLFxuICBlcnJvcixcbn06IHtcbiAgZmlsZW5hbWU6IHN0cmluZ1xuICBlcnJvcjogRXJyb3Jcbn0pIHtcbiAgcmV0dXJuIGBcbiR7Y2hhbGsucmVkLmJvbGQoXCIqKkVSUk9SKipcIil9ICR7Y2hhbGsucmVkKFxuICAgIGBGYWlsZWQgdG8gYXBwbHkgcGF0Y2ggZmlsZSAke2NoYWxrLmJvbGQoZmlsZW5hbWUpfWAsXG4gICl9XG4gIFxuJHtlcnJvci5zdGFja31cblxuICBgXG59XG4iXX0=