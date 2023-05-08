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

    public sealed class RdbInstancePrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint of the private network.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP of the endpoint.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        /// <summary>
        /// Port of the endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public RdbInstancePrivateNetworkGetArgs()
        {
        }
        public static new RdbInstancePrivateNetworkGetArgs Empty => new RdbInstancePrivateNetworkGetArgs();
    }
}
