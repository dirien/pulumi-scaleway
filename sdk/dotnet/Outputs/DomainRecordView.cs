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
    public sealed class DomainRecordView
    {
        /// <summary>
        /// The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The subnet of the view
        /// </summary>
        public readonly string Subnet;

        [OutputConstructor]
        private DomainRecordView(
            string data,

            string subnet)
        {
            Data = data;
            Subnet = subnet;
        }
    }
}
