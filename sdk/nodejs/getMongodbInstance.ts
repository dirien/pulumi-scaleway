// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a MongoDB® Instance.
 *
 * For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
 */
export function getMongodbInstance(args?: GetMongodbInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetMongodbInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getMongodbInstance:getMongodbInstance", {
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getMongodbInstance.
 */
export interface GetMongodbInstanceArgs {
    /**
     * The MongoDB® instance ID.
     *
     * > **Note** You must specify at least one: `name` or `instanceId`.
     */
    instanceId?: string;
    /**
     * The name of the MongoDB® instance.
     */
    name?: string;
    /**
     * The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
     */
    projectId?: string;
    /**
     * `region`) The region in which the MongoDB® Instance exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getMongodbInstance.
 */
export interface GetMongodbInstanceResult {
    /**
     * The date and time the MongoDB® instance was created.
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    /**
     * The name of the MongoDB® instance.
     */
    readonly name?: string;
    /**
     * The number of nodes in the MongoDB® cluster.
     */
    readonly nodeNumber: number;
    /**
     * The type of MongoDB® node.
     */
    readonly nodeType: string;
    readonly password: string;
    /**
     * The ID of the project the instance belongs to.
     */
    readonly projectId?: string;
    /**
     * The details of the public network configuration, if applicable.
     */
    readonly publicNetworks: outputs.GetMongodbInstancePublicNetwork[];
    readonly region?: string;
    readonly settings: {[key: string]: string};
    readonly snapshotId: string;
    /**
     * A list of tags attached to the MongoDB® instance.
     */
    readonly tags: string[];
    readonly updatedAt: string;
    readonly userName: string;
    /**
     * The version of MongoDB® running on the instance.
     */
    readonly version: string;
    /**
     * The size of the attached volume, in GB.
     */
    readonly volumeSizeInGb: number;
    /**
     * The type of volume attached to the MongoDB® instance.
     */
    readonly volumeType: string;
}
/**
 * Gets information about a MongoDB® Instance.
 *
 * For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
 */
export function getMongodbInstanceOutput(args?: GetMongodbInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMongodbInstanceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getMongodbInstance:getMongodbInstance", {
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getMongodbInstance.
 */
export interface GetMongodbInstanceOutputArgs {
    /**
     * The MongoDB® instance ID.
     *
     * > **Note** You must specify at least one: `name` or `instanceId`.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the MongoDB® instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the MongoDB® Instance exists.
     */
    region?: pulumi.Input<string>;
}