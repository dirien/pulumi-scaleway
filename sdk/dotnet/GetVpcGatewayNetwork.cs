// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway
{
    public static class GetVpcGatewayNetwork
    {
        /// <summary>
        /// Gets information about a gateway network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = ediri.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///         DhcpId = scaleway_vpc_public_gateway_dhcp.Dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpcGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayNetworkId = main.Id,
        ///     });
        /// 
        ///     var byGatewayAndPn = Scaleway.GetVpcGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcGatewayNetworkResult> InvokeAsync(GetVpcGatewayNetworkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcGatewayNetworkResult>("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", args ?? new GetVpcGatewayNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a gateway network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// using Scaleway = ediri.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///         DhcpId = scaleway_vpc_public_gateway_dhcp.Dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var byId = Scaleway.GetVpcGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayNetworkId = main.Id,
        ///     });
        /// 
        ///     var byGatewayAndPn = Scaleway.GetVpcGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcGatewayNetworkResult> Invoke(GetVpcGatewayNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcGatewayNetworkResult>("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", args ?? new GetVpcGatewayNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcGatewayNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the public gateway DHCP config
        /// </summary>
        [Input("dhcpId")]
        public string? DhcpId { get; set; }

        /// <summary>
        /// If masquerade is enabled on requested network
        /// </summary>
        [Input("enableMasquerade")]
        public bool? EnableMasquerade { get; set; }

        /// <summary>
        /// ID of the public gateway the gateway network is linked to
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        /// <summary>
        /// ID of the gateway network.
        /// </summary>
        [Input("gatewayNetworkId")]
        public string? GatewayNetworkId { get; set; }

        /// <summary>
        /// ID of the private network the gateway network is linked to
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        public GetVpcGatewayNetworkArgs()
        {
        }
        public static new GetVpcGatewayNetworkArgs Empty => new GetVpcGatewayNetworkArgs();
    }

    public sealed class GetVpcGatewayNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the public gateway DHCP config
        /// </summary>
        [Input("dhcpId")]
        public Input<string>? DhcpId { get; set; }

        /// <summary>
        /// If masquerade is enabled on requested network
        /// </summary>
        [Input("enableMasquerade")]
        public Input<bool>? EnableMasquerade { get; set; }

        /// <summary>
        /// ID of the public gateway the gateway network is linked to
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// ID of the gateway network.
        /// </summary>
        [Input("gatewayNetworkId")]
        public Input<string>? GatewayNetworkId { get; set; }

        /// <summary>
        /// ID of the private network the gateway network is linked to
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        public GetVpcGatewayNetworkInvokeArgs()
        {
        }
        public static new GetVpcGatewayNetworkInvokeArgs Empty => new GetVpcGatewayNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcGatewayNetworkResult
    {
        public readonly bool CleanupDhcp;
        public readonly string CreatedAt;
        public readonly string? DhcpId;
        public readonly bool EnableDhcp;
        public readonly bool? EnableMasquerade;
        public readonly string? GatewayId;
        public readonly string? GatewayNetworkId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MacAddress;
        public readonly string? PrivateNetworkId;
        public readonly string StaticAddress;
        public readonly string UpdatedAt;
        public readonly string Zone;

        [OutputConstructor]
        private GetVpcGatewayNetworkResult(
            bool cleanupDhcp,

            string createdAt,

            string? dhcpId,

            bool enableDhcp,

            bool? enableMasquerade,

            string? gatewayId,

            string? gatewayNetworkId,

            string id,

            string macAddress,

            string? privateNetworkId,

            string staticAddress,

            string updatedAt,

            string zone)
        {
            CleanupDhcp = cleanupDhcp;
            CreatedAt = createdAt;
            DhcpId = dhcpId;
            EnableDhcp = enableDhcp;
            EnableMasquerade = enableMasquerade;
            GatewayId = gatewayId;
            GatewayNetworkId = gatewayNetworkId;
            Id = id;
            MacAddress = macAddress;
            PrivateNetworkId = privateNetworkId;
            StaticAddress = staticAddress;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
