# fun-redirect

## File Descriptions - 
- `fun.reg` -> Registry entry to "create" `fun` protocol such that all `fun://` entries redirect to your specified executable file possibly with arguments - `%1`, `%2`, etc 
- `main.go` -> Http redirect server such that it redirects all `/xyz` paths sent to `http://localhost` to a pre-defined url such as `https://youtube.com`
- `redirect.bat` -> Batch script specified as executable when via `fun://` protocol call - extracts the parameter path ( suppose - `/xyz` ) and redirects to localhost with that path appended
- `go.mod`,`go.sum` -> Go files to specify language version and dependencies, etc