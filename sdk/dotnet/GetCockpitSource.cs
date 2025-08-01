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
    public static class GetCockpitSource
    {
        /// <summary>
        /// The `scaleway.CockpitSource` data source allows you to retrieve information about a specific [data source](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
        /// 
        /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve a specific data source by ID
        /// 
        /// The following example retrieves a Cockpit data source by its unique ID.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Scaleway.GetCockpitSource.Invoke(new()
        ///     {
        ///         Id = "fr-par/11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCockpitSourceResult> InvokeAsync(GetCockpitSourceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCockpitSourceResult>("scaleway:index/getCockpitSource:getCockpitSource", args ?? new GetCockpitSourceArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.CockpitSource` data source allows you to retrieve information about a specific [data source](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
        /// 
        /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve a specific data source by ID
        /// 
        /// The following example retrieves a Cockpit data source by its unique ID.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Scaleway.GetCockpitSource.Invoke(new()
        ///     {
        ///         Id = "fr-par/11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCockpitSourceResult> Invoke(GetCockpitSourceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCockpitSourceResult>("scaleway:index/getCockpitSource:getCockpitSource", args ?? new GetCockpitSourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.CockpitSource` data source allows you to retrieve information about a specific [data source](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
        /// 
        /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
        /// 
        /// ## Example Usage
        /// 
        /// ### Retrieve a specific data source by ID
        /// 
        /// The following example retrieves a Cockpit data source by its unique ID.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Scaleway.GetCockpitSource.Invoke(new()
        ///     {
        ///         Id = "fr-par/11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCockpitSourceResult> Invoke(GetCockpitSourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCockpitSourceResult>("scaleway:index/getCockpitSource:getCockpitSource", args ?? new GetCockpitSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCockpitSourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Cockpit data source in the `{region}/{id}` format. If specified, other filters are ignored.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The origin of the data source. Possible values are:
        /// </summary>
        [Input("origin")]
        public string? Origin { get; set; }

        /// <summary>
        /// The ID of the Project the data source is associated with. Defaults to the Project ID specified in the provider configuration.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetCockpitSourceArgs()
        {
        }
        public static new GetCockpitSourceArgs Empty => new GetCockpitSourceArgs();
    }

    public sealed class GetCockpitSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Cockpit data source in the `{region}/{id}` format. If specified, other filters are ignored.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The origin of the data source. Possible values are:
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// The ID of the Project the data source is associated with. Defaults to the Project ID specified in the provider configuration.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetCockpitSourceInvokeArgs()
        {
        }
        public static new GetCockpitSourceInvokeArgs Empty => new GetCockpitSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetCockpitSourceResult
    {
        /// <summary>
        /// The date and time the data source was created (in RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The unique identifier of the data source in the `{region}/{id}` format.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The origin of the data source.
        /// </summary>
        public readonly string Origin;
        public readonly string ProjectId;
        public readonly string Region;
        /// <summary>
        /// The number of days the data is retained in the data source.
        /// </summary>
        public readonly int RetentionDays;
        /// <summary>
        /// Indicates whether the data source is synchronized with Grafana.
        /// </summary>
        public readonly bool SynchronizedWithGrafana;
        public readonly string Type;
        /// <summary>
        /// The date and time the data source was last updated (in RFC 3339 format).
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The URL of the Cockpit data source.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetCockpitSourceResult(
            string createdAt,

            string id,

            string name,

            string origin,

            string projectId,

            string region,

            int retentionDays,

            bool synchronizedWithGrafana,

            string type,

            string updatedAt,

            string url)
        {
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            Origin = origin;
            ProjectId = projectId;
            Region = region;
            RetentionDays = retentionDays;
            SynchronizedWithGrafana = synchronizedWithGrafana;
            Type = type;
            UpdatedAt = updatedAt;
            Url = url;
        }
    }
}
