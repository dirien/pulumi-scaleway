// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple IPs managed by IPAM service.
//
// ## Examples
//
// ### By tag
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
//			_, err := scaleway.GetIpamIps(ctx, &scaleway.GetIpamIpsArgs{
//				Tags: []string{
//					"tag",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### By type and resource
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
//			vpc01, err := scaleway.NewVpc(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", &scaleway.VpcPrivateNetworkArgs{
//				VpcId: vpc01.ID(),
//				Ipv4Subnet: &scaleway.VpcPrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.32.0/22"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			redis01, err := scaleway.NewRedisCluster(ctx, "redis01", &scaleway.RedisClusterArgs{
//				Version:     pulumi.String("7.0.5"),
//				NodeType:    pulumi.String("RED1-XS"),
//				UserName:    pulumi.String("my_initial_user"),
//				Password:    pulumi.String("thiZ_is_v&ry_s3cret"),
//				ClusterSize: pulumi.Int(3),
//				PrivateNetworks: scaleway.RedisClusterPrivateNetworkArray{
//					&scaleway.RedisClusterPrivateNetworkArgs{
//						Id: pn01.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.GetIpamIpsOutput(ctx, scaleway.GetIpamIpsOutputArgs{
//				Type: pulumi.String("ipv4"),
//				Resource: &scaleway.GetIpamIpsResourceArgs{
//					Id:   redis01.ID(),
//					Type: pulumi.String("redis_cluster"),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetIpamIps(ctx *pulumi.Context, args *GetIpamIpsArgs, opts ...pulumi.InvokeOption) (*GetIpamIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpamIpsResult
	err := ctx.Invoke("scaleway:index/getIpamIps:getIpamIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpamIps.
type GetIpamIpsArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource.
	Attached *bool `pulumi:"attached"`
	// The Mac Address used as filter.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network used as filter.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the project used as filter.
	ProjectId *string `pulumi:"projectId"`
	// The region used as filter.
	Region *string `pulumi:"region"`
	// Filter by resource ID, type or name.
	Resource *GetIpamIpsResource `pulumi:"resource"`
	// The tags used as filter.
	Tags []string `pulumi:"tags"`
	// The type of IP used as filter (ipv4, ipv6).
	Type *string `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal *string `pulumi:"zonal"`
}

// A collection of values returned by getIpamIps.
type GetIpamIpsResult struct {
	Attached *bool `pulumi:"attached"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of found IPs
	Ips []GetIpamIpsIp `pulumi:"ips"`
	// The mac address.
	MacAddress       *string `pulumi:"macAddress"`
	OrganizationId   string  `pulumi:"organizationId"`
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the project the server is associated with.
	ProjectId string `pulumi:"projectId"`
	// The region in which the IP is.
	Region string `pulumi:"region"`
	// The list of public IPs of the server.
	Resource *GetIpamIpsResource `pulumi:"resource"`
	// The tags associated with the IP.
	Tags []string `pulumi:"tags"`
	// The type of resource.
	Type  *string `pulumi:"type"`
	Zonal string  `pulumi:"zonal"`
}

func GetIpamIpsOutput(ctx *pulumi.Context, args GetIpamIpsOutputArgs, opts ...pulumi.InvokeOption) GetIpamIpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpamIpsResult, error) {
			args := v.(GetIpamIpsArgs)
			r, err := GetIpamIps(ctx, &args, opts...)
			var s GetIpamIpsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpamIpsResultOutput)
}

// A collection of arguments for invoking getIpamIps.
type GetIpamIpsOutputArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource.
	Attached pulumi.BoolPtrInput `pulumi:"attached"`
	// The Mac Address used as filter.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the private network used as filter.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// The ID of the project used as filter.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region used as filter.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Filter by resource ID, type or name.
	Resource GetIpamIpsResourcePtrInput `pulumi:"resource"`
	// The tags used as filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The type of IP used as filter (ipv4, ipv6).
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal pulumi.StringPtrInput `pulumi:"zonal"`
}

func (GetIpamIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamIpsArgs)(nil)).Elem()
}

// A collection of values returned by getIpamIps.
type GetIpamIpsResultOutput struct{ *pulumi.OutputState }

func (GetIpamIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamIpsResult)(nil)).Elem()
}

func (o GetIpamIpsResultOutput) ToGetIpamIpsResultOutput() GetIpamIpsResultOutput {
	return o
}

func (o GetIpamIpsResultOutput) ToGetIpamIpsResultOutputWithContext(ctx context.Context) GetIpamIpsResultOutput {
	return o
}

func (o GetIpamIpsResultOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIpamIpsResult) *bool { return v.Attached }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpamIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of found IPs
func (o GetIpamIpsResultOutput) Ips() GetIpamIpsIpArrayOutput {
	return o.ApplyT(func(v GetIpamIpsResult) []GetIpamIpsIp { return v.Ips }).(GetIpamIpsIpArrayOutput)
}

// The mac address.
func (o GetIpamIpsResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpamIpsResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o GetIpamIpsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetIpamIpsResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpamIpsResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

// The ID of the project the server is associated with.
func (o GetIpamIpsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which the IP is.
func (o GetIpamIpsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpsResult) string { return v.Region }).(pulumi.StringOutput)
}

// The list of public IPs of the server.
func (o GetIpamIpsResultOutput) Resource() GetIpamIpsResourcePtrOutput {
	return o.ApplyT(func(v GetIpamIpsResult) *GetIpamIpsResource { return v.Resource }).(GetIpamIpsResourcePtrOutput)
}

// The tags associated with the IP.
func (o GetIpamIpsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpamIpsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of resource.
func (o GetIpamIpsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpamIpsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetIpamIpsResultOutput) Zonal() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamIpsResult) string { return v.Zonal }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpamIpsResultOutput{})
}