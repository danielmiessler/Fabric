"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getPackageResolution = void 0;
const path_1 = require("./path");
const PackageDetails_1 = require("./PackageDetails");
const detectPackageManager_1 = require("./detectPackageManager");
const fs_extra_1 = require("fs-extra");
const lockfile_1 = require("@yarnpkg/lockfile");
const yaml_1 = __importDefault(require("yaml"));
const find_yarn_workspace_root_1 = __importDefault(require("find-yarn-workspace-root"));
const getPackageVersion_1 = require("./getPackageVersion");
const coerceSemVer_1 = require("./coerceSemVer");
function getPackageResolution({ packageDetails, packageManager, appPath, }) {
    if (packageManager === "yarn") {
        let lockFilePath = "yarn.lock";
        if (!fs_extra_1.existsSync(lockFilePath)) {
            const workspaceRoot = find_yarn_workspace_root_1.default();
            if (!workspaceRoot) {
                throw new Error("Can't find yarn.lock file");
            }
            lockFilePath = path_1.join(workspaceRoot, "yarn.lock");
        }
        if (!fs_extra_1.existsSync(lockFilePath)) {
            throw new Error("Can't find yarn.lock file");
        }
        const lockFileString = fs_extra_1.readFileSync(lockFilePath).toString();
        let appLockFile;
        if (lockFileString.includes("yarn lockfile v1")) {
            const parsedYarnLockFile = lockfile_1.parse(lockFileString);
            if (parsedYarnLockFile.type !== "success") {
                throw new Error("Could not parse yarn v1 lock file");
            }
            else {
                appLockFile = parsedYarnLockFile.object;
            }
        }
        else {
            try {
                appLockFile = yaml_1.default.parse(lockFileString);
            }
            catch (e) {
                console.log(e);
                throw new Error("Could not parse yarn v2 lock file");
            }
        }
        const installedVersion = getPackageVersion_1.getPackageVersion(path_1.join(path_1.resolve(appPath, packageDetails.path), "package.json"));
        const entries = Object.entries(appLockFile).filter(([k, v]) => k.startsWith(packageDetails.name + "@") &&
            // @ts-ignore
            coerceSemVer_1.coerceSemVer(v.version) === coerceSemVer_1.coerceSemVer(installedVersion));
        const resolutions = entries.map(([_, v]) => {
            // @ts-ignore
            return v.resolved;
        });
        if (resolutions.length === 0) {
            throw new Error(`\`${packageDetails.pathSpecifier}\`'s installed version is ${installedVersion} but a lockfile entry for it couldn't be found. Your lockfile is likely to be corrupt or you forgot to reinstall your packages.`);
        }
        if (new Set(resolutions).size !== 1) {
            console.log(`Ambigious lockfile entries for ${packageDetails.pathSpecifier}. Using version ${installedVersion}`);
            return installedVersion;
        }
        if (resolutions[0]) {
            return resolutions[0];
        }
        const resolution = entries[0][0].slice(packageDetails.name.length + 1);
        // resolve relative file path
        if (resolution.startsWith("file:.")) {
            return `file:${path_1.resolve(appPath, resolution.slice("file:".length))}`;
        }
        if (resolution.startsWith("npm:")) {
            return resolution.replace("npm:", "");
        }
        return resolution;
    }
    else {
        const lockfile = require(path_1.join(appPath, packageManager === "npm-shrinkwrap"
            ? "npm-shrinkwrap.json"
            : "package-lock.json"));
        const lockFileStack = [lockfile];
        for (const name of packageDetails.packageNames.slice(0, -1)) {
            const child = lockFileStack[0].dependencies;
            if (child && name in child) {
                lockFileStack.push(child[name]);
            }
        }
        lockFileStack.reverse();
        const relevantStackEntry = lockFileStack.find((entry) => {
            if (entry.dependencies) {
                return entry.dependencies && packageDetails.name in entry.dependencies;
            }
            else if (entry.packages) {
                return entry.packages && packageDetails.path in entry.packages;
            }
            throw new Error("Cannot find dependencies or packages in lockfile");
        });
        const pkg = relevantStackEntry.dependencies
            ? relevantStackEntry.dependencies[packageDetails.name]
            : relevantStackEntry.packages[packageDetails.path];
        return pkg.resolved || pkg.version || pkg.from;
    }
}
exports.getPackageResolution = getPackageResolution;
if (require.main === module) {
    const packageDetails = PackageDetails_1.getPatchDetailsFromCliString(process.argv[2]);
    if (!packageDetails) {
        console.log(`Can't find package ${process.argv[2]}`);
        process.exit(1);
    }
    console.log(getPackageResolution({
        appPath: process.cwd(),
        packageDetails,
        packageManager: detectPackageManager_1.detectPackageManager(process.cwd(), null),
    }));
}
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZ2V0UGFja2FnZVJlc29sdXRpb24uanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvZ2V0UGFja2FnZVJlc29sdXRpb24udHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7O0FBQUEsaUNBQXNDO0FBQ3RDLHFEQUErRTtBQUMvRSxpRUFBNkU7QUFDN0UsdUNBQW1EO0FBQ25ELGdEQUE4RDtBQUM5RCxnREFBdUI7QUFDdkIsd0ZBQXdEO0FBQ3hELDJEQUF1RDtBQUN2RCxpREFBNkM7QUFFN0MsU0FBZ0Isb0JBQW9CLENBQUMsRUFDbkMsY0FBYyxFQUNkLGNBQWMsRUFDZCxPQUFPLEdBS1I7SUFDQyxJQUFJLGNBQWMsS0FBSyxNQUFNLEVBQUU7UUFDN0IsSUFBSSxZQUFZLEdBQUcsV0FBVyxDQUFBO1FBQzlCLElBQUksQ0FBQyxxQkFBVSxDQUFDLFlBQVksQ0FBQyxFQUFFO1lBQzdCLE1BQU0sYUFBYSxHQUFHLGtDQUFpQixFQUFFLENBQUE7WUFDekMsSUFBSSxDQUFDLGFBQWEsRUFBRTtnQkFDbEIsTUFBTSxJQUFJLEtBQUssQ0FBQywyQkFBMkIsQ0FBQyxDQUFBO2FBQzdDO1lBQ0QsWUFBWSxHQUFHLFdBQUksQ0FBQyxhQUFhLEVBQUUsV0FBVyxDQUFDLENBQUE7U0FDaEQ7UUFDRCxJQUFJLENBQUMscUJBQVUsQ0FBQyxZQUFZLENBQUMsRUFBRTtZQUM3QixNQUFNLElBQUksS0FBSyxDQUFDLDJCQUEyQixDQUFDLENBQUE7U0FDN0M7UUFDRCxNQUFNLGNBQWMsR0FBRyx1QkFBWSxDQUFDLFlBQVksQ0FBQyxDQUFDLFFBQVEsRUFBRSxDQUFBO1FBQzVELElBQUksV0FBVyxDQUFBO1FBQ2YsSUFBSSxjQUFjLENBQUMsUUFBUSxDQUFDLGtCQUFrQixDQUFDLEVBQUU7WUFDL0MsTUFBTSxrQkFBa0IsR0FBRyxnQkFBaUIsQ0FBQyxjQUFjLENBQUMsQ0FBQTtZQUM1RCxJQUFJLGtCQUFrQixDQUFDLElBQUksS0FBSyxTQUFTLEVBQUU7Z0JB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SUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUE7SUFDcEUsSUFBSSxDQUFDLGNBQWMsRUFBRTtRQUNuQixPQUFPLENBQUMsR0FBRyxDQUFDLHNCQUFzQixPQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQTtRQUNwRCxPQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFBO0tBQ2hCO0lBQ0QsT0FBTyxDQUFDLEdBQUcsQ0FDVCxvQkFBb0IsQ0FBQztRQUNuQixPQUFPLEVBQUUsT0FBTyxDQUFDLEdBQUcsRUFBRTtRQUN0QixjQUFjO1FBQ2QsY0FBYyxFQUFFLDJDQUFvQixDQUFDLE9BQU8sQ0FBQyxHQUFHLEVBQUUsRUFBRSxJQUFJLENBQUM7S0FDMUQsQ0FBQyxDQUNILENBQUE7Q0FDRiIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGpvaW4sIHJlc29sdmUgfSBmcm9tIFwiLi9wYXRoXCJcbmltcG9ydCB7IFBhY2thZ2VEZXRhaWxzLCBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nIH0gZnJvbSBcIi4vUGFja2FnZURldGFpbHNcIlxuaW1wb3J0IHsgUGFja2FnZU1hbmFnZXIsIGRldGVjdFBhY2thZ2VNYW5hZ2VyIH0gZnJvbSBcIi4vZGV0ZWN0UGFja2FnZU1hbmFnZXJcIlxuaW1wb3J0IHsgcmVhZEZpbGVTeW5jLCBleGlzdHNTeW5jIH0gZnJvbSBcImZzLWV4dHJhXCJcbmltcG9ydCB7IHBhcnNlIGFzIHBhcnNlWWFybkxvY2tGaWxlIH0gZnJvbSBcIkB5YXJucGtnL2xvY2tmaWxlXCJcbmltcG9ydCB5YW1sIGZyb20gXCJ5YW1sXCJcbmltcG9ydCBmaW5kV29ya3NwYWNlUm9vdCBmcm9tIFwiZmluZC15YXJuLXdvcmtzcGFjZS1yb290XCJcbmltcG9ydCB7IGdldFBhY2thZ2VWZXJzaW9uIH0gZnJvbSBcIi4vZ2V0UGFja2FnZVZlcnNpb25cIlxuaW1wb3J0IHsgY29lcmNlU2VtVmVyIH0gZnJvbSBcIi4vY29lcmNlU2VtVmVyXCJcblxuZXhwb3J0IGZ1bmN0aW9uIGdldFBhY2thZ2VSZXNvbHV0aW9uKHtcbiAgcGFja2FnZURldGFpbHMsXG4gIHBhY2thZ2VNYW5hZ2VyLFxuICBhcHBQYXRoLFxufToge1xuICBwYWNrYWdlRGV0YWlsczogUGFja2FnZURldGFpbHNcbiAgcGFja2FnZU1hbmFnZXI6IFBhY2thZ2VNYW5hZ2VyXG4gIGFwcFBhdGg6IHN0cmluZ1xufSkge1xuICBpZiAocGFja2FnZU1hbmFnZXIgPT09IFwieWFyblwiKSB7XG4gICAgbGV0IGxvY2tGaWxlUGF0aCA9IFwieWFybi5sb2NrXCJcbiAgICBpZiAoIWV4aXN0c1N5bmMobG9ja0ZpbGVQYXRoKSkge1xuICAgICAgY29uc3Qgd29ya3NwYWNlUm9vdCA9IGZpbmRXb3Jrc3BhY2VSb290KClcbiAgICAgIGlmICghd29ya3NwYWNlUm9vdCkge1xuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoXCJDYW4ndCBmaW5kIHlhcm4ubG9jayBmaWxlXCIpXG4gICAgICB9XG4gICAgICBsb2NrRmlsZVBhdGggPSBqb2luKHdvcmtzcGFjZVJvb3QsIFwieWFybi5sb2NrXCIpXG4gICAgfVxuICAgIGlmICghZXhpc3RzU3luYyhsb2NrRmlsZVBhdGgpKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoXCJDYW4ndCBmaW5kIHlhcm4ubG9jayBmaWxlXCIpXG4gICAgfVxuICAgIGNvbnN0IGxvY2tGaWxlU3RyaW5nID0gcmVhZEZpbGVTeW5jKGxvY2tGaWxlUGF0aCkudG9TdHJpbmcoKVxuICAgIGxldCBhcHBMb2NrRmlsZVxuICAgIGlmIChsb2NrRmlsZVN0cmluZy5pbmNsdWRlcyhcInlhcm4gbG9ja2ZpbGUgdjFcIikpIHtcbiAgICAgIGNvbnN0IHBhcnNlZFlhcm5Mb2NrRmlsZSA9IHBhcnNlWWFybkxvY2tGaWxlKGxvY2tGaWxlU3RyaW5nKVxuICAgICAgaWYgKHBhcnNlZFlhcm5Mb2NrRmlsZS50eXBlICE9PSBcInN1Y2Nlc3NcIikge1xuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoXCJDb3VsZCBub3QgcGFyc2UgeWFybiB2MSBsb2NrIGZpbGVcIilcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGFwcExvY2tGaWxlID0gcGFyc2VkWWFybkxvY2tGaWxlLm9iamVjdFxuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICB0cnkge1xuICAgICAgICBhcHBMb2NrRmlsZSA9IHlhbWwucGFyc2UobG9ja0ZpbGVTdHJpbmcpXG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKGUpXG4gICAgICAgIHRocm93IG5ldyBFcnJvcihcIkNvdWxkIG5vdCBwYXJzZSB5YXJuIHYyIGxvY2sgZmlsZVwiKVxuICAgICAgfVxuICAgIH1cblxuICAgIGNvbnN0IGluc3RhbGxlZFZlcnNpb24gPSBnZXRQYWNrYWdlVmVyc2lvbihcbiAgICAgIGpvaW4ocmVzb2x2ZShhcHBQYXRoLCBwYWNrYWdlRGV0YWlscy5wYXRoKSwgXCJwYWNrYWdlLmpzb25cIiksXG4gICAgKVxuXG4gICAgY29uc3QgZW50cmllcyA9IE9iamVjdC5lbnRyaWVzKGFwcExvY2tGaWxlKS5maWx0ZXIoXG4gICAgICAoW2ssIHZdKSA9PlxuICAgICAgICBrLnN0YXJ0c1dpdGgocGFja2FnZURldGFpbHMubmFtZSArIFwiQFwiKSAmJlxuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIGNvZXJjZVNlbVZlcih2LnZlcnNpb24pID09PSBjb2VyY2VTZW1WZXIoaW5zdGFsbGVkVmVyc2lvbiksXG4gICAgKVxuXG4gICAgY29uc3QgcmVzb2x1dGlvbnMgPSBlbnRyaWVzLm1hcCgoW18sIHZdKSA9PiB7XG4gICAgICAvLyBAdHMtaWdub3JlXG4gICAgICByZXR1cm4gdi5yZXNvbHZlZFxuICAgIH0pXG5cbiAgICBpZiAocmVzb2x1dGlvbnMubGVuZ3RoID09PSAwKSB7XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoXG4gICAgICAgIGBcXGAke3BhY2thZ2VEZXRhaWxzLnBhdGhTcGVjaWZpZXJ9XFxgJ3MgaW5zdGFsbGVkIHZlcnNpb24gaXMgJHtpbnN0YWxsZWRWZXJzaW9ufSBidXQgYSBsb2NrZmlsZSBlbnRyeSBmb3IgaXQgY291bGRuJ3QgYmUgZm91bmQuIFlvdXIgbG9ja2ZpbGUgaXMgbGlrZWx5IHRvIGJlIGNvcnJ1cHQgb3IgeW91IGZvcmdvdCB0byByZWluc3RhbGwgeW91ciBwYWNrYWdlcy5gLFxuICAgICAgKVxuICAgIH1cblxuICAgIGlmIChuZXcgU2V0KHJlc29sdXRpb25zKS5zaXplICE9PSAxKSB7XG4gICAgICBjb25zb2xlLmxvZyhcbiAgICAgICAgYEFtYmlnaW91cyBsb2NrZmlsZSBlbnRyaWVzIGZvciAke3BhY2thZ2VEZXRhaWxzLnBhdGhTcGVjaWZpZXJ9LiBVc2luZyB2ZXJzaW9uICR7aW5zdGFsbGVkVmVyc2lvbn1gLFxuICAgICAgKVxuICAgICAgcmV0dXJuIGluc3RhbGxlZFZlcnNpb25cbiAgICB9XG5cbiAgICBpZiAocmVzb2x1dGlvbnNbMF0pIHtcbiAgICAgIHJldHVybiByZXNvbHV0aW9uc1swXVxuICAgIH1cblxuICAgIGNvbnN0IHJlc29sdXRpb24gPSBlbnRyaWVzWzBdWzBdLnNsaWNlKHBhY2thZ2VEZXRhaWxzLm5hbWUubGVuZ3RoICsgMSlcblxuICAgIC8vIHJlc29sdmUgcmVsYXRpdmUgZmlsZSBwYXRoXG4gICAgaWYgKHJlc29sdXRpb24uc3RhcnRzV2l0aChcImZpbGU6LlwiKSkge1xuICAgICAgcmV0dXJuIGBmaWxlOiR7cmVzb2x2ZShhcHBQYXRoLCByZXNvbHV0aW9uLnNsaWNlKFwiZmlsZTpcIi5sZW5ndGgpKX1gXG4gICAgfVxuXG4gICAgaWYgKHJlc29sdXRpb24uc3RhcnRzV2l0aChcIm5wbTpcIikpIHtcbiAgICAgIHJldHVybiByZXNvbHV0aW9uLnJlcGxhY2UoXCJucG06XCIsIFwiXCIpXG4gICAgfVxuXG4gICAgcmV0dXJuIHJlc29sdXRpb25cbiAgfSBlbHNlIHtcbiAgICBjb25zdCBsb2NrZmlsZSA9IHJlcXVpcmUoam9pbihcbiAgICAgIGFwcFBhdGgsXG4gICAgICBwYWNrYWdlTWFuYWdlciA9PT0gXCJucG0tc2hyaW5rd3JhcFwiXG4gICAgICAgID8gXCJucG0tc2hyaW5rd3JhcC5qc29uXCJcbiAgICAgICAgOiBcInBhY2thZ2UtbG9jay5qc29uXCIsXG4gICAgKSlcbiAgICBjb25zdCBsb2NrRmlsZVN0YWNrID0gW2xvY2tmaWxlXVxuICAgIGZvciAoY29uc3QgbmFtZSBvZiBwYWNrYWdlRGV0YWlscy5wYWNrYWdlTmFtZXMuc2xpY2UoMCwgLTEpKSB7XG4gICAgICBjb25zdCBjaGlsZCA9IGxvY2tGaWxlU3RhY2tbMF0uZGVwZW5kZW5jaWVzXG4gICAgICBpZiAoY2hpbGQgJiYgbmFtZSBpbiBjaGlsZCkge1xuICAgICAgICBsb2NrRmlsZVN0YWNrLnB1c2goY2hpbGRbbmFtZV0pXG4gICAgICB9XG4gICAgfVxuICAgIGxvY2tGaWxlU3RhY2sucmV2ZXJzZSgpXG4gICAgY29uc3QgcmVsZXZhbnRTdGFja0VudHJ5ID0gbG9ja0ZpbGVTdGFjay5maW5kKChlbnRyeSkgPT4ge1xuICAgICAgaWYgKGVudHJ5LmRlcGVuZGVuY2llcykge1xuICAgICAgICByZXR1cm4gZW50cnkuZGVwZW5kZW5jaWVzICYmIHBhY2thZ2VEZXRhaWxzLm5hbWUgaW4gZW50cnkuZGVwZW5kZW5jaWVzXG4gICAgICB9IGVsc2UgaWYgKGVudHJ5LnBhY2thZ2VzKSB7XG4gICAgICAgIHJldHVybiBlbnRyeS5wYWNrYWdlcyAmJiBwYWNrYWdlRGV0YWlscy5wYXRoIGluIGVudHJ5LnBhY2thZ2VzXG4gICAgICB9XG4gICAgICB0aHJvdyBuZXcgRXJyb3IoXCJDYW5ub3QgZmluZCBkZXBlbmRlbmNpZXMgb3IgcGFja2FnZXMgaW4gbG9ja2ZpbGVcIilcbiAgICB9KVxuICAgIGNvbnN0IHBrZyA9IHJlbGV2YW50U3RhY2tFbnRyeS5kZXBlbmRlbmNpZXNcbiAgICAgID8gcmVsZXZhbnRTdGFja0VudHJ5LmRlcGVuZGVuY2llc1twYWNrYWdlRGV0YWlscy5uYW1lXVxuICAgICAgOiByZWxldmFudFN0YWNrRW50cnkucGFja2FnZXNbcGFja2FnZURldGFpbHMucGF0aF1cbiAgICByZXR1cm4gcGtnLnJlc29sdmVkIHx8IHBrZy52ZXJzaW9uIHx8IHBrZy5mcm9tXG4gIH1cbn1cblxuaWYgKHJlcXVpcmUubWFpbiA9PT0gbW9kdWxlKSB7XG4gIGNvbnN0IHBhY2thZ2VEZXRhaWxzID0gZ2V0UGF0Y2hEZXRhaWxzRnJvbUNsaVN0cmluZyhwcm9jZXNzLmFyZ3ZbMl0pXG4gIGlmICghcGFja2FnZURldGFpbHMpIHtcbiAgICBjb25zb2xlLmxvZyhgQ2FuJ3QgZmluZCBwYWNrYWdlICR7cHJvY2Vzcy5hcmd2WzJdfWApXG4gICAgcHJvY2Vzcy5leGl0KDEpXG4gIH1cbiAgY29uc29sZS5sb2coXG4gICAgZ2V0UGFja2FnZVJlc29sdXRpb24oe1xuICAgICAgYXBwUGF0aDogcHJvY2Vzcy5jd2QoKSxcbiAgICAgIHBhY2thZ2VEZXRhaWxzLFxuICAgICAgcGFja2FnZU1hbmFnZXI6IGRldGVjdFBhY2thZ2VNYW5hZ2VyKHByb2Nlc3MuY3dkKCksIG51bGwpLFxuICAgIH0pLFxuICApXG59XG4iXX0=