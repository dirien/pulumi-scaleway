// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load Balancers.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/concepts/#load-balancers) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-list-load-balancers).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.LbIp("main", {zone: "fr-par-1"});
 * const base = new scaleway.Lb("base", {
 *     ipIds: [main.id],
 *     zone: main.zone,
 *     type: "LB-S",
 * });
 * ```
 *
 * ### Private LB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const base = new scaleway.Lb("base", {
 *     assignFlexibleIp: false,
 *     type: "LB-S",
 * });
 * ```
 *
 * ### With IPv6
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const v4 = new scaleway.LbIp("v4", {});
 * const v6 = new scaleway.LbIp("v6", {isIpv6: true});
 * const main = new scaleway.Lb("main", {
 *     ipIds: [
 *         v4.id,
 *         v6.id,
 *     ],
 *     type: "LB-S",
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancers can be imported using `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/lb:Lb main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 *
 * Be aware that you will also need to import the `scaleway_lb_ip` resource.
 */
export class Lb extends pulumi.CustomResource {
    /**
     * Get an existing Lb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbState, opts?: pulumi.CustomResourceOptions): Lb {
        return new Lb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lb:Lb';

    /**
     * Returns true if the given object is an instance of Lb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lb.__pulumiType;
    }

    /**
     * Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
     */
    public readonly assignFlexibleIp!: pulumi.Output<boolean | undefined>;
    /**
     * Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
     */
    public readonly assignFlexibleIpv6!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the Load Balancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Load Balancer public IPv4 address.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Please use `ipIds`. The ID of the associated Load Balancer IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the Load Balancer.
     *
     * @deprecated Please use ip_ids
     */
    public readonly ipId!: pulumi.Output<string>;
    /**
     * The List of IP IDs to attach to the Load Balancer.
     *
     * > **Important:** Make sure to use a `scaleway.LbIp` resource to create the IPs.
     */
    public readonly ipIds!: pulumi.Output<string[]>;
    /**
     * The Load Balancer public IPv6 address.
     */
    public /*out*/ readonly ipv6Address!: pulumi.Output<string>;
    /**
     * The name of the Load Balancer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the Organization ID the Load Balancer is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * List of private network to connect with your load balancer.
     */
    public readonly privateNetworks!: pulumi.Output<outputs.LbPrivateNetwork[] | undefined>;
    /**
     * `projectId`) The ID of the Project the Load Balancer is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region of the resource
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The `releaseIp` allow the release of the IP address associated with the Load Balancer.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    public readonly releaseIp!: pulumi.Output<boolean | undefined>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    public readonly sslCompatibilityLevel!: pulumi.Output<string | undefined>;
    /**
     * The tags associated with the Load Balancer.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of the Load Balancer. Please check the migration section to upgrade the type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * `zone`) The zone of the Load Balancer.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Lb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbArgs | LbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbState | undefined;
            resourceInputs["assignFlexibleIp"] = state ? state.assignFlexibleIp : undefined;
            resourceInputs["assignFlexibleIpv6"] = state ? state.assignFlexibleIpv6 : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ipId"] = state ? state.ipId : undefined;
            resourceInputs["ipIds"] = state ? state.ipIds : undefined;
            resourceInputs["ipv6Address"] = state ? state.ipv6Address : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["privateNetworks"] = state ? state.privateNetworks : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["releaseIp"] = state ? state.releaseIp : undefined;
            resourceInputs["sslCompatibilityLevel"] = state ? state.sslCompatibilityLevel : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as LbArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["assignFlexibleIp"] = args ? args.assignFlexibleIp : undefined;
            resourceInputs["assignFlexibleIpv6"] = args ? args.assignFlexibleIpv6 : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipId"] = args ? args.ipId : undefined;
            resourceInputs["ipIds"] = args ? args.ipIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateNetworks"] = args ? args.privateNetworks : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["releaseIp"] = args ? args.releaseIp : undefined;
            resourceInputs["sslCompatibilityLevel"] = args ? args.sslCompatibilityLevel : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["ipv6Address"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Lb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lb resources.
 */
export interface LbState {
    /**
     * Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
     */
    assignFlexibleIp?: pulumi.Input<boolean>;
    /**
     * Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
     */
    assignFlexibleIpv6?: pulumi.Input<boolean>;
    /**
     * The description of the Load Balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The Load Balancer public IPv4 address.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Please use `ipIds`. The ID of the associated Load Balancer IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the Load Balancer.
     *
     * @deprecated Please use ip_ids
     */
    ipId?: pulumi.Input<string>;
    /**
     * The List of IP IDs to attach to the Load Balancer.
     *
     * > **Important:** Make sure to use a `scaleway.LbIp` resource to create the IPs.
     */
    ipIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Load Balancer public IPv6 address.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * The name of the Load Balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the Organization ID the Load Balancer is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * List of private network to connect with your load balancer.
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.LbPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the Project the Load Balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region of the resource
     */
    region?: pulumi.Input<string>;
    /**
     * The `releaseIp` allow the release of the IP address associated with the Load Balancer.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    releaseIp?: pulumi.Input<boolean>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    sslCompatibilityLevel?: pulumi.Input<string>;
    /**
     * The tags associated with the Load Balancer.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the Load Balancer. Please check the migration section to upgrade the type.
     */
    type?: pulumi.Input<string>;
    /**
     * `zone`) The zone of the Load Balancer.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Lb resource.
 */
export interface LbArgs {
    /**
     * Defines whether to automatically assign a flexible public IPv4 to the Load Balancer.
     */
    assignFlexibleIp?: pulumi.Input<boolean>;
    /**
     * Defines whether to automatically assign a flexible public IPv6 to the Load Balancer.
     */
    assignFlexibleIpv6?: pulumi.Input<boolean>;
    /**
     * The description of the Load Balancer.
     */
    description?: pulumi.Input<string>;
    /**
     * Please use `ipIds`. The ID of the associated Load Balancer IP. See below.
     *
     * > **Important:** Updates to `ipId` will recreate the Load Balancer.
     *
     * @deprecated Please use ip_ids
     */
    ipId?: pulumi.Input<string>;
    /**
     * The List of IP IDs to attach to the Load Balancer.
     *
     * > **Important:** Make sure to use a `scaleway.LbIp` resource to create the IPs.
     */
    ipIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Load Balancer.
     */
    name?: pulumi.Input<string>;
    /**
     * List of private network to connect with your load balancer.
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.LbPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the Project the Load Balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The `releaseIp` allow the release of the IP address associated with the Load Balancer.
     *
     * @deprecated The resource ip will be destroyed by it's own resource. Please set this to `false`
     */
    releaseIp?: pulumi.Input<boolean>;
    /**
     * Enforces minimal SSL version (in SSL/TLS offloading context). Please check [possible values](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-create-a-load-balancer).
     */
    sslCompatibilityLevel?: pulumi.Input<string>;
    /**
     * The tags associated with the Load Balancer.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the Load Balancer. Please check the migration section to upgrade the type.
     */
    type: pulumi.Input<string>;
    /**
     * `zone`) The zone of the Load Balancer.
     */
    zone?: pulumi.Input<string>;
}
