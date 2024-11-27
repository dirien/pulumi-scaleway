// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an instance image.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.getInstanceImage({
 *     imageId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceImage(args?: GetInstanceImageArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getInstanceImage:getInstanceImage", {
        "architecture": args.architecture,
        "imageId": args.imageId,
        "latest": args.latest,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceImage.
 */
export interface GetInstanceImageArgs {
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: string;
    /**
     * The image id. Only one of `name` and `imageId` should be specified.
     */
    imageId?: string;
    /**
     * Use the latest image ID.
     */
    latest?: boolean;
    /**
     * The image name. Only one of `name` and `imageId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: string;
    /**
     * `zone`) The zone in which the image exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstanceImage.
 */
export interface GetInstanceImageResult {
    /**
     * IDs of the additional volumes in this image.
     */
    readonly additionalVolumeIds: string[];
    readonly architecture?: string;
    /**
     * Date of the image creation.
     */
    readonly creationDate: string;
    /**
     * ID of the default bootscript for this image.
     */
    readonly defaultBootscriptId: string;
    /**
     * ID of the server the image if based from.
     */
    readonly fromServerId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId?: string;
    readonly latest?: boolean;
    /**
     * Date of image latest update.
     */
    readonly modificationDate: string;
    readonly name?: string;
    /**
     * The ID of the organization the image is associated with.
     */
    readonly organizationId: string;
    /**
     * The ID of the project the image is associated with.
     */
    readonly projectId: string;
    /**
     * Set to `true` if the image is public.
     */
    readonly public: boolean;
    /**
     * ID of the root volume in this image.
     */
    readonly rootVolumeId: string;
    /**
     * State of the image. Possible values are: `available`, `creating` or `error`.
     */
    readonly state: string;
    readonly zone: string;
}
/**
 * Gets information about an instance image.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.getInstanceImage({
 *     imageId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceImageOutput(args?: GetInstanceImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getInstanceImage:getInstanceImage", {
        "architecture": args.architecture,
        "imageId": args.imageId,
        "latest": args.latest,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceImage.
 */
export interface GetInstanceImageOutputArgs {
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * The image id. Only one of `name` and `imageId` should be specified.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Use the latest image ID.
     */
    latest?: pulumi.Input<boolean>;
    /**
     * The image name. Only one of `name` and `imageId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the image exists.
     */
    zone?: pulumi.Input<string>;
}
