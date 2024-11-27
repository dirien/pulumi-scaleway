// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.FunctionNamespace` data source is used to retrieve information about a Serverless Functions namespace.
 *
 * Refer to the Serverless Functions [product documentation](https://www.scaleway.com/en/docs/serverless/functions/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/) for more information.
 *
 * ## Retrieve a Serverless Functions namespace
 *
 * The following commands allow you to:
 *
 * - retrieve a namespace by its name
 * - retrieve a namespace by its ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myNamespace = scaleway.getFunctionNamespace({
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getFunctionNamespace(args?: GetFunctionNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionNamespaceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getFunctionNamespace:getFunctionNamespace", {
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionNamespace.
 */
export interface GetFunctionNamespaceArgs {
    /**
     * The name of the namespace. Only one of `name` and `namespaceId` should be specified.
     */
    name?: string;
    /**
     * The unique identifier of the namespace. Only one of `name` and `namespaceId` should be specified.
     */
    namespaceId?: string;
    /**
     * `projectId`) The unique identifier of the project with which the namespace is associated.
     */
    projectId?: string;
    /**
     * `region`) The region in which the namespace exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getFunctionNamespace.
 */
export interface GetFunctionNamespaceResult {
    /**
     * The description of the namespace.
     */
    readonly description: string;
    /**
     * The environment variables of the namespace.
     */
    readonly environmentVariables: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly namespaceId?: string;
    /**
     * The unique identifier of the organization with which the namespace is associated.
     */
    readonly organizationId: string;
    readonly projectId?: string;
    readonly region?: string;
    /**
     * The registry endpoint of the namespace.
     */
    readonly registryEndpoint: string;
    /**
     * The unique identifier of the registry namespace of the Serverless Functions namespace.
     */
    readonly registryNamespaceId: string;
    readonly secretEnvironmentVariables: {[key: string]: string};
}
/**
 * The `scaleway.FunctionNamespace` data source is used to retrieve information about a Serverless Functions namespace.
 *
 * Refer to the Serverless Functions [product documentation](https://www.scaleway.com/en/docs/serverless/functions/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/) for more information.
 *
 * ## Retrieve a Serverless Functions namespace
 *
 * The following commands allow you to:
 *
 * - retrieve a namespace by its name
 * - retrieve a namespace by its ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myNamespace = scaleway.getFunctionNamespace({
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getFunctionNamespaceOutput(args?: GetFunctionNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionNamespaceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getFunctionNamespace:getFunctionNamespace", {
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionNamespace.
 */
export interface GetFunctionNamespaceOutputArgs {
    /**
     * The name of the namespace. Only one of `name` and `namespaceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The unique identifier of the namespace. Only one of `name` and `namespaceId` should be specified.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * `projectId`) The unique identifier of the project with which the namespace is associated.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the namespace exists.
     */
    region?: pulumi.Input<string>;
}
