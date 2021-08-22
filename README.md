# Semantic Release Go Commit Lint Git Hook

This library is a git hook to validate commit messages
following Semantic Release convention.

If the git commit does not follow Semantic Release convention, the commit is aborted,
and a description with the reasons is returned.

## Install

```bash
go build -o pre-commit
# move the binary to your:
# project-name/.git/hooks/
```

## Use

```bash
git commit -m 'unknown-type: hello'
```

It will throw an error with the root cause.


