// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a Container Registry image.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.getRegistryImage({
 *     imageId: "11111111-1111-1111-1111-111111111111",
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRegistryImage(args?: GetRegistryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getRegistryImage:getRegistryImage", {
        "imageId": args.imageId,
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageArgs {
    /**
     * The image ID.
     *
     * > **Note** You must specify at least one: `name` and/or `imageId`.
     */
    imageId?: string;
    /**
     * The image name.
     */
    name?: string;
    /**
     * The namespace ID in which the image is.
     */
    namespaceId?: string;
    /**
     * `projectId`) The ID of the project the image is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the image exists.
     */
    region?: string;
    /**
     * The tags associated with the registry image
     */
    tags?: string[];
}

/**
 * A collection of values returned by getRegistryImage.
 */
export interface GetRegistryImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId?: string;
    readonly name?: string;
    readonly namespaceId: string;
    /**
     * The organization ID the image is associated with.
     */
    readonly organizationId: string;
    readonly projectId: string;
    readonly region: string;
    /**
     * The size of the registry image.
     */
    readonly size: number;
    /**
     * The tags associated with the registry image
     */
    readonly tags: string[];
    /**
     * The date the image of the last update
     */
    readonly updatedAt: string;
    /**
     * The privacy policy of the registry image.
     */
    readonly visibility: string;
}
/**
 * Gets information about a Container Registry image.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.getRegistryImage({
 *     imageId: "11111111-1111-1111-1111-111111111111",
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRegistryImageOutput(args?: GetRegistryImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRegistryImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getRegistryImage:getRegistryImage", {
        "imageId": args.imageId,
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageOutputArgs {
    /**
     * The image ID.
     *
     * > **Note** You must specify at least one: `name` and/or `imageId`.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The image name.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace ID in which the image is.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the image exists.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the registry image
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
