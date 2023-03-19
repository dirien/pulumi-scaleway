// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Function Token.
 * For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/#tokens-26b085).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainFunctionNamespace = new scaleway.FunctionNamespace("mainFunctionNamespace", {});
 * const mainFunction = new scaleway.Function("mainFunction", {
 *     namespaceId: mainFunctionNamespace.id,
 *     runtime: "go118",
 *     handler: "Handle",
 *     privacy: "private",
 * });
 * // Namespace Token
 * const namespace = new scaleway.FunctionToken("namespace", {
 *     namespaceId: mainFunctionNamespace.id,
 *     expiresAt: "2022-10-18T11:35:15+02:00",
 * });
 * // Function Token
 * const _function = new scaleway.FunctionToken("function", {functionId: mainFunction.id});
 * ```
 *
 * ## Import
 *
 * Tokens can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/functionToken:FunctionToken main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class FunctionToken extends pulumi.CustomResource {
    /**
     * Get an existing FunctionToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionTokenState, opts?: pulumi.CustomResourceOptions): FunctionToken {
        return new FunctionToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/functionToken:FunctionToken';

    /**
     * Returns true if the given object is an instance of FunctionToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionToken.__pulumiType;
    }

    /**
     * The description of the token.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The expiration date of the token.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The ID of the function.
     */
    public readonly functionId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the function namespace.
     */
    public readonly namespaceId!: pulumi.Output<string | undefined>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;

    /**
     * Create a FunctionToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FunctionTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionTokenArgs | FunctionTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionTokenState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["functionId"] = state ? state.functionId : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as FunctionTokenArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["functionId"] = args ? args.functionId : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FunctionToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionToken resources.
 */
export interface FunctionTokenState {
    /**
     * The description of the token.
     */
    description?: pulumi.Input<string>;
    /**
     * The expiration date of the token.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID of the function.
     */
    functionId?: pulumi.Input<string>;
    /**
     * The ID of the function namespace.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The token.
     */
    token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionToken resource.
 */
export interface FunctionTokenArgs {
    /**
     * The description of the token.
     */
    description?: pulumi.Input<string>;
    /**
     * The expiration date of the token.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID of the function.
     */
    functionId?: pulumi.Input<string>;
    /**
     * The ID of the function namespace.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
}
