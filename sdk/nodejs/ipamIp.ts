// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Books and manages Scaleway IPAM IPs.
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const vpc01 = new scaleway.Vpc("vpc01", {});
 * const pn01 = new scaleway.VpcPrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const ip01 = new scaleway.IpamIp("ip01", {sources: [{
 *     privateNetworkId: pn01.id,
 * }]});
 * ```
 *
 * ### Request a specific IPv4
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const vpc01 = new scaleway.Vpc("vpc01", {});
 * const pn01 = new scaleway.VpcPrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const ip01 = new scaleway.IpamIp("ip01", {
 *     address: "172.16.32.7/22",
 *     sources: [{
 *         privateNetworkId: pn01.id,
 *     }],
 * });
 * ```
 *
 * ### Request an IPv6
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const vpc01 = new scaleway.Vpc("vpc01", {});
 * const pn01 = new scaleway.VpcPrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv6Subnets: [{
 *         subnet: "fd46:78ab:30b8:177c::/64",
 *     }],
 * });
 * const ip01 = new scaleway.IpamIp("ip01", {
 *     isIpv6: true,
 *     sources: [{
 *         privateNetworkId: pn01.id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * IPAM IPs can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/ipamIp:IpamIp ip_demo fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class IpamIp extends pulumi.CustomResource {
    /**
     * Get an existing IpamIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpamIpState, opts?: pulumi.CustomResourceOptions): IpamIp {
        return new IpamIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/ipamIp:IpamIp';

    /**
     * Returns true if the given object is an instance of IpamIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpamIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpamIp.__pulumiType;
    }

    /**
     * Request a specific IP in the requested source pool.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Date and time of IP's creation (RFC 3339 format).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Defines whether to request an IPv6 instead of an IPv4.
     */
    public readonly isIpv6!: pulumi.Output<boolean | undefined>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region of the IP.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The IP resource.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.IpamIpResource[]>;
    /**
     * The reverses DNS for this IP.
     */
    public /*out*/ readonly reverses!: pulumi.Output<outputs.IpamIpReverse[]>;
    /**
     * The source in which to book the IP.
     */
    public readonly sources!: pulumi.Output<outputs.IpamIpSource[]>;
    /**
     * The tags associated with the IP.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Date and time of IP's last update (RFC 3339 format).
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The zone of the IP.
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a IpamIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpamIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpamIpArgs | IpamIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpamIpState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["isIpv6"] = state ? state.isIpv6 : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["reverses"] = state ? state.reverses : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as IpamIpArgs | undefined;
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["isIpv6"] = args ? args.isIpv6 : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["reverses"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpamIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpamIp resources.
 */
export interface IpamIpState {
    /**
     * Request a specific IP in the requested source pool.
     */
    address?: pulumi.Input<string>;
    /**
     * Date and time of IP's creation (RFC 3339 format).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Defines whether to request an IPv6 instead of an IPv4.
     */
    isIpv6?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the IP.
     */
    region?: pulumi.Input<string>;
    /**
     * The IP resource.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.IpamIpResource>[]>;
    /**
     * The reverses DNS for this IP.
     */
    reverses?: pulumi.Input<pulumi.Input<inputs.IpamIpReverse>[]>;
    /**
     * The source in which to book the IP.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.IpamIpSource>[]>;
    /**
     * The tags associated with the IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date and time of IP's last update (RFC 3339 format).
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The zone of the IP.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpamIp resource.
 */
export interface IpamIpArgs {
    /**
     * Request a specific IP in the requested source pool.
     */
    address?: pulumi.Input<string>;
    /**
     * Defines whether to request an IPv6 instead of an IPv4.
     */
    isIpv6?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the IP is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the IP.
     */
    region?: pulumi.Input<string>;
    /**
     * The source in which to book the IP.
     */
    sources: pulumi.Input<pulumi.Input<inputs.IpamIpSource>[]>;
    /**
     * The tags associated with the IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
