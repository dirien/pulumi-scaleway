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

    public sealed class ObjectBucketAclAccessControlPolicyGrantGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("grantee")]
        public Input<Inputs.ObjectBucketAclAccessControlPolicyGrantGranteeGetArgs>? Grantee { get; set; }

        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        public ObjectBucketAclAccessControlPolicyGrantGetArgs()
        {
        }
        public static new ObjectBucketAclAccessControlPolicyGrantGetArgs Empty => new ObjectBucketAclAccessControlPolicyGrantGetArgs();
    }
}
