// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Load-Balancer Backends.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const backend01 = new scaleway.LbBackend("backend01", {
 *     lbId: scaleway_lb.lb01.id,
 *     forwardProtocol: "http",
 *     forwardPort: 80,
 * });
 * ```
 *
 * ### With HTTP Health Check
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const backend01 = new scaleway.LbBackend("backend01", {
 *     lbId: scaleway_lb.lb01.id,
 *     forwardProtocol: "http",
 *     forwardPort: 80,
 *     healthCheckHttp: {
 *         uri: "www.test.com/health",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Load-Balancer backend can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/lbBackend:LbBackend backend01 fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class LbBackend extends pulumi.CustomResource {
    /**
     * Get an existing LbBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbBackendState, opts?: pulumi.CustomResourceOptions): LbBackend {
        return new LbBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/lbBackend:LbBackend';

    /**
     * Returns true if the given object is an instance of LbBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbBackend.__pulumiType;
    }

    /**
     * Scaleway S3 bucket website to be served in case all backend servers are down.
     * > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
     * e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
     */
    public readonly failoverHost!: pulumi.Output<string | undefined>;
    /**
     * User sessions will be forwarded to this port of backend servers.
     */
    public readonly forwardPort!: pulumi.Output<number>;
    /**
     * Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
     */
    public readonly forwardPortAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * Backend protocol. Possible values are: `tcp` or `http`.
     */
    public readonly forwardProtocol!: pulumi.Output<string>;
    /**
     * Interval between two HC requests.
     */
    public readonly healthCheckDelay!: pulumi.Output<string | undefined>;
    /**
     * This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    public readonly healthCheckHttp!: pulumi.Output<outputs.LbBackendHealthCheckHttp | undefined>;
    /**
     * This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    public readonly healthCheckHttps!: pulumi.Output<outputs.LbBackendHealthCheckHttps | undefined>;
    /**
     * Number of allowed failed HC requests before the backend server is marked down.
     */
    public readonly healthCheckMaxRetries!: pulumi.Output<number | undefined>;
    /**
     * Port the HC requests will be send to.
     */
    public readonly healthCheckPort!: pulumi.Output<number>;
    /**
     * This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    public readonly healthCheckTcp!: pulumi.Output<outputs.LbBackendHealthCheckTcp>;
    /**
     * Timeout before we consider a HC request failed.
     */
    public readonly healthCheckTimeout!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
     */
    public readonly ignoreSslServerVerify!: pulumi.Output<boolean | undefined>;
    /**
     * The load-balancer ID this backend is attached to.
     * > **Important:** Updates to `lbId` will recreate the backend.
     */
    public readonly lbId!: pulumi.Output<string>;
    /**
     * Maximum number of connections allowed per backend server.
     */
    public readonly maxConnections!: pulumi.Output<number | undefined>;
    /**
     * Number of retries when a backend server connection failed.
     */
    public readonly maxRetries!: pulumi.Output<number | undefined>;
    /**
     * The name of the load-balancer backend.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
     */
    public readonly onMarkedDownAction!: pulumi.Output<string | undefined>;
    /**
     * Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
     */
    public readonly proxyProtocol!: pulumi.Output<string | undefined>;
    /**
     * Whether to use another backend server on each attempt.
     */
    public readonly redispatchAttemptCount!: pulumi.Output<number | undefined>;
    /**
     * DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
     *
     * @deprecated Please use proxy_protocol instead
     */
    public readonly sendProxyV2!: pulumi.Output<boolean>;
    /**
     * List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
     */
    public readonly serverIps!: pulumi.Output<string[] | undefined>;
    /**
     * Enables SSL between load balancer and backend servers.
     */
    public readonly sslBridging!: pulumi.Output<boolean | undefined>;
    /**
     * The type of sticky sessions. The only current possible values are: `none`, `cookie` and `table`.
     */
    public readonly stickySessions!: pulumi.Output<string | undefined>;
    /**
     * Cookie name for sticky sessions. Only applicable when stickySessions is set to `cookie`.
     */
    public readonly stickySessionsCookieName!: pulumi.Output<string | undefined>;
    /**
     * Maximum initial server connection establishment time. (e.g.: `1s`)
     */
    public readonly timeoutConnect!: pulumi.Output<string | undefined>;
    /**
     * Maximum time for a request to be left pending in queue when `maxConnections` is reached. (e.g.: `1s`)
     */
    public readonly timeoutQueue!: pulumi.Output<string | undefined>;
    /**
     * Maximum server connection inactivity time. (e.g.: `1s`)
     */
    public readonly timeoutServer!: pulumi.Output<string | undefined>;
    /**
     * Maximum tunnel inactivity time. (e.g.: `1s`)
     */
    public readonly timeoutTunnel!: pulumi.Output<string | undefined>;

    /**
     * Create a LbBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbBackendArgs | LbBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbBackendState | undefined;
            resourceInputs["failoverHost"] = state ? state.failoverHost : undefined;
            resourceInputs["forwardPort"] = state ? state.forwardPort : undefined;
            resourceInputs["forwardPortAlgorithm"] = state ? state.forwardPortAlgorithm : undefined;
            resourceInputs["forwardProtocol"] = state ? state.forwardProtocol : undefined;
            resourceInputs["healthCheckDelay"] = state ? state.healthCheckDelay : undefined;
            resourceInputs["healthCheckHttp"] = state ? state.healthCheckHttp : undefined;
            resourceInputs["healthCheckHttps"] = state ? state.healthCheckHttps : undefined;
            resourceInputs["healthCheckMaxRetries"] = state ? state.healthCheckMaxRetries : undefined;
            resourceInputs["healthCheckPort"] = state ? state.healthCheckPort : undefined;
            resourceInputs["healthCheckTcp"] = state ? state.healthCheckTcp : undefined;
            resourceInputs["healthCheckTimeout"] = state ? state.healthCheckTimeout : undefined;
            resourceInputs["ignoreSslServerVerify"] = state ? state.ignoreSslServerVerify : undefined;
            resourceInputs["lbId"] = state ? state.lbId : undefined;
            resourceInputs["maxConnections"] = state ? state.maxConnections : undefined;
            resourceInputs["maxRetries"] = state ? state.maxRetries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onMarkedDownAction"] = state ? state.onMarkedDownAction : undefined;
            resourceInputs["proxyProtocol"] = state ? state.proxyProtocol : undefined;
            resourceInputs["redispatchAttemptCount"] = state ? state.redispatchAttemptCount : undefined;
            resourceInputs["sendProxyV2"] = state ? state.sendProxyV2 : undefined;
            resourceInputs["serverIps"] = state ? state.serverIps : undefined;
            resourceInputs["sslBridging"] = state ? state.sslBridging : undefined;
            resourceInputs["stickySessions"] = state ? state.stickySessions : undefined;
            resourceInputs["stickySessionsCookieName"] = state ? state.stickySessionsCookieName : undefined;
            resourceInputs["timeoutConnect"] = state ? state.timeoutConnect : undefined;
            resourceInputs["timeoutQueue"] = state ? state.timeoutQueue : undefined;
            resourceInputs["timeoutServer"] = state ? state.timeoutServer : undefined;
            resourceInputs["timeoutTunnel"] = state ? state.timeoutTunnel : undefined;
        } else {
            const args = argsOrState as LbBackendArgs | undefined;
            if ((!args || args.forwardPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forwardPort'");
            }
            if ((!args || args.forwardProtocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forwardProtocol'");
            }
            if ((!args || args.lbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbId'");
            }
            resourceInputs["failoverHost"] = args ? args.failoverHost : undefined;
            resourceInputs["forwardPort"] = args ? args.forwardPort : undefined;
            resourceInputs["forwardPortAlgorithm"] = args ? args.forwardPortAlgorithm : undefined;
            resourceInputs["forwardProtocol"] = args ? args.forwardProtocol : undefined;
            resourceInputs["healthCheckDelay"] = args ? args.healthCheckDelay : undefined;
            resourceInputs["healthCheckHttp"] = args ? args.healthCheckHttp : undefined;
            resourceInputs["healthCheckHttps"] = args ? args.healthCheckHttps : undefined;
            resourceInputs["healthCheckMaxRetries"] = args ? args.healthCheckMaxRetries : undefined;
            resourceInputs["healthCheckPort"] = args ? args.healthCheckPort : undefined;
            resourceInputs["healthCheckTcp"] = args ? args.healthCheckTcp : undefined;
            resourceInputs["healthCheckTimeout"] = args ? args.healthCheckTimeout : undefined;
            resourceInputs["ignoreSslServerVerify"] = args ? args.ignoreSslServerVerify : undefined;
            resourceInputs["lbId"] = args ? args.lbId : undefined;
            resourceInputs["maxConnections"] = args ? args.maxConnections : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onMarkedDownAction"] = args ? args.onMarkedDownAction : undefined;
            resourceInputs["proxyProtocol"] = args ? args.proxyProtocol : undefined;
            resourceInputs["redispatchAttemptCount"] = args ? args.redispatchAttemptCount : undefined;
            resourceInputs["sendProxyV2"] = args ? args.sendProxyV2 : undefined;
            resourceInputs["serverIps"] = args ? args.serverIps : undefined;
            resourceInputs["sslBridging"] = args ? args.sslBridging : undefined;
            resourceInputs["stickySessions"] = args ? args.stickySessions : undefined;
            resourceInputs["stickySessionsCookieName"] = args ? args.stickySessionsCookieName : undefined;
            resourceInputs["timeoutConnect"] = args ? args.timeoutConnect : undefined;
            resourceInputs["timeoutQueue"] = args ? args.timeoutQueue : undefined;
            resourceInputs["timeoutServer"] = args ? args.timeoutServer : undefined;
            resourceInputs["timeoutTunnel"] = args ? args.timeoutTunnel : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbBackend resources.
 */
export interface LbBackendState {
    /**
     * Scaleway S3 bucket website to be served in case all backend servers are down.
     * > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
     * e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
     */
    failoverHost?: pulumi.Input<string>;
    /**
     * User sessions will be forwarded to this port of backend servers.
     */
    forwardPort?: pulumi.Input<number>;
    /**
     * Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
     */
    forwardPortAlgorithm?: pulumi.Input<string>;
    /**
     * Backend protocol. Possible values are: `tcp` or `http`.
     */
    forwardProtocol?: pulumi.Input<string>;
    /**
     * Interval between two HC requests.
     */
    healthCheckDelay?: pulumi.Input<string>;
    /**
     * This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckHttp?: pulumi.Input<inputs.LbBackendHealthCheckHttp>;
    /**
     * This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckHttps?: pulumi.Input<inputs.LbBackendHealthCheckHttps>;
    /**
     * Number of allowed failed HC requests before the backend server is marked down.
     */
    healthCheckMaxRetries?: pulumi.Input<number>;
    /**
     * Port the HC requests will be send to.
     */
    healthCheckPort?: pulumi.Input<number>;
    /**
     * This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckTcp?: pulumi.Input<inputs.LbBackendHealthCheckTcp>;
    /**
     * Timeout before we consider a HC request failed.
     */
    healthCheckTimeout?: pulumi.Input<string>;
    /**
     * Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
     */
    ignoreSslServerVerify?: pulumi.Input<boolean>;
    /**
     * The load-balancer ID this backend is attached to.
     * > **Important:** Updates to `lbId` will recreate the backend.
     */
    lbId?: pulumi.Input<string>;
    /**
     * Maximum number of connections allowed per backend server.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Number of retries when a backend server connection failed.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * The name of the load-balancer backend.
     */
    name?: pulumi.Input<string>;
    /**
     * Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
     */
    onMarkedDownAction?: pulumi.Input<string>;
    /**
     * Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
     */
    proxyProtocol?: pulumi.Input<string>;
    /**
     * Whether to use another backend server on each attempt.
     */
    redispatchAttemptCount?: pulumi.Input<number>;
    /**
     * DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
     *
     * @deprecated Please use proxy_protocol instead
     */
    sendProxyV2?: pulumi.Input<boolean>;
    /**
     * List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
     */
    serverIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables SSL between load balancer and backend servers.
     */
    sslBridging?: pulumi.Input<boolean>;
    /**
     * The type of sticky sessions. The only current possible values are: `none`, `cookie` and `table`.
     */
    stickySessions?: pulumi.Input<string>;
    /**
     * Cookie name for sticky sessions. Only applicable when stickySessions is set to `cookie`.
     */
    stickySessionsCookieName?: pulumi.Input<string>;
    /**
     * Maximum initial server connection establishment time. (e.g.: `1s`)
     */
    timeoutConnect?: pulumi.Input<string>;
    /**
     * Maximum time for a request to be left pending in queue when `maxConnections` is reached. (e.g.: `1s`)
     */
    timeoutQueue?: pulumi.Input<string>;
    /**
     * Maximum server connection inactivity time. (e.g.: `1s`)
     */
    timeoutServer?: pulumi.Input<string>;
    /**
     * Maximum tunnel inactivity time. (e.g.: `1s`)
     */
    timeoutTunnel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbBackend resource.
 */
export interface LbBackendArgs {
    /**
     * Scaleway S3 bucket website to be served in case all backend servers are down.
     * > **Note:** Only the host part of the Scaleway S3 bucket website is expected:
     * e.g. 'failover-website.s3-website.fr-par.scw.cloud' if your bucket website URL is 'https://failover-website.s3-website.fr-par.scw.cloud/'.
     */
    failoverHost?: pulumi.Input<string>;
    /**
     * User sessions will be forwarded to this port of backend servers.
     */
    forwardPort: pulumi.Input<number>;
    /**
     * Load balancing algorithm. Possible values are: `roundrobin`, `leastconn` and `first`.
     */
    forwardPortAlgorithm?: pulumi.Input<string>;
    /**
     * Backend protocol. Possible values are: `tcp` or `http`.
     */
    forwardProtocol: pulumi.Input<string>;
    /**
     * Interval between two HC requests.
     */
    healthCheckDelay?: pulumi.Input<string>;
    /**
     * This block enable HTTP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckHttp?: pulumi.Input<inputs.LbBackendHealthCheckHttp>;
    /**
     * This block enable HTTPS health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckHttps?: pulumi.Input<inputs.LbBackendHealthCheckHttps>;
    /**
     * Number of allowed failed HC requests before the backend server is marked down.
     */
    healthCheckMaxRetries?: pulumi.Input<number>;
    /**
     * Port the HC requests will be send to.
     */
    healthCheckPort?: pulumi.Input<number>;
    /**
     * This block enable TCP health check. Only one of `healthCheckTcp`, `healthCheckHttp` and `healthCheckHttps` should be specified.
     */
    healthCheckTcp?: pulumi.Input<inputs.LbBackendHealthCheckTcp>;
    /**
     * Timeout before we consider a HC request failed.
     */
    healthCheckTimeout?: pulumi.Input<string>;
    /**
     * Specifies whether the Load Balancer should check the backend server’s certificate before initiating a connection.
     */
    ignoreSslServerVerify?: pulumi.Input<boolean>;
    /**
     * The load-balancer ID this backend is attached to.
     * > **Important:** Updates to `lbId` will recreate the backend.
     */
    lbId: pulumi.Input<string>;
    /**
     * Maximum number of connections allowed per backend server.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Number of retries when a backend server connection failed.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * The name of the load-balancer backend.
     */
    name?: pulumi.Input<string>;
    /**
     * Modify what occurs when a backend server is marked down. Possible values are: `none` and `shutdownSessions`.
     */
    onMarkedDownAction?: pulumi.Input<string>;
    /**
     * Choose the type of PROXY protocol to enable (`none`, `v1`, `v2`, `v2Ssl`, `v2SslCn`)
     */
    proxyProtocol?: pulumi.Input<string>;
    /**
     * Whether to use another backend server on each attempt.
     */
    redispatchAttemptCount?: pulumi.Input<number>;
    /**
     * DEPRECATED please use `proxyProtocol` instead - (Default: `false`) Enables PROXY protocol version 2.
     *
     * @deprecated Please use proxy_protocol instead
     */
    sendProxyV2?: pulumi.Input<boolean>;
    /**
     * List of backend server IP addresses. Addresses can be either IPv4 or IPv6.
     */
    serverIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables SSL between load balancer and backend servers.
     */
    sslBridging?: pulumi.Input<boolean>;
    /**
     * The type of sticky sessions. The only current possible values are: `none`, `cookie` and `table`.
     */
    stickySessions?: pulumi.Input<string>;
    /**
     * Cookie name for sticky sessions. Only applicable when stickySessions is set to `cookie`.
     */
    stickySessionsCookieName?: pulumi.Input<string>;
    /**
     * Maximum initial server connection establishment time. (e.g.: `1s`)
     */
    timeoutConnect?: pulumi.Input<string>;
    /**
     * Maximum time for a request to be left pending in queue when `maxConnections` is reached. (e.g.: `1s`)
     */
    timeoutQueue?: pulumi.Input<string>;
    /**
     * Maximum server connection inactivity time. (e.g.: `1s`)
     */
    timeoutServer?: pulumi.Input<string>;
    /**
     * Maximum tunnel inactivity time. (e.g.: `1s`)
     */
    timeoutTunnel?: pulumi.Input<string>;
}
