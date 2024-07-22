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
    public sealed class MnqSqsCredentialsPermissions
    {
        /// <summary>
        /// . Defines whether the user can manage the associated resource(s).
        /// </summary>
        public readonly bool? CanManage;
        /// <summary>
        /// . Defines whether the user can publish messages to the service.
        /// </summary>
        public readonly bool? CanPublish;
        /// <summary>
        /// . Defines whether the user can receive messages from the service.
        /// </summary>
        public readonly bool? CanReceive;

        [OutputConstructor]
        private MnqSqsCredentialsPermissions(
            bool? canManage,

            bool? canPublish,

            bool? canReceive)
        {
            CanManage = canManage;
            CanPublish = canPublish;
            CanReceive = canReceive;
        }
    }
}
