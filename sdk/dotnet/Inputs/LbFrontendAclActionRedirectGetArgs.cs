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

    public sealed class LbFrontendAclActionRedirectGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP redirect code to use. Valid values are `301`, `302`, `303`, `307` and `308`.
        /// </summary>
        [Input("code")]
        public Input<int>? Code { get; set; }

        /// <summary>
        /// A URL can be used in case of a location redirect (e.g. `https://scaleway.com` will redirect to this same URL). A scheme name (e.g. `https`, `http`, `ftp`, `git`) will replace the request's original scheme.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The redirect type. Possible values are: `location` or `scheme`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LbFrontendAclActionRedirectGetArgs()
        {
        }
        public static new LbFrontendAclActionRedirectGetArgs Empty => new LbFrontendAclActionRedirectGetArgs();
    }
}
