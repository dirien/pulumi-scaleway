// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class EdgeServicesRouteStageRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the backend stage that requests matching the rule should be forwarded to.
        /// </summary>
        [Input("backendStageId", required: true)]
        public Input<string> BackendStageId { get; set; } = null!;

        /// <summary>
        /// The rule condition to be matched. Requests matching the condition defined here will be directly forwarded to the backend specified by the `backend_stage_id` field. Requests that do not match will be checked by the next rule's condition.
        /// </summary>
        [Input("ruleHttpMatch")]
        public Input<Inputs.EdgeServicesRouteStageRuleRuleHttpMatchArgs>? RuleHttpMatch { get; set; }

        public EdgeServicesRouteStageRuleArgs()
        {
        }
        public static new EdgeServicesRouteStageRuleArgs Empty => new EdgeServicesRouteStageRuleArgs();
    }
}
