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
    /// Creates and manages Scaleway Load-Balancer Routes.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
    /// 
    /// ## Examples
    /// 
    /// ### With SNI for direction to TCP backends
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ip01 = new Scaleway.LbIp("ip01");
    /// 
    ///     var lb01 = new Scaleway.Lb("lb01", new()
    ///     {
    ///         IpId = ip01.Id,
    ///         Type = "lb-s",
    ///     });
    /// 
    ///     var bkd01 = new Scaleway.LbBackend("bkd01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         ForwardProtocol = "tcp",
    ///         ForwardPort = 80,
    ///         ProxyProtocol = "none",
    ///     });
    /// 
    ///     var frt01 = new Scaleway.LbFrontend("frt01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         BackendId = bkd01.Id,
    ///         InboundPort = 80,
    ///     });
    /// 
    ///     var rt01 = new Scaleway.LbRoute("rt01", new()
    ///     {
    ///         FrontendId = frt01.Id,
    ///         BackendId = bkd01.Id,
    ///         MatchSni = "sni.scaleway.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With host-header for direction to HTTP backends
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ip01 = new Scaleway.LbIp("ip01");
    /// 
    ///     var lb01 = new Scaleway.Lb("lb01", new()
    ///     {
    ///         IpId = ip01.Id,
    ///         Type = "lb-s",
    ///     });
    /// 
    ///     var bkd01 = new Scaleway.LbBackend("bkd01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         ForwardProtocol = "http",
    ///         ForwardPort = 80,
    ///         ProxyProtocol = "none",
    ///     });
    /// 
    ///     var frt01 = new Scaleway.LbFrontend("frt01", new()
    ///     {
    ///         LbId = lb01.Id,
    ///         BackendId = bkd01.Id,
    ///         InboundPort = 80,
    ///     });
    /// 
    ///     var rt01 = new Scaleway.LbRoute("rt01", new()
    ///     {
    ///         FrontendId = frt01.Id,
    ///         BackendId = bkd01.Id,
    ///         MatchHostHeader = "host.scaleway.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/lbRoute:LbRoute main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/lbRoute:LbRoute")]
    public partial class LbRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the backend to which the route is associated.
        /// </summary>
        [Output("backendId")]
        public Output<string> BackendId { get; private set; } = null!;

        /// <summary>
        /// The date at which the route was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the frontend to which the route is associated.
        /// </summary>
        [Output("frontendId")]
        public Output<string> FrontendId { get; private set; } = null!;

        /// <summary>
        /// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on HTTP Load Balancers.
        /// </summary>
        [Output("matchHostHeader")]
        public Output<string?> MatchHostHeader { get; private set; } = null!;

        /// <summary>
        /// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on TCP Load Balancers.
        /// </summary>
        [Output("matchSni")]
        public Output<string?> MatchSni { get; private set; } = null!;

        /// <summary>
        /// The date at which the route was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a LbRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbRoute(string name, LbRouteArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/lbRoute:LbRoute", name, args ?? new LbRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbRoute(string name, Input<string> id, LbRouteState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/lbRoute:LbRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbRoute Get(string name, Input<string> id, LbRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new LbRoute(name, id, state, options);
        }
    }

    public sealed class LbRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the backend to which the route is associated.
        /// </summary>
        [Input("backendId", required: true)]
        public Input<string> BackendId { get; set; } = null!;

        /// <summary>
        /// The ID of the frontend to which the route is associated.
        /// </summary>
        [Input("frontendId", required: true)]
        public Input<string> FrontendId { get; set; } = null!;

        /// <summary>
        /// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on HTTP Load Balancers.
        /// </summary>
        [Input("matchHostHeader")]
        public Input<string>? MatchHostHeader { get; set; }

        /// <summary>
        /// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on TCP Load Balancers.
        /// </summary>
        [Input("matchSni")]
        public Input<string>? MatchSni { get; set; }

        public LbRouteArgs()
        {
        }
        public static new LbRouteArgs Empty => new LbRouteArgs();
    }

    public sealed class LbRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the backend to which the route is associated.
        /// </summary>
        [Input("backendId")]
        public Input<string>? BackendId { get; set; }

        /// <summary>
        /// The date at which the route was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the frontend to which the route is associated.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        /// <summary>
        /// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on HTTP Load Balancers.
        /// </summary>
        [Input("matchHostHeader")]
        public Input<string>? MatchHostHeader { get; set; }

        /// <summary>
        /// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        /// Only one of `match_sni` and `match_host_header` should be specified.
        /// 
        /// &gt; **Important:** This field should be set for routes on TCP Load Balancers.
        /// </summary>
        [Input("matchSni")]
        public Input<string>? MatchSni { get; set; }

        /// <summary>
        /// The date at which the route was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public LbRouteState()
        {
        }
        public static new LbRouteState Empty => new LbRouteState();
    }
}
