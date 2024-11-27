// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class ContainerTriggerNats
    {
        /// <summary>
        /// unique identifier of the Messaging and Queuing NATS account.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// THe ID of the project that contains the Messaging and Queuing NATS account (defaults to provider `project_id`)
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// Region where the Messaging and Queuing NATS account is enabled (defaults to provider `region`)
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The subject to listen to.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private ContainerTriggerNats(
            string? accountId,

            string? projectId,

            string? region,

            string subject)
        {
            AccountId = accountId;
            ProjectId = projectId;
            Region = region;
            Subject = subject;
        }
    }
}
