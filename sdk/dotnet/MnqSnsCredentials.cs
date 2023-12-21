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
    /// Creates and manages Scaleway Messaging and queuing SNS Credentials.
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
    ///     var mainMnqSns = new Scaleway.MnqSns("mainMnqSns");
    /// 
    ///     var mainMnqSnsCredentials = new Scaleway.MnqSnsCredentials("mainMnqSnsCredentials", new()
    ///     {
    ///         ProjectId = mainMnqSns.ProjectId,
    ///         Permissions = new Scaleway.Inputs.MnqSnsCredentialsPermissionsArgs
    ///         {
    ///             CanManage = false,
    ///             CanReceive = true,
    ///             CanPublish = false,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SNS credentials can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/mnqSnsCredentials:MnqSnsCredentials main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqSnsCredentials:MnqSnsCredentials")]
    public partial class MnqSnsCredentials : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the key.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The unique name of the sns credentials.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Output("permissions")]
        public Output<Outputs.MnqSnsCredentialsPermissions> Permissions { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which sns is enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret value of the key.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;


        /// <summary>
        /// Create a MnqSnsCredentials resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqSnsCredentials(string name, MnqSnsCredentialsArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSnsCredentials:MnqSnsCredentials", name, args ?? new MnqSnsCredentialsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqSnsCredentials(string name, Input<string> id, MnqSnsCredentialsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqSnsCredentials:MnqSnsCredentials", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqSnsCredentials resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqSnsCredentials Get(string name, Input<string> id, MnqSnsCredentialsState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqSnsCredentials(name, id, state, options);
        }
    }

    public sealed class MnqSnsCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the sns credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.MnqSnsCredentialsPermissionsArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sns is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqSnsCredentialsArgs()
        {
        }
        public static new MnqSnsCredentialsArgs Empty => new MnqSnsCredentialsArgs();
    }

    public sealed class MnqSnsCredentialsState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The ID of the key.
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
        /// The unique name of the sns credentials.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// . List of permissions associated to these credentials. Only one of permissions may be set.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.MnqSnsCredentialsPermissionsGetArgs>? Permissions { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns is enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which sns is enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret value of the key.
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

        public MnqSnsCredentialsState()
        {
        }
        public static new MnqSnsCredentialsState Empty => new MnqSnsCredentialsState();
    }
}
