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
    public sealed class LbPrivateNetwork
    {
        /// <summary>
        /// (Deprecated) Please use `ipam_ids`. Set to `true` if you want to let DHCP assign IP addresses. See below.
        /// </summary>
        public readonly bool? DhcpConfig;
        /// <summary>
        /// (Optional) IPAM ID of a pre-reserved IP address to assign to the Load Balancer on this Private Network.
        /// </summary>
        public readonly string? IpamIds;
        /// <summary>
        /// (Required) The ID of the Private Network to attach to.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// (Deprecated) Please use `ipam_ids`. Define a local ip address of your choice for the load balancer instance.
        /// </summary>
        public readonly string? StaticConfig;
        /// <summary>
        /// The status of private network connection
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// `zone`) The zone of the Load Balancer.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private LbPrivateNetwork(
            bool? dhcpConfig,

            string? ipamIds,

            string privateNetworkId,

            string? staticConfig,

            string? status,

            string? zone)
        {
            DhcpConfig = dhcpConfig;
            IpamIds = ipamIds;
            PrivateNetworkId = privateNetworkId;
            StaticConfig = staticConfig;
            Status = status;
            Zone = zone;
        }
    }
}
