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

    public sealed class LbPrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Deprecated) Please use `ipam_ids`. Set to `true` if you want to let DHCP assign IP addresses. See below.
        /// </summary>
        [Input("dhcpConfig")]
        public Input<bool>? DhcpConfig { get; set; }

        /// <summary>
        /// (Optional) IPAM ID of a pre-reserved IP address to assign to the Load Balancer on this Private Network.
        /// </summary>
        [Input("ipamIds")]
        public Input<string>? IpamIds { get; set; }

        /// <summary>
        /// (Required) The ID of the Private Network to attach to.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        /// <summary>
        /// (Deprecated) Please use `ipam_ids`. Define a local ip address of your choice for the load balancer instance.
        /// </summary>
        [Input("staticConfig")]
        public Input<string>? StaticConfig { get; set; }

        /// <summary>
        /// The status of private network connection
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// `zone`) The zone of the Load Balancer.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LbPrivateNetworkGetArgs()
        {
        }
        public static new LbPrivateNetworkGetArgs Empty => new LbPrivateNetworkGetArgs();
    }
}
