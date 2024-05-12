// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Projects can be imported using the `id`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/accountProject:AccountProject project 11111111-1111-1111-1111-111111111111
// ```
type AccountProject struct {
	pulumi.CustomResourceState

	// The Project creation time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the Project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Project.
	Name pulumi.StringOutput `pulumi:"name"`
	// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The Project last update time.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewAccountProject registers a new resource with the given unique name, arguments, and options.
func NewAccountProject(ctx *pulumi.Context,
	name string, args *AccountProjectArgs, opts ...pulumi.ResourceOption) (*AccountProject, error) {
	if args == nil {
		args = &AccountProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountProject
	err := ctx.RegisterResource("scaleway:index/accountProject:AccountProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountProject gets an existing AccountProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountProjectState, opts ...pulumi.ResourceOption) (*AccountProject, error) {
	var resource AccountProject
	err := ctx.ReadResource("scaleway:index/accountProject:AccountProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountProject resources.
type accountProjectState struct {
	// The Project creation time.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the Project.
	Description *string `pulumi:"description"`
	// The name of the Project.
	Name *string `pulumi:"name"`
	// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
	OrganizationId *string `pulumi:"organizationId"`
	// The Project last update time.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type AccountProjectState struct {
	// The Project creation time.
	CreatedAt pulumi.StringPtrInput
	// The description of the Project.
	Description pulumi.StringPtrInput
	// The name of the Project.
	Name pulumi.StringPtrInput
	// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
	OrganizationId pulumi.StringPtrInput
	// The Project last update time.
	UpdatedAt pulumi.StringPtrInput
}

func (AccountProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountProjectState)(nil)).Elem()
}

type accountProjectArgs struct {
	// The description of the Project.
	Description *string `pulumi:"description"`
	// The name of the Project.
	Name *string `pulumi:"name"`
	// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
	OrganizationId *string `pulumi:"organizationId"`
}

// The set of arguments for constructing a AccountProject resource.
type AccountProjectArgs struct {
	// The description of the Project.
	Description pulumi.StringPtrInput
	// The name of the Project.
	Name pulumi.StringPtrInput
	// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
	OrganizationId pulumi.StringPtrInput
}

func (AccountProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountProjectArgs)(nil)).Elem()
}

type AccountProjectInput interface {
	pulumi.Input

	ToAccountProjectOutput() AccountProjectOutput
	ToAccountProjectOutputWithContext(ctx context.Context) AccountProjectOutput
}

func (*AccountProject) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountProject)(nil)).Elem()
}

func (i *AccountProject) ToAccountProjectOutput() AccountProjectOutput {
	return i.ToAccountProjectOutputWithContext(context.Background())
}

func (i *AccountProject) ToAccountProjectOutputWithContext(ctx context.Context) AccountProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountProjectOutput)
}

// AccountProjectArrayInput is an input type that accepts AccountProjectArray and AccountProjectArrayOutput values.
// You can construct a concrete instance of `AccountProjectArrayInput` via:
//
//	AccountProjectArray{ AccountProjectArgs{...} }
type AccountProjectArrayInput interface {
	pulumi.Input

	ToAccountProjectArrayOutput() AccountProjectArrayOutput
	ToAccountProjectArrayOutputWithContext(context.Context) AccountProjectArrayOutput
}

type AccountProjectArray []AccountProjectInput

func (AccountProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountProject)(nil)).Elem()
}

func (i AccountProjectArray) ToAccountProjectArrayOutput() AccountProjectArrayOutput {
	return i.ToAccountProjectArrayOutputWithContext(context.Background())
}

func (i AccountProjectArray) ToAccountProjectArrayOutputWithContext(ctx context.Context) AccountProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountProjectArrayOutput)
}

// AccountProjectMapInput is an input type that accepts AccountProjectMap and AccountProjectMapOutput values.
// You can construct a concrete instance of `AccountProjectMapInput` via:
//
//	AccountProjectMap{ "key": AccountProjectArgs{...} }
type AccountProjectMapInput interface {
	pulumi.Input

	ToAccountProjectMapOutput() AccountProjectMapOutput
	ToAccountProjectMapOutputWithContext(context.Context) AccountProjectMapOutput
}

type AccountProjectMap map[string]AccountProjectInput

func (AccountProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountProject)(nil)).Elem()
}

func (i AccountProjectMap) ToAccountProjectMapOutput() AccountProjectMapOutput {
	return i.ToAccountProjectMapOutputWithContext(context.Background())
}

func (i AccountProjectMap) ToAccountProjectMapOutputWithContext(ctx context.Context) AccountProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountProjectMapOutput)
}

type AccountProjectOutput struct{ *pulumi.OutputState }

func (AccountProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountProject)(nil)).Elem()
}

func (o AccountProjectOutput) ToAccountProjectOutput() AccountProjectOutput {
	return o
}

func (o AccountProjectOutput) ToAccountProjectOutputWithContext(ctx context.Context) AccountProjectOutput {
	return o
}

// The Project creation time.
func (o AccountProjectOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountProject) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the Project.
func (o AccountProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountProject) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the Project.
func (o AccountProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountProject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
func (o AccountProjectOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountProject) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The Project last update time.
func (o AccountProjectOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountProject) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type AccountProjectArrayOutput struct{ *pulumi.OutputState }

func (AccountProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountProject)(nil)).Elem()
}

func (o AccountProjectArrayOutput) ToAccountProjectArrayOutput() AccountProjectArrayOutput {
	return o
}

func (o AccountProjectArrayOutput) ToAccountProjectArrayOutputWithContext(ctx context.Context) AccountProjectArrayOutput {
	return o
}

func (o AccountProjectArrayOutput) Index(i pulumi.IntInput) AccountProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountProject {
		return vs[0].([]*AccountProject)[vs[1].(int)]
	}).(AccountProjectOutput)
}

type AccountProjectMapOutput struct{ *pulumi.OutputState }

func (AccountProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountProject)(nil)).Elem()
}

func (o AccountProjectMapOutput) ToAccountProjectMapOutput() AccountProjectMapOutput {
	return o
}

func (o AccountProjectMapOutput) ToAccountProjectMapOutputWithContext(ctx context.Context) AccountProjectMapOutput {
	return o
}

func (o AccountProjectMapOutput) MapIndex(k pulumi.StringInput) AccountProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountProject {
		return vs[0].(map[string]*AccountProject)[vs[1].(string)]
	}).(AccountProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountProjectInput)(nil)).Elem(), &AccountProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountProjectArrayInput)(nil)).Elem(), AccountProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountProjectMapInput)(nil)).Elem(), AccountProjectMap{})
	pulumi.RegisterOutputType(AccountProjectOutput{})
	pulumi.RegisterOutputType(AccountProjectArrayOutput{})
	pulumi.RegisterOutputType(AccountProjectMapOutput{})
}
