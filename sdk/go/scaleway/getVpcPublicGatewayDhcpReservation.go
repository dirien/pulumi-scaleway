// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a DHCP entry. For further information, please see the
// API [documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-dhcp-entries-list-dhcp-entries)/
//
// ## Example Dynamic
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
//			mainVpcPrivateNetwork, err := scaleway.NewVpcPrivateNetwork(ctx, "mainVpcPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			mainInstanceServer, err := scaleway.NewInstanceServer(ctx, "mainInstanceServer", &scaleway.InstanceServerArgs{
//				Image: pulumi.String("ubuntu_jammy"),
//				Type:  pulumi.String("DEV1-S"),
//				Zone:  pulumi.String("fr-par-1"),
//			})
//			if err != nil {
//				return err
//			}
//			mainInstancePrivateNic, err := scaleway.NewInstancePrivateNic(ctx, "mainInstancePrivateNic", &scaleway.InstancePrivateNicArgs{
//				ServerId:         mainInstanceServer.ID(),
//				PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGatewayIp, err := scaleway.NewVpcPublicGatewayIp(ctx, "mainVpcPublicGatewayIp", nil)
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGatewayDhcp, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "mainVpcPublicGatewayDhcp", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("192.168.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGateway, err := scaleway.NewVpcPublicGateway(ctx, "mainVpcPublicGateway", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: mainVpcPublicGatewayIp.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcGatewayNetwork, err := scaleway.NewVpcGatewayNetwork(ctx, "mainVpcGatewayNetwork", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        mainVpcPublicGateway.ID(),
//				PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//				DhcpId:           mainVpcPublicGatewayDhcp.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupVpcPublicGatewayDhcpReservationOutput(ctx, scaleway.GetVpcPublicGatewayDhcpReservationOutputArgs{
//				MacAddress:       mainInstancePrivateNic.MacAddress,
//				GatewayNetworkId: mainVpcGatewayNetwork.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// ## Example Static and PAT Rule
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
//			mainVpcPrivateNetwork, err := scaleway.NewVpcPrivateNetwork(ctx, "mainVpcPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			mainInstanceSecurityGroup, err := scaleway.NewInstanceSecurityGroup(ctx, "mainInstanceSecurityGroup", &scaleway.InstanceSecurityGroupArgs{
//				InboundDefaultPolicy:  pulumi.String("drop"),
//				OutboundDefaultPolicy: pulumi.String("accept"),
//				InboundRules: scaleway.InstanceSecurityGroupInboundRuleArray{
//					&scaleway.InstanceSecurityGroupInboundRuleArgs{
//						Action: pulumi.String("accept"),
//						Port:   pulumi.Int(22),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			mainInstanceServer, err := scaleway.NewInstanceServer(ctx, "mainInstanceServer", &scaleway.InstanceServerArgs{
//				Image:           pulumi.String("ubuntu_jammy"),
//				Type:            pulumi.String("DEV1-S"),
//				Zone:            pulumi.String("fr-par-1"),
//				SecurityGroupId: mainInstanceSecurityGroup.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainInstancePrivateNic, err := scaleway.NewInstancePrivateNic(ctx, "mainInstancePrivateNic", &scaleway.InstancePrivateNicArgs{
//				ServerId:         mainInstanceServer.ID(),
//				PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGatewayIp, err := scaleway.NewVpcPublicGatewayIp(ctx, "mainVpcPublicGatewayIp", nil)
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGatewayDhcp, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "mainVpcPublicGatewayDhcp", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("192.168.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGateway, err := scaleway.NewVpcPublicGateway(ctx, "mainVpcPublicGateway", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: mainVpcPublicGatewayIp.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcGatewayNetwork, err := scaleway.NewVpcGatewayNetwork(ctx, "mainVpcGatewayNetwork", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        mainVpcPublicGateway.ID(),
//				PrivateNetworkId: mainVpcPrivateNetwork.ID(),
//				DhcpId:           mainVpcPublicGatewayDhcp.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			mainVpcPublicGatewayDhcpReservation, err := scaleway.NewVpcPublicGatewayDhcpReservation(ctx, "mainVpcPublicGatewayDhcpReservation", &scaleway.VpcPublicGatewayDhcpReservationArgs{
//				GatewayNetworkId: mainVpcGatewayNetwork.ID(),
//				MacAddress:       mainInstancePrivateNic.MacAddress,
//				IpAddress:        pulumi.String("192.168.1.4"),
//			})
//			if err != nil {
//				return err
//			}
//			// ## VPC PAT RULE
//			_, err = scaleway.NewVpcPublicGatewayPatRule(ctx, "mainVpcPublicGatewayPatRule", &scaleway.VpcPublicGatewayPatRuleArgs{
//				GatewayId:   mainVpcPublicGateway.ID(),
//				PrivateIp:   mainVpcPublicGatewayDhcpReservation.IpAddress,
//				PrivatePort: pulumi.Int(22),
//				PublicPort:  pulumi.Int(2222),
//				Protocol:    pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupVpcPublicGatewayDhcpReservationOutput(ctx, scaleway.GetVpcPublicGatewayDhcpReservationOutputArgs{
//				ReservationId: mainVpcPublicGatewayDhcpReservation.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupVpcPublicGatewayDhcpReservation(ctx *pulumi.Context, args *LookupVpcPublicGatewayDhcpReservationArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayDhcpReservationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayDhcpReservationResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationArgs struct {
	// The ID of the owning GatewayNetwork.
	//
	// > Only one of `reservationId` or `macAddress` with `gatewayNetworkId` should be specified.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the reservation (DHCP entry) to retrieve.
	ReservationId *string `pulumi:"reservationId"`
	// Whether to wait for `macAddress` to exist in DHCP.
	WaitForDhcp *bool `pulumi:"waitForDhcp"`
	// `zone`). The zone in which the reservation exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationResult struct {
	CreatedAt        string  `pulumi:"createdAt"`
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	Hostname         string  `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	IpAddress     string  `pulumi:"ipAddress"`
	MacAddress    *string `pulumi:"macAddress"`
	ReservationId *string `pulumi:"reservationId"`
	Type          string  `pulumi:"type"`
	UpdatedAt     string  `pulumi:"updatedAt"`
	WaitForDhcp   *bool   `pulumi:"waitForDhcp"`
	Zone          *string `pulumi:"zone"`
}

func LookupVpcPublicGatewayDhcpReservationOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayDhcpReservationOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayDhcpReservationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayDhcpReservationResultOutput, error) {
			args := v.(LookupVpcPublicGatewayDhcpReservationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation", args, LookupVpcPublicGatewayDhcpReservationResultOutput{}, options).(LookupVpcPublicGatewayDhcpReservationResultOutput), nil
		}).(LookupVpcPublicGatewayDhcpReservationResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationOutputArgs struct {
	// The ID of the owning GatewayNetwork.
	//
	// > Only one of `reservationId` or `macAddress` with `gatewayNetworkId` should be specified.
	GatewayNetworkId pulumi.StringPtrInput `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the reservation (DHCP entry) to retrieve.
	ReservationId pulumi.StringPtrInput `pulumi:"reservationId"`
	// Whether to wait for `macAddress` to exist in DHCP.
	WaitForDhcp pulumi.BoolPtrInput `pulumi:"waitForDhcp"`
	// `zone`). The zone in which the reservation exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupVpcPublicGatewayDhcpReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpReservationArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayDhcpReservation.
type LookupVpcPublicGatewayDhcpReservationResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayDhcpReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpReservationResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ToLookupVpcPublicGatewayDhcpReservationResultOutput() LookupVpcPublicGatewayDhcpReservationResultOutput {
	return o
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ToLookupVpcPublicGatewayDhcpReservationResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayDhcpReservationResultOutput {
	return o
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) GatewayNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.GatewayNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) ReservationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.ReservationId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) WaitForDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *bool { return v.WaitForDhcp }).(pulumi.BoolPtrOutput)
}

func (o LookupVpcPublicGatewayDhcpReservationResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpReservationResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayDhcpReservationResultOutput{})
}
