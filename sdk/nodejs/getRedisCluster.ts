// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a Redis™ cluster.
 *
 * For further information refer to the Managed Database for Redis™ [API documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myCluster = scaleway.getRedisCluster({
 *     clusterId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRedisCluster(args?: GetRedisClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetRedisClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getRedisCluster:getRedisCluster", {
        "clusterId": args.clusterId,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisCluster.
 */
export interface GetRedisClusterArgs {
    /**
     * The Redis cluster ID.
     *
     * > **Note** You must specify at least one: `name` and/or `clusterId`.
     */
    clusterId?: string;
    /**
     * The name of the Redis cluster.
     */
    name?: string;
    /**
     * The ID of the project the Redis cluster is associated with.
     */
    projectId?: string;
    /**
     * `region`) The zone in which the server exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getRedisCluster.
 */
export interface GetRedisClusterResult {
    /**
     * List of acl rules.
     */
    readonly acls: outputs.GetRedisClusterAcl[];
    /**
     * The PEM of the certificate used by redis, only when `tlsEnabled` is true.
     */
    readonly certificate: string;
    readonly clusterId?: string;
    /**
     * The number of nodes in the Redis Cluster.
     */
    readonly clusterSize: number;
    /**
     * The date and time of creation of the Redis Cluster.
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The type of Redis Cluster (e.g. `RED1-M`).
     */
    readonly nodeType: string;
    /**
     * Password of the first user of the Redis Cluster.
     */
    readonly password: string;
    /**
     * List of private networks endpoints of the Redis Cluster.
     */
    readonly privateNetworks: outputs.GetRedisClusterPrivateNetwork[];
    readonly projectId?: string;
    /**
     * Public network details.
     */
    readonly publicNetworks: outputs.GetRedisClusterPublicNetwork[];
    /**
     * Map of settings for redis cluster.
     */
    readonly settings: {[key: string]: string};
    /**
     * The tags associated with the Redis Cluster.
     */
    readonly tags: string[];
    /**
     * Whether TLS is enabled or not.
     */
    readonly tlsEnabled: boolean;
    /**
     * The date and time of the last update of the Redis Cluster.
     */
    readonly updatedAt: string;
    /**
     * The first user of the Redis Cluster.
     */
    readonly userName: string;
    /**
     * Redis's Cluster version (e.g. `6.2.7`).
     */
    readonly version: string;
    readonly zone?: string;
}
/**
 * Gets information about a Redis™ cluster.
 *
 * For further information refer to the Managed Database for Redis™ [API documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myCluster = scaleway.getRedisCluster({
 *     clusterId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getRedisClusterOutput(args?: GetRedisClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRedisClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getRedisCluster:getRedisCluster", {
        "clusterId": args.clusterId,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisCluster.
 */
export interface GetRedisClusterOutputArgs {
    /**
     * The Redis cluster ID.
     *
     * > **Note** You must specify at least one: `name` and/or `clusterId`.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The name of the Redis cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the Redis cluster is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The zone in which the server exists.
     */
    zone?: pulumi.Input<string>;
}
