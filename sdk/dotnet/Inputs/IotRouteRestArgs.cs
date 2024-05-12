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

    public sealed class IotRouteRestArgs : global::Pulumi.ResourceArgs
    {
        [Input("headers", required: true)]
        private InputMap<string>? _headers;

        /// <summary>
        /// a map of the extra headers to send with the HTTP call (e.g. `X-Header = Value`).
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// The URI of the Rest endpoint (e.g. `https://internal.mycompany.com/ingest/mqttdata`).
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        /// <summary>
        /// The HTTP Verb used to call Rest URI (e.g. `post`).
        /// </summary>
        [Input("verb", required: true)]
        public Input<string> Verb { get; set; } = null!;

        public IotRouteRestArgs()
        {
        }
        public static new IotRouteRestArgs Empty => new IotRouteRestArgs();
    }
}
