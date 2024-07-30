// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load Balancer IP addresses.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const ip = new scaleway.LbIp("ip", {reverse: "my-reverse.com"});
 * ```
 *
 * ### With IPv6
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const ipv6 = new scaleway.LbIp("ipv6", {isIpv6: true});
 * ```
 *
 * ## Import
 *
 * IPs can be imported using `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/lbIp:LbIp ip01 fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class LbIp extends pulumi.CustomResource {
    /**
     * Get an existing LbIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbIpState, opts?: pulumi.CustomResourceOptions): LbIp {
        return new LbIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lbIp:LbIp';

    /**
     * Returns true if the given object is an instance of LbIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbIp.__pulumiType;
    }

    /**
     * The IP address
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * If true, creates a flexible IP with an IPv6 address.
     */
    public readonly isIpv6!: pulumi.Output<boolean | undefined>;
    /**
     * The associated Load Balancer ID if any
     */
    public /*out*/ readonly lbId!: pulumi.Output<string>;
    /**
     * The organizationId you want to attach the resource to
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the Project the IP is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region of the resource
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The reverse domain associated with this IP.
     */
    public readonly reverse!: pulumi.Output<string>;
    /**
     * The tags associated with this IP.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a LbIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LbIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbIpArgs | LbIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbIpState | undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["isIpv6"] = state ? state.isIpv6 : undefined;
            resourceInputs["lbId"] = state ? state.lbId : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["reverse"] = state ? state.reverse : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as LbIpArgs | undefined;
            resourceInputs["isIpv6"] = args ? args.isIpv6 : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["reverse"] = args ? args.reverse : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["lbId"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbIp resources.
 */
export interface LbIpState {
    /**
     * The IP address
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * If true, creates a flexible IP with an IPv6 address.
     */
    isIpv6?: pulumi.Input<boolean>;
    /**
     * The associated Load Balancer ID if any
     */
    lbId?: pulumi.Input<string>;
    /**
     * The organizationId you want to attach the resource to
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the Project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region of the resource
     */
    region?: pulumi.Input<string>;
    /**
     * The reverse domain associated with this IP.
     */
    reverse?: pulumi.Input<string>;
    /**
     * The tags associated with this IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbIp resource.
 */
export interface LbIpArgs {
    /**
     * If true, creates a flexible IP with an IPv6 address.
     */
    isIpv6?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the Project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The reverse domain associated with this IP.
     */
    reverse?: pulumi.Input<string>;
    /**
     * The tags associated with this IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the IP should be reserved.
     */
    zone?: pulumi.Input<string>;
}
