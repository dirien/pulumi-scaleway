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

// Creates and manages Scaleway Container.
//
// For more information consult the [documentation](https://www.scaleway.com/en/docs/faq/serverless-containers/).
//
// For more details about the limitation check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/).
//
// You can check also our [containers guide](https://www.scaleway.com/en/docs/compute/containers/concepts/).
//
// ## Example Usage
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
//			mainContainerNamespace, err := scaleway.NewContainerNamespace(ctx, "mainContainerNamespace", &scaleway.ContainerNamespaceArgs{
//				Description: pulumi.String("test container"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewContainer(ctx, "mainContainer", &scaleway.ContainerArgs{
//				Description: pulumi.String("environment variables test"),
//				NamespaceId: mainContainerNamespace.ID(),
//				RegistryImage: mainContainerNamespace.RegistryEndpoint.ApplyT(func(registryEndpoint string) (string, error) {
//					return fmt.Sprintf("%v/alpine:test", registryEndpoint), nil
//				}).(pulumi.StringOutput),
//				Port:           pulumi.Int(9997),
//				CpuLimit:       pulumi.Int(140),
//				MemoryLimit:    pulumi.Int(256),
//				MinScale:       pulumi.Int(3),
//				MaxScale:       pulumi.Int(5),
//				Timeout:        pulumi.Int(600),
//				MaxConcurrency: pulumi.Int(80),
//				Privacy:        pulumi.String("private"),
//				Protocol:       pulumi.String("http1"),
//				Deploy:         pulumi.Bool(true),
//				EnvironmentVariables: pulumi.StringMap{
//					"foo": pulumi.String("var"),
//				},
//				SecretEnvironmentVariables: pulumi.StringMap{
//					"key": pulumi.String("secret"),
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
// ## Protocols
//
// The supported protocols are:
//
// * `h2c`: HTTP/2 over TCP.
// * `http1`: Hypertext Transfer Protocol.
//
// **Important:** For details about the protocols check [this](https://httpd.apache.org/docs/2.4/howto/http2.html)
//
// ## Privacy
//
// By default, creating a container will make it `public`, meaning that anybody knowing the endpoint could execute it.
// A container can be made `private` with the privacy parameter.
//
// Please check our [authentication](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) section
//
// ## Memory and vCPUs configuration
//
// The vCPU represents a portion or share of the underlying, physical CPU that is assigned to a particular virtual machine (VM).
//
// You may decide how much computing resources to allocate to each container.
// The `memoryLimit` (in MB) must correspond with the right amount of vCPU.
//
// **Important:** The right choice for your container's resources is very important, as you will be billed based on compute usage over time and the number of Containers executions.
//
// Please check our [price](https://www.scaleway.com/en/docs/faq/serverless-containers/#prices) section for more details.
//
// | Memory (in MB) | vCPU |
// |----------------|------|
// | 128            | 70m  |
// | 256            | 140m |
// | 512            | 280m |
// | 1024           | 560m |
//
// **Note:** 560mCPU accounts roughly for half of one CPU power of a Scaleway General Purpose instance
//
// ## Import
//
// Container can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/container:Container main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Container struct {
	pulumi.CustomResourceState

	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	CpuLimit pulumi.IntOutput `pulumi:"cpuLimit"`
	// The cron status of the container.
	CronStatus pulumi.StringOutput `pulumi:"cronStatus"`
	// Boolean controlling whether the container is on a production environment.
	//
	// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
	Deploy pulumi.BoolPtrOutput `pulumi:"deploy"`
	// The description of the container.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The native domain name of the container
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// The error message of the container.
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// HTTP traffic configuration
	HttpOption pulumi.StringPtrOutput `pulumi:"httpOption"`
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency pulumi.IntOutput `pulumi:"maxConcurrency"`
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale pulumi.IntOutput `pulumi:"maxScale"`
	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	MemoryLimit pulumi.IntOutput `pulumi:"memoryLimit"`
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale pulumi.IntOutput `pulumi:"minScale"`
	// The unique name of the container name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	//
	// The following arguments are optional:
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The port to expose the container. Defaults to 8080.
	Port pulumi.IntOutput `pulumi:"port"`
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy pulumi.StringPtrOutput `pulumi:"privacy"`
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage pulumi.StringOutput `pulumi:"registryImage"`
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256 pulumi.StringPtrOutput `pulumi:"registrySha256"`
	// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
	SecretEnvironmentVariables pulumi.StringMapOutput `pulumi:"secretEnvironmentVariables"`
	// The container status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.SecretEnvironmentVariables != nil {
		args.SecretEnvironmentVariables = pulumi.ToSecret(args.SecretEnvironmentVariables).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretEnvironmentVariables",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Container
	err := ctx.RegisterResource("scaleway:index/container:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("scaleway:index/container:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	CpuLimit *int `pulumi:"cpuLimit"`
	// The cron status of the container.
	CronStatus *string `pulumi:"cronStatus"`
	// Boolean controlling whether the container is on a production environment.
	//
	// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
	Deploy *bool `pulumi:"deploy"`
	// The description of the container.
	Description *string `pulumi:"description"`
	// The native domain name of the container
	DomainName *string `pulumi:"domainName"`
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The error message of the container.
	ErrorMessage *string `pulumi:"errorMessage"`
	// HTTP traffic configuration
	HttpOption *string `pulumi:"httpOption"`
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency *int `pulumi:"maxConcurrency"`
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale *int `pulumi:"maxScale"`
	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	MemoryLimit *int `pulumi:"memoryLimit"`
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale *int `pulumi:"minScale"`
	// The unique name of the container name.
	Name *string `pulumi:"name"`
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	//
	// The following arguments are optional:
	NamespaceId *string `pulumi:"namespaceId"`
	// The port to expose the container. Defaults to 8080.
	Port *int `pulumi:"port"`
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy *string `pulumi:"privacy"`
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol *string `pulumi:"protocol"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region *string `pulumi:"region"`
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage *string `pulumi:"registryImage"`
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256 *string `pulumi:"registrySha256"`
	// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
	// The container status.
	Status *string `pulumi:"status"`
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout *int `pulumi:"timeout"`
}

type ContainerState struct {
	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	CpuLimit pulumi.IntPtrInput
	// The cron status of the container.
	CronStatus pulumi.StringPtrInput
	// Boolean controlling whether the container is on a production environment.
	//
	// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
	Deploy pulumi.BoolPtrInput
	// The description of the container.
	Description pulumi.StringPtrInput
	// The native domain name of the container
	DomainName pulumi.StringPtrInput
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables pulumi.StringMapInput
	// The error message of the container.
	ErrorMessage pulumi.StringPtrInput
	// HTTP traffic configuration
	HttpOption pulumi.StringPtrInput
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency pulumi.IntPtrInput
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale pulumi.IntPtrInput
	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	MemoryLimit pulumi.IntPtrInput
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale pulumi.IntPtrInput
	// The unique name of the container name.
	Name pulumi.StringPtrInput
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	//
	// The following arguments are optional:
	NamespaceId pulumi.StringPtrInput
	// The port to expose the container. Defaults to 8080.
	Port pulumi.IntPtrInput
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy pulumi.StringPtrInput
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol pulumi.StringPtrInput
	// (Defaults to provider `region`) The region in which the container was created.
	Region pulumi.StringPtrInput
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage pulumi.StringPtrInput
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256 pulumi.StringPtrInput
	// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
	SecretEnvironmentVariables pulumi.StringMapInput
	// The container status.
	Status pulumi.StringPtrInput
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout pulumi.IntPtrInput
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	CpuLimit *int `pulumi:"cpuLimit"`
	// Boolean controlling whether the container is on a production environment.
	//
	// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
	Deploy *bool `pulumi:"deploy"`
	// The description of the container.
	Description *string `pulumi:"description"`
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// HTTP traffic configuration
	HttpOption *string `pulumi:"httpOption"`
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency *int `pulumi:"maxConcurrency"`
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale *int `pulumi:"maxScale"`
	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	MemoryLimit *int `pulumi:"memoryLimit"`
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale *int `pulumi:"minScale"`
	// The unique name of the container name.
	Name *string `pulumi:"name"`
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	//
	// The following arguments are optional:
	NamespaceId string `pulumi:"namespaceId"`
	// The port to expose the container. Defaults to 8080.
	Port *int `pulumi:"port"`
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy *string `pulumi:"privacy"`
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol *string `pulumi:"protocol"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region *string `pulumi:"region"`
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage *string `pulumi:"registryImage"`
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256 *string `pulumi:"registrySha256"`
	// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
	// The container status.
	Status *string `pulumi:"status"`
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	CpuLimit pulumi.IntPtrInput
	// Boolean controlling whether the container is on a production environment.
	//
	// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
	Deploy pulumi.BoolPtrInput
	// The description of the container.
	Description pulumi.StringPtrInput
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables pulumi.StringMapInput
	// HTTP traffic configuration
	HttpOption pulumi.StringPtrInput
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency pulumi.IntPtrInput
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale pulumi.IntPtrInput
	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	MemoryLimit pulumi.IntPtrInput
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale pulumi.IntPtrInput
	// The unique name of the container name.
	Name pulumi.StringPtrInput
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	//
	// The following arguments are optional:
	NamespaceId pulumi.StringInput
	// The port to expose the container. Defaults to 8080.
	Port pulumi.IntPtrInput
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy pulumi.StringPtrInput
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol pulumi.StringPtrInput
	// (Defaults to provider `region`) The region in which the container was created.
	Region pulumi.StringPtrInput
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage pulumi.StringPtrInput
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256 pulumi.StringPtrInput
	// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
	SecretEnvironmentVariables pulumi.StringMapInput
	// The container status.
	Status pulumi.StringPtrInput
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout pulumi.IntPtrInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (*Container) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i *Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

// ContainerArrayInput is an input type that accepts ContainerArray and ContainerArrayOutput values.
// You can construct a concrete instance of `ContainerArrayInput` via:
//
//	ContainerArray{ ContainerArgs{...} }
type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

// ContainerMapInput is an input type that accepts ContainerMap and ContainerMapOutput values.
// You can construct a concrete instance of `ContainerMapInput` via:
//
//	ContainerMap{ "key": ContainerArgs{...} }
type ContainerMapInput interface {
	pulumi.Input

	ToContainerMapOutput() ContainerMapOutput
	ToContainerMapOutputWithContext(context.Context) ContainerMapOutput
}

type ContainerMap map[string]ContainerInput

func (ContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Container)(nil)).Elem()
}

func (i ContainerMap) ToContainerMapOutput() ContainerMapOutput {
	return i.ToContainerMapOutputWithContext(context.Background())
}

func (i ContainerMap) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerMapOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
func (o ContainerOutput) CpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.CpuLimit }).(pulumi.IntOutput)
}

// The cron status of the container.
func (o ContainerOutput) CronStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.CronStatus }).(pulumi.StringOutput)
}

// Boolean controlling whether the container is on a production environment.
//
// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
func (o ContainerOutput) Deploy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.BoolPtrOutput { return v.Deploy }).(pulumi.BoolPtrOutput)
}

// The description of the container.
func (o ContainerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The native domain name of the container
func (o ContainerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
func (o ContainerOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Container) pulumi.StringMapOutput { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The error message of the container.
func (o ContainerOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// HTTP traffic configuration
func (o ContainerOutput) HttpOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.HttpOption }).(pulumi.StringPtrOutput)
}

// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
func (o ContainerOutput) MaxConcurrency() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.MaxConcurrency }).(pulumi.IntOutput)
}

// The maximum of number of instances this container can scale to. Default to 20.
func (o ContainerOutput) MaxScale() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.MaxScale }).(pulumi.IntOutput)
}

// The memory computing resources in MB to allocate to each container. Defaults to 256.
func (o ContainerOutput) MemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.MemoryLimit }).(pulumi.IntOutput)
}

// The minimum of running container instances continuously. Defaults to 0.
func (o ContainerOutput) MinScale() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.MinScale }).(pulumi.IntOutput)
}

// The unique name of the container name.
func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The container namespace ID of the container.
//
// > **Important** Updates to `name` will recreate the container.
//
// The following arguments are optional:
func (o ContainerOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// The port to expose the container. Defaults to 8080.
func (o ContainerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
func (o ContainerOutput) Privacy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.Privacy }).(pulumi.StringPtrOutput)
}

// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
func (o ContainerOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// (Defaults to provider `region`) The region in which the container was created.
func (o ContainerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
func (o ContainerOutput) RegistryImage() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.RegistryImage }).(pulumi.StringOutput)
}

// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
func (o ContainerOutput) RegistrySha256() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.RegistrySha256 }).(pulumi.StringPtrOutput)
}

// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
func (o ContainerOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Container) pulumi.StringMapOutput { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

// The container status.
func (o ContainerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
func (o ContainerOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Container) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Container)(nil)).Elem()
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Container {
		return vs[0].([]*Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerMapOutput struct{ *pulumi.OutputState }

func (ContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Container)(nil)).Elem()
}

func (o ContainerMapOutput) ToContainerMapOutput() ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) MapIndex(k pulumi.StringInput) ContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Container {
		return vs[0].(map[string]*Container)[vs[1].(string)]
	}).(ContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInput)(nil)).Elem(), &Container{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerArrayInput)(nil)).Elem(), ContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerMapInput)(nil)).Elem(), ContainerMap{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerMapOutput{})
}
