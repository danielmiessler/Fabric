"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const chalk_1 = __importDefault(require("chalk"));
const process_1 = __importDefault(require("process"));
const minimist_1 = __importDefault(require("minimist"));
const applyPatches_1 = require("./applyPatches");
const getAppRootPath_1 = require("./getAppRootPath");
const makePatch_1 = require("./makePatch");
const makeRegExp_1 = require("./makeRegExp");
const detectPackageManager_1 = require("./detectPackageManager");
const path_1 = require("./path");
const path_2 = require("path");
const slash = require("slash");
const ci_info_1 = require("ci-info");
const rebase_1 = require("./rebase");
const appPath = getAppRootPath_1.getAppRootPath();
const argv = minimist_1.default(process_1.default.argv.slice(2), {
    boolean: [
        "use-yarn",
        "case-sensitive-path-filtering",
        "reverse",
        "help",
        "version",
        "error-on-fail",
        "error-on-warn",
        "create-issue",
        "partial",
        "",
    ],
    string: ["patch-dir", "append", "rebase"],
});
const packageNames = argv._;
console.log(chalk_1.default.bold("patch-package"), 
// tslint:disable-next-line:no-var-requires
require(path_1.join(__dirname, "../package.json")).version);
if (argv.version || argv.v) {
    // noop
}
else if (argv.help || argv.h) {
    printHelp();
}
else {
    const patchDir = slash(path_2.normalize((argv["patch-dir"] || "patches") + path_2.sep));
    if (patchDir.startsWith("/")) {
        throw new Error("--patch-dir must be a relative path");
    }
    if ("rebase" in argv) {
        if (!argv.rebase) {
            console.log(chalk_1.default.red("You must specify a patch file name or number when rebasing patches"));
            process_1.default.exit(1);
        }
        if (packageNames.length !== 1) {
            console.log(chalk_1.default.red("You must specify exactly one package name when rebasing patches"));
            process_1.default.exit(1);
        }
        rebase_1.rebase({
            appPath,
            packagePathSpecifier: packageNames[0],
            patchDir,
            targetPatch: argv.rebase,
        });
    }
    else if (packageNames.length) {
        const includePaths = makeRegExp_1.makeRegExp(argv.include, "include", /.*/, argv["case-sensitive-path-filtering"]);
        const excludePaths = makeRegExp_1.makeRegExp(argv.exclude, "exclude", /^package\.json$/, argv["case-sensitive-path-filtering"]);
        const packageManager = detectPackageManager_1.detectPackageManager(appPath, argv["use-yarn"] ? "yarn" : null);
        const createIssue = argv["create-issue"];
        packageNames.forEach((packagePathSpecifier) => {
            makePatch_1.makePatch({
                packagePathSpecifier,
                appPath,
                packageManager,
                includePaths,
                excludePaths,
                patchDir,
                createIssue,
                mode: "append" in argv
                    ? { type: "append", name: argv.append || undefined }
                    : { type: "overwrite_last" },
            });
        });
    }
    else {
        console.log("Applying patches...");
        const reverse = !!argv["reverse"];
        // don't want to exit(1) on postinstall locally.
        // see https://github.com/ds300/patch-package/issues/86
        const shouldExitWithError = !!argv["error-on-fail"] ||
            (process_1.default.env.NODE_ENV === "production" && ci_info_1.isCI) ||
            (ci_info_1.isCI && !process_1.default.env.PATCH_PACKAGE_INTEGRATION_TEST) ||
            process_1.default.env.NODE_ENV === "test";
        const shouldExitWithWarning = !!argv["error-on-warn"];
        applyPatches_1.applyPatchesForApp({
            appPath,
            reverse,
            patchDir,
            shouldExitWithError,
            shouldExitWithWarning,
            bestEffort: argv.partial,
        });
    }
}
function printHelp() {
    console.log(`
Usage:

  1. Patching packages
  ====================

    ${chalk_1.default.bold("patch-package")}

  Without arguments, the ${chalk_1.default.bold("patch-package")} command will attempt to find and apply
  patch files to your project. It looks for files named like

     ./patches/<package-name>+<version>.patch

  Options:

    ${chalk_1.default.bold("--patch-dir <dirname>")}

      Specify the name for the directory in which the patch files are located.
      
    ${chalk_1.default.bold("--error-on-fail")}
    
      Forces patch-package to exit with code 1 after failing.
    
      When running locally patch-package always exits with 0 by default.
      This happens even after failing to apply patches because otherwise 
      yarn.lock and package.json might get out of sync with node_modules,
      which can be very confusing.
      
      --error-on-fail is ${chalk_1.default.bold("switched on")} by default on CI.
      
      See https://github.com/ds300/patch-package/issues/86 for background.
      
    ${chalk_1.default.bold("--error-on-warn")}
    
      Forces patch-package to exit with code 1 after warning.
      
      See https://github.com/ds300/patch-package/issues/314 for background.

    ${chalk_1.default.bold("--reverse")}
        
      Un-applies all patches.

      Note that this will fail if the patched files have changed since being
      patched. In that case, you'll probably need to re-install 'node_modules'.

      This option was added to help people using CircleCI avoid an issue around caching
      and patch file updates (https://github.com/ds300/patch-package/issues/37),
      but might be useful in other contexts too.
      

  2. Creating patch files
  =======================

    ${chalk_1.default.bold("patch-package")} <package-name>${chalk_1.default.italic("[ <package-name>]")}

  When given package names as arguments, patch-package will create patch files
  based on any changes you've made to the versions installed by yarn/npm.

  Options:
  
    ${chalk_1.default.bold("--create-issue")}
    
       For packages whose source is hosted on GitHub this option opens a web
       browser with a draft issue based on your diff.

    ${chalk_1.default.bold("--use-yarn")}

        By default, patch-package checks whether you use npm or yarn based on
        which lockfile you have. If you have both, it uses npm by default.
        Set this option to override that default and always use yarn.

    ${chalk_1.default.bold("--exclude <regexp>")}

        Ignore paths matching the regexp when creating patch files.
        Paths are relative to the root dir of the package to be patched.

        Default: 'package\\.json$'

    ${chalk_1.default.bold("--include <regexp>")}

        Only consider paths matching the regexp when creating patch files.
        Paths are relative to the root dir of the package to be patched.

        Default '.*'

    ${chalk_1.default.bold("--case-sensitive-path-filtering")}

        Make regexps used in --include or --exclude filters case-sensitive.
    
    ${chalk_1.default.bold("--patch-dir")}

        Specify the name for the directory in which to put the patch files.
`);
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5kZXguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvaW5kZXgudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7QUFBQSxrREFBeUI7QUFDekIsc0RBQTZCO0FBQzdCLHdEQUErQjtBQUUvQixpREFBbUQ7QUFDbkQscURBQWlEO0FBQ2pELDJDQUF1QztBQUN2Qyw2Q0FBeUM7QUFDekMsaUVBQTZEO0FBQzdELGlDQUE2QjtBQUM3QiwrQkFBcUM7QUFDckMsK0JBQStCO0FBQy9CLHFDQUE4QjtBQUM5QixxQ0FBaUM7QUFFakMsTUFBTSxPQUFPLEdBQUcsK0JBQWMsRUFBRSxDQUFBO0FBQ2hDLE1BQU0sSUFBSSxHQUFHLGtCQUFRLENBQUMsaUJBQU8sQ0FBQyxJQUFJLENBQUMsS0FBSyxDQUFDLENBQUMsQ0FBQyxFQUFFO0lBQzNDLE9BQU8sRUFBRTtRQUNQLFVBQVU7UUFDViwrQkFBK0I7UUFDL0IsU0FBUztRQUNULE1BQU07UUFDTixTQUFTO1FBQ1QsZUFBZTtRQUNmLGVBQWU7UUFDZixjQUFjO1FBQ2QsU0FBUztRQUNULEVBQUU7S0FDSDtJQUNELE1BQU0sRUFBRSxDQUFDLFdBQVcsRUFBRSxRQUFRLEVBQUUsUUFBUSxDQUFDO0NBQzFDLENBQUMsQ0FBQTtBQUNGLE1BQU0sWUFBWSxHQUFHLElBQUksQ0FBQyxDQUFDLENBQUE7QUFFM0IsT0FBTyxDQUFDLEdBQUcsQ0FDVCxlQUFLLENBQUMsSUFBSSxDQUFDLGVBQWUsQ0FBQztBQUMzQiwyQ0FBMkM7QUFDM0MsT0FBTyxDQUFDLFdBQUksQ0FBQyxTQUFTLEVBQUUsaUJBQWlCLENBQUMsQ0FBQyxDQUFDLE9BQU8sQ0FDcEQsQ0FBQTtBQUVELElBQUksSUFBSSxDQUFDLE9BQU8sSUFBSSxJQUFJLENBQUMsQ0FBQyxFQUFFO0lBQzFCLE9BQU87Q0FDUjtLQUFNLElBQUksSUFBSSxDQUFDLElBQUksSUFBSSxJQUFJLENBQUMsQ0FBQyxFQUFFO0lBQzlCLFNBQVMsRUFBRSxDQUFBO0NB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YWxrXCJcbmltcG9ydCBwcm9jZXNzIGZyb20gXCJwcm9jZXNzXCJcbmltcG9ydCBtaW5pbWlzdCBmcm9tIFwibWluaW1pc3RcIlxuXG5pbXBvcnQgeyBhcHBseVBhdGNoZXNGb3JBcHAgfSBmcm9tIFwiLi9hcHBseVBhdGNoZXNcIlxuaW1wb3J0IHsgZ2V0QXBwUm9vdFBhdGggfSBmcm9tIFwiLi9nZXRBcHBSb290UGF0aFwiXG5pbXBvcnQgeyBtYWtlUGF0Y2ggfSBmcm9tIFwiLi9tYWtlUGF0Y2hcIlxuaW1wb3J0IHsgbWFrZVJlZ0V4cCB9IGZyb20gXCIuL21ha2VSZWdFeHBcIlxuaW1wb3J0IHsgZGV0ZWN0UGFja2FnZU1hbmFnZXIgfSBmcm9tIFwiLi9kZXRlY3RQYWNrYWdlTWFuYWdlclwiXG5pbXBvcnQgeyBqb2luIH0gZnJvbSBcIi4vcGF0aFwiXG5pbXBvcnQgeyBub3JtYWxpemUsIHNlcCB9IGZyb20gXCJwYXRoXCJcbmltcG9ydCBzbGFzaCA9IHJlcXVpcmUoXCJzbGFzaFwiKVxuaW1wb3J0IHsgaXNDSSB9IGZyb20gXCJjaS1pbmZvXCJcbmltcG9ydCB7IHJlYmFzZSB9IGZyb20gXCIuL3JlYmFzZVwiXG5cbmNvbnN0IGFwcFBhdGggPSBnZXRBcHBSb290UGF0aCgpXG5jb25zdCBhcmd2ID0gbWluaW1pc3QocHJvY2Vzcy5hcmd2LnNsaWNlKDIpLCB7XG4gIGJvb2xlYW46IFtcbiAgICBcInVzZS15YXJuXCIsXG4gICAgXCJjYXNlLXNlbnNpdGl2ZS1wYXRoLWZpbHRlcmluZ1wiLFxuICAgIFwicmV2ZXJzZVwiLFxuICAgIFwiaGVscFwiLFxuICAgIFwidmVyc2lvblwiLFxuICAgIFwiZXJyb3Itb24tZmFpbFwiLFxuICAgIFwiZXJyb3Itb24td2FyblwiLFxuICAgIFwiY3JlYXRlLWlzc3VlXCIsXG4gICAgXCJwYXJ0aWFsXCIsXG4gICAgXCJcIixcbiAgXSxcbiAgc3RyaW5nOiBbXCJwYXRjaC1kaXJcIiwgXCJhcHBlbmRcIiwgXCJyZWJhc2VcIl0sXG59KVxuY29uc3QgcGFja2FnZU5hbWVzID0gYXJndi5fXG5cbmNvbnNvbGUubG9nKFxuICBjaGFsay5ib2xkKFwicGF0Y2gtcGFja2FnZVwiKSxcbiAgLy8gdHNsaW50OmRpc2FibGUtbmV4dC1saW5lOm5vLXZhci1yZXF1aXJlc1xuICByZXF1aXJlKGpvaW4oX19kaXJuYW1lLCBcIi4uL3BhY2thZ2UuanNvblwiKSkudmVyc2lvbixcbilcblxuaWYgKGFyZ3YudmVyc2lvbiB8fCBhcmd2LnYpIHtcbiAgLy8gbm9vcFxufSBlbHNlIGlmIChhcmd2LmhlbHAgfHwgYXJndi5oKSB7XG4gIHByaW50SGVscCgpXG59IGVsc2Uge1xuICBjb25zdCBwYXRjaERpciA9IHNsYXNoKG5vcm1hbGl6ZSgoYXJndltcInBhdGNoLWRpclwiXSB8fCBcInBhdGNoZXNcIikgKyBzZXApKVxuICBpZiAocGF0Y2hEaXIuc3RhcnRzV2l0aChcIi9cIikpIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoXCItLXBhdGNoLWRpciBtdXN0IGJlIGEgcmVsYXRpdmUgcGF0aFwiKVxuICB9XG4gIGlmIChcInJlYmFzZVwiIGluIGFyZ3YpIHtcbiAgICBpZiAoIWFyZ3YucmViYXNlKSB7XG4gICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgY2hhbGsucmVkKFxuICAgICAgICAgIFwiWW91IG11c3Qgc3BlY2lmeSBhIHBhdGNoIGZpbGUgbmFtZSBvciBudW1iZXIgd2hlbiByZWJhc2luZyBwYXRjaGVzXCIsXG4gICAgICAgICksXG4gICAgICApXG4gICAgICBwcm9jZXNzLmV4aXQoMSlcbiAgICB9XG4gICAgaWYgKHBhY2thZ2VOYW1lcy5sZW5ndGggIT09IDEpIHtcbiAgICAgIGNvbnNvbGUubG9nKFxuICAgICAgICBjaGFsay5yZWQoXG4gICAgICAgICAgXCJZb3UgbXVzdCBzcGVjaWZ5IGV4YWN0bHkgb25lIHBhY2thZ2UgbmFtZSB3aGVuIHJlYmFzaW5nIHBhdGNoZXNcIixcbiAgICAgICAgKSxcbiAgICAgIClcbiAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgIH1cbiAgICByZWJhc2Uoe1xuICAgICAgYXBwUGF0aCxcbiAgICAgIHBhY2thZ2VQYXRoU3BlY2lmaWVyOiBwYWNrYWdlTmFtZXNbMF0sXG4gICAgICBwYXRjaERpcixcbiAgICAgIHRhcmdldFBhdGNoOiBhcmd2LnJlYmFzZSxcbiAgICB9KVxuICB9IGVsc2UgaWYgKHBhY2thZ2VOYW1lcy5sZW5ndGgpIHtcbiAgICBjb25zdCBpbmNsdWRlUGF0aHMgPSBtYWtlUmVnRXhwKFxuICAgICAgYXJndi5pbmNsdWRlLFxuICAgICAgXCJpbmNsdWRlXCIsXG4gICAgICAvLiovLFxuICAgICAgYXJndltcImNhc2Utc2Vuc2l0aXZlLXBhdGgtZmlsdGVyaW5nXCJdLFxuICAgIClcbiAgICBjb25zdCBleGNsdWRlUGF0aHMgPSBtYWtlUmVnRXhwKFxuICAgICAgYXJndi5leGNsdWRlLFxuICAgICAgXCJleGNsdWRlXCIsXG4gICAgICAvXnBhY2thZ2VcXC5qc29uJC8sXG4gICAgICBhcmd2W1wiY2FzZS1zZW5zaXRpdmUtcGF0aC1maWx0ZXJpbmdcIl0sXG4gICAgKVxuICAgIGNvbnN0IHBhY2thZ2VNYW5hZ2VyID0gZGV0ZWN0UGFja2FnZU1hbmFnZXIoXG4gICAgICBhcHBQYXRoLFxuICAgICAgYXJndltcInVzZS15YXJuXCJdID8gXCJ5YXJuXCIgOiBudWxsLFxuICAgIClcbiAgICBjb25zdCBjcmVhdGVJc3N1ZSA9IGFyZ3ZbXCJjcmVhdGUtaXNzdWVcIl1cbiAgICBwYWNrYWdlTmFtZXMuZm9yRWFjaCgocGFja2FnZVBhdGhTcGVjaWZpZXI6IHN0cmluZykgPT4ge1xuICAgICAgbWFrZVBhdGNoKHtcbiAgICAgICAgcGFja2FnZVBhdGhTcGVjaWZpZXIsXG4gICAgICAgIGFwcFBhdGgsXG4gICAgICAgIHBhY2thZ2VNYW5hZ2VyLFxuICAgICAgICBpbmNsdWRlUGF0aHMsXG4gICAgICAgIGV4Y2x1ZGVQYXRocyxcbiAgICAgICAgcGF0Y2hEaXIsXG4gICAgICAgIGNyZWF0ZUlzc3VlLFxuICAgICAgICBtb2RlOlxuICAgICAgICAgIFwiYXBwZW5kXCIgaW4gYXJndlxuICAgICAgICAgICAgPyB7IHR5cGU6IFwiYXBwZW5kXCIsIG5hbWU6IGFyZ3YuYXBwZW5kIHx8IHVuZGVmaW5lZCB9XG4gICAgICAgICAgICA6IHsgdHlwZTogXCJvdmVyd3JpdGVfbGFzdFwiIH0sXG4gICAgICB9KVxuICAgIH0pXG4gIH0gZWxzZSB7XG4gICAgY29uc29sZS5sb2coXCJBcHBseWluZyBwYXRjaGVzLi4uXCIpXG4gICAgY29uc3QgcmV2ZXJzZSA9ICEhYXJndltcInJldmVyc2VcIl1cbiAgICAvLyBkb24ndCB3YW50IHRvIGV4aXQoMSkgb24gcG9zdGluc3RhbGwgbG9jYWxseS5cbiAgICAvLyBzZWUgaHR0cHM6Ly9naXRodWIuY29tL2RzMzAwL3BhdGNoLXBhY2thZ2UvaXNzdWVzLzg2XG4gICAgY29uc3Qgc2hvdWxkRXhpdFdpdGhFcnJvciA9XG4gICAgICAhIWFyZ3ZbXCJlcnJvci1vbi1mYWlsXCJdIHx8XG4gICAgICAocHJvY2Vzcy5lbnYuTk9ERV9FTlYgPT09IFwicHJvZHVjdGlvblwiICYmIGlzQ0kpIHx8XG4gICAgICAoaXNDSSAmJiAhcHJvY2Vzcy5lbnYuUEFUQ0hfUEFDS0FHRV9JTlRFR1JBVElPTl9URVNUKSB8fFxuICAgICAgcHJvY2Vzcy5lbnYuTk9ERV9FTlYgPT09IFwidGVzdFwiXG5cbiAgICBjb25zdCBzaG91bGRFeGl0V2l0aFdhcm5pbmcgPSAhIWFyZ3ZbXCJlcnJvci1vbi13YXJuXCJdXG5cbiAgICBhcHBseVBhdGNoZXNGb3JBcHAoe1xuICAgICAgYXBwUGF0aCxcbiAgICAgIHJldmVyc2UsXG4gICAgICBwYXRjaERpcixcbiAgICAgIHNob3VsZEV4aXRXaXRoRXJyb3IsXG4gICAgICBzaG91bGRFeGl0V2l0aFdhcm5pbmcsXG4gICAgICBiZXN0RWZmb3J0OiBhcmd2LnBhcnRpYWwsXG4gICAgfSlcbiAgfVxufVxuXG5mdW5jdGlvbiBwcmludEhlbHAoKSB7XG4gIGNvbnNvbGUubG9nKGBcblVzYWdlOlxuXG4gIDEuIFBhdGNoaW5nIHBhY2thZ2VzXG4gID09PT09PT09PT09PT09PT09PT09XG5cbiAgICAke2NoYWxrLmJvbGQoXCJwYXRjaC1wYWNrYWdlXCIpfVxuXG4gIFdpdGhvdXQgYXJndW1lbnRzLCB0aGUgJHtjaGFsay5ib2xkKFxuICAgIFwicGF0Y2gtcGFja2FnZVwiLFxuICApfSBjb21tYW5kIHdpbGwgYXR0ZW1wdCB0byBmaW5kIGFuZCBhcHBseVxuICBwYXRjaCBmaWxlcyB0byB5b3VyIHByb2plY3QuIEl0IGxvb2tzIGZvciBmaWxlcyBuYW1lZCBsaWtlXG5cbiAgICAgLi9wYXRjaGVzLzxwYWNrYWdlLW5hbWU+Kzx2ZXJzaW9uPi5wYXRjaFxuXG4gIE9wdGlvbnM6XG5cbiAgICAke2NoYWxrLmJvbGQoXCItLXBhdGNoLWRpciA8ZGlybmFtZT5cIil9XG5cbiAgICAgIFNwZWNpZnkgdGhlIG5hbWUgZm9yIHRoZSBkaXJlY3RvcnkgaW4gd2hpY2ggdGhlIHBhdGNoIGZpbGVzIGFyZSBsb2NhdGVkLlxuICAgICAgXG4gICAgJHtjaGFsay5ib2xkKFwiLS1lcnJvci1vbi1mYWlsXCIpfVxuICAgIFxuICAgICAgRm9yY2VzIHBhdGNoLXBhY2thZ2UgdG8gZXhpdCB3aXRoIGNvZGUgMSBhZnRlciBmYWlsaW5nLlxuICAgIFxuICAgICAgV2hlbiBydW5uaW5nIGxvY2FsbHkgcGF0Y2gtcGFja2FnZSBhbHdheXMgZXhpdHMgd2l0aCAwIGJ5IGRlZmF1bHQuXG4gICAgICBUaGlzIGhhcHBlbnMgZXZlbiBhZnRlciBmYWlsaW5nIHRvIGFwcGx5IHBhdGNoZXMgYmVjYXVzZSBvdGhlcndpc2UgXG4gICAgICB5YXJuLmxvY2sgYW5kIHBhY2thZ2UuanNvbiBtaWdodCBnZXQgb3V0IG9mIHN5bmMgd2l0aCBub2RlX21vZHVsZXMsXG4gICAgICB3aGljaCBjYW4gYmUgdmVyeSBjb25mdXNpbmcuXG4gICAgICBcbiAgICAgIC0tZXJyb3Itb24tZmFpbCBpcyAke2NoYWxrLmJvbGQoXCJzd2l0Y2hlZCBvblwiKX0gYnkgZGVmYXVsdCBvbiBDSS5cbiAgICAgIFxuICAgICAgU2VlIGh0dHBzOi8vZ2l0aHViLmNvbS9kczMwMC9wYXRjaC1wYWNrYWdlL2lzc3Vlcy84NiBmb3IgYmFja2dyb3VuZC5cbiAgICAgIFxuICAgICR7Y2hhbGsuYm9sZChcIi0tZXJyb3Itb24td2FyblwiKX1cbiAgICBcbiAgICAgIEZvcmNlcyBwYXRjaC1wYWNrYWdlIHRvIGV4aXQgd2l0aCBjb2RlIDEgYWZ0ZXIgd2FybmluZy5cbiAgICAgIFxuICAgICAgU2VlIGh0dHBzOi8vZ2l0aHViLmNvbS9kczMwMC9wYXRjaC1wYWNrYWdlL2lzc3Vlcy8zMTQgZm9yIGJhY2tncm91bmQuXG5cbiAgICAke2NoYWxrLmJvbGQoXCItLXJldmVyc2VcIil9XG4gICAgICAgIFxuICAgICAgVW4tYXBwbGllcyBhbGwgcGF0Y2hlcy5cblxuICAgICAgTm90ZSB0aGF0IHRoaXMgd2lsbCBmYWlsIGlmIHRoZSBwYXRjaGVkIGZpbGVzIGhhdmUgY2hhbmdlZCBzaW5jZSBiZWluZ1xuICAgICAgcGF0Y2hlZC4gSW4gdGhhdCBjYXNlLCB5b3UnbGwgcHJvYmFibHkgbmVlZCB0byByZS1pbnN0YWxsICdub2RlX21vZHVsZXMnLlxuXG4gICAgICBUaGlzIG9wdGlvbiB3YXMgYWRkZWQgdG8gaGVscCBwZW9wbGUgdXNpbmcgQ2lyY2xlQ0kgYXZvaWQgYW4gaXNzdWUgYXJvdW5kIGNhY2hpbmdcbiAgICAgIGFuZCBwYXRjaCBmaWxlIHVwZGF0ZXMgKGh0dHBzOi8vZ2l0aHViLmNvbS9kczMwMC9wYXRjaC1wYWNrYWdlL2lzc3Vlcy8zNyksXG4gICAgICBidXQgbWlnaHQgYmUgdXNlZnVsIGluIG90aGVyIGNvbnRleHRzIHRvby5cbiAgICAgIFxuXG4gIDIuIENyZWF0aW5nIHBhdGNoIGZpbGVzXG4gID09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgICAke2NoYWxrLmJvbGQoXCJwYXRjaC1wYWNrYWdlXCIpfSA8cGFja2FnZS1uYW1lPiR7Y2hhbGsuaXRhbGljKFxuICAgIFwiWyA8cGFja2FnZS1uYW1lPl1cIixcbiAgKX1cblxuICBXaGVuIGdpdmVuIHBhY2thZ2UgbmFtZXMgYXMgYXJndW1lbnRzLCBwYXRjaC1wYWNrYWdlIHdpbGwgY3JlYXRlIHBhdGNoIGZpbGVzXG4gIGJhc2VkIG9uIGFueSBjaGFuZ2VzIHlvdSd2ZSBtYWRlIHRvIHRoZSB2ZXJzaW9ucyBpbnN0YWxsZWQgYnkgeWFybi9ucG0uXG5cbiAgT3B0aW9uczpcbiAgXG4gICAgJHtjaGFsay5ib2xkKFwiLS1jcmVhdGUtaXNzdWVcIil9XG4gICAgXG4gICAgICAgRm9yIHBhY2thZ2VzIHdob3NlIHNvdXJjZSBpcyBob3N0ZWQgb24gR2l0SHViIHRoaXMgb3B0aW9uIG9wZW5zIGEgd2ViXG4gICAgICAgYnJvd3NlciB3aXRoIGEgZHJhZnQgaXNzdWUgYmFzZWQgb24geW91ciBkaWZmLlxuXG4gICAgJHtjaGFsay5ib2xkKFwiLS11c2UteWFyblwiKX1cblxuICAgICAgICBCeSBkZWZhdWx0LCBwYXRjaC1wYWNrYWdlIGNoZWNrcyB3aGV0aGVyIHlvdSB1c2UgbnBtIG9yIHlhcm4gYmFzZWQgb25cbiAgICAgICAgd2hpY2ggbG9ja2ZpbGUgeW91IGhhdmUuIElmIHlvdSBoYXZlIGJvdGgsIGl0IHVzZXMgbnBtIGJ5IGRlZmF1bHQuXG4gICAgICAgIFNldCB0aGlzIG9wdGlvbiB0byBvdmVycmlkZSB0aGF0IGRlZmF1bHQgYW5kIGFsd2F5cyB1c2UgeWFybi5cblxuICAgICR7Y2hhbGsuYm9sZChcIi0tZXhjbHVkZSA8cmVnZXhwPlwiKX1cblxuICAgICAgICBJZ25vcmUgcGF0aHMgbWF0Y2hpbmcgdGhlIHJlZ2V4cCB3aGVuIGNyZWF0aW5nIHBhdGNoIGZpbGVzLlxuICAgICAgICBQYXRocyBhcmUgcmVsYXRpdmUgdG8gdGhlIHJvb3QgZGlyIG9mIHRoZSBwYWNrYWdlIHRvIGJlIHBhdGNoZWQuXG5cbiAgICAgICAgRGVmYXVsdDogJ3BhY2thZ2VcXFxcLmpzb24kJ1xuXG4gICAgJHtjaGFsay5ib2xkKFwiLS1pbmNsdWRlIDxyZWdleHA+XCIpfVxuXG4gICAgICAgIE9ubHkgY29uc2lkZXIgcGF0aHMgbWF0Y2hpbmcgdGhlIHJlZ2V4cCB3aGVuIGNyZWF0aW5nIHBhdGNoIGZpbGVzLlxuICAgICAgICBQYXRocyBhcmUgcmVsYXRpdmUgdG8gdGhlIHJvb3QgZGlyIG9mIHRoZSBwYWNrYWdlIHRvIGJlIHBhdGNoZWQuXG5cbiAgICAgICAgRGVmYXVsdCAnLionXG5cbiAgICAke2NoYWxrLmJvbGQoXCItLWNhc2Utc2Vuc2l0aXZlLXBhdGgtZmlsdGVyaW5nXCIpfVxuXG4gICAgICAgIE1ha2UgcmVnZXhwcyB1c2VkIGluIC0taW5jbHVkZSBvciAtLWV4Y2x1ZGUgZmlsdGVycyBjYXNlLXNlbnNpdGl2ZS5cbiAgICBcbiAgICAke2NoYWxrLmJvbGQoXCItLXBhdGNoLWRpclwiKX1cblxuICAgICAgICBTcGVjaWZ5IHRoZSBuYW1lIGZvciB0aGUgZGlyZWN0b3J5IGluIHdoaWNoIHRvIHB1dCB0aGUgcGF0Y2ggZmlsZXMuXG5gKVxufVxuIl19