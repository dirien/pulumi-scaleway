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
    public static class GetLbRoutes
    {
        /// <summary>
        /// Gets information about multiple Load Balancer Routes.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byFrontendID = Scaleway.GetLbRoutes.Invoke(new()
        ///     {
        ///         FrontendId = scaleway_lb_frontend.Frt01.Id,
        ///     });
        /// 
        ///     var myKey = Scaleway.GetLbRoutes.Invoke(new()
        ///     {
        ///         FrontendId = "11111111-1111-1111-1111-111111111111",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLbRoutesResult> InvokeAsync(GetLbRoutesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbRoutesResult>("scaleway:index/getLbRoutes:getLbRoutes", args ?? new GetLbRoutesArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer Routes.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byFrontendID = Scaleway.GetLbRoutes.Invoke(new()
        ///     {
        ///         FrontendId = scaleway_lb_frontend.Frt01.Id,
        ///     });
        /// 
        ///     var myKey = Scaleway.GetLbRoutes.Invoke(new()
        ///     {
        ///         FrontendId = "11111111-1111-1111-1111-111111111111",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLbRoutesResult> Invoke(GetLbRoutesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbRoutesResult>("scaleway:index/getLbRoutes:getLbRoutes", args ?? new GetLbRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbRoutesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID origin of redirection used as a filter. routes with a frontend ID like it are listed.
        /// </summary>
        [Input("frontendId")]
        public string? FrontendId { get; set; }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which routes exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetLbRoutesArgs()
        {
        }
        public static new GetLbRoutesArgs Empty => new GetLbRoutesArgs();
    }

    public sealed class GetLbRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend ID origin of redirection used as a filter. routes with a frontend ID like it are listed.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which routes exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetLbRoutesInvokeArgs()
        {
        }
        public static new GetLbRoutesInvokeArgs Empty => new GetLbRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbRoutesResult
    {
        public readonly string? FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        /// <summary>
        /// List of found routes
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbRoutesRouteResult> Routes;
        public readonly string Zone;

        [OutputConstructor]
        private GetLbRoutesResult(
            string? frontendId,

            string id,

            string organizationId,

            string projectId,

            ImmutableArray<Outputs.GetLbRoutesRouteResult> routes,

            string zone)
        {
            FrontendId = frontendId;
            Id = id;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Routes = routes;
            Zone = zone;
        }
    }
}
