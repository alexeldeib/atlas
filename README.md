# Atlas

This is a simple controller to read data from Azure Instance Metadata Service (IMDS) and write it to a config map in every Kubernetes namespace. The idea would be that we can use data from IMDS to bootstrap our regional differentation and hand off to app developers for configuration tailored to their specific scenario.

Here's an example of what the config map written out currently looks like: 

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: region-config
  namespace: default
data:
  env: int
  geography: unitedstates
  location: eastus2
  resourceGroup: foo-bar
  shortName: eus2
```

Here is an example of almost all of the IMDS return body:

```json
{
    "compute": {
        "azEnvironment": "AzurePublicCloud",
        "customData": "HUGE_BASE64_BLOB",
        "location": "eastus2",
        "name": "baz",
        "offer": "UbuntuServer",
        "osType": "Linux",
        "placementGroupId": "foo",
        "plan": {
            "name": "",
            "product": "",
            "publisher": ""
        },
        "platformFaultDomain": "0",
        "platformUpdateDomain": "0",
        "provider": "Microsoft.Compute",
        "publicKeys": [
            {
                "keyData": "SSH_PUB_KEY",
                "path": "/home/azureuser/.ssh/authorized_keys"
            }
        ],
        "publisher": "Canonical",
        "resourceGroupName": "bar",
        "resourceId": "/subscriptions/foo/resourceGroups/bar/providers/Microsoft.Compute/virtualMachines/baz",
        "sku": "18.04-LTS",
        "subscriptionId": "foo",
        "tags": "acsengineVersion:v0.0.0;clusterName:bar;orchestrator:kubernetes",
        "version": "18.04.201909030",
        "vmId": "foo",
        "vmScaleSetName": "default-ae90a13e-5da68c2c",
        "vmSize": "Standard_E4s_v3",
        "zone": "1"
    }
}
```

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