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

    public sealed class IotRouteS3GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the S3 route's destination bucket (e.g. `my-object-storage`).
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The region of the S3 route's destination bucket (e.g. `fr-par`).
        /// </summary>
        [Input("bucketRegion", required: true)]
        public Input<string> BucketRegion { get; set; } = null!;

        /// <summary>
        /// The string to prefix object names with (e.g. `mykeyprefix-`).
        /// </summary>
        [Input("objectPrefix")]
        public Input<string>? ObjectPrefix { get; set; }

        /// <summary>
        /// How the S3 route's objects will be created (e.g. `per_topic`). See [documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Messages-Store-Strategies) for behaviour details.
        /// </summary>
        [Input("strategy", required: true)]
        public Input<string> Strategy { get; set; } = null!;

        public IotRouteS3GetArgs()
        {
        }
        public static new IotRouteS3GetArgs Empty => new IotRouteS3GetArgs();
    }
}
