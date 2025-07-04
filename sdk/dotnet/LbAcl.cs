// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// Creates and manages Scaleway Load Balancer ACLs.
    /// 
    /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/acls/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls-get-an-acl).
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
    ///     var acl01 = new Scaleway.LbAcl("acl01", new()
    ///     {
    ///         FrontendId = scaleway_lb_frontend.Frt01.Id,
    ///         Description = "Exclude well-known IPs",
    ///         Index = 0,
    ///         Action = new Scaleway.Inputs.LbAclActionArgs
    ///         {
    ///             Type = "allow",
    ///         },
    ///         Match = new Scaleway.Inputs.LbAclMatchArgs
    ///         {
    ///             IpSubnets = new[]
    ///             {
    ///                 "192.168.0.1",
    ///                 "192.168.0.2",
    ///                 "192.168.10.0/24",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer ACLs can be imported using `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/lbAcl:LbAcl acl01 fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/lbAcl:LbAcl")]
    public partial class LbAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        [Output("action")]
        public Output<Outputs.LbAclAction> Action { get; private set; } = null!;

        /// <summary>
        /// IsDate and time of ACL's creation (RFC 3339 format)
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ACL description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the Load Balancer frontend to attach the ACL to.
        /// </summary>
        [Output("frontendId")]
        public Output<string> FrontendId { get; private set; } = null!;

        /// <summary>
        /// The priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        /// </summary>
        [Output("index")]
        public Output<int> Index { get; private set; } = null!;

        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        [Output("match")]
        public Output<Outputs.LbAclMatch?> Match { get; private set; } = null!;

        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IsDate and time of ACL's update (RFC 3339 format)
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a LbAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbAcl(string name, LbAclArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/lbAcl:LbAcl", name, args ?? new LbAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbAcl(string name, Input<string> id, LbAclState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/lbAcl:LbAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbAcl Get(string name, Input<string> id, LbAclState? state = null, CustomResourceOptions? options = null)
        {
            return new LbAcl(name, id, state, options);
        }
    }

    public sealed class LbAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.LbAclActionArgs> Action { get; set; } = null!;

        /// <summary>
        /// The ACL description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Load Balancer frontend to attach the ACL to.
        /// </summary>
        [Input("frontendId", required: true)]
        public Input<string> FrontendId { get; set; } = null!;

        /// <summary>
        /// The priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        /// </summary>
        [Input("index", required: true)]
        public Input<int> Index { get; set; } = null!;

        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        [Input("match")]
        public Input<Inputs.LbAclMatchArgs>? Match { get; set; }

        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LbAclArgs()
        {
        }
        public static new LbAclArgs Empty => new LbAclArgs();
    }

    public sealed class LbAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        [Input("action")]
        public Input<Inputs.LbAclActionGetArgs>? Action { get; set; }

        /// <summary>
        /// IsDate and time of ACL's creation (RFC 3339 format)
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ACL description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Load Balancer frontend to attach the ACL to.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        /// <summary>
        /// The priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        [Input("match")]
        public Input<Inputs.LbAclMatchGetArgs>? Match { get; set; }

        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IsDate and time of ACL's update (RFC 3339 format)
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public LbAclState()
        {
        }
        public static new LbAclState Empty => new LbAclState();
    }
}
