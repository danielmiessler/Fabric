"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.makeRegExp = void 0;
const chalk_1 = __importDefault(require("chalk"));
const makeRegExp = (reString, name, defaultValue, caseSensitive) => {
    if (!reString) {
        return defaultValue;
    }
    else {
        try {
            return new RegExp(reString, caseSensitive ? "" : "i");
        }
        catch (_) {
            console.log(`${chalk_1.default.red.bold("***ERROR***")}
Invalid format for option --${name}

  Unable to convert the string ${JSON.stringify(reString)} to a regular expression.
`);
            process.exit(1);
            return /unreachable/;
        }
    }
};
exports.makeRegExp = makeRegExp;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibWFrZVJlZ0V4cC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9tYWtlUmVnRXhwLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7OztBQUFBLGtEQUF5QjtBQUVsQixNQUFNLFVBQVUsR0FBRyxDQUN4QixRQUFnQixFQUNoQixJQUFZLEVBQ1osWUFBb0IsRUFDcEIsYUFBc0IsRUFDZCxFQUFFO0lBQ1YsSUFBSSxDQUFDLFFBQVEsRUFBRTtRQUNiLE9BQU8sWUFBWSxDQUFBO0tBQ3BCO1NBQU07UUFDTCxJQUFJO1lBQ0YsT0FBTyxJQUFJLE1BQU0sQ0FBQyxRQUFRLEVBQUUsYUFBYSxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQyxDQUFDLEdBQUcsQ0FBQyxDQUFBO1NBQ3REO1FBQUMsT0FBTyxDQUFDLEVBQUU7WUFDVixPQUFPLENBQUMsR0FBRyxDQUFDLEdBQUcsZUFBSyxDQUFDLEdBQUcsQ0FBQyxJQUFJLENBQUMsYUFBYSxDQUFDOzhCQUNwQixJQUFJOztpQ0FFRCxJQUFJLENBQUMsU0FBUyxDQUMzQyxRQUFRLENBQ1Q7Q0FDRixDQUFDLENBQUE7WUFFSSxPQUFPLENBQUMsSUFBSSxDQUFDLENBQUMsQ0FBQyxDQUFBO1lBQ2YsT0FBTyxhQUFhLENBQUE7U0FDckI7S0FDRjtBQUNILENBQUMsQ0FBQTtBQXhCWSxRQUFBLFVBQVUsY0F3QnRCIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IGNoYWxrIGZyb20gXCJjaGFsa1wiXG5cbmV4cG9ydCBjb25zdCBtYWtlUmVnRXhwID0gKFxuICByZVN0cmluZzogc3RyaW5nLFxuICBuYW1lOiBzdHJpbmcsXG4gIGRlZmF1bHRWYWx1ZTogUmVnRXhwLFxuICBjYXNlU2Vuc2l0aXZlOiBib29sZWFuLFxuKTogUmVnRXhwID0+IHtcbiAgaWYgKCFyZVN0cmluZykge1xuICAgIHJldHVybiBkZWZhdWx0VmFsdWVcbiAgfSBlbHNlIHtcbiAgICB0cnkge1xuICAgICAgcmV0dXJuIG5ldyBSZWdFeHAocmVTdHJpbmcsIGNhc2VTZW5zaXRpdmUgPyBcIlwiIDogXCJpXCIpXG4gICAgfSBjYXRjaCAoXykge1xuICAgICAgY29uc29sZS5sb2coYCR7Y2hhbGsucmVkLmJvbGQoXCIqKipFUlJPUioqKlwiKX1cbkludmFsaWQgZm9ybWF0IGZvciBvcHRpb24gLS0ke25hbWV9XG5cbiAgVW5hYmxlIHRvIGNvbnZlcnQgdGhlIHN0cmluZyAke0pTT04uc3RyaW5naWZ5KFxuICAgIHJlU3RyaW5nLFxuICApfSB0byBhIHJlZ3VsYXIgZXhwcmVzc2lvbi5cbmApXG5cbiAgICAgIHByb2Nlc3MuZXhpdCgxKVxuICAgICAgcmV0dXJuIC91bnJlYWNoYWJsZS9cbiAgICB9XG4gIH1cbn1cbiJdfQ==