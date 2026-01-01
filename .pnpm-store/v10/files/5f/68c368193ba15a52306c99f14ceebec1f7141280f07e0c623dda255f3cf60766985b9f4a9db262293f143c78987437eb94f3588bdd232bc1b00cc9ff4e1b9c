"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.detectPackageManager = void 0;
const fs_extra_1 = __importDefault(require("fs-extra"));
const path_1 = require("./path");
const chalk_1 = __importDefault(require("chalk"));
const process_1 = __importDefault(require("process"));
const find_yarn_workspace_root_1 = __importDefault(require("find-yarn-workspace-root"));
function printNoYarnLockfileError() {
    console.log(`
${chalk_1.default.red.bold("**ERROR**")} ${chalk_1.default.red(`The --use-yarn option was specified but there is no yarn.lock file`)}
`);
}
function printNoLockfilesError() {
    console.log(`
${chalk_1.default.red.bold("**ERROR**")} ${chalk_1.default.red(`No package-lock.json, npm-shrinkwrap.json, or yarn.lock file.

You must use either npm@>=5, yarn, or npm-shrinkwrap to manage this project's
dependencies.`)}
`);
}
function printSelectingDefaultMessage() {
    console.info(`${chalk_1.default.bold("patch-package")}: you have both yarn.lock and package-lock.json
Defaulting to using ${chalk_1.default.bold("npm")}
You can override this setting by passing --use-yarn or deleting
package-lock.json if you don't need it
`);
}
const detectPackageManager = (appRootPath, overridePackageManager) => {
    const packageLockExists = fs_extra_1.default.existsSync(path_1.join(appRootPath, "package-lock.json"));
    const shrinkWrapExists = fs_extra_1.default.existsSync(path_1.join(appRootPath, "npm-shrinkwrap.json"));
    const yarnLockExists = fs_extra_1.default.existsSync(path_1.join(appRootPath, "yarn.lock"));
    if ((packageLockExists || shrinkWrapExists) && yarnLockExists) {
        if (overridePackageManager) {
            return overridePackageManager;
        }
        else {
            printSelectingDefaultMessage();
            return shrinkWrapExists ? "npm-shrinkwrap" : "npm";
        }
    }
    else if (packageLockExists || shrinkWrapExists) {
        if (overridePackageManager === "yarn") {
            printNoYarnLockfileError();
            process_1.default.exit(1);
        }
        else {
            return shrinkWrapExists ? "npm-shrinkwrap" : "npm";
        }
    }
    else if (yarnLockExists || find_yarn_workspace_root_1.default()) {
        return "yarn";
    }
    else {
        printNoLockfilesError();
        process_1.default.exit(1);
    }
    throw Error();
};
exports.detectPackageManager = detectPackageManager;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZGV0ZWN0UGFja2FnZU1hbmFnZXIuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvZGV0ZWN0UGFja2FnZU1hbmFnZXIudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7O0FBQUEsd0RBQXlCO0FBQ3pCLGlDQUE2QjtBQUM3QixrREFBeUI7QUFDekIsc0RBQTZCO0FBQzdCLHdGQUF3RDtBQUl4RCxTQUFTLHdCQUF3QjtJQUMvQixPQUFPLENBQUMsR0FBRyxDQUFDO0VBQ1osZUFBSyxDQUFDLEdBQUcsQ0FBQyxJQUFJLENBQUMsV0FBVyxDQUFDLElBQUksZUFBSyxDQUFDLEdBQUcsQ0FDdEMsb0VBQW9FLENBQ3JFO0NBQ0YsQ0FBQyxDQUFBO0FBQ0YsQ0FBQztBQUVELFNBQVMscUJBQXFCO0lBQzVCLE9BQU8sQ0FBQyxHQUFHLENBQUM7RUFDWixlQUFLLENBQUMsR0FBRyxDQUFDLElBQUksQ0FBQyxXQUFXLENBQUMsSUFBSSxlQUFLLENBQUMsR0FBRyxDQUN0Qzs7O2NBR1UsQ0FDWDtDQUNGLENBQUMsQ0FBQTtBQUNGLENBQUM7QUFFRCxTQUFTLDRCQUE0QjtJQUNuQyxPQUFPLENBQUMsSUFBSSxDQUNWLEdBQUcsZUFBSyxDQUFDLElBQUksQ0FDWCxlQUFlLENBQ2hCO3NCQUNpQixlQUFLLENBQUMsSUFBSSxDQUFDLEtBQUssQ0FBQzs7O0NBR3RDLENBQ0UsQ0FBQTtBQUNILENBQUM7QUFFTSxNQUFNLG9CQUFvQixHQUFHLENBQ2xDLFdBQW1CLEVBQ25CLHNCQUE2QyxFQUM3QixFQUFFO0lBQ2xCLE1BQU0saUJBQWlCLEdBQUcsa0JBQUUsQ0FBQyxVQUFVLENBQ3JDLFdBQUksQ0FBQyxXQUFXLEVBQUUsbUJBQW1CLENBQUMsQ0FDdkMsQ0FBQTtJQUNELE1BQU0sZ0JBQWdCLEdBQUcsa0JBQUUsQ0FBQyxVQUFVLENBQ3BDLFdBQUksQ0FBQyxXQUFXLEVBQUUscUJBQXFCLENBQUMsQ0FDekMsQ0FBQTtJQUNELE1BQU0sY0FBYyxHQUFHLGtCQUFFLENBQUMsVUFBVSxDQUFDLFdBQUksQ0FBQyxXQUFXLEVBQUUsV0FBVyxDQUFDLENBQUMsQ0FBQTtJQUNwRSxJQUFJLENBQUMsaUJBQWlCLElBQUksZ0JBQWdCLENBQUMsSUFBSSxjQUFjLEVBQUU7UUFDN0QsSUFBSSxzQkFBc0IsRUFBRTtZQUMxQixPQUFPLHNCQUFzQixDQUFBO1NBQzlCO2FBQU07WUFDTCw0QkFBNEIsRUFBRSxDQUFBO1lBQzlCLE9BQU8sZ0JBQWdCLENBQUMsQ0FBQyxDQUFDLGdCQUFnQixDQUFDLENBQUMsQ0FBQyxLQUFLLENBQUE7U0FDbkQ7S0FDRjtTQUFNLElBQUksaUJBQWlCLElBQUksZ0JBQWdCLEVBQUU7UUFDaEQsSUFBSSxzQkFBc0IsS0FBSyxNQUFNLEVBQUU7WUFDckMsd0JBQXdCLEVBQUUsQ0FBQTtZQUMxQixpQkFBTyxDQUFDLElBQUksQ0FBQyxDQUFDLENBQUMsQ0FBQTtTQUNoQjthQUFNO1lBQ0wsT0FBTyxnQkFBZ0IsQ0FBQyxDQUFDLENBQUMsZ0JBQWdCLENBQUMsQ0FBQyxDQUFDLEtBQUssQ0FBQTtTQUNuRDtLQUNGO1NBQU0sSUFBSSxjQUFjLElBQUksa0NBQWlCLEVBQUUsRUFBRTtRQUNoRCxPQUFPLE1BQU0sQ0FBQTtLQUNkO1NBQU07UUFDTCxxQkFBcUIsRUFBRSxDQUFBO1FBQ3ZCLGlCQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFBO0tBQ2hCO0lBQ0QsTUFBTSxLQUFLLEVBQUUsQ0FBQTtBQUNmLENBQUMsQ0FBQTtBQWhDWSxRQUFBLG9CQUFvQix3QkFnQ2hDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IGZzIGZyb20gXCJmcy1leHRyYVwiXG5pbXBvcnQgeyBqb2luIH0gZnJvbSBcIi4vcGF0aFwiXG5pbXBvcnQgY2hhbGsgZnJvbSBcImNoYWxrXCJcbmltcG9ydCBwcm9jZXNzIGZyb20gXCJwcm9jZXNzXCJcbmltcG9ydCBmaW5kV29ya3NwYWNlUm9vdCBmcm9tIFwiZmluZC15YXJuLXdvcmtzcGFjZS1yb290XCJcblxuZXhwb3J0IHR5cGUgUGFja2FnZU1hbmFnZXIgPSBcInlhcm5cIiB8IFwibnBtXCIgfCBcIm5wbS1zaHJpbmt3cmFwXCJcblxuZnVuY3Rpb24gcHJpbnROb1lhcm5Mb2NrZmlsZUVycm9yKCkge1xuICBjb25zb2xlLmxvZyhgXG4ke2NoYWxrLnJlZC5ib2xkKFwiKipFUlJPUioqXCIpfSAke2NoYWxrLnJlZChcbiAgICBgVGhlIC0tdXNlLXlhcm4gb3B0aW9uIHdhcyBzcGVjaWZpZWQgYnV0IHRoZXJlIGlzIG5vIHlhcm4ubG9jayBmaWxlYCxcbiAgKX1cbmApXG59XG5cbmZ1bmN0aW9uIHByaW50Tm9Mb2NrZmlsZXNFcnJvcigpIHtcbiAgY29uc29sZS5sb2coYFxuJHtjaGFsay5yZWQuYm9sZChcIioqRVJST1IqKlwiKX0gJHtjaGFsay5yZWQoXG4gICAgYE5vIHBhY2thZ2UtbG9jay5qc29uLCBucG0tc2hyaW5rd3JhcC5qc29uLCBvciB5YXJuLmxvY2sgZmlsZS5cblxuWW91IG11c3QgdXNlIGVpdGhlciBucG1APj01LCB5YXJuLCBvciBucG0tc2hyaW5rd3JhcCB0byBtYW5hZ2UgdGhpcyBwcm9qZWN0J3NcbmRlcGVuZGVuY2llcy5gLFxuICApfVxuYClcbn1cblxuZnVuY3Rpb24gcHJpbnRTZWxlY3RpbmdEZWZhdWx0TWVzc2FnZSgpIHtcbiAgY29uc29sZS5pbmZvKFxuICAgIGAke2NoYWxrLmJvbGQoXG4gICAgICBcInBhdGNoLXBhY2thZ2VcIixcbiAgICApfTogeW91IGhhdmUgYm90aCB5YXJuLmxvY2sgYW5kIHBhY2thZ2UtbG9jay5qc29uXG5EZWZhdWx0aW5nIHRvIHVzaW5nICR7Y2hhbGsuYm9sZChcIm5wbVwiKX1cbllvdSBjYW4gb3ZlcnJpZGUgdGhpcyBzZXR0aW5nIGJ5IHBhc3NpbmcgLS11c2UteWFybiBvciBkZWxldGluZ1xucGFja2FnZS1sb2NrLmpzb24gaWYgeW91IGRvbid0IG5lZWQgaXRcbmAsXG4gIClcbn1cblxuZXhwb3J0IGNvbnN0IGRldGVjdFBhY2thZ2VNYW5hZ2VyID0gKFxuICBhcHBSb290UGF0aDogc3RyaW5nLFxuICBvdmVycmlkZVBhY2thZ2VNYW5hZ2VyOiBQYWNrYWdlTWFuYWdlciB8IG51bGwsXG4pOiBQYWNrYWdlTWFuYWdlciA9PiB7XG4gIGNvbnN0IHBhY2thZ2VMb2NrRXhpc3RzID0gZnMuZXhpc3RzU3luYyhcbiAgICBqb2luKGFwcFJvb3RQYXRoLCBcInBhY2thZ2UtbG9jay5qc29uXCIpLFxuICApXG4gIGNvbnN0IHNocmlua1dyYXBFeGlzdHMgPSBmcy5leGlzdHNTeW5jKFxuICAgIGpvaW4oYXBwUm9vdFBhdGgsIFwibnBtLXNocmlua3dyYXAuanNvblwiKSxcbiAgKVxuICBjb25zdCB5YXJuTG9ja0V4aXN0cyA9IGZzLmV4aXN0c1N5bmMoam9pbihhcHBSb290UGF0aCwgXCJ5YXJuLmxvY2tcIikpXG4gIGlmICgocGFja2FnZUxvY2tFeGlzdHMgfHwgc2hyaW5rV3JhcEV4aXN0cykgJiYgeWFybkxvY2tFeGlzdHMpIHtcbiAgICBpZiAob3ZlcnJpZGVQYWNrYWdlTWFuYWdlcikge1xuICAgICAgcmV0dXJuIG92ZXJyaWRlUGFja2FnZU1hbmFnZXJcbiAgICB9IGVsc2Uge1xuICAgICAgcHJpbnRTZWxlY3RpbmdEZWZhdWx0TWVzc2FnZSgpXG4gICAgICByZXR1cm4gc2hyaW5rV3JhcEV4aXN0cyA/IFwibnBtLXNocmlua3dyYXBcIiA6IFwibnBtXCJcbiAgICB9XG4gIH0gZWxzZSBpZiAocGFja2FnZUxvY2tFeGlzdHMgfHwgc2hyaW5rV3JhcEV4aXN0cykge1xuICAgIGlmIChvdmVycmlkZVBhY2thZ2VNYW5hZ2VyID09PSBcInlhcm5cIikge1xuICAgICAgcHJpbnROb1lhcm5Mb2NrZmlsZUVycm9yKClcbiAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgIH0gZWxzZSB7XG4gICAgICByZXR1cm4gc2hyaW5rV3JhcEV4aXN0cyA/IFwibnBtLXNocmlua3dyYXBcIiA6IFwibnBtXCJcbiAgICB9XG4gIH0gZWxzZSBpZiAoeWFybkxvY2tFeGlzdHMgfHwgZmluZFdvcmtzcGFjZVJvb3QoKSkge1xuICAgIHJldHVybiBcInlhcm5cIlxuICB9IGVsc2Uge1xuICAgIHByaW50Tm9Mb2NrZmlsZXNFcnJvcigpXG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cbiAgdGhyb3cgRXJyb3IoKVxufVxuIl19