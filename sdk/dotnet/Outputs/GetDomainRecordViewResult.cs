// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class GetDomainRecordViewResult
    {
        /// <summary>
        /// The content of the record (e.g., an IPv4 address for an `A` record or a string for a `TXT` record). Cannot be used with `record_id`.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The subnet of the view
        /// </summary>
        public readonly string Subnet;

        [OutputConstructor]
        private GetDomainRecordViewResult(
            string data,

            string subnet)
        {
            Data = data;
            Subnet = subnet;
        }
    }
}
