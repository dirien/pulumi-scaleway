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
    public sealed class GetMongodbInstancePublicNetworkResult
    {
        /// <summary>
        /// The DNS record of your endpoint
        /// </summary>
        public readonly string DnsRecord;
        /// <summary>
        /// The ID of the MongoDB® Instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// TCP port of the endpoint
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetMongodbInstancePublicNetworkResult(
            string dnsRecord,

            string id,

            int port)
        {
            DnsRecord = dnsRecord;
            Id = id;
            Port = port;
        }
    }
}
