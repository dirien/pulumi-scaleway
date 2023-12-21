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
    /// Manage Scaleway Messaging and queuing SNS Topic Subscriptions.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
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
    ///     var mainMnqSns = new Scaleway.MnqSns("mainMnqSns");
    /// 
    ///     var mainMnqSnsCredentials = new Scaleway.MnqSnsCredentials("mainMnqSnsCredentials", new()
    ///     {
    ///         ProjectId = mainMnqSns.ProjectId,
    ///         Permissions = new Scaleway.Inputs.MnqSnsCredentialsPermissionsArgs
    ///         {
    ///             CanManage = true,
    ///             CanPublish = true,
    ///             CanReceive = true,
    ///         },
    ///     });
    /// 
    ///     var topic = new Scaleway.MnqSnsTopic("topic", new()
    ///     {
    ///         ProjectId = mainMnqSns.ProjectId,
    ///         AccessKey = mainMnqSnsCredentials.AccessKey,
    ///         SecretKey = mainMnqSnsCredentials.SecretKey,
    ///     });
    /// 
    ///     var mainMnqSnsTopicSubscription = new Scaleway.MnqSnsTopicSubscription("mainMnqSnsTopicSubscription", new()
    ///     {
    ///         ProjectId = mainMnqSns.ProjectId,
    ///         AccessKey = mainMnqSnsCredentials.AccessKey,
    ///         SecretKey = mainMnqSnsCredentials.SecretKey,
    ///         TopicId = topic.Id,
    ///         Protocol = "http",
    ///         Endpoint = "http://example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SNS topic subscriptions can be imported using the `{region}/{project-id}/{topic-name}/{subscription-id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription main fr-par/11111111111111111111111111111111/my-topic/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription")]
    public partial class MnqSnsTopicSubscription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access key of the SNS credentials.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The ARN of the topic subscription
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Endpoint of the subscription
        /// </summary>
        [Output("endpoint")]
        public Output<string?> Endpoint { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Protocol of the SNS Topic Subscription.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Activate JSON Redrive Policy.
        /// </summary>
        [Output("redrivePolicy")]
        public Output<bool> RedrivePolicy { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which sns is enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret key of the SNS credentials.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        /// </summary>
        [Output("snsEndpoint")]
        public Output<string?> SnsEndpoint { get; private set; } = null!;

        /// <summary>
        /// The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        /// </summary>
        [Output("topicArn")]
        public Output<string?> TopicArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        /// </summary>
        [Output("topicId")]
        public Output<string?> TopicId { get; private set; } = null!;


        /// <summary>
        /// Create a MnqSnsTopicSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqSnsTopicSubscription(string name, MnqSnsTopicSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription", name, args ?? new MnqSnsTopicSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqSnsTopicSubscription(string name, Input<string> id, MnqSnsTopicSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-scaleway",
                AdditionalSecretOutputs =
                {
                    "accessKey",
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MnqSnsTopicSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqSnsTopicSubscription Get(string name, Input<string> id, MnqSnsTopicSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqSnsTopicSubscription(name, id, state, options);
        }
    }

    public sealed class MnqSnsTopicSubscriptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKey", required: true)]
        private Input<string>? _accessKey;

        /// <summary>
        /// The access key of the SNS credentials.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Endpoint of the subscription
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Protocol of the SNS Topic Subscription.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Activate JSON Redrive Policy.
        /// </summary>
        [Input("redrivePolicy")]
        public Input<bool>? RedrivePolicy { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret key of the SNS credentials.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        /// </summary>
        [Input("snsEndpoint")]
        public Input<string>? SnsEndpoint { get; set; }

        /// <summary>
        /// The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        /// <summary>
        /// The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public MnqSnsTopicSubscriptionArgs()
        {
        }
        public static new MnqSnsTopicSubscriptionArgs Empty => new MnqSnsTopicSubscriptionArgs();
    }

    public sealed class MnqSnsTopicSubscriptionState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The access key of the SNS credentials.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ARN of the topic subscription
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Endpoint of the subscription
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Protocol of the SNS Topic Subscription.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Activate JSON Redrive Policy.
        /// </summary>
        [Input("redrivePolicy")]
        public Input<bool>? RedrivePolicy { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret key of the SNS credentials.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        /// </summary>
        [Input("snsEndpoint")]
        public Input<string>? SnsEndpoint { get; set; }

        /// <summary>
        /// The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        /// </summary>
        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        /// <summary>
        /// The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public MnqSnsTopicSubscriptionState()
        {
        }
        public static new MnqSnsTopicSubscriptionState Empty => new MnqSnsTopicSubscriptionState();
    }
}
