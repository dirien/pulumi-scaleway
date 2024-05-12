// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway
{
    /// <summary>
    /// ## Import
    /// 
    /// Cockpits can be imported using the `{project_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/cockpit:Cockpit main 11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/cockpit:Cockpit")]
    public partial class Cockpit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoints.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.CockpitEndpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// Name or ID of the plan to use.
        /// </summary>
        [Output("plan")]
        public Output<string?> Plan { get; private set; } = null!;

        /// <summary>
        /// The ID of the current plan.
        /// </summary>
        [Output("planId")]
        public Output<string> PlanId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a Cockpit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cockpit(string name, CockpitArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpit:Cockpit", name, args ?? new CockpitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cockpit(string name, Input<string> id, CockpitState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpit:Cockpit", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-scaleway",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cockpit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cockpit Get(string name, Input<string> id, CockpitState? state = null, CustomResourceOptions? options = null)
        {
            return new Cockpit(name, id, state, options);
        }
    }

    public sealed class CockpitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name or ID of the plan to use.
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public CockpitArgs()
        {
        }
        public static new CockpitArgs Empty => new CockpitArgs();
    }

    public sealed class CockpitState : global::Pulumi.ResourceArgs
    {
        [Input("endpoints")]
        private InputList<Inputs.CockpitEndpointGetArgs>? _endpoints;

        /// <summary>
        /// Endpoints.
        /// </summary>
        public InputList<Inputs.CockpitEndpointGetArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.CockpitEndpointGetArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Name or ID of the plan to use.
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// The ID of the current plan.
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public CockpitState()
        {
        }
        public static new CockpitState Empty => new CockpitState();
    }
}
