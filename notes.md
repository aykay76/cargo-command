https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/

`brew upgrade operator-sdk`

`operator-sdk init --domain cargo-command.com --repo github.com/aykay76/cargo-command --plugins=go/v4`

`operator-sdk create api --group simulation --version v1alpha1 --kind Simulation --resource --controller --plugins=go/v4`
Add additional fields to spec and status for simulation types in `api/v1alpha1/simulation_types.go`
Add creation/cleanup logic to `internal/controller/simulation_controller.go`
After modifying types file always run:
`make generate`
`make manifests`

To build the image:

`make docker-build docker-push`

To create flux manifests:
`make deploy`

Then when the operator is running, create a CR:
https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/#create-a-memcached-cr

Sample CRs is in `config/samples`

apply it using kubectl apply



To check images in Kind cluster:
`kubectl get nodes`
Should yield something like:
```
NAME                 STATUS   ROLES           AGE   VERSION
kind-control-plane   Ready    control-plane   26h   v1.31.2
```

Then execute:
`podman exec -it kind-control-plane bash`
`crictl images`

