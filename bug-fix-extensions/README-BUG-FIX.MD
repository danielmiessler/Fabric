
# Bug Fix of {{input}} in extensions

## NOTES

- diff two repos with bc

- run tests and coverage before fix
go test -v -coverprofile=coverage.out ./... > bug-fix-extensions/test_results_with_coverage_before_fix.txt 2>&1;rm coverage.out

- `go mod tidy` will remove this from go.sum
 -github.com/anthropics/anthropic-sdk-go v1.12.0 h1:xPqlGnq7rWrTiHazIvCiumA0u7mGQnwDQtvA1M82h9U=
 -github.com/anthropics/anthropic-sdk-go v1.12.0/go.mod h1:WTz31rIUHUHqai2UslPpw5CwXrQP3geYBioRV4WOLvE=

- ☑️ Compare Fabric Fix and Fabric in BC - see big picture
- Bring Test over that shows it fails as is
- Learn how to debug word generator
- Compare Java to Go in table
- create word counter
- show it fails
- create a test that shows fail in general
- make the fix
- show the test passes
- Make full code execution calling claude and chat via api
- create PR together with your thoughts to Daniel Miessler
