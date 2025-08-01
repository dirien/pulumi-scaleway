// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetLbRoutesRouteResult
    {
        /// <summary>
        /// The backend ID to redirect to
        /// </summary>
        public readonly string BackendId;
        /// <summary>
        /// The date on which the route was created (RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
        /// </summary>
        public readonly string FrontendId;
        /// <summary>
        /// The associated route ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the host of the server to which the request is being sent.
        /// </summary>
        public readonly string MatchHostHeader;
        /// <summary>
        /// Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer.
        /// </summary>
        public readonly string MatchSni;
        /// <summary>
        /// If true, all subdomains will match.
        /// </summary>
        public readonly bool MatchSubdomains;
        /// <summary>
        /// The date on which the route was last updated (RFC 3339 format).
        /// </summary>
        public readonly string UpdateAt;

        [OutputConstructor]
        private GetLbRoutesRouteResult(
            string backendId,

            string createdAt,

            string frontendId,

            string id,

            string matchHostHeader,

            string matchSni,

            bool matchSubdomains,

            string updateAt)
        {
            BackendId = backendId;
            CreatedAt = createdAt;
            FrontendId = frontendId;
            Id = id;
            MatchHostHeader = matchHostHeader;
            MatchSni = matchSni;
            MatchSubdomains = matchSubdomains;
            UpdateAt = updateAt;
        }
    }
}
