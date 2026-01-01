"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const resolveRelativeFileDependencies_1 = require("./resolveRelativeFileDependencies");
describe("resolveRelativeFileDependencies", () => {
    it("works for package.json", () => {
        const appRootPath = "/foo/bar";
        const resolutions = {
            absolute: "file:/not-foo/bar",
            relative: "file:../baz",
            remote: "git+https://blah.com/blah.git",
            version: "^434.34.34",
        };
        const expected = {
            absolute: "file:/not-foo/bar",
            relative: "file:/foo/baz",
            remote: "git+https://blah.com/blah.git",
            version: "^434.34.34",
        };
        expect(resolveRelativeFileDependencies_1.resolveRelativeFileDependencies(appRootPath, JSON.parse(JSON.stringify(resolutions)))).toEqual(expected);
    });
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicmVzb2x2ZVJlbGF0aXZlRmlsZURlcGVuZGVuY2llcy50ZXN0LmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiLi4vc3JjL3Jlc29sdmVSZWxhdGl2ZUZpbGVEZXBlbmRlbmNpZXMudGVzdC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUFBLHVGQUFtRjtBQUVuRixRQUFRLENBQUMsaUNBQWlDLEVBQUUsR0FBRyxFQUFFO0lBQy9DLEVBQUUsQ0FBQyx3QkFBd0IsRUFBRSxHQUFHLEVBQUU7UUFDaEMsTUFBTSxXQUFXLEdBQUcsVUFBVSxDQUFBO1FBRTlCLE1BQU0sV0FBVyxHQUFHO1lBQ2xCLFFBQVEsRUFBRSxtQkFBbUI7WUFDN0IsUUFBUSxFQUFFLGFBQWE7WUFDdkIsTUFBTSxFQUFFLCtCQUErQjtZQUN2QyxPQUFPLEVBQUUsWUFBWTtTQUN0QixDQUFBO1FBRUQsTUFBTSxRQUFRLEdBQUc7WUFDZixRQUFRLEVBQUUsbUJBQW1CO1lBQzdCLFFBQVEsRUFBRSxlQUFlO1lBQ3pCLE1BQU0sRUFBRSwrQkFBK0I7WUFDdkMsT0FBTyxFQUFFLFlBQVk7U0FDdEIsQ0FBQTtRQUVELE1BQU0sQ0FDSixpRUFBK0IsQ0FDN0IsV0FBVyxFQUNYLElBQUksQ0FBQyxLQUFLLENBQUMsSUFBSSxDQUFDLFNBQVMsQ0FBQyxXQUFXLENBQUMsQ0FBQyxDQUN4QyxDQUNGLENBQUMsT0FBTyxDQUFDLFFBQVEsQ0FBQyxDQUFBO0lBQ3JCLENBQUMsQ0FBQyxDQUFBO0FBQ0osQ0FBQyxDQUFDLENBQUEiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyByZXNvbHZlUmVsYXRpdmVGaWxlRGVwZW5kZW5jaWVzIH0gZnJvbSBcIi4vcmVzb2x2ZVJlbGF0aXZlRmlsZURlcGVuZGVuY2llc1wiXG5cbmRlc2NyaWJlKFwicmVzb2x2ZVJlbGF0aXZlRmlsZURlcGVuZGVuY2llc1wiLCAoKSA9PiB7XG4gIGl0KFwid29ya3MgZm9yIHBhY2thZ2UuanNvblwiLCAoKSA9PiB7XG4gICAgY29uc3QgYXBwUm9vdFBhdGggPSBcIi9mb28vYmFyXCJcblxuICAgIGNvbnN0IHJlc29sdXRpb25zID0ge1xuICAgICAgYWJzb2x1dGU6IFwiZmlsZTovbm90LWZvby9iYXJcIixcbiAgICAgIHJlbGF0aXZlOiBcImZpbGU6Li4vYmF6XCIsXG4gICAgICByZW1vdGU6IFwiZ2l0K2h0dHBzOi8vYmxhaC5jb20vYmxhaC5naXRcIixcbiAgICAgIHZlcnNpb246IFwiXjQzNC4zNC4zNFwiLFxuICAgIH1cblxuICAgIGNvbnN0IGV4cGVjdGVkID0ge1xuICAgICAgYWJzb2x1dGU6IFwiZmlsZTovbm90LWZvby9iYXJcIixcbiAgICAgIHJlbGF0aXZlOiBcImZpbGU6L2Zvby9iYXpcIixcbiAgICAgIHJlbW90ZTogXCJnaXQraHR0cHM6Ly9ibGFoLmNvbS9ibGFoLmdpdFwiLFxuICAgICAgdmVyc2lvbjogXCJeNDM0LjM0LjM0XCIsXG4gICAgfVxuXG4gICAgZXhwZWN0KFxuICAgICAgcmVzb2x2ZVJlbGF0aXZlRmlsZURlcGVuZGVuY2llcyhcbiAgICAgICAgYXBwUm9vdFBhdGgsXG4gICAgICAgIEpTT04ucGFyc2UoSlNPTi5zdHJpbmdpZnkocmVzb2x1dGlvbnMpKSxcbiAgICAgICksXG4gICAgKS50b0VxdWFsKGV4cGVjdGVkKVxuICB9KVxufSlcbiJdfQ==