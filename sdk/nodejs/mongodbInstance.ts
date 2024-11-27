// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway MongoDB® instance.
 * For more information refer to [the API documentation](https://www.scaleway.com/en/docs/managed-databases/mongodb/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.MongodbInstance("main", {
 *     nodeNumber: 1,
 *     nodeType: "MGDB-PLAY2-NANO",
 *     password: "thiZ_is_v&ry_s3cret",
 *     userName: "my_initial_user",
 *     version: "7.0.12",
 *     volumeSizeInGb: 5,
 * });
 * ```
 *
 * ### Restore From Snapshot
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const restoredInstance = new scaleway.MongodbInstance("restoredInstance", {
 *     nodeNumber: 1,
 *     nodeType: "MGDB-PLAY2-NANO",
 *     snapshotId: scaleway_vpc_private_network.pn.idscaleway_mongodb_snapshot.main_snapshot.id,
 * });
 * ```
 *
 * ## Import
 *
 * MongoDB® instance can be imported using the `id`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/mongodbInstance:MongodbInstance main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class MongodbInstance extends pulumi.CustomResource {
    /**
     * Get an existing MongodbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MongodbInstanceState, opts?: pulumi.CustomResourceOptions): MongodbInstance {
        return new MongodbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mongodbInstance:MongodbInstance';

    /**
     * Returns true if the given object is an instance of MongodbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MongodbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MongodbInstance.__pulumiType;
    }

    /**
     * The date and time of the creation of the MongoDB® instance.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the MongoDB® instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of nodes in the instance
     */
    public readonly nodeNumber!: pulumi.Output<number>;
    /**
     * The type of MongoDB® intance to create.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Password of the user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The projectId you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Public network specs details.
     */
    public readonly publicNetwork!: pulumi.Output<outputs.MongodbInstancePublicNetwork>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Map of settings to define for the instance.
     */
    public readonly settings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Snapshot ID to restore the MongoDB® instance from.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * List of tags attached to the MongoDB® instance.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time of the last update of the MongoDB® instance.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Name of the user created when the intance is created.
     */
    public readonly userName!: pulumi.Output<string | undefined>;
    /**
     * MongoDB® version of the instance.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Volume size in GB.
     */
    public readonly volumeSizeInGb!: pulumi.Output<number>;
    /**
     * Volume type of the instance.
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a MongodbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MongodbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MongodbInstanceArgs | MongodbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MongodbInstanceState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeNumber"] = state ? state.nodeNumber : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publicNetwork"] = state ? state.publicNetwork : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["volumeSizeInGb"] = state ? state.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as MongodbInstanceArgs | undefined;
            if ((!args || args.nodeNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeNumber'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeNumber"] = args ? args.nodeNumber : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publicNetwork"] = args ? args.publicNetwork : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["volumeSizeInGb"] = args ? args.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MongodbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MongodbInstance resources.
 */
export interface MongodbInstanceState {
    /**
     * The date and time of the creation of the MongoDB® instance.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the MongoDB® instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes in the instance
     */
    nodeNumber?: pulumi.Input<number>;
    /**
     * The type of MongoDB® intance to create.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * The projectId you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Public network specs details.
     */
    publicNetwork?: pulumi.Input<inputs.MongodbInstancePublicNetwork>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * Map of settings to define for the instance.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Snapshot ID to restore the MongoDB® instance from.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * List of tags attached to the MongoDB® instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time of the last update of the MongoDB® instance.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Name of the user created when the intance is created.
     */
    userName?: pulumi.Input<string>;
    /**
     * MongoDB® version of the instance.
     */
    version?: pulumi.Input<string>;
    /**
     * Volume size in GB.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Volume type of the instance.
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MongodbInstance resource.
 */
export interface MongodbInstanceArgs {
    /**
     * Name of the MongoDB® instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes in the instance
     */
    nodeNumber: pulumi.Input<number>;
    /**
     * The type of MongoDB® intance to create.
     */
    nodeType: pulumi.Input<string>;
    /**
     * Password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * The projectId you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Public network specs details.
     */
    publicNetwork?: pulumi.Input<inputs.MongodbInstancePublicNetwork>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * Map of settings to define for the instance.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Snapshot ID to restore the MongoDB® instance from.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * List of tags attached to the MongoDB® instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the user created when the intance is created.
     */
    userName?: pulumi.Input<string>;
    /**
     * MongoDB® version of the instance.
     */
    version?: pulumi.Input<string>;
    /**
     * Volume size in GB.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Volume type of the instance.
     */
    volumeType?: pulumi.Input<string>;
}
