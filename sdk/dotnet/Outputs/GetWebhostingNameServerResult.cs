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
    public sealed class GetWebhostingNameServerResult
    {
        public readonly string Hostname;
        public readonly bool IsDefault;
        public readonly string Status;

        [OutputConstructor]
        private GetWebhostingNameServerResult(
            string hostname,

            bool isDefault,

            string status)
        {
            Hostname = hostname;
            IsDefault = isDefault;
            Status = status;
        }
    }
}
