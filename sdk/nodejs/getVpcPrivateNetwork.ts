// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a Private Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myName = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 * });
 * const myNameAndVpcId = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 *     vpcId: "11111111-1111-1111-1111-111111111111",
 * });
 * const myId = scaleway.getVpcPrivateNetwork({
 *     privateNetworkId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getVpcPrivateNetwork(args?: GetVpcPrivateNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPrivateNetworkResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", {
        "name": args.name,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "region": args.region,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkArgs {
    /**
     * Name of the Private Network. Cannot be used with `privateNetworkId`.
     */
    name?: string;
    /**
     * ID of the Private Network. Cannot be used with `name` or `vpcId`.
     */
    privateNetworkId?: string;
    /**
     * The ID of the Project the Private Network is associated with.
     */
    projectId?: string;
    region?: string;
    /**
     * ID of the VPC the Private Network is in. Cannot be used with `privateNetworkId`.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkResult {
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IPv4 subnet associated with the Private Network.
     */
    readonly ipv4Subnets: outputs.GetVpcPrivateNetworkIpv4Subnet[];
    /**
     * The IPv6 subnets associated with the Private Network.
     */
    readonly ipv6Subnets: outputs.GetVpcPrivateNetworkIpv6Subnet[];
    readonly isRegional: boolean;
    readonly name?: string;
    readonly organizationId: string;
    readonly privateNetworkId?: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly tags: string[];
    readonly updatedAt: string;
    readonly vpcId?: string;
    readonly zone: string;
}
/**
 * Gets information about a Private Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myName = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 * });
 * const myNameAndVpcId = scaleway.getVpcPrivateNetwork({
 *     name: "foobar",
 *     vpcId: "11111111-1111-1111-1111-111111111111",
 * });
 * const myId = scaleway.getVpcPrivateNetwork({
 *     privateNetworkId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getVpcPrivateNetworkOutput(args?: GetVpcPrivateNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcPrivateNetworkResult> {
    return pulumi.output(args).apply((a: any) => getVpcPrivateNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getVpcPrivateNetwork.
 */
export interface GetVpcPrivateNetworkOutputArgs {
    /**
     * Name of the Private Network. Cannot be used with `privateNetworkId`.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the Private Network. Cannot be used with `name` or `vpcId`.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * The ID of the Project the Private Network is associated with.
     */
    projectId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * ID of the VPC the Private Network is in. Cannot be used with `privateNetworkId`.
     */
    vpcId?: pulumi.Input<string>;
}
