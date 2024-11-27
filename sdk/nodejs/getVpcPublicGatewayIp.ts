// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a Public Gateway public flexible IP address.
 *
 * For further information, please see the API [documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.VpcPublicGatewayIp("main", {});
 * const ipById = scaleway.getVpcPublicGatewayIpOutput({
 *     ipId: main.id,
 * });
 * ```
 */
export function getVpcPublicGatewayIp(args?: GetVpcPublicGatewayIpArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPublicGatewayIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcPublicGatewayIp:getVpcPublicGatewayIp", {
        "ipId": args.ipId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPublicGatewayIp.
 */
export interface GetVpcPublicGatewayIpArgs {
    ipId?: string;
}

/**
 * A collection of values returned by getVpcPublicGatewayIp.
 */
export interface GetVpcPublicGatewayIpResult {
    readonly address: string;
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipId?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly reverse: string;
    readonly tags: string[];
    readonly updatedAt: string;
    readonly zone: string;
}
/**
 * Gets information about a Public Gateway public flexible IP address.
 *
 * For further information, please see the API [documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.VpcPublicGatewayIp("main", {});
 * const ipById = scaleway.getVpcPublicGatewayIpOutput({
 *     ipId: main.id,
 * });
 * ```
 */
export function getVpcPublicGatewayIpOutput(args?: GetVpcPublicGatewayIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcPublicGatewayIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getVpcPublicGatewayIp:getVpcPublicGatewayIp", {
        "ipId": args.ipId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPublicGatewayIp.
 */
export interface GetVpcPublicGatewayIpOutputArgs {
    ipId?: pulumi.Input<string>;
}
