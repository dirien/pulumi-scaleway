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

    public sealed class GetIpamIpResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the resource that the IP is attached to.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the resource the IP is attached to.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The type of the resource the IP is attached to. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetIpamIpResourceArgs()
        {
        }
        public static new GetIpamIpResourceArgs Empty => new GetIpamIpResourceArgs();
    }
}
