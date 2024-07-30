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

    public sealed class IamPolicyRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of organization scoped to the rule, this can be used to create a rule for all projects in an organization.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("permissionSetNames", required: true)]
        private InputList<string>? _permissionSetNames;

        /// <summary>
        /// Names of permission sets bind to the rule.
        /// 
        /// **_TIP:_** You can use the Scaleway CLI to list the permissions details. e.g:
        /// 
        /// ```shell
        /// scw IAM permission-set list
        /// ```
        /// </summary>
        public InputList<string> PermissionSetNames
        {
            get => _permissionSetNames ?? (_permissionSetNames = new InputList<string>());
            set => _permissionSetNames = value;
        }

        [Input("projectIds")]
        private InputList<string>? _projectIds;

        /// <summary>
        /// List of project IDs scoped to the rule.
        /// 
        /// &gt; **Important** One `organization_id` or `project_ids` must be set per rule.
        /// </summary>
        public InputList<string> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<string>());
            set => _projectIds = value;
        }

        public IamPolicyRuleGetArgs()
        {
        }
        public static new IamPolicyRuleGetArgs Empty => new IamPolicyRuleGetArgs();
    }
}
