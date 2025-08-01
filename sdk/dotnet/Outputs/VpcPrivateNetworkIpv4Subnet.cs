// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class VpcPrivateNetworkIpv4Subnet
    {
        /// <summary>
        /// The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The date and time of the creation of the subnet.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The subnet ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        /// </summary>
        public readonly int? PrefixLength;
        /// <summary>
        /// The subnet CIDR.
        /// </summary>
        public readonly string? Subnet;
        /// <summary>
        /// The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        /// </summary>
        public readonly string? SubnetMask;
        /// <summary>
        /// The date and time of the last update of the subnet.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private VpcPrivateNetworkIpv4Subnet(
            string? address,

            string? createdAt,

            string? id,

            int? prefixLength,

            string? subnet,

            string? subnetMask,

            string? updatedAt)
        {
            Address = address;
            CreatedAt = createdAt;
            Id = id;
            PrefixLength = prefixLength;
            Subnet = subnet;
            SubnetMask = subnetMask;
            UpdatedAt = updatedAt;
        }
    }
}
