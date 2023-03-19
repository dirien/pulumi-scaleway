// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getAccountSshKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getAccountSshKey(args?: GetAccountSshKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountSshKeyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getAccountSshKey:getAccountSshKey", {
        "name": args.name,
        "sshKeyId": args.sshKeyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountSshKey.
 */
export interface GetAccountSshKeyArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: string;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: string;
}

/**
 * A collection of values returned by getAccountSshKey.
 */
export interface GetAccountSshKeyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The ID of the organization the SSH key is associated with.
     */
    readonly organizationId: string;
    readonly projectId: string;
    /**
     * The SSH public key string
     */
    readonly publicKey: string;
    readonly sshKeyId?: string;
}
/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getAccountSshKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getAccountSshKeyOutput(args?: GetAccountSshKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountSshKeyResult> {
    return pulumi.output(args).apply((a: any) => getAccountSshKey(a, opts))
}

/**
 * A collection of arguments for invoking getAccountSshKey.
 */
export interface GetAccountSshKeyOutputArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: pulumi.Input<string>;
}
