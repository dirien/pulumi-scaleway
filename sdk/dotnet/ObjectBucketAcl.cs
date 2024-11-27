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
    ///     var someBucket = new Scaleway.ObjectBucket("someBucket");
    /// 
    ///     var main = new Scaleway.ObjectBucketAcl("main", new()
    ///     {
    ///         Bucket = scaleway_object_bucket.Main.Id,
    ///         Acl = "private",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// For more information, refer to the [PutBucketAcl API call documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-operations/#putbucketacl).
    /// 
    /// ### With Grants
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainObjectBucket = new Scaleway.ObjectBucket("mainObjectBucket");
    /// 
    ///     var mainObjectBucketAcl = new Scaleway.ObjectBucketAcl("mainObjectBucketAcl", new()
    ///     {
    ///         Bucket = mainObjectBucket.Id,
    ///         AccessControlPolicy = new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyArgs
    ///         {
    ///             Grants = new[]
    ///             {
    ///                 new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyGrantArgs
    ///                 {
    ///                     Grantee = new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyGrantGranteeArgs
    ///                     {
    ///                         Id = "&lt;project-id&gt;:&lt;project-id&gt;",
    ///                         Type = "CanonicalUser",
    ///                     },
    ///                     Permission = "FULL_CONTROL",
    ///                 },
    ///                 new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyGrantArgs
    ///                 {
    ///                     Grantee = new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyGrantGranteeArgs
    ///                     {
    ///                         Id = "&lt;project-id&gt;",
    ///                         Type = "CanonicalUser",
    ///                     },
    ///                     Permission = "WRITE",
    ///                 },
    ///             },
    ///             Owner = new Scaleway.Inputs.ObjectBucketAclAccessControlPolicyOwnerArgs
    ///             {
    ///                 Id = "&lt;project-id&gt;",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## The ACL
    /// 
    /// Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
    /// 
    /// ## The access control policy
    /// 
    /// The `access_control_policy` configuration block supports the following arguments:
    /// 
    /// * `grant` - (Required) Set of grant configuration blocks documented below.
    /// * `owner` - (Required) Configuration block of the bucket owner's display name and ID documented below.
    /// 
    /// ## The grant
    /// 
    /// The `grant` configuration block supports the following arguments:
    /// 
    /// * `grantee` - (Required) Configuration block for the project being granted permissions documented below.
    /// * `permission` - (Required) Logging permissions assigned to the grantee for the bucket.
    /// 
    /// ## The permission
    /// 
    /// The following list shows each access policy permissions supported.
    /// 
    /// `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`
    /// 
    /// For more information about ACL permissions in the S3 bucket, see [ACL permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html).
    /// 
    /// ## The owner
    /// 
    /// The `owner` configuration block supports the following arguments:
    /// 
    /// * `id` - (Required) The ID of the project owner.
    /// * `display_name` - (Optional) The display name of the owner.
    /// 
    /// ## the grantee
    /// 
    /// The `grantee` configuration block supports the following arguments:
    /// 
    /// * `id` - (Required) The canonical user ID of the grantee.
    /// * `type` - (Required) Type of grantee. Valid values: CanonicalUser.
    /// 
    /// ## Import
    /// 
    /// Bucket ACLs can be imported using the `{region}/{bucketName}/{acl}` identifier, as shown below:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private
    /// ```
    /// 
    /// ~&gt; **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
    /// 
    /// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/objectBucketAcl:ObjectBucketAcl")]
    public partial class ObjectBucketAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A configuration block that sets the ACL permissions for an object per grantee documented below.
        /// </summary>
        [Output("accessControlPolicy")]
        public Output<Outputs.ObjectBucketAclAccessControlPolicy> AccessControlPolicy { get; private set; } = null!;

        /// <summary>
        /// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The project ID of the expected bucket owner.
        /// </summary>
        [Output("expectedBucketOwner")]
        public Output<string?> ExpectedBucketOwner { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectBucketAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectBucketAcl(string name, ObjectBucketAclArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketAcl:ObjectBucketAcl", name, args ?? new ObjectBucketAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectBucketAcl(string name, Input<string> id, ObjectBucketAclState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketAcl:ObjectBucketAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectBucketAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectBucketAcl Get(string name, Input<string> id, ObjectBucketAclState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectBucketAcl(name, id, state, options);
        }
    }

    public sealed class ObjectBucketAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A configuration block that sets the ACL permissions for an object per grantee documented below.
        /// </summary>
        [Input("accessControlPolicy")]
        public Input<Inputs.ObjectBucketAclAccessControlPolicyArgs>? AccessControlPolicy { get; set; }

        /// <summary>
        /// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The project ID of the expected bucket owner.
        /// </summary>
        [Input("expectedBucketOwner")]
        public Input<string>? ExpectedBucketOwner { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObjectBucketAclArgs()
        {
        }
        public static new ObjectBucketAclArgs Empty => new ObjectBucketAclArgs();
    }

    public sealed class ObjectBucketAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A configuration block that sets the ACL permissions for an object per grantee documented below.
        /// </summary>
        [Input("accessControlPolicy")]
        public Input<Inputs.ObjectBucketAclAccessControlPolicyGetArgs>? AccessControlPolicy { get; set; }

        /// <summary>
        /// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The project ID of the expected bucket owner.
        /// </summary>
        [Input("expectedBucketOwner")]
        public Input<string>? ExpectedBucketOwner { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ObjectBucketAclState()
        {
        }
        public static new ObjectBucketAclState Empty => new ObjectBucketAclState();
    }
}
