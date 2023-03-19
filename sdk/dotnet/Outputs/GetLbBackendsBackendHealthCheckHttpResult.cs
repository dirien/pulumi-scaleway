// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class GetLbBackendsBackendHealthCheckHttpResult
    {
        /// <summary>
        /// The expected HTTP status code.
        /// </summary>
        public readonly int Code;
        /// <summary>
        /// The HTTP host header to use for HC requests.
        /// </summary>
        public readonly string HostHeader;
        /// <summary>
        /// The HTTP method to use for HC requests.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// The SNI to use for HC requests over SSL.
        /// </summary>
        public readonly string Sni;
        /// <summary>
        /// The HTTPS endpoint URL to call for HC requests.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetLbBackendsBackendHealthCheckHttpResult(
            int code,

            string hostHeader,

            string method,

            string sni,

            string uri)
        {
            Code = code;
            HostHeader = hostHeader;
            Method = method;
            Sni = sni;
            Uri = uri;
        }
    }
}
