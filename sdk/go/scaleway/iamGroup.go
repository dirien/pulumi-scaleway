// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway IAM Groups.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
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
//			_, err := scaleway.NewIamGroup(ctx, "basic", &scaleway.IamGroupArgs{
//				ApplicationIds: pulumi.StringArray{},
//				Description:    pulumi.String("basic description"),
//				UserIds:        pulumi.StringArray{},
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
// ### With applications
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
//			app, err := scaleway.NewIamApplication(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewIamGroup(ctx, "withApp", &scaleway.IamGroupArgs{
//				ApplicationIds: pulumi.StringArray{
//					app.ID(),
//				},
//				UserIds: pulumi.StringArray{},
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
// IAM groups can be imported using the `{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/iamGroup:IamGroup basic 11111111-1111-1111-1111-111111111111
//
// ```
type IamGroup struct {
	pulumi.CustomResourceState

	// The list of IDs of the applications attached to the group.
	ApplicationIds pulumi.StringArrayOutput `pulumi:"applicationIds"`
	// The date and time of the creation of the group
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the IAM group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the IAM group.
	Name pulumi.StringOutput `pulumi:"name"`
	// `organizationId`) The ID of the organization the group is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The date and time of the last update of the group
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The list of IDs of the users attached to the group.
	UserIds pulumi.StringArrayOutput `pulumi:"userIds"`
}

// NewIamGroup registers a new resource with the given unique name, arguments, and options.
func NewIamGroup(ctx *pulumi.Context,
	name string, args *IamGroupArgs, opts ...pulumi.ResourceOption) (*IamGroup, error) {
	if args == nil {
		args = &IamGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IamGroup
	err := ctx.RegisterResource("scaleway:index/iamGroup:IamGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamGroup gets an existing IamGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamGroupState, opts ...pulumi.ResourceOption) (*IamGroup, error) {
	var resource IamGroup
	err := ctx.ReadResource("scaleway:index/iamGroup:IamGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamGroup resources.
type iamGroupState struct {
	// The list of IDs of the applications attached to the group.
	ApplicationIds []string `pulumi:"applicationIds"`
	// The date and time of the creation of the group
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the IAM group.
	Description *string `pulumi:"description"`
	// The name of the IAM group.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the group is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The date and time of the last update of the group
	UpdatedAt *string `pulumi:"updatedAt"`
	// The list of IDs of the users attached to the group.
	UserIds []string `pulumi:"userIds"`
}

type IamGroupState struct {
	// The list of IDs of the applications attached to the group.
	ApplicationIds pulumi.StringArrayInput
	// The date and time of the creation of the group
	CreatedAt pulumi.StringPtrInput
	// The description of the IAM group.
	Description pulumi.StringPtrInput
	// The name of the IAM group.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the group is associated with.
	OrganizationId pulumi.StringPtrInput
	// The date and time of the last update of the group
	UpdatedAt pulumi.StringPtrInput
	// The list of IDs of the users attached to the group.
	UserIds pulumi.StringArrayInput
}

func (IamGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupState)(nil)).Elem()
}

type iamGroupArgs struct {
	// The list of IDs of the applications attached to the group.
	ApplicationIds []string `pulumi:"applicationIds"`
	// The description of the IAM group.
	Description *string `pulumi:"description"`
	// The name of the IAM group.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the organization the group is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The list of IDs of the users attached to the group.
	UserIds []string `pulumi:"userIds"`
}

// The set of arguments for constructing a IamGroup resource.
type IamGroupArgs struct {
	// The list of IDs of the applications attached to the group.
	ApplicationIds pulumi.StringArrayInput
	// The description of the IAM group.
	Description pulumi.StringPtrInput
	// The name of the IAM group.
	Name pulumi.StringPtrInput
	// `organizationId`) The ID of the organization the group is associated with.
	OrganizationId pulumi.StringPtrInput
	// The list of IDs of the users attached to the group.
	UserIds pulumi.StringArrayInput
}

func (IamGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupArgs)(nil)).Elem()
}

type IamGroupInput interface {
	pulumi.Input

	ToIamGroupOutput() IamGroupOutput
	ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput
}

func (*IamGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroup)(nil)).Elem()
}

func (i *IamGroup) ToIamGroupOutput() IamGroupOutput {
	return i.ToIamGroupOutputWithContext(context.Background())
}

func (i *IamGroup) ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupOutput)
}

// IamGroupArrayInput is an input type that accepts IamGroupArray and IamGroupArrayOutput values.
// You can construct a concrete instance of `IamGroupArrayInput` via:
//
//	IamGroupArray{ IamGroupArgs{...} }
type IamGroupArrayInput interface {
	pulumi.Input

	ToIamGroupArrayOutput() IamGroupArrayOutput
	ToIamGroupArrayOutputWithContext(context.Context) IamGroupArrayOutput
}

type IamGroupArray []IamGroupInput

func (IamGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroup)(nil)).Elem()
}

func (i IamGroupArray) ToIamGroupArrayOutput() IamGroupArrayOutput {
	return i.ToIamGroupArrayOutputWithContext(context.Background())
}

func (i IamGroupArray) ToIamGroupArrayOutputWithContext(ctx context.Context) IamGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupArrayOutput)
}

// IamGroupMapInput is an input type that accepts IamGroupMap and IamGroupMapOutput values.
// You can construct a concrete instance of `IamGroupMapInput` via:
//
//	IamGroupMap{ "key": IamGroupArgs{...} }
type IamGroupMapInput interface {
	pulumi.Input

	ToIamGroupMapOutput() IamGroupMapOutput
	ToIamGroupMapOutputWithContext(context.Context) IamGroupMapOutput
}

type IamGroupMap map[string]IamGroupInput

func (IamGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroup)(nil)).Elem()
}

func (i IamGroupMap) ToIamGroupMapOutput() IamGroupMapOutput {
	return i.ToIamGroupMapOutputWithContext(context.Background())
}

func (i IamGroupMap) ToIamGroupMapOutputWithContext(ctx context.Context) IamGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupMapOutput)
}

type IamGroupOutput struct{ *pulumi.OutputState }

func (IamGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroup)(nil)).Elem()
}

func (o IamGroupOutput) ToIamGroupOutput() IamGroupOutput {
	return o
}

func (o IamGroupOutput) ToIamGroupOutputWithContext(ctx context.Context) IamGroupOutput {
	return o
}

// The list of IDs of the applications attached to the group.
func (o IamGroupOutput) ApplicationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringArrayOutput { return v.ApplicationIds }).(pulumi.StringArrayOutput)
}

// The date and time of the creation of the group
func (o IamGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the IAM group.
func (o IamGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the IAM group.
func (o IamGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `organizationId`) The ID of the organization the group is associated with.
func (o IamGroupOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The date and time of the last update of the group
func (o IamGroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The list of IDs of the users attached to the group.
func (o IamGroupOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IamGroup) pulumi.StringArrayOutput { return v.UserIds }).(pulumi.StringArrayOutput)
}

type IamGroupArrayOutput struct{ *pulumi.OutputState }

func (IamGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroup)(nil)).Elem()
}

func (o IamGroupArrayOutput) ToIamGroupArrayOutput() IamGroupArrayOutput {
	return o
}

func (o IamGroupArrayOutput) ToIamGroupArrayOutputWithContext(ctx context.Context) IamGroupArrayOutput {
	return o
}

func (o IamGroupArrayOutput) Index(i pulumi.IntInput) IamGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamGroup {
		return vs[0].([]*IamGroup)[vs[1].(int)]
	}).(IamGroupOutput)
}

type IamGroupMapOutput struct{ *pulumi.OutputState }

func (IamGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroup)(nil)).Elem()
}

func (o IamGroupMapOutput) ToIamGroupMapOutput() IamGroupMapOutput {
	return o
}

func (o IamGroupMapOutput) ToIamGroupMapOutputWithContext(ctx context.Context) IamGroupMapOutput {
	return o
}

func (o IamGroupMapOutput) MapIndex(k pulumi.StringInput) IamGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamGroup {
		return vs[0].(map[string]*IamGroup)[vs[1].(string)]
	}).(IamGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupInput)(nil)).Elem(), &IamGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupArrayInput)(nil)).Elem(), IamGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupMapInput)(nil)).Elem(), IamGroupMap{})
	pulumi.RegisterOutputType(IamGroupOutput{})
	pulumi.RegisterOutputType(IamGroupArrayOutput{})
	pulumi.RegisterOutputType(IamGroupMapOutput{})
}
