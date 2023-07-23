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

    public sealed class WebHostingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The option ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The option name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WebHostingOptionArgs()
        {
        }
        public static new WebHostingOptionArgs Empty => new WebHostingOptionArgs();
    }
}
