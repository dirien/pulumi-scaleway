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
    /// Creates and manages Scaleway Instance Private NICs. For more information, see
    /// [the documentation](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea).
    /// 
    /// ## Examples
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
    ///     var pnic01 = new Scaleway.InstancePrivateNic("pnic01", new()
    ///     {
    ///         PrivateNetworkId = "fr-par-1/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
    ///         ServerId = "fr-par-1/11111111-1111-1111-1111-111111111111",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01", new()
    ///     {
    ///         Zone = "fr-par-2",
    ///     });
    /// 
    ///     var @base = new Scaleway.InstanceServer("base", new()
    ///     {
    ///         Image = "ubuntu_jammy",
    ///         Type = "DEV1-S",
    ///         Zone = pn01.Zone,
    ///     });
    /// 
    ///     var pnic01 = new Scaleway.InstancePrivateNic("pnic01", new()
    ///     {
    ///         ServerId = @base.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         Zone = pn01.Zone,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Private NICs can be imported using the `{zone}/{server_id}/{private_nic_id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/instancePrivateNic:InstancePrivateNic pnic01 fr-par-1/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instancePrivateNic:InstancePrivateNic")]
    public partial class InstancePrivateNic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IPAM ip list, should be for internal use only
        /// </summary>
        [Output("ipIds")]
        public Output<ImmutableArray<string>> IpIds { get; private set; } = null!;

        /// <summary>
        /// The MAC address of the private NIC.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the private network attached to.
        /// </summary>
        [Output("privateNetworkId")]
        public Output<string> PrivateNetworkId { get; private set; } = null!;

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the private NIC.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the server must be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstancePrivateNic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstancePrivateNic(string name, InstancePrivateNicArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/instancePrivateNic:InstancePrivateNic", name, args ?? new InstancePrivateNicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstancePrivateNic(string name, Input<string> id, InstancePrivateNicState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instancePrivateNic:InstancePrivateNic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstancePrivateNic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstancePrivateNic Get(string name, Input<string> id, InstancePrivateNicState? state = null, CustomResourceOptions? options = null)
        {
            return new InstancePrivateNic(name, id, state, options);
        }
    }

    public sealed class InstancePrivateNicArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// IPAM ip list, should be for internal use only
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The ID of the private network attached to.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the server must be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstancePrivateNicArgs()
        {
        }
        public static new InstancePrivateNicArgs Empty => new InstancePrivateNicArgs();
    }

    public sealed class InstancePrivateNicState : global::Pulumi.ResourceArgs
    {
        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// IPAM ip list, should be for internal use only
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The MAC address of the private NIC.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network attached to.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the server must be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstancePrivateNicState()
        {
        }
        public static new InstancePrivateNicState Empty => new InstancePrivateNicState();
    }
}
