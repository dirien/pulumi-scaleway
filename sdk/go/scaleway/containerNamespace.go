// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway Serverless Container Namespace.
// For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#namespaces-cdce79).
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
//			_, err := scaleway.NewContainerNamespace(ctx, "main", &scaleway.ContainerNamespaceArgs{
//				Description: pulumi.String("Main container namespace"),
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
// Namespaces can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/containerNamespace:ContainerNamespace main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type ContainerNamespace struct {
	pulumi.CustomResourceState

	// The description of the namespace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Destroy registry on deletion
	//
	// Deprecated: Registry namespace is automatically destroyed with namespace
	DestroyRegistry pulumi.BoolPtrOutput `pulumi:"destroyRegistry"`
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// The unique name of the container namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint pulumi.StringOutput `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId pulumi.StringOutput `pulumi:"registryNamespaceId"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapOutput `pulumi:"secretEnvironmentVariables"`
}

// NewContainerNamespace registers a new resource with the given unique name, arguments, and options.
func NewContainerNamespace(ctx *pulumi.Context,
	name string, args *ContainerNamespaceArgs, opts ...pulumi.ResourceOption) (*ContainerNamespace, error) {
	if args == nil {
		args = &ContainerNamespaceArgs{}
	}

	if args.SecretEnvironmentVariables != nil {
		args.SecretEnvironmentVariables = pulumi.ToSecret(args.SecretEnvironmentVariables).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretEnvironmentVariables",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerNamespace
	err := ctx.RegisterResource("scaleway:index/containerNamespace:ContainerNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerNamespace gets an existing ContainerNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerNamespaceState, opts ...pulumi.ResourceOption) (*ContainerNamespace, error) {
	var resource ContainerNamespace
	err := ctx.ReadResource("scaleway:index/containerNamespace:ContainerNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerNamespace resources.
type containerNamespaceState struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Destroy registry on deletion
	//
	// Deprecated: Registry namespace is automatically destroyed with namespace
	DestroyRegistry *bool `pulumi:"destroyRegistry"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The unique name of the container namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name *string `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint *string `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId *string `pulumi:"registryNamespaceId"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

type ContainerNamespaceState struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Destroy registry on deletion
	//
	// Deprecated: Registry namespace is automatically destroyed with namespace
	DestroyRegistry pulumi.BoolPtrInput
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapInput
	// The unique name of the container namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringPtrInput
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
	// The registry endpoint of the namespace.
	RegistryEndpoint pulumi.StringPtrInput
	// The registry namespace ID of the namespace.
	RegistryNamespaceId pulumi.StringPtrInput
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapInput
}

func (ContainerNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerNamespaceState)(nil)).Elem()
}

type containerNamespaceArgs struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Destroy registry on deletion
	//
	// Deprecated: Registry namespace is automatically destroyed with namespace
	DestroyRegistry *bool `pulumi:"destroyRegistry"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The unique name of the container namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

// The set of arguments for constructing a ContainerNamespace resource.
type ContainerNamespaceArgs struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Destroy registry on deletion
	//
	// Deprecated: Registry namespace is automatically destroyed with namespace
	DestroyRegistry pulumi.BoolPtrInput
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapInput
	// The unique name of the container namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapInput
}

func (ContainerNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerNamespaceArgs)(nil)).Elem()
}

type ContainerNamespaceInput interface {
	pulumi.Input

	ToContainerNamespaceOutput() ContainerNamespaceOutput
	ToContainerNamespaceOutputWithContext(ctx context.Context) ContainerNamespaceOutput
}

func (*ContainerNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerNamespace)(nil)).Elem()
}

func (i *ContainerNamespace) ToContainerNamespaceOutput() ContainerNamespaceOutput {
	return i.ToContainerNamespaceOutputWithContext(context.Background())
}

func (i *ContainerNamespace) ToContainerNamespaceOutputWithContext(ctx context.Context) ContainerNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerNamespaceOutput)
}

func (i *ContainerNamespace) ToOutput(ctx context.Context) pulumix.Output[*ContainerNamespace] {
	return pulumix.Output[*ContainerNamespace]{
		OutputState: i.ToContainerNamespaceOutputWithContext(ctx).OutputState,
	}
}

// ContainerNamespaceArrayInput is an input type that accepts ContainerNamespaceArray and ContainerNamespaceArrayOutput values.
// You can construct a concrete instance of `ContainerNamespaceArrayInput` via:
//
//	ContainerNamespaceArray{ ContainerNamespaceArgs{...} }
type ContainerNamespaceArrayInput interface {
	pulumi.Input

	ToContainerNamespaceArrayOutput() ContainerNamespaceArrayOutput
	ToContainerNamespaceArrayOutputWithContext(context.Context) ContainerNamespaceArrayOutput
}

type ContainerNamespaceArray []ContainerNamespaceInput

func (ContainerNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerNamespace)(nil)).Elem()
}

func (i ContainerNamespaceArray) ToContainerNamespaceArrayOutput() ContainerNamespaceArrayOutput {
	return i.ToContainerNamespaceArrayOutputWithContext(context.Background())
}

func (i ContainerNamespaceArray) ToContainerNamespaceArrayOutputWithContext(ctx context.Context) ContainerNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerNamespaceArrayOutput)
}

func (i ContainerNamespaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*ContainerNamespace] {
	return pulumix.Output[[]*ContainerNamespace]{
		OutputState: i.ToContainerNamespaceArrayOutputWithContext(ctx).OutputState,
	}
}

// ContainerNamespaceMapInput is an input type that accepts ContainerNamespaceMap and ContainerNamespaceMapOutput values.
// You can construct a concrete instance of `ContainerNamespaceMapInput` via:
//
//	ContainerNamespaceMap{ "key": ContainerNamespaceArgs{...} }
type ContainerNamespaceMapInput interface {
	pulumi.Input

	ToContainerNamespaceMapOutput() ContainerNamespaceMapOutput
	ToContainerNamespaceMapOutputWithContext(context.Context) ContainerNamespaceMapOutput
}

type ContainerNamespaceMap map[string]ContainerNamespaceInput

func (ContainerNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerNamespace)(nil)).Elem()
}

func (i ContainerNamespaceMap) ToContainerNamespaceMapOutput() ContainerNamespaceMapOutput {
	return i.ToContainerNamespaceMapOutputWithContext(context.Background())
}

func (i ContainerNamespaceMap) ToContainerNamespaceMapOutputWithContext(ctx context.Context) ContainerNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerNamespaceMapOutput)
}

func (i ContainerNamespaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContainerNamespace] {
	return pulumix.Output[map[string]*ContainerNamespace]{
		OutputState: i.ToContainerNamespaceMapOutputWithContext(ctx).OutputState,
	}
}

type ContainerNamespaceOutput struct{ *pulumi.OutputState }

func (ContainerNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerNamespace)(nil)).Elem()
}

func (o ContainerNamespaceOutput) ToContainerNamespaceOutput() ContainerNamespaceOutput {
	return o
}

func (o ContainerNamespaceOutput) ToContainerNamespaceOutputWithContext(ctx context.Context) ContainerNamespaceOutput {
	return o
}

func (o ContainerNamespaceOutput) ToOutput(ctx context.Context) pulumix.Output[*ContainerNamespace] {
	return pulumix.Output[*ContainerNamespace]{
		OutputState: o.OutputState,
	}
}

// The description of the namespace.
func (o ContainerNamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Destroy registry on deletion
//
// Deprecated: Registry namespace is automatically destroyed with namespace
func (o ContainerNamespaceOutput) DestroyRegistry() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.BoolPtrOutput { return v.DestroyRegistry }).(pulumi.BoolPtrOutput)
}

// The environment variables of the namespace.
func (o ContainerNamespaceOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringMapOutput { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The unique name of the container namespace.
//
// > **Important** Updates to `name` will recreate the namespace.
func (o ContainerNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the namespace is associated with.
func (o ContainerNamespaceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the namespace is associated with.
func (o ContainerNamespaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region in which the namespace should be created.
func (o ContainerNamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The registry endpoint of the namespace.
func (o ContainerNamespaceOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

// The registry namespace ID of the namespace.
func (o ContainerNamespaceOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringOutput { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

// The secret environment variables of the namespace.
func (o ContainerNamespaceOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerNamespace) pulumi.StringMapOutput { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

type ContainerNamespaceArrayOutput struct{ *pulumi.OutputState }

func (ContainerNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerNamespace)(nil)).Elem()
}

func (o ContainerNamespaceArrayOutput) ToContainerNamespaceArrayOutput() ContainerNamespaceArrayOutput {
	return o
}

func (o ContainerNamespaceArrayOutput) ToContainerNamespaceArrayOutputWithContext(ctx context.Context) ContainerNamespaceArrayOutput {
	return o
}

func (o ContainerNamespaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ContainerNamespace] {
	return pulumix.Output[[]*ContainerNamespace]{
		OutputState: o.OutputState,
	}
}

func (o ContainerNamespaceArrayOutput) Index(i pulumi.IntInput) ContainerNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerNamespace {
		return vs[0].([]*ContainerNamespace)[vs[1].(int)]
	}).(ContainerNamespaceOutput)
}

type ContainerNamespaceMapOutput struct{ *pulumi.OutputState }

func (ContainerNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerNamespace)(nil)).Elem()
}

func (o ContainerNamespaceMapOutput) ToContainerNamespaceMapOutput() ContainerNamespaceMapOutput {
	return o
}

func (o ContainerNamespaceMapOutput) ToContainerNamespaceMapOutputWithContext(ctx context.Context) ContainerNamespaceMapOutput {
	return o
}

func (o ContainerNamespaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContainerNamespace] {
	return pulumix.Output[map[string]*ContainerNamespace]{
		OutputState: o.OutputState,
	}
}

func (o ContainerNamespaceMapOutput) MapIndex(k pulumi.StringInput) ContainerNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerNamespace {
		return vs[0].(map[string]*ContainerNamespace)[vs[1].(string)]
	}).(ContainerNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerNamespaceInput)(nil)).Elem(), &ContainerNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerNamespaceArrayInput)(nil)).Elem(), ContainerNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerNamespaceMapInput)(nil)).Elem(), ContainerNamespaceMap{})
	pulumi.RegisterOutputType(ContainerNamespaceOutput{})
	pulumi.RegisterOutputType(ContainerNamespaceArrayOutput{})
	pulumi.RegisterOutputType(ContainerNamespaceMapOutput{})
}
