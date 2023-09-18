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

// The provider type for the scaleway package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The Scaleway access key.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The Scaleway API URL to use.
	ApiUrl pulumi.StringPtrOutput `pulumi:"apiUrl"`
	// The Scaleway organization ID.
	OrganizationId pulumi.StringPtrOutput `pulumi:"organizationId"`
	// The Scaleway profile to use.
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// The Scaleway project ID.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The Scaleway secret Key.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.AccessKey == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_ACCESS_KEY"); d != nil {
			args.AccessKey = pulumi.StringPtr(d.(string))
		}
	}
	if args.OrganizationId == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_ORGANIZATION_ID"); d != nil {
			args.OrganizationId = pulumi.StringPtr(d.(string))
		}
	}
	if args.ProjectId == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_PROJECT_ID"); d != nil {
			args.ProjectId = pulumi.StringPtr(d.(string))
		}
	}
	if args.Region == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_REGION"); d != nil {
			args.Region = pulumi.StringPtr(d.(string))
		}
	}
	if args.SecretKey == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_SECRET_KEY"); d != nil {
			args.SecretKey = pulumi.StringPtr(d.(string))
		}
	}
	if args.Zone == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "SCW_DEFAULT_ZONE"); d != nil {
			args.Zone = pulumi.StringPtr(d.(string))
		}
	}
	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringPtrInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:scaleway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The Scaleway access key.
	AccessKey *string `pulumi:"accessKey"`
	// The Scaleway API URL to use.
	ApiUrl *string `pulumi:"apiUrl"`
	// The Scaleway organization ID.
	OrganizationId *string `pulumi:"organizationId"`
	// The Scaleway profile to use.
	Profile *string `pulumi:"profile"`
	// The Scaleway project ID.
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The Scaleway secret Key.
	SecretKey *string `pulumi:"secretKey"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The Scaleway access key.
	AccessKey pulumi.StringPtrInput
	// The Scaleway API URL to use.
	ApiUrl pulumi.StringPtrInput
	// The Scaleway organization ID.
	OrganizationId pulumi.StringPtrInput
	// The Scaleway profile to use.
	Profile pulumi.StringPtrInput
	// The Scaleway project ID.
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The Scaleway secret Key.
	SecretKey pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: i.ToProviderOutputWithContext(ctx).OutputState,
	}
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: o.OutputState,
	}
}

// The Scaleway access key.
func (o ProviderOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The Scaleway API URL to use.
func (o ProviderOutput) ApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiUrl }).(pulumi.StringPtrOutput)
}

// The Scaleway organization ID.
func (o ProviderOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

// The Scaleway profile to use.
func (o ProviderOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Profile }).(pulumi.StringPtrOutput)
}

// The Scaleway project ID.
func (o ProviderOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The region you want to attach the resource to
func (o ProviderOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The Scaleway secret Key.
func (o ProviderOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SecretKey }).(pulumi.StringPtrOutput)
}

// The zone you want to attach the resource to
func (o ProviderOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
