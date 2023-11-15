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
    /// Creates and manages Scaleway object storage buckets.
    /// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
    ///     var someBucket = new Scaleway.ObjectBucket("someBucket", new()
    ///     {
    ///         Tags = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Creating the bucket in a specific project
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var someBucket = new Scaleway.ObjectBucket("someBucket", new()
    ///     {
    ///         ProjectId = "11111111-1111-1111-1111-111111111111",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Using object lifecycle
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.ObjectBucket("main", new()
    ///     {
    ///         LifecycleRules = new[]
    ///         {
    ///             new Scaleway.Inputs.ObjectBucketLifecycleRuleArgs
    ///             {
    ///                 Enabled = true,
    ///                 Expiration = new Scaleway.Inputs.ObjectBucketLifecycleRuleExpirationArgs
    ///                 {
    ///                     Days = 365,
    ///                 },
    ///                 Id = "id1",
    ///                 Prefix = "path1/",
    ///                 Transitions = new[]
    ///                 {
    ///                     new Scaleway.Inputs.ObjectBucketLifecycleRuleTransitionArgs
    ///                     {
    ///                         Days = 120,
    ///                         StorageClass = "GLACIER",
    ///                     },
    ///                 },
    ///             },
    ///             new Scaleway.Inputs.ObjectBucketLifecycleRuleArgs
    ///             {
    ///                 Enabled = true,
    ///                 Expiration = new Scaleway.Inputs.ObjectBucketLifecycleRuleExpirationArgs
    ///                 {
    ///                     Days = 50,
    ///                 },
    ///                 Id = "id2",
    ///                 Prefix = "path2/",
    ///             },
    ///             new Scaleway.Inputs.ObjectBucketLifecycleRuleArgs
    ///             {
    ///                 Enabled = false,
    ///                 Expiration = new Scaleway.Inputs.ObjectBucketLifecycleRuleExpirationArgs
    ///                 {
    ///                     Days = 1,
    ///                 },
    ///                 Id = "id3",
    ///                 Prefix = "path3/",
    ///                 Tags = 
    ///                 {
    ///                     { "tagKey", "tagValue" },
    ///                     { "terraform", "hashicorp" },
    ///                 },
    ///             },
    ///             new Scaleway.Inputs.ObjectBucketLifecycleRuleArgs
    ///             {
    ///                 Enabled = true,
    ///                 Id = "id4",
    ///                 Tags = 
    ///                 {
    ///                     { "tag1", "value1" },
    ///                 },
    ///                 Transitions = new[]
    ///                 {
    ///                     new Scaleway.Inputs.ObjectBucketLifecycleRuleTransitionArgs
    ///                     {
    ///                         Days = 0,
    ///                         StorageClass = "GLACIER",
    ///                     },
    ///                 },
    ///             },
    ///             new Scaleway.Inputs.ObjectBucketLifecycleRuleArgs
    ///             {
    ///                 AbortIncompleteMultipartUploadDays = 30,
    ///                 Enabled = true,
    ///             },
    ///         },
    ///         Region = "fr-par",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket
    /// ```
    /// 
    ///  If you are importing a bucket from a specific project (that is not your default project), you can use the following syntaxbash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/objectBucket:ObjectBucket")]
    public partial class ObjectBucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Deprecated) The canned ACL you want to apply to the bucket.
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        /// <summary>
        /// API URL of the bucket
        /// </summary>
        [Output("apiEndpoint")]
        public Output<string> ApiEndpoint { get; private set; } = null!;

        /// <summary>
        /// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        /// </summary>
        [Output("corsRules")]
        public Output<ImmutableArray<Outputs.ObjectBucketCorsRule>> CorsRules { get; private set; } = null!;

        /// <summary>
        /// The endpoint URL of the bucket
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
        /// </summary>
        [Output("lifecycleRules")]
        public Output<ImmutableArray<Outputs.ObjectBucketLifecycleRule>> LifecycleRules { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable object lock
        /// </summary>
        [Output("objectLockEnabled")]
        public Output<bool?> ObjectLockEnabled { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the bucket is associated with.
        /// 
        /// The `acl` attribute is deprecated. See scaleway.ObjectBucketAcl resource documentation.
        /// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of tags (key / value) for the bucket.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        /// </summary>
        [Output("versioning")]
        public Output<Outputs.ObjectBucketVersioning> Versioning { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectBucket(string name, ObjectBucketArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucket:ObjectBucket", name, args ?? new ObjectBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectBucket(string name, Input<string> id, ObjectBucketState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucket:ObjectBucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectBucket Get(string name, Input<string> id, ObjectBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectBucket(name, id, state, options);
        }
    }

    public sealed class ObjectBucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Deprecated) The canned ACL you want to apply to the bucket.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("corsRules")]
        private InputList<Inputs.ObjectBucketCorsRuleArgs>? _corsRules;

        /// <summary>
        /// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        /// </summary>
        public InputList<Inputs.ObjectBucketCorsRuleArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.ObjectBucketCorsRuleArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.ObjectBucketLifecycleRuleArgs>? _lifecycleRules;

        /// <summary>
        /// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
        /// </summary>
        public InputList<Inputs.ObjectBucketLifecycleRuleArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.ObjectBucketLifecycleRuleArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable object lock
        /// </summary>
        [Input("objectLockEnabled")]
        public Input<bool>? ObjectLockEnabled { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the bucket is associated with.
        /// 
        /// The `acl` attribute is deprecated. See scaleway.ObjectBucketAcl resource documentation.
        /// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of tags (key / value) for the bucket.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.ObjectBucketVersioningArgs>? Versioning { get; set; }

        public ObjectBucketArgs()
        {
        }
        public static new ObjectBucketArgs Empty => new ObjectBucketArgs();
    }

    public sealed class ObjectBucketState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Deprecated) The canned ACL you want to apply to the bucket.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// API URL of the bucket
        /// </summary>
        [Input("apiEndpoint")]
        public Input<string>? ApiEndpoint { get; set; }

        [Input("corsRules")]
        private InputList<Inputs.ObjectBucketCorsRuleGetArgs>? _corsRules;

        /// <summary>
        /// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
        /// </summary>
        public InputList<Inputs.ObjectBucketCorsRuleGetArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.ObjectBucketCorsRuleGetArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// The endpoint URL of the bucket
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.ObjectBucketLifecycleRuleGetArgs>? _lifecycleRules;

        /// <summary>
        /// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
        /// </summary>
        public InputList<Inputs.ObjectBucketLifecycleRuleGetArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.ObjectBucketLifecycleRuleGetArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable object lock
        /// </summary>
        [Input("objectLockEnabled")]
        public Input<bool>? ObjectLockEnabled { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the bucket is associated with.
        /// 
        /// The `acl` attribute is deprecated. See scaleway.ObjectBucketAcl resource documentation.
        /// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A list of tags (key / value) for the bucket.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.ObjectBucketVersioningGetArgs>? Versioning { get; set; }

        public ObjectBucketState()
        {
        }
        public static new ObjectBucketState Empty => new ObjectBucketState();
    }
}
