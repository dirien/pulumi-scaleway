// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
 */
export function getRdbInstance(args?: GetRdbInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetRdbInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getRdbInstance:getRdbInstance", {
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdbInstance.
 */
export interface GetRdbInstanceArgs {
    /**
     * The RDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: string;
    /**
     * The name of the RDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
     */
    projectId?: string;
    /**
     * `region`) The region in which the RDB instance exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getRdbInstance.
 */
export interface GetRdbInstanceResult {
    readonly backupSameRegion: boolean;
    readonly backupScheduleFrequency: number;
    readonly backupScheduleRetention: number;
    readonly certificate: string;
    readonly disableBackup: boolean;
    readonly endpointIp: string;
    readonly endpointPort: number;
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly initSettings: {[key: string]: string};
    readonly instanceId?: string;
    readonly isHaCluster: boolean;
    readonly loadBalancers: outputs.GetRdbInstanceLoadBalancer[];
    readonly name?: string;
    readonly nodeType: string;
    readonly organizationId: string;
    readonly password: string;
    readonly privateNetworks: outputs.GetRdbInstancePrivateNetwork[];
    readonly projectId: string;
    readonly readReplicas: outputs.GetRdbInstanceReadReplica[];
    readonly region?: string;
    readonly settings: {[key: string]: string};
    readonly tags: string[];
    readonly userName: string;
    readonly volumeSizeInGb: number;
    readonly volumeType: string;
}
/**
 * Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
 */
export function getRdbInstanceOutput(args?: GetRdbInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRdbInstanceResult> {
    return pulumi.output(args).apply((a: any) => getRdbInstance(a, opts))
}

/**
 * A collection of arguments for invoking getRdbInstance.
 */
export interface GetRdbInstanceOutputArgs {
    /**
     * The RDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the RDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the RDB instance exists.
     */
    region?: pulumi.Input<string>;
}
