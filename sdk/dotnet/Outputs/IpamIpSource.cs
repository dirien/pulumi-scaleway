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
    public sealed class IpamIpSource
    {
        /// <summary>
        /// The Private Network of the IP (if the IP is a private IP).
        /// </summary>
        public readonly string? PrivateNetworkId;
        /// <summary>
        /// The Private Network subnet of the IP (if the IP is a private IP).
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The zone of the IP (if the IP is public and zoned, rather than private and/or regional)
        /// </summary>
        public readonly string? Zonal;

        [OutputConstructor]
        private IpamIpSource(
            string? privateNetworkId,

            string? subnetId,

            string? zonal)
        {
            PrivateNetworkId = privateNetworkId;
            SubnetId = subnetId;
            Zonal = zonal;
        }
    }
}
