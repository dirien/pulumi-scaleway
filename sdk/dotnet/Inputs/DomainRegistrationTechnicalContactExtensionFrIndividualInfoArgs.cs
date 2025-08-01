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

    public sealed class DomainRegistrationTechnicalContactExtensionFrIndividualInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the individual contact has opted into WHOIS publishing.
        /// </summary>
        [Input("whoisOptIn")]
        public Input<bool>? WhoisOptIn { get; set; }

        public DomainRegistrationTechnicalContactExtensionFrIndividualInfoArgs()
        {
        }
        public static new DomainRegistrationTechnicalContactExtensionFrIndividualInfoArgs Empty => new DomainRegistrationTechnicalContactExtensionFrIndividualInfoArgs();
    }
}
