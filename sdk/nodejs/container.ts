// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Container.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/faq/serverless-containers/).
 *
 * For more details about the limitation check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/).
 *
 * You can check also our [containers guide](https://www.scaleway.com/en/docs/compute/containers/concepts/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainContainerNamespace = new scaleway.ContainerNamespace("mainContainerNamespace", {description: "test container"});
 * const mainContainer = new scaleway.Container("mainContainer", {
 *     description: "environment variables test",
 *     namespaceId: mainContainerNamespace.id,
 *     registryImage: pulumi.interpolate`${mainContainerNamespace.registryEndpoint}/alpine:test`,
 *     port: 9997,
 *     cpuLimit: 140,
 *     memoryLimit: 256,
 *     minScale: 3,
 *     maxScale: 5,
 *     timeout: 600,
 *     maxConcurrency: 80,
 *     privacy: "private",
 *     protocol: "h2c",
 *     deploy: true,
 *     environmentVariables: {
 *         foo: "var",
 *     },
 *     secretEnvironmentVariables: {
 *         key: "secret",
 *     },
 * });
 * ```
 * ## Protocols
 *
 * The supported protocols are:
 *
 * * `h2c`: HTTP/2 over TCP.
 * * `http1`: Hypertext Transfer Protocol.
 *
 * **Important:** For details about the protocols check [this](https://httpd.apache.org/docs/2.4/howto/http2.html)
 *
 * ## Privacy
 *
 * By default, creating a container will make it `public`, meaning that anybody knowing the endpoint could execute it.
 * A container can be made `private` with the privacy parameter.
 *
 * Please check our [authentication](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) section
 *
 * ## Memory and vCPUs configuration
 *
 * The vCPU represents a portion or share of the underlying, physical CPU that is assigned to a particular virtual machine (VM).
 *
 * You may decide how much computing resources to allocate to each container.
 * The `memoryLimit` (in MB) must correspond with the right amount of vCPU.
 *
 * **Important:** The right choice for your container's resources is very important, as you will be billed based on compute usage over time and the number of Containers executions.
 *
 * Please check our [price](https://www.scaleway.com/en/docs/faq/serverless-containers/#prices) section for more details.
 *
 * | Memory (in MB) | vCPU |
 * |----------------|------|
 * | 128            | 70m  |
 * | 256            | 140m |
 * | 512            | 280m |
 * | 1024           | 560m |
 *
 * **Note:** 560mCPU accounts roughly for half of one CPU power of a Scaleway General Purpose instance
 *
 * ## Import
 *
 * Container can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/container:Container main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState, opts?: pulumi.CustomResourceOptions): Container {
        return new Container(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/container:Container';

    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Container {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Container.__pulumiType;
    }

    /**
     * The amount of vCPU computing resources to allocate to each container. Defaults to 70.
     */
    public readonly cpuLimit!: pulumi.Output<number>;
    /**
     * The cron status of the container.
     */
    public /*out*/ readonly cronStatus!: pulumi.Output<string>;
    /**
     * Boolean controlling whether the container is on a production environment.
     */
    public readonly deploy!: pulumi.Output<boolean | undefined>;
    /**
     * The description of the container.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The native domain name of the container
     */
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    /**
     * The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string}>;
    /**
     * The error message of the container.
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * HTTP traffic configuration
     */
    public readonly httpOption!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     */
    public readonly maxConcurrency!: pulumi.Output<number>;
    /**
     * The maximum of number of instances this container can scale to. Default to 20.
     */
    public readonly maxScale!: pulumi.Output<number>;
    /**
     * The memory computing resources in MB to allocate to each container. Defaults to 128.
     */
    public readonly memoryLimit!: pulumi.Output<number>;
    /**
     * The minimum of running container instances continuously. Defaults to 0.
     */
    public readonly minScale!: pulumi.Output<number>;
    /**
     * The unique name of the container name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The container namespace ID of the container.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * The port to expose the container. Defaults to 8080.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     */
    public readonly privacy!: pulumi.Output<string | undefined>;
    /**
     * The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * (Defaults to provider `region`) The region in which the container was created.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
     */
    public readonly registryImage!: pulumi.Output<string>;
    /**
     * The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
     */
    public readonly registrySha256!: pulumi.Output<string | undefined>;
    /**
     * The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
     */
    public readonly secretEnvironmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The container status.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     */
    public readonly timeout!: pulumi.Output<number>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerArgs | ContainerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerState | undefined;
            resourceInputs["cpuLimit"] = state ? state.cpuLimit : undefined;
            resourceInputs["cronStatus"] = state ? state.cronStatus : undefined;
            resourceInputs["deploy"] = state ? state.deploy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["errorMessage"] = state ? state.errorMessage : undefined;
            resourceInputs["httpOption"] = state ? state.httpOption : undefined;
            resourceInputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            resourceInputs["maxScale"] = state ? state.maxScale : undefined;
            resourceInputs["memoryLimit"] = state ? state.memoryLimit : undefined;
            resourceInputs["minScale"] = state ? state.minScale : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privacy"] = state ? state.privacy : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registryImage"] = state ? state.registryImage : undefined;
            resourceInputs["registrySha256"] = state ? state.registrySha256 : undefined;
            resourceInputs["secretEnvironmentVariables"] = state ? state.secretEnvironmentVariables : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as ContainerArgs | undefined;
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            resourceInputs["cpuLimit"] = args ? args.cpuLimit : undefined;
            resourceInputs["deploy"] = args ? args.deploy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["httpOption"] = args ? args.httpOption : undefined;
            resourceInputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            resourceInputs["maxScale"] = args ? args.maxScale : undefined;
            resourceInputs["memoryLimit"] = args ? args.memoryLimit : undefined;
            resourceInputs["minScale"] = args ? args.minScale : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["privacy"] = args ? args.privacy : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["registryImage"] = args ? args.registryImage : undefined;
            resourceInputs["registrySha256"] = args ? args.registrySha256 : undefined;
            resourceInputs["secretEnvironmentVariables"] = args?.secretEnvironmentVariables ? pulumi.secret(args.secretEnvironmentVariables) : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["cronStatus"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secretEnvironmentVariables"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Container.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Container resources.
 */
export interface ContainerState {
    /**
     * The amount of vCPU computing resources to allocate to each container. Defaults to 70.
     */
    cpuLimit?: pulumi.Input<number>;
    /**
     * The cron status of the container.
     */
    cronStatus?: pulumi.Input<string>;
    /**
     * Boolean controlling whether the container is on a production environment.
     */
    deploy?: pulumi.Input<boolean>;
    /**
     * The description of the container.
     */
    description?: pulumi.Input<string>;
    /**
     * The native domain name of the container
     */
    domainName?: pulumi.Input<string>;
    /**
     * The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The error message of the container.
     */
    errorMessage?: pulumi.Input<string>;
    /**
     * HTTP traffic configuration
     */
    httpOption?: pulumi.Input<string>;
    /**
     * The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * The maximum of number of instances this container can scale to. Default to 20.
     */
    maxScale?: pulumi.Input<number>;
    /**
     * The memory computing resources in MB to allocate to each container. Defaults to 128.
     */
    memoryLimit?: pulumi.Input<number>;
    /**
     * The minimum of running container instances continuously. Defaults to 0.
     */
    minScale?: pulumi.Input<number>;
    /**
     * The unique name of the container name.
     */
    name?: pulumi.Input<string>;
    /**
     * The container namespace ID of the container.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The port to expose the container. Defaults to 8080.
     */
    port?: pulumi.Input<number>;
    /**
     * The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     */
    privacy?: pulumi.Input<string>;
    /**
     * The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     */
    protocol?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in which the container was created.
     */
    region?: pulumi.Input<string>;
    /**
     * The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
     */
    registryImage?: pulumi.Input<string>;
    /**
     * The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
     */
    registrySha256?: pulumi.Input<string>;
    /**
     * The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The container status.
     */
    status?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * The amount of vCPU computing resources to allocate to each container. Defaults to 70.
     */
    cpuLimit?: pulumi.Input<number>;
    /**
     * Boolean controlling whether the container is on a production environment.
     */
    deploy?: pulumi.Input<boolean>;
    /**
     * The description of the container.
     */
    description?: pulumi.Input<string>;
    /**
     * The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * HTTP traffic configuration
     */
    httpOption?: pulumi.Input<string>;
    /**
     * The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * The maximum of number of instances this container can scale to. Default to 20.
     */
    maxScale?: pulumi.Input<number>;
    /**
     * The memory computing resources in MB to allocate to each container. Defaults to 128.
     */
    memoryLimit?: pulumi.Input<number>;
    /**
     * The minimum of running container instances continuously. Defaults to 0.
     */
    minScale?: pulumi.Input<number>;
    /**
     * The unique name of the container name.
     */
    name?: pulumi.Input<string>;
    /**
     * The container namespace ID of the container.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * The port to expose the container. Defaults to 8080.
     */
    port?: pulumi.Input<number>;
    /**
     * The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     */
    privacy?: pulumi.Input<string>;
    /**
     * The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
     */
    registryImage?: pulumi.Input<string>;
    /**
     * The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
     */
    registrySha256?: pulumi.Input<string>;
    /**
     * The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The container status.
     */
    status?: pulumi.Input<string>;
    /**
     * The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     */
    timeout?: pulumi.Input<number>;
}
