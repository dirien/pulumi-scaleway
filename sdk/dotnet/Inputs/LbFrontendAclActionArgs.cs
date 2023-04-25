// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class LbFrontendAclActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("redirects")]
        private InputList<Inputs.LbFrontendAclActionRedirectArgs>? _redirects;

        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public InputList<Inputs.LbFrontendAclActionRedirectArgs> Redirects
        {
            get => _redirects ?? (_redirects = new InputList<Inputs.LbFrontendAclActionRedirectArgs>());
            set => _redirects = value;
        }

        /// <summary>
        /// The redirect type. Possible values are: `location` or `scheme`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LbFrontendAclActionArgs()
        {
        }
        public static new LbFrontendAclActionArgs Empty => new LbFrontendAclActionArgs();
    }
}
