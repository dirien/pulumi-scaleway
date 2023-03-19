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

    public sealed class BaremetalServerPrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the private network.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The id of the private network to attach.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The private network status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time of the last update of the private network.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The VLAN ID associated to the private network.
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public BaremetalServerPrivateNetworkArgs()
        {
        }
        public static new BaremetalServerPrivateNetworkArgs Empty => new BaremetalServerPrivateNetworkArgs();
    }
}
