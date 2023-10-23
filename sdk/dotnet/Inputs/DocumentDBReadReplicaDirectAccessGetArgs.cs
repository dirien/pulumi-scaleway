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

    public sealed class DocumentDBReadReplicaDirectAccessGetArgs : global::Pulumi.ResourceArgs
    {
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

        public DocumentDBReadReplicaDirectAccessGetArgs()
        {
        }
        public static new DocumentDBReadReplicaDirectAccessGetArgs Empty => new DocumentDBReadReplicaDirectAccessGetArgs();
    }
}
