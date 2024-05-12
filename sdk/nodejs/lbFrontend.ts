// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load-Balancer Frontends. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const frontend01 = new scaleway.LbFrontend("frontend01", {
 *     lbId: scaleway_lb.lb01.id,
 *     backendId: scaleway_lb_backend.backend01.id,
 *     inboundPort: 80,
 * });
 * ```
 *
 * ## With ACLs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const frontend01 = new scaleway.LbFrontend("frontend01", {
 *     lbId: scaleway_lb.lb01.id,
 *     backendId: scaleway_lb_backend.backend01.id,
 *     inboundPort: 80,
 *     acls: [
 *         {
 *             name: "blacklist wellknwon IPs",
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 ipSubnets: [
 *                     "192.168.0.1",
 *                     "192.168.0.2",
 *                     "192.168.10.0/24",
 *                 ],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "deny",
 *             },
 *             match: {
 *                 ipSubnets: ["51.51.51.51"],
 *                 httpFilter: "regex",
 *                 httpFilterValues: ["^foo*bar$"],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: [
 *                     "foo",
 *                     "bar",
 *                 ],
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: ["hi"],
 *                 invert: true,
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "allow",
 *             },
 *             match: {
 *                 httpFilter: "http_header_match",
 *                 httpFilterValues: "foo",
 *                 httpFilterOption: "bar",
 *             },
 *         },
 *         {
 *             action: {
 *                 type: "redirect",
 *                 redirects: [{
 *                     type: "location",
 *                     target: "https://example.com",
 *                     code: 307,
 *                 }],
 *             },
 *             match: {
 *                 ipSubnets: ["10.0.0.10"],
 *                 httpFilter: "path_begin",
 *                 httpFilterValues: [
 *                     "foo",
 *                     "bar",
 *                 ],
 *             },
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/lbFrontend:LbFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class LbFrontend extends pulumi.CustomResource {
    /**
     * Get an existing LbFrontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbFrontendState, opts?: pulumi.CustomResourceOptions): LbFrontend {
        return new LbFrontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lbFrontend:LbFrontend';

    /**
     * Returns true if the given object is an instance of LbFrontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbFrontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbFrontend.__pulumiType;
    }

    /**
     * A list of ACL rules to apply to the load-balancer frontend.  Defined below.
     */
    public readonly acls!: pulumi.Output<outputs.LbFrontendAcl[] | undefined>;
    /**
     * The load-balancer backend ID this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    public readonly backendId!: pulumi.Output<string>;
    /**
     * (Deprecated) first certificate ID used by the frontend.
     *
     * @deprecated Please use certificate_ids
     */
    public /*out*/ readonly certificateId!: pulumi.Output<string>;
    /**
     * List of Certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    public readonly certificateIds!: pulumi.Output<string[] | undefined>;
    /**
     * Activates HTTP/3 protocol.
     */
    public readonly enableHttp3!: pulumi.Output<boolean | undefined>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
     */
    public readonly externalAcls!: pulumi.Output<boolean | undefined>;
    /**
     * TCP port to listen on the front side.
     */
    public readonly inboundPort!: pulumi.Output<number>;
    /**
     * The load-balancer ID this frontend is attached to.
     */
    public readonly lbId!: pulumi.Output<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Maximum inactivity time on the client side. (e.g.: `1s`)
     */
    public readonly timeoutClient!: pulumi.Output<string | undefined>;

    /**
     * Create a LbFrontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbFrontendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbFrontendArgs | LbFrontendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbFrontendState | undefined;
            resourceInputs["acls"] = state ? state.acls : undefined;
            resourceInputs["backendId"] = state ? state.backendId : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["certificateIds"] = state ? state.certificateIds : undefined;
            resourceInputs["enableHttp3"] = state ? state.enableHttp3 : undefined;
            resourceInputs["externalAcls"] = state ? state.externalAcls : undefined;
            resourceInputs["inboundPort"] = state ? state.inboundPort : undefined;
            resourceInputs["lbId"] = state ? state.lbId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["timeoutClient"] = state ? state.timeoutClient : undefined;
        } else {
            const args = argsOrState as LbFrontendArgs | undefined;
            if ((!args || args.backendId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendId'");
            }
            if ((!args || args.inboundPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inboundPort'");
            }
            if ((!args || args.lbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbId'");
            }
            resourceInputs["acls"] = args ? args.acls : undefined;
            resourceInputs["backendId"] = args ? args.backendId : undefined;
            resourceInputs["certificateIds"] = args ? args.certificateIds : undefined;
            resourceInputs["enableHttp3"] = args ? args.enableHttp3 : undefined;
            resourceInputs["externalAcls"] = args ? args.externalAcls : undefined;
            resourceInputs["inboundPort"] = args ? args.inboundPort : undefined;
            resourceInputs["lbId"] = args ? args.lbId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["timeoutClient"] = args ? args.timeoutClient : undefined;
            resourceInputs["certificateId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbFrontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbFrontend resources.
 */
export interface LbFrontendState {
    /**
     * A list of ACL rules to apply to the load-balancer frontend.  Defined below.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.LbFrontendAcl>[]>;
    /**
     * The load-balancer backend ID this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    backendId?: pulumi.Input<string>;
    /**
     * (Deprecated) first certificate ID used by the frontend.
     *
     * @deprecated Please use certificate_ids
     */
    certificateId?: pulumi.Input<string>;
    /**
     * List of Certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Activates HTTP/3 protocol.
     */
    enableHttp3?: pulumi.Input<boolean>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
     */
    externalAcls?: pulumi.Input<boolean>;
    /**
     * TCP port to listen on the front side.
     */
    inboundPort?: pulumi.Input<number>;
    /**
     * The load-balancer ID this frontend is attached to.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum inactivity time on the client side. (e.g.: `1s`)
     */
    timeoutClient?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbFrontend resource.
 */
export interface LbFrontendArgs {
    /**
     * A list of ACL rules to apply to the load-balancer frontend.  Defined below.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.LbFrontendAcl>[]>;
    /**
     * The load-balancer backend ID this frontend is attached to.
     *
     * > **Important:** Updates to `lbId` or `backendId` will recreate the frontend.
     */
    backendId: pulumi.Input<string>;
    /**
     * List of Certificate IDs that should be used by the frontend.
     *
     * > **Important:** Certificates are not allowed on port 80.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Activates HTTP/3 protocol.
     */
    enableHttp3?: pulumi.Input<boolean>;
    /**
     * A boolean to specify whether to use lb_acl.
     * If `externalAcls` is set to `true`, `acl` can not be set directly in the lb frontend.
     */
    externalAcls?: pulumi.Input<boolean>;
    /**
     * TCP port to listen on the front side.
     */
    inboundPort: pulumi.Input<number>;
    /**
     * The load-balancer ID this frontend is attached to.
     */
    lbId: pulumi.Input<string>;
    /**
     * The ACL name. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum inactivity time on the client side. (e.g.: `1s`)
     */
    timeoutClient?: pulumi.Input<string>;
}
