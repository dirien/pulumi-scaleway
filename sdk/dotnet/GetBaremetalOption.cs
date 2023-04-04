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
    public static class GetBaremetalOption
    {
        /// <summary>
        /// Gets information about a baremetal option.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.GetBaremetalOption.Invoke(new()
        ///     {
        ///         Name = "Remote Access",
        ///     });
        /// 
        ///     var byId = Scaleway.GetBaremetalOption.Invoke(new()
        ///     {
        ///         OptionId = "931df052-d713-4674-8b58-96a63244c8e2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBaremetalOptionResult> InvokeAsync(GetBaremetalOptionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBaremetalOptionResult>("scaleway:index/getBaremetalOption:getBaremetalOption", args ?? new GetBaremetalOptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a baremetal option.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.GetBaremetalOption.Invoke(new()
        ///     {
        ///         Name = "Remote Access",
        ///     });
        /// 
        ///     var byId = Scaleway.GetBaremetalOption.Invoke(new()
        ///     {
        ///         OptionId = "931df052-d713-4674-8b58-96a63244c8e2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBaremetalOptionResult> Invoke(GetBaremetalOptionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBaremetalOptionResult>("scaleway:index/getBaremetalOption:getBaremetalOption", args ?? new GetBaremetalOptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBaremetalOptionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The option name. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The option id. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("optionId")]
        public string? OptionId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the option exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetBaremetalOptionArgs()
        {
        }
        public static new GetBaremetalOptionArgs Empty => new GetBaremetalOptionArgs();
    }

    public sealed class GetBaremetalOptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The option name. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The option id. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("optionId")]
        public Input<string>? OptionId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the option exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBaremetalOptionInvokeArgs()
        {
        }
        public static new GetBaremetalOptionInvokeArgs Empty => new GetBaremetalOptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetBaremetalOptionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Is false if the option could not be added or removed.
        /// </summary>
        public readonly bool Manageable;
        /// <summary>
        /// The name of the option.
        /// </summary>
        public readonly string? Name;
        public readonly string? OptionId;
        public readonly string Zone;

        [OutputConstructor]
        private GetBaremetalOptionResult(
            string id,

            bool manageable,

            string? name,

            string? optionId,

            string zone)
        {
            Id = id;
            Manageable = manageable;
            Name = name;
            OptionId = optionId;
            Zone = zone;
        }
    }
}
