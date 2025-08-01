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

    public sealed class DomainRegistrationTechnicalContactExtensionFrCodeAuthAfnicInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AFNIC authorization code for the contact (specific to French domains).
        /// </summary>
        [Input("codeAuthAfnic")]
        public Input<string>? CodeAuthAfnic { get; set; }

        public DomainRegistrationTechnicalContactExtensionFrCodeAuthAfnicInfoGetArgs()
        {
        }
        public static new DomainRegistrationTechnicalContactExtensionFrCodeAuthAfnicInfoGetArgs Empty => new DomainRegistrationTechnicalContactExtensionFrCodeAuthAfnicInfoGetArgs();
    }
}
