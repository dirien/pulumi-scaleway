// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
 *
 * ## Examples
 *
 * ### IPAM IP ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byId = scaleway.getIpamIp({
 *     ipamIpId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 *
 * ### Instance Private Network IP
 *
 * Get Instance IP in a private network.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Connect your instance to a private network using a private nic.
 * const nic = new scaleway.InstancePrivateNic("nic", {
 *     serverId: scaleway_instance_server.server.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn.id,
 * });
 * const byMac = scaleway.getIpamIpOutput({
 *     macAddress: nic.macAddress,
 *     type: "ipv4",
 * });
 * const byId = scaleway.getIpamIpOutput({
 *     resource: {
 *         id: nic.id,
 *         type: "instance_private_nic",
 *     },
 *     type: "ipv4",
 * });
 * ```
 *
 * ### RDB instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Find the private IPv4 using resource name
 * const pn = new scaleway.VpcPrivateNetwork("pn", {});
 * const main = new scaleway.RdbInstance("main", {
 *     nodeType: "DB-DEV-S",
 *     engine: "PostgreSQL-15",
 *     isHaCluster: true,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     privateNetwork: {
 *         pnId: pn.id,
 *     },
 * });
 * const byName = scaleway.getIpamIpOutput({
 *     resource: {
 *         name: main.name,
 *         type: "rdb_instance",
 *     },
 *     type: "ipv4",
 * });
 * ```
 */
export function getIpamIp(args?: GetIpamIpArgs, opts?: pulumi.InvokeOptions): Promise<GetIpamIpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIpamIp:getIpamIp", {
        "attached": args.attached,
        "ipamIpId": args.ipamIpId,
        "macAddress": args.macAddress,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "region": args.region,
        "resource": args.resource,
        "tags": args.tags,
        "type": args.type,
        "zonal": args.zonal,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpamIp.
 */
export interface GetIpamIpArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipamIpId`.
     */
    attached?: boolean;
    /**
     * The IPAM IP ID. Cannot be used with the rest of the arguments.
     */
    ipamIpId?: string;
    /**
     * The Mac Address linked to the IP. Cannot be used with `ipamIpId`.
     */
    macAddress?: string;
    /**
     * The ID of the private network the IP belong to. Cannot be used with `ipamIpId`.
     */
    privateNetworkId?: string;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the IP exists.
     */
    region?: string;
    /**
     * Filter by resource ID, type or name. Cannot be used with `ipamIpId`.
     * If specified, `type` is required, and at least one of `id` or `name` must be set.
     */
    resource?: inputs.GetIpamIpResource;
    /**
     * The tags associated with the IP. Cannot be used with `ipamIpId`.
     * As datasource only returns one IP, the search with given tags must return only one result.
     */
    tags?: string[];
    /**
     * The type of IP to search for (ipv4, ipv6). Cannot be used with `ipamIpId`.
     */
    type?: string;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: string;
}

/**
 * A collection of values returned by getIpamIp.
 */
export interface GetIpamIpResult {
    /**
     * The IP address.
     */
    readonly address: string;
    /**
     * the IP address with a CIDR notation.
     */
    readonly addressCidr: string;
    readonly attached?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipamIpId?: string;
    readonly macAddress?: string;
    readonly organizationId: string;
    readonly privateNetworkId?: string;
    readonly projectId: string;
    readonly region: string;
    readonly resource?: outputs.GetIpamIpResource;
    readonly tags?: string[];
    readonly type?: string;
    readonly zonal: string;
}
/**
 * Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
 *
 * ## Examples
 *
 * ### IPAM IP ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byId = scaleway.getIpamIp({
 *     ipamIpId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 *
 * ### Instance Private Network IP
 *
 * Get Instance IP in a private network.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Connect your instance to a private network using a private nic.
 * const nic = new scaleway.InstancePrivateNic("nic", {
 *     serverId: scaleway_instance_server.server.id,
 *     privateNetworkId: scaleway_vpc_private_network.pn.id,
 * });
 * const byMac = scaleway.getIpamIpOutput({
 *     macAddress: nic.macAddress,
 *     type: "ipv4",
 * });
 * const byId = scaleway.getIpamIpOutput({
 *     resource: {
 *         id: nic.id,
 *         type: "instance_private_nic",
 *     },
 *     type: "ipv4",
 * });
 * ```
 *
 * ### RDB instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Find the private IPv4 using resource name
 * const pn = new scaleway.VpcPrivateNetwork("pn", {});
 * const main = new scaleway.RdbInstance("main", {
 *     nodeType: "DB-DEV-S",
 *     engine: "PostgreSQL-15",
 *     isHaCluster: true,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     privateNetwork: {
 *         pnId: pn.id,
 *     },
 * });
 * const byName = scaleway.getIpamIpOutput({
 *     resource: {
 *         name: main.name,
 *         type: "rdb_instance",
 *     },
 *     type: "ipv4",
 * });
 * ```
 */
export function getIpamIpOutput(args?: GetIpamIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpamIpResult> {
    return pulumi.output(args).apply((a: any) => getIpamIp(a, opts))
}

/**
 * A collection of arguments for invoking getIpamIp.
 */
export interface GetIpamIpOutputArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipamIpId`.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * The IPAM IP ID. Cannot be used with the rest of the arguments.
     */
    ipamIpId?: pulumi.Input<string>;
    /**
     * The Mac Address linked to the IP. Cannot be used with `ipamIpId`.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The ID of the private network the IP belong to. Cannot be used with `ipamIpId`.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the IP exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Filter by resource ID, type or name. Cannot be used with `ipamIpId`.
     * If specified, `type` is required, and at least one of `id` or `name` must be set.
     */
    resource?: pulumi.Input<inputs.GetIpamIpResourceArgs>;
    /**
     * The tags associated with the IP. Cannot be used with `ipamIpId`.
     * As datasource only returns one IP, the search with given tags must return only one result.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of IP to search for (ipv4, ipv6). Cannot be used with `ipamIpId`.
     */
    type?: pulumi.Input<string>;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: pulumi.Input<string>;
}
