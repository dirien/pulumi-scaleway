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

    public sealed class VpcGatewayNetworkIpamConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the default route is enabled on that Gateway Network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Input("pushDefaultRoute")]
        public Input<bool>? PushDefaultRoute { get; set; }

        public VpcGatewayNetworkIpamConfigGetArgs()
        {
        }
        public static new VpcGatewayNetworkIpamConfigGetArgs Empty => new VpcGatewayNetworkIpamConfigGetArgs();
    }
}
