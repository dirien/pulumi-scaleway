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

    public sealed class DomainRecordHttpServiceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ips", required: true)]
        private InputList<string>? _ips;

        /// <summary>
        /// List of IPs to check
        /// </summary>
        public InputList<string> Ips
        {
            get => _ips ?? (_ips = new InputList<string>());
            set => _ips = value;
        }

        /// <summary>
        /// Text to search
        /// </summary>
        [Input("mustContain", required: true)]
        public Input<string> MustContain { get; set; } = null!;

        /// <summary>
        /// Strategy to return an IP from the IPs list. Can be `random`, `hashed` or `all`
        /// </summary>
        [Input("strategy", required: true)]
        public Input<string> Strategy { get; set; } = null!;

        /// <summary>
        /// URL to match the `must_contain` text to validate an IP
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// User-agent used when checking the URL
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public DomainRecordHttpServiceGetArgs()
        {
        }
        public static new DomainRecordHttpServiceGetArgs Empty => new DomainRecordHttpServiceGetArgs();
    }
}
