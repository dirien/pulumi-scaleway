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
    public static class GetLbRoute
    {
        /// <summary>
        /// Get information about Scaleway Load-Balancer Routes.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#route-ff94b7).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var byID = Scaleway.GetLbRoute.Invoke(new()
        ///     {
        ///         RouteId = rt01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLbRouteResult> InvokeAsync(GetLbRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbRouteResult>("scaleway:index/getLbRoute:getLbRoute", args ?? new GetLbRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load-Balancer Routes.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#route-ff94b7).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var byID = Scaleway.GetLbRoute.Invoke(new()
        ///     {
        ///         RouteId = rt01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLbRouteResult> Invoke(GetLbRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbRouteResult>("scaleway:index/getLbRoute:getLbRoute", args ?? new GetLbRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbRouteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The route id.
        /// </summary>
        [Input("routeId", required: true)]
        public string RouteId { get; set; } = null!;

        public GetLbRouteArgs()
        {
        }
        public static new GetLbRouteArgs Empty => new GetLbRouteArgs();
    }

    public sealed class GetLbRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The route id.
        /// </summary>
        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        public GetLbRouteInvokeArgs()
        {
        }
        public static new GetLbRouteInvokeArgs Empty => new GetLbRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbRouteResult
    {
        public readonly string BackendId;
        public readonly string CreatedAt;
        public readonly string FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MatchHostHeader;
        public readonly string MatchSni;
        public readonly string RouteId;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetLbRouteResult(
            string backendId,

            string createdAt,

            string frontendId,

            string id,

            string matchHostHeader,

            string matchSni,

            string routeId,

            string updatedAt)
        {
            BackendId = backendId;
            CreatedAt = createdAt;
            FrontendId = frontendId;
            Id = id;
            MatchHostHeader = matchHostHeader;
            MatchSni = matchSni;
            RouteId = routeId;
            UpdatedAt = updatedAt;
        }
    }
}
