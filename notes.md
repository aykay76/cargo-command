https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/

`brew upgrade operator-sdk`

`operator-sdk init --domain cargo-command.com --repo github.com/aykay76/cargo-command --plugins=go/v4`

`operator-sdk create api --group simulation --version v1alpha1 --kind Simulation --resource --controller`