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
    /// Creates and manages Scaleway VPC Public Gateway PAT (Port Address Translation).
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway/#pat-rules-e75d10).
    /// 
    /// ## Example Usage
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
    ///         InboundDefaultPolicy = "drop",
    ///         OutboundDefaultPolicy = "accept",
    ///         InboundRules = new[]
    ///         {
    ///             new Scaleway.Inputs.InstanceSecurityGroupInboundRuleArgs
    ///             {
    ///                 Action = "accept",
    ///                 Port = 22,
    ///                 Protocol = "TCP",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var srv01 = new Scaleway.InstanceServer("srv01", new()
    ///     {
    ///         Type = "PLAY2-NANO",
    ///         Image = "ubuntu_jammy",
    ///         SecurityGroupId = sg01.Id,
    ///     });
    /// 
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01");
    /// 
    ///     var pnic01 = new Scaleway.InstancePrivateNic("pnic01", new()
    ///     {
    ///         ServerId = srv01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///     });
    /// 
    ///     var dhcp01 = new Scaleway.VpcPublicGatewayDhcp("dhcp01", new()
    ///     {
    ///         Subnet = "192.168.0.0/24",
    ///     });
    /// 
    ///     var ip01 = new Scaleway.VpcPublicGatewayIp("ip01");
    /// 
    ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
    ///     {
    ///         Type = "VPC-GW-S",
    ///         IpId = ip01.Id,
    ///     });
    /// 
    ///     var gn01 = new Scaleway.VpcGatewayNetwork("gn01", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         DhcpId = dhcp01.Id,
    ///         CleanupDhcp = true,
    ///         EnableMasquerade = true,
    ///     });
    /// 
    ///     var rsv01 = new Scaleway.VpcPublicGatewayDhcpReservation("rsv01", new()
    ///     {
    ///         GatewayNetworkId = gn01.Id,
    ///         MacAddress = pnic01.MacAddress,
    ///         IpAddress = "192.168.0.7",
    ///     });
    /// 
    ///     // PAT rule for SSH traffic
    ///     var pat01 = new Scaleway.VpcPublicGatewayPatRule("pat01", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateIp = rsv01.IpAddress,
    ///         PrivatePort = 22,
    ///         PublicPort = 2202,
    ///         Protocol = "tcp",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Public gateway PAT rules config can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule")]
    public partial class VpcPublicGatewayPatRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time of the creation of the pat rule config.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// The organization ID the pat rule config is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The Private IP to forward data to (IP address).
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// The Private port to translate to.
        /// </summary>
        [Output("privatePort")]
        public Output<int> PrivatePort { get; private set; } = null!;

        /// <summary>
        /// The Protocol the rule should apply to. Possible values are both, tcp and udp.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The Public port to listen on.
        /// </summary>
        [Output("publicPort")]
        public Output<int> PublicPort { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the pat rule config.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGatewayPatRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGatewayPatRule(string name, VpcPublicGatewayPatRuleArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, args ?? new VpcPublicGatewayPatRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGatewayPatRule(string name, Input<string> id, VpcPublicGatewayPatRuleState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPublicGatewayPatRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGatewayPatRule Get(string name, Input<string> id, VpcPublicGatewayPatRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGatewayPatRule(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayPatRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// The Private IP to forward data to (IP address).
        /// </summary>
        [Input("privateIp", required: true)]
        public Input<string> PrivateIp { get; set; } = null!;

        /// <summary>
        /// The Private port to translate to.
        /// </summary>
        [Input("privatePort", required: true)]
        public Input<int> PrivatePort { get; set; } = null!;

        /// <summary>
        /// The Protocol the rule should apply to. Possible values are both, tcp and udp.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The Public port to listen on.
        /// </summary>
        [Input("publicPort", required: true)]
        public Input<int> PublicPort { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayPatRuleArgs()
        {
        }
        public static new VpcPublicGatewayPatRuleArgs Empty => new VpcPublicGatewayPatRuleArgs();
    }

    public sealed class VpcPublicGatewayPatRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the pat rule config.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// The organization ID the pat rule config is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The Private IP to forward data to (IP address).
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        /// <summary>
        /// The Private port to translate to.
        /// </summary>
        [Input("privatePort")]
        public Input<int>? PrivatePort { get; set; }

        /// <summary>
        /// The Protocol the rule should apply to. Possible values are both, tcp and udp.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The Public port to listen on.
        /// </summary>
        [Input("publicPort")]
        public Input<int>? PublicPort { get; set; }

        /// <summary>
        /// The date and time of the last update of the pat rule config.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayPatRuleState()
        {
        }
        public static new VpcPublicGatewayPatRuleState Empty => new VpcPublicGatewayPatRuleState();
    }
}
