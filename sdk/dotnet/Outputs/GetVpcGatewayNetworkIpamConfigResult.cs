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
    public sealed class GetVpcGatewayNetworkIpamConfigResult
    {
        public readonly string IpamIpId;
        public readonly bool PushDefaultRoute;

        [OutputConstructor]
        private GetVpcGatewayNetworkIpamConfigResult(
            string ipamIpId,

            bool pushDefaultRoute)
        {
            IpamIpId = ipamIpId;
            PushDefaultRoute = pushDefaultRoute;
        }
    }
}
