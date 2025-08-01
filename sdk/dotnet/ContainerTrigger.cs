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
    /// <summary>
    /// The `scaleway.ContainerTrigger` resource allows you to create and manage triggers for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
    /// 
    /// Refer to the Containers triggers [documentation](https://www.scaleway.com/en/docs/serverless/containers/how-to/add-trigger-to-a-container/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers-list-all-triggers) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ### SQS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.ContainerTrigger("main", new()
    ///     {
    ///         ContainerId = scaleway_container.Main.Id,
    ///         Sqs = new Scaleway.Inputs.ContainerTriggerSqsArgs
    ///         {
    ///             ProjectId = scaleway_mnq_sqs.Main.Project_id,
    ///             Queue = "MyQueue",
    ///             Region = scaleway_mnq_sqs.Main.Region,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### NATS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.ContainerTrigger("main", new()
    ///     {
    ///         ContainerId = scaleway_container.Main.Id,
    ///         Nats = new Scaleway.Inputs.ContainerTriggerNatsArgs
    ///         {
    ///             AccountId = scaleway_mnq_nats_account.Main.Id,
    ///             Subject = "MySubject",
    ///             Region = scaleway_mnq_nats_account.Main.Region,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Container Triggers can be imported using `{region}/{id}`, as shown below:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/containerTrigger:ContainerTrigger main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/containerTrigger:ContainerTrigger")]
    public partial class ContainerTrigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier of the container to create a trigger for.
        /// </summary>
        [Output("containerId")]
        public Output<string> ContainerId { get; private set; } = null!;

        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The unique name of the trigger. If not provided, a random name is generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The configuration for the Scaleway NATS account used by the trigger
        /// </summary>
        [Output("nats")]
        public Output<Outputs.ContainerTriggerNats?> Nats { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace is created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The configuration of the Scaleway SQS queue used by the trigger
        /// </summary>
        [Output("sqs")]
        public Output<Outputs.ContainerTriggerSqs?> Sqs { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerTrigger(string name, ContainerTriggerArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/containerTrigger:ContainerTrigger", name, args ?? new ContainerTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerTrigger(string name, Input<string> id, ContainerTriggerState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/containerTrigger:ContainerTrigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerTrigger Get(string name, Input<string> id, ContainerTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerTrigger(name, id, state, options);
        }
    }

    public sealed class ContainerTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container to create a trigger for.
        /// </summary>
        [Input("containerId", required: true)]
        public Input<string> ContainerId { get; set; } = null!;

        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique name of the trigger. If not provided, a random name is generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The configuration for the Scaleway NATS account used by the trigger
        /// </summary>
        [Input("nats")]
        public Input<Inputs.ContainerTriggerNatsArgs>? Nats { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace is created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The configuration of the Scaleway SQS queue used by the trigger
        /// </summary>
        [Input("sqs")]
        public Input<Inputs.ContainerTriggerSqsArgs>? Sqs { get; set; }

        public ContainerTriggerArgs()
        {
        }
        public static new ContainerTriggerArgs Empty => new ContainerTriggerArgs();
    }

    public sealed class ContainerTriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container to create a trigger for.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// The description of the trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique name of the trigger. If not provided, a random name is generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The configuration for the Scaleway NATS account used by the trigger
        /// </summary>
        [Input("nats")]
        public Input<Inputs.ContainerTriggerNatsGetArgs>? Nats { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace is created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The configuration of the Scaleway SQS queue used by the trigger
        /// </summary>
        [Input("sqs")]
        public Input<Inputs.ContainerTriggerSqsGetArgs>? Sqs { get; set; }

        public ContainerTriggerState()
        {
        }
        public static new ContainerTriggerState Empty => new ContainerTriggerState();
    }
}
