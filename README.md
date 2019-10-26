# Atlas

This is a simple controller to read data from Azure Instance Metadata Service (IMDS) and write it to a config map in every Kubernetes namespace.

## Development

Run setup.ps1 to grab a few dependencies. Currently there's a [Taskfile][taskfile] with common tasks captured.

### Tasks

Install the Taskfile runner (go-task). It will be installed already if you ran setup.ps1.

```
# set $env:GO111MODULE="on" if necessary
go get github.com/go-task/task/v2/cmd/task@v2.7.0
```

Build:
```
task build
```

All tests (**will spin up a k8s docker cluster**):
```
# alias for `task e2e unit`
task test
```

Unit tests
```
task unit
```

E2E tests
```
task e2e
```

[taskfile]: https://taskfile.dev/#/