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
    public sealed class EdgeServicesBackendStageS3BackendConfig
    {
        /// <summary>
        /// The name of the Bucket.
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// The region of the Bucket.
        /// </summary>
        public readonly string? BucketRegion;
        /// <summary>
        /// Defines whether the bucket website feature is enabled.
        /// </summary>
        public readonly bool? IsWebsite;

        [OutputConstructor]
        private EdgeServicesBackendStageS3BackendConfig(
            string? bucketName,

            string? bucketRegion,

            bool? isWebsite)
        {
            BucketName = bucketName;
            BucketRegion = bucketRegion;
            IsWebsite = isWebsite;
        }
    }
}
