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
    public sealed class IpamIpCustomResource
    {
        /// <summary>
        /// The MAC address of the resource the IP is attached to.
        /// </summary>
        public readonly string MacAddress;
        /// <summary>
        /// The name of the resource the IP is attached to.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private IpamIpCustomResource(
            string macAddress,

            string? name)
        {
            MacAddress = macAddress;
            Name = name;
        }
    }
}
