// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a container namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.getContainerNamespace({
 *     name: "my-namespace-name",
 * });
 * const byId = scaleway.getContainerNamespace({
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getContainerNamespace(args?: GetContainerNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerNamespaceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getContainerNamespace:getContainerNamespace", {
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerNamespace.
 */
export interface GetContainerNamespaceArgs {
    /**
     * The namespace name.
     * Only one of `name` and `namespaceId` should be specified.
     */
    name?: string;
    /**
     * The namespace id.
     * Only one of `name` and `namespaceId` should be specified.
     */
    namespaceId?: string;
    /**
     * `projectId`) The ID of the project the namespace is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the namespace exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getContainerNamespace.
 */
export interface GetContainerNamespaceResult {
    /**
     * The description of the namespace.
     */
    readonly description: string;
    readonly destroyRegistry: boolean;
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
     * The organization ID the namespace is associated with.
     */
    readonly organizationId: string;
    readonly projectId?: string;
    readonly region?: string;
    /**
     * The registry endpoint of the namespace.
     */
    readonly registryEndpoint: string;
    /**
     * The registry namespace ID of the namespace.
     */
    readonly registryNamespaceId: string;
    readonly secretEnvironmentVariables: {[key: string]: string};
}
/**
 * Gets information about a container namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.getContainerNamespace({
 *     name: "my-namespace-name",
 * });
 * const byId = scaleway.getContainerNamespace({
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getContainerNamespaceOutput(args?: GetContainerNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContainerNamespaceResult> {
    return pulumi.output(args).apply((a: any) => getContainerNamespace(a, opts))
}

/**
 * A collection of arguments for invoking getContainerNamespace.
 */
export interface GetContainerNamespaceOutputArgs {
    /**
     * The namespace name.
     * Only one of `name` and `namespaceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace id.
     * Only one of `name` and `namespaceId` should be specified.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the namespace is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the namespace exists.
     */
    region?: pulumi.Input<string>;
}
