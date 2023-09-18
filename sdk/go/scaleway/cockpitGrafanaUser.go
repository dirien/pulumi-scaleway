// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway Cockpit Grafana Users.
//
// For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users).
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
//			mainCockpit, err := scaleway.LookupCockpit(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewCockpitGrafanaUser(ctx, "mainCockpitGrafanaUser", &scaleway.CockpitGrafanaUserArgs{
//				ProjectId: *pulumi.String(mainCockpit.ProjectId),
//				Login:     pulumi.String("my-awesome-user"),
//				Role:      pulumi.String("editor"),
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
// Cockpits Grafana Users can be imported using the project ID and the grafana user ID formatted `{project_id}/{grafana_user_id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser main 11111111-1111-1111-1111-111111111111/2
//
// ```
type CockpitGrafanaUser struct {
	pulumi.CustomResourceState

	// The login of the grafana user.
	Login pulumi.StringOutput `pulumi:"login"`
	// The password of the grafana user
	Password pulumi.StringOutput `pulumi:"password"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The role of the grafana user. Must be `editor` or `viewer`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCockpitGrafanaUser registers a new resource with the given unique name, arguments, and options.
func NewCockpitGrafanaUser(ctx *pulumi.Context,
	name string, args *CockpitGrafanaUserArgs, opts ...pulumi.ResourceOption) (*CockpitGrafanaUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CockpitGrafanaUser
	err := ctx.RegisterResource("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCockpitGrafanaUser gets an existing CockpitGrafanaUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCockpitGrafanaUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CockpitGrafanaUserState, opts ...pulumi.ResourceOption) (*CockpitGrafanaUser, error) {
	var resource CockpitGrafanaUser
	err := ctx.ReadResource("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CockpitGrafanaUser resources.
type cockpitGrafanaUserState struct {
	// The login of the grafana user.
	Login *string `pulumi:"login"`
	// The password of the grafana user
	Password *string `pulumi:"password"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The role of the grafana user. Must be `editor` or `viewer`.
	Role *string `pulumi:"role"`
}

type CockpitGrafanaUserState struct {
	// The login of the grafana user.
	Login pulumi.StringPtrInput
	// The password of the grafana user
	Password pulumi.StringPtrInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// The role of the grafana user. Must be `editor` or `viewer`.
	Role pulumi.StringPtrInput
}

func (CockpitGrafanaUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitGrafanaUserState)(nil)).Elem()
}

type cockpitGrafanaUserArgs struct {
	// The login of the grafana user.
	Login string `pulumi:"login"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The role of the grafana user. Must be `editor` or `viewer`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CockpitGrafanaUser resource.
type CockpitGrafanaUserArgs struct {
	// The login of the grafana user.
	Login pulumi.StringInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// The role of the grafana user. Must be `editor` or `viewer`.
	Role pulumi.StringInput
}

func (CockpitGrafanaUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitGrafanaUserArgs)(nil)).Elem()
}

type CockpitGrafanaUserInput interface {
	pulumi.Input

	ToCockpitGrafanaUserOutput() CockpitGrafanaUserOutput
	ToCockpitGrafanaUserOutputWithContext(ctx context.Context) CockpitGrafanaUserOutput
}

func (*CockpitGrafanaUser) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitGrafanaUser)(nil)).Elem()
}

func (i *CockpitGrafanaUser) ToCockpitGrafanaUserOutput() CockpitGrafanaUserOutput {
	return i.ToCockpitGrafanaUserOutputWithContext(context.Background())
}

func (i *CockpitGrafanaUser) ToCockpitGrafanaUserOutputWithContext(ctx context.Context) CockpitGrafanaUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitGrafanaUserOutput)
}

func (i *CockpitGrafanaUser) ToOutput(ctx context.Context) pulumix.Output[*CockpitGrafanaUser] {
	return pulumix.Output[*CockpitGrafanaUser]{
		OutputState: i.ToCockpitGrafanaUserOutputWithContext(ctx).OutputState,
	}
}

// CockpitGrafanaUserArrayInput is an input type that accepts CockpitGrafanaUserArray and CockpitGrafanaUserArrayOutput values.
// You can construct a concrete instance of `CockpitGrafanaUserArrayInput` via:
//
//	CockpitGrafanaUserArray{ CockpitGrafanaUserArgs{...} }
type CockpitGrafanaUserArrayInput interface {
	pulumi.Input

	ToCockpitGrafanaUserArrayOutput() CockpitGrafanaUserArrayOutput
	ToCockpitGrafanaUserArrayOutputWithContext(context.Context) CockpitGrafanaUserArrayOutput
}

type CockpitGrafanaUserArray []CockpitGrafanaUserInput

func (CockpitGrafanaUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitGrafanaUser)(nil)).Elem()
}

func (i CockpitGrafanaUserArray) ToCockpitGrafanaUserArrayOutput() CockpitGrafanaUserArrayOutput {
	return i.ToCockpitGrafanaUserArrayOutputWithContext(context.Background())
}

func (i CockpitGrafanaUserArray) ToCockpitGrafanaUserArrayOutputWithContext(ctx context.Context) CockpitGrafanaUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitGrafanaUserArrayOutput)
}

func (i CockpitGrafanaUserArray) ToOutput(ctx context.Context) pulumix.Output[[]*CockpitGrafanaUser] {
	return pulumix.Output[[]*CockpitGrafanaUser]{
		OutputState: i.ToCockpitGrafanaUserArrayOutputWithContext(ctx).OutputState,
	}
}

// CockpitGrafanaUserMapInput is an input type that accepts CockpitGrafanaUserMap and CockpitGrafanaUserMapOutput values.
// You can construct a concrete instance of `CockpitGrafanaUserMapInput` via:
//
//	CockpitGrafanaUserMap{ "key": CockpitGrafanaUserArgs{...} }
type CockpitGrafanaUserMapInput interface {
	pulumi.Input

	ToCockpitGrafanaUserMapOutput() CockpitGrafanaUserMapOutput
	ToCockpitGrafanaUserMapOutputWithContext(context.Context) CockpitGrafanaUserMapOutput
}

type CockpitGrafanaUserMap map[string]CockpitGrafanaUserInput

func (CockpitGrafanaUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitGrafanaUser)(nil)).Elem()
}

func (i CockpitGrafanaUserMap) ToCockpitGrafanaUserMapOutput() CockpitGrafanaUserMapOutput {
	return i.ToCockpitGrafanaUserMapOutputWithContext(context.Background())
}

func (i CockpitGrafanaUserMap) ToCockpitGrafanaUserMapOutputWithContext(ctx context.Context) CockpitGrafanaUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitGrafanaUserMapOutput)
}

func (i CockpitGrafanaUserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CockpitGrafanaUser] {
	return pulumix.Output[map[string]*CockpitGrafanaUser]{
		OutputState: i.ToCockpitGrafanaUserMapOutputWithContext(ctx).OutputState,
	}
}

type CockpitGrafanaUserOutput struct{ *pulumi.OutputState }

func (CockpitGrafanaUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitGrafanaUser)(nil)).Elem()
}

func (o CockpitGrafanaUserOutput) ToCockpitGrafanaUserOutput() CockpitGrafanaUserOutput {
	return o
}

func (o CockpitGrafanaUserOutput) ToCockpitGrafanaUserOutputWithContext(ctx context.Context) CockpitGrafanaUserOutput {
	return o
}

func (o CockpitGrafanaUserOutput) ToOutput(ctx context.Context) pulumix.Output[*CockpitGrafanaUser] {
	return pulumix.Output[*CockpitGrafanaUser]{
		OutputState: o.OutputState,
	}
}

// The login of the grafana user.
func (o CockpitGrafanaUserOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitGrafanaUser) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// The password of the grafana user
func (o CockpitGrafanaUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitGrafanaUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the cockpit is associated with.
func (o CockpitGrafanaUserOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitGrafanaUser) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The role of the grafana user. Must be `editor` or `viewer`.
func (o CockpitGrafanaUserOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitGrafanaUser) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type CockpitGrafanaUserArrayOutput struct{ *pulumi.OutputState }

func (CockpitGrafanaUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitGrafanaUser)(nil)).Elem()
}

func (o CockpitGrafanaUserArrayOutput) ToCockpitGrafanaUserArrayOutput() CockpitGrafanaUserArrayOutput {
	return o
}

func (o CockpitGrafanaUserArrayOutput) ToCockpitGrafanaUserArrayOutputWithContext(ctx context.Context) CockpitGrafanaUserArrayOutput {
	return o
}

func (o CockpitGrafanaUserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CockpitGrafanaUser] {
	return pulumix.Output[[]*CockpitGrafanaUser]{
		OutputState: o.OutputState,
	}
}

func (o CockpitGrafanaUserArrayOutput) Index(i pulumi.IntInput) CockpitGrafanaUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CockpitGrafanaUser {
		return vs[0].([]*CockpitGrafanaUser)[vs[1].(int)]
	}).(CockpitGrafanaUserOutput)
}

type CockpitGrafanaUserMapOutput struct{ *pulumi.OutputState }

func (CockpitGrafanaUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitGrafanaUser)(nil)).Elem()
}

func (o CockpitGrafanaUserMapOutput) ToCockpitGrafanaUserMapOutput() CockpitGrafanaUserMapOutput {
	return o
}

func (o CockpitGrafanaUserMapOutput) ToCockpitGrafanaUserMapOutputWithContext(ctx context.Context) CockpitGrafanaUserMapOutput {
	return o
}

func (o CockpitGrafanaUserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CockpitGrafanaUser] {
	return pulumix.Output[map[string]*CockpitGrafanaUser]{
		OutputState: o.OutputState,
	}
}

func (o CockpitGrafanaUserMapOutput) MapIndex(k pulumi.StringInput) CockpitGrafanaUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CockpitGrafanaUser {
		return vs[0].(map[string]*CockpitGrafanaUser)[vs[1].(string)]
	}).(CockpitGrafanaUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitGrafanaUserInput)(nil)).Elem(), &CockpitGrafanaUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitGrafanaUserArrayInput)(nil)).Elem(), CockpitGrafanaUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitGrafanaUserMapInput)(nil)).Elem(), CockpitGrafanaUserMap{})
	pulumi.RegisterOutputType(CockpitGrafanaUserOutput{})
	pulumi.RegisterOutputType(CockpitGrafanaUserArrayOutput{})
	pulumi.RegisterOutputType(CockpitGrafanaUserMapOutput{})
}
