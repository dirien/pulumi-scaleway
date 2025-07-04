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
    public static class GetAccountProjects
    {
        /// <summary>
        /// The `scaleway.getAccountProjects` data source is used to list all Scaleway projects in an Organization.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/) and [API documentation](https://www.scaleway.com/en/developers/api/account/project-api/) for more information.
        /// 
        /// 
        /// ## Retrieve a Scaleway Projects
        /// 
        /// The following commands allow you to:
        /// 
        /// - retrieve all Projects in an Organization
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Scaleway.GetAccountProjects.Invoke(new()
        ///     {
        ///         OrganizationId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Usage
        /// 
        /// ### Deploy an SSH key in all your organization's projects
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
        ///     var all = Scaleway.GetAccountProjects.Invoke();
        /// 
        ///     var main = new List&lt;Scaleway.AccountSshKey&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects).Length; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         main.Add(new Scaleway.AccountSshKey($"main-{range.Value}", new()
        ///         {
        ///             PublicKey = local.Public_key,
        ///             ProjectId = all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects)[range.Value].Id,
        ///         }));
        ///     }
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccountProjectsResult> InvokeAsync(GetAccountProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountProjectsResult>("scaleway:index/getAccountProjects:getAccountProjects", args ?? new GetAccountProjectsArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.getAccountProjects` data source is used to list all Scaleway projects in an Organization.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/) and [API documentation](https://www.scaleway.com/en/developers/api/account/project-api/) for more information.
        /// 
        /// 
        /// ## Retrieve a Scaleway Projects
        /// 
        /// The following commands allow you to:
        /// 
        /// - retrieve all Projects in an Organization
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Scaleway.GetAccountProjects.Invoke(new()
        ///     {
        ///         OrganizationId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Usage
        /// 
        /// ### Deploy an SSH key in all your organization's projects
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
        ///     var all = Scaleway.GetAccountProjects.Invoke();
        /// 
        ///     var main = new List&lt;Scaleway.AccountSshKey&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects).Length; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         main.Add(new Scaleway.AccountSshKey($"main-{range.Value}", new()
        ///         {
        ///             PublicKey = local.Public_key,
        ///             ProjectId = all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects)[range.Value].Id,
        ///         }));
        ///     }
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccountProjectsResult> Invoke(GetAccountProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountProjectsResult>("scaleway:index/getAccountProjects:getAccountProjects", args ?? new GetAccountProjectsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.getAccountProjects` data source is used to list all Scaleway projects in an Organization.
        /// 
        /// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/) and [API documentation](https://www.scaleway.com/en/developers/api/account/project-api/) for more information.
        /// 
        /// 
        /// ## Retrieve a Scaleway Projects
        /// 
        /// The following commands allow you to:
        /// 
        /// - retrieve all Projects in an Organization
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Scaleway.GetAccountProjects.Invoke(new()
        ///     {
        ///         OrganizationId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example Usage
        /// 
        /// ### Deploy an SSH key in all your organization's projects
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
        ///     var all = Scaleway.GetAccountProjects.Invoke();
        /// 
        ///     var main = new List&lt;Scaleway.AccountSshKey&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects).Length; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         main.Add(new Scaleway.AccountSshKey($"main-{range.Value}", new()
        ///         {
        ///             PublicKey = local.Public_key,
        ///             ProjectId = all.Apply(getAccountProjectsResult =&gt; getAccountProjectsResult.Projects)[range.Value].Id,
        ///         }));
        ///     }
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccountProjectsResult> Invoke(GetAccountProjectsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountProjectsResult>("scaleway:index/getAccountProjects:getAccountProjects", args ?? new GetAccountProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountProjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Organization with which the Projects are associated.
        /// If no default `organization_id` is set, one must be set explicitly in this datasource
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        public GetAccountProjectsArgs()
        {
        }
        public static new GetAccountProjectsArgs Empty => new GetAccountProjectsArgs();
    }

    public sealed class GetAccountProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the Organization with which the Projects are associated.
        /// If no default `organization_id` is set, one must be set explicitly in this datasource
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public GetAccountProjectsInvokeArgs()
        {
        }
        public static new GetAccountProjectsInvokeArgs Empty => new GetAccountProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountProjectsResult
    {
        /// <summary>
        /// (Computed) The date and time when the project was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// (Computed) The description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The name of the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Computed) The unique identifier of the organization with which the project is associated.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// (Computed) A list of projects. Each project has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountProjectsProjectResult> Projects;
        /// <summary>
        /// (Computed) The date and time when the project was updated.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetAccountProjectsResult(
            string createdAt,

            string description,

            string id,

            string name,

            string organizationId,

            ImmutableArray<Outputs.GetAccountProjectsProjectResult> projects,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            Projects = projects;
            UpdatedAt = updatedAt;
        }
    }
}
