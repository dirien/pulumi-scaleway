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
    /// Creates and manages Scaleway VPC Public Gateway IP.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
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
    ///     var main = new Scaleway.VpcPublicGatewayIp("main", new()
    ///     {
    ///         Reverse = "tf.example.com",
    ///     });
    /// 
    ///     var tfA = new Scaleway.DomainRecord("tfA", new()
    ///     {
    ///         Data = main.Address,
    ///         DnsZone = "example.com",
    ///         Priority = 1,
    ///         Ttl = 3600,
    ///         Type = "A",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Public gateway can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp")]
    public partial class VpcPublicGatewayIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address itself.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the public gateway ip.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The organization ID the public gateway ip is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the public gateway ip is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The reverse domain name for the IP address
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the public gateway IP.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the public gateway ip.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the public gateway ip should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGatewayIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGatewayIp(string name, VpcPublicGatewayIpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, args ?? new VpcPublicGatewayIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGatewayIp(string name, Input<string> id, VpcPublicGatewayIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPublicGatewayIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGatewayIp Get(string name, Input<string> id, VpcPublicGatewayIpState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGatewayIp(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the public gateway ip is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse domain name for the IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the public gateway IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the public gateway ip should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpArgs()
        {
        }
        public static new VpcPublicGatewayIpArgs Empty => new VpcPublicGatewayIpArgs();
    }

    public sealed class VpcPublicGatewayIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address itself.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The date and time of the creation of the public gateway ip.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The organization ID the public gateway ip is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the public gateway ip is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse domain name for the IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the public gateway IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time of the last update of the public gateway ip.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// `zone`) The zone in which the public gateway ip should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpState()
        {
        }
        public static new VpcPublicGatewayIpState Empty => new VpcPublicGatewayIpState();
    }
}
