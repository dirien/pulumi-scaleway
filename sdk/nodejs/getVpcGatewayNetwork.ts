// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a GatewayNetwork (a connection between a Public Gateway and a Private Network).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.VpcGatewayNetwork("main", {
 *     gatewayId: scaleway_vpc_public_gateway.pg01.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn01.id,
 *     dhcpId: scaleway_vpc_public_gateway_dhcp.dhcp01.id,
 *     cleanupDhcp: true,
 *     enableMasquerade: true,
 * });
 * const byId = scaleway.getVpcGatewayNetworkOutput({
 *     gatewayNetworkId: main.id,
 * });
 * const byGatewayAndPn = scaleway.getVpcGatewayNetwork({
 *     gatewayId: scaleway_vpc_public_gateway.pg01.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn01.id,
 * });
 * ```
 */
export function getVpcGatewayNetwork(args?: GetVpcGatewayNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcGatewayNetworkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", {
        "dhcpId": args.dhcpId,
        "enableMasquerade": args.enableMasquerade,
        "gatewayId": args.gatewayId,
        "gatewayNetworkId": args.gatewayNetworkId,
        "privateNetworkId": args.privateNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcGatewayNetwork.
 */
export interface GetVpcGatewayNetworkArgs {
    dhcpId?: string;
    enableMasquerade?: boolean;
    gatewayId?: string;
    /**
     * ID of the GatewayNetwork.
     */
    gatewayNetworkId?: string;
    privateNetworkId?: string;
}

/**
 * A collection of values returned by getVpcGatewayNetwork.
 */
export interface GetVpcGatewayNetworkResult {
    readonly cleanupDhcp: boolean;
    readonly createdAt: string;
    readonly dhcpId?: string;
    readonly enableDhcp: boolean;
    readonly enableMasquerade?: boolean;
    readonly gatewayId?: string;
    readonly gatewayNetworkId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipamConfigs: outputs.GetVpcGatewayNetworkIpamConfig[];
    readonly macAddress: string;
    readonly privateNetworkId?: string;
    readonly staticAddress: string;
    readonly status: string;
    readonly updatedAt: string;
    readonly zone: string;
}
/**
 * Gets information about a GatewayNetwork (a connection between a Public Gateway and a Private Network).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.VpcGatewayNetwork("main", {
 *     gatewayId: scaleway_vpc_public_gateway.pg01.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn01.id,
 *     dhcpId: scaleway_vpc_public_gateway_dhcp.dhcp01.id,
 *     cleanupDhcp: true,
 *     enableMasquerade: true,
 * });
 * const byId = scaleway.getVpcGatewayNetworkOutput({
 *     gatewayNetworkId: main.id,
 * });
 * const byGatewayAndPn = scaleway.getVpcGatewayNetwork({
 *     gatewayId: scaleway_vpc_public_gateway.pg01.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn01.id,
 * });
 * ```
 */
export function getVpcGatewayNetworkOutput(args?: GetVpcGatewayNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVpcGatewayNetworkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", {
        "dhcpId": args.dhcpId,
        "enableMasquerade": args.enableMasquerade,
        "gatewayId": args.gatewayId,
        "gatewayNetworkId": args.gatewayNetworkId,
        "privateNetworkId": args.privateNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcGatewayNetwork.
 */
export interface GetVpcGatewayNetworkOutputArgs {
    dhcpId?: pulumi.Input<string>;
    enableMasquerade?: pulumi.Input<boolean>;
    gatewayId?: pulumi.Input<string>;
    /**
     * ID of the GatewayNetwork.
     */
    gatewayNetworkId?: pulumi.Input<string>;
    privateNetworkId?: pulumi.Input<string>;
}
