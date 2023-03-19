// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetInstanceServersServerResult
    {
        /// <summary>
        /// The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
        /// </summary>
        public readonly string BootType;
        /// <summary>
        /// The ID of the bootscript.
        /// </summary>
        public readonly string BootscriptId;
        /// <summary>
        /// If true a dynamic IP will be attached to the server.
        /// </summary>
        public readonly bool EnableDynamicIp;
        /// <summary>
        /// Determines if IPv6 is enabled for the server.
        /// </summary>
        public readonly bool EnableIpv6;
        /// <summary>
        /// The ID of the server.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The UUID or the label of the base image used by the server.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly string Ipv6Gateway;
        /// <summary>
        /// The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly int Ipv6PrefixLength;
        /// <summary>
        /// The server name used as filter. Servers with a name like it are listed.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The organization ID the server is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        /// </summary>
        public readonly string PlacementGroupId;
        public readonly bool PlacementGroupPolicyRespected;
        /// <summary>
        /// The Scaleway internal IP address of the server.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The ID of the project the server is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The public IPv4 address of the server.
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The state of the server. Possible values are: `started`, `stopped` or `standby`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// List of tags used as filter. Servers with these exact tags are listed.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The commercial type of the server.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// `zone`) The zone in which servers exist.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceServersServerResult(
            string bootType,

            string bootscriptId,

            bool enableDynamicIp,

            bool enableIpv6,

            string id,

            string image,

            string ipv6Address,

            string ipv6Gateway,

            int ipv6PrefixLength,

            string name,

            string organizationId,

            string placementGroupId,

            bool placementGroupPolicyRespected,

            string privateIp,

            string projectId,

            string publicIp,

            string securityGroupId,

            string state,

            ImmutableArray<string> tags,

            string type,

            string zone)
        {
            BootType = bootType;
            BootscriptId = bootscriptId;
            EnableDynamicIp = enableDynamicIp;
            EnableIpv6 = enableIpv6;
            Id = id;
            Image = image;
            Ipv6Address = ipv6Address;
            Ipv6Gateway = ipv6Gateway;
            Ipv6PrefixLength = ipv6PrefixLength;
            Name = name;
            OrganizationId = organizationId;
            PlacementGroupId = placementGroupId;
            PlacementGroupPolicyRespected = placementGroupPolicyRespected;
            PrivateIp = privateIp;
            ProjectId = projectId;
            PublicIp = publicIp;
            SecurityGroupId = securityGroupId;
            State = state;
            Tags = tags;
            Type = type;
            Zone = zone;
        }
    }
}
