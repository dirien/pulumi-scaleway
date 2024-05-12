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

    public sealed class RdbReadReplicaPrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, the IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// 
        /// &gt; **Important:** One of `service_ip` or `enable_ipam=true` must be set.
        /// </summary>
        [Input("enableIpam")]
        public Input<bool>? EnableIpam { get; set; }

        /// <summary>
        /// The ID of the endpoint of the read replica.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Hostname of the endpoint. Only one of ip and hostname may be set.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// TCP port of the endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// UUID of the private network to be connected to the read replica.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a CIDR notation. If not set, The IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// </summary>
        [Input("serviceIp")]
        public Input<string>? ServiceIp { get; set; }

        /// <summary>
        /// Private network zone
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public RdbReadReplicaPrivateNetworkArgs()
        {
        }
        public static new RdbReadReplicaPrivateNetworkArgs Empty => new RdbReadReplicaPrivateNetworkArgs();
    }
}
