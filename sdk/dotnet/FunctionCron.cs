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
    /// Creates and manages Scaleway Function Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
    /// 
    /// For more details about the limitation
    /// check [functions-limitations](https://www.scaleway.com/en/docs/compute/functions/reference-content/functions-limitations/).
    /// 
    /// You can check also
    /// our [functions cron api documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#crons-942bf4).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainFunctionNamespace = new Scaleway.FunctionNamespace("mainFunctionNamespace");
    /// 
    ///     var mainFunction = new Scaleway.Function("mainFunction", new()
    ///     {
    ///         NamespaceId = mainFunctionNamespace.Id,
    ///         Runtime = "node14",
    ///         Privacy = "private",
    ///         Handler = "handler.handle",
    ///     });
    /// 
    ///     var mainFunctionCron = new Scaleway.FunctionCron("mainFunctionCron", new()
    ///     {
    ///         FunctionId = mainFunction.Id,
    ///         Schedule = "0 0 * * *",
    ///         Args = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["test"] = "scw",
    ///         }),
    ///     });
    /// 
    ///     var func = new Scaleway.FunctionCron("func", new()
    ///     {
    ///         FunctionId = mainFunction.Id,
    ///         Schedule = "0 1 * * *",
    ///         Args = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["my_var"] = "terraform",
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Container Cron can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/functionCron:FunctionCron")]
    public partial class FunctionCron : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The key-value mapping to define arguments that will be passed to your function’s event object
        /// during
        /// </summary>
        [Output("args")]
        public Output<string> Args { get; private set; } = null!;

        /// <summary>
        /// The function ID to link with your cron.
        /// </summary>
        [Output("functionId")]
        public Output<string> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The name of the cron. If not provided, the name is generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in where the job was created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
        /// executed.
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// The cron status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionCron resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionCron(string name, FunctionCronArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/functionCron:FunctionCron", name, args ?? new FunctionCronArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionCron(string name, Input<string> id, FunctionCronState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionCron:FunctionCron", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionCron resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionCron Get(string name, Input<string> id, FunctionCronState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionCron(name, id, state, options);
        }
    }

    public sealed class FunctionCronArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key-value mapping to define arguments that will be passed to your function’s event object
        /// during
        /// </summary>
        [Input("args", required: true)]
        public Input<string> Args { get; set; } = null!;

        /// <summary>
        /// The function ID to link with your cron.
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        /// <summary>
        /// The name of the cron. If not provided, the name is generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region
        /// in where the job was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
        /// executed.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        public FunctionCronArgs()
        {
        }
        public static new FunctionCronArgs Empty => new FunctionCronArgs();
    }

    public sealed class FunctionCronState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key-value mapping to define arguments that will be passed to your function’s event object
        /// during
        /// </summary>
        [Input("args")]
        public Input<string>? Args { get; set; }

        /// <summary>
        /// The function ID to link with your cron.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The name of the cron. If not provided, the name is generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region
        /// in where the job was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
        /// executed.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// The cron status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public FunctionCronState()
        {
        }
        public static new FunctionCronState Empty => new FunctionCronState();
    }
}
