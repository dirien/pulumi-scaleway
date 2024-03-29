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
    public sealed class RdbInstancePrivateNetwork
    {
        /// <summary>
        /// Whether the endpoint should be configured with IPAM. Defaults to `false` if `ip_net` is defined, `true` otherwise.
        /// </summary>
        public readonly bool? EnableIpam;
        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        public readonly string? Ip;
        public readonly string? IpNet;
        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the private network.
        /// </summary>
        public readonly string PnId;
        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        public readonly int? Port;
        public readonly string? Zone;

        [OutputConstructor]
        private RdbInstancePrivateNetwork(
            bool? enableIpam,

            string? endpointId,

            string? hostname,

            string? ip,

            string? ipNet,

            string? name,

            string pnId,

            int? port,

            string? zone)
        {
            EnableIpam = enableIpam;
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            IpNet = ipNet;
            Name = name;
            PnId = pnId;
            Port = port;
            Zone = zone;
        }
    }
}
