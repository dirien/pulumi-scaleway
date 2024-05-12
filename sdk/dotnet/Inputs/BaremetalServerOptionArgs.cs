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

    public sealed class BaremetalServerOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auto expiration date for compatible options
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The id of the option to enable. Use [this endpoint](https://www.scaleway.com/en/developers/api/elastic-metal/#get-012dcc) to find the available options IDs.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BaremetalServerOptionArgs()
        {
        }
        public static new BaremetalServerOptionArgs Empty => new BaremetalServerOptionArgs();
    }
}
