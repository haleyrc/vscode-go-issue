# vscode-go-issue

This is a minimal example demonstrating an output issue when using the "run test"
shortcut from a Go test file using the Go extension for VS Code.

## Description

In an effort to make file links clickable in the VS Code output panel, the Go
extension converts short filenames to absolute paths. When the filename is not
the first element on the log line, however, the additional path information is
still prepended to the line, as shown below:

```bash
Running tool: /usr/local/go/bin/go test -v -tags=integration -coverprofile=/tmp/vscode-goOzDjv1/go-code-cover -timeout 30s github.com/haleyrc/vscode-go-issue

=== RUN   TestDemoLogIssue
/home/ryan/go/src/github.com/haleyrc/vscode-go-issue/2018/10/17 11:38:59.141114 /home/ryan/go/src/github.com/haleyrc/vscode-go-issue/demo.go:11: hello
--- PASS: TestDemoLogIssue (0.00s)
PASS
coverage: 100.0% of statements
ok  	github.com/haleyrc/vscode-go-issue	0.001s	coverage: 100.0% of statements
Success: Tests passed.
```

The same test run from the terminal produces the desired output:

```bash
=== RUN   TestDemoLogIssue
2018/10/17 11:39:44.000783 /home/ryan/go/src/github.com/haleyrc/vscode-go-issue/demo.go:11: hello
--- PASS: TestDemoLogIssue (0.00s)
PASS
ok      github.com/haleyrc/vscode-go-issue      0.001s
```

## Environment

- Version: 1.28.1
- Commit: 3368db6750222d319c851f6d90eb619d886e08f5
- Date: 2018-10-11T18:09:20.566Z
- Electron: 2.0.9
- Chrome: 61.0.3163.100
- Node.js: 8.9.3
- V8: 6.1.534.41
- Architecture: x64
- Extenstion Version: 0.69.1
