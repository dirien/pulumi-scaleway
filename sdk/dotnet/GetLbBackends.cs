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
    public static class GetLbBackends
    {
        /// <summary>
        /// Gets information about multiple Load Balancer Backends.
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
        ///     var byLBID = Scaleway.GetLbBackends.Invoke(new()
        ///     {
        ///         LbId = scaleway_lb.Lb01.Id,
        ///     });
        /// 
        ///     var byLBIDAndName = Scaleway.GetLbBackends.Invoke(new()
        ///     {
        ///         LbId = scaleway_lb.Lb01.Id,
        ///         Name = "tf-backend-datasource",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLbBackendsResult> InvokeAsync(GetLbBackendsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbBackendsResult>("scaleway:index/getLbBackends:getLbBackends", args ?? new GetLbBackendsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Load Balancer Backends.
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
        ///     var byLBID = Scaleway.GetLbBackends.Invoke(new()
        ///     {
        ///         LbId = scaleway_lb.Lb01.Id,
        ///     });
        /// 
        ///     var byLBIDAndName = Scaleway.GetLbBackends.Invoke(new()
        ///     {
        ///         LbId = scaleway_lb.Lb01.Id,
        ///         Name = "tf-backend-datasource",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbBackendsResult> Invoke(GetLbBackendsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbBackendsResult>("scaleway:index/getLbBackends:getLbBackends", args ?? new GetLbBackendsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbBackendsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
        /// </summary>
        [Input("lbId", required: true)]
        public string LbId { get; set; } = null!;

        /// <summary>
        /// The backend name used as filter. Backends with a name like it are listed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which backends exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetLbBackendsArgs()
        {
        }
        public static new GetLbBackendsArgs Empty => new GetLbBackendsArgs();
    }

    public sealed class GetLbBackendsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
        /// </summary>
        [Input("lbId", required: true)]
        public Input<string> LbId { get; set; } = null!;

        /// <summary>
        /// The backend name used as filter. Backends with a name like it are listed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which backends exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetLbBackendsInvokeArgs()
        {
        }
        public static new GetLbBackendsInvokeArgs Empty => new GetLbBackendsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbBackendsResult
    {
        /// <summary>
        /// List of found backends
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLbBackendsBackendResult> Backends;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LbId;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Zone;

        [OutputConstructor]
        private GetLbBackendsResult(
            ImmutableArray<Outputs.GetLbBackendsBackendResult> backends,

            string id,

            string lbId,

            string? name,

            string organizationId,

            string projectId,

            string zone)
        {
            Backends = backends;
            Id = id;
            LbId = lbId;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Zone = zone;
        }
    }
}
