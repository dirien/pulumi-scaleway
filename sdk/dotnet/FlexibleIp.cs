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
    /// Creates and manages Scaleway flexible IPs.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).
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
    ///     var main = new Scaleway.FlexibleIp("main", new()
    ///     {
    ///         Reverse = "my-reverse.com",
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
    ///     var main = new Scaleway.FlexibleIp("main", new()
    ///     {
    ///         Zone = "fr-par-2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With IPv6
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.FlexibleIp("main", new()
    ///     {
    ///         IsIpv6 = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With baremetal server
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainAccountSshKey = new Scaleway.AccountSshKey("mainAccountSshKey", new()
    ///     {
    ///         PublicKey = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILHy/M5FVm5ydLGcal3e5LNcfTalbeN7QL/ZGCvDEdqJ foobar@example.com",
    ///     });
    /// 
    ///     var byId = Scaleway.GetBaremetalOs.Invoke(new()
    ///     {
    ///         Zone = "fr-par-2",
    ///         Name = "Ubuntu",
    ///         Version = "20.04 LTS (Focal Fossa)",
    ///     });
    /// 
    ///     var myOffer = Scaleway.GetBaremetalOffer.Invoke(new()
    ///     {
    ///         Zone = "fr-par-2",
    ///         Name = "EM-A210R-HDD",
    ///     });
    /// 
    ///     var @base = new Scaleway.BaremetalServer("base", new()
    ///     {
    ///         Zone = "fr-par-2",
    ///         Offer = myOffer.Apply(getBaremetalOfferResult =&gt; getBaremetalOfferResult.OfferId),
    ///         Os = byId.Apply(getBaremetalOsResult =&gt; getBaremetalOsResult.OsId),
    ///         SshKeyIds = mainAccountSshKey.Id,
    ///     });
    /// 
    ///     var mainFlexibleIp = new Scaleway.FlexibleIp("mainFlexibleIp", new()
    ///     {
    ///         ServerId = @base.Id,
    ///         Zone = "fr-par-2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Flexible IPs can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/flexibleIp:FlexibleIp main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/flexibleIp:FlexibleIp")]
    public partial class FlexibleIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time of the creation of the Flexible IP (Format ISO 8601).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A description of the flexible IP.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IP address of the Flexible IP.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Defines whether the flexible IP has an IPv6 address.
        /// </summary>
        [Output("isIpv6")]
        public Output<bool?> IsIpv6 { get; private set; } = null!;

        /// <summary>
        /// The organization of the Flexible IP.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The project of the Flexible IP.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The reverse domain associated with this flexible IP.
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated server.
        /// </summary>
        [Output("serverId")]
        public Output<string?> ServerId { get; private set; } = null!;

        /// <summary>
        /// The status of the flexible IP.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the flexible IP.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the Flexible IP (Format ISO 8601).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The zone of the Flexible IP.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a FlexibleIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlexibleIp(string name, FlexibleIpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/flexibleIp:FlexibleIp", name, args ?? new FlexibleIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlexibleIp(string name, Input<string> id, FlexibleIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/flexibleIp:FlexibleIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FlexibleIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlexibleIp Get(string name, Input<string> id, FlexibleIpState? state = null, CustomResourceOptions? options = null)
        {
            return new FlexibleIp(name, id, state, options);
        }
    }

    public sealed class FlexibleIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the flexible IP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Defines whether the flexible IP has an IPv6 address.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// The project of the Flexible IP.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse domain associated with this flexible IP.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// The ID of the associated server.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the flexible IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone of the Flexible IP.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public FlexibleIpArgs()
        {
        }
        public static new FlexibleIpArgs Empty => new FlexibleIpArgs();
    }

    public sealed class FlexibleIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the Flexible IP (Format ISO 8601).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A description of the flexible IP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address of the Flexible IP.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Defines whether the flexible IP has an IPv6 address.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// The organization of the Flexible IP.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The project of the Flexible IP.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse domain associated with this flexible IP.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// The ID of the associated server.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// The status of the flexible IP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the flexible IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time of the last update of the Flexible IP (Format ISO 8601).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The zone of the Flexible IP.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public FlexibleIpState()
        {
        }
        public static new FlexibleIpState Empty => new FlexibleIpState();
    }
}
