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
    public sealed class InferenceModelNodesSupport
    {
        /// <summary>
        /// The type of node supported.
        /// </summary>
        public readonly string? NodeTypeName;
        /// <summary>
        /// A list of supported quantization options, including:
        /// </summary>
        public readonly ImmutableArray<Outputs.InferenceModelNodesSupportQuantization> Quantizations;

        [OutputConstructor]
        private InferenceModelNodesSupport(
            string? nodeTypeName,

            ImmutableArray<Outputs.InferenceModelNodesSupportQuantization> quantizations)
        {
            NodeTypeName = nodeTypeName;
            Quantizations = quantizations;
        }
    }
}
