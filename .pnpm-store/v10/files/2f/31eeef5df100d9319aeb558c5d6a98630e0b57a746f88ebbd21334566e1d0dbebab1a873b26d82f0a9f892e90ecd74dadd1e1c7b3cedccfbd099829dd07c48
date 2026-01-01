"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.spawnSafeSync = void 0;
const cross_spawn_1 = require("cross-spawn");
const defaultOptions = {
    logStdErrOnError: true,
    throwOnError: true,
};
const spawnSafeSync = (command, args, options) => {
    const mergedOptions = Object.assign({}, defaultOptions, options);
    const result = cross_spawn_1.sync(command, args, options);
    if (result.error || result.status !== 0) {
        if (mergedOptions.logStdErrOnError) {
            if (result.stderr) {
                console.log(result.stderr.toString());
            }
            else if (result.error) {
                console.log(result.error);
            }
        }
        if (mergedOptions.throwOnError) {
            throw result;
        }
    }
    return result;
};
exports.spawnSafeSync = spawnSafeSync;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoic3Bhd25TYWZlLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3NwYXduU2FmZS50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7QUFBQSw2Q0FBK0M7QUFTL0MsTUFBTSxjQUFjLEdBQXFCO0lBQ3ZDLGdCQUFnQixFQUFFLElBQUk7SUFDdEIsWUFBWSxFQUFFLElBQUk7Q0FDbkIsQ0FBQTtBQUVNLE1BQU0sYUFBYSxHQUFHLENBQzNCLE9BQWUsRUFDZixJQUFlLEVBQ2YsT0FBMEIsRUFDMUIsRUFBRTtJQUNGLE1BQU0sYUFBYSxHQUFHLE1BQU0sQ0FBQyxNQUFNLENBQUMsRUFBRSxFQUFFLGNBQWMsRUFBRSxPQUFPLENBQUMsQ0FBQTtJQUNoRSxNQUFNLE1BQU0sR0FBRyxrQkFBUyxDQUFDLE9BQU8sRUFBRSxJQUFJLEVBQUUsT0FBTyxDQUFDLENBQUE7SUFDaEQsSUFBSSxNQUFNLENBQUMsS0FBSyxJQUFJLE1BQU0sQ0FBQyxNQUFNLEtBQUssQ0FBQyxFQUFFO1FBQ3ZDLElBQUksYUFBYSxDQUFDLGdCQUFnQixFQUFFO1lBQ2xDLElBQUksTUFBTSxDQUFDLE1BQU0sRUFBRTtnQkFDakIsT0FBTyxDQUFDLEdBQUcsQ0FBQyxNQUFNLENBQUMsTUFBTSxDQUFDLFFBQVEsRUFBRSxDQUFDLENBQUE7YUFDdEM7aUJBQU0sSUFBSSxNQUFNLENBQUMsS0FBSyxFQUFFO2dCQUN2QixPQUFPLENBQUMsR0FBRyxDQUFDLE1BQU0sQ0FBQyxLQUFLLENBQUMsQ0FBQTthQUMxQjtTQUNGO1FBQ0QsSUFBSSxhQUFhLENBQUMsWUFBWSxFQUFFO1lBQzlCLE1BQU0sTUFBTSxDQUFBO1NBQ2I7S0FDRjtJQUNELE9BQU8sTUFBTSxDQUFBO0FBQ2YsQ0FBQyxDQUFBO0FBcEJZLFFBQUEsYUFBYSxpQkFvQnpCIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgc3luYyBhcyBzcGF3blN5bmMgfSBmcm9tIFwiY3Jvc3Mtc3Bhd25cIlxuaW1wb3J0IHsgU3Bhd25PcHRpb25zIH0gZnJvbSBcImNoaWxkX3Byb2Nlc3NcIlxuXG5leHBvcnQgaW50ZXJmYWNlIFNwYXduU2FmZU9wdGlvbnMgZXh0ZW5kcyBTcGF3bk9wdGlvbnMge1xuICB0aHJvd09uRXJyb3I/OiBib29sZWFuXG4gIGxvZ1N0ZEVyck9uRXJyb3I/OiBib29sZWFuXG4gIG1heEJ1ZmZlcj86IG51bWJlclxufVxuXG5jb25zdCBkZWZhdWx0T3B0aW9uczogU3Bhd25TYWZlT3B0aW9ucyA9IHtcbiAgbG9nU3RkRXJyT25FcnJvcjogdHJ1ZSxcbiAgdGhyb3dPbkVycm9yOiB0cnVlLFxufVxuXG5leHBvcnQgY29uc3Qgc3Bhd25TYWZlU3luYyA9IChcbiAgY29tbWFuZDogc3RyaW5nLFxuICBhcmdzPzogc3RyaW5nW10sXG4gIG9wdGlvbnM/OiBTcGF3blNhZmVPcHRpb25zLFxuKSA9PiB7XG4gIGNvbnN0IG1lcmdlZE9wdGlvbnMgPSBPYmplY3QuYXNzaWduKHt9LCBkZWZhdWx0T3B0aW9ucywgb3B0aW9ucylcbiAgY29uc3QgcmVzdWx0ID0gc3Bhd25TeW5jKGNvbW1hbmQsIGFyZ3MsIG9wdGlvbnMpXG4gIGlmIChyZXN1bHQuZXJyb3IgfHwgcmVzdWx0LnN0YXR1cyAhPT0gMCkge1xuICAgIGlmIChtZXJnZWRPcHRpb25zLmxvZ1N0ZEVyck9uRXJyb3IpIHtcbiAgICAgIGlmIChyZXN1bHQuc3RkZXJyKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKHJlc3VsdC5zdGRlcnIudG9TdHJpbmcoKSlcbiAgICAgIH0gZWxzZSBpZiAocmVzdWx0LmVycm9yKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKHJlc3VsdC5lcnJvcilcbiAgICAgIH1cbiAgICB9XG4gICAgaWYgKG1lcmdlZE9wdGlvbnMudGhyb3dPbkVycm9yKSB7XG4gICAgICB0aHJvdyByZXN1bHRcbiAgICB9XG4gIH1cbiAgcmV0dXJuIHJlc3VsdFxufVxuIl19