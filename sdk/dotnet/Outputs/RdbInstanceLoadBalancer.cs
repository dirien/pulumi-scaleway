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
    public sealed class RdbInstanceLoadBalancer
    {
        /// <summary>
        /// The ID of the endpoint of the private network.
        /// </summary>
        public readonly string? EndpointId;
        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// IP of the endpoint.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Port of the endpoint.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private RdbInstanceLoadBalancer(
            string? endpointId,

            string? hostname,

            string? ip,

            string? name,

            int? port)
        {
            EndpointId = endpointId;
            Hostname = hostname;
            Ip = ip;
            Name = name;
            Port = port;
        }
    }
}
