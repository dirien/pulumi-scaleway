// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Kubernetes clusters. For more information, see [the documentation](https://developers.scaleway.com/en/products/k8s/api/).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			jack, err := scaleway.NewK8sCluster(ctx, "jack", &scaleway.K8sClusterArgs{
//				Version:                   pulumi.String("1.24.3"),
//				Cni:                       pulumi.String("cilium"),
//				DeleteAdditionalResources: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewK8sPool(ctx, "john", &scaleway.K8sPoolArgs{
//				ClusterId: jack.ID(),
//				NodeType:  pulumi.String("DEV1-M"),
//				Size:      pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Multicloud
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			henry, err := scaleway.NewK8sCluster(ctx, "henry", &scaleway.K8sClusterArgs{
//				Type:                      pulumi.String("multicloud"),
//				Version:                   pulumi.String("1.24.3"),
//				Cni:                       pulumi.String("kilo"),
//				DeleteAdditionalResources: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewK8sPool(ctx, "friendFromOuterSpace", &scaleway.K8sPoolArgs{
//				ClusterId: henry.ID(),
//				NodeType:  pulumi.String("external"),
//				Size:      pulumi.Int(0),
//				MinSize:   pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// For a detailed example of how to add or run Elastic Metal servers instead of instances on your cluster, please refer to this guide.
//
// ### With additional configuration
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			johnK8sCluster, err := scaleway.NewK8sCluster(ctx, "johnK8sCluster", &scaleway.K8sClusterArgs{
//				Description: pulumi.String("my awesome cluster"),
//				Version:     pulumi.String("1.24.3"),
//				Cni:         pulumi.String("calico"),
//				Tags: pulumi.StringArray{
//					pulumi.String("i'm an awesome tag"),
//					pulumi.String("yay"),
//				},
//				DeleteAdditionalResources: pulumi.Bool(false),
//				AutoscalerConfig: &scaleway.K8sClusterAutoscalerConfigArgs{
//					DisableScaleDown:             pulumi.Bool(false),
//					ScaleDownDelayAfterAdd:       pulumi.String("5m"),
//					Estimator:                    pulumi.String("binpacking"),
//					Expander:                     pulumi.String("random"),
//					IgnoreDaemonsetsUtilization:  pulumi.Bool(true),
//					BalanceSimilarNodeGroups:     pulumi.Bool(true),
//					ExpendablePodsPriorityCutoff: -5,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewK8sPool(ctx, "johnK8sPool", &scaleway.K8sPoolArgs{
//				ClusterId:   johnK8sCluster.ID(),
//				NodeType:    pulumi.String("DEV1-M"),
//				Size:        pulumi.Int(3),
//				Autoscaling: pulumi.Bool(true),
//				Autohealing: pulumi.Bool(true),
//				MinSize:     pulumi.Int(1),
//				MaxSize:     pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Kubernetes clusters can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/k8sCluster:K8sCluster mycluster fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type K8sCluster struct {
	pulumi.CustomResourceState

	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
	AdmissionPlugins pulumi.StringArrayOutput `pulumi:"admissionPlugins"`
	// Additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans pulumi.StringArrayOutput `pulumi:"apiserverCertSans"`
	// The URL of the Kubernetes API server.
	ApiserverUrl pulumi.StringOutput `pulumi:"apiserverUrl"`
	// The auto upgrade configuration.
	AutoUpgrade K8sClusterAutoUpgradeOutput `pulumi:"autoUpgrade"`
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfig K8sClusterAutoscalerConfigOutput `pulumi:"autoscalerConfig"`
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	// > **Important:** Updates to this field will recreate a new resource.
	Cni pulumi.StringOutput `pulumi:"cni"`
	// The creation date of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
	// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
	// If you prefer keeping it, you should instead set it as `false`.
	DeleteAdditionalResources pulumi.BoolOutput `pulumi:"deleteAdditionalResources"`
	// A description for the Kubernetes cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
	FeatureGates pulumi.StringArrayOutput `pulumi:"featureGates"`
	// The kubeconfig configuration file of the Kubernetes cluster
	Kubeconfigs K8sClusterKubeconfigArrayOutput `pulumi:"kubeconfigs"`
	// The name for the Kubernetes cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OpenID Connect configuration of the cluster
	OpenIdConnectConfig K8sClusterOpenIdConnectConfigOutput `pulumi:"openIdConnectConfig"`
	// The organization ID the cluster is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The ID of the private network of the cluster.
	//
	// > **Important:** This field can only be set at cluster creation and cannot be updated later.
	// Changes to this field will cause the cluster to be destroyed then recreated.
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the cluster is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region in which the cluster should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of the Kubernetes cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags associated with the Kubernetes cluster.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The last update date of the cluster.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Set to `true` if a newer Kubernetes version is available.
	UpgradeAvailable pulumi.BoolOutput `pulumi:"upgradeAvailable"`
	// The version of the Kubernetes cluster.
	Version pulumi.StringOutput `pulumi:"version"`
	// The DNS wildcard that points to all ready nodes.
	WildcardDns pulumi.StringOutput `pulumi:"wildcardDns"`
}

// NewK8sCluster registers a new resource with the given unique name, arguments, and options.
func NewK8sCluster(ctx *pulumi.Context,
	name string, args *K8sClusterArgs, opts ...pulumi.ResourceOption) (*K8sCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cni == nil {
		return nil, errors.New("invalid value for required argument 'Cni'")
	}
	if args.DeleteAdditionalResources == nil {
		return nil, errors.New("invalid value for required argument 'DeleteAdditionalResources'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"kubeconfigs",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource K8sCluster
	err := ctx.RegisterResource("scaleway:index/k8sCluster:K8sCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetK8sCluster gets an existing K8sCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetK8sCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *K8sClusterState, opts ...pulumi.ResourceOption) (*K8sCluster, error) {
	var resource K8sCluster
	err := ctx.ReadResource("scaleway:index/k8sCluster:K8sCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering K8sCluster resources.
type k8sClusterState struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
	AdmissionPlugins []string `pulumi:"admissionPlugins"`
	// Additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans []string `pulumi:"apiserverCertSans"`
	// The URL of the Kubernetes API server.
	ApiserverUrl *string `pulumi:"apiserverUrl"`
	// The auto upgrade configuration.
	AutoUpgrade *K8sClusterAutoUpgrade `pulumi:"autoUpgrade"`
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfig *K8sClusterAutoscalerConfig `pulumi:"autoscalerConfig"`
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	// > **Important:** Updates to this field will recreate a new resource.
	Cni *string `pulumi:"cni"`
	// The creation date of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
	// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
	// If you prefer keeping it, you should instead set it as `false`.
	DeleteAdditionalResources *bool `pulumi:"deleteAdditionalResources"`
	// A description for the Kubernetes cluster.
	Description *string `pulumi:"description"`
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
	FeatureGates []string `pulumi:"featureGates"`
	// The kubeconfig configuration file of the Kubernetes cluster
	Kubeconfigs []K8sClusterKubeconfig `pulumi:"kubeconfigs"`
	// The name for the Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The OpenID Connect configuration of the cluster
	OpenIdConnectConfig *K8sClusterOpenIdConnectConfig `pulumi:"openIdConnectConfig"`
	// The organization ID the cluster is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The ID of the private network of the cluster.
	//
	// > **Important:** This field can only be set at cluster creation and cannot be updated later.
	// Changes to this field will cause the cluster to be destroyed then recreated.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the cluster is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the cluster should be created.
	Region *string `pulumi:"region"`
	// The status of the Kubernetes cluster.
	Status *string `pulumi:"status"`
	// The tags associated with the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
	// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
	Type *string `pulumi:"type"`
	// The last update date of the cluster.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Set to `true` if a newer Kubernetes version is available.
	UpgradeAvailable *bool `pulumi:"upgradeAvailable"`
	// The version of the Kubernetes cluster.
	Version *string `pulumi:"version"`
	// The DNS wildcard that points to all ready nodes.
	WildcardDns *string `pulumi:"wildcardDns"`
}

type K8sClusterState struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
	AdmissionPlugins pulumi.StringArrayInput
	// Additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans pulumi.StringArrayInput
	// The URL of the Kubernetes API server.
	ApiserverUrl pulumi.StringPtrInput
	// The auto upgrade configuration.
	AutoUpgrade K8sClusterAutoUpgradePtrInput
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfig K8sClusterAutoscalerConfigPtrInput
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	// > **Important:** Updates to this field will recreate a new resource.
	Cni pulumi.StringPtrInput
	// The creation date of the cluster.
	CreatedAt pulumi.StringPtrInput
	// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
	// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
	// If you prefer keeping it, you should instead set it as `false`.
	DeleteAdditionalResources pulumi.BoolPtrInput
	// A description for the Kubernetes cluster.
	Description pulumi.StringPtrInput
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
	FeatureGates pulumi.StringArrayInput
	// The kubeconfig configuration file of the Kubernetes cluster
	Kubeconfigs K8sClusterKubeconfigArrayInput
	// The name for the Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The OpenID Connect configuration of the cluster
	OpenIdConnectConfig K8sClusterOpenIdConnectConfigPtrInput
	// The organization ID the cluster is associated with.
	OrganizationId pulumi.StringPtrInput
	// The ID of the private network of the cluster.
	//
	// > **Important:** This field can only be set at cluster creation and cannot be updated later.
	// Changes to this field will cause the cluster to be destroyed then recreated.
	PrivateNetworkId pulumi.StringPtrInput
	// `projectId`) The ID of the project the cluster is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the cluster should be created.
	Region pulumi.StringPtrInput
	// The status of the Kubernetes cluster.
	Status pulumi.StringPtrInput
	// The tags associated with the Kubernetes cluster.
	Tags pulumi.StringArrayInput
	// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
	Type pulumi.StringPtrInput
	// The last update date of the cluster.
	UpdatedAt pulumi.StringPtrInput
	// Set to `true` if a newer Kubernetes version is available.
	UpgradeAvailable pulumi.BoolPtrInput
	// The version of the Kubernetes cluster.
	Version pulumi.StringPtrInput
	// The DNS wildcard that points to all ready nodes.
	WildcardDns pulumi.StringPtrInput
}

func (K8sClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*k8sClusterState)(nil)).Elem()
}

type k8sClusterArgs struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
	AdmissionPlugins []string `pulumi:"admissionPlugins"`
	// Additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans []string `pulumi:"apiserverCertSans"`
	// The auto upgrade configuration.
	AutoUpgrade *K8sClusterAutoUpgrade `pulumi:"autoUpgrade"`
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfig *K8sClusterAutoscalerConfig `pulumi:"autoscalerConfig"`
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	// > **Important:** Updates to this field will recreate a new resource.
	Cni string `pulumi:"cni"`
	// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
	// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
	// If you prefer keeping it, you should instead set it as `false`.
	DeleteAdditionalResources bool `pulumi:"deleteAdditionalResources"`
	// A description for the Kubernetes cluster.
	Description *string `pulumi:"description"`
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
	FeatureGates []string `pulumi:"featureGates"`
	// The name for the Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The OpenID Connect configuration of the cluster
	OpenIdConnectConfig *K8sClusterOpenIdConnectConfig `pulumi:"openIdConnectConfig"`
	// The ID of the private network of the cluster.
	//
	// > **Important:** This field can only be set at cluster creation and cannot be updated later.
	// Changes to this field will cause the cluster to be destroyed then recreated.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the cluster is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the cluster should be created.
	Region *string `pulumi:"region"`
	// The tags associated with the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
	// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
	Type *string `pulumi:"type"`
	// The version of the Kubernetes cluster.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a K8sCluster resource.
type K8sClusterArgs struct {
	// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
	AdmissionPlugins pulumi.StringArrayInput
	// Additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans pulumi.StringArrayInput
	// The auto upgrade configuration.
	AutoUpgrade K8sClusterAutoUpgradePtrInput
	// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
	AutoscalerConfig K8sClusterAutoscalerConfigPtrInput
	// The Container Network Interface (CNI) for the Kubernetes cluster.
	// > **Important:** Updates to this field will recreate a new resource.
	Cni pulumi.StringInput
	// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
	// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
	// If you prefer keeping it, you should instead set it as `false`.
	DeleteAdditionalResources pulumi.BoolInput
	// A description for the Kubernetes cluster.
	Description pulumi.StringPtrInput
	// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
	FeatureGates pulumi.StringArrayInput
	// The name for the Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The OpenID Connect configuration of the cluster
	OpenIdConnectConfig K8sClusterOpenIdConnectConfigPtrInput
	// The ID of the private network of the cluster.
	//
	// > **Important:** This field can only be set at cluster creation and cannot be updated later.
	// Changes to this field will cause the cluster to be destroyed then recreated.
	PrivateNetworkId pulumi.StringPtrInput
	// `projectId`) The ID of the project the cluster is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the cluster should be created.
	Region pulumi.StringPtrInput
	// The tags associated with the Kubernetes cluster.
	Tags pulumi.StringArrayInput
	// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
	Type pulumi.StringPtrInput
	// The version of the Kubernetes cluster.
	Version pulumi.StringInput
}

func (K8sClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*k8sClusterArgs)(nil)).Elem()
}

type K8sClusterInput interface {
	pulumi.Input

	ToK8sClusterOutput() K8sClusterOutput
	ToK8sClusterOutputWithContext(ctx context.Context) K8sClusterOutput
}

func (*K8sCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sCluster)(nil)).Elem()
}

func (i *K8sCluster) ToK8sClusterOutput() K8sClusterOutput {
	return i.ToK8sClusterOutputWithContext(context.Background())
}

func (i *K8sCluster) ToK8sClusterOutputWithContext(ctx context.Context) K8sClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sClusterOutput)
}

// K8sClusterArrayInput is an input type that accepts K8sClusterArray and K8sClusterArrayOutput values.
// You can construct a concrete instance of `K8sClusterArrayInput` via:
//
//	K8sClusterArray{ K8sClusterArgs{...} }
type K8sClusterArrayInput interface {
	pulumi.Input

	ToK8sClusterArrayOutput() K8sClusterArrayOutput
	ToK8sClusterArrayOutputWithContext(context.Context) K8sClusterArrayOutput
}

type K8sClusterArray []K8sClusterInput

func (K8sClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8sCluster)(nil)).Elem()
}

func (i K8sClusterArray) ToK8sClusterArrayOutput() K8sClusterArrayOutput {
	return i.ToK8sClusterArrayOutputWithContext(context.Background())
}

func (i K8sClusterArray) ToK8sClusterArrayOutputWithContext(ctx context.Context) K8sClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sClusterArrayOutput)
}

// K8sClusterMapInput is an input type that accepts K8sClusterMap and K8sClusterMapOutput values.
// You can construct a concrete instance of `K8sClusterMapInput` via:
//
//	K8sClusterMap{ "key": K8sClusterArgs{...} }
type K8sClusterMapInput interface {
	pulumi.Input

	ToK8sClusterMapOutput() K8sClusterMapOutput
	ToK8sClusterMapOutputWithContext(context.Context) K8sClusterMapOutput
}

type K8sClusterMap map[string]K8sClusterInput

func (K8sClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8sCluster)(nil)).Elem()
}

func (i K8sClusterMap) ToK8sClusterMapOutput() K8sClusterMapOutput {
	return i.ToK8sClusterMapOutputWithContext(context.Background())
}

func (i K8sClusterMap) ToK8sClusterMapOutputWithContext(ctx context.Context) K8sClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8sClusterMapOutput)
}

type K8sClusterOutput struct{ *pulumi.OutputState }

func (K8sClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8sCluster)(nil)).Elem()
}

func (o K8sClusterOutput) ToK8sClusterOutput() K8sClusterOutput {
	return o
}

func (o K8sClusterOutput) ToK8sClusterOutputWithContext(ctx context.Context) K8sClusterOutput {
	return o
}

// The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
func (o K8sClusterOutput) AdmissionPlugins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringArrayOutput { return v.AdmissionPlugins }).(pulumi.StringArrayOutput)
}

// Additional Subject Alternative Names for the Kubernetes API server certificate
func (o K8sClusterOutput) ApiserverCertSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringArrayOutput { return v.ApiserverCertSans }).(pulumi.StringArrayOutput)
}

// The URL of the Kubernetes API server.
func (o K8sClusterOutput) ApiserverUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.ApiserverUrl }).(pulumi.StringOutput)
}

// The auto upgrade configuration.
func (o K8sClusterOutput) AutoUpgrade() K8sClusterAutoUpgradeOutput {
	return o.ApplyT(func(v *K8sCluster) K8sClusterAutoUpgradeOutput { return v.AutoUpgrade }).(K8sClusterAutoUpgradeOutput)
}

// The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
func (o K8sClusterOutput) AutoscalerConfig() K8sClusterAutoscalerConfigOutput {
	return o.ApplyT(func(v *K8sCluster) K8sClusterAutoscalerConfigOutput { return v.AutoscalerConfig }).(K8sClusterAutoscalerConfigOutput)
}

// The Container Network Interface (CNI) for the Kubernetes cluster.
// > **Important:** Updates to this field will recreate a new resource.
func (o K8sClusterOutput) Cni() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Cni }).(pulumi.StringOutput)
}

// The creation date of the cluster.
func (o K8sClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
// > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
// If you prefer keeping it, you should instead set it as `false`.
func (o K8sClusterOutput) DeleteAdditionalResources() pulumi.BoolOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.BoolOutput { return v.DeleteAdditionalResources }).(pulumi.BoolOutput)
}

// A description for the Kubernetes cluster.
func (o K8sClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
func (o K8sClusterOutput) FeatureGates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringArrayOutput { return v.FeatureGates }).(pulumi.StringArrayOutput)
}

// The kubeconfig configuration file of the Kubernetes cluster
func (o K8sClusterOutput) Kubeconfigs() K8sClusterKubeconfigArrayOutput {
	return o.ApplyT(func(v *K8sCluster) K8sClusterKubeconfigArrayOutput { return v.Kubeconfigs }).(K8sClusterKubeconfigArrayOutput)
}

// The name for the Kubernetes cluster.
func (o K8sClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The OpenID Connect configuration of the cluster
func (o K8sClusterOutput) OpenIdConnectConfig() K8sClusterOpenIdConnectConfigOutput {
	return o.ApplyT(func(v *K8sCluster) K8sClusterOpenIdConnectConfigOutput { return v.OpenIdConnectConfig }).(K8sClusterOpenIdConnectConfigOutput)
}

// The organization ID the cluster is associated with.
func (o K8sClusterOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the private network of the cluster.
//
// > **Important:** This field can only be set at cluster creation and cannot be updated later.
// Changes to this field will cause the cluster to be destroyed then recreated.
func (o K8sClusterOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the cluster is associated with.
func (o K8sClusterOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region in which the cluster should be created.
func (o K8sClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The status of the Kubernetes cluster.
func (o K8sClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the Kubernetes cluster.
func (o K8sClusterOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
func (o K8sClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last update date of the cluster.
func (o K8sClusterOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Set to `true` if a newer Kubernetes version is available.
func (o K8sClusterOutput) UpgradeAvailable() pulumi.BoolOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.BoolOutput { return v.UpgradeAvailable }).(pulumi.BoolOutput)
}

// The version of the Kubernetes cluster.
func (o K8sClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The DNS wildcard that points to all ready nodes.
func (o K8sClusterOutput) WildcardDns() pulumi.StringOutput {
	return o.ApplyT(func(v *K8sCluster) pulumi.StringOutput { return v.WildcardDns }).(pulumi.StringOutput)
}

type K8sClusterArrayOutput struct{ *pulumi.OutputState }

func (K8sClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8sCluster)(nil)).Elem()
}

func (o K8sClusterArrayOutput) ToK8sClusterArrayOutput() K8sClusterArrayOutput {
	return o
}

func (o K8sClusterArrayOutput) ToK8sClusterArrayOutputWithContext(ctx context.Context) K8sClusterArrayOutput {
	return o
}

func (o K8sClusterArrayOutput) Index(i pulumi.IntInput) K8sClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *K8sCluster {
		return vs[0].([]*K8sCluster)[vs[1].(int)]
	}).(K8sClusterOutput)
}

type K8sClusterMapOutput struct{ *pulumi.OutputState }

func (K8sClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8sCluster)(nil)).Elem()
}

func (o K8sClusterMapOutput) ToK8sClusterMapOutput() K8sClusterMapOutput {
	return o
}

func (o K8sClusterMapOutput) ToK8sClusterMapOutputWithContext(ctx context.Context) K8sClusterMapOutput {
	return o
}

func (o K8sClusterMapOutput) MapIndex(k pulumi.StringInput) K8sClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *K8sCluster {
		return vs[0].(map[string]*K8sCluster)[vs[1].(string)]
	}).(K8sClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*K8sClusterInput)(nil)).Elem(), &K8sCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8sClusterArrayInput)(nil)).Elem(), K8sClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8sClusterMapInput)(nil)).Elem(), K8sClusterMap{})
	pulumi.RegisterOutputType(K8sClusterOutput{})
	pulumi.RegisterOutputType(K8sClusterArrayOutput{})
	pulumi.RegisterOutputType(K8sClusterMapOutput{})
}
