// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scaleway

import (
	"context"
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/dirien/pulumi-scaleway/provider/v2/pkg/version"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/scaleway/terraform-provider-scaleway/v2/shim"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "scaleway"
	// modules:
	mainMod = "index" // the scaleway module
)

//go:embed cmd/pulumi-resource-scaleway/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider

	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(shim.NewProvider()),
		shim.Framework()(),
	)

	prov := tfbridge.ProviderInfo{
		P:       p,
		Version: version.Version,
		Name:    "scaleway",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Scaleway",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "dirien",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/dirien/pulumi-scaleway",
		Description:       "A Pulumi package for creating and managing Scaleway resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:     []string{"pulumi", "scaleway", "category/utility"},
		License:      "Apache-2.0",
		Homepage:     "https://www.pulumi.com",
		Repository:   "https://github.com/dirien/pulumi-scaleway",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "scaleway",
		Config: map[string]*tfbridge.SchemaInfo{
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_ACCESS_KEY"},
				},
				Secret: tfbridge.BoolRef(true),
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_SECRET_KEY"},
				},
				Secret: tfbridge.BoolRef(true),
			},
			"project_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_PROJECT_ID"},
				},
			},
			"organization_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ORGANIZATION_ID"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_REGION"},
				},
			},
			"zone": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ZONE"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"scaleway_object": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectItem")},
		},
		DataSources:             map[string]*tfbridge.DataSourceInfo{},
		TFProviderModuleVersion: "v2",
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@ediri/scaleway",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "ediri_scaleway",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/dirien/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "ediri",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "io.dirien",
		},
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("scaleway_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))

	prov.SetAutonaming(255, "-")

	return prov
}
