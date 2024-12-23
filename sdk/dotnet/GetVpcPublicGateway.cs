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
    public static class GetVpcPublicGateway
    {
        /// <summary>
        /// Gets information about a Public Gateway.
        /// 
        /// ## Example Usage
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
        ///     var main = new Scaleway.VpcPublicGateway("main", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestByName = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         Name = main.Name,
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestById = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         PublicGatewayId = main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVpcPublicGatewayResult> InvokeAsync(GetVpcPublicGatewayArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPublicGatewayResult>("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", args ?? new GetVpcPublicGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Public Gateway.
        /// 
        /// ## Example Usage
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
        ///     var main = new Scaleway.VpcPublicGateway("main", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestByName = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         Name = main.Name,
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestById = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         PublicGatewayId = main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcPublicGatewayResult> Invoke(GetVpcPublicGatewayInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicGatewayResult>("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", args ?? new GetVpcPublicGatewayInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Public Gateway.
        /// 
        /// ## Example Usage
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
        ///     var main = new Scaleway.VpcPublicGateway("main", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestByName = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         Name = main.Name,
        ///         Zone = "nl-ams-1",
        ///     });
        /// 
        ///     var pgTestById = Scaleway.GetVpcPublicGateway.Invoke(new()
        ///     {
        ///         PublicGatewayId = main.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcPublicGatewayResult> Invoke(GetVpcPublicGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPublicGatewayResult>("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", args ?? new GetVpcPublicGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPublicGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Exact name of the Public Gateway.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the Project the Public Gateway is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("publicGatewayId")]
        public string? PublicGatewayId { get; set; }

        /// <summary>
        /// `zone`) The Public Gateway's zone.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetVpcPublicGatewayArgs()
        {
        }
        public static new GetVpcPublicGatewayArgs Empty => new GetVpcPublicGatewayArgs();
    }

    public sealed class GetVpcPublicGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Exact name of the Public Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Project the Public Gateway is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("publicGatewayId")]
        public Input<string>? PublicGatewayId { get; set; }

        /// <summary>
        /// `zone`) The Public Gateway's zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetVpcPublicGatewayInvokeArgs()
        {
        }
        public static new GetVpcPublicGatewayInvokeArgs Empty => new GetVpcPublicGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPublicGatewayResult
    {
        public readonly bool BastionEnabled;
        public readonly int BastionPort;
        public readonly string CreatedAt;
        public readonly bool EnableSmtp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpId;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        public readonly string? PublicGatewayId;
        public readonly string RefreshSshKeys;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string Type;
        public readonly string UpdatedAt;
        public readonly ImmutableArray<string> UpstreamDnsServers;
        public readonly string? Zone;

        [OutputConstructor]
        private GetVpcPublicGatewayResult(
            bool bastionEnabled,

            int bastionPort,

            string createdAt,

            bool enableSmtp,

            string id,

            string ipId,

            string? name,

            string organizationId,

            string? projectId,

            string? publicGatewayId,

            string refreshSshKeys,

            string status,

            ImmutableArray<string> tags,

            string type,

            string updatedAt,

            ImmutableArray<string> upstreamDnsServers,

            string? zone)
        {
            BastionEnabled = bastionEnabled;
            BastionPort = bastionPort;
            CreatedAt = createdAt;
            EnableSmtp = enableSmtp;
            Id = id;
            IpId = ipId;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            PublicGatewayId = publicGatewayId;
            RefreshSshKeys = refreshSshKeys;
            Status = status;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
            UpstreamDnsServers = upstreamDnsServers;
            Zone = zone;
        }
    }
}
