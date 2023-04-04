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
    public static class GetMarketplaceImage
    {
        /// <summary>
        /// Gets local image ID of an image from its label name.
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
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myImage = Scaleway.GetMarketplaceImage.Invoke(new()
        ///     {
        ///         Label = "ubuntu_jammy",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMarketplaceImageResult> InvokeAsync(GetMarketplaceImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMarketplaceImageResult>("scaleway:index/getMarketplaceImage:getMarketplaceImage", args ?? new GetMarketplaceImageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets local image ID of an image from its label name.
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
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myImage = Scaleway.GetMarketplaceImage.Invoke(new()
        ///     {
        ///         Label = "ubuntu_jammy",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMarketplaceImageResult> Invoke(GetMarketplaceImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMarketplaceImageResult>("scaleway:index/getMarketplaceImage:getMarketplaceImage", args ?? new GetMarketplaceImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMarketplaceImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance type the image is compatible with.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
        /// to find the right `label`.
        /// </summary>
        [Input("label", required: true)]
        public string Label { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetMarketplaceImageArgs()
        {
        }
        public static new GetMarketplaceImageArgs Empty => new GetMarketplaceImageArgs();
    }

    public sealed class GetMarketplaceImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance type the image is compatible with.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
        /// to find the right `label`.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetMarketplaceImageInvokeArgs()
        {
        }
        public static new GetMarketplaceImageInvokeArgs Empty => new GetMarketplaceImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetMarketplaceImageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceType;
        public readonly string Label;
        public readonly string Zone;

        [OutputConstructor]
        private GetMarketplaceImageResult(
            string id,

            string? instanceType,

            string label,

            string zone)
        {
            Id = id;
            InstanceType = instanceType;
            Label = label;
            Zone = zone;
        }
    }
}
