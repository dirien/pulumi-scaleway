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
    public sealed class GetRdbInstanceReadReplicaResult
    {
        /// <summary>
        /// IP of the replica
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The name of the RDB instance.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Port of the replica
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetRdbInstanceReadReplicaResult(
            string ip,

            string name,

            int port)
        {
            Ip = ip;
            Name = name;
            Port = port;
        }
    }
}
