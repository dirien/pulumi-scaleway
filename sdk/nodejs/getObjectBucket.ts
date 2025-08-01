// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `scaleway.ObjectBucket` data source is used to retrieve information about an Object Storage bucket.
 *
 * Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/object-storage/how-to/create-a-bucket/) for more information.
 *
 * ## Retrieve an Object Storage bucket
 *
 * The following commands allow you to:
 *
 * - retrieve a bucket by its name
 * - retrieve a bucket by its ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.ObjectBucket("main", {tags: {
 *     foo: "bar",
 * }});
 * const selected = scaleway.getObjectBucketOutput({
 *     name: main.id,
 * });
 * ```
 *
 * ## Retrieve a bucket from a specific project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const selected = scaleway.getObjectBucket({
 *     name: "bucket.test.com",
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getObjectBucket(args?: GetObjectBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectBucketResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getObjectBucket:getObjectBucket", {
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectBucket.
 */
export interface GetObjectBucketArgs {
    name?: string;
    /**
     * `projectId`) The ID of the project with which the bucket is associated.
     */
    projectId?: string;
    /**
     * `region`) The region in which the bucket exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getObjectBucket.
 */
export interface GetObjectBucketResult {
    readonly acl: string;
    readonly apiEndpoint: string;
    readonly corsRules: outputs.GetObjectBucketCorsRule[];
    /**
     * The endpoint URL of the bucket
     */
    readonly endpoint: string;
    readonly forceDestroy: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lifecycleRules: outputs.GetObjectBucketLifecycleRule[];
    readonly name?: string;
    readonly objectLockEnabled: boolean;
    readonly projectId?: string;
    readonly region?: string;
    readonly tags: {[key: string]: string};
    readonly versionings: outputs.GetObjectBucketVersioning[];
}
/**
 * The `scaleway.ObjectBucket` data source is used to retrieve information about an Object Storage bucket.
 *
 * Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/object-storage/how-to/create-a-bucket/) for more information.
 *
 * ## Retrieve an Object Storage bucket
 *
 * The following commands allow you to:
 *
 * - retrieve a bucket by its name
 * - retrieve a bucket by its ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.ObjectBucket("main", {tags: {
 *     foo: "bar",
 * }});
 * const selected = scaleway.getObjectBucketOutput({
 *     name: main.id,
 * });
 * ```
 *
 * ## Retrieve a bucket from a specific project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const selected = scaleway.getObjectBucket({
 *     name: "bucket.test.com",
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getObjectBucketOutput(args?: GetObjectBucketOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetObjectBucketResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getObjectBucket:getObjectBucket", {
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getObjectBucket.
 */
export interface GetObjectBucketOutputArgs {
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project with which the bucket is associated.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the bucket exists.
     */
    region?: pulumi.Input<string>;
}
