// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Add members to an IAM group.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#groups-f592eb).
//
// ## Example Usage
//
// ### Application Membership
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
//			group, err := scaleway.NewIamGroup(ctx, "group", &scaleway.IamGroupArgs{
//				ExternalMembership: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			app, err := scaleway.NewIamApplication(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewIamGroupMembership(ctx, "member", &scaleway.IamGroupMembershipArgs{
//				GroupId:       group.ID(),
//				ApplicationId: app.ID(),
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
// IAM group memberships can be imported using two format:
//
// - For user: `{group_id}/user/{user_id}`
//
// - For application: `{group_id}/app/{application_id}`
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/iamGroupMembership:IamGroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
// ```
type IamGroupMembership struct {
	pulumi.CustomResourceState

	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewIamGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewIamGroupMembership(ctx *pulumi.Context,
	name string, args *IamGroupMembershipArgs, opts ...pulumi.ResourceOption) (*IamGroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamGroupMembership
	err := ctx.RegisterResource("scaleway:index/iamGroupMembership:IamGroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamGroupMembership gets an existing IamGroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamGroupMembershipState, opts ...pulumi.ResourceOption) (*IamGroupMembership, error) {
	var resource IamGroupMembership
	err := ctx.ReadResource("scaleway:index/iamGroupMembership:IamGroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamGroupMembership resources.
type iamGroupMembershipState struct {
	// The ID of the application that will be added to the group.
	ApplicationId *string `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId *string `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId *string `pulumi:"userId"`
}

type IamGroupMembershipState struct {
	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrInput
	// ID of the group to add members to.
	GroupId pulumi.StringPtrInput
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrInput
}

func (IamGroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupMembershipState)(nil)).Elem()
}

type iamGroupMembershipArgs struct {
	// The ID of the application that will be added to the group.
	ApplicationId *string `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId string `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a IamGroupMembership resource.
type IamGroupMembershipArgs struct {
	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrInput
	// ID of the group to add members to.
	GroupId pulumi.StringInput
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrInput
}

func (IamGroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupMembershipArgs)(nil)).Elem()
}

type IamGroupMembershipInput interface {
	pulumi.Input

	ToIamGroupMembershipOutput() IamGroupMembershipOutput
	ToIamGroupMembershipOutputWithContext(ctx context.Context) IamGroupMembershipOutput
}

func (*IamGroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupMembership)(nil)).Elem()
}

func (i *IamGroupMembership) ToIamGroupMembershipOutput() IamGroupMembershipOutput {
	return i.ToIamGroupMembershipOutputWithContext(context.Background())
}

func (i *IamGroupMembership) ToIamGroupMembershipOutputWithContext(ctx context.Context) IamGroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupMembershipOutput)
}

// IamGroupMembershipArrayInput is an input type that accepts IamGroupMembershipArray and IamGroupMembershipArrayOutput values.
// You can construct a concrete instance of `IamGroupMembershipArrayInput` via:
//
//	IamGroupMembershipArray{ IamGroupMembershipArgs{...} }
type IamGroupMembershipArrayInput interface {
	pulumi.Input

	ToIamGroupMembershipArrayOutput() IamGroupMembershipArrayOutput
	ToIamGroupMembershipArrayOutputWithContext(context.Context) IamGroupMembershipArrayOutput
}

type IamGroupMembershipArray []IamGroupMembershipInput

func (IamGroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroupMembership)(nil)).Elem()
}

func (i IamGroupMembershipArray) ToIamGroupMembershipArrayOutput() IamGroupMembershipArrayOutput {
	return i.ToIamGroupMembershipArrayOutputWithContext(context.Background())
}

func (i IamGroupMembershipArray) ToIamGroupMembershipArrayOutputWithContext(ctx context.Context) IamGroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupMembershipArrayOutput)
}

// IamGroupMembershipMapInput is an input type that accepts IamGroupMembershipMap and IamGroupMembershipMapOutput values.
// You can construct a concrete instance of `IamGroupMembershipMapInput` via:
//
//	IamGroupMembershipMap{ "key": IamGroupMembershipArgs{...} }
type IamGroupMembershipMapInput interface {
	pulumi.Input

	ToIamGroupMembershipMapOutput() IamGroupMembershipMapOutput
	ToIamGroupMembershipMapOutputWithContext(context.Context) IamGroupMembershipMapOutput
}

type IamGroupMembershipMap map[string]IamGroupMembershipInput

func (IamGroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroupMembership)(nil)).Elem()
}

func (i IamGroupMembershipMap) ToIamGroupMembershipMapOutput() IamGroupMembershipMapOutput {
	return i.ToIamGroupMembershipMapOutputWithContext(context.Background())
}

func (i IamGroupMembershipMap) ToIamGroupMembershipMapOutputWithContext(ctx context.Context) IamGroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupMembershipMapOutput)
}

type IamGroupMembershipOutput struct{ *pulumi.OutputState }

func (IamGroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupMembership)(nil)).Elem()
}

func (o IamGroupMembershipOutput) ToIamGroupMembershipOutput() IamGroupMembershipOutput {
	return o
}

func (o IamGroupMembershipOutput) ToIamGroupMembershipOutputWithContext(ctx context.Context) IamGroupMembershipOutput {
	return o
}

// The ID of the application that will be added to the group.
func (o IamGroupMembershipOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamGroupMembership) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// ID of the group to add members to.
func (o IamGroupMembershipOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroupMembership) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The ID of the user that will be added to the group
//
// - > Only one of `applicationId` or `userId` must be specified
func (o IamGroupMembershipOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamGroupMembership) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type IamGroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (IamGroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroupMembership)(nil)).Elem()
}

func (o IamGroupMembershipArrayOutput) ToIamGroupMembershipArrayOutput() IamGroupMembershipArrayOutput {
	return o
}

func (o IamGroupMembershipArrayOutput) ToIamGroupMembershipArrayOutputWithContext(ctx context.Context) IamGroupMembershipArrayOutput {
	return o
}

func (o IamGroupMembershipArrayOutput) Index(i pulumi.IntInput) IamGroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamGroupMembership {
		return vs[0].([]*IamGroupMembership)[vs[1].(int)]
	}).(IamGroupMembershipOutput)
}

type IamGroupMembershipMapOutput struct{ *pulumi.OutputState }

func (IamGroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroupMembership)(nil)).Elem()
}

func (o IamGroupMembershipMapOutput) ToIamGroupMembershipMapOutput() IamGroupMembershipMapOutput {
	return o
}

func (o IamGroupMembershipMapOutput) ToIamGroupMembershipMapOutputWithContext(ctx context.Context) IamGroupMembershipMapOutput {
	return o
}

func (o IamGroupMembershipMapOutput) MapIndex(k pulumi.StringInput) IamGroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamGroupMembership {
		return vs[0].(map[string]*IamGroupMembership)[vs[1].(string)]
	}).(IamGroupMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupMembershipInput)(nil)).Elem(), &IamGroupMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupMembershipArrayInput)(nil)).Elem(), IamGroupMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupMembershipMapInput)(nil)).Elem(), IamGroupMembershipMap{})
	pulumi.RegisterOutputType(IamGroupMembershipOutput{})
	pulumi.RegisterOutputType(IamGroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(IamGroupMembershipMapOutput{})
}
