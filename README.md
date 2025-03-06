# cloudcontroller

Kubernetes Operator for deploying Cloud resources via Kubernetes CRDs. Written from scratch using controller-runtime library.

## Instructions:

### Common

- Register all new types in `cmd/cloudcontroller/api/v1`

- Controllers (Reconcilers) code found in `cmd/cloudcontroller/pkg/controllers`

- Cloudprovider APIs interaction found in `cmd/cloudcontroller/pkg/cloudproviders`

### Makefile:

- `IMG=<image>:<tag> make build`: Build CloudController Operator Docker image
- `IMG=<image>:<tag> make push`: Push CloudController Operator Docker image
- `IMG=<image>:<tag> make build-push`: Build and push CloudController Operator Docker image
- `make mock`: Generate mocks
- `make generate`: Generate CRDs
- `make install-crds`: Install CRDs
- `make deploy`: Deploy CloudController Operator

## Clouds

### GCP

- Place GCP Service Account credentials JSON file in `creds` secret
## AWS

 TBD
 
## Azure

  TBD