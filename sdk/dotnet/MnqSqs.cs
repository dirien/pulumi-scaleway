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
    /// Activate Scaleway Messaging and queuing SQS for a project.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // For default project in default region
    ///     var main = new Scaleway.MnqSqs("main");
    /// 
    ///     // For specific project in default region
    ///     var forProject = new Scaleway.MnqSqs("forProject", new()
    ///     {
    ///         ProjectId = scaleway_account_project.Main.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQS status can be imported using the `{region}/{project_id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/mnqSqs:MnqSqs main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqSqs:MnqSqs")]
    public partial class MnqSqs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the SQS service for this project.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a MnqSqs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqSqs(string name, MnqSqsArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSqs:MnqSqs", name, args ?? new MnqSqsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqSqs(string name, Input<string> id, MnqSqsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSqs:MnqSqs", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqSqs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqSqs Get(string name, Input<string> id, MnqSqsState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqSqs(name, id, state, options);
        }
    }

    public sealed class MnqSqsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSqsArgs()
        {
        }
        public static new MnqSqsArgs Empty => new MnqSqsArgs();
    }

    public sealed class MnqSqsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the SQS service for this project.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSqsState()
        {
        }
        public static new MnqSqsState Empty => new MnqSqsState();
    }
}
