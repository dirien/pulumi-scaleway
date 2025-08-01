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

    public sealed class ObjectBucketWebsiteConfigurationIndexDocumentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A suffix that is appended to a request targeting a specific directory on the website endpoint.
        /// 
        /// &gt; **Important:** The suffix must not be empty and must not include a slash character. The routing is not supported.
        /// </summary>
        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public ObjectBucketWebsiteConfigurationIndexDocumentArgs()
        {
        }
        public static new ObjectBucketWebsiteConfigurationIndexDocumentArgs Empty => new ObjectBucketWebsiteConfigurationIndexDocumentArgs();
    }
}
