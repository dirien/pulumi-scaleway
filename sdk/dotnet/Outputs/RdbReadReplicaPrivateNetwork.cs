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
    public sealed class RdbReadReplicaPrivateNetwork
    {
        /// <summary>
        /// Whether or not the private network endpoint should be configured with IPAM
        /// </summary>
        public readonly bool? EnableIpam;
        /// <summary>
        /// The ID of the endpoint of the read replica.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// Hostname of the endpoint. Only one of ip and hostname may be set.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// TCP port of the endpoint.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// UUID of the private network to be connected to the read replica.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a
        /// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        /// service if not set.
        /// </summary>
        public readonly string? ServiceIp;
        /// <summary>
        /// Private network zone
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private RdbReadReplicaPrivateNetwork(
            bool? enableIpam,

            string? endpointId,

            string? hostname,

            string? ip,

            string? name,

            int? port,

            string privateNetworkId,

            string? serviceIp,

            string? zone)
        {
            EnableIpam = enableIpam;
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            Name = name;
            Port = port;
            PrivateNetworkId = privateNetworkId;
            ServiceIp = serviceIp;
            Zone = zone;
        }
    }
}
