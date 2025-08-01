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

    public sealed class RdbInstanceLogsPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The max age (in days) of remote logs to keep on the Database Instance
        /// </summary>
        [Input("maxAgeRetention")]
        public Input<int>? MaxAgeRetention { get; set; }

        /// <summary>
        /// The max disk size of remote logs to keep on the Database Instance.
        /// </summary>
        [Input("totalDiskRetention")]
        public Input<int>? TotalDiskRetention { get; set; }

        public RdbInstanceLogsPolicyGetArgs()
        {
        }
        public static new RdbInstanceLogsPolicyGetArgs Empty => new RdbInstanceLogsPolicyGetArgs();
    }
}
