// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class DomainRegistrationDsRecordPublicKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public key value.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public DomainRegistrationDsRecordPublicKeyArgs()
        {
        }
        public static new DomainRegistrationDsRecordPublicKeyArgs Empty => new DomainRegistrationDsRecordPublicKeyArgs();
    }
}
