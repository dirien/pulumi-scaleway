// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the Scaleway Container.
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
//			_ = scaleway.LookupContainerOutput(ctx, scaleway.GetContainerOutputArgs{
//				NamespaceId: mainContainerNamespace.ID(),
//				Name:        mainContainer.Name,
//			}, nil)
//			_ = scaleway.LookupContainerOutput(ctx, scaleway.GetContainerOutputArgs{
//				NamespaceId: mainContainerNamespace.ID(),
//				ContainerId: mainContainer.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupContainer(ctx *pulumi.Context, args *LookupContainerArgs, opts ...pulumi.InvokeOption) (*LookupContainerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerResult
	err := ctx.Invoke("scaleway:index/getContainer:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainer.
type LookupContainerArgs struct {
	ContainerId *string `pulumi:"containerId"`
	// The unique name of the container name.
	Name *string `pulumi:"name"`
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	NamespaceId string `pulumi:"namespaceId"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getContainer.
type LookupContainerResult struct {
	ContainerId *string `pulumi:"containerId"`
	// The amount of vCPU computing resources to allocate to each container. Defaults  to 70.
	CpuLimit int `pulumi:"cpuLimit"`
	// The cron status of the container.
	CronStatus string `pulumi:"cronStatus"`
	// Boolean indicating whether the container is on a production environment.
	Deploy bool `pulumi:"deploy"`
	// The description of the container.
	Description string `pulumi:"description"`
	// The container domain name.
	DomainName string `pulumi:"domainName"`
	// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The error message of the container.
	ErrorMessage string `pulumi:"errorMessage"`
	HttpOption   string `pulumi:"httpOption"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency int `pulumi:"maxConcurrency"`
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale int `pulumi:"maxScale"`
	// The memory computing resources in MB to allocate to each container. Defaults to 128.
	MemoryLimit int `pulumi:"memoryLimit"`
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale    int     `pulumi:"minScale"`
	Name        *string `pulumi:"name"`
	NamespaceId string  `pulumi:"namespaceId"`
	// The port to expose the container. Defaults to 8080.
	Port int `pulumi:"port"`
	// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
	Privacy string `pulumi:"privacy"`
	// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
	Protocol string `pulumi:"protocol"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region *string `pulumi:"region"`
	// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
	RegistryImage string `pulumi:"registryImage"`
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	RegistrySha256             string            `pulumi:"registrySha256"`
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
	// The container status.
	Status string `pulumi:"status"`
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout int `pulumi:"timeout"`
}

func LookupContainerOutput(ctx *pulumi.Context, args LookupContainerOutputArgs, opts ...pulumi.InvokeOption) LookupContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerResult, error) {
			args := v.(LookupContainerArgs)
			r, err := LookupContainer(ctx, &args, opts...)
			var s LookupContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerResultOutput)
}

// A collection of arguments for invoking getContainer.
type LookupContainerOutputArgs struct {
	ContainerId pulumi.StringPtrInput `pulumi:"containerId"`
	// The unique name of the container name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The container namespace ID of the container.
	//
	// > **Important** Updates to `name` will recreate the container.
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
	// (Defaults to provider `region`) The region in which the container was created.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerArgs)(nil)).Elem()
}

// A collection of values returned by getContainer.
type LookupContainerResultOutput struct{ *pulumi.OutputState }

func (LookupContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerResult)(nil)).Elem()
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutput() LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutputWithContext(ctx context.Context) LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

// The amount of vCPU computing resources to allocate to each container. Defaults  to 70.
func (o LookupContainerResultOutput) CpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.CpuLimit }).(pulumi.IntOutput)
}

// The cron status of the container.
func (o LookupContainerResultOutput) CronStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.CronStatus }).(pulumi.StringOutput)
}

// Boolean indicating whether the container is on a production environment.
func (o LookupContainerResultOutput) Deploy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupContainerResult) bool { return v.Deploy }).(pulumi.BoolOutput)
}

// The description of the container.
func (o LookupContainerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Description }).(pulumi.StringOutput)
}

// The container domain name.
func (o LookupContainerResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
func (o LookupContainerResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The error message of the container.
func (o LookupContainerResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) HttpOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.HttpOption }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
func (o LookupContainerResultOutput) MaxConcurrency() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.MaxConcurrency }).(pulumi.IntOutput)
}

// The maximum of number of instances this container can scale to. Default to 20.
func (o LookupContainerResultOutput) MaxScale() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.MaxScale }).(pulumi.IntOutput)
}

// The memory computing resources in MB to allocate to each container. Defaults to 128.
func (o LookupContainerResultOutput) MemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.MemoryLimit }).(pulumi.IntOutput)
}

// The minimum of running container instances continuously. Defaults to 0.
func (o LookupContainerResultOutput) MinScale() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.MinScale }).(pulumi.IntOutput)
}

func (o LookupContainerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupContainerResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// The port to expose the container. Defaults to 8080.
func (o LookupContainerResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.Port }).(pulumi.IntOutput)
}

// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
func (o LookupContainerResultOutput) Privacy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Privacy }).(pulumi.StringOutput)
}

// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
func (o LookupContainerResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// (Defaults to provider `region`) The region in which the container was created.
func (o LookupContainerResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
func (o LookupContainerResultOutput) RegistryImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.RegistryImage }).(pulumi.StringOutput)
}

// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
func (o LookupContainerResultOutput) RegistrySha256() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.RegistrySha256 }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContainerResult) map[string]string { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

// The container status.
func (o LookupContainerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Status }).(pulumi.StringOutput)
}

// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
func (o LookupContainerResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupContainerResult) int { return v.Timeout }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerResultOutput{})
}
