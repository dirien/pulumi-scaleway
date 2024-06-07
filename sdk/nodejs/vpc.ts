// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Virtual Private Clouds.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const vpc01 = new scaleway.Vpc("vpc01", {tags: [
 *     "demo",
 *     "terraform",
 * ]});
 * ```
 *
 * ### Enable routing
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const vpc01 = new scaleway.Vpc("vpc01", {
 *     enableRouting: true,
 *     tags: [
 *         "demo",
 *         "terraform",
 *         "routing",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * VPCs can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/vpc:Vpc vpc_demo fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Vpc extends pulumi.CustomResource {
    /**
     * Get an existing Vpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcState, opts?: pulumi.CustomResourceOptions): Vpc {
        return new Vpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/vpc:Vpc';

    /**
     * Returns true if the given object is an instance of Vpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vpc.__pulumiType;
    }

    /**
     * Date and time of VPC's creation (RFC 3339 format).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
     */
    public readonly enableRouting!: pulumi.Output<boolean>;
    /**
     * Defines whether the VPC is the default one for its Project.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * The name of the VPC. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the VPC is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the VPC is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region of the VPC.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The tags associated with the VPC.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Date and time of VPC's last update (RFC 3339 format).
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Vpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcArgs | VpcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enableRouting"] = state ? state.enableRouting : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as VpcArgs | undefined;
            resourceInputs["enableRouting"] = args ? args.enableRouting : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vpc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vpc resources.
 */
export interface VpcState {
    /**
     * Date and time of VPC's creation (RFC 3339 format).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
     */
    enableRouting?: pulumi.Input<boolean>;
    /**
     * Defines whether the VPC is the default one for its Project.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of the VPC. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the VPC is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the VPC is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the VPC.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the VPC.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date and time of VPC's last update (RFC 3339 format).
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vpc resource.
 */
export interface VpcArgs {
    /**
     * Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
     */
    enableRouting?: pulumi.Input<boolean>;
    /**
     * The name of the VPC. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the VPC is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the VPC.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the VPC.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
