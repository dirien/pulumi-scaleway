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

    public sealed class FunctionTriggerSqsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the mnq namespace. Deprecated.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// ID of the project that contain the mnq nats account, defaults to provider's project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Name of the queue
        /// </summary>
        [Input("queue", required: true)]
        public Input<string> Queue { get; set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public FunctionTriggerSqsGetArgs()
        {
        }
        public static new FunctionTriggerSqsGetArgs Empty => new FunctionTriggerSqsGetArgs();
    }
}
