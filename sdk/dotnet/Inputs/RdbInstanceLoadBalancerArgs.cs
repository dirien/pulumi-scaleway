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

    public sealed class RdbInstanceLoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public RdbInstanceLoadBalancerArgs()
        {
        }
        public static new RdbInstanceLoadBalancerArgs Empty => new RdbInstanceLoadBalancerArgs();
    }
}
