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

    public sealed class LbCertificateCustomCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateChain", required: true)]
        private Input<string>? _certificateChain;

        /// <summary>
        /// The full PEM-formatted certificate chain
        /// </summary>
        public Input<string>? CertificateChain
        {
            get => _certificateChain;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _certificateChain = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public LbCertificateCustomCertificateArgs()
        {
        }
        public static new LbCertificateCustomCertificateArgs Empty => new LbCertificateCustomCertificateArgs();
    }
}
