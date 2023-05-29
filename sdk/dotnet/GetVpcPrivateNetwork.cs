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
    public static class GetVpcPrivateNetwork
    {
        /// <summary>
        /// Gets information about a private network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// 
        /// N/A, the usage will be meaningful in the next releases of VPC.
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcPrivateNetworkResult> InvokeAsync(GetVpcPrivateNetworkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPrivateNetworkResult>("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", args ?? new GetVpcPrivateNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a private network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// 
        /// N/A, the usage will be meaningful in the next releases of VPC.
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVpcPrivateNetworkResult> Invoke(GetVpcPrivateNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPrivateNetworkResult>("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", args ?? new GetVpcPrivateNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPrivateNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the private network. One of `name` and `private_network_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// ID of the private network. One of `name` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        public GetVpcPrivateNetworkArgs()
        {
        }
        public static new GetVpcPrivateNetworkArgs Empty => new GetVpcPrivateNetworkArgs();
    }

    public sealed class GetVpcPrivateNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the private network. One of `name` and `private_network_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the private network. One of `name` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        public GetVpcPrivateNetworkInvokeArgs()
        {
        }
        public static new GetVpcPrivateNetworkInvokeArgs Empty => new GetVpcPrivateNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPrivateNetworkResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        public readonly string ProjectId;
        /// <summary>
        /// The subnets CIDR associated with private network.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string Zone;

        [OutputConstructor]
        private GetVpcPrivateNetworkResult(
            string createdAt,

            string id,

            string? name,

            string organizationId,

            string? privateNetworkId,

            string projectId,

            ImmutableArray<string> subnets,

            ImmutableArray<string> tags,

            string updatedAt,

            string zone)
        {
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Subnets = subnets;
            Tags = tags;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
