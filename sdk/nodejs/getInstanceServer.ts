// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about an instance server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getInstanceServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceServer(args?: GetInstanceServerArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceServerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getInstanceServer:getInstanceServer", {
        "name": args.name,
        "projectId": args.projectId,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceServer.
 */
export interface GetInstanceServerArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the instance server is associated with.
     */
    projectId?: string;
    /**
     * The server id. Only one of `name` and `serverId` should be specified.
     */
    serverId?: string;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstanceServer.
 */
export interface GetInstanceServerResult {
    /**
     * The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
     * attached to the server.
     */
    readonly additionalVolumeIds: string[];
    readonly bootType: string;
    readonly bootscriptId: string;
    /**
     * The cloud init script associated with this server.
     */
    readonly cloudInit: string;
    /**
     * True if dynamic IP in enable on the server.
     */
    readonly enableDynamicIp: boolean;
    /**
     * Determines if IPv6 is enabled for the server.
     */
    readonly enableIpv6: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The UUID and the label of the base image used by the server.
     */
    readonly image: string;
    readonly ipId: string;
    readonly ipIds: string[];
    /**
     * The default ipv6 address routed to the server. ( Only set when enableIpv6 is set to true )
     */
    readonly ipv6Address: string;
    /**
     * The ipv6 gateway address. ( Only set when enableIpv6 is set to true )
     */
    readonly ipv6Gateway: string;
    /**
     * The prefix length of the ipv6 subnet routed to the server. ( Only set when enableIpv6 is set to true )
     */
    readonly ipv6PrefixLength: number;
    readonly name?: string;
    /**
     * The ID of the organization the server is associated with.
     */
    readonly organizationId: string;
    /**
     * The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
     */
    readonly placementGroupId: string;
    /**
     * True when the placement group policy is respected.
     */
    readonly placementGroupPolicyRespected: boolean;
    /**
     * The Scaleway internal IP address of the server.
     */
    readonly privateIp: string;
    readonly privateNetworks: outputs.GetInstanceServerPrivateNetwork[];
    readonly projectId?: string;
    /**
     * The public IP address of the server.
     */
    readonly publicIp: string;
    /**
     * The list of public IPs of the server
     */
    readonly publicIps: outputs.GetInstanceServerPublicIp[];
    readonly replaceOnTypeChange: boolean;
    readonly rootVolumes: outputs.GetInstanceServerRootVolume[];
    /**
     * True if the server support routed ip only.
     */
    readonly routedIpEnabled: boolean;
    /**
     * The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
     */
    readonly securityGroupId: string;
    readonly serverId?: string;
    /**
     * The state of the server. Possible values are: `started`, `stopped` or `standby`.
     */
    readonly state: string;
    /**
     * The tags associated with the server.
     */
    readonly tags: string[];
    /**
     * The commercial type of the server.
     * You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
     */
    readonly type: string;
    /**
     * The user data associated with the server.
     */
    readonly userData: {[key: string]: string};
    readonly zone?: string;
}
/**
 * Gets information about an instance server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.getInstanceServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceServerOutput(args?: GetInstanceServerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstanceServerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getInstanceServer:getInstanceServer", {
        "name": args.name,
        "projectId": args.projectId,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceServer.
 */
export interface GetInstanceServerOutputArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the instance server is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The server id. Only one of `name` and `serverId` should be specified.
     */
    serverId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: pulumi.Input<string>;
}
