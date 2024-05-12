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
    public static class GetTemDomain
    {
        /// <summary>
        /// Gets information about a transactional email domain.
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
        ///     var myDomain = Scaleway.GetTemDomain.Invoke(new()
        ///     {
        ///         DomainId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTemDomainResult> InvokeAsync(GetTemDomainArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemDomainResult>("scaleway:index/getTemDomain:getTemDomain", args ?? new GetTemDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a transactional email domain.
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
        ///     var myDomain = Scaleway.GetTemDomain.Invoke(new()
        ///     {
        ///         DomainId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTemDomainResult> Invoke(GetTemDomainInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemDomainResult>("scaleway:index/getTemDomain:getTemDomain", args ?? new GetTemDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain id.
        /// Only one of `name` and `domain_id` should be specified.
        /// </summary>
        [Input("domainId")]
        public string? DomainId { get; set; }

        /// <summary>
        /// The domain name.
        /// Only one of `name` and `domain_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the domain exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetTemDomainArgs()
        {
        }
        public static new GetTemDomainArgs Empty => new GetTemDomainArgs();
    }

    public sealed class GetTemDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain id.
        /// Only one of `name` and `domain_id` should be specified.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The domain name.
        /// Only one of `name` and `domain_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the domain exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetTemDomainInvokeArgs()
        {
        }
        public static new GetTemDomainInvokeArgs Empty => new GetTemDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemDomainResult
    {
        public readonly bool AcceptTos;
        public readonly string CreatedAt;
        public readonly string DkimConfig;
        public readonly string? DomainId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LastError;
        public readonly string LastValidAt;
        public readonly string MxBlackhole;
        public readonly string? Name;
        public readonly string NextCheckAt;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly ImmutableArray<Outputs.GetTemDomainReputationResult> Reputations;
        public readonly string RevokedAt;
        public readonly string SmtpHost;
        public readonly int SmtpPort;
        public readonly int SmtpPortAlternative;
        public readonly int SmtpPortUnsecure;
        public readonly string SmtpsAuthUser;
        public readonly int SmtpsPort;
        public readonly int SmtpsPortAlternative;
        public readonly string SpfConfig;
        public readonly string Status;

        [OutputConstructor]
        private GetTemDomainResult(
            bool acceptTos,

            string createdAt,

            string dkimConfig,

            string? domainId,

            string id,

            string lastError,

            string lastValidAt,

            string mxBlackhole,

            string? name,

            string nextCheckAt,

            string? projectId,

            string? region,

            ImmutableArray<Outputs.GetTemDomainReputationResult> reputations,

            string revokedAt,

            string smtpHost,

            int smtpPort,

            int smtpPortAlternative,

            int smtpPortUnsecure,

            string smtpsAuthUser,

            int smtpsPort,

            int smtpsPortAlternative,

            string spfConfig,

            string status)
        {
            AcceptTos = acceptTos;
            CreatedAt = createdAt;
            DkimConfig = dkimConfig;
            DomainId = domainId;
            Id = id;
            LastError = lastError;
            LastValidAt = lastValidAt;
            MxBlackhole = mxBlackhole;
            Name = name;
            NextCheckAt = nextCheckAt;
            ProjectId = projectId;
            Region = region;
            Reputations = reputations;
            RevokedAt = revokedAt;
            SmtpHost = smtpHost;
            SmtpPort = smtpPort;
            SmtpPortAlternative = smtpPortAlternative;
            SmtpPortUnsecure = smtpPortUnsecure;
            SmtpsAuthUser = smtpsAuthUser;
            SmtpsPort = smtpsPort;
            SmtpsPortAlternative = smtpsPortAlternative;
            SpfConfig = spfConfig;
            Status = status;
        }
    }
}
