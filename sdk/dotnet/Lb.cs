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
    /// Creates and manages Scaleway Load-Balancers.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api).
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
    ///     var main = new Scaleway.LbIp("main", new()
    ///     {
    ///         Zone = "fr-par-1",
    ///     });
    /// 
    ///     var @base = new Scaleway.Lb("base", new()
    ///     {
    ///         IpIds = new[]
    ///         {
    ///             main.Id,
    ///         },
    ///         Zone = main.Zone,
    ///         Type = "LB-S",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Private LB
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @base = new Scaleway.Lb("base", new()
    ///     {
    ///         AssignFlexibleIp = false,
    ///         Type = "LB-S",
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
    ///     var v4 = new Scaleway.LbIp("v4");
    /// 
    ///     var v6 = new Scaleway.LbIp("v6", new()
    ///     {
    ///         IsIpv6 = true,
    ///     });
    /// 
    ///     var main = new Scaleway.Lb("main", new()
    ///     {
    ///         IpIds = new[]
    ///         {
    ///             v4.Id,
    ///             v6.Id,
    ///         },
    ///         Type = "LB-S",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## IP ID
    /// 
    /// Since v1.15.0, `ip_id` is a required field. This means that now a separate `scaleway.LbIp` is required.
    /// When importing, the IP needs to be imported as well as the LB.
    /// When upgrading to v1.15.0, you will need to create a new `scaleway.LbIp` resource and import it.
    /// 
    /// For instance, if you had the following:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.Lb("main", new()
    ///     {
    ///         Type = "LB-S",
    ///         Zone = "fr-par-1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// You will need to update it to:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainLbIp = new Scaleway.LbIp("mainLbIp");
    /// 
    ///     var mainLb = new Scaleway.Lb("mainLb", new()
    ///     {
    ///         IpId = mainLbIp.Id,
    ///         Zone = "fr-par-1",
    ///         Type = "LB-S",
    ///         ReleaseIp = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load-Balancer can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/lb:Lb main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    /// Be aware that you will also need to import the `scaleway_lb_ip` resource.
    /// </summary>
    [ScalewayResourceType("scaleway:index/lb:Lb")]
    public partial class Lb : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defines whether to automatically assign a flexible public IP to the load-balancer.
        /// </summary>
        [Output("assignFlexibleIp")]
        public Output<bool?> AssignFlexibleIp { get; private set; } = null!;

        /// <summary>
        /// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
        /// </summary>
        [Output("assignFlexibleIpv6")]
        public Output<bool?> AssignFlexibleIpv6 { get; private set; } = null!;

        /// <summary>
        /// The description of the load-balancer.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The load-balancer public IPv4 Address.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated LB IP. See below.
        /// 
        /// &gt; **Important:** Updates to `ip_id` will recreate the load-balancer.
        /// </summary>
        [Output("ipId")]
        public Output<string> IpId { get; private set; } = null!;

        /// <summary>
        /// The List of IP IDs to attach to the Load Balancer.
        /// </summary>
        [Output("ipIds")]
        public Output<ImmutableArray<string>> IpIds { get; private set; } = null!;

        /// <summary>
        /// The load-balancer public IPv6 Address.
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// The name of the load-balancer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the load-balancer is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// List of private network to connect with your load balancer
        /// </summary>
        [Output("privateNetworks")]
        public Output<ImmutableArray<Outputs.LbPrivateNetwork>> PrivateNetworks { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the load-balancer is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region of the resource
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The release_ip allow release the ip address associated with the load-balancers.
        /// </summary>
        [Output("releaseIp")]
        public Output<bool?> ReleaseIp { get; private set; } = null!;

        /// <summary>
        /// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
        /// </summary>
        [Output("sslCompatibilityLevel")]
        public Output<string?> SslCompatibilityLevel { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the load-balancers.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the load-balancer. Please check the migration section to upgrade the type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone of the load-balancer.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Lb resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lb(string name, LbArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/lb:Lb", name, args ?? new LbArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lb(string name, Input<string> id, LbState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/lb:Lb", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Lb resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lb Get(string name, Input<string> id, LbState? state = null, CustomResourceOptions? options = null)
        {
            return new Lb(name, id, state, options);
        }
    }

    public sealed class LbArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether to automatically assign a flexible public IP to the load-balancer.
        /// </summary>
        [Input("assignFlexibleIp")]
        public Input<bool>? AssignFlexibleIp { get; set; }

        /// <summary>
        /// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
        /// </summary>
        [Input("assignFlexibleIpv6")]
        public Input<bool>? AssignFlexibleIpv6 { get; set; }

        /// <summary>
        /// The description of the load-balancer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the associated LB IP. See below.
        /// 
        /// &gt; **Important:** Updates to `ip_id` will recreate the load-balancer.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// The List of IP IDs to attach to the Load Balancer.
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The name of the load-balancer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateNetworks")]
        private InputList<Inputs.LbPrivateNetworkArgs>? _privateNetworks;

        /// <summary>
        /// List of private network to connect with your load balancer
        /// </summary>
        public InputList<Inputs.LbPrivateNetworkArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.LbPrivateNetworkArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the load-balancer is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The release_ip allow release the ip address associated with the load-balancers.
        /// </summary>
        [Input("releaseIp")]
        public Input<bool>? ReleaseIp { get; set; }

        /// <summary>
        /// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
        /// </summary>
        [Input("sslCompatibilityLevel")]
        public Input<string>? SslCompatibilityLevel { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the load-balancers.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the load-balancer. Please check the migration section to upgrade the type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone of the load-balancer.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LbArgs()
        {
        }
        public static new LbArgs Empty => new LbArgs();
    }

    public sealed class LbState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether to automatically assign a flexible public IP to the load-balancer.
        /// </summary>
        [Input("assignFlexibleIp")]
        public Input<bool>? AssignFlexibleIp { get; set; }

        /// <summary>
        /// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
        /// </summary>
        [Input("assignFlexibleIpv6")]
        public Input<bool>? AssignFlexibleIpv6 { get; set; }

        /// <summary>
        /// The description of the load-balancer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The load-balancer public IPv4 Address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The ID of the associated LB IP. See below.
        /// 
        /// &gt; **Important:** Updates to `ip_id` will recreate the load-balancer.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// The List of IP IDs to attach to the Load Balancer.
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The load-balancer public IPv6 Address.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        /// <summary>
        /// The name of the load-balancer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the load-balancer is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("privateNetworks")]
        private InputList<Inputs.LbPrivateNetworkGetArgs>? _privateNetworks;

        /// <summary>
        /// List of private network to connect with your load balancer
        /// </summary>
        public InputList<Inputs.LbPrivateNetworkGetArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.LbPrivateNetworkGetArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the load-balancer is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region of the resource
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The release_ip allow release the ip address associated with the load-balancers.
        /// </summary>
        [Input("releaseIp")]
        public Input<bool>? ReleaseIp { get; set; }

        /// <summary>
        /// Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
        /// </summary>
        [Input("sslCompatibilityLevel")]
        public Input<string>? SslCompatibilityLevel { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the load-balancers.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the load-balancer. Please check the migration section to upgrade the type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// `zone`) The zone of the load-balancer.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LbState()
        {
        }
        public static new LbState Empty => new LbState();
    }
}
