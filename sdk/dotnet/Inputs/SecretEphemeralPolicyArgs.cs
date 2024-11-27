// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class SecretEphemeralPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to perform when the version of a secret expires. Available values can be found in [SDK constants](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/secret/v1beta1#pkg-constants).
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// True if the secret version expires after a single user access.
        /// </summary>
        [Input("expiresOnceAccessed")]
        public Input<bool>? ExpiresOnceAccessed { get; set; }

        /// <summary>
        /// Time frame, from one second and up to one year, during which the secret's versions are valid. Has to be specified in [Go Duration format](https://pkg.go.dev/time#ParseDuration) (ex: "30m", "24h").
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public SecretEphemeralPolicyArgs()
        {
        }
        public static new SecretEphemeralPolicyArgs Empty => new SecretEphemeralPolicyArgs();
    }
}