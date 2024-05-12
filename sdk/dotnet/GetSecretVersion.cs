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
    public static class GetSecretVersion
    {
        /// <summary>
        /// Gets information about Scaleway a Secret Version.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).
        /// 
        /// ## Examples
        /// 
        /// ### Basic
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
        ///     var mainSecret = new Scaleway.Secret("mainSecret", new()
        ///     {
        ///         Description = "barr",
        ///     });
        /// 
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
        /// ## Data
        /// 
        /// Note: This Data Source give you **access** to the secret payload encoded en base64.
        /// 
        /// Be aware that this is a sensitive attribute. For more information,
        /// see Sensitive Data in State.
        /// 
        /// &gt; **Important:**  This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public static Task<GetSecretVersionResult> InvokeAsync(GetSecretVersionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretVersionResult>("scaleway:index/getSecretVersion:getSecretVersion", args ?? new GetSecretVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about Scaleway a Secret Version.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).
        /// 
        /// ## Examples
        /// 
        /// ### Basic
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
        ///     var mainSecret = new Scaleway.Secret("mainSecret", new()
        ///     {
        ///         Description = "barr",
        ///     });
        /// 
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
        /// ## Data
        /// 
        /// Note: This Data Source give you **access** to the secret payload encoded en base64.
        /// 
        /// Be aware that this is a sensitive attribute. For more information,
        /// see Sensitive Data in State.
        /// 
        /// &gt; **Important:**  This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public static Output<GetSecretVersionResult> Invoke(GetSecretVersionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretVersionResult>("scaleway:index/getSecretVersion:getSecretVersion", args ?? new GetSecretVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project the Secret version is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the resource exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The revision for this Secret Version.
        /// </summary>
        [Input("revision")]
        public string? Revision { get; set; }

        /// <summary>
        /// The Secret ID associated wit the secret version.
        /// Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretId")]
        public string? SecretId { get; set; }

        /// <summary>
        /// The Name of Secret associated wit the secret version.
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
        /// <summary>
        /// The ID of the project the Secret version is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The revision for this Secret Version.
        /// </summary>
        [Input("revision")]
        public Input<string>? Revision { get; set; }

        /// <summary>
        /// The Secret ID associated wit the secret version.
        /// Only one of `secret_id` and `secret_name` should be specified.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        /// <summary>
        /// The Name of Secret associated wit the secret version.
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
        /// Date and time of secret version's creation (RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The data payload of the secret version. more on the data section
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// (Optional) Description of the secret version (e.g. `my-new-description`).
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly string? Revision;
        public readonly string? SecretId;
        public readonly string? SecretName;
        /// <summary>
        /// The status of the Secret Version.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Date and time of secret version's last update (RFC 3339 format).
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetSecretVersionResult(
            string createdAt,

            string data,

            string description,

            string id,

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
