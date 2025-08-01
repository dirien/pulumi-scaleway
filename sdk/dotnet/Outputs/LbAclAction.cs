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
    public sealed class LbAclAction
    {
        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public readonly ImmutableArray<Outputs.LbAclActionRedirect> Redirects;
        /// <summary>
        /// The action type. Possible values are: `allow` or `deny` or `redirect`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private LbAclAction(
            ImmutableArray<Outputs.LbAclActionRedirect> redirects,

            string type)
        {
            Redirects = redirects;
            Type = type;
        }
    }
}
