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
	"fmt"
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"path/filepath"

	"github.com/dirien/pulumi-scaleway/provider/v2/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/scaleway/terraform-provider-scaleway/v2/scaleway"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "scaleway"
	// modules:
	mainMod = "index" // the scaleway module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	providerConfig := scaleway.ProviderConfig{}
	p := shimv2.NewProvider(scaleway.Provider(&providerConfig)())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "scaleway",
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
		Keywords:   []string{"pulumi", "scaleway", "category/utility"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/dirien/pulumi-scaleway",
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
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"scaleway_account_project":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AccountProject")},
			"scaleway_account_ssh_key":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AccountSshKey")},
			"scaleway_apple_silicon_server":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AppleSiliconServer")},
			"scaleway_baremetal_server":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BaremetalServer")},
			"scaleway_cockpit":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Cockpit")},
			"scaleway_cockpit_grafana_user":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CockpitGrafanaUser")},
			"scaleway_cockpit_token":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CockpitToken")},
			"scaleway_container":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Container")},
			"scaleway_container_cron":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContainerCron")},
			"scaleway_container_domain":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContainerDomain")},
			"scaleway_container_namespace":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContainerNamespace")},
			"scaleway_container_token":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContainerToken")},
			"scaleway_domain_record":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainRecord")},
			"scaleway_domain_zone":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DomainZone")},
			"scaleway_flexible_ip":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FlexibleIp")},
			"scaleway_function":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Function")},
			"scaleway_function_cron":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FunctionCron")},
			"scaleway_function_domain":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FunctionDomain")},
			"scaleway_function_namespace":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FunctionNamespace")},
			"scaleway_function_token":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FunctionToken")},
			"scaleway_iam_api_key":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamApiKey")},
			"scaleway_iam_user":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamUser")},
			"scaleway_iam_application":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamApplication")},
			"scaleway_iam_group":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamGroup")},
			"scaleway_iam_policy":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamPolicy")},
			"scaleway_iam_ssh_key":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamSshKey")},
			"scaleway_instance_image":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceImage")},
			"scaleway_instance_ip":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceIp")},
			"scaleway_instance_ip_reverse_dns":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceIpReverseDns")},
			"scaleway_instance_placement_group":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstancePlacementGroup")},
			"scaleway_instance_private_nic":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstancePrivateNic")},
			"scaleway_instance_security_group":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceSecurityGroup")},
			"scaleway_instance_security_group_rules":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceSecurityGroupRules")},
			"scaleway_instance_server":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceServer")},
			"scaleway_instance_snapshot":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceSnapshot")},
			"scaleway_instance_user_data":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceUserData")},
			"scaleway_instance_volume":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceVolume")},
			"scaleway_iot_device":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IotDevice")},
			"scaleway_iot_hub":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IotHub")},
			"scaleway_iot_network":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IotNetwork")},
			"scaleway_iot_route":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IotRoute")},
			"scaleway_k8s_cluster":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "K8sCluster")},
			"scaleway_k8s_pool":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "K8sPool")},
			"scaleway_lb":                                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Lb")},
			"scaleway_lb_backend":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbBackend")},
			"scaleway_lb_certificate":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbCertificate")},
			"scaleway_lb_frontend":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbFrontend")},
			"scaleway_lb_ip":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbIp")},
			"scaleway_lb_route":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbRoute")},
			"scaleway_mnq_credential":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqCredential")},
			"scaleway_mnq_namespace":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqNamespace")},
			"scaleway_mnq_queue":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqQueue")},
			"scaleway_object":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectItem")},
			"scaleway_object_bucket":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectBucket")},
			"scaleway_object_bucket_acl":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectBucketAcl")},
			"scaleway_object_bucket_lock_configuration":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectBucketLockConfiguration")},
			"scaleway_object_bucket_policy":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectBucketPolicy")},
			"scaleway_object_bucket_website_configuration": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectBucketWebsiteConfiguration")},
			"scaleway_rdb_acl":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbAcl")},
			"scaleway_rdb_database":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbDatabase")},
			"scaleway_rdb_database_backup":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbDatabaseBackup")},
			"scaleway_rdb_instance":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbInstance")},
			"scaleway_rdb_privilege":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbPrivilege")},
			"scaleway_rdb_read_replica":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbReadReplica")},
			"scaleway_rdb_user":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RdbUser")},
			"scaleway_redis_cluster":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RedisCluster")},
			"scaleway_registry_namespace":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RegistryNamespace")},
			"scaleway_secret":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Secret")},
			"scaleway_secret_version":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SecretVersion")},
			"scaleway_tem_domain":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TemDomain")},
			"scaleway_vpc_gateway_network":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcGatewayNetwork")},
			"scaleway_vpc_private_network":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPrivateNetwork")},
			"scaleway_vpc_public_gateway":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGateway")},
			"scaleway_vpc_public_gateway_dhcp":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGatewayDhcp")},
			"scaleway_vpc_public_gateway_dhcp_reservation": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGatewayDhcpReservation")},
			"scaleway_vpc_public_gateway_ip":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGatewayIp")},
			"scaleway_vpc_public_gateway_ip_reverse_dns":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGatewayIpReverseDns")},
			"scaleway_vpc_public_gateway_pat_rule":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "VpcPublicGatewayPatRule")},
			"scaleway_function_trigger":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FunctionTrigger")},
			"scaleway_lb_acl":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LbAcl")},
			"scaleway_vpc":                                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vpc")},
			"scaleway_container_trigger":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContainerTrigger")},
			"scaleway_webhosting":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "WebHosting")},
			"scaleway_iam_group_membership":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamGroupMembership")},
			"scaleway_flexible_ip_mac_address":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FlexibleIpMacAddress")},
			"scaleway_documentdb_database":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBDatabase")},
			"scaleway_documentdb_instance":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBInstance")},
			"scaleway_documentdb_private_network_endpoint": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBPrivateNetworkEndpoint")},
			"scaleway_documentdb_privilege":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBPrivilege")},
			"scaleway_documentdb_read_replica":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBReadReplica")},
			"scaleway_documentdb_user":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DocumentDBUser")},
			"scaleway_mnq_nats_account":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqNatsAccount")},
			"scaleway_mnq_nats_credentials":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqNatsCredentials")},
			"scaleway_mnq_sqs":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqSqs")},
			"scaleway_mnq_sqs_credentials":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqSqsCredentials")},
			"scaleway_mnq_sqs_queue":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MnqSqsQueue")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"scaleway_account_project":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccountProject")},
			"scaleway_account_ssh_key":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccountSshKey")},
			"scaleway_baremetal_offer":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBaremetalOffer")},
			"scaleway_baremetal_option":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBaremetalOption")},
			"scaleway_baremetal_os":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBaremetalOs")},
			"scaleway_baremetal_server":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBaremetalServer")},
			"scaleway_cockpit":                             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCockpit")},
			"scaleway_container":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getContainer")},
			"scaleway_container_namespace":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getContainerNamespace")},
			"scaleway_domain_record":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDomainRecord")},
			"scaleway_domain_zone":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDomainZone")},
			"scaleway_function":                            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFunction")},
			"scaleway_flexible_ip":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFlexibleIp")},
			"scaleway_function_namespace":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFunctionNamespace")},
			"scaleway_iam_application":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIamApplication")},
			"scaleway_iam_group":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIamGroup")},
			"scaleway_iam_ssh_key":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIamSshKey")},
			"scaleway_iam_user":                            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIamUser")},
			"scaleway_instance_image":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceImage")},
			"scaleway_instance_ip":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceIp")},
			"scaleway_instance_private_nic":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstancePrivateNic")},
			"scaleway_instance_security_group":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceSecurityGroup")},
			"scaleway_instance_server":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceServer")},
			"scaleway_instance_servers":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceServers")},
			"scaleway_instance_snapshot":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceSnapshot")},
			"scaleway_instance_volume":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceVolume")},
			"scaleway_iot_device":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIotDevice")},
			"scaleway_iot_hub":                             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIotHub")},
			"scaleway_k8s_cluster":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getK8sCluster")},
			"scaleway_k8s_pool":                            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getK8sPool")},
			"scaleway_k8s_version":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getK8sVersion")},
			"scaleway_lb":                                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLb")},
			"scaleway_lb_acls":                             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbAcls")},
			"scaleway_lb_backend":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbBackend")},
			"scaleway_lb_backends":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbBackends")},
			"scaleway_lb_certificate":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbCertificate")},
			"scaleway_lb_frontend":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbFrontend")},
			"scaleway_lb_frontends":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbFrontends")},
			"scaleway_lb_ip":                               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbIp")},
			"scaleway_lb_ips":                              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbIps")},
			"scaleway_lb_route":                            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbRoute")},
			"scaleway_lb_routes":                           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbRoutes")},
			"scaleway_lbs":                                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLbs")},
			"scaleway_marketplace_image":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMarketplaceImage")},
			"scaleway_object_bucket":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getObjectBucket")},
			"scaleway_object_bucket_policy":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getObjectBucketPolicy")},
			"scaleway_rdb_acl":                             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRdbAcl")},
			"scaleway_rdb_database":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRdbDatabase")},
			"scaleway_rdb_database_backup":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRdbDatabaseBackup")},
			"scaleway_rdb_instance":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRdbInstance")},
			"scaleway_rdb_privilege":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRdbPrivilege")},
			"scaleway_redis_cluster":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRedisCluster")},
			"scaleway_registry_image":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegistryImage")},
			"scaleway_registry_namespace":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegistryNamespace")},
			"scaleway_secret":                              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSecret")},
			"scaleway_secret_version":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSecretVersion")},
			"scaleway_tem_domain":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTemDomain")},
			"scaleway_vpc_gateway_network":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcGatewayNetwork")},
			"scaleway_vpc_private_network":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPrivateNetwork")},
			"scaleway_vpc_public_gateway":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPublicGateway")},
			"scaleway_vpc_public_gateway_dhcp":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPublicGatewayDhcp")},
			"scaleway_vpc_public_gateway_dhcp_reservation": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPublicGatewayDhcpReservation")},
			"scaleway_vpc_public_gateway_ip":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPublicGatewayIp")},
			"scaleway_vpc_public_gateway_pat_rule":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcPublicGatewayPatRule")},
			"scaleway_webhosting_offer":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebHostOffer")},
			"scaleway_availability_zones":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAvailabilityZones")},
			"scaleway_cockpit_plan":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCockpitPlan")},
			"scaleway_vpc":                                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpc")},
			"scaleway_vpcs":                                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpcs")},
			"scaleway_webhosting":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebhosting")},
			"scaleway_flexible_ips":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFlexibleIps")},
			"scaleway_ipam_ip":                             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIpamIp")},
			"scaleway_documentdb_database":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDocumentDBDatabase")},
			"scaleway_documentdb_instance":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDocumentDBInstance")},
			"scaleway_documentdb_load_balancer_endpoint":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDocumentDBLoadBalancerEndpoint")},
			"scaleway_billing_consumptions":				{Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBillingConsumptions")},
			"scaleway_billing_invoices":					{Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBillingInvoices")},
			"scaleway_mnq_sqs":								{Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMnqSqs")},
		},
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
		}, MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}
	err := x.AutoAliasing(&prov, prov.GetMetadata())
	contract.AssertNoErrorf(err, "auto aliasing apply failed")

	prov.SetAutonaming(255, "-")

	return prov
}

//go:embed cmd/pulumi-resource-scaleway/bridge-metadata.json
var metadata []byte
