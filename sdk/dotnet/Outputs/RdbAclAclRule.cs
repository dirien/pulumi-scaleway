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
    public sealed class RdbAclAclRule
    {
        /// <summary>
        /// A text describing this rule. Default description: `IP allowed`
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The IP range to whitelist in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private RdbAclAclRule(
            string? description,

            string ip)
        {
            Description = description;
            Ip = ip;
        }
    }
}
