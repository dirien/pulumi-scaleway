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

    public sealed class ObjectBucketAclAccessControlPolicyOwnerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project ID of the grantee.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The `region`,`bucket` and `acl` separated by (`/`).
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public ObjectBucketAclAccessControlPolicyOwnerArgs()
        {
        }
        public static new ObjectBucketAclAccessControlPolicyOwnerArgs Empty => new ObjectBucketAclAccessControlPolicyOwnerArgs();
    }
}
