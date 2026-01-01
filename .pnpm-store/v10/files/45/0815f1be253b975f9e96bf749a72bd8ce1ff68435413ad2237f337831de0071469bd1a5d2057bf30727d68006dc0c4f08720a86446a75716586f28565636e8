"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.rebase = void 0;
const chalk_1 = __importDefault(require("chalk"));
const path_1 = require("path");
const applyPatches_1 = require("./applyPatches");
const hash_1 = require("./hash");
const patchFs_1 = require("./patchFs");
const stateFile_1 = require("./stateFile");
function rebase({ appPath, patchDir, packagePathSpecifier, targetPatch, }) {
    const patchesDirectory = path_1.join(appPath, patchDir);
    const groupedPatches = patchFs_1.getGroupedPatches(patchesDirectory);
    if (groupedPatches.numPatchFiles === 0) {
        console.log(chalk_1.default.blueBright("No patch files found"));
        process.exit(1);
    }
    const packagePatches = groupedPatches.pathSpecifierToPatchFiles[packagePathSpecifier];
    if (!packagePatches) {
        console.log(chalk_1.default.blueBright("No patch files found for package"), packagePathSpecifier);
        process.exit(1);
    }
    const state = stateFile_1.getPatchApplicationState(packagePatches[0]);
    if (!state) {
        console.log(chalk_1.default.blueBright("No patch state found"), "Did you forget to run", chalk_1.default.bold("patch-package"), "(without arguments) first?");
        process.exit(1);
    }
    if (state.isRebasing) {
        console.log(chalk_1.default.blueBright("Already rebasing"), "Make changes to the files in", chalk_1.default.bold(packagePatches[0].path), "and then run `patch-package", packagePathSpecifier, "--continue` to", packagePatches.length === state.patches.length
            ? "append a patch file"
            : `update the ${packagePatches[packagePatches.length - 1].patchFilename} file`);
        console.log(`💡 To remove a broken patch file, delete it and reinstall node_modules`);
        process.exit(1);
    }
    if (state.patches.length !== packagePatches.length) {
        console.log(chalk_1.default.blueBright("Some patches have not been applied."), "Reinstall node_modules and try again.");
    }
    // check hashes
    stateFile_1.verifyAppliedPatches({ appPath, patchDir, state });
    if (targetPatch === "0") {
        // unapply all
        unApplyPatches({
            patches: packagePatches,
            appPath,
            patchDir,
        });
        stateFile_1.savePatchApplicationState({
            packageDetails: packagePatches[0],
            isRebasing: true,
            patches: [],
        });
        console.log(`
Make any changes you need inside ${chalk_1.default.bold(packagePatches[0].path)}

When you are done, run

  ${chalk_1.default.bold(`patch-package ${packagePathSpecifier} --append 'MyChangeDescription'`)}
  
to insert a new patch file.
`);
        return;
    }
    // find target patch
    const target = packagePatches.find((p) => {
        if (p.patchFilename === targetPatch) {
            return true;
        }
        if (path_1.resolve(process.cwd(), targetPatch) ===
            path_1.join(patchesDirectory, p.patchFilename)) {
            return true;
        }
        if (targetPatch === p.sequenceName) {
            return true;
        }
        const n = Number(targetPatch.replace(/^0+/g, ""));
        if (!isNaN(n) && n === p.sequenceNumber) {
            return true;
        }
        return false;
    });
    if (!target) {
        console.log(chalk_1.default.red("Could not find target patch file"), chalk_1.default.bold(targetPatch));
        console.log();
        console.log("The list of available patch files is:");
        packagePatches.forEach((p) => {
            console.log(`  - ${p.patchFilename}`);
        });
        process.exit(1);
    }
    const currentHash = hash_1.hashFile(path_1.join(patchesDirectory, target.patchFilename));
    const prevApplication = state.patches.find((p) => p.patchContentHash === currentHash);
    if (!prevApplication) {
        console.log(chalk_1.default.red("Could not find previous application of patch file"), chalk_1.default.bold(target.patchFilename));
        console.log();
        console.log("You should reinstall node_modules and try again.");
        process.exit(1);
    }
    // ok, we are good to start undoing all the patches that were applied up to but not including the target patch
    const targetIdx = state.patches.indexOf(prevApplication);
    unApplyPatches({
        patches: packagePatches.slice(targetIdx + 1),
        appPath,
        patchDir,
    });
    stateFile_1.savePatchApplicationState({
        packageDetails: packagePatches[0],
        isRebasing: true,
        patches: packagePatches.slice(0, targetIdx + 1).map((p) => ({
            patchFilename: p.patchFilename,
            patchContentHash: hash_1.hashFile(path_1.join(patchesDirectory, p.patchFilename)),
            didApply: true,
        })),
    });
    console.log(`
Make any changes you need inside ${chalk_1.default.bold(packagePatches[0].path)}

When you are done, do one of the following:

  To update ${chalk_1.default.bold(packagePatches[targetIdx].patchFilename)} run

    ${chalk_1.default.bold(`patch-package ${packagePathSpecifier}`)}
    
  To create a new patch file after ${chalk_1.default.bold(packagePatches[targetIdx].patchFilename)} run
  
    ${chalk_1.default.bold(`patch-package ${packagePathSpecifier} --append 'MyChangeDescription'`)}

  `);
}
exports.rebase = rebase;
function unApplyPatches({ patches, appPath, patchDir, }) {
    for (const patch of patches.slice().reverse()) {
        if (!applyPatches_1.applyPatch({
            patchFilePath: path_1.join(appPath, patchDir, patch.patchFilename),
            reverse: true,
            patchDetails: patch,
            patchDir,
            cwd: process.cwd(),
            bestEffort: false,
        })) {
            console.log(chalk_1.default.red("Failed to un-apply patch file"), chalk_1.default.bold(patch.patchFilename), "Try completely reinstalling node_modules.");
            process.exit(1);
        }
        console.log(chalk_1.default.cyan.bold("Un-applied"), patch.patchFilename);
    }
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicmViYXNlLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3JlYmFzZS50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7Ozs7QUFBQSxrREFBeUI7QUFDekIsK0JBQW9DO0FBQ3BDLGlEQUEyQztBQUMzQyxpQ0FBaUM7QUFFakMsdUNBQTZDO0FBQzdDLDJDQUlvQjtBQUVwQixTQUFnQixNQUFNLENBQUMsRUFDckIsT0FBTyxFQUNQLFFBQVEsRUFDUixvQkFBb0IsRUFDcEIsV0FBVyxHQU1aO0lBQ0MsTUFBTSxnQkFBZ0IsR0FBRyxXQUFJLENBQUMsT0FBTyxFQUFFLFFBQVEsQ0FBQyxDQUFBO0lBQ2hELE1BQU0sY0FBYyxHQUFHLDJCQUFpQixDQUFDLGdCQUFnQixDQUFDLENBQUE7SUFFMUQsSUFBSSxjQUFjLENBQUMsYUFBYSxLQUFLLENBQUMsRUFBRTtRQUN0QyxPQUFPLENBQUMsR0FBRyxDQUFDLGVBQUssQ0FBQyxVQUFVLENBQUMsc0JBQXNCLENBQUMsQ0FBQyxDQUFBO1FBQ3JELE9BQU8sQ0FBQyxJQUFJLENBQUMsQ0FBQyxDQUFDLENBQUE7S0FDaEI7SUFFRCxNQUFNLGNBQWMsR0FDbEIsY0FBYyxDQUFDLHlCQUF5QixDQUFDLG9CQUFvQixDQUFDLENBQUE7SUFDaEUsSUFBSSxDQUFDLGNBQWMsRUFBRTtRQUNuQixPQUFPLENBQUMsR0FBRyxDQUNULGVBQUssQ0FBQyxVQUFVLENBQUMsa0NBQWtDLENBQUMsRUFDcEQsb0JBQW9CLENBQ3JCLENBQUE7UUFDRCxPQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFBO0tBQ2hCO0lBRUQsTUFBTSxLQUFLLEdBQUcsb0NBQXdCLENBQUMsY0FBYyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUE7SUFFekQsSUFBSSxDQUFDLEtBQUssRUFBRTtRQUNWLE9BQU8sQ0FBQyxHQUFHLENBQ1QsZUFBSyxDQUFDLFVBQVUsQ0FBQyxzQkFBc0IsQ0FBQyxFQUN4Qyx1QkFBdUIsRUFDdkIsZUFB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RSxXQUFJLENBQUMsT0FBTyxFQUFFLFFBQVEsRUFBRSxLQUFLLENBQUMsYUFBYSxDQUFXO1lBQ3JFLE9BQU8sRUFBRSxJQUFJO1lBQ2IsWUFBWSxFQUFFLEtBQUs7WUFDbkIsUUFBUTtZQUNSLEdBQUcsRUFBRSxPQUFPLENBQUMsR0FBRyxFQUFFO1lBQ2xCLFVBQVUsRUFBRSxLQUFLO1NBQ2xCLENBQUMsRUFDRjtZQUNBLE9BQU8sQ0FBQyxHQUFHLENBQ1QsZUFBSyxDQUFDLEdBQUcsQ0FBQywrQkFBK0IsQ0FBQyxFQUMxQyxlQUFLLENBQUMsSUFBSSxDQUFDLEtBQUssQ0FBQyxhQUFhLENBQUMsRUFDL0IsMkNBQTJDLENBQzVDLENBQUE7WUFDRCxPQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFBO1NBQ2hCO1FBQ0QsT0FBTyxDQUFDLEdBQUcsQ0FBQyxlQUFLLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxZQUFZLENBQUMsRUFBRSxLQUFLLENBQUMsYUFBYSxDQUFDLENBQUE7S0FDaEU7QUFDSCxDQUFDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IGNoYWxrIGZyb20gXCJjaGFsa1wiXG5pbXBvcnQgeyBqb2luLCByZXNvbHZlIH0gZnJvbSBcInBhdGhcIlxuaW1wb3J0IHsgYXBwbHlQYXRjaCB9IGZyb20gXCIuL2FwcGx5UGF0Y2hlc1wiXG5pbXBvcnQgeyBoYXNoRmlsZSB9IGZyb20gXCIuL2hhc2hcIlxuaW1wb3J0IHsgUGF0Y2hlZFBhY2thZ2VEZXRhaWxzIH0gZnJvbSBcIi4vUGFja2FnZURldGFpbHNcIlxuaW1wb3J0IHsgZ2V0R3JvdXBlZFBhdGNoZXMgfSBmcm9tIFwiLi9wYXRjaEZzXCJcbmltcG9ydCB7XG4gIGdldFBhdGNoQXBwbGljYXRpb25TdGF0ZSxcbiAgc2F2ZVBhdGNoQXBwbGljYXRpb25TdGF0ZSxcbiAgdmVyaWZ5QXBwbGllZFBhdGNoZXMsXG59IGZyb20gXCIuL3N0YXRlRmlsZVwiXG5cbmV4cG9ydCBmdW5jdGlvbiByZWJhc2Uoe1xuICBhcHBQYXRoLFxuICBwYXRjaERpcixcbiAgcGFja2FnZVBhdGhTcGVjaWZpZXIsXG4gIHRhcmdldFBhdGNoLFxufToge1xuICBhcHBQYXRoOiBzdHJpbmdcbiAgcGF0Y2hEaXI6IHN0cmluZ1xuICBwYWNrYWdlUGF0aFNwZWNpZmllcjogc3RyaW5nXG4gIHRhcmdldFBhdGNoOiBzdHJpbmdcbn0pOiB2b2lkIHtcbiAgY29uc3QgcGF0Y2hlc0RpcmVjdG9yeSA9IGpvaW4oYXBwUGF0aCwgcGF0Y2hEaXIpXG4gIGNvbnN0IGdyb3VwZWRQYXRjaGVzID0gZ2V0R3JvdXBlZFBhdGNoZXMocGF0Y2hlc0RpcmVjdG9yeSlcblxuICBpZiAoZ3JvdXBlZFBhdGNoZXMubnVtUGF0Y2hGaWxlcyA9PT0gMCkge1xuICAgIGNvbnNvbGUubG9nKGNoYWxrLmJsdWVCcmlnaHQoXCJObyBwYXRjaCBmaWxlcyBmb3VuZFwiKSlcbiAgICBwcm9jZXNzLmV4aXQoMSlcbiAgfVxuXG4gIGNvbnN0IHBhY2thZ2VQYXRjaGVzID1cbiAgICBncm91cGVkUGF0Y2hlcy5wYXRoU3BlY2lmaWVyVG9QYXRjaEZpbGVzW3BhY2thZ2VQYXRoU3BlY2lmaWVyXVxuICBpZiAoIXBhY2thZ2VQYXRjaGVzKSB7XG4gICAgY29uc29sZS5sb2coXG4gICAgICBjaGFsay5ibHVlQnJpZ2h0KFwiTm8gcGF0Y2ggZmlsZXMgZm91bmQgZm9yIHBhY2thZ2VcIiksXG4gICAgICBwYWNrYWdlUGF0aFNwZWNpZmllcixcbiAgICApXG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cblxuICBjb25zdCBzdGF0ZSA9IGdldFBhdGNoQXBwbGljYXRpb25TdGF0ZShwYWNrYWdlUGF0Y2hlc1swXSlcblxuICBpZiAoIXN0YXRlKSB7XG4gICAgY29uc29sZS5sb2coXG4gICAgICBjaGFsay5ibHVlQnJpZ2h0KFwiTm8gcGF0Y2ggc3RhdGUgZm91bmRcIiksXG4gICAgICBcIkRpZCB5b3UgZm9yZ2V0IHRvIHJ1blwiLFxuICAgICAgY2hhbGsuYm9sZChcInBhdGNoLXBhY2thZ2VcIiksXG4gICAgICBcIih3aXRob3V0IGFyZ3VtZW50cykgZmlyc3Q/XCIsXG4gICAgKVxuICAgIHByb2Nlc3MuZXhpdCgxKVxuICB9XG4gIGlmIChzdGF0ZS5pc1JlYmFzaW5nKSB7XG4gICAgY29uc29sZS5sb2coXG4gICAgICBjaGFsay5ibHVlQnJpZ2h0KFwiQWxyZWFkeSByZWJhc2luZ1wiKSxcbiAgICAgIFwiTWFrZSBjaGFuZ2VzIHRvIHRoZSBmaWxlcyBpblwiLFxuICAgICAgY2hhbGsuYm9sZChwYWNrYWdlUGF0Y2hlc1swXS5wYXRoKSxcbiAgICAgIFwiYW5kIHRoZW4gcnVuIGBwYXRjaC1wYWNrYWdlXCIsXG4gICAgICBwYWNrYWdlUGF0aFNwZWNpZmllcixcbiAgICAgIFwiLS1jb250aW51ZWAgdG9cIixcbiAgICAgIHBhY2thZ2VQYXRjaGVzLmxlbmd0aCA9PT0gc3RhdGUucGF0Y2hlcy5sZW5ndGhcbiAgICAgICAgPyBcImFwcGVuZCBhIHBhdGNoIGZpbGVcIlxuICAgICAgICA6IGB1cGRhdGUgdGhlICR7XG4gICAgICAgICAgICBwYWNrYWdlUGF0Y2hlc1twYWNrYWdlUGF0Y2hlcy5sZW5ndGggLSAxXS5wYXRjaEZpbGVuYW1lXG4gICAgICAgICAgfSBmaWxlYCxcbiAgICApXG4gICAgY29uc29sZS5sb2coXG4gICAgICBg8J+SoSBUbyByZW1vdmUgYSBicm9rZW4gcGF0Y2ggZmlsZSwgZGVsZXRlIGl0IGFuZCByZWluc3RhbGwgbm9kZV9tb2R1bGVzYCxcbiAgICApXG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cbiAgaWYgKHN0YXRlLnBhdGNoZXMubGVuZ3RoICE9PSBwYWNrYWdlUGF0Y2hlcy5sZW5ndGgpIHtcbiAgICBjb25zb2xlLmxvZyhcbiAgICAgIGNoYWxrLmJsdWVCcmlnaHQoXCJTb21lIHBhdGNoZXMgaGF2ZSBub3QgYmVlbiBhcHBsaWVkLlwiKSxcbiAgICAgIFwiUmVpbnN0YWxsIG5vZGVfbW9kdWxlcyBhbmQgdHJ5IGFnYWluLlwiLFxuICAgIClcbiAgfVxuICAvLyBjaGVjayBoYXNoZXNcbiAgdmVyaWZ5QXBwbGllZFBhdGNoZXMoeyBhcHBQYXRoLCBwYXRjaERpciwgc3RhdGUgfSlcblxuICBpZiAodGFyZ2V0UGF0Y2ggPT09IFwiMFwiKSB7XG4gICAgLy8gdW5hcHBseSBhbGxcbiAgICB1bkFwcGx5UGF0Y2hlcyh7XG4gICAgICBwYXRjaGVzOiBwYWNrYWdlUGF0Y2hlcyxcbiAgICAgIGFwcFBhdGgsXG4gICAgICBwYXRjaERpcixcbiAgICB9KVxuICAgIHNhdmVQYXRjaEFwcGxpY2F0aW9uU3RhdGUoe1xuICAgICAgcGFja2FnZURldGFpbHM6IHBhY2thZ2VQYXRjaGVzWzBdLFxuICAgICAgaXNSZWJhc2luZzogdHJ1ZSxcbiAgICAgIHBhdGNoZXM6IFtdLFxuICAgIH0pXG4gICAgY29uc29sZS5sb2coYFxuTWFrZSBhbnkgY2hhbmdlcyB5b3UgbmVlZCBpbnNpZGUgJHtjaGFsay5ib2xkKHBhY2thZ2VQYXRjaGVzWzBdLnBhdGgpfVxuXG5XaGVuIHlvdSBhcmUgZG9uZSwgcnVuXG5cbiAgJHtjaGFsay5ib2xkKFxuICAgIGBwYXRjaC1wYWNrYWdlICR7cGFja2FnZVBhdGhTcGVjaWZpZXJ9IC0tYXBwZW5kICdNeUNoYW5nZURlc2NyaXB0aW9uJ2AsXG4gICl9XG4gIFxudG8gaW5zZXJ0IGEgbmV3IHBhdGNoIGZpbGUuXG5gKVxuICAgIHJldHVyblxuICB9XG5cbiAgLy8gZmluZCB0YXJnZXQgcGF0Y2hcbiAgY29uc3QgdGFyZ2V0ID0gcGFja2FnZVBhdGNoZXMuZmluZCgocCkgPT4ge1xuICAgIGlmIChwLnBhdGNoRmlsZW5hbWUgPT09IHRhcmdldFBhdGNoKSB7XG4gICAgICByZXR1cm4gdHJ1ZVxuICAgIH1cbiAgICBpZiAoXG4gICAgICByZXNvbHZlKHByb2Nlc3MuY3dkKCksIHRhcmdldFBhdGNoKSA9PT1cbiAgICAgIGpvaW4ocGF0Y2hlc0RpcmVjdG9yeSwgcC5wYXRjaEZpbGVuYW1lKVxuICAgICkge1xuICAgICAgcmV0dXJuIHRydWVcbiAgICB9XG5cbiAgICBpZiAodGFyZ2V0UGF0Y2ggPT09IHAuc2VxdWVuY2VOYW1lKSB7XG4gICAgICByZXR1cm4gdHJ1ZVxuICAgIH1cbiAgICBjb25zdCBuID0gTnVtYmVyKHRhcmdldFBhdGNoLnJlcGxhY2UoL14wKy9nLCBcIlwiKSlcbiAgICBpZiAoIWlzTmFOKG4pICYmIG4gPT09IHAuc2VxdWVuY2VOdW1iZXIpIHtcbiAgICAgIHJldHVybiB0cnVlXG4gICAgfVxuICAgIHJldHVybiBmYWxzZVxuICB9KVxuXG4gIGlmICghdGFyZ2V0KSB7XG4gICAgY29uc29sZS5sb2coXG4gICAgICBjaGFsay5yZWQoXCJDb3VsZCBub3QgZmluZCB0YXJnZXQgcGF0Y2ggZmlsZVwiKSxcbiAgICAgIGNoYWxrLmJvbGQodGFyZ2V0UGF0Y2gpLFxuICAgIClcbiAgICBjb25zb2xlLmxvZygpXG4gICAgY29uc29sZS5sb2coXCJUaGUgbGlzdCBvZiBhdmFpbGFibGUgcGF0Y2ggZmlsZXMgaXM6XCIpXG4gICAgcGFja2FnZVBhdGNoZXMuZm9yRWFjaCgocCkgPT4ge1xuICAgICAgY29uc29sZS5sb2coYCAgLSAke3AucGF0Y2hGaWxlbmFtZX1gKVxuICAgIH0pXG5cbiAgICBwcm9jZXNzLmV4aXQoMSlcbiAgfVxuICBjb25zdCBjdXJyZW50SGFzaCA9IGhhc2hGaWxlKGpvaW4ocGF0Y2hlc0RpcmVjdG9yeSwgdGFyZ2V0LnBhdGNoRmlsZW5hbWUpKVxuXG4gIGNvbnN0IHByZXZBcHBsaWNhdGlvbiA9IHN0YXRlLnBhdGNoZXMuZmluZChcbiAgICAocCkgPT4gcC5wYXRjaENvbnRlbnRIYXNoID09PSBjdXJyZW50SGFzaCxcbiAgKVxuICBpZiAoIXByZXZBcHBsaWNhdGlvbikge1xuICAgIGNvbnNvbGUubG9nKFxuICAgICAgY2hhbGsucmVkKFwiQ291bGQgbm90IGZpbmQgcHJldmlvdXMgYXBwbGljYXRpb24gb2YgcGF0Y2ggZmlsZVwiKSxcbiAgICAgIGNoYWxrLmJvbGQodGFyZ2V0LnBhdGNoRmlsZW5hbWUpLFxuICAgIClcbiAgICBjb25zb2xlLmxvZygpXG4gICAgY29uc29sZS5sb2coXCJZb3Ugc2hvdWxkIHJlaW5zdGFsbCBub2RlX21vZHVsZXMgYW5kIHRyeSBhZ2Fpbi5cIilcbiAgICBwcm9jZXNzLmV4aXQoMSlcbiAgfVxuXG4gIC8vIG9rLCB3ZSBhcmUgZ29vZCB0byBzdGFydCB1bmRvaW5nIGFsbCB0aGUgcGF0Y2hlcyB0aGF0IHdlcmUgYXBwbGllZCB1cCB0byBidXQgbm90IGluY2x1ZGluZyB0aGUgdGFyZ2V0IHBhdGNoXG4gIGNvbnN0IHRhcmdldElkeCA9IHN0YXRlLnBhdGNoZXMuaW5kZXhPZihwcmV2QXBwbGljYXRpb24pXG5cbiAgdW5BcHBseVBhdGNoZXMoe1xuICAgIHBhdGNoZXM6IHBhY2thZ2VQYXRjaGVzLnNsaWNlKHRhcmdldElkeCArIDEpLFxuICAgIGFwcFBhdGgsXG4gICAgcGF0Y2hEaXIsXG4gIH0pXG4gIHNhdmVQYXRjaEFwcGxpY2F0aW9uU3RhdGUoe1xuICAgIHBhY2thZ2VEZXRhaWxzOiBwYWNrYWdlUGF0Y2hlc1swXSxcbiAgICBpc1JlYmFzaW5nOiB0cnVlLFxuICAgIHBhdGNoZXM6IHBhY2thZ2VQYXRjaGVzLnNsaWNlKDAsIHRhcmdldElkeCArIDEpLm1hcCgocCkgPT4gKHtcbiAgICAgIHBhdGNoRmlsZW5hbWU6IHAucGF0Y2hGaWxlbmFtZSxcbiAgICAgIHBhdGNoQ29udGVudEhhc2g6IGhhc2hGaWxlKGpvaW4ocGF0Y2hlc0RpcmVjdG9yeSwgcC5wYXRjaEZpbGVuYW1lKSksXG4gICAgICBkaWRBcHBseTogdHJ1ZSxcbiAgICB9KSksXG4gIH0pXG5cbiAgY29uc29sZS5sb2coYFxuTWFrZSBhbnkgY2hhbmdlcyB5b3UgbmVlZCBpbnNpZGUgJHtjaGFsay5ib2xkKHBhY2thZ2VQYXRjaGVzWzBdLnBhdGgpfVxuXG5XaGVuIHlvdSBhcmUgZG9uZSwgZG8gb25lIG9mIHRoZSBmb2xsb3dpbmc6XG5cbiAgVG8gdXBkYXRlICR7Y2hhbGsuYm9sZChwYWNrYWdlUGF0Y2hlc1t0YXJnZXRJZHhdLnBhdGNoRmlsZW5hbWUpfSBydW5cblxuICAgICR7Y2hhbGsuYm9sZChgcGF0Y2gtcGFja2FnZSAke3BhY2thZ2VQYXRoU3BlY2lmaWVyfWApfVxuICAgIFxuICBUbyBjcmVhdGUgYSBuZXcgcGF0Y2ggZmlsZSBhZnRlciAke2NoYWxrLmJvbGQoXG4gICAgcGFja2FnZVBhdGNoZXNbdGFyZ2V0SWR4XS5wYXRjaEZpbGVuYW1lLFxuICApfSBydW5cbiAgXG4gICAgJHtjaGFsay5ib2xkKFxuICAgICAgYHBhdGNoLXBhY2thZ2UgJHtwYWNrYWdlUGF0aFNwZWNpZmllcn0gLS1hcHBlbmQgJ015Q2hhbmdlRGVzY3JpcHRpb24nYCxcbiAgICApfVxuXG4gIGApXG59XG5cbmZ1bmN0aW9uIHVuQXBwbHlQYXRjaGVzKHtcbiAgcGF0Y2hlcyxcbiAgYXBwUGF0aCxcbiAgcGF0Y2hEaXIsXG59OiB7XG4gIHBhdGNoZXM6IFBhdGNoZWRQYWNrYWdlRGV0YWlsc1tdXG4gIGFwcFBhdGg6IHN0cmluZ1xuICBwYXRjaERpcjogc3RyaW5nXG59KSB7XG4gIGZvciAoY29uc3QgcGF0Y2ggb2YgcGF0Y2hlcy5zbGljZSgpLnJldmVyc2UoKSkge1xuICAgIGlmIChcbiAgICAgICFhcHBseVBhdGNoKHtcbiAgICAgICAgcGF0Y2hGaWxlUGF0aDogam9pbihhcHBQYXRoLCBwYXRjaERpciwgcGF0Y2gucGF0Y2hGaWxlbmFtZSkgYXMgc3RyaW5nLFxuICAgICAgICByZXZlcnNlOiB0cnVlLFxuICAgICAgICBwYXRjaERldGFpbHM6IHBhdGNoLFxuICAgICAgICBwYXRjaERpcixcbiAgICAgICAgY3dkOiBwcm9jZXNzLmN3ZCgpLFxuICAgICAgICBiZXN0RWZmb3J0OiBmYWxzZSxcbiAgICAgIH0pXG4gICAgKSB7XG4gICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgY2hhbGsucmVkKFwiRmFpbGVkIHRvIHVuLWFwcGx5IHBhdGNoIGZpbGVcIiksXG4gICAgICAgIGNoYWxrLmJvbGQocGF0Y2gucGF0Y2hGaWxlbmFtZSksXG4gICAgICAgIFwiVHJ5IGNvbXBsZXRlbHkgcmVpbnN0YWxsaW5nIG5vZGVfbW9kdWxlcy5cIixcbiAgICAgIClcbiAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgIH1cbiAgICBjb25zb2xlLmxvZyhjaGFsay5jeWFuLmJvbGQoXCJVbi1hcHBsaWVkXCIpLCBwYXRjaC5wYXRjaEZpbGVuYW1lKVxuICB9XG59XG4iXX0=