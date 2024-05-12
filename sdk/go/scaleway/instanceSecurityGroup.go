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
// Instance security group can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/instanceSecurityGroup:InstanceSecurityGroup web fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type InstanceSecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the security group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
	EnableDefaultSecurity pulumi.BoolPtrOutput `pulumi:"enableDefaultSecurity"`
	// A boolean to specify whether to use instance_security_group_rules.
	// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
	ExternalRules pulumi.BoolPtrOutput `pulumi:"externalRules"`
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"inboundDefaultPolicy"`
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupInboundRuleArrayOutput `pulumi:"inboundRules"`
	// The name of the security group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the security group is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"outboundDefaultPolicy"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// `projectId`) The ID of the project the security group is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// A boolean to specify whether the security group should be stateful or not.
	Stateful pulumi.BoolPtrOutput `pulumi:"stateful"`
	// The tags associated with the security group
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// `zone`) The zone in which the security group should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceSecurityGroup(ctx *pulumi.Context,
	name string, args *InstanceSecurityGroupArgs, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	if args == nil {
		args = &InstanceSecurityGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceSecurityGroup
	err := ctx.RegisterResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceSecurityGroup gets an existing InstanceSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceSecurityGroupState, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	var resource InstanceSecurityGroup
	err := ctx.ReadResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceSecurityGroup resources.
type instanceSecurityGroupState struct {
	// The description of the security group.
	Description *string `pulumi:"description"`
	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	// A boolean to specify whether to use instance_security_group_rules.
	// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
	ExternalRules *bool `pulumi:"externalRules"`
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group.
	Name *string `pulumi:"name"`
	// The organization ID the security group is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// `projectId`) The ID of the project the security group is associated with.
	ProjectId *string `pulumi:"projectId"`
	// A boolean to specify whether the security group should be stateful or not.
	Stateful *bool `pulumi:"stateful"`
	// The tags associated with the security group
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the security group should be created.
	Zone *string `pulumi:"zone"`
}

type InstanceSecurityGroupState struct {
	// The description of the security group.
	Description pulumi.StringPtrInput
	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
	EnableDefaultSecurity pulumi.BoolPtrInput
	// A boolean to specify whether to use instance_security_group_rules.
	// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
	ExternalRules pulumi.BoolPtrInput
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy pulumi.StringPtrInput
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group.
	Name pulumi.StringPtrInput
	// The organization ID the security group is associated with.
	OrganizationId pulumi.StringPtrInput
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy pulumi.StringPtrInput
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// `projectId`) The ID of the project the security group is associated with.
	ProjectId pulumi.StringPtrInput
	// A boolean to specify whether the security group should be stateful or not.
	Stateful pulumi.BoolPtrInput
	// The tags associated with the security group
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the security group should be created.
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupState)(nil)).Elem()
}

type instanceSecurityGroupArgs struct {
	// The description of the security group.
	Description *string `pulumi:"description"`
	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	// A boolean to specify whether to use instance_security_group_rules.
	// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
	ExternalRules *bool `pulumi:"externalRules"`
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group.
	Name *string `pulumi:"name"`
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// `projectId`) The ID of the project the security group is associated with.
	ProjectId *string `pulumi:"projectId"`
	// A boolean to specify whether the security group should be stateful or not.
	Stateful *bool `pulumi:"stateful"`
	// The tags associated with the security group
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the security group should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceSecurityGroup resource.
type InstanceSecurityGroupArgs struct {
	// The description of the security group.
	Description pulumi.StringPtrInput
	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
	EnableDefaultSecurity pulumi.BoolPtrInput
	// A boolean to specify whether to use instance_security_group_rules.
	// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
	ExternalRules pulumi.BoolPtrInput
	// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
	InboundDefaultPolicy pulumi.StringPtrInput
	// A list of inbound rule to add to the security group. (Structure is documented below.)
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group.
	Name pulumi.StringPtrInput
	// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
	OutboundDefaultPolicy pulumi.StringPtrInput
	// A list of outbound rule to add to the security group. (Structure is documented below.)
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// `projectId`) The ID of the project the security group is associated with.
	ProjectId pulumi.StringPtrInput
	// A boolean to specify whether the security group should be stateful or not.
	Stateful pulumi.BoolPtrInput
	// The tags associated with the security group
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the security group should be created.
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupArgs)(nil)).Elem()
}

type InstanceSecurityGroupInput interface {
	pulumi.Input

	ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput
	ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput
}

func (*InstanceSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroup)(nil)).Elem()
}

func (i *InstanceSecurityGroup) ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput {
	return i.ToInstanceSecurityGroupOutputWithContext(context.Background())
}

func (i *InstanceSecurityGroup) ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupOutput)
}

// InstanceSecurityGroupArrayInput is an input type that accepts InstanceSecurityGroupArray and InstanceSecurityGroupArrayOutput values.
// You can construct a concrete instance of `InstanceSecurityGroupArrayInput` via:
//
//	InstanceSecurityGroupArray{ InstanceSecurityGroupArgs{...} }
type InstanceSecurityGroupArrayInput interface {
	pulumi.Input

	ToInstanceSecurityGroupArrayOutput() InstanceSecurityGroupArrayOutput
	ToInstanceSecurityGroupArrayOutputWithContext(context.Context) InstanceSecurityGroupArrayOutput
}

type InstanceSecurityGroupArray []InstanceSecurityGroupInput

func (InstanceSecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSecurityGroup)(nil)).Elem()
}

func (i InstanceSecurityGroupArray) ToInstanceSecurityGroupArrayOutput() InstanceSecurityGroupArrayOutput {
	return i.ToInstanceSecurityGroupArrayOutputWithContext(context.Background())
}

func (i InstanceSecurityGroupArray) ToInstanceSecurityGroupArrayOutputWithContext(ctx context.Context) InstanceSecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupArrayOutput)
}

// InstanceSecurityGroupMapInput is an input type that accepts InstanceSecurityGroupMap and InstanceSecurityGroupMapOutput values.
// You can construct a concrete instance of `InstanceSecurityGroupMapInput` via:
//
//	InstanceSecurityGroupMap{ "key": InstanceSecurityGroupArgs{...} }
type InstanceSecurityGroupMapInput interface {
	pulumi.Input

	ToInstanceSecurityGroupMapOutput() InstanceSecurityGroupMapOutput
	ToInstanceSecurityGroupMapOutputWithContext(context.Context) InstanceSecurityGroupMapOutput
}

type InstanceSecurityGroupMap map[string]InstanceSecurityGroupInput

func (InstanceSecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSecurityGroup)(nil)).Elem()
}

func (i InstanceSecurityGroupMap) ToInstanceSecurityGroupMapOutput() InstanceSecurityGroupMapOutput {
	return i.ToInstanceSecurityGroupMapOutputWithContext(context.Background())
}

func (i InstanceSecurityGroupMap) ToInstanceSecurityGroupMapOutputWithContext(ctx context.Context) InstanceSecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupMapOutput)
}

type InstanceSecurityGroupOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroup)(nil)).Elem()
}

func (o InstanceSecurityGroupOutput) ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput {
	return o
}

func (o InstanceSecurityGroupOutput) ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput {
	return o
}

// The description of the security group.
func (o InstanceSecurityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
func (o InstanceSecurityGroupOutput) EnableDefaultSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.EnableDefaultSecurity }).(pulumi.BoolPtrOutput)
}

// A boolean to specify whether to use instance_security_group_rules.
// If `externalRules` is set to `true`, `inboundRule` and `outboundRule` can not be set directly in the security group.
func (o InstanceSecurityGroupOutput) ExternalRules() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.ExternalRules }).(pulumi.BoolPtrOutput)
}

// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
func (o InstanceSecurityGroupOutput) InboundDefaultPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.InboundDefaultPolicy }).(pulumi.StringPtrOutput)
}

// A list of inbound rule to add to the security group. (Structure is documented below.)
func (o InstanceSecurityGroupOutput) InboundRules() InstanceSecurityGroupInboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) InstanceSecurityGroupInboundRuleArrayOutput { return v.InboundRules }).(InstanceSecurityGroupInboundRuleArrayOutput)
}

// The name of the security group.
func (o InstanceSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the security group is associated with.
func (o InstanceSecurityGroupOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
func (o InstanceSecurityGroupOutput) OutboundDefaultPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.OutboundDefaultPolicy }).(pulumi.StringPtrOutput)
}

// A list of outbound rule to add to the security group. (Structure is documented below.)
func (o InstanceSecurityGroupOutput) OutboundRules() InstanceSecurityGroupOutboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) InstanceSecurityGroupOutboundRuleArrayOutput { return v.OutboundRules }).(InstanceSecurityGroupOutboundRuleArrayOutput)
}

// `projectId`) The ID of the project the security group is associated with.
func (o InstanceSecurityGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// A boolean to specify whether the security group should be stateful or not.
func (o InstanceSecurityGroupOutput) Stateful() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.Stateful }).(pulumi.BoolPtrOutput)
}

// The tags associated with the security group
func (o InstanceSecurityGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// `zone`) The zone in which the security group should be created.
func (o InstanceSecurityGroupOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceSecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceSecurityGroup)(nil)).Elem()
}

func (o InstanceSecurityGroupArrayOutput) ToInstanceSecurityGroupArrayOutput() InstanceSecurityGroupArrayOutput {
	return o
}

func (o InstanceSecurityGroupArrayOutput) ToInstanceSecurityGroupArrayOutputWithContext(ctx context.Context) InstanceSecurityGroupArrayOutput {
	return o
}

func (o InstanceSecurityGroupArrayOutput) Index(i pulumi.IntInput) InstanceSecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceSecurityGroup {
		return vs[0].([]*InstanceSecurityGroup)[vs[1].(int)]
	}).(InstanceSecurityGroupOutput)
}

type InstanceSecurityGroupMapOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceSecurityGroup)(nil)).Elem()
}

func (o InstanceSecurityGroupMapOutput) ToInstanceSecurityGroupMapOutput() InstanceSecurityGroupMapOutput {
	return o
}

func (o InstanceSecurityGroupMapOutput) ToInstanceSecurityGroupMapOutputWithContext(ctx context.Context) InstanceSecurityGroupMapOutput {
	return o
}

func (o InstanceSecurityGroupMapOutput) MapIndex(k pulumi.StringInput) InstanceSecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceSecurityGroup {
		return vs[0].(map[string]*InstanceSecurityGroup)[vs[1].(string)]
	}).(InstanceSecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupInput)(nil)).Elem(), &InstanceSecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupArrayInput)(nil)).Elem(), InstanceSecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupMapInput)(nil)).Elem(), InstanceSecurityGroupMap{})
	pulumi.RegisterOutputType(InstanceSecurityGroupOutput{})
	pulumi.RegisterOutputType(InstanceSecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(InstanceSecurityGroupMapOutput{})
}
