// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class IpamIpSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Private Network of the IP (if the IP is a private IP).
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The Private Network subnet of the IP (if the IP is a private IP).
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The zone of the IP (if the IP is public and zoned, rather than private and/or regional)
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public IpamIpSourceGetArgs()
        {
        }
        public static new IpamIpSourceGetArgs Empty => new IpamIpSourceGetArgs();
    }
}
