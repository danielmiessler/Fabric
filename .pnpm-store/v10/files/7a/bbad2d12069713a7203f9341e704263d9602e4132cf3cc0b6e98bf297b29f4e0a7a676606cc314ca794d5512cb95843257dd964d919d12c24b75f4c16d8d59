"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getAppRootPath = void 0;
const path_1 = require("./path");
const process_1 = __importDefault(require("process"));
const fs_extra_1 = require("fs-extra");
const getAppRootPath = () => {
    let cwd = process_1.default.cwd();
    while (!fs_extra_1.existsSync(path_1.join(cwd, "package.json"))) {
        const up = path_1.resolve(cwd, "../");
        if (up === cwd) {
            throw new Error("no package.json found for this project");
        }
        cwd = up;
    }
    return cwd;
};
exports.getAppRootPath = getAppRootPath;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZ2V0QXBwUm9vdFBhdGguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvZ2V0QXBwUm9vdFBhdGgudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7O0FBQUEsaUNBQXNDO0FBQ3RDLHNEQUE2QjtBQUM3Qix1Q0FBcUM7QUFFOUIsTUFBTSxjQUFjLEdBQUcsR0FBVyxFQUFFO0lBQ3pDLElBQUksR0FBRyxHQUFHLGlCQUFPLENBQUMsR0FBRyxFQUFFLENBQUE7SUFDdkIsT0FBTyxDQUFDLHFCQUFVLENBQUMsV0FBSSxDQUFDLEdBQUcsRUFBRSxjQUFjLENBQUMsQ0FBQyxFQUFFO1FBQzdDLE1BQU0sRUFBRSxHQUFHLGNBQU8sQ0FBQyxHQUFHLEVBQUUsS0FBSyxDQUFDLENBQUE7UUFDOUIsSUFBSSxFQUFFLEtBQUssR0FBRyxFQUFFO1lBQ2QsTUFBTSxJQUFJLEtBQUssQ0FBQyx3Q0FBd0MsQ0FBQyxDQUFBO1NBQzFEO1FBQ0QsR0FBRyxHQUFHLEVBQUUsQ0FBQTtLQUNUO0lBQ0QsT0FBTyxHQUFHLENBQUE7QUFDWixDQUFDLENBQUE7QUFWWSxRQUFBLGNBQWMsa0JBVTFCIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgam9pbiwgcmVzb2x2ZSB9IGZyb20gXCIuL3BhdGhcIlxuaW1wb3J0IHByb2Nlc3MgZnJvbSBcInByb2Nlc3NcIlxuaW1wb3J0IHsgZXhpc3RzU3luYyB9IGZyb20gXCJmcy1leHRyYVwiXG5cbmV4cG9ydCBjb25zdCBnZXRBcHBSb290UGF0aCA9ICgpOiBzdHJpbmcgPT4ge1xuICBsZXQgY3dkID0gcHJvY2Vzcy5jd2QoKVxuICB3aGlsZSAoIWV4aXN0c1N5bmMoam9pbihjd2QsIFwicGFja2FnZS5qc29uXCIpKSkge1xuICAgIGNvbnN0IHVwID0gcmVzb2x2ZShjd2QsIFwiLi4vXCIpXG4gICAgaWYgKHVwID09PSBjd2QpIHtcbiAgICAgIHRocm93IG5ldyBFcnJvcihcIm5vIHBhY2thZ2UuanNvbiBmb3VuZCBmb3IgdGhpcyBwcm9qZWN0XCIpXG4gICAgfVxuICAgIGN3ZCA9IHVwXG4gIH1cbiAgcmV0dXJuIGN3ZFxufVxuIl19