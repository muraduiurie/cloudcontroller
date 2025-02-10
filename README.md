# cloudcontroller

Kubernetes Operator for deploying Cloud resources via Kubernetes CRDs

## GCP

#### Instructions:

- Place GCP Service Account credentials JSON file in `creds` secret

#### Makefile commands:

- `IMG=<image>:<tag> make build`: Build CloudController Operator Docker image
- `make generate`: generate CRDs