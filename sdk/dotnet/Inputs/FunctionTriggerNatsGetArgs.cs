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

    public sealed class FunctionTriggerNatsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// unique identifier of the Messaging and Queuing NATS account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// THe ID of the project that contains the Messaging and Queuing NATS account (defaults to provider `project_id`)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Region where the Messaging and Queuing NATS account is enabled (defaults to provider `region`)
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The subject to listen to.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        public FunctionTriggerNatsGetArgs()
        {
        }
        public static new FunctionTriggerNatsGetArgs Empty => new FunctionTriggerNatsGetArgs();
    }
}
