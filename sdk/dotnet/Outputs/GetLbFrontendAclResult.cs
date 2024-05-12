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
    public sealed class GetLbFrontendAclResult
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbFrontendAclActionResult> Actions;
        /// <summary>
        /// Date and time of ACL's creation (RFC 3339 format)
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Description of the ACL
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ACL match rule
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbFrontendAclMatchResult> Matches;
        /// <summary>
        /// The name of the frontend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Date and time of ACL's update (RFC 3339 format)
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetLbFrontendAclResult(
            ImmutableArray<Outputs.GetLbFrontendAclActionResult> actions,

            string createdAt,

            string description,

            ImmutableArray<Outputs.GetLbFrontendAclMatchResult> matches,

            string name,

            string updatedAt)
        {
            Actions = actions;
            CreatedAt = createdAt;
            Description = description;
            Matches = matches;
            Name = name;
            UpdatedAt = updatedAt;
        }
    }
}
