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
    /// Creates and manages Scaleway compute Instance security group rules. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#path-security-groups-list-security-groups).
    /// 
    /// This resource can be used to externalize rules from a `scaleway.InstanceSecurityGroup` to solve circular dependency problems. When using this resource do not forget to set `external_rules = true` on the security group.
    /// 
    /// &gt; **Warning:** In order to guaranty rules order in a given security group only one scaleway.InstanceSecurityGroupRules is allowed per security group.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sg01 = new Scaleway.InstanceSecurityGroup("sg01", new()
    ///     {
    ///         ExternalRules = true,
    ///     });
    /// 
    ///     var sgrs01 = new Scaleway.InstanceSecurityGroupRules("sgrs01", new()
    ///     {
    ///         SecurityGroupId = sg01.Id,
    ///         InboundRules = new[]
    ///         {
    ///             new Scaleway.Inputs.InstanceSecurityGroupRulesInboundRuleArgs
    ///             {
    ///                 Action = "accept",
    ///                 Port = 80,
    ///                 IpRange = "0.0.0.0/0",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance security group rules can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules web fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules")]
    public partial class InstanceSecurityGroupRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        [Output("inboundRules")]
        public Output<ImmutableArray<Outputs.InstanceSecurityGroupRulesInboundRule>> InboundRules { get; private set; } = null!;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        [Output("outboundRules")]
        public Output<ImmutableArray<Outputs.InstanceSecurityGroupRulesOutboundRule>> OutboundRules { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceSecurityGroupRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceSecurityGroupRules(string name, InstanceSecurityGroupRulesArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules", name, args ?? new InstanceSecurityGroupRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceSecurityGroupRules(string name, Input<string> id, InstanceSecurityGroupRulesState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceSecurityGroupRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceSecurityGroupRules Get(string name, Input<string> id, InstanceSecurityGroupRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceSecurityGroupRules(name, id, state, options);
        }
    }

    public sealed class InstanceSecurityGroupRulesArgs : global::Pulumi.ResourceArgs
    {
        [Input("inboundRules")]
        private InputList<Inputs.InstanceSecurityGroupRulesInboundRuleArgs>? _inboundRules;

        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupRulesInboundRuleArgs> InboundRules
        {
            get => _inboundRules ?? (_inboundRules = new InputList<Inputs.InstanceSecurityGroupRulesInboundRuleArgs>());
            set => _inboundRules = value;
        }

        [Input("outboundRules")]
        private InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleArgs>? _outboundRules;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleArgs>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public InstanceSecurityGroupRulesArgs()
        {
        }
        public static new InstanceSecurityGroupRulesArgs Empty => new InstanceSecurityGroupRulesArgs();
    }

    public sealed class InstanceSecurityGroupRulesState : global::Pulumi.ResourceArgs
    {
        [Input("inboundRules")]
        private InputList<Inputs.InstanceSecurityGroupRulesInboundRuleGetArgs>? _inboundRules;

        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupRulesInboundRuleGetArgs> InboundRules
        {
            get => _inboundRules ?? (_inboundRules = new InputList<Inputs.InstanceSecurityGroupRulesInboundRuleGetArgs>());
            set => _inboundRules = value;
        }

        [Input("outboundRules")]
        private InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleGetArgs>? _outboundRules;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleGetArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<Inputs.InstanceSecurityGroupRulesOutboundRuleGetArgs>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public InstanceSecurityGroupRulesState()
        {
        }
        public static new InstanceSecurityGroupRulesState Empty => new InstanceSecurityGroupRulesState();
    }
}
