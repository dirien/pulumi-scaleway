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
    /// Creates and manages Scaleway IAM API Keys. For more information, please
    /// check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#api-keys-3665ae)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ciCd = new Scaleway.IamApplication("ciCd");
    /// 
    ///     var main = new Scaleway.IamApiKey("main", new()
    ///     {
    ///         ApplicationId = scaleway_iam_application.Main.Id,
    ///         Description = "a description",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api keys can be imported using the `{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/iamApiKey:IamApiKey main 11111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iamApiKey:IamApiKey")]
    public partial class IamApiKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access key of the iam api key.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// ID of the application attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the iam api key.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The IP Address of the device which created the API key.
        /// </summary>
        [Output("creationIp")]
        public Output<string> CreationIp { get; private set; } = null!;

        /// <summary>
        /// The default project ID to use with object storage.
        /// </summary>
        [Output("defaultProjectId")]
        public Output<string> DefaultProjectId { get; private set; } = null!;

        /// <summary>
        /// The description of the iam api key.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the iam api key is editable.
        /// </summary>
        [Output("editable")]
        public Output<bool> Editable { get; private set; } = null!;

        /// <summary>
        /// The date and time of the expiration of the iam api key. Please note that in case of change,
        /// the resource will be recreated.
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The secret Key of the iam api key.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the iam api key.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// ID of the user attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a IamApiKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamApiKey(string name, IamApiKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamApiKey:IamApiKey", name, args ?? new IamApiKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamApiKey(string name, Input<string> id, IamApiKeyState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamApiKey:IamApiKey", name, state, MakeResourceOptions(options, id))
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
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IamApiKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamApiKey Get(string name, Input<string> id, IamApiKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamApiKey(name, id, state, options);
        }
    }

    public sealed class IamApiKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the application attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The default project ID to use with object storage.
        /// </summary>
        [Input("defaultProjectId")]
        public Input<string>? DefaultProjectId { get; set; }

        /// <summary>
        /// The description of the iam api key.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The date and time of the expiration of the iam api key. Please note that in case of change,
        /// the resource will be recreated.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// ID of the user attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamApiKeyArgs()
        {
        }
        public static new IamApiKeyArgs Empty => new IamApiKeyArgs();
    }

    public sealed class IamApiKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key of the iam api key.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// ID of the application attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The date and time of the creation of the iam api key.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The IP Address of the device which created the API key.
        /// </summary>
        [Input("creationIp")]
        public Input<string>? CreationIp { get; set; }

        /// <summary>
        /// The default project ID to use with object storage.
        /// </summary>
        [Input("defaultProjectId")]
        public Input<string>? DefaultProjectId { get; set; }

        /// <summary>
        /// The description of the iam api key.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the iam api key is editable.
        /// </summary>
        [Input("editable")]
        public Input<bool>? Editable { get; set; }

        /// <summary>
        /// The date and time of the expiration of the iam api key. Please note that in case of change,
        /// the resource will be recreated.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The secret Key of the iam api key.
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
        /// The date and time of the last update of the iam api key.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// ID of the user attached to the api key.
        /// Only one of the `application_id` and `user_id` should be specified.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamApiKeyState()
        {
        }
        public static new IamApiKeyState Empty => new IamApiKeyState();
    }
}
