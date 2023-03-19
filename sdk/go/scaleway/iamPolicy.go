// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway IAM Policies. For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#policies-54b8a7).
//
// ## Example Usage
// ### Create a policy for an organization's project
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
//			_default, err := scaleway.LookupAccountProject(ctx, &scaleway.LookupAccountProjectArgs{
//				Name: pulumi.StringRef("default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			app, err := scaleway.NewIamApplication(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewIamPolicy(ctx, "objectReadOnly", &scaleway.IamPolicyArgs{
//				Description:   pulumi.String("gives app readonly access to object storage in project"),
//				ApplicationId: app.ID(),
//				Rules: scaleway.IamPolicyRuleArray{
//					&scaleway.IamPolicyRuleArgs{
//						ProjectIds: pulumi.StringArray{
//							*pulumi.String(_default.Id),
//						},
//						PermissionSetNames: pulumi.StringArray{
//							pulumi.String("ObjectStorageReadOnly"),
//						},
//					},
//				},
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
// Policies can be imported using the `{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/iamPolicy:IamPolicy main 11111111-1111-1111-1111-111111111111
//
// ```
type IamPolicy struct {
	pulumi.CustomResourceState

	// ID of the Application the policy will be linked to
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// The date and time of the creation of the policy.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the iam policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the policy is editable.
	Editable pulumi.BoolOutput `pulumi:"editable"`
	// ID of the Group the policy will be linked to
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// .The name of the iam policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// If the policy doesn't apply to a principal.
	NoPrincipal pulumi.BoolPtrOutput `pulumi:"noPrincipal"`
	// ID of organization scoped to the rule.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of rules in the policy.
	Rules IamPolicyRuleArrayOutput `pulumi:"rules"`
	// The date and time of the last update of the policy.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// ID of the User the policy will be linked to
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewIamPolicy(ctx *pulumi.Context,
	name string, args *IamPolicyArgs, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IamPolicy
	err := ctx.RegisterResource("scaleway:index/iamPolicy:IamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamPolicy gets an existing IamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamPolicyState, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	var resource IamPolicy
	err := ctx.ReadResource("scaleway:index/iamPolicy:IamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamPolicy resources.
type iamPolicyState struct {
	// ID of the Application the policy will be linked to
	ApplicationId *string `pulumi:"applicationId"`
	// The date and time of the creation of the policy.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the iam policy.
	Description *string `pulumi:"description"`
	// Whether the policy is editable.
	Editable *bool `pulumi:"editable"`
	// ID of the Group the policy will be linked to
	GroupId *string `pulumi:"groupId"`
	// .The name of the iam policy.
	Name *string `pulumi:"name"`
	// If the policy doesn't apply to a principal.
	NoPrincipal *bool `pulumi:"noPrincipal"`
	// ID of organization scoped to the rule.
	OrganizationId *string `pulumi:"organizationId"`
	// List of rules in the policy.
	Rules []IamPolicyRule `pulumi:"rules"`
	// The date and time of the last update of the policy.
	UpdatedAt *string `pulumi:"updatedAt"`
	// ID of the User the policy will be linked to
	UserId *string `pulumi:"userId"`
}

type IamPolicyState struct {
	// ID of the Application the policy will be linked to
	ApplicationId pulumi.StringPtrInput
	// The date and time of the creation of the policy.
	CreatedAt pulumi.StringPtrInput
	// The description of the iam policy.
	Description pulumi.StringPtrInput
	// Whether the policy is editable.
	Editable pulumi.BoolPtrInput
	// ID of the Group the policy will be linked to
	GroupId pulumi.StringPtrInput
	// .The name of the iam policy.
	Name pulumi.StringPtrInput
	// If the policy doesn't apply to a principal.
	NoPrincipal pulumi.BoolPtrInput
	// ID of organization scoped to the rule.
	OrganizationId pulumi.StringPtrInput
	// List of rules in the policy.
	Rules IamPolicyRuleArrayInput
	// The date and time of the last update of the policy.
	UpdatedAt pulumi.StringPtrInput
	// ID of the User the policy will be linked to
	UserId pulumi.StringPtrInput
}

func (IamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyState)(nil)).Elem()
}

type iamPolicyArgs struct {
	// ID of the Application the policy will be linked to
	ApplicationId *string `pulumi:"applicationId"`
	// The description of the iam policy.
	Description *string `pulumi:"description"`
	// ID of the Group the policy will be linked to
	GroupId *string `pulumi:"groupId"`
	// .The name of the iam policy.
	Name *string `pulumi:"name"`
	// If the policy doesn't apply to a principal.
	NoPrincipal *bool `pulumi:"noPrincipal"`
	// ID of organization scoped to the rule.
	OrganizationId *string `pulumi:"organizationId"`
	// List of rules in the policy.
	Rules []IamPolicyRule `pulumi:"rules"`
	// ID of the User the policy will be linked to
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a IamPolicy resource.
type IamPolicyArgs struct {
	// ID of the Application the policy will be linked to
	ApplicationId pulumi.StringPtrInput
	// The description of the iam policy.
	Description pulumi.StringPtrInput
	// ID of the Group the policy will be linked to
	GroupId pulumi.StringPtrInput
	// .The name of the iam policy.
	Name pulumi.StringPtrInput
	// If the policy doesn't apply to a principal.
	NoPrincipal pulumi.BoolPtrInput
	// ID of organization scoped to the rule.
	OrganizationId pulumi.StringPtrInput
	// List of rules in the policy.
	Rules IamPolicyRuleArrayInput
	// ID of the User the policy will be linked to
	UserId pulumi.StringPtrInput
}

func (IamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyArgs)(nil)).Elem()
}

type IamPolicyInput interface {
	pulumi.Input

	ToIamPolicyOutput() IamPolicyOutput
	ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput
}

func (*IamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicy)(nil)).Elem()
}

func (i *IamPolicy) ToIamPolicyOutput() IamPolicyOutput {
	return i.ToIamPolicyOutputWithContext(context.Background())
}

func (i *IamPolicy) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyOutput)
}

// IamPolicyArrayInput is an input type that accepts IamPolicyArray and IamPolicyArrayOutput values.
// You can construct a concrete instance of `IamPolicyArrayInput` via:
//
//	IamPolicyArray{ IamPolicyArgs{...} }
type IamPolicyArrayInput interface {
	pulumi.Input

	ToIamPolicyArrayOutput() IamPolicyArrayOutput
	ToIamPolicyArrayOutputWithContext(context.Context) IamPolicyArrayOutput
}

type IamPolicyArray []IamPolicyInput

func (IamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicy)(nil)).Elem()
}

func (i IamPolicyArray) ToIamPolicyArrayOutput() IamPolicyArrayOutput {
	return i.ToIamPolicyArrayOutputWithContext(context.Background())
}

func (i IamPolicyArray) ToIamPolicyArrayOutputWithContext(ctx context.Context) IamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyArrayOutput)
}

// IamPolicyMapInput is an input type that accepts IamPolicyMap and IamPolicyMapOutput values.
// You can construct a concrete instance of `IamPolicyMapInput` via:
//
//	IamPolicyMap{ "key": IamPolicyArgs{...} }
type IamPolicyMapInput interface {
	pulumi.Input

	ToIamPolicyMapOutput() IamPolicyMapOutput
	ToIamPolicyMapOutputWithContext(context.Context) IamPolicyMapOutput
}

type IamPolicyMap map[string]IamPolicyInput

func (IamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicy)(nil)).Elem()
}

func (i IamPolicyMap) ToIamPolicyMapOutput() IamPolicyMapOutput {
	return i.ToIamPolicyMapOutputWithContext(context.Background())
}

func (i IamPolicyMap) ToIamPolicyMapOutputWithContext(ctx context.Context) IamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyMapOutput)
}

type IamPolicyOutput struct{ *pulumi.OutputState }

func (IamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicy)(nil)).Elem()
}

func (o IamPolicyOutput) ToIamPolicyOutput() IamPolicyOutput {
	return o
}

func (o IamPolicyOutput) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return o
}

// ID of the Application the policy will be linked to
func (o IamPolicyOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// The date and time of the creation of the policy.
func (o IamPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the iam policy.
func (o IamPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the policy is editable.
func (o IamPolicyOutput) Editable() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.BoolOutput { return v.Editable }).(pulumi.BoolOutput)
}

// ID of the Group the policy will be linked to
func (o IamPolicyOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// .The name of the iam policy.
func (o IamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If the policy doesn't apply to a principal.
func (o IamPolicyOutput) NoPrincipal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.BoolPtrOutput { return v.NoPrincipal }).(pulumi.BoolPtrOutput)
}

// ID of organization scoped to the rule.
func (o IamPolicyOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// List of rules in the policy.
func (o IamPolicyOutput) Rules() IamPolicyRuleArrayOutput {
	return o.ApplyT(func(v *IamPolicy) IamPolicyRuleArrayOutput { return v.Rules }).(IamPolicyRuleArrayOutput)
}

// The date and time of the last update of the policy.
func (o IamPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// ID of the User the policy will be linked to
func (o IamPolicyOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type IamPolicyArrayOutput struct{ *pulumi.OutputState }

func (IamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicy)(nil)).Elem()
}

func (o IamPolicyArrayOutput) ToIamPolicyArrayOutput() IamPolicyArrayOutput {
	return o
}

func (o IamPolicyArrayOutput) ToIamPolicyArrayOutputWithContext(ctx context.Context) IamPolicyArrayOutput {
	return o
}

func (o IamPolicyArrayOutput) Index(i pulumi.IntInput) IamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamPolicy {
		return vs[0].([]*IamPolicy)[vs[1].(int)]
	}).(IamPolicyOutput)
}

type IamPolicyMapOutput struct{ *pulumi.OutputState }

func (IamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicy)(nil)).Elem()
}

func (o IamPolicyMapOutput) ToIamPolicyMapOutput() IamPolicyMapOutput {
	return o
}

func (o IamPolicyMapOutput) ToIamPolicyMapOutputWithContext(ctx context.Context) IamPolicyMapOutput {
	return o
}

func (o IamPolicyMapOutput) MapIndex(k pulumi.StringInput) IamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamPolicy {
		return vs[0].(map[string]*IamPolicy)[vs[1].(string)]
	}).(IamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyInput)(nil)).Elem(), &IamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyArrayInput)(nil)).Elem(), IamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyMapInput)(nil)).Elem(), IamPolicyMap{})
	pulumi.RegisterOutputType(IamPolicyOutput{})
	pulumi.RegisterOutputType(IamPolicyArrayOutput{})
	pulumi.RegisterOutputType(IamPolicyMapOutput{})
}
