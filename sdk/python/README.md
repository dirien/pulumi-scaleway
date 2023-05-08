# Scaleway Resource Provider

The Scaleway Resource Provider lets you manage [Scaleway](https://www.scaleway.com/en/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)


To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @ediri/scaleway
```

or `yarn`:

```bash
yarn add @ediri/scaleway
```

### Python

To use from Python, install using `pip`:

```bash
pip install ediri-scaleway
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/dirien/pulumi-scaleway/sdk/v2
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package ediri.Scaleway
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Scaleway provider, you need to have Scaleway credentials. Scaleway maintains
documentation on how to create API
keys [here](https://www.scaleway.com/en/docs/console/my-project/how-to/generate-api-key/)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Scaleway:

```bash
export SCW_ACCESS_KEY=<SCW_ACCESS_KEY>
export SCW_SECRET_KEY=<SCW_SECRET_KEY>
```

```powershell
$env:SCW_ACCESS_KEY = "<SCW_ACCESS_KEY>"
$env:SCW_SECRET_KEY = "<SCW_SECRET_KEY>"
```

## Configuration Options

Use `pulumi config set scaleway:<option>` to set the following configuration options.

| Option            | Environment Variables         | Required/Optional | Description                                                                                                                                                                                        |
|-------------------|-------------------------------|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `access_key`      | `SCW_ACCESS_KEY`              | Required          | [Scaleway access key](https://console.scaleway.com/project/credentials)                                                                                                                            |
| `secret_key`      | `SCW_SECRET_KEY`              | Required          | [Scaleway secret key](https://console.scaleway.com/project/credentials)                                                                                                                            |
| `project_id`      | `SCW_DEFAULT_PROJECT_ID`      | Required          | The [project ID](https://console.scaleway.com/project/settings) that will be used as default value for all resources.                                                                              |
| `organization_id` | `SCW_DEFAULT_ORGANIZATION_ID` | Optional          | The [organization ID](https://console.scaleway.com/organization/settings) that will be used as default value for all resources.                                                                    | 
| `region`          | `SCW_DEFAULT_REGION`          | Optional          | The [region](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#regions) that will be used as default value for all resources. (`fr-par` if none specified) |
| `zone`            | `SCW_DEFAULT_ZONE`            | Optional          | The [zone](https://registry.terraform.io/providers/scaleway/scaleway/latest/guides/regions_and_zones#zones) that will be used as default value for all resources. (`fr-par-1` if none specified)   |
