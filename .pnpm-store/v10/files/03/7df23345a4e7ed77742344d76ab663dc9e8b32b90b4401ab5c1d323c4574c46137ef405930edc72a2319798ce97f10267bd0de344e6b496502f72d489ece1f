"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reversePatch = void 0;
const parse_1 = require("./parse");
const assertNever_1 = require("../assertNever");
function reverseHunk(hunk) {
    const header = {
        original: hunk.header.patched,
        patched: hunk.header.original,
    };
    const parts = [];
    for (const part of hunk.parts) {
        switch (part.type) {
            case "context":
                parts.push(part);
                break;
            case "deletion":
                parts.push({
                    type: "insertion",
                    lines: part.lines,
                    noNewlineAtEndOfFile: part.noNewlineAtEndOfFile,
                });
                break;
            case "insertion":
                parts.push({
                    type: "deletion",
                    lines: part.lines,
                    noNewlineAtEndOfFile: part.noNewlineAtEndOfFile,
                });
                break;
            default:
                assertNever_1.assertNever(part.type);
        }
    }
    // swap insertions and deletions over so deletions always come first
    for (let i = 0; i < parts.length - 1; i++) {
        if (parts[i].type === "insertion" && parts[i + 1].type === "deletion") {
            const tmp = parts[i];
            parts[i] = parts[i + 1];
            parts[i + 1] = tmp;
            i += 1;
        }
    }
    const result = {
        header,
        parts,
        source: hunk.source,
    };
    parse_1.verifyHunkIntegrity(result);
    return result;
}
function reversePatchPart(part) {
    switch (part.type) {
        case "file creation":
            return {
                type: "file deletion",
                path: part.path,
                hash: part.hash,
                hunk: part.hunk && reverseHunk(part.hunk),
                mode: part.mode,
            };
        case "file deletion":
            return {
                type: "file creation",
                path: part.path,
                hunk: part.hunk && reverseHunk(part.hunk),
                mode: part.mode,
                hash: part.hash,
            };
        case "rename":
            return {
                type: "rename",
                fromPath: part.toPath,
                toPath: part.fromPath,
            };
        case "patch":
            return {
                type: "patch",
                path: part.path,
                hunks: part.hunks.map(reverseHunk),
                beforeHash: part.afterHash,
                afterHash: part.beforeHash,
            };
        case "mode change":
            return {
                type: "mode change",
                path: part.path,
                newMode: part.oldMode,
                oldMode: part.newMode,
            };
    }
}
const reversePatch = (patch) => {
    return patch.map(reversePatchPart).reverse();
};
exports.reversePatch = reversePatch;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicmV2ZXJzZS5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uLy4uL3NyYy9wYXRjaC9yZXZlcnNlLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7OztBQUFBLG1DQU1nQjtBQUNoQixnREFBNEM7QUFFNUMsU0FBUyxXQUFXLENBQUMsSUFBVTtJQUM3QixNQUFNLE1BQU0sR0FBZTtRQUN6QixRQUFRLEVBQUUsSUFBSSxDQUFDLE1BQU0sQ0FBQyxPQUFPO1FBQzdCLE9BQU8sRUFBRSxJQUFJLENBQUMsTUFBTSxDQUFDLFFBQVE7S0FDOUIsQ0FBQTtJQUNELE1BQU0sS0FBSyxHQUFrQixFQUFFLENBQUE7SUFFL0IsS0FBSyxNQUFNLElBQUksSUFBSSxJQUFJLENBQUMsS0FBSyxFQUFFO1FBQzdCLFFBQVEsSUFBSSxDQUFDLElBQUksRUFBRTtZQUNqQixLQUFLLFNBQVM7Z0JBQ1osS0FBSyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUMsQ0FBQTtnQkFDaEIsTUFBSztZQUNQLEtBQUssVUFBVTtnQkFDYixLQUFLLENBQUMsSUFBSSxDQUFDO29CQUNULElBQUksRUFBRSxXQUFXO29CQUNqQixLQUFLLEVBQUUsSUFBSSxDQUFDLEtBQUs7b0JBQ2pCLG9CQUFvQixFQUFFLElBQUksQ0FBQyxvQkFBb0I7aUJBQ2hELENBQUMsQ0FBQTtnQkFDRixNQUFLO1lBQ1AsS0FBSyxXQUFXO2dCQUNkLEtBQUssQ0FBQyxJQUFJLENBQUM7b0JBQ1QsSUFBSSxFQUFFLFVBQVU7b0JBQ2hCLEtBQUssRUFBRSxJQUFJLENBQUMsS0FBSztvQkFDakIsb0JBQW9CLEVBQUUsSUFBSSxDQUFDLG9CQUFvQjtpQkFDaEQsQ0FBQyxDQUFBO2dCQUNGLE1BQUs7WUFDUDtnQkFDRSx5QkFBVyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUMsQ0FBQTtTQUN6QjtLQUNGO0lBRUQsb0VBQW9FO0lBQ3BFLEtBQUssSUFBSSxDQUFDLEdBQUcsQ0FBQyxFQUFFLENBQUMsR0FBRyxLQUFLLENBQUMsTUFBTSxHQUFHLENBQUMsRUFBRSxDQUFDLEVBQUUsRUFBRTtRQUN6QyxJQUFJLEtBQUssQ0FBQyxDQUFDLENBQUMsQ0FBQyxJQUFJLEtBQUssV0FBVyxJQUFJLEtBQUssQ0FBQyxDQUFDLEdBQUcsQ0FBQyxDQUFDLENBQUMsSUFBSSxLQUFLLFVBQVUsRUFBRTtZQUNyRSxNQUFNLEdBQUcsR0FBRyxLQUFLLENBQUMsQ0FBQyxDQUFDLENBQUE7WUFDcEIsS0FBSyxDQUFDLENBQUMsQ0FBQyxHQUFHLEtBQUssQ0FBQyxDQUFDLEdBQUcsQ0FBQyxDQUFDLENBQUE7WUFDdkIsS0FBSyxDQUFDLENBQUMsR0FBRyxDQUFDLENBQUMsR0FBRyxHQUFHLENBQUE7WUFDbEIsQ0FBQyxJQUFJLENBQUMsQ0FBQTtTQUNQO0tBQ0Y7SUFFRCxNQUFNLE1BQU0sR0FBUztRQUNuQixNQUFNO1FBQ04sS0FBSztRQUNMLE1BQU0sRUFBRSxJQUFJLENBQUMsTUFBTTtLQUNwQixDQUFBO0lBRUQsMkJBQW1CLENBQUMsTUFBTSxDQUFDLENBQUE7SUFFM0IsT0FBTyxNQUFNLENBQUE7QUFDZixDQUFDO0FBRUQsU0FBUyxnQkFBZ0IsQ0FBQyxJQUFtQjtJQUMzQyxRQUFRLElBQUksQ0FBQyxJQUFJLEVBQUU7UUFDakIsS0FBSyxlQUFlO1lBQ2xCLE9BQU87Z0JBQ0wsSUFBSSxFQUFFLGVBQWU7Z0JBQ3JCLElBQUksRUFBRSxJQUFJLENBQUMsSUFBSTtnQkFDZixJQUFJLEVBQUUsSUFBSSxDQUFDLElBQUk7Z0JBQ2YsSUFBSSxFQUFFLElBQUksQ0FBQyxJQUFJLElBQUksV0FBVyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUM7Z0JBQ3pDLElBQUksRUFBRSxJQUFJLENBQUMsSUFBSTthQUNoQixDQUFBO1FBQ0gsS0FBSyxlQUFlO1lBQ2xCLE9BQU87Z0JBQ0wsSUFBSSxFQUFFLGVBQWU7Z0JBQ3JCLElBQUksRUFBRSxJQUFJLENBQUMsSUFBSTtnQkFDZixJQUFJLEVBQUUsSUFBSSxDQUFDLElBQUksSUFBSSxXQUFXLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQztnQkFDekMsSUFBSSxFQUFFLElBQUksQ0FBQyxJQUFJO2dCQUNmLElBQUksRUFBRSxJQUFJLENBQUMsSUFBSTthQUNoQixDQUFBO1FBQ0gsS0FBSyxRQUFRO1lBQ1gsT0FBTztnQkFDTCxJQUFJLEVBQUUsUUFBUTtnQkFDZCxRQUFRLEVBQUUsSUFBSSxDQUFDLE1BQU07Z0JBQ3JCLE1BQU0sRUFBRSxJQUFJLENBQUMsUUFBUTthQUN0QixDQUFBO1FBQ0gsS0FBSyxPQUFPO1lBQ1YsT0FBTztnQkFDTCxJQUFJLEVBQUUsT0FBTztnQkFDYixJQUFJLEVBQUUsSUFBSSxDQUFDLElBQUk7Z0JBQ2YsS0FBSyxFQUFFLElBQUksQ0FBQyxLQUFLLENBQUMsR0FBRyxDQUFDLFdBQVcsQ0FBQztnQkFDbEMsVUFBVSxFQUFFLElBQUksQ0FBQyxTQUFTO2dCQUMxQixTQUFTLEVBQUUsSUFBSSxDQUFDLFVBQVU7YUFDM0IsQ0FBQTtRQUNILEtBQUssYUFBYTtZQUNoQixPQUFPO2dCQUNMLElBQUksRUFBRSxhQUFhO2dCQUNuQixJQUFJLEVBQUUsSUFBSSxDQUFDLElBQUk7Z0JBQ2YsT0FBTyxFQUFFLElBQUksQ0FBQyxPQUFPO2dCQUNyQixPQUFPLEVBQUUsSUFBSSxDQUFDLE9BQU87YUFDdEIsQ0FBQTtLQUNKO0FBQ0gsQ0FBQztBQUVNLE1BQU0sWUFBWSxHQUFHLENBQUMsS0FBc0IsRUFBbUIsRUFBRTtJQUN0RSxPQUFPLEtBQUssQ0FBQyxHQUFHLENBQUMsZ0JBQWdCLENBQUMsQ0FBQyxPQUFPLEVBQUUsQ0FBQTtBQUM5QyxDQUFDLENBQUE7QUFGWSxRQUFBLFlBQVksZ0JBRXhCIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHtcbiAgUGFyc2VkUGF0Y2hGaWxlLFxuICBQYXRjaEZpbGVQYXJ0LFxuICBIdW5rLFxuICBIdW5rSGVhZGVyLFxuICB2ZXJpZnlIdW5rSW50ZWdyaXR5LFxufSBmcm9tIFwiLi9wYXJzZVwiXG5pbXBvcnQgeyBhc3NlcnROZXZlciB9IGZyb20gXCIuLi9hc3NlcnROZXZlclwiXG5cbmZ1bmN0aW9uIHJldmVyc2VIdW5rKGh1bms6IEh1bmspOiBIdW5rIHtcbiAgY29uc3QgaGVhZGVyOiBIdW5rSGVhZGVyID0ge1xuICAgIG9yaWdpbmFsOiBodW5rLmhlYWRlci5wYXRjaGVkLFxuICAgIHBhdGNoZWQ6IGh1bmsuaGVhZGVyLm9yaWdpbmFsLFxuICB9XG4gIGNvbnN0IHBhcnRzOiBIdW5rW1wicGFydHNcIl0gPSBbXVxuXG4gIGZvciAoY29uc3QgcGFydCBvZiBodW5rLnBhcnRzKSB7XG4gICAgc3dpdGNoIChwYXJ0LnR5cGUpIHtcbiAgICAgIGNhc2UgXCJjb250ZXh0XCI6XG4gICAgICAgIHBhcnRzLnB1c2gocGFydClcbiAgICAgICAgYnJlYWtcbiAgICAgIGNhc2UgXCJkZWxldGlvblwiOlxuICAgICAgICBwYXJ0cy5wdXNoKHtcbiAgICAgICAgICB0eXBlOiBcImluc2VydGlvblwiLFxuICAgICAgICAgIGxpbmVzOiBwYXJ0LmxpbmVzLFxuICAgICAgICAgIG5vTmV3bGluZUF0RW5kT2ZGaWxlOiBwYXJ0Lm5vTmV3bGluZUF0RW5kT2ZGaWxlLFxuICAgICAgICB9KVxuICAgICAgICBicmVha1xuICAgICAgY2FzZSBcImluc2VydGlvblwiOlxuICAgICAgICBwYXJ0cy5wdXNoKHtcbiAgICAgICAgICB0eXBlOiBcImRlbGV0aW9uXCIsXG4gICAgICAgICAgbGluZXM6IHBhcnQubGluZXMsXG4gICAgICAgICAgbm9OZXdsaW5lQXRFbmRPZkZpbGU6IHBhcnQubm9OZXdsaW5lQXRFbmRPZkZpbGUsXG4gICAgICAgIH0pXG4gICAgICAgIGJyZWFrXG4gICAgICBkZWZhdWx0OlxuICAgICAgICBhc3NlcnROZXZlcihwYXJ0LnR5cGUpXG4gICAgfVxuICB9XG5cbiAgLy8gc3dhcCBpbnNlcnRpb25zIGFuZCBkZWxldGlvbnMgb3ZlciBzbyBkZWxldGlvbnMgYWx3YXlzIGNvbWUgZmlyc3RcbiAgZm9yIChsZXQgaSA9IDA7IGkgPCBwYXJ0cy5sZW5ndGggLSAxOyBpKyspIHtcbiAgICBpZiAocGFydHNbaV0udHlwZSA9PT0gXCJpbnNlcnRpb25cIiAmJiBwYXJ0c1tpICsgMV0udHlwZSA9PT0gXCJkZWxldGlvblwiKSB7XG4gICAgICBjb25zdCB0bXAgPSBwYXJ0c1tpXVxuICAgICAgcGFydHNbaV0gPSBwYXJ0c1tpICsgMV1cbiAgICAgIHBhcnRzW2kgKyAxXSA9IHRtcFxuICAgICAgaSArPSAxXG4gICAgfVxuICB9XG5cbiAgY29uc3QgcmVzdWx0OiBIdW5rID0ge1xuICAgIGhlYWRlcixcbiAgICBwYXJ0cyxcbiAgICBzb3VyY2U6IGh1bmsuc291cmNlLFxuICB9XG5cbiAgdmVyaWZ5SHVua0ludGVncml0eShyZXN1bHQpXG5cbiAgcmV0dXJuIHJlc3VsdFxufVxuXG5mdW5jdGlvbiByZXZlcnNlUGF0Y2hQYXJ0KHBhcnQ6IFBhdGNoRmlsZVBhcnQpOiBQYXRjaEZpbGVQYXJ0IHtcbiAgc3dpdGNoIChwYXJ0LnR5cGUpIHtcbiAgICBjYXNlIFwiZmlsZSBjcmVhdGlvblwiOlxuICAgICAgcmV0dXJuIHtcbiAgICAgICAgdHlwZTogXCJmaWxlIGRlbGV0aW9uXCIsXG4gICAgICAgIHBhdGg6IHBhcnQucGF0aCxcbiAgICAgICAgaGFzaDogcGFydC5oYXNoLFxuICAgICAgICBodW5rOiBwYXJ0Lmh1bmsgJiYgcmV2ZXJzZUh1bmsocGFydC5odW5rKSxcbiAgICAgICAgbW9kZTogcGFydC5tb2RlLFxuICAgICAgfVxuICAgIGNhc2UgXCJmaWxlIGRlbGV0aW9uXCI6XG4gICAgICByZXR1cm4ge1xuICAgICAgICB0eXBlOiBcImZpbGUgY3JlYXRpb25cIixcbiAgICAgICAgcGF0aDogcGFydC5wYXRoLFxuICAgICAgICBodW5rOiBwYXJ0Lmh1bmsgJiYgcmV2ZXJzZUh1bmsocGFydC5odW5rKSxcbiAgICAgICAgbW9kZTogcGFydC5tb2RlLFxuICAgICAgICBoYXNoOiBwYXJ0Lmhhc2gsXG4gICAgICB9XG4gICAgY2FzZSBcInJlbmFtZVwiOlxuICAgICAgcmV0dXJuIHtcbiAgICAgICAgdHlwZTogXCJyZW5hbWVcIixcbiAgICAgICAgZnJvbVBhdGg6IHBhcnQudG9QYXRoLFxuICAgICAgICB0b1BhdGg6IHBhcnQuZnJvbVBhdGgsXG4gICAgICB9XG4gICAgY2FzZSBcInBhdGNoXCI6XG4gICAgICByZXR1cm4ge1xuICAgICAgICB0eXBlOiBcInBhdGNoXCIsXG4gICAgICAgIHBhdGg6IHBhcnQucGF0aCxcbiAgICAgICAgaHVua3M6IHBhcnQuaHVua3MubWFwKHJldmVyc2VIdW5rKSxcbiAgICAgICAgYmVmb3JlSGFzaDogcGFydC5hZnRlckhhc2gsXG4gICAgICAgIGFmdGVySGFzaDogcGFydC5iZWZvcmVIYXNoLFxuICAgICAgfVxuICAgIGNhc2UgXCJtb2RlIGNoYW5nZVwiOlxuICAgICAgcmV0dXJuIHtcbiAgICAgICAgdHlwZTogXCJtb2RlIGNoYW5nZVwiLFxuICAgICAgICBwYXRoOiBwYXJ0LnBhdGgsXG4gICAgICAgIG5ld01vZGU6IHBhcnQub2xkTW9kZSxcbiAgICAgICAgb2xkTW9kZTogcGFydC5uZXdNb2RlLFxuICAgICAgfVxuICB9XG59XG5cbmV4cG9ydCBjb25zdCByZXZlcnNlUGF0Y2ggPSAocGF0Y2g6IFBhcnNlZFBhdGNoRmlsZSk6IFBhcnNlZFBhdGNoRmlsZSA9PiB7XG4gIHJldHVybiBwYXRjaC5tYXAocmV2ZXJzZVBhdGNoUGFydCkucmV2ZXJzZSgpXG59XG4iXX0=