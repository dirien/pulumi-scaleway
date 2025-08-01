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
    public sealed class EdgeServicesCacheStagePurgeRequest
    {
        /// <summary>
        /// Defines whether to purge all content.
        /// </summary>
        public readonly bool? All;
        /// <summary>
        /// The list of asserts to purge.
        /// </summary>
        public readonly ImmutableArray<string> Assets;
        /// <summary>
        /// The pipeline ID in which the purge request will be created.
        /// </summary>
        public readonly string? PipelineId;

        [OutputConstructor]
        private EdgeServicesCacheStagePurgeRequest(
            bool? all,

            ImmutableArray<string> assets,

            string? pipelineId)
        {
            All = all;
            Assets = assets;
            PipelineId = pipelineId;
        }
    }
}
