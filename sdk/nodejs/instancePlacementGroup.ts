// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Compute Instance Placement Groups. For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-list-placement-groups).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const availabilityGroup = new scaleway.InstancePlacementGroup("availabilityGroup", {});
 * ```
 *
 * ## Import
 *
 * Placement groups can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/instancePlacementGroup:InstancePlacementGroup availability_group fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstancePlacementGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstancePlacementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstancePlacementGroupState, opts?: pulumi.CustomResourceOptions): InstancePlacementGroup {
        return new InstancePlacementGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instancePlacementGroup:InstancePlacementGroup';

    /**
     * Returns true if the given object is an instance of InstancePlacementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstancePlacementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstancePlacementGroup.__pulumiType;
    }

    /**
     * The name of the placement group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the placement group is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The [policy mode](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-group) of the placement group. Possible values are: `optional` or `enforced`.
     */
    public readonly policyMode!: pulumi.Output<string | undefined>;
    /**
     * Is true when the policy is respected.
     */
    public /*out*/ readonly policyRespected!: pulumi.Output<boolean>;
    /**
     * The [policy type](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-grou) of the placement group. Possible values are: `lowLatency` or `maxAvailability`.
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * `projectId`) The ID of the project the placement group is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * A list of tags to apply to the placement group.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * `zone`) The zone in which the placement group should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstancePlacementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstancePlacementGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstancePlacementGroupArgs | InstancePlacementGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstancePlacementGroupState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["policyMode"] = state ? state.policyMode : undefined;
            resourceInputs["policyRespected"] = state ? state.policyRespected : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstancePlacementGroupArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyMode"] = args ? args.policyMode : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["policyRespected"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstancePlacementGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstancePlacementGroup resources.
 */
export interface InstancePlacementGroupState {
    /**
     * The name of the placement group.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the placement group is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The [policy mode](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-group) of the placement group. Possible values are: `optional` or `enforced`.
     */
    policyMode?: pulumi.Input<string>;
    /**
     * Is true when the policy is respected.
     */
    policyRespected?: pulumi.Input<boolean>;
    /**
     * The [policy type](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-grou) of the placement group. Possible values are: `lowLatency` or `maxAvailability`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the placement group is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the placement group.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the placement group should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstancePlacementGroup resource.
 */
export interface InstancePlacementGroupArgs {
    /**
     * The name of the placement group.
     */
    name?: pulumi.Input<string>;
    /**
     * The [policy mode](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-group) of the placement group. Possible values are: `optional` or `enforced`.
     */
    policyMode?: pulumi.Input<string>;
    /**
     * The [policy type](https://www.scaleway.com/en/developers/api/instance/#path-placement-groups-create-a-placement-grou) of the placement group. Possible values are: `lowLatency` or `maxAvailability`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the placement group is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the placement group.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the placement group should be created.
     */
    zone?: pulumi.Input<string>;
}
