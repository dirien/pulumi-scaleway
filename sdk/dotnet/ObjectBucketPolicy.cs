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
    /// The `scaleway.ObjectBucketPolicy` resource allows you to create and manage bucket policies for [Scaleway Object storage](https://www.scaleway.com/en/docs/storage/object/).
    /// 
    /// Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/) for more information on Object Storage bucket policies.
    /// 
    /// ## Example Usage
    /// 
    /// ### Example Usage with an IAM user
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     var user = Scaleway.GetIamUser.Invoke(new()
    ///     {
    ///         Email = "user@scaleway.com",
    ///     });
    /// 
    ///     var policyIamPolicy = new Scaleway.IamPolicy("policyIamPolicy", new()
    ///     {
    ///         UserId = user.Apply(getIamUserResult =&gt; getIamUserResult.Id),
    ///         Rules = new[]
    ///         {
    ///             new Scaleway.Inputs.IamPolicyRuleArgs
    ///             {
    ///                 ProjectIds = new[]
    ///                 {
    ///                     @default.Apply(@default =&gt; @default.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id)),
    ///                 },
    ///                 PermissionSetNames = new[]
    ///                 {
    ///                     "ObjectStorageFullAccess",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Object storage configuration
    ///     var bucket = new Scaleway.ObjectBucket("bucket");
    /// 
    ///     var policyObjectBucketPolicy = new Scaleway.ObjectBucketPolicy("policyObjectBucketPolicy", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Policy = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2023-04-17",
    ///             ["Id"] = "MyBucketPolicy",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Effect"] = "Allow",
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "s3:*",
    ///                     },
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["SCW"] = $"user_id:{user.Apply(getIamUserResult =&gt; getIamUserResult.Id)}",
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         bucket.Name,
    ///                         bucket.Name.Apply(name =&gt; $"{name}/*"),
    ///                     },
    ///                 },
    ///             },
    ///         })),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with an IAM application
    /// 
    /// ### Creating a bucket and delegating read access to an application
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     // IAM configuration
    ///     var reading_app = new Scaleway.IamApplication("reading-app");
    /// 
    ///     var policyIamPolicy = new Scaleway.IamPolicy("policyIamPolicy", new()
    ///     {
    ///         ApplicationId = reading_app.Id,
    ///         Rules = new[]
    ///         {
    ///             new Scaleway.Inputs.IamPolicyRuleArgs
    ///             {
    ///                 ProjectIds = new[]
    ///                 {
    ///                     @default.Apply(@default =&gt; @default.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id)),
    ///                 },
    ///                 PermissionSetNames = new[]
    ///                 {
    ///                     "ObjectStorageBucketsRead",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Object storage configuration
    ///     var bucket = new Scaleway.ObjectBucket("bucket");
    /// 
    ///     var policyObjectBucketPolicy = new Scaleway.ObjectBucketPolicy("policyObjectBucketPolicy", new()
    ///     {
    ///         Bucket = bucket.Id,
    ///         Policy = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2023-04-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Sid"] = "Delegate read access",
    ///                     ["Effect"] = "Allow",
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["SCW"] = reading_app.Id.Apply(id =&gt; $"application_id:{id}"),
    ///                     },
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "s3:ListBucket",
    ///                         "s3:GetObject",
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         bucket.Name,
    ///                         bucket.Name.Apply(name =&gt; $"{name}/*"),
    ///                     },
    ///                 },
    ///             },
    ///         })),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Reading the bucket with the application
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var reading_app = Scaleway.GetIamApplication.Invoke(new()
    ///     {
    ///         Name = "reading-app",
    ///     });
    /// 
    ///     var reading_api_key = new Scaleway.IamApiKey("reading-api-key", new()
    ///     {
    ///         ApplicationId = reading_app.Apply(reading_app =&gt; reading_app.Apply(getIamApplicationResult =&gt; getIamApplicationResult.Id)),
    ///     });
    /// 
    ///     var reading_profile = new Scaleway.Provider("reading-profile", new()
    ///     {
    ///         AccessKey = reading_api_key.AccessKey,
    ///         SecretKey = reading_api_key.SecretKey,
    ///     });
    /// 
    ///     var bucket = Scaleway.GetObjectBucket.Invoke(new()
    ///     {
    ///         Name = "some-unique-name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with AWS provider
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     // Object storage configuration
    ///     var bucket = new Scaleway.ObjectBucket("bucket");
    /// 
    ///     var policy = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Version = "2012-10-17",
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Sid = "Delegate access",
    ///                 Effect = "Allow",
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "SCW",
    ///                         Identifiers = new[]
    ///                         {
    ///                             $"project_id:{@default.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id)}",
    ///                         },
    ///                     },
    ///                 },
    ///                 Actions = new[]
    ///                 {
    ///                     "s3:ListBucket",
    ///                 },
    ///                 Resources = new[]
    ///                 {
    ///                     bucket.Name,
    ///                     $"{bucket.Name}/*",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var main = new Scaleway.ObjectBucketPolicy("main", new()
    ///     {
    ///         Bucket = bucket.Id,
    ///         Policy = policy.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with deprecated version 2012-10-17
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Scaleway.GetAccountProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     // Object storage configuration
    ///     var bucket = new Scaleway.ObjectBucket("bucket", new()
    ///     {
    ///         Region = "fr-par",
    ///     });
    /// 
    ///     var policy = new Scaleway.ObjectBucketPolicy("policy", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Policy = Output.JsonSerialize(Output.Create(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Effect"] = "Allow",
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "s3:ListBucket",
    ///                         "s3:GetObjectTagging",
    ///                     },
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["SCW"] = @default.Apply(@default =&gt; $"project_id:{@default.Apply(getAccountProjectResult =&gt; getAccountProjectResult.Id)}"),
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         bucket.Name,
    ///                         bucket.Name.Apply(name =&gt; $"{name}/*"),
    ///                     },
    ///                 },
    ///             },
    ///         })),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// **NB:** To configure the AWS provider with Scaleway credentials, refer to the [dedicated documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/object-storage-aws-cli/).
    /// 
    /// ## Import
    /// 
    /// Bucket policies can be imported using the `{region}/{bucketName}` identifier, as shown below:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketPolicy:ObjectBucketPolicy some_bucket fr-par/some-bucket
    /// ```
    /// 
    /// ~&gt; **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
    /// 
    /// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketPolicy:ObjectBucketPolicy some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/objectBucketPolicy:ObjectBucketPolicy")]
    public partial class ObjectBucketPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The text of the policy.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectBucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectBucketPolicy(string name, ObjectBucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketPolicy:ObjectBucketPolicy", name, args ?? new ObjectBucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectBucketPolicy(string name, Input<string> id, ObjectBucketPolicyState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketPolicy:ObjectBucketPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectBucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectBucketPolicy Get(string name, Input<string> id, ObjectBucketPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectBucketPolicy(name, id, state, options);
        }
    }

    public sealed class ObjectBucketPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The text of the policy.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObjectBucketPolicyArgs()
        {
        }
        public static new ObjectBucketPolicyArgs Empty => new ObjectBucketPolicyArgs();
    }

    public sealed class ObjectBucketPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The text of the policy.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The Scaleway region this bucket resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObjectBucketPolicyState()
        {
        }
        public static new ObjectBucketPolicyState Empty => new ObjectBucketPolicyState();
    }
}
