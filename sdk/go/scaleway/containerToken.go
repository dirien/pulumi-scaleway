// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ContainerToken` resource allows you to create and manage authentication tokens for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
//
// Refer to the Containers tokens [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/create-auth-token-from-console/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-tokens-list-all-tokens) for more information.
//
// ## Example Usage
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
//			mainContainerNamespace, err := scaleway.NewContainerNamespace(ctx, "mainContainerNamespace", nil)
//			if err != nil {
//				return err
//			}
//			mainContainer, err := scaleway.NewContainer(ctx, "mainContainer", &scaleway.ContainerArgs{
//				NamespaceId: mainContainerNamespace.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			// Namespace Token
//			_, err = scaleway.NewContainerToken(ctx, "namespace", &scaleway.ContainerTokenArgs{
//				NamespaceId: mainContainerNamespace.ID(),
//				ExpiresAt:   pulumi.String("2022-10-18T11:35:15+02:00"),
//			})
//			if err != nil {
//				return err
//			}
//			// Container Token
//			_, err = scaleway.NewContainerToken(ctx, "container", &scaleway.ContainerTokenArgs{
//				ContainerId: mainContainer.ID(),
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
// Tokens can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/containerToken:ContainerToken main fr-par/11111111-1111-1111-1111-111111111111
// ```
type ContainerToken struct {
	pulumi.CustomResourceState

	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrOutput `pulumi:"containerId"`
	// The description of the token.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrOutput `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringOutput `pulumi:"region"`
	// The token.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewContainerToken registers a new resource with the given unique name, arguments, and options.
func NewContainerToken(ctx *pulumi.Context,
	name string, args *ContainerTokenArgs, opts ...pulumi.ResourceOption) (*ContainerToken, error) {
	if args == nil {
		args = &ContainerTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerToken
	err := ctx.RegisterResource("scaleway:index/containerToken:ContainerToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerToken gets an existing ContainerToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerTokenState, opts ...pulumi.ResourceOption) (*ContainerToken, error) {
	var resource ContainerToken
	err := ctx.ReadResource("scaleway:index/containerToken:ContainerToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerToken resources.
type containerTokenState struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId *string `pulumi:"containerId"`
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region *string `pulumi:"region"`
	// The token.
	Token *string `pulumi:"token"`
}

type ContainerTokenState struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrInput
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringPtrInput
	// The token.
	Token pulumi.StringPtrInput
}

func (ContainerTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerTokenState)(nil)).Elem()
}

type containerTokenArgs struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId *string `pulumi:"containerId"`
	// The description of the token.
	Description *string `pulumi:"description"`
	// The expiration date of the token.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The unique identifier of the Containers namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ContainerToken resource.
type ContainerTokenArgs struct {
	// The unique identifier of the container.
	//
	// > Only one of `namespaceId` or `containerId` must be set.
	ContainerId pulumi.StringPtrInput
	// The description of the token.
	Description pulumi.StringPtrInput
	// The expiration date of the token.
	ExpiresAt pulumi.StringPtrInput
	// The unique identifier of the Containers namespace.
	NamespaceId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	//
	// > **Important** Updating any of the arguments above will recreate the token.
	Region pulumi.StringPtrInput
}

func (ContainerTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerTokenArgs)(nil)).Elem()
}

type ContainerTokenInput interface {
	pulumi.Input

	ToContainerTokenOutput() ContainerTokenOutput
	ToContainerTokenOutputWithContext(ctx context.Context) ContainerTokenOutput
}

func (*ContainerToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerToken)(nil)).Elem()
}

func (i *ContainerToken) ToContainerTokenOutput() ContainerTokenOutput {
	return i.ToContainerTokenOutputWithContext(context.Background())
}

func (i *ContainerToken) ToContainerTokenOutputWithContext(ctx context.Context) ContainerTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTokenOutput)
}

// ContainerTokenArrayInput is an input type that accepts ContainerTokenArray and ContainerTokenArrayOutput values.
// You can construct a concrete instance of `ContainerTokenArrayInput` via:
//
//	ContainerTokenArray{ ContainerTokenArgs{...} }
type ContainerTokenArrayInput interface {
	pulumi.Input

	ToContainerTokenArrayOutput() ContainerTokenArrayOutput
	ToContainerTokenArrayOutputWithContext(context.Context) ContainerTokenArrayOutput
}

type ContainerTokenArray []ContainerTokenInput

func (ContainerTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerToken)(nil)).Elem()
}

func (i ContainerTokenArray) ToContainerTokenArrayOutput() ContainerTokenArrayOutput {
	return i.ToContainerTokenArrayOutputWithContext(context.Background())
}

func (i ContainerTokenArray) ToContainerTokenArrayOutputWithContext(ctx context.Context) ContainerTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTokenArrayOutput)
}

// ContainerTokenMapInput is an input type that accepts ContainerTokenMap and ContainerTokenMapOutput values.
// You can construct a concrete instance of `ContainerTokenMapInput` via:
//
//	ContainerTokenMap{ "key": ContainerTokenArgs{...} }
type ContainerTokenMapInput interface {
	pulumi.Input

	ToContainerTokenMapOutput() ContainerTokenMapOutput
	ToContainerTokenMapOutputWithContext(context.Context) ContainerTokenMapOutput
}

type ContainerTokenMap map[string]ContainerTokenInput

func (ContainerTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerToken)(nil)).Elem()
}

func (i ContainerTokenMap) ToContainerTokenMapOutput() ContainerTokenMapOutput {
	return i.ToContainerTokenMapOutputWithContext(context.Background())
}

func (i ContainerTokenMap) ToContainerTokenMapOutputWithContext(ctx context.Context) ContainerTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerTokenMapOutput)
}

type ContainerTokenOutput struct{ *pulumi.OutputState }

func (ContainerTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerToken)(nil)).Elem()
}

func (o ContainerTokenOutput) ToContainerTokenOutput() ContainerTokenOutput {
	return o
}

func (o ContainerTokenOutput) ToContainerTokenOutputWithContext(ctx context.Context) ContainerTokenOutput {
	return o
}

// The unique identifier of the container.
//
// > Only one of `namespaceId` or `containerId` must be set.
func (o ContainerTokenOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringPtrOutput { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The description of the token.
func (o ContainerTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The expiration date of the token.
func (o ContainerTokenOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The unique identifier of the Containers namespace.
func (o ContainerTokenOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringPtrOutput { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

// `region`). The region in which the namespace is created.
//
// > **Important** Updating any of the arguments above will recreate the token.
func (o ContainerTokenOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The token.
func (o ContainerTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type ContainerTokenArrayOutput struct{ *pulumi.OutputState }

func (ContainerTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerToken)(nil)).Elem()
}

func (o ContainerTokenArrayOutput) ToContainerTokenArrayOutput() ContainerTokenArrayOutput {
	return o
}

func (o ContainerTokenArrayOutput) ToContainerTokenArrayOutputWithContext(ctx context.Context) ContainerTokenArrayOutput {
	return o
}

func (o ContainerTokenArrayOutput) Index(i pulumi.IntInput) ContainerTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerToken {
		return vs[0].([]*ContainerToken)[vs[1].(int)]
	}).(ContainerTokenOutput)
}

type ContainerTokenMapOutput struct{ *pulumi.OutputState }

func (ContainerTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerToken)(nil)).Elem()
}

func (o ContainerTokenMapOutput) ToContainerTokenMapOutput() ContainerTokenMapOutput {
	return o
}

func (o ContainerTokenMapOutput) ToContainerTokenMapOutputWithContext(ctx context.Context) ContainerTokenMapOutput {
	return o
}

func (o ContainerTokenMapOutput) MapIndex(k pulumi.StringInput) ContainerTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerToken {
		return vs[0].(map[string]*ContainerToken)[vs[1].(string)]
	}).(ContainerTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTokenInput)(nil)).Elem(), &ContainerToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTokenArrayInput)(nil)).Elem(), ContainerTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTokenMapInput)(nil)).Elem(), ContainerTokenMap{})
	pulumi.RegisterOutputType(ContainerTokenOutput{})
	pulumi.RegisterOutputType(ContainerTokenArrayOutput{})
	pulumi.RegisterOutputType(ContainerTokenMapOutput{})
}
