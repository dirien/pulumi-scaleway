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
    /// The `scaleway.CockpitSource` resource allows you to create and manage [data sources](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
    /// 
    /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a data source
    /// 
    /// The following command allows you to create a [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) data source named `my-data-source` in a given Project.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Scaleway.AccountProject("project");
    /// 
    ///     var main = new Scaleway.CockpitSource("main", new()
    ///     {
    ///         ProjectId = project.Id,
    ///         Type = "metrics",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This section explains how to import a data source using the ID of the region it is located in, in the `{region}/{id}` format.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/cockpitSource:CockpitSource main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/cockpitSource:CockpitSource")]
    public partial class CockpitSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the data source was created (in RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The origin of the Cockpit data source.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// ) The ID of the Project the data source is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The URL endpoint used for pushing data to the Cockpit data source.
        /// </summary>
        [Output("pushUrl")]
        public Output<string> PushUrl { get; private set; } = null!;

        /// <summary>
        /// ) The region where the data source is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the data source is synchronized with Grafana.
        /// </summary>
        [Output("synchronizedWithGrafana")]
        public Output<bool> SynchronizedWithGrafana { get; private set; } = null!;

        /// <summary>
        /// The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The date and time the data source was last updated (in RFC 3339 format).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL of the Cockpit data source.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a CockpitSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CockpitSource(string name, CockpitSourceArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitSource:CockpitSource", name, args ?? new CockpitSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CockpitSource(string name, Input<string> id, CockpitSourceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitSource:CockpitSource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CockpitSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CockpitSource Get(string name, Input<string> id, CockpitSourceState? state = null, CustomResourceOptions? options = null)
        {
            return new CockpitSource(name, id, state, options);
        }
    }

    public sealed class CockpitSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ) The ID of the Project the data source is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// ) The region where the data source is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CockpitSourceArgs()
        {
        }
        public static new CockpitSourceArgs Empty => new CockpitSourceArgs();
    }

    public sealed class CockpitSourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the data source was created (in RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The origin of the Cockpit data source.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// ) The ID of the Project the data source is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The URL endpoint used for pushing data to the Cockpit data source.
        /// </summary>
        [Input("pushUrl")]
        public Input<string>? PushUrl { get; set; }

        /// <summary>
        /// ) The region where the data source is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether the data source is synchronized with Grafana.
        /// </summary>
        [Input("synchronizedWithGrafana")]
        public Input<bool>? SynchronizedWithGrafana { get; set; }

        /// <summary>
        /// The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date and time the data source was last updated (in RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The URL of the Cockpit data source.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public CockpitSourceState()
        {
        }
        public static new CockpitSourceState Empty => new CockpitSourceState();
    }
}
