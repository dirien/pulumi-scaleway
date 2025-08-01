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
    public static class GetK8sVersion
    {
        /// <summary>
        /// Gets information about a Kubernetes version.
        /// For more information, see the [API documentation](https://developers.scaleway.com/en/products/k8s/api).
        /// 
        /// You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.
        /// 
        /// ## Example Usage
        /// 
        /// ### Use the latest version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var latest = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "latest",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Use a specific version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "1.26.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetK8sVersionResult> InvokeAsync(GetK8sVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetK8sVersionResult>("scaleway:index/getK8sVersion:getK8sVersion", args ?? new GetK8sVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes version.
        /// For more information, see the [API documentation](https://developers.scaleway.com/en/products/k8s/api).
        /// 
        /// You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.
        /// 
        /// ## Example Usage
        /// 
        /// ### Use the latest version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var latest = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "latest",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Use a specific version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "1.26.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetK8sVersionResult> Invoke(GetK8sVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetK8sVersionResult>("scaleway:index/getK8sVersion:getK8sVersion", args ?? new GetK8sVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes version.
        /// For more information, see the [API documentation](https://developers.scaleway.com/en/products/k8s/api).
        /// 
        /// You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.
        /// 
        /// ## Example Usage
        /// 
        /// ### Use the latest version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var latest = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "latest",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Use a specific version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetK8sVersion.Invoke(new()
        ///     {
        ///         Name = "1.26.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetK8sVersionResult> Invoke(GetK8sVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetK8sVersionResult>("scaleway:index/getK8sVersion:getK8sVersion", args ?? new GetK8sVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetK8sVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kubernetes version.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the version exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetK8sVersionArgs()
        {
        }
        public static new GetK8sVersionArgs Empty => new GetK8sVersionArgs();
    }

    public sealed class GetK8sVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kubernetes version.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the version exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetK8sVersionInvokeArgs()
        {
        }
        public static new GetK8sVersionInvokeArgs Empty => new GetK8sVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetK8sVersionResult
    {
        /// <summary>
        /// The list of supported Container Network Interface (CNI) plugins for this version.
        /// </summary>
        public readonly ImmutableArray<string> AvailableCnis;
        /// <summary>
        /// The list of supported container runtimes for this version.
        /// </summary>
        public readonly ImmutableArray<string> AvailableContainerRuntimes;
        /// <summary>
        /// The list of supported feature gates for this version.
        /// </summary>
        public readonly ImmutableArray<string> AvailableFeatureGates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Region;

        [OutputConstructor]
        private GetK8sVersionResult(
            ImmutableArray<string> availableCnis,

            ImmutableArray<string> availableContainerRuntimes,

            ImmutableArray<string> availableFeatureGates,

            string id,

            string name,

            string region)
        {
            AvailableCnis = availableCnis;
            AvailableContainerRuntimes = availableContainerRuntimes;
            AvailableFeatureGates = availableFeatureGates;
            Id = id;
            Name = name;
            Region = region;
        }
    }
}
