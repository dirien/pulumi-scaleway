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
    public static class GetWebhostingOffer
    {
        /// <summary>
        /// Gets information about a webhosting offer.
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
        ///     var byName = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         Name = "performance",
        ///     });
        /// 
        ///     var byId = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         OfferId = "de2426b4-a9e9-11ec-b909-0242ac120002",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetWebhostingOfferResult> InvokeAsync(GetWebhostingOfferArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebhostingOfferResult>("scaleway:index/getWebhostingOffer:getWebhostingOffer", args ?? new GetWebhostingOfferArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a webhosting offer.
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
        ///     var byName = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         Name = "performance",
        ///     });
        /// 
        ///     var byId = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         OfferId = "de2426b4-a9e9-11ec-b909-0242ac120002",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWebhostingOfferResult> Invoke(GetWebhostingOfferInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhostingOfferResult>("scaleway:index/getWebhostingOffer:getWebhostingOffer", args ?? new GetWebhostingOfferInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a webhosting offer.
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
        ///     var byName = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         Name = "performance",
        ///     });
        /// 
        ///     var byId = Scaleway.GetWebhostingOffer.Invoke(new()
        ///     {
        ///         OfferId = "de2426b4-a9e9-11ec-b909-0242ac120002",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWebhostingOfferResult> Invoke(GetWebhostingOfferInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebhostingOfferResult>("scaleway:index/getWebhostingOffer:getWebhostingOffer", args ?? new GetWebhostingOfferInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebhostingOfferArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The offer name. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The offer id. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("offerId")]
        public string? OfferId { get; set; }

        /// <summary>
        /// `region`) The region in which offer exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetWebhostingOfferArgs()
        {
        }
        public static new GetWebhostingOfferArgs Empty => new GetWebhostingOfferArgs();
    }

    public sealed class GetWebhostingOfferInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The offer name. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The offer id. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("offerId")]
        public Input<string>? OfferId { get; set; }

        /// <summary>
        /// `region`) The region in which offer exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetWebhostingOfferInvokeArgs()
        {
        }
        public static new GetWebhostingOfferInvokeArgs Empty => new GetWebhostingOfferInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebhostingOfferResult
    {
        /// <summary>
        /// The unique identifier used for billing.
        /// </summary>
        public readonly string BillingOperationPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? OfferId;
        /// <summary>
        /// The offer price.
        /// </summary>
        public readonly string Price;
        /// <summary>
        /// The offer product.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWebhostingOfferProductResult> Products;
        public readonly string Region;

        [OutputConstructor]
        private GetWebhostingOfferResult(
            string billingOperationPath,

            string id,

            string? name,

            string? offerId,

            string price,

            ImmutableArray<Outputs.GetWebhostingOfferProductResult> products,

            string region)
        {
            BillingOperationPath = billingOperationPath;
            Id = id;
            Name = name;
            OfferId = offerId;
            Price = price;
            Products = products;
            Region = region;
        }
    }
}
