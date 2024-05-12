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
    public static class GetLbs
    {
        /// <summary>
        /// Gets information about multiple Load Balancers.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetLbs.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLbsResult> InvokeAsync(GetLbsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbsResult>("scaleway:index/getLbs:getLbs", args ?? new GetLbsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancers.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetLbs.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbsResult> Invoke(GetLbsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbsResult>("scaleway:index/getLbs:getLbs", args ?? new GetLbsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The load balancer name used as a filter. LBs with a name like it are listed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the load-balancer is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which LBs exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetLbsArgs()
        {
        }
        public static new GetLbsArgs Empty => new GetLbsArgs();
    }

    public sealed class GetLbsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The load balancer name used as a filter. LBs with a name like it are listed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the load-balancer is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which LBs exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetLbsInvokeArgs()
        {
        }
        public static new GetLbsInvokeArgs Empty => new GetLbsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of found LBs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbsLbResult> Lbs;
        /// <summary>
        /// The name of the load-balancer.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The organization ID the load-balancer is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the project the load-balancer is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The zone in which the load-balancer is.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetLbsResult(
            string id,

            ImmutableArray<Outputs.GetLbsLbResult> lbs,

            string? name,

            string organizationId,

            string projectId,

            string zone)
        {
            Id = id;
            Lbs = lbs;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Zone = zone;
        }
    }
}
