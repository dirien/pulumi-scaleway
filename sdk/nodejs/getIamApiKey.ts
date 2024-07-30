// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getIamApiKey({
 *     accessKey: "SCWABCDEFGHIJKLMNOPQ",
 * });
 * ```
 */
export function getIamApiKey(args: GetIamApiKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetIamApiKeyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIamApiKey:getIamApiKey", {
        "accessKey": args.accessKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamApiKey.
 */
export interface GetIamApiKeyArgs {
    /**
     * The access key of the IAM API key which is also the ID of the API key.
     */
    accessKey: string;
}

/**
 * A collection of values returned by getIamApiKey.
 */
export interface GetIamApiKeyResult {
    readonly accessKey: string;
    readonly applicationId: string;
    readonly createdAt: string;
    readonly creationIp: string;
    readonly defaultProjectId: string;
    readonly description: string;
    readonly editable: boolean;
    readonly expiresAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly updatedAt: string;
    readonly userId: string;
}
/**
 * Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getIamApiKey({
 *     accessKey: "SCWABCDEFGHIJKLMNOPQ",
 * });
 * ```
 */
export function getIamApiKeyOutput(args: GetIamApiKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIamApiKeyResult> {
    return pulumi.output(args).apply((a: any) => getIamApiKey(a, opts))
}

/**
 * A collection of arguments for invoking getIamApiKey.
 */
export interface GetIamApiKeyOutputArgs {
    /**
     * The access key of the IAM API key which is also the ID of the API key.
     */
    accessKey: pulumi.Input<string>;
}
