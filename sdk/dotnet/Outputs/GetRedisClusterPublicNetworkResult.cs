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
    public sealed class GetRedisClusterPublicNetworkResult
    {
        /// <summary>
        /// The ID of the Redis cluster.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ips;
        /// <summary>
        /// TCP port of the endpoint
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetRedisClusterPublicNetworkResult(
            string id,

            ImmutableArray<string> ips,

            int port)
        {
            Id = id;
            Ips = ips;
            Port = port;
        }
    }
}
