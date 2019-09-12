# riff System

Controllers for riff CRDs

## CRDs

riff System contains three API groups with CustomResourceDefinitions that are the public API for riff.

- `build.projectriff.io/v1alpha1`
  - `Application` - applications built from source using application buildpacks
  - `Function` - functions built from source using function buildpacks
  - `Container` - watch a container repository for the latest image
- `core.projectriff.io/v1alpha1`
  - `Deployer` - deployers map HTTP requests to applications, functions, containers or images with Kubernetes core resources
- `streaming.projectriff.io/v1alpha1`
  - `Provider` - an abstraction over a message broker
  - `Stream` - streams of messages
  - `Processor` - processors apply functions to messages on streams
- `knative.projectriff.io/v1alpha1`
  - `Adapter` - adapters map applications, functions or container images into an existing Knative Service or Configuration.
  - `Deployer` - deployers map HTTP requests to applications, functions, containers or images with Knative

### Runtime

A `controller` and `webhook` Deployments exist in the `riff-system` namespace to validate and reconcile the riff CRDs.

### RBAC

Two ClusterRoles are defined to grant access to the riff CRDs.

- `projectriff` - read/write access to all riff CRDs
- `projectriff-readonly` - read access to all riff CRDs

These roles are aggregated to the `edit` and `view` ClusterRoles respectively.

See the Kuberneties [Using RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) for more information.

## Install

riff System is not installed directly for typical use. See the [Getting Started](https://projectriff.io/docs/getting-started/) docs to learn how to install riff.

## Development

After making any changes to source files in `./pkgs/apis` it is necessary to regenerate the API client by running:

```sh
./hack/update-codegen.sh
```

Dependencies are managed with dep. After importing a new package or modifying Gopkg.toml, run:

```sh
./hack/update-deps.sh
```

To run the unit tests locally:

```sh
go test ./...
```

To deploy to a development cluster with [ko](https://github.com/google/ko):

```sh
ko apply -f config/
```

Additional dependencies must be installed independently into the cluster including:

- riff build templates
- Knative Build
- Knative Serving
- Istio

A common practice is to start with a standard riff install and then incrementally update riff System from source.

### Releases

Releases are generated by the CI server by running `./hack/release.sh`, and published to:

> `https://storage.googleapis.com/projectriff/riff-system/riff-system-{version}.yaml`

## Code of Conduct

Please refer to the [Contributor Code of Conduct](CODE_OF_CONDUCT.adoc).
