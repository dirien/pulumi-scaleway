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

    public sealed class ObjectBucketLockConfigurationRuleDefaultRetentionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days you want to specify for the default retention period.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// The default object lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are `GOVERNANCE` or `COMPLIANCE`. Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/#retention-modes) for more information on retention modes.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// The number of years you want to specify for the default retention period.
        /// </summary>
        [Input("years")]
        public Input<int>? Years { get; set; }

        public ObjectBucketLockConfigurationRuleDefaultRetentionGetArgs()
        {
        }
        public static new ObjectBucketLockConfigurationRuleDefaultRetentionGetArgs Empty => new ObjectBucketLockConfigurationRuleDefaultRetentionGetArgs();
    }
}
