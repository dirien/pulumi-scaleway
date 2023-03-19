// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class RedisClusterPrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The UUID of the private network resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("serviceIps", required: true)]
        private InputList<string>? _serviceIps;

        /// <summary>
        /// Endpoint IPv4 addresses
        /// in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at
        /// least one IP per node.
        /// </summary>
        public InputList<string> ServiceIps
        {
            get => _serviceIps ?? (_serviceIps = new InputList<string>());
            set => _serviceIps = value;
        }

        /// <summary>
        /// `zone`) The zone in which the
        /// Redis Cluster should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public RedisClusterPrivateNetworkArgs()
        {
        }
        public static new RedisClusterPrivateNetworkArgs Empty => new RedisClusterPrivateNetworkArgs();
    }
}
