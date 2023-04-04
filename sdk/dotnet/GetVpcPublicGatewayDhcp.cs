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
    public static class GetVpcPublicGatewayDhcp
    {
        /// <summary>
        /// Gets information about a public gateway DHCP.
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
        ///     var main = new Scaleway.VpcPublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var dhcpById = Scaleway.GetVpcPublicGatewayDhcp.Invoke(new()
        ///     {
        ///         DhcpId = main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcPublicGatewayDhcpResult> InvokeAsync(GetVpcPublicGatewayDhcpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicGatewayDhcpResult>("scaleway:index/getVpcPublicGatewayDhcp:getVpcPublicGatewayDhcp", args ?? new GetVpcPublicGatewayDhcpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a public gateway DHCP.
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
        ///     var main = new Scaleway.VpcPublicGatewayDhcp("main", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var dhcpById = Scaleway.GetVpcPublicGatewayDhcp.Invoke(new()
        ///     {
        ///         DhcpId = main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcPublicGatewayDhcpResult> Invoke(GetVpcPublicGatewayDhcpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicGatewayDhcpResult>("scaleway:index/getVpcPublicGatewayDhcp:getVpcPublicGatewayDhcp", args ?? new GetVpcPublicGatewayDhcpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPublicGatewayDhcpArgs : global::Pulumi.InvokeArgs
    {
        [Input("dhcpId", required: true)]
        public string DhcpId { get; set; } = null!;

        public GetVpcPublicGatewayDhcpArgs()
        {
        }
        public static new GetVpcPublicGatewayDhcpArgs Empty => new GetVpcPublicGatewayDhcpArgs();
    }

    public sealed class GetVpcPublicGatewayDhcpInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dhcpId", required: true)]
        public Input<string> DhcpId { get; set; } = null!;

        public GetVpcPublicGatewayDhcpInvokeArgs()
        {
        }
        public static new GetVpcPublicGatewayDhcpInvokeArgs Empty => new GetVpcPublicGatewayDhcpInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPublicGatewayDhcpResult
    {
        public readonly string Address;
        public readonly string CreatedAt;
        public readonly string DhcpId;
        public readonly string DnsLocalName;
        public readonly ImmutableArray<string> DnsSearches;
        public readonly ImmutableArray<string> DnsServersOverrides;
        public readonly bool EnableDynamic;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string PoolHigh;
        public readonly string PoolLow;
        public readonly string ProjectId;
        public readonly bool PushDefaultRoute;
        public readonly bool PushDnsServer;
        public readonly int RebindTimer;
        public readonly int RenewTimer;
        public readonly string Subnet;
        public readonly string UpdatedAt;
        public readonly int ValidLifetime;
        public readonly string Zone;

        [OutputConstructor]
        private GetVpcPublicGatewayDhcpResult(
            string address,

            string createdAt,

            string dhcpId,

            string dnsLocalName,

            ImmutableArray<string> dnsSearches,

            ImmutableArray<string> dnsServersOverrides,

            bool enableDynamic,

            string id,

            string organizationId,

            string poolHigh,

            string poolLow,

            string projectId,

            bool pushDefaultRoute,

            bool pushDnsServer,

            int rebindTimer,

            int renewTimer,

            string subnet,

            string updatedAt,

            int validLifetime,

            string zone)
        {
            Address = address;
            CreatedAt = createdAt;
            DhcpId = dhcpId;
            DnsLocalName = dnsLocalName;
            DnsSearches = dnsSearches;
            DnsServersOverrides = dnsServersOverrides;
            EnableDynamic = enableDynamic;
            Id = id;
            OrganizationId = organizationId;
            PoolHigh = poolHigh;
            PoolLow = poolLow;
            ProjectId = projectId;
            PushDefaultRoute = pushDefaultRoute;
            PushDnsServer = pushDnsServer;
            RebindTimer = rebindTimer;
            RenewTimer = renewTimer;
            Subnet = subnet;
            UpdatedAt = updatedAt;
            ValidLifetime = validLifetime;
            Zone = zone;
        }
    }
}
