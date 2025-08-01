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
    /// ## Import
    /// 
    /// This section explains how to import a secret using the `{region}/{id}` format.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/secret:Secret main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/secret:Secret")]
    public partial class Secret : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date and time of the secret's creation (in RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the secret (e.g. `my-new-description`).
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Ephemeral policy of the secret. Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
        /// </summary>
        [Output("ephemeralPolicies")]
        public Output<ImmutableArray<Outputs.SecretEphemeralPolicy>> EphemeralPolicies { get; private set; } = null!;

        /// <summary>
        /// Name of the secret (e.g. `my-secret`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Path of the secret, defaults to `/`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The project ID containing is the secret.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// True if secret protection is enabled on a given secret. A protected secret cannot be deleted.
        /// </summary>
        [Output("protected")]
        public Output<bool?> Protected { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The status of the secret.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags of the secret (e.g. `["tag", "secret"]`).
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the secret. If not specified, the type is Opaque. Available values can be found in [SDK Constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// Date and time of the secret's last update (in RFC 3339 format).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The amount of secret versions.
        /// </summary>
        [Output("versionCount")]
        public Output<int> VersionCount { get; private set; } = null!;

        [Output("versions")]
        public Output<ImmutableArray<Outputs.SecretVersion>> Versions { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/secret:Secret", name, args ?? new SecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/secret:Secret", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the secret (e.g. `my-new-description`).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ephemeralPolicies")]
        private InputList<Inputs.SecretEphemeralPolicyArgs>? _ephemeralPolicies;

        /// <summary>
        /// Ephemeral policy of the secret. Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
        /// </summary>
        public InputList<Inputs.SecretEphemeralPolicyArgs> EphemeralPolicies
        {
            get => _ephemeralPolicies ?? (_ephemeralPolicies = new InputList<Inputs.SecretEphemeralPolicyArgs>());
            set => _ephemeralPolicies = value;
        }

        /// <summary>
        /// Name of the secret (e.g. `my-secret`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path of the secret, defaults to `/`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The project ID containing is the secret.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// True if secret protection is enabled on a given secret. A protected secret cannot be deleted.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags of the secret (e.g. `["tag", "secret"]`).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of the secret. If not specified, the type is Opaque. Available values can be found in [SDK Constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SecretArgs()
        {
        }
        public static new SecretArgs Empty => new SecretArgs();
    }

    public sealed class SecretState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date and time of the secret's creation (in RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the secret (e.g. `my-new-description`).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ephemeralPolicies")]
        private InputList<Inputs.SecretEphemeralPolicyGetArgs>? _ephemeralPolicies;

        /// <summary>
        /// Ephemeral policy of the secret. Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
        /// </summary>
        public InputList<Inputs.SecretEphemeralPolicyGetArgs> EphemeralPolicies
        {
            get => _ephemeralPolicies ?? (_ephemeralPolicies = new InputList<Inputs.SecretEphemeralPolicyGetArgs>());
            set => _ephemeralPolicies = value;
        }

        /// <summary>
        /// Name of the secret (e.g. `my-secret`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path of the secret, defaults to `/`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The project ID containing is the secret.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// True if secret protection is enabled on a given secret. A protected secret cannot be deleted.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The status of the secret.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags of the secret (e.g. `["tag", "secret"]`).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of the secret. If not specified, the type is Opaque. Available values can be found in [SDK Constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Date and time of the secret's last update (in RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The amount of secret versions.
        /// </summary>
        [Input("versionCount")]
        public Input<int>? VersionCount { get; set; }

        [Input("versions")]
        private InputList<Inputs.SecretVersionGetArgs>? _versions;
        public InputList<Inputs.SecretVersionGetArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.SecretVersionGetArgs>());
            set => _versions = value;
        }

        public SecretState()
        {
        }
        public static new SecretState Empty => new SecretState();
    }
}
