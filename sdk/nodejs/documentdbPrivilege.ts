// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Create and manage Scaleway DocumentDB database privilege.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.DocumentdbPrivilege("main", {
 *     databaseName: "my-db-name",
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     permission: "all",
 *     userName: "my-db-user",
 * });
 * ```
 *
 * ## Import
 *
 * The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/documentdbPrivilege:DocumentdbPrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
 * ```
 */
export class DocumentdbPrivilege extends pulumi.CustomResource {
    /**
     * Get an existing DocumentdbPrivilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentdbPrivilegeState, opts?: pulumi.CustomResourceOptions): DocumentdbPrivilege {
        return new DocumentdbPrivilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentdbPrivilege:DocumentdbPrivilege';

    /**
     * Returns true if the given object is an instance of DocumentdbPrivilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentdbPrivilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentdbPrivilege.__pulumiType;
    }

    /**
     * Name of the database (e.g. `my-db-name`).
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * UUID of the rdb instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a DocumentdbPrivilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentdbPrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentdbPrivilegeArgs | DocumentdbPrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentdbPrivilegeState | undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as DocumentdbPrivilegeArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DocumentdbPrivilege.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentdbPrivilege resources.
 */
export interface DocumentdbPrivilegeState {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName?: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentdbPrivilege resource.
 */
export interface DocumentdbPrivilegeArgs {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName: pulumi.Input<string>;
}
