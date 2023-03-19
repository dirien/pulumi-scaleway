// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Load-Balancer Backends.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewLbBackend(ctx, "backend01", &scaleway.LbBackendArgs{
//				LbId:            pulumi.Any(scaleway_lb.Lb01.Id),
//				ForwardProtocol: pulumi.String("http"),
//				ForwardPort:     pulumi.Int(80),
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
// ### With HTTP Health Check
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewLbBackend(ctx, "backend01", &scaleway.LbBackendArgs{
//				LbId:            pulumi.Any(scaleway_lb.Lb01.Id),
//				ForwardProtocol: pulumi.String("http"),
//				ForwardPort:     pulumi.Int(80),
//				HealthCheckHttp: &scaleway.LbBackendHealthCheckHttpArgs{
//					Uri: pulumi.String("www.test.com/health"),
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
// Load-Balancer backend can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/lbBackend:LbBackend backend01 fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type LbBackend struct {
	pulumi.CustomResourceState

	// Scaleway S3 bucket website to be served in case all backend servers are down.
	// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
	// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
	FailoverHost pulumi.StringPtrOutput `pulumi:"failoverHost"`
	// User sessions will be forwarded to this port of backend servers.
	ForwardPort pulumi.IntOutput `pulumi:"forwardPort"`
	// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
	ForwardPortAlgorithm pulumi.StringPtrOutput `pulumi:"forwardPortAlgorithm"`
	// Backend protocol. Possible values are: `tcp` or `http`.
	ForwardProtocol pulumi.StringOutput `pulumi:"forwardProtocol"`
	// Interval between two HC requests.
	HealthCheckDelay pulumi.StringPtrOutput `pulumi:"healthCheckDelay"`
	// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttp LbBackendHealthCheckHttpPtrOutput `pulumi:"healthCheckHttp"`
	// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttps LbBackendHealthCheckHttpsPtrOutput `pulumi:"healthCheckHttps"`
	// Number of allowed failed HC requests before the backend server is marked down.
	HealthCheckMaxRetries pulumi.IntPtrOutput `pulumi:"healthCheckMaxRetries"`
	// Port the HC requests will be send to.
	HealthCheckPort pulumi.IntOutput `pulumi:"healthCheckPort"`
	// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckTcp LbBackendHealthCheckTcpOutput `pulumi:"healthCheckTcp"`
	// Timeout before we consider a HC request failed.
	HealthCheckTimeout pulumi.StringPtrOutput `pulumi:"healthCheckTimeout"`
	// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
	IgnoreSslServerVerify pulumi.BoolPtrOutput `pulumi:"ignoreSslServerVerify"`
	// The load-balancer ID this backend is attached to.
	// > **Important:** Updates to `lbId` will recreate the backend.
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The name of the load-balancer backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
	OnMarkedDownAction pulumi.StringPtrOutput `pulumi:"onMarkedDownAction"`
	// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
	ProxyProtocol pulumi.StringPtrOutput `pulumi:"proxyProtocol"`
	// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
	//
	// Deprecated: Please use proxy_protocol instead
	SendProxyV2 pulumi.BoolPtrOutput `pulumi:"sendProxyV2"`
	// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
	ServerIps pulumi.StringArrayOutput `pulumi:"serverIps"`
	// Enables SSL between load balancer and backend servers.
	SslBridging pulumi.BoolPtrOutput `pulumi:"sslBridging"`
	// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
	StickySessions pulumi.StringPtrOutput `pulumi:"stickySessions"`
	// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
	StickySessionsCookieName pulumi.StringPtrOutput `pulumi:"stickySessionsCookieName"`
	// Maximum initial server connection establishment time. (e.g.: `1s`)
	TimeoutConnect pulumi.StringPtrOutput `pulumi:"timeoutConnect"`
	// Maximum server connection inactivity time. (e.g.: `1s`)
	TimeoutServer pulumi.StringPtrOutput `pulumi:"timeoutServer"`
	// Maximum tunnel inactivity time. (e.g.: `1s`)
	TimeoutTunnel pulumi.StringPtrOutput `pulumi:"timeoutTunnel"`
}

// NewLbBackend registers a new resource with the given unique name, arguments, and options.
func NewLbBackend(ctx *pulumi.Context,
	name string, args *LbBackendArgs, opts ...pulumi.ResourceOption) (*LbBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ForwardPort == nil {
		return nil, errors.New("invalid value for required argument 'ForwardPort'")
	}
	if args.ForwardProtocol == nil {
		return nil, errors.New("invalid value for required argument 'ForwardProtocol'")
	}
	if args.LbId == nil {
		return nil, errors.New("invalid value for required argument 'LbId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LbBackend
	err := ctx.RegisterResource("scaleway:index/lbBackend:LbBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbBackend gets an existing LbBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbBackendState, opts ...pulumi.ResourceOption) (*LbBackend, error) {
	var resource LbBackend
	err := ctx.ReadResource("scaleway:index/lbBackend:LbBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbBackend resources.
type lbBackendState struct {
	// Scaleway S3 bucket website to be served in case all backend servers are down.
	// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
	// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
	FailoverHost *string `pulumi:"failoverHost"`
	// User sessions will be forwarded to this port of backend servers.
	ForwardPort *int `pulumi:"forwardPort"`
	// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
	ForwardPortAlgorithm *string `pulumi:"forwardPortAlgorithm"`
	// Backend protocol. Possible values are: `tcp` or `http`.
	ForwardProtocol *string `pulumi:"forwardProtocol"`
	// Interval between two HC requests.
	HealthCheckDelay *string `pulumi:"healthCheckDelay"`
	// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttp *LbBackendHealthCheckHttp `pulumi:"healthCheckHttp"`
	// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttps *LbBackendHealthCheckHttps `pulumi:"healthCheckHttps"`
	// Number of allowed failed HC requests before the backend server is marked down.
	HealthCheckMaxRetries *int `pulumi:"healthCheckMaxRetries"`
	// Port the HC requests will be send to.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckTcp *LbBackendHealthCheckTcp `pulumi:"healthCheckTcp"`
	// Timeout before we consider a HC request failed.
	HealthCheckTimeout *string `pulumi:"healthCheckTimeout"`
	// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
	IgnoreSslServerVerify *bool `pulumi:"ignoreSslServerVerify"`
	// The load-balancer ID this backend is attached to.
	// > **Important:** Updates to `lbId` will recreate the backend.
	LbId *string `pulumi:"lbId"`
	// The name of the load-balancer backend.
	Name *string `pulumi:"name"`
	// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
	OnMarkedDownAction *string `pulumi:"onMarkedDownAction"`
	// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
	ProxyProtocol *string `pulumi:"proxyProtocol"`
	// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
	//
	// Deprecated: Please use proxy_protocol instead
	SendProxyV2 *bool `pulumi:"sendProxyV2"`
	// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
	ServerIps []string `pulumi:"serverIps"`
	// Enables SSL between load balancer and backend servers.
	SslBridging *bool `pulumi:"sslBridging"`
	// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
	StickySessions *string `pulumi:"stickySessions"`
	// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
	StickySessionsCookieName *string `pulumi:"stickySessionsCookieName"`
	// Maximum initial server connection establishment time. (e.g.: `1s`)
	TimeoutConnect *string `pulumi:"timeoutConnect"`
	// Maximum server connection inactivity time. (e.g.: `1s`)
	TimeoutServer *string `pulumi:"timeoutServer"`
	// Maximum tunnel inactivity time. (e.g.: `1s`)
	TimeoutTunnel *string `pulumi:"timeoutTunnel"`
}

type LbBackendState struct {
	// Scaleway S3 bucket website to be served in case all backend servers are down.
	// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
	// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
	FailoverHost pulumi.StringPtrInput
	// User sessions will be forwarded to this port of backend servers.
	ForwardPort pulumi.IntPtrInput
	// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
	ForwardPortAlgorithm pulumi.StringPtrInput
	// Backend protocol. Possible values are: `tcp` or `http`.
	ForwardProtocol pulumi.StringPtrInput
	// Interval between two HC requests.
	HealthCheckDelay pulumi.StringPtrInput
	// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttp LbBackendHealthCheckHttpPtrInput
	// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttps LbBackendHealthCheckHttpsPtrInput
	// Number of allowed failed HC requests before the backend server is marked down.
	HealthCheckMaxRetries pulumi.IntPtrInput
	// Port the HC requests will be send to.
	HealthCheckPort pulumi.IntPtrInput
	// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckTcp LbBackendHealthCheckTcpPtrInput
	// Timeout before we consider a HC request failed.
	HealthCheckTimeout pulumi.StringPtrInput
	// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
	IgnoreSslServerVerify pulumi.BoolPtrInput
	// The load-balancer ID this backend is attached to.
	// > **Important:** Updates to `lbId` will recreate the backend.
	LbId pulumi.StringPtrInput
	// The name of the load-balancer backend.
	Name pulumi.StringPtrInput
	// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
	OnMarkedDownAction pulumi.StringPtrInput
	// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
	ProxyProtocol pulumi.StringPtrInput
	// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
	//
	// Deprecated: Please use proxy_protocol instead
	SendProxyV2 pulumi.BoolPtrInput
	// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
	ServerIps pulumi.StringArrayInput
	// Enables SSL between load balancer and backend servers.
	SslBridging pulumi.BoolPtrInput
	// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
	StickySessions pulumi.StringPtrInput
	// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
	StickySessionsCookieName pulumi.StringPtrInput
	// Maximum initial server connection establishment time. (e.g.: `1s`)
	TimeoutConnect pulumi.StringPtrInput
	// Maximum server connection inactivity time. (e.g.: `1s`)
	TimeoutServer pulumi.StringPtrInput
	// Maximum tunnel inactivity time. (e.g.: `1s`)
	TimeoutTunnel pulumi.StringPtrInput
}

func (LbBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbBackendState)(nil)).Elem()
}

type lbBackendArgs struct {
	// Scaleway S3 bucket website to be served in case all backend servers are down.
	// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
	// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
	FailoverHost *string `pulumi:"failoverHost"`
	// User sessions will be forwarded to this port of backend servers.
	ForwardPort int `pulumi:"forwardPort"`
	// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
	ForwardPortAlgorithm *string `pulumi:"forwardPortAlgorithm"`
	// Backend protocol. Possible values are: `tcp` or `http`.
	ForwardProtocol string `pulumi:"forwardProtocol"`
	// Interval between two HC requests.
	HealthCheckDelay *string `pulumi:"healthCheckDelay"`
	// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttp *LbBackendHealthCheckHttp `pulumi:"healthCheckHttp"`
	// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttps *LbBackendHealthCheckHttps `pulumi:"healthCheckHttps"`
	// Number of allowed failed HC requests before the backend server is marked down.
	HealthCheckMaxRetries *int `pulumi:"healthCheckMaxRetries"`
	// Port the HC requests will be send to.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckTcp *LbBackendHealthCheckTcp `pulumi:"healthCheckTcp"`
	// Timeout before we consider a HC request failed.
	HealthCheckTimeout *string `pulumi:"healthCheckTimeout"`
	// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
	IgnoreSslServerVerify *bool `pulumi:"ignoreSslServerVerify"`
	// The load-balancer ID this backend is attached to.
	// > **Important:** Updates to `lbId` will recreate the backend.
	LbId string `pulumi:"lbId"`
	// The name of the load-balancer backend.
	Name *string `pulumi:"name"`
	// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
	OnMarkedDownAction *string `pulumi:"onMarkedDownAction"`
	// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
	ProxyProtocol *string `pulumi:"proxyProtocol"`
	// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
	//
	// Deprecated: Please use proxy_protocol instead
	SendProxyV2 *bool `pulumi:"sendProxyV2"`
	// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
	ServerIps []string `pulumi:"serverIps"`
	// Enables SSL between load balancer and backend servers.
	SslBridging *bool `pulumi:"sslBridging"`
	// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
	StickySessions *string `pulumi:"stickySessions"`
	// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
	StickySessionsCookieName *string `pulumi:"stickySessionsCookieName"`
	// Maximum initial server connection establishment time. (e.g.: `1s`)
	TimeoutConnect *string `pulumi:"timeoutConnect"`
	// Maximum server connection inactivity time. (e.g.: `1s`)
	TimeoutServer *string `pulumi:"timeoutServer"`
	// Maximum tunnel inactivity time. (e.g.: `1s`)
	TimeoutTunnel *string `pulumi:"timeoutTunnel"`
}

// The set of arguments for constructing a LbBackend resource.
type LbBackendArgs struct {
	// Scaleway S3 bucket website to be served in case all backend servers are down.
	// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
	// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
	FailoverHost pulumi.StringPtrInput
	// User sessions will be forwarded to this port of backend servers.
	ForwardPort pulumi.IntInput
	// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
	ForwardPortAlgorithm pulumi.StringPtrInput
	// Backend protocol. Possible values are: `tcp` or `http`.
	ForwardProtocol pulumi.StringInput
	// Interval between two HC requests.
	HealthCheckDelay pulumi.StringPtrInput
	// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttp LbBackendHealthCheckHttpPtrInput
	// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckHttps LbBackendHealthCheckHttpsPtrInput
	// Number of allowed failed HC requests before the backend server is marked down.
	HealthCheckMaxRetries pulumi.IntPtrInput
	// Port the HC requests will be send to.
	HealthCheckPort pulumi.IntPtrInput
	// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
	HealthCheckTcp LbBackendHealthCheckTcpPtrInput
	// Timeout before we consider a HC request failed.
	HealthCheckTimeout pulumi.StringPtrInput
	// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
	IgnoreSslServerVerify pulumi.BoolPtrInput
	// The load-balancer ID this backend is attached to.
	// > **Important:** Updates to `lbId` will recreate the backend.
	LbId pulumi.StringInput
	// The name of the load-balancer backend.
	Name pulumi.StringPtrInput
	// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
	OnMarkedDownAction pulumi.StringPtrInput
	// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
	ProxyProtocol pulumi.StringPtrInput
	// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
	//
	// Deprecated: Please use proxy_protocol instead
	SendProxyV2 pulumi.BoolPtrInput
	// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
	ServerIps pulumi.StringArrayInput
	// Enables SSL between load balancer and backend servers.
	SslBridging pulumi.BoolPtrInput
	// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
	StickySessions pulumi.StringPtrInput
	// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
	StickySessionsCookieName pulumi.StringPtrInput
	// Maximum initial server connection establishment time. (e.g.: `1s`)
	TimeoutConnect pulumi.StringPtrInput
	// Maximum server connection inactivity time. (e.g.: `1s`)
	TimeoutServer pulumi.StringPtrInput
	// Maximum tunnel inactivity time. (e.g.: `1s`)
	TimeoutTunnel pulumi.StringPtrInput
}

func (LbBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbBackendArgs)(nil)).Elem()
}

type LbBackendInput interface {
	pulumi.Input

	ToLbBackendOutput() LbBackendOutput
	ToLbBackendOutputWithContext(ctx context.Context) LbBackendOutput
}

func (*LbBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**LbBackend)(nil)).Elem()
}

func (i *LbBackend) ToLbBackendOutput() LbBackendOutput {
	return i.ToLbBackendOutputWithContext(context.Background())
}

func (i *LbBackend) ToLbBackendOutputWithContext(ctx context.Context) LbBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbBackendOutput)
}

// LbBackendArrayInput is an input type that accepts LbBackendArray and LbBackendArrayOutput values.
// You can construct a concrete instance of `LbBackendArrayInput` via:
//
//	LbBackendArray{ LbBackendArgs{...} }
type LbBackendArrayInput interface {
	pulumi.Input

	ToLbBackendArrayOutput() LbBackendArrayOutput
	ToLbBackendArrayOutputWithContext(context.Context) LbBackendArrayOutput
}

type LbBackendArray []LbBackendInput

func (LbBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbBackend)(nil)).Elem()
}

func (i LbBackendArray) ToLbBackendArrayOutput() LbBackendArrayOutput {
	return i.ToLbBackendArrayOutputWithContext(context.Background())
}

func (i LbBackendArray) ToLbBackendArrayOutputWithContext(ctx context.Context) LbBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbBackendArrayOutput)
}

// LbBackendMapInput is an input type that accepts LbBackendMap and LbBackendMapOutput values.
// You can construct a concrete instance of `LbBackendMapInput` via:
//
//	LbBackendMap{ "key": LbBackendArgs{...} }
type LbBackendMapInput interface {
	pulumi.Input

	ToLbBackendMapOutput() LbBackendMapOutput
	ToLbBackendMapOutputWithContext(context.Context) LbBackendMapOutput
}

type LbBackendMap map[string]LbBackendInput

func (LbBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbBackend)(nil)).Elem()
}

func (i LbBackendMap) ToLbBackendMapOutput() LbBackendMapOutput {
	return i.ToLbBackendMapOutputWithContext(context.Background())
}

func (i LbBackendMap) ToLbBackendMapOutputWithContext(ctx context.Context) LbBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbBackendMapOutput)
}

type LbBackendOutput struct{ *pulumi.OutputState }

func (LbBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbBackend)(nil)).Elem()
}

func (o LbBackendOutput) ToLbBackendOutput() LbBackendOutput {
	return o
}

func (o LbBackendOutput) ToLbBackendOutputWithContext(ctx context.Context) LbBackendOutput {
	return o
}

// Scaleway S3 bucket website to be served in case all backend servers are down.
// > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
// e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
func (o LbBackendOutput) FailoverHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.FailoverHost }).(pulumi.StringPtrOutput)
}

// User sessions will be forwarded to this port of backend servers.
func (o LbBackendOutput) ForwardPort() pulumi.IntOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.IntOutput { return v.ForwardPort }).(pulumi.IntOutput)
}

// Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
func (o LbBackendOutput) ForwardPortAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.ForwardPortAlgorithm }).(pulumi.StringPtrOutput)
}

// Backend protocol. Possible values are: `tcp` or `http`.
func (o LbBackendOutput) ForwardProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringOutput { return v.ForwardProtocol }).(pulumi.StringOutput)
}

// Interval between two HC requests.
func (o LbBackendOutput) HealthCheckDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.HealthCheckDelay }).(pulumi.StringPtrOutput)
}

// This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
func (o LbBackendOutput) HealthCheckHttp() LbBackendHealthCheckHttpPtrOutput {
	return o.ApplyT(func(v *LbBackend) LbBackendHealthCheckHttpPtrOutput { return v.HealthCheckHttp }).(LbBackendHealthCheckHttpPtrOutput)
}

// This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
func (o LbBackendOutput) HealthCheckHttps() LbBackendHealthCheckHttpsPtrOutput {
	return o.ApplyT(func(v *LbBackend) LbBackendHealthCheckHttpsPtrOutput { return v.HealthCheckHttps }).(LbBackendHealthCheckHttpsPtrOutput)
}

// Number of allowed failed HC requests before the backend server is marked down.
func (o LbBackendOutput) HealthCheckMaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.IntPtrOutput { return v.HealthCheckMaxRetries }).(pulumi.IntPtrOutput)
}

// Port the HC requests will be send to.
func (o LbBackendOutput) HealthCheckPort() pulumi.IntOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.IntOutput { return v.HealthCheckPort }).(pulumi.IntOutput)
}

// This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
func (o LbBackendOutput) HealthCheckTcp() LbBackendHealthCheckTcpOutput {
	return o.ApplyT(func(v *LbBackend) LbBackendHealthCheckTcpOutput { return v.HealthCheckTcp }).(LbBackendHealthCheckTcpOutput)
}

// Timeout before we consider a HC request failed.
func (o LbBackendOutput) HealthCheckTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.HealthCheckTimeout }).(pulumi.StringPtrOutput)
}

// Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
func (o LbBackendOutput) IgnoreSslServerVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.BoolPtrOutput { return v.IgnoreSslServerVerify }).(pulumi.BoolPtrOutput)
}

// The load-balancer ID this backend is attached to.
// > **Important:** Updates to `lbId` will recreate the backend.
func (o LbBackendOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// The name of the load-balancer backend.
func (o LbBackendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
func (o LbBackendOutput) OnMarkedDownAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.OnMarkedDownAction }).(pulumi.StringPtrOutput)
}

// Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
func (o LbBackendOutput) ProxyProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.ProxyProtocol }).(pulumi.StringPtrOutput)
}

// DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
//
// Deprecated: Please use proxy_protocol instead
func (o LbBackendOutput) SendProxyV2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.BoolPtrOutput { return v.SendProxyV2 }).(pulumi.BoolPtrOutput)
}

// List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
func (o LbBackendOutput) ServerIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringArrayOutput { return v.ServerIps }).(pulumi.StringArrayOutput)
}

// Enables SSL between load balancer and backend servers.
func (o LbBackendOutput) SslBridging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.BoolPtrOutput { return v.SslBridging }).(pulumi.BoolPtrOutput)
}

// Load balancing algorithm. Possible values are: `none`, `cookie` and `table`.
func (o LbBackendOutput) StickySessions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.StickySessions }).(pulumi.StringPtrOutput)
}

// Cookie name for for sticky sessions. Only applicable when stickySessions is set to `cookie`.
func (o LbBackendOutput) StickySessionsCookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.StickySessionsCookieName }).(pulumi.StringPtrOutput)
}

// Maximum initial server connection establishment time. (e.g.: `1s`)
func (o LbBackendOutput) TimeoutConnect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.TimeoutConnect }).(pulumi.StringPtrOutput)
}

// Maximum server connection inactivity time. (e.g.: `1s`)
func (o LbBackendOutput) TimeoutServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.TimeoutServer }).(pulumi.StringPtrOutput)
}

// Maximum tunnel inactivity time. (e.g.: `1s`)
func (o LbBackendOutput) TimeoutTunnel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbBackend) pulumi.StringPtrOutput { return v.TimeoutTunnel }).(pulumi.StringPtrOutput)
}

type LbBackendArrayOutput struct{ *pulumi.OutputState }

func (LbBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbBackend)(nil)).Elem()
}

func (o LbBackendArrayOutput) ToLbBackendArrayOutput() LbBackendArrayOutput {
	return o
}

func (o LbBackendArrayOutput) ToLbBackendArrayOutputWithContext(ctx context.Context) LbBackendArrayOutput {
	return o
}

func (o LbBackendArrayOutput) Index(i pulumi.IntInput) LbBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbBackend {
		return vs[0].([]*LbBackend)[vs[1].(int)]
	}).(LbBackendOutput)
}

type LbBackendMapOutput struct{ *pulumi.OutputState }

func (LbBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbBackend)(nil)).Elem()
}

func (o LbBackendMapOutput) ToLbBackendMapOutput() LbBackendMapOutput {
	return o
}

func (o LbBackendMapOutput) ToLbBackendMapOutputWithContext(ctx context.Context) LbBackendMapOutput {
	return o
}

func (o LbBackendMapOutput) MapIndex(k pulumi.StringInput) LbBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbBackend {
		return vs[0].(map[string]*LbBackend)[vs[1].(string)]
	}).(LbBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbBackendInput)(nil)).Elem(), &LbBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbBackendArrayInput)(nil)).Elem(), LbBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbBackendMapInput)(nil)).Elem(), LbBackendMap{})
	pulumi.RegisterOutputType(LbBackendOutput{})
	pulumi.RegisterOutputType(LbBackendArrayOutput{})
	pulumi.RegisterOutputType(LbBackendMapOutput{})
}
