// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ContainerDomain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
//
// Refer to the Containers domain [documentation](https://www.scaleway.com/en/docs/serverless-containers/how-to/add-a-custom-domain-to-a-container/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-domains-list-all-domain-name-bindings) for more information.
//
// ## Example Usage
//
// The commands below shows how to bind a custom domain name to a container.
//
// ### Simple
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
//			appContainer, err := scaleway.NewContainer(ctx, "appContainer", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewContainerDomain(ctx, "appContainerDomain", &scaleway.ContainerDomainArgs{
//				ContainerId: appContainer.ID(),
//				Hostname:    pulumi.String("container.domain.tld"),
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
// ### Complete example with domain
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := scaleway.NewContainerNamespace(ctx, "main", &scaleway.ContainerNamespaceArgs{
//				Description: pulumi.String("test container"),
//			})
//			if err != nil {
//				return err
//			}
//			appContainer, err := scaleway.NewContainer(ctx, "appContainer", &scaleway.ContainerArgs{
//				NamespaceId: main.ID(),
//				RegistryImage: main.RegistryEndpoint.ApplyT(func(registryEndpoint string) (string, error) {
//					return fmt.Sprintf("%v/nginx:alpine", registryEndpoint), nil
//				}).(pulumi.StringOutput),
//				Port:           pulumi.Int(80),
//				CpuLimit:       pulumi.Int(140),
//				MemoryLimit:    pulumi.Int(256),
//				MinScale:       pulumi.Int(1),
//				MaxScale:       pulumi.Int(1),
//				Timeout:        pulumi.Int(600),
//				MaxConcurrency: pulumi.Int(80),
//				Privacy:        pulumi.String("public"),
//				Protocol:       pulumi.String("http1"),
//				Deploy:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			appDomainRecord, err := scaleway.NewDomainRecord(ctx, "appDomainRecord", &scaleway.DomainRecordArgs{
//				DnsZone: pulumi.String("domain.tld"),
//				Type:    pulumi.String("CNAME"),
//				Data: appContainer.DomainName.ApplyT(func(domainName string) (string, error) {
//					return fmt.Sprintf("%v.", domainName), nil
//				}).(pulumi.StringOutput),
//				Ttl: pulumi.Int(3600),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewContainerDomain(ctx, "appContainerDomain", &scaleway.ContainerDomainArgs{
//				ContainerId: appContainer.ID(),
//				Hostname: pulumi.All(appDomainRecord.Name, appDomainRecord.DnsZone).ApplyT(func(_args []interface{}) (string, error) {
//					name := _args[0].(string)
//					dnsZone := _args[1].(string)
//					return fmt.Sprintf("%v.%v", name, dnsZone), nil
//				}).(pulumi.StringOutput),
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
// Container domain binding can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/containerDomain:ContainerDomain main fr-par/11111111-1111-1111-1111-111111111111
// ```
type ContainerDomain struct {
	pulumi.CustomResourceState

	// The unique identifier of the container.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL used to query the container.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewContainerDomain registers a new resource with the given unique name, arguments, and options.
func NewContainerDomain(ctx *pulumi.Context,
	name string, args *ContainerDomainArgs, opts ...pulumi.ResourceOption) (*ContainerDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerDomain
	err := ctx.RegisterResource("scaleway:index/containerDomain:ContainerDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerDomain gets an existing ContainerDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerDomainState, opts ...pulumi.ResourceOption) (*ContainerDomain, error) {
	var resource ContainerDomain
	err := ctx.ReadResource("scaleway:index/containerDomain:ContainerDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerDomain resources.
type containerDomainState struct {
	// The unique identifier of the container.
	ContainerId *string `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname *string `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region *string `pulumi:"region"`
	// The URL used to query the container.
	Url *string `pulumi:"url"`
}

type ContainerDomainState struct {
	// The unique identifier of the container.
	ContainerId pulumi.StringPtrInput
	// The hostname with a CNAME record.
	Hostname pulumi.StringPtrInput
	// `region`) The region in which the container exists.
	Region pulumi.StringPtrInput
	// The URL used to query the container.
	Url pulumi.StringPtrInput
}

func (ContainerDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerDomainState)(nil)).Elem()
}

type containerDomainArgs struct {
	// The unique identifier of the container.
	ContainerId string `pulumi:"containerId"`
	// The hostname with a CNAME record.
	Hostname string `pulumi:"hostname"`
	// `region`) The region in which the container exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ContainerDomain resource.
type ContainerDomainArgs struct {
	// The unique identifier of the container.
	ContainerId pulumi.StringInput
	// The hostname with a CNAME record.
	Hostname pulumi.StringInput
	// `region`) The region in which the container exists.
	Region pulumi.StringPtrInput
}

func (ContainerDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerDomainArgs)(nil)).Elem()
}

type ContainerDomainInput interface {
	pulumi.Input

	ToContainerDomainOutput() ContainerDomainOutput
	ToContainerDomainOutputWithContext(ctx context.Context) ContainerDomainOutput
}

func (*ContainerDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerDomain)(nil)).Elem()
}

func (i *ContainerDomain) ToContainerDomainOutput() ContainerDomainOutput {
	return i.ToContainerDomainOutputWithContext(context.Background())
}

func (i *ContainerDomain) ToContainerDomainOutputWithContext(ctx context.Context) ContainerDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerDomainOutput)
}

// ContainerDomainArrayInput is an input type that accepts ContainerDomainArray and ContainerDomainArrayOutput values.
// You can construct a concrete instance of `ContainerDomainArrayInput` via:
//
//	ContainerDomainArray{ ContainerDomainArgs{...} }
type ContainerDomainArrayInput interface {
	pulumi.Input

	ToContainerDomainArrayOutput() ContainerDomainArrayOutput
	ToContainerDomainArrayOutputWithContext(context.Context) ContainerDomainArrayOutput
}

type ContainerDomainArray []ContainerDomainInput

func (ContainerDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerDomain)(nil)).Elem()
}

func (i ContainerDomainArray) ToContainerDomainArrayOutput() ContainerDomainArrayOutput {
	return i.ToContainerDomainArrayOutputWithContext(context.Background())
}

func (i ContainerDomainArray) ToContainerDomainArrayOutputWithContext(ctx context.Context) ContainerDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerDomainArrayOutput)
}

// ContainerDomainMapInput is an input type that accepts ContainerDomainMap and ContainerDomainMapOutput values.
// You can construct a concrete instance of `ContainerDomainMapInput` via:
//
//	ContainerDomainMap{ "key": ContainerDomainArgs{...} }
type ContainerDomainMapInput interface {
	pulumi.Input

	ToContainerDomainMapOutput() ContainerDomainMapOutput
	ToContainerDomainMapOutputWithContext(context.Context) ContainerDomainMapOutput
}

type ContainerDomainMap map[string]ContainerDomainInput

func (ContainerDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerDomain)(nil)).Elem()
}

func (i ContainerDomainMap) ToContainerDomainMapOutput() ContainerDomainMapOutput {
	return i.ToContainerDomainMapOutputWithContext(context.Background())
}

func (i ContainerDomainMap) ToContainerDomainMapOutputWithContext(ctx context.Context) ContainerDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerDomainMapOutput)
}

type ContainerDomainOutput struct{ *pulumi.OutputState }

func (ContainerDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerDomain)(nil)).Elem()
}

func (o ContainerDomainOutput) ToContainerDomainOutput() ContainerDomainOutput {
	return o
}

func (o ContainerDomainOutput) ToContainerDomainOutputWithContext(ctx context.Context) ContainerDomainOutput {
	return o
}

// The unique identifier of the container.
func (o ContainerDomainOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerDomain) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// The hostname with a CNAME record.
func (o ContainerDomainOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerDomain) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// `region`) The region in which the container exists.
func (o ContainerDomainOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerDomain) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URL used to query the container.
func (o ContainerDomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerDomain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ContainerDomainArrayOutput struct{ *pulumi.OutputState }

func (ContainerDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerDomain)(nil)).Elem()
}

func (o ContainerDomainArrayOutput) ToContainerDomainArrayOutput() ContainerDomainArrayOutput {
	return o
}

func (o ContainerDomainArrayOutput) ToContainerDomainArrayOutputWithContext(ctx context.Context) ContainerDomainArrayOutput {
	return o
}

func (o ContainerDomainArrayOutput) Index(i pulumi.IntInput) ContainerDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerDomain {
		return vs[0].([]*ContainerDomain)[vs[1].(int)]
	}).(ContainerDomainOutput)
}

type ContainerDomainMapOutput struct{ *pulumi.OutputState }

func (ContainerDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerDomain)(nil)).Elem()
}

func (o ContainerDomainMapOutput) ToContainerDomainMapOutput() ContainerDomainMapOutput {
	return o
}

func (o ContainerDomainMapOutput) ToContainerDomainMapOutputWithContext(ctx context.Context) ContainerDomainMapOutput {
	return o
}

func (o ContainerDomainMapOutput) MapIndex(k pulumi.StringInput) ContainerDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerDomain {
		return vs[0].(map[string]*ContainerDomain)[vs[1].(string)]
	}).(ContainerDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerDomainInput)(nil)).Elem(), &ContainerDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerDomainArrayInput)(nil)).Elem(), ContainerDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerDomainMapInput)(nil)).Elem(), ContainerDomainMap{})
	pulumi.RegisterOutputType(ContainerDomainOutput{})
	pulumi.RegisterOutputType(ContainerDomainArrayOutput{})
	pulumi.RegisterOutputType(ContainerDomainMapOutput{})
}
