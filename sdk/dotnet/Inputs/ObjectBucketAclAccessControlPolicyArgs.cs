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

    public sealed class ObjectBucketAclAccessControlPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("grants")]
        private InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantArgs>? _grants;
        public InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantArgs> Grants
        {
            get => _grants ?? (_grants = new InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantArgs>());
            set => _grants = value;
        }

        /// <summary>
        /// Configuration block of the bucket project owner's display organization ID.
        /// </summary>
        [Input("owner", required: true)]
        public Input<Inputs.ObjectBucketAclAccessControlPolicyOwnerArgs> Owner { get; set; } = null!;

        public ObjectBucketAclAccessControlPolicyArgs()
        {
        }
        public static new ObjectBucketAclAccessControlPolicyArgs Empty => new ObjectBucketAclAccessControlPolicyArgs();
    }
}
