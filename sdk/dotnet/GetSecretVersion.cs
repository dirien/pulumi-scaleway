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
    public static class GetSecretVersion
    {
        /// <summary>
        /// The `scaleway.SecretVersion` data source is used to get information about a specific secret version stored in Scaleway Secret Manager.
        /// 
        /// Refer to the Secret Manager [product documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/) and [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/) for more information.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### Use Secret Manager
        /// 
        /// The following commands allow you to:
        /// 
        /// - create a secret named `fooii`
        /// - create a new version of `fooii` containing data (`your_secret`)
        /// - retrieve the secret version specified by the secret ID and the desired version
        /// - retrieve the secret version specified by the secret name and the desired version
        /// 
        /// The output blocks display the sensitive data contained in your secret version.
        /// 
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
        ///     // Create a secret named fooii
        ///     var mainSecret = new Scaleway.Secret("mainSecret", new()
        ///     {
        ///         Description = "barr",
        ///     });
        /// 
        ///     // Create a version of fooii containing data
        ///     var mainSecretVersion = new Scaleway.SecretVersion("mainSecretVersion", new()
        ///     {
        ///         Description = "your description",
        ///         SecretId = mainSecret.Id,
        ///         Data = "your_secret",
        ///     });
        /// 
        ///     var dataBySecretId = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretId = mainSecret.Id,
        ///         Revision = "1",
        ///     });
        /// 
        ///     var dataBySecretName = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretName = mainSecret.Name,
        ///         Revision = "1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["scalewaySecretAccessPayload"] = dataBySecretName.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///         ["scalewaySecretAccessPayloadById"] = dataBySecretId.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Data information
        /// 
        /// Note: This data source provides you with access to the secret payload, which is encoded in base64.
        /// 
        /// Keep in mind that this is a sensitive attribute. For more information,
        /// see Sensitive Data in State.
        /// 
        /// &gt; **Important:**  This property is sensitive and will not be displayed in the pulumi preview, for security reasons.
        /// </summary>
        public static Task<GetSecretVersionResult> InvokeAsync(GetSecretVersionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("scaleway:index/getSecretVersion:getSecretVersion", args ?? new GetSecretVersionArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.SecretVersion` data source is used to get information about a specific secret version stored in Scaleway Secret Manager.
        /// 
        /// Refer to the Secret Manager [product documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/) and [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/) for more information.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### Use Secret Manager
        /// 
        /// The following commands allow you to:
        /// 
        /// - create a secret named `fooii`
        /// - create a new version of `fooii` containing data (`your_secret`)
        /// - retrieve the secret version specified by the secret ID and the desired version
        /// - retrieve the secret version specified by the secret name and the desired version
        /// 
        /// The output blocks display the sensitive data contained in your secret version.
        /// 
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
        ///     // Create a secret named fooii
        ///     var mainSecret = new Scaleway.Secret("mainSecret", new()
        ///     {
        ///         Description = "barr",
        ///     });
        /// 
        ///     // Create a version of fooii containing data
        ///     var mainSecretVersion = new Scaleway.SecretVersion("mainSecretVersion", new()
        ///     {
        ///         Description = "your description",
        ///         SecretId = mainSecret.Id,
        ///         Data = "your_secret",
        ///     });
        /// 
        ///     var dataBySecretId = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretId = mainSecret.Id,
        ///         Revision = "1",
        ///     });
        /// 
        ///     var dataBySecretName = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretName = mainSecret.Name,
        ///         Revision = "1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["scalewaySecretAccessPayload"] = dataBySecretName.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///         ["scalewaySecretAccessPayloadById"] = dataBySecretId.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Data information
        /// 
        /// Note: This data source provides you with access to the secret payload, which is encoded in base64.
        /// 
        /// Keep in mind that this is a sensitive attribute. For more information,
        /// see Sensitive Data in State.
        /// 
        /// &gt; **Important:**  This property is sensitive and will not be displayed in the pulumi preview, for security reasons.
        /// </summary>
        public static Output<GetSecretVersionResult> Invoke(GetSecretVersionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretVersionResult>("scaleway:index/getSecretVersion:getSecretVersion", args ?? new GetSecretVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.SecretVersion` data source is used to get information about a specific secret version stored in Scaleway Secret Manager.
        /// 
        /// Refer to the Secret Manager [product documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/) and [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/) for more information.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### Use Secret Manager
        /// 
        /// The following commands allow you to:
        /// 
        /// - create a secret named `fooii`
        /// - create a new version of `fooii` containing data (`your_secret`)
        /// - retrieve the secret version specified by the secret ID and the desired version
        /// - retrieve the secret version specified by the secret name and the desired version
        /// 
        /// The output blocks display the sensitive data contained in your secret version.
        /// 
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
        ///     // Create a secret named fooii
        ///     var mainSecret = new Scaleway.Secret("mainSecret", new()
        ///     {
        ///         Description = "barr",
        ///     });
        /// 
        ///     // Create a version of fooii containing data
        ///     var mainSecretVersion = new Scaleway.SecretVersion("mainSecretVersion", new()
        ///     {
        ///         Description = "your description",
        ///         SecretId = mainSecret.Id,
        ///         Data = "your_secret",
        ///     });
        /// 
        ///     var dataBySecretId = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretId = mainSecret.Id,
        ///         Revision = "1",
        ///     });
        /// 
        ///     var dataBySecretName = Scaleway.GetSecretVersion.Invoke(new()
        ///     {
        ///         SecretName = mainSecret.Name,
        ///         Revision = "1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["scalewaySecretAccessPayload"] = dataBySecretName.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///         ["scalewaySecretAccessPayloadById"] = dataBySecretId.Apply(getSecretVersionResult =&gt; getSecretVersionResult.Data),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Data information
        /// 
        /// Note: This data source provides you with access to the secret payload, which is encoded in base64.
        /// 
        /// Keep in mind that this is a sensitive attribute. For more information,
        /// see Sensitive Data in State.
        /// 
        /// &gt; **Important:**  This property is sensitive and will not be displayed in the pulumi preview, for security reasons.
        /// </summary>
        public static Output<GetSecretVersionResult> Invoke(GetSecretVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretVersionResult>("scaleway:index/getSecretVersion:getSecretVersion", args ?? new GetSecretVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// The ID of the Scaleway Project associated with the secret version.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The revision for this secret version. Refer to alternative values (ex: `latest`) in the [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/#path-secret-versions-access-a-secrets-version-using-the-secrets-id)
        /// </summary>
        [Input("revision")]
        public string? Revision { get; set; }

        /// <summary>
        /// The ID of the secret associated with the secret version. Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretId")]
        public string? SecretId { get; set; }

        /// <summary>
        /// The name of the secret associated with the secret version.
        /// Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretName")]
        public string? SecretName { get; set; }

        public GetSecretVersionArgs()
        {
        }
        public static new GetSecretVersionArgs Empty => new GetSecretVersionArgs();
    }

    public sealed class GetSecretVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The ID of the Scaleway Project associated with the secret version.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The revision for this secret version. Refer to alternative values (ex: `latest`) in the [API documentation](https://www.scaleway.com/en/developers/api/secret-manager/#path-secret-versions-access-a-secrets-version-using-the-secrets-id)
        /// </summary>
        [Input("revision")]
        public Input<string>? Revision { get; set; }

        /// <summary>
        /// The ID of the secret associated with the secret version. Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        /// <summary>
        /// The name of the secret associated with the secret version.
        /// Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public GetSecretVersionInvokeArgs()
        {
        }
        public static new GetSecretVersionInvokeArgs Empty => new GetSecretVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretVersionResult
    {
        /// <summary>
        /// The date and time of the secret version's creation in RFC 3339 format.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The data payload of the secret version. This is a sensitive attribute containing the secret value. Learn more in the [data section](https://www.terraform.io/#data-information).
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// (Optional) The description of the secret version (e.g. `my-new-description`).
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly string? Revision;
        public readonly string? SecretId;
        public readonly string? SecretName;
        /// <summary>
        /// The status of the secret version.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The date and time of the secret version's last update in RFC 3339 format.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetSecretVersionResult(
            string createdAt,

            string data,

            string description,

            string id,

            string organizationId,

            string? projectId,

            string? region,

            string? revision,

            string? secretId,

            string? secretName,

            string status,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Data = data;
            Description = description;
            Id = id;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Revision = revision;
            SecretId = secretId;
            SecretName = secretName;
            Status = status;
            UpdatedAt = updatedAt;
        }
    }
}
