// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an existing IAM group. For more information, please
 * check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const findByName = scaleway.getIamGroup({
 *     name: "foobar",
 * });
 * const findById = scaleway.getIamGroup({
 *     groupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getIamGroup(args?: GetIamGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetIamGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIamGroup:getIamGroup", {
        "groupId": args.groupId,
        "name": args.name,
        "organizationId": args.organizationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamGroup.
 */
export interface GetIamGroupArgs {
    /**
     * The ID of the IAM group.
     * Only one of the `name` and `groupId` should be specified.
     */
    groupId?: string;
    /**
     * The name of the IAM group.
     * Only one of the `name` and `groupId` should be specified.
     */
    name?: string;
    /**
     * `organizationId`) The ID of the
     * organization the group is associated with.
     */
    organizationId?: string;
}

/**
 * A collection of values returned by getIamGroup.
 */
export interface GetIamGroupResult {
    readonly applicationIds: string[];
    readonly createdAt: string;
    readonly description: string;
    readonly externalMembership: boolean;
    readonly groupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly organizationId?: string;
    readonly updatedAt: string;
    readonly userIds: string[];
}
/**
 * Gets information about an existing IAM group. For more information, please
 * check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const findByName = scaleway.getIamGroup({
 *     name: "foobar",
 * });
 * const findById = scaleway.getIamGroup({
 *     groupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getIamGroupOutput(args?: GetIamGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIamGroupResult> {
    return pulumi.output(args).apply((a: any) => getIamGroup(a, opts))
}

/**
 * A collection of arguments for invoking getIamGroup.
 */
export interface GetIamGroupOutputArgs {
    /**
     * The ID of the IAM group.
     * Only one of the `name` and `groupId` should be specified.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of the IAM group.
     * Only one of the `name` and `groupId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`) The ID of the
     * organization the group is associated with.
     */
    organizationId?: pulumi.Input<string>;
}
