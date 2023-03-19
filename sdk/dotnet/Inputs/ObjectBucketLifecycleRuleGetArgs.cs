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

    public sealed class ObjectBucketLifecycleRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
        /// </summary>
        [Input("abortIncompleteMultipartUploadDays")]
        public Input<int>? AbortIncompleteMultipartUploadDays { get; set; }

        /// <summary>
        /// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        [Input("expiration")]
        public Input<Inputs.ObjectBucketLifecycleRuleExpirationGetArgs>? Expiration { get; set; }

        /// <summary>
        /// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Specifies object tags key and value.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("transitions")]
        private InputList<Inputs.ObjectBucketLifecycleRuleTransitionGetArgs>? _transitions;

        /// <summary>
        /// Specifies a period in the object's transitions (documented below).
        /// </summary>
        public InputList<Inputs.ObjectBucketLifecycleRuleTransitionGetArgs> Transitions
        {
            get => _transitions ?? (_transitions = new InputList<Inputs.ObjectBucketLifecycleRuleTransitionGetArgs>());
            set => _transitions = value;
        }

        public ObjectBucketLifecycleRuleGetArgs()
        {
        }
        public static new ObjectBucketLifecycleRuleGetArgs Empty => new ObjectBucketLifecycleRuleGetArgs();
    }
}
