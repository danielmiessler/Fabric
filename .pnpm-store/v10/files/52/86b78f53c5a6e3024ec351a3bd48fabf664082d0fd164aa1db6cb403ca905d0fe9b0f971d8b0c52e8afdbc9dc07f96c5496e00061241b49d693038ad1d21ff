"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const PackageDetails_1 = require("./PackageDetails");
describe("getPackageDetailsFromPatchFilename", () => {
    it("parses new-style patch filenames", () => {
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("banana++apple+0.4.2.patch"))
            .toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "banana => apple",
  "isDevOnly": false,
  "isNested": true,
  "name": "apple",
  "packageNames": Array [
    "banana",
    "apple",
  ],
  "patchFilename": "banana++apple+0.4.2.patch",
  "path": "node_modules/banana/node_modules/apple",
  "pathSpecifier": "banana/apple",
  "sequenceName": undefined,
  "sequenceNumber": undefined,
  "version": "0.4.2",
}
`);
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("@types+banana++@types+apple++@mollusc+man+0.4.2-banana-tree.patch")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@types/banana => @types/apple => @mollusc/man",
  "isDevOnly": false,
  "isNested": true,
  "name": "@mollusc/man",
  "packageNames": Array [
    "@types/banana",
    "@types/apple",
    "@mollusc/man",
  ],
  "patchFilename": "@types+banana++@types+apple++@mollusc+man+0.4.2-banana-tree.patch",
  "path": "node_modules/@types/banana/node_modules/@types/apple/node_modules/@mollusc/man",
  "pathSpecifier": "@types/banana/@types/apple/@mollusc/man",
  "sequenceName": undefined,
  "sequenceNumber": undefined,
  "version": "0.4.2-banana-tree",
}
`);
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("@types+banana.patch++hello+0.4.2-banana-tree.patch")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@types/banana.patch => hello",
  "isDevOnly": false,
  "isNested": true,
  "name": "hello",
  "packageNames": Array [
    "@types/banana.patch",
    "hello",
  ],
  "patchFilename": "@types+banana.patch++hello+0.4.2-banana-tree.patch",
  "path": "node_modules/@types/banana.patch/node_modules/hello",
  "pathSpecifier": "@types/banana.patch/hello",
  "sequenceName": undefined,
  "sequenceNumber": undefined,
  "version": "0.4.2-banana-tree",
}
`);
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("@types+banana.patch++hello+0.4.2-banana-tree.dev.patch")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@types/banana.patch => hello",
  "isDevOnly": true,
  "isNested": true,
  "name": "hello",
  "packageNames": Array [
    "@types/banana.patch",
    "hello",
  ],
  "patchFilename": "@types+banana.patch++hello+0.4.2-banana-tree.dev.patch",
  "path": "node_modules/@types/banana.patch/node_modules/hello",
  "pathSpecifier": "@types/banana.patch/hello",
  "sequenceName": undefined,
  "sequenceNumber": undefined,
  "version": "0.4.2-banana-tree",
}
`);
    });
    it("works for ordered patches", () => {
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("left-pad+1.3.0+02+world"))
            .toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "left-pad",
  "isDevOnly": false,
  "isNested": false,
  "name": "left-pad",
  "packageNames": Array [
    "left-pad",
  ],
  "patchFilename": "left-pad+1.3.0+02+world",
  "path": "node_modules/left-pad",
  "pathSpecifier": "left-pad",
  "sequenceName": "world",
  "sequenceNumber": 2,
  "version": "1.3.0",
}
`);
        expect(PackageDetails_1.getPackageDetailsFromPatchFilename("@microsoft/api-extractor+2.0.0+01+FixThing")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@microsoft/api-extractor",
  "isDevOnly": false,
  "isNested": false,
  "name": "@microsoft/api-extractor",
  "packageNames": Array [
    "@microsoft/api-extractor",
  ],
  "patchFilename": "@microsoft/api-extractor+2.0.0+01+FixThing",
  "path": "node_modules/@microsoft/api-extractor",
  "pathSpecifier": "@microsoft/api-extractor",
  "sequenceName": "FixThing",
  "sequenceNumber": 1,
  "version": "2.0.0",
}
`);
    });
});
describe("getPatchDetailsFromCliString", () => {
    it("handles a minimal package name", () => {
        expect(PackageDetails_1.getPatchDetailsFromCliString("patch-package")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "patch-package",
  "isNested": false,
  "name": "patch-package",
  "packageNames": Array [
    "patch-package",
  ],
  "path": "node_modules/patch-package",
  "pathSpecifier": "patch-package",
}
`);
    });
    it("handles a scoped package name", () => {
        expect(PackageDetails_1.getPatchDetailsFromCliString("@david/patch-package")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@david/patch-package",
  "isNested": false,
  "name": "@david/patch-package",
  "packageNames": Array [
    "@david/patch-package",
  ],
  "path": "node_modules/@david/patch-package",
  "pathSpecifier": "@david/patch-package",
}
`);
    });
    it("handles a nested package name", () => {
        expect(PackageDetails_1.getPatchDetailsFromCliString("david/patch-package")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "david => patch-package",
  "isNested": true,
  "name": "patch-package",
  "packageNames": Array [
    "david",
    "patch-package",
  ],
  "path": "node_modules/david/node_modules/patch-package",
  "pathSpecifier": "david/patch-package",
}
`);
    });
    it("handles a nested package name with scopes", () => {
        expect(PackageDetails_1.getPatchDetailsFromCliString("@david/patch-package/banana")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@david/patch-package => banana",
  "isNested": true,
  "name": "banana",
  "packageNames": Array [
    "@david/patch-package",
    "banana",
  ],
  "path": "node_modules/@david/patch-package/node_modules/banana",
  "pathSpecifier": "@david/patch-package/banana",
}
`);
        expect(PackageDetails_1.getPatchDetailsFromCliString("@david/patch-package/@david/banana")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "@david/patch-package => @david/banana",
  "isNested": true,
  "name": "@david/banana",
  "packageNames": Array [
    "@david/patch-package",
    "@david/banana",
  ],
  "path": "node_modules/@david/patch-package/node_modules/@david/banana",
  "pathSpecifier": "@david/patch-package/@david/banana",
}
`);
        expect(PackageDetails_1.getPatchDetailsFromCliString("david/patch-package/@david/banana")).toMatchInlineSnapshot(`
Object {
  "humanReadablePathSpecifier": "david => patch-package => @david/banana",
  "isNested": true,
  "name": "@david/banana",
  "packageNames": Array [
    "david",
    "patch-package",
    "@david/banana",
  ],
  "path": "node_modules/david/node_modules/patch-package/node_modules/@david/banana",
  "pathSpecifier": "david/patch-package/@david/banana",
}
`);
    });
});
describe("parseNameAndVersion", () => {
    it("works for good-looking names", () => {
        expect(PackageDetails_1.parseNameAndVersion("lodash+2.3.4")).toMatchInlineSnapshot(`
Object {
  "packageName": "lodash",
  "version": "2.3.4",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("patch-package+2.0.0-alpha.3"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "patch-package",
  "version": "2.0.0-alpha.3",
}
`);
    });
    it("works for scoped package names", () => {
        expect(PackageDetails_1.parseNameAndVersion("@react-spring+rafz+2.0.0-alpha.3"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@react-spring/rafz",
  "version": "2.0.0-alpha.3",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("@microsoft+api-extractor+2.2.3"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@microsoft/api-extractor",
  "version": "2.2.3",
}
`);
    });
    it("works for ordered patches", () => {
        expect(PackageDetails_1.parseNameAndVersion("patch-package+2.0.0+01"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "patch-package",
  "sequenceNumber": 1,
  "version": "2.0.0",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("@react-spring+rafz+2.0.0-alpha.3+23"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@react-spring/rafz",
  "sequenceNumber": 23,
  "version": "2.0.0-alpha.3",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("@microsoft+api-extractor+2.0.0+001"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@microsoft/api-extractor",
  "sequenceNumber": 1,
  "version": "2.0.0",
}
`);
    });
    it("works for ordered patches with names", () => {
        expect(PackageDetails_1.parseNameAndVersion("patch-package+2.0.0+021+FixImportantThing"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "patch-package",
  "sequenceName": "FixImportantThing",
  "sequenceNumber": 21,
  "version": "2.0.0",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("@react-spring+rafz+2.0.0-alpha.3+000023+Foo"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@react-spring/rafz",
  "sequenceName": "Foo",
  "sequenceNumber": 23,
  "version": "2.0.0-alpha.3",
}
`);
        expect(PackageDetails_1.parseNameAndVersion("@microsoft+api-extractor+2.0.0+001+Bar"))
            .toMatchInlineSnapshot(`
Object {
  "packageName": "@microsoft/api-extractor",
  "sequenceName": "Bar",
  "sequenceNumber": 1,
  "version": "2.0.0",
}
`);
    });
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiUGFja2FnZURldGFpbHMudGVzdC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbIi4uL3NyYy9QYWNrYWdlRGV0YWlscy50ZXN0LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7O0FBQUEscURBSXlCO0FBRXpCLFFBQVEsQ0FBQyxvQ0FBb0MsRUFBRSxHQUFHLEVBQUU7SUFDbEQsRUFBRSxDQUFDLGtDQUFrQyxFQUFFLEdBQUcsRUFBRTtRQUMxQyxNQUFNLENBQUMsbURBQWtDLENBQUMsMkJBQTJCLENBQUMsQ0FBQzthQUNwRSxxQkFBcUIsQ0FBQzs7Ozs7Ozs7Ozs7Ozs7Ozs7Q0FpQjVCLENBQUMsQ0FBQTtRQUVFLE1BQU0sQ0FDSixtREFBa0MsQ0FDaEMsbUVBQW1FLENBQ3BFLENBQ0YsQ0FBQyxxQkFBcUIsQ0FBQzs7Ozs7Ozs7Ozs7Ozs7Ozs7O0NBa0IzQixDQUFDLENBQUE7UUFFRSxNQUFNLENBQ0osbURBQWtDLENBQ2hDLG9EQUFvRCxDQUNyRCxDQUNGLENBQUMscUJBQXFCLENBQUM7Ozs7Ozs7Ozs7Ozs7Ozs7O0NBaUIzQixDQUFDLENBQUE7UUFFRSxNQUFNLENBQ0osbURBQWtDLENBQ2hDLHdEQUF3RCxDQUN6RCxDQUNGLENBQUMscUJBQXFCLENBQUM7Ozs7Ozs7Ozs7Ozs7Ozs7O0NBaUIzQixDQUFDLENBQUE7SUFDQSxDQUFDLENBQUMsQ0FBQTtJQUVGLEVBQUUsQ0FBQywyQkFBMkIsRUFBRSxHQUFHLEVBQUU7UUFDbkMsTUFBTSxDQUFDLG1EQUFrQyxDQUFDLHlCQUF5QixDQUFDLENBQUM7YUFDbEUscUJBQXFCLENBQUM7Ozs7Ozs7Ozs7Ozs7Ozs7Q0FnQjVCLENBQUMsQ0FBQTtRQUVFLE1BQU0sQ0FDSixtREFBa0MsQ0FDaEMsNENBQTRDLENBQzdDLENBQ0YsQ0FBQyxxQkFBcUIsQ0FBQzs7Ozs7Ozs7Ozs7Ozs7OztDQWdCM0IsQ0FBQyxDQUFBO0lBQ0EsQ0FBQyxDQUFDLENBQUE7QUFDSixDQUFDLENBQUMsQ0FBQTtBQUVGLFFBQVEs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QTtRQUNFLE1BQU0sQ0FBQyxvQ0FBbUIsQ0FBQyx3Q0FBd0MsQ0FBQyxDQUFDO2FBQ2xFLHFCQUFxQixDQUFDOzs7Ozs7O0NBTzVCLENBQUMsQ0FBQTtJQUNBLENBQUMsQ0FBQyxDQUFBO0FBQ0osQ0FBQyxDQUFDLENBQUEiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQge1xuICBnZXRQYWNrYWdlRGV0YWlsc0Zyb21QYXRjaEZpbGVuYW1lLFxuICBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nLFxuICBwYXJzZU5hbWVBbmRWZXJzaW9uLFxufSBmcm9tIFwiLi9QYWNrYWdlRGV0YWlsc1wiXG5cbmRlc2NyaWJlKFwiZ2V0UGFja2FnZURldGFpbHNGcm9tUGF0Y2hGaWxlbmFtZVwiLCAoKSA9PiB7XG4gIGl0KFwicGFyc2VzIG5ldy1zdHlsZSBwYXRjaCBmaWxlbmFtZXNcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChnZXRQYWNrYWdlRGV0YWlsc0Zyb21QYXRjaEZpbGVuYW1lKFwiYmFuYW5hKythcHBsZSswLjQuMi5wYXRjaFwiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcImJhbmFuYSA9PiBhcHBsZVwiLFxuICBcImlzRGV2T25seVwiOiBmYWxzZSxcbiAgXCJpc05lc3RlZFwiOiB0cnVlLFxuICBcIm5hbWVcIjogXCJhcHBsZVwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJiYW5hbmFcIixcbiAgICBcImFwcGxlXCIsXG4gIF0sXG4gIFwicGF0Y2hGaWxlbmFtZVwiOiBcImJhbmFuYSsrYXBwbGUrMC40LjIucGF0Y2hcIixcbiAgXCJwYXRoXCI6IFwibm9kZV9tb2R1bGVzL2JhbmFuYS9ub2RlX21vZHVsZXMvYXBwbGVcIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwiYmFuYW5hL2FwcGxlXCIsXG4gIFwic2VxdWVuY2VOYW1lXCI6IHVuZGVmaW5lZCxcbiAgXCJzZXF1ZW5jZU51bWJlclwiOiB1bmRlZmluZWQsXG4gIFwidmVyc2lvblwiOiBcIjAuNC4yXCIsXG59XG5gKVxuXG4gICAgZXhwZWN0KFxuICAgICAgZ2V0UGFja2FnZURldGFpbHNGcm9tUGF0Y2hGaWxlbmFtZShcbiAgICAgICAgXCJAdHlwZXMrYmFuYW5hKytAdHlwZXMrYXBwbGUrK0Btb2xsdXNjK21hbiswLjQuMi1iYW5hbmEtdHJlZS5wYXRjaFwiLFxuICAgICAgKSxcbiAgICApLnRvTWF0Y2hJbmxpbmVTbmFwc2hvdChgXG5PYmplY3Qge1xuICBcImh1bWFuUmVhZGFibGVQYXRoU3BlY2lmaWVyXCI6IFwiQHR5cGVzL2JhbmFuYSA9PiBAdHlwZXMvYXBwbGUgPT4gQG1vbGx1c2MvbWFuXCIsXG4gIFwiaXNEZXZPbmx5XCI6IGZhbHNlLFxuICBcImlzTmVzdGVkXCI6IHRydWUsXG4gIFwibmFtZVwiOiBcIkBtb2xsdXNjL21hblwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJAdHlwZXMvYmFuYW5hXCIsXG4gICAgXCJAdHlwZXMvYXBwbGVcIixcbiAgICBcIkBtb2xsdXNjL21hblwiLFxuICBdLFxuICBcInBhdGNoRmlsZW5hbWVcIjogXCJAdHlwZXMrYmFuYW5hKytAdHlwZXMrYXBwbGUrK0Btb2xsdXNjK21hbiswLjQuMi1iYW5hbmEtdHJlZS5wYXRjaFwiLFxuICBcInBhdGhcIjogXCJub2RlX21vZHVsZXMvQHR5cGVzL2JhbmFuYS9ub2RlX21vZHVsZXMvQHR5cGVzL2FwcGxlL25vZGVfbW9kdWxlcy9AbW9sbHVzYy9tYW5cIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwiQHR5cGVzL2JhbmFuYS9AdHlwZXMvYXBwbGUvQG1vbGx1c2MvbWFuXCIsXG4gIFwic2VxdWVuY2VOYW1lXCI6IHVuZGVmaW5lZCxcbiAgXCJzZXF1ZW5jZU51bWJlclwiOiB1bmRlZmluZWQsXG4gIFwidmVyc2lvblwiOiBcIjAuNC4yLWJhbmFuYS10cmVlXCIsXG59XG5gKVxuXG4gICAgZXhwZWN0KFxuICAgICAgZ2V0UGFja2FnZURldGFpbHNGcm9tUGF0Y2hGaWxlbmFtZShcbiAgICAgICAgXCJAdHlwZXMrYmFuYW5hLnBhdGNoKytoZWxsbyswLjQuMi1iYW5hbmEtdHJlZS5wYXRjaFwiLFxuICAgICAgKSxcbiAgICApLnRvTWF0Y2hJbmxpbmVTbmFwc2hvdChgXG5PYmplY3Qge1xuICBcImh1bWFuUmVhZGFibGVQYXRoU3BlY2lmaWVyXCI6IFwiQHR5cGVzL2JhbmFuYS5wYXRjaCA9PiBoZWxsb1wiLFxuICBcImlzRGV2T25seVwiOiBmYWxzZSxcbiAgXCJpc05lc3RlZFwiOiB0cnVlLFxuICBcIm5hbWVcIjogXCJoZWxsb1wiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJAdHlwZXMvYmFuYW5hLnBhdGNoXCIsXG4gICAgXCJoZWxsb1wiLFxuICBdLFxuICBcInBhdGNoRmlsZW5hbWVcIjogXCJAdHlwZXMrYmFuYW5hLnBhdGNoKytoZWxsbyswLjQuMi1iYW5hbmEtdHJlZS5wYXRjaFwiLFxuICBcInBhdGhcIjogXCJub2RlX21vZHVsZXMvQHR5cGVzL2JhbmFuYS5wYXRjaC9ub2RlX21vZHVsZXMvaGVsbG9cIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwiQHR5cGVzL2JhbmFuYS5wYXRjaC9oZWxsb1wiLFxuICBcInNlcXVlbmNlTmFtZVwiOiB1bmRlZmluZWQsXG4gIFwic2VxdWVuY2VOdW1iZXJcIjogdW5kZWZpbmVkLFxuICBcInZlcnNpb25cIjogXCIwLjQuMi1iYW5hbmEtdHJlZVwiLFxufVxuYClcblxuICAgIGV4cGVjdChcbiAgICAgIGdldFBhY2thZ2VEZXRhaWxzRnJvbVBhdGNoRmlsZW5hbWUoXG4gICAgICAgIFwiQHR5cGVzK2JhbmFuYS5wYXRjaCsraGVsbG8rMC40LjItYmFuYW5hLXRyZWUuZGV2LnBhdGNoXCIsXG4gICAgICApLFxuICAgICkudG9NYXRjaElubGluZVNuYXBzaG90KGBcbk9iamVjdCB7XG4gIFwiaHVtYW5SZWFkYWJsZVBhdGhTcGVjaWZpZXJcIjogXCJAdHlwZXMvYmFuYW5hLnBhdGNoID0+IGhlbGxvXCIsXG4gIFwiaXNEZXZPbmx5XCI6IHRydWUsXG4gIFwiaXNOZXN0ZWRcIjogdHJ1ZSxcbiAgXCJuYW1lXCI6IFwiaGVsbG9cIixcbiAgXCJwYWNrYWdlTmFtZXNcIjogQXJyYXkgW1xuICAgIFwiQHR5cGVzL2JhbmFuYS5wYXRjaFwiLFxuICAgIFwiaGVsbG9cIixcbiAgXSxcbiAgXCJwYXRjaEZpbGVuYW1lXCI6IFwiQHR5cGVzK2JhbmFuYS5wYXRjaCsraGVsbG8rMC40LjItYmFuYW5hLXRyZWUuZGV2LnBhdGNoXCIsXG4gIFwicGF0aFwiOiBcIm5vZGVfbW9kdWxlcy9AdHlwZXMvYmFuYW5hLnBhdGNoL25vZGVfbW9kdWxlcy9oZWxsb1wiLFxuICBcInBhdGhTcGVjaWZpZXJcIjogXCJAdHlwZXMvYmFuYW5hLnBhdGNoL2hlbGxvXCIsXG4gIFwic2VxdWVuY2VOYW1lXCI6IHVuZGVmaW5lZCxcbiAgXCJzZXF1ZW5jZU51bWJlclwiOiB1bmRlZmluZWQsXG4gIFwidmVyc2lvblwiOiBcIjAuNC4yLWJhbmFuYS10cmVlXCIsXG59XG5gKVxuICB9KVxuXG4gIGl0KFwid29ya3MgZm9yIG9yZGVyZWQgcGF0Y2hlc1wiLCAoKSA9PiB7XG4gICAgZXhwZWN0KGdldFBhY2thZ2VEZXRhaWxzRnJvbVBhdGNoRmlsZW5hbWUoXCJsZWZ0LXBhZCsxLjMuMCswMit3b3JsZFwiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcImxlZnQtcGFkXCIsXG4gIFwiaXNEZXZPbmx5XCI6IGZhbHNlLFxuICBcImlzTmVzdGVkXCI6IGZhbHNlLFxuICBcIm5hbWVcIjogXCJsZWZ0LXBhZFwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJsZWZ0LXBhZFwiLFxuICBdLFxuICBcInBhdGNoRmlsZW5hbWVcIjogXCJsZWZ0LXBhZCsxLjMuMCswMit3b3JsZFwiLFxuICBcInBhdGhcIjogXCJub2RlX21vZHVsZXMvbGVmdC1wYWRcIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwibGVmdC1wYWRcIixcbiAgXCJzZXF1ZW5jZU5hbWVcIjogXCJ3b3JsZFwiLFxuICBcInNlcXVlbmNlTnVtYmVyXCI6IDIsXG4gIFwidmVyc2lvblwiOiBcIjEuMy4wXCIsXG59XG5gKVxuXG4gICAgZXhwZWN0KFxuICAgICAgZ2V0UGFja2FnZURldGFpbHNGcm9tUGF0Y2hGaWxlbmFtZShcbiAgICAgICAgXCJAbWljcm9zb2Z0L2FwaS1leHRyYWN0b3IrMi4wLjArMDErRml4VGhpbmdcIixcbiAgICAgICksXG4gICAgKS50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcIkBtaWNyb3NvZnQvYXBpLWV4dHJhY3RvclwiLFxuICBcImlzRGV2T25seVwiOiBmYWxzZSxcbiAgXCJpc05lc3RlZFwiOiBmYWxzZSxcbiAgXCJuYW1lXCI6IFwiQG1pY3Jvc29mdC9hcGktZXh0cmFjdG9yXCIsXG4gIFwicGFja2FnZU5hbWVzXCI6IEFycmF5IFtcbiAgICBcIkBtaWNyb3NvZnQvYXBpLWV4dHJhY3RvclwiLFxuICBdLFxuICBcInBhdGNoRmlsZW5hbWVcIjogXCJAbWljcm9zb2Z0L2FwaS1leHRyYWN0b3IrMi4wLjArMDErRml4VGhpbmdcIixcbiAgXCJwYXRoXCI6IFwibm9kZV9tb2R1bGVzL0BtaWNyb3NvZnQvYXBpLWV4dHJhY3RvclwiLFxuICBcInBhdGhTcGVjaWZpZXJcIjogXCJAbWljcm9zb2Z0L2FwaS1leHRyYWN0b3JcIixcbiAgXCJzZXF1ZW5jZU5hbWVcIjogXCJGaXhUaGluZ1wiLFxuICBcInNlcXVlbmNlTnVtYmVyXCI6IDEsXG4gIFwidmVyc2lvblwiOiBcIjIuMC4wXCIsXG59XG5gKVxuICB9KVxufSlcblxuZGVzY3JpYmUoXCJnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nXCIsICgpID0+IHtcbiAgaXQoXCJoYW5kbGVzIGEgbWluaW1hbCBwYWNrYWdlIG5hbWVcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nKFwicGF0Y2gtcGFja2FnZVwiKSkudG9NYXRjaElubGluZVNuYXBzaG90KFxuICAgICAgYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcInBhdGNoLXBhY2thZ2VcIixcbiAgXCJpc05lc3RlZFwiOiBmYWxzZSxcbiAgXCJuYW1lXCI6IFwicGF0Y2gtcGFja2FnZVwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJwYXRjaC1wYWNrYWdlXCIsXG4gIF0sXG4gIFwicGF0aFwiOiBcIm5vZGVfbW9kdWxlcy9wYXRjaC1wYWNrYWdlXCIsXG4gIFwicGF0aFNwZWNpZmllclwiOiBcInBhdGNoLXBhY2thZ2VcIixcbn1cbmAsXG4gICAgKVxuICB9KVxuXG4gIGl0KFwiaGFuZGxlcyBhIHNjb3BlZCBwYWNrYWdlIG5hbWVcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChcbiAgICAgIGdldFBhdGNoRGV0YWlsc0Zyb21DbGlTdHJpbmcoXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiKSxcbiAgICApLnRvTWF0Y2hJbmxpbmVTbmFwc2hvdChcbiAgICAgIGBcbk9iamVjdCB7XG4gIFwiaHVtYW5SZWFkYWJsZVBhdGhTcGVjaWZpZXJcIjogXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiLFxuICBcImlzTmVzdGVkXCI6IGZhbHNlLFxuICBcIm5hbWVcIjogXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiLFxuICBdLFxuICBcInBhdGhcIjogXCJub2RlX21vZHVsZXMvQGRhdmlkL3BhdGNoLXBhY2thZ2VcIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwiQGRhdmlkL3BhdGNoLXBhY2thZ2VcIixcbn1cbmAsXG4gICAgKVxuICB9KVxuXG4gIGl0KFwiaGFuZGxlcyBhIG5lc3RlZCBwYWNrYWdlIG5hbWVcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChcbiAgICAgIGdldFBhdGNoRGV0YWlsc0Zyb21DbGlTdHJpbmcoXCJkYXZpZC9wYXRjaC1wYWNrYWdlXCIpLFxuICAgICkudG9NYXRjaElubGluZVNuYXBzaG90KFxuICAgICAgYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcImRhdmlkID0+IHBhdGNoLXBhY2thZ2VcIixcbiAgXCJpc05lc3RlZFwiOiB0cnVlLFxuICBcIm5hbWVcIjogXCJwYXRjaC1wYWNrYWdlXCIsXG4gIFwicGFja2FnZU5hbWVzXCI6IEFycmF5IFtcbiAgICBcImRhdmlkXCIsXG4gICAgXCJwYXRjaC1wYWNrYWdlXCIsXG4gIF0sXG4gIFwicGF0aFwiOiBcIm5vZGVfbW9kdWxlcy9kYXZpZC9ub2RlX21vZHVsZXMvcGF0Y2gtcGFja2FnZVwiLFxuICBcInBhdGhTcGVjaWZpZXJcIjogXCJkYXZpZC9wYXRjaC1wYWNrYWdlXCIsXG59XG5gLFxuICAgIClcbiAgfSlcblxuICBpdChcImhhbmRsZXMgYSBuZXN0ZWQgcGFja2FnZSBuYW1lIHdpdGggc2NvcGVzXCIsICgpID0+IHtcbiAgICBleHBlY3QoXG4gICAgICBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nKFwiQGRhdmlkL3BhdGNoLXBhY2thZ2UvYmFuYW5hXCIpLFxuICAgICkudG9NYXRjaElubGluZVNuYXBzaG90KFxuICAgICAgYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcIkBkYXZpZC9wYXRjaC1wYWNrYWdlID0+IGJhbmFuYVwiLFxuICBcImlzTmVzdGVkXCI6IHRydWUsXG4gIFwibmFtZVwiOiBcImJhbmFuYVwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiLFxuICAgIFwiYmFuYW5hXCIsXG4gIF0sXG4gIFwicGF0aFwiOiBcIm5vZGVfbW9kdWxlcy9AZGF2aWQvcGF0Y2gtcGFja2FnZS9ub2RlX21vZHVsZXMvYmFuYW5hXCIsXG4gIFwicGF0aFNwZWNpZmllclwiOiBcIkBkYXZpZC9wYXRjaC1wYWNrYWdlL2JhbmFuYVwiLFxufVxuYCxcbiAgICApXG5cbiAgICBleHBlY3QoXG4gICAgICBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nKFwiQGRhdmlkL3BhdGNoLXBhY2thZ2UvQGRhdmlkL2JhbmFuYVwiKSxcbiAgICApLnRvTWF0Y2hJbmxpbmVTbmFwc2hvdChcbiAgICAgIGBcbk9iamVjdCB7XG4gIFwiaHVtYW5SZWFkYWJsZVBhdGhTcGVjaWZpZXJcIjogXCJAZGF2aWQvcGF0Y2gtcGFja2FnZSA9PiBAZGF2aWQvYmFuYW5hXCIsXG4gIFwiaXNOZXN0ZWRcIjogdHJ1ZSxcbiAgXCJuYW1lXCI6IFwiQGRhdmlkL2JhbmFuYVwiLFxuICBcInBhY2thZ2VOYW1lc1wiOiBBcnJheSBbXG4gICAgXCJAZGF2aWQvcGF0Y2gtcGFja2FnZVwiLFxuICAgIFwiQGRhdmlkL2JhbmFuYVwiLFxuICBdLFxuICBcInBhdGhcIjogXCJub2RlX21vZHVsZXMvQGRhdmlkL3BhdGNoLXBhY2thZ2Uvbm9kZV9tb2R1bGVzL0BkYXZpZC9iYW5hbmFcIixcbiAgXCJwYXRoU3BlY2lmaWVyXCI6IFwiQGRhdmlkL3BhdGNoLXBhY2thZ2UvQGRhdmlkL2JhbmFuYVwiLFxufVxuYCxcbiAgICApXG5cbiAgICBleHBlY3QoXG4gICAgICBnZXRQYXRjaERldGFpbHNGcm9tQ2xpU3RyaW5nKFwiZGF2aWQvcGF0Y2gtcGFja2FnZS9AZGF2aWQvYmFuYW5hXCIpLFxuICAgICkudG9NYXRjaElubGluZVNuYXBzaG90KFxuICAgICAgYFxuT2JqZWN0IHtcbiAgXCJodW1hblJlYWRhYmxlUGF0aFNwZWNpZmllclwiOiBcImRhdmlkID0+IHBhdGNoLXBhY2thZ2UgPT4gQGRhdmlkL2JhbmFuYVwiLFxuICBcImlzTmVzdGVkXCI6IHRydWUsXG4gIFwibmFtZVwiOiBcIkBkYXZpZC9iYW5hbmFcIixcbiAgXCJwYWNrYWdlTmFtZXNcIjogQXJyYXkgW1xuICAgIFwiZGF2aWRcIixcbiAgICBcInBhdGNoLXBhY2thZ2VcIixcbiAgICBcIkBkYXZpZC9iYW5hbmFcIixcbiAgXSxcbiAgXCJwYXRoXCI6IFwibm9kZV9tb2R1bGVzL2RhdmlkL25vZGVfbW9kdWxlcy9wYXRjaC1wYWNrYWdlL25vZGVfbW9kdWxlcy9AZGF2aWQvYmFuYW5hXCIsXG4gIFwicGF0aFNwZWNpZmllclwiOiBcImRhdmlkL3BhdGNoLXBhY2thZ2UvQGRhdmlkL2JhbmFuYVwiLFxufVxuYCxcbiAgICApXG4gIH0pXG59KVxuXG5kZXNjcmliZShcInBhcnNlTmFtZUFuZFZlcnNpb25cIiwgKCkgPT4ge1xuICBpdChcIndvcmtzIGZvciBnb29kLWxvb2tpbmcgbmFtZXNcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChwYXJzZU5hbWVBbmRWZXJzaW9uKFwibG9kYXNoKzIuMy40XCIpKS50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcImxvZGFzaFwiLFxuICBcInZlcnNpb25cIjogXCIyLjMuNFwiLFxufVxuYClcbiAgICBleHBlY3QocGFyc2VOYW1lQW5kVmVyc2lvbihcInBhdGNoLXBhY2thZ2UrMi4wLjAtYWxwaGEuM1wiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcInBhdGNoLXBhY2thZ2VcIixcbiAgXCJ2ZXJzaW9uXCI6IFwiMi4wLjAtYWxwaGEuM1wiLFxufVxuYClcbiAgfSlcbiAgaXQoXCJ3b3JrcyBmb3Igc2NvcGVkIHBhY2thZ2UgbmFtZXNcIiwgKCkgPT4ge1xuICAgIGV4cGVjdChwYXJzZU5hbWVBbmRWZXJzaW9uKFwiQHJlYWN0LXNwcmluZytyYWZ6KzIuMC4wLWFscGhhLjNcIikpXG4gICAgICAudG9NYXRjaElubGluZVNuYXBzaG90KGBcbk9iamVjdCB7XG4gIFwicGFja2FnZU5hbWVcIjogXCJAcmVhY3Qtc3ByaW5nL3JhZnpcIixcbiAgXCJ2ZXJzaW9uXCI6IFwiMi4wLjAtYWxwaGEuM1wiLFxufVxuYClcbiAgICBleHBlY3QocGFyc2VOYW1lQW5kVmVyc2lvbihcIkBtaWNyb3NvZnQrYXBpLWV4dHJhY3RvcisyLjIuM1wiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcIkBtaWNyb3NvZnQvYXBpLWV4dHJhY3RvclwiLFxuICBcInZlcnNpb25cIjogXCIyLjIuM1wiLFxufVxuYClcbiAgfSlcbiAgaXQoXCJ3b3JrcyBmb3Igb3JkZXJlZCBwYXRjaGVzXCIsICgpID0+IHtcbiAgICBleHBlY3QocGFyc2VOYW1lQW5kVmVyc2lvbihcInBhdGNoLXBhY2thZ2UrMi4wLjArMDFcIikpXG4gICAgICAudG9NYXRjaElubGluZVNuYXBzaG90KGBcbk9iamVjdCB7XG4gIFwicGFja2FnZU5hbWVcIjogXCJwYXRjaC1wYWNrYWdlXCIsXG4gIFwic2VxdWVuY2VOdW1iZXJcIjogMSxcbiAgXCJ2ZXJzaW9uXCI6IFwiMi4wLjBcIixcbn1cbmApXG4gICAgZXhwZWN0KHBhcnNlTmFtZUFuZFZlcnNpb24oXCJAcmVhY3Qtc3ByaW5nK3JhZnorMi4wLjAtYWxwaGEuMysyM1wiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcIkByZWFjdC1zcHJpbmcvcmFmelwiLFxuICBcInNlcXVlbmNlTnVtYmVyXCI6IDIzLFxuICBcInZlcnNpb25cIjogXCIyLjAuMC1hbHBoYS4zXCIsXG59XG5gKVxuICAgIGV4cGVjdChwYXJzZU5hbWVBbmRWZXJzaW9uKFwiQG1pY3Jvc29mdCthcGktZXh0cmFjdG9yKzIuMC4wKzAwMVwiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcIkBtaWNyb3NvZnQvYXBpLWV4dHJhY3RvclwiLFxuICBcInNlcXVlbmNlTnVtYmVyXCI6IDEsXG4gIFwidmVyc2lvblwiOiBcIjIuMC4wXCIsXG59XG5gKVxuICB9KVxuXG4gIGl0KFwid29ya3MgZm9yIG9yZGVyZWQgcGF0Y2hlcyB3aXRoIG5hbWVzXCIsICgpID0+IHtcbiAgICBleHBlY3QocGFyc2VOYW1lQW5kVmVyc2lvbihcInBhdGNoLXBhY2thZ2UrMi4wLjArMDIxK0ZpeEltcG9ydGFudFRoaW5nXCIpKVxuICAgICAgLnRvTWF0Y2hJbmxpbmVTbmFwc2hvdChgXG5PYmplY3Qge1xuICBcInBhY2thZ2VOYW1lXCI6IFwicGF0Y2gtcGFja2FnZVwiLFxuICBcInNlcXVlbmNlTmFtZVwiOiBcIkZpeEltcG9ydGFudFRoaW5nXCIsXG4gIFwic2VxdWVuY2VOdW1iZXJcIjogMjEsXG4gIFwidmVyc2lvblwiOiBcIjIuMC4wXCIsXG59XG5gKVxuICAgIGV4cGVjdChwYXJzZU5hbWVBbmRWZXJzaW9uKFwiQHJlYWN0LXNwcmluZytyYWZ6KzIuMC4wLWFscGhhLjMrMDAwMDIzK0Zvb1wiKSlcbiAgICAgIC50b01hdGNoSW5saW5lU25hcHNob3QoYFxuT2JqZWN0IHtcbiAgXCJwYWNrYWdlTmFtZVwiOiBcIkByZWFjdC1zcHJpbmcvcmFmelwiLFxuICBcInNlcXVlbmNlTmFtZVwiOiBcIkZvb1wiLFxuICBcInNlcXVlbmNlTnVtYmVyXCI6IDIzLFxuICBcInZlcnNpb25cIjogXCIyLjAuMC1hbHBoYS4zXCIsXG59XG5gKVxuICAgIGV4cGVjdChwYXJzZU5hbWVBbmRWZXJzaW9uKFwiQG1pY3Jvc29mdCthcGktZXh0cmFjdG9yKzIuMC4wKzAwMStCYXJcIikpXG4gICAgICAudG9NYXRjaElubGluZVNuYXBzaG90KGBcbk9iamVjdCB7XG4gIFwicGFja2FnZU5hbWVcIjogXCJAbWljcm9zb2Z0L2FwaS1leHRyYWN0b3JcIixcbiAgXCJzZXF1ZW5jZU5hbWVcIjogXCJCYXJcIixcbiAgXCJzZXF1ZW5jZU51bWJlclwiOiAxLFxuICBcInZlcnNpb25cIjogXCIyLjAuMFwiLFxufVxuYClcbiAgfSlcbn0pXG4iXX0=