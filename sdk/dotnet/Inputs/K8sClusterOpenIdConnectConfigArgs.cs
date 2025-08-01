// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class K8sClusterOpenIdConnectConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A client id that all tokens must be issued for
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("groupsClaims")]
        private InputList<string>? _groupsClaims;

        /// <summary>
        /// JWT claim to use as the user's group
        /// </summary>
        public InputList<string> GroupsClaims
        {
            get => _groupsClaims ?? (_groupsClaims = new InputList<string>());
            set => _groupsClaims = value;
        }

        /// <summary>
        /// Prefix prepended to group claims
        /// </summary>
        [Input("groupsPrefix")]
        public Input<string>? GroupsPrefix { get; set; }

        /// <summary>
        /// URL of the provider which allows the API server to discover public signing keys
        /// </summary>
        [Input("issuerUrl", required: true)]
        public Input<string> IssuerUrl { get; set; } = null!;

        [Input("requiredClaims")]
        private InputList<string>? _requiredClaims;

        /// <summary>
        /// Multiple key=value pairs that describes a required claim in the ID Token
        /// </summary>
        public InputList<string> RequiredClaims
        {
            get => _requiredClaims ?? (_requiredClaims = new InputList<string>());
            set => _requiredClaims = value;
        }

        /// <summary>
        /// JWT claim to use as the user name
        /// </summary>
        [Input("usernameClaim")]
        public Input<string>? UsernameClaim { get; set; }

        /// <summary>
        /// Prefix prepended to username
        /// </summary>
        [Input("usernamePrefix")]
        public Input<string>? UsernamePrefix { get; set; }

        public K8sClusterOpenIdConnectConfigArgs()
        {
        }
        public static new K8sClusterOpenIdConnectConfigArgs Empty => new K8sClusterOpenIdConnectConfigArgs();
    }
}
