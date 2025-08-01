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

    public sealed class LbAclActionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("redirects")]
        private InputList<Inputs.LbAclActionRedirectGetArgs>? _redirects;

        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public InputList<Inputs.LbAclActionRedirectGetArgs> Redirects
        {
            get => _redirects ?? (_redirects = new InputList<Inputs.LbAclActionRedirectGetArgs>());
            set => _redirects = value;
        }

        /// <summary>
        /// The action type. Possible values are: `allow` or `deny` or `redirect`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LbAclActionGetArgs()
        {
        }
        public static new LbAclActionGetArgs Empty => new LbAclActionGetArgs();
    }
}
