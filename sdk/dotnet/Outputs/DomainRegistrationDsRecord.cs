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
    public sealed class DomainRegistrationDsRecord
    {
        /// <summary>
        /// The algorithm used for dnssec (e.g., rsasha256, ecdsap256sha256).
        /// </summary>
        public readonly string? Algorithm;
        /// <summary>
        /// Details about the digest.
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainRegistrationDsRecordDigest> Digests;
        /// <summary>
        /// The identifier for the dnssec key.
        /// </summary>
        public readonly int? KeyId;
        /// <summary>
        /// Public key associated with the dnssec record.
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainRegistrationDsRecordPublicKey> PublicKeys;

        [OutputConstructor]
        private DomainRegistrationDsRecord(
            string? algorithm,

            ImmutableArray<Outputs.DomainRegistrationDsRecordDigest> digests,

            int? keyId,

            ImmutableArray<Outputs.DomainRegistrationDsRecordPublicKey> publicKeys)
        {
            Algorithm = algorithm;
            Digests = digests;
            KeyId = keyId;
            PublicKeys = publicKeys;
        }
    }
}
