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
    public static class GetIamApplication
    {
        /// <summary>
        /// Gets information about an existing IAM application.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetIamApplication.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var findById = Scaleway.GetIamApplication.Invoke(new()
        ///     {
        ///         ApplicationId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIamApplicationResult> InvokeAsync(GetIamApplicationArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamApplicationResult>("scaleway:index/getIamApplication:getIamApplication", args ?? new GetIamApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an existing IAM application.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetIamApplication.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var findById = Scaleway.GetIamApplication.Invoke(new()
        ///     {
        ///         ApplicationId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamApplicationResult> Invoke(GetIamApplicationInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamApplicationResult>("scaleway:index/getIamApplication:getIamApplication", args ?? new GetIamApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the IAM application.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `application_id`.
        /// </summary>
        [Input("applicationId")]
        public string? ApplicationId { get; set; }

        /// <summary>
        /// The name of the IAM application.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the
        /// Organization the application is associated with.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        public GetIamApplicationArgs()
        {
        }
        public static new GetIamApplicationArgs Empty => new GetIamApplicationArgs();
    }

    public sealed class GetIamApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the IAM application.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `application_id`.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The name of the IAM application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `organization_id`) The ID of the
        /// Organization the application is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public GetIamApplicationInvokeArgs()
        {
        }
        public static new GetIamApplicationInvokeArgs Empty => new GetIamApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamApplicationResult
    {
        public readonly string? ApplicationId;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly bool Editable;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? OrganizationId;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetIamApplicationResult(
            string? applicationId,

            string createdAt,

            string description,

            bool editable,

            string id,

            string? name,

            string? organizationId,

            ImmutableArray<string> tags,

            string updatedAt)
        {
            ApplicationId = applicationId;
            CreatedAt = createdAt;
            Description = description;
            Editable = editable;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
