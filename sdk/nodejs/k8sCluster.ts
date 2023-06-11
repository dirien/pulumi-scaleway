// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Kubernetes clusters. For more information, see [the documentation](https://developers.scaleway.com/en/products/k8s/api/).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const jack = new scaleway.K8sCluster("jack", {
 *     version: "1.24.3",
 *     cni: "cilium",
 *     deleteAdditionalResources: false,
 * });
 * const john = new scaleway.K8sPool("john", {
 *     clusterId: jack.id,
 *     nodeType: "DEV1-M",
 *     size: 1,
 * });
 * ```
 *
 * ### Multicloud
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const henry = new scaleway.K8sCluster("henry", {
 *     type: "multicloud",
 *     version: "1.24.3",
 *     cni: "kilo",
 *     deleteAdditionalResources: false,
 * });
 * const friendFromOuterSpace = new scaleway.K8sPool("friendFromOuterSpace", {
 *     clusterId: henry.id,
 *     nodeType: "external",
 *     size: 0,
 *     minSize: 0,
 * });
 * ```
 *
 * For a detailed example of how to add or run Elastic Metal servers instead of instances on your cluster, please refer to this guide.
 *
 * ### With additional configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const johnK8sCluster = new scaleway.K8sCluster("johnK8sCluster", {
 *     description: "my awesome cluster",
 *     version: "1.24.3",
 *     cni: "calico",
 *     tags: [
 *         "i'm an awesome tag",
 *         "yay",
 *     ],
 *     deleteAdditionalResources: false,
 *     autoscalerConfig: {
 *         disableScaleDown: false,
 *         scaleDownDelayAfterAdd: "5m",
 *         estimator: "binpacking",
 *         expander: "random",
 *         ignoreDaemonsetsUtilization: true,
 *         balanceSimilarNodeGroups: true,
 *         expendablePodsPriorityCutoff: -5,
 *     },
 * });
 * const johnK8sPool = new scaleway.K8sPool("johnK8sPool", {
 *     clusterId: johnK8sCluster.id,
 *     nodeType: "DEV1-M",
 *     size: 3,
 *     autoscaling: true,
 *     autohealing: true,
 *     minSize: 1,
 *     maxSize: 5,
 * });
 * ```
 *
 * ## Import
 *
 * Kubernetes clusters can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/k8sCluster:K8sCluster mycluster fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class K8sCluster extends pulumi.CustomResource {
    /**
     * Get an existing K8sCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: K8sClusterState, opts?: pulumi.CustomResourceOptions): K8sCluster {
        return new K8sCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/k8sCluster:K8sCluster';

    /**
     * Returns true if the given object is an instance of K8sCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is K8sCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === K8sCluster.__pulumiType;
    }

    /**
     * The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
     */
    public readonly admissionPlugins!: pulumi.Output<string[] | undefined>;
    /**
     * Additional Subject Alternative Names for the Kubernetes API server certificate
     */
    public readonly apiserverCertSans!: pulumi.Output<string[] | undefined>;
    /**
     * The URL of the Kubernetes API server.
     */
    public /*out*/ readonly apiserverUrl!: pulumi.Output<string>;
    /**
     * The auto upgrade configuration.
     */
    public readonly autoUpgrade!: pulumi.Output<outputs.K8sClusterAutoUpgrade>;
    /**
     * The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
     */
    public readonly autoscalerConfig!: pulumi.Output<outputs.K8sClusterAutoscalerConfig>;
    /**
     * The Container Network Interface (CNI) for the Kubernetes cluster.
     * > **Important:** Updates to this field will recreate a new resource.
     */
    public readonly cni!: pulumi.Output<string>;
    /**
     * The creation date of the cluster.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
     * > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
     * If you prefer keeping it, you should instead set it as `false`.
     */
    public readonly deleteAdditionalResources!: pulumi.Output<boolean>;
    /**
     * A description for the Kubernetes cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
     */
    public readonly featureGates!: pulumi.Output<string[] | undefined>;
    /**
     * The kubeconfig configuration file of the Kubernetes cluster
     */
    public /*out*/ readonly kubeconfigs!: pulumi.Output<outputs.K8sClusterKubeconfig[]>;
    /**
     * The name for the Kubernetes cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The OpenID Connect configuration of the cluster
     */
    public readonly openIdConnectConfig!: pulumi.Output<outputs.K8sClusterOpenIdConnectConfig>;
    /**
     * The organization ID the cluster is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The ID of the private network of the cluster.
     *
     * > **Important:** This field can only be set at cluster creation and cannot be updated later.
     * Changes to this field will cause the cluster to be destroyed then recreated.
     */
    public readonly privateNetworkId!: pulumi.Output<string | undefined>;
    /**
     * `projectId`) The ID of the project the cluster is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region in which the cluster should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of the Kubernetes cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags associated with the Kubernetes cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The last update date of the cluster.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Set to `true` if a newer Kubernetes version is available.
     */
    public /*out*/ readonly upgradeAvailable!: pulumi.Output<boolean>;
    /**
     * The version of the Kubernetes cluster.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The DNS wildcard that points to all ready nodes.
     */
    public /*out*/ readonly wildcardDns!: pulumi.Output<string>;

    /**
     * Create a K8sCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: K8sClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: K8sClusterArgs | K8sClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as K8sClusterState | undefined;
            resourceInputs["admissionPlugins"] = state ? state.admissionPlugins : undefined;
            resourceInputs["apiserverCertSans"] = state ? state.apiserverCertSans : undefined;
            resourceInputs["apiserverUrl"] = state ? state.apiserverUrl : undefined;
            resourceInputs["autoUpgrade"] = state ? state.autoUpgrade : undefined;
            resourceInputs["autoscalerConfig"] = state ? state.autoscalerConfig : undefined;
            resourceInputs["cni"] = state ? state.cni : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deleteAdditionalResources"] = state ? state.deleteAdditionalResources : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["featureGates"] = state ? state.featureGates : undefined;
            resourceInputs["kubeconfigs"] = state ? state.kubeconfigs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openIdConnectConfig"] = state ? state.openIdConnectConfig : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["privateNetworkId"] = state ? state.privateNetworkId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["upgradeAvailable"] = state ? state.upgradeAvailable : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["wildcardDns"] = state ? state.wildcardDns : undefined;
        } else {
            const args = argsOrState as K8sClusterArgs | undefined;
            if ((!args || args.cni === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cni'");
            }
            if ((!args || args.deleteAdditionalResources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deleteAdditionalResources'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["admissionPlugins"] = args ? args.admissionPlugins : undefined;
            resourceInputs["apiserverCertSans"] = args ? args.apiserverCertSans : undefined;
            resourceInputs["autoUpgrade"] = args ? args.autoUpgrade : undefined;
            resourceInputs["autoscalerConfig"] = args ? args.autoscalerConfig : undefined;
            resourceInputs["cni"] = args ? args.cni : undefined;
            resourceInputs["deleteAdditionalResources"] = args ? args.deleteAdditionalResources : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["featureGates"] = args ? args.featureGates : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openIdConnectConfig"] = args ? args.openIdConnectConfig : undefined;
            resourceInputs["privateNetworkId"] = args ? args.privateNetworkId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["apiserverUrl"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["kubeconfigs"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["upgradeAvailable"] = undefined /*out*/;
            resourceInputs["wildcardDns"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubeconfigs"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(K8sCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering K8sCluster resources.
 */
export interface K8sClusterState {
    /**
     * The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
     */
    admissionPlugins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Additional Subject Alternative Names for the Kubernetes API server certificate
     */
    apiserverCertSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the Kubernetes API server.
     */
    apiserverUrl?: pulumi.Input<string>;
    /**
     * The auto upgrade configuration.
     */
    autoUpgrade?: pulumi.Input<inputs.K8sClusterAutoUpgrade>;
    /**
     * The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
     */
    autoscalerConfig?: pulumi.Input<inputs.K8sClusterAutoscalerConfig>;
    /**
     * The Container Network Interface (CNI) for the Kubernetes cluster.
     * > **Important:** Updates to this field will recreate a new resource.
     */
    cni?: pulumi.Input<string>;
    /**
     * The creation date of the cluster.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
     * > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
     * If you prefer keeping it, you should instead set it as `false`.
     */
    deleteAdditionalResources?: pulumi.Input<boolean>;
    /**
     * A description for the Kubernetes cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
     */
    featureGates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The kubeconfig configuration file of the Kubernetes cluster
     */
    kubeconfigs?: pulumi.Input<pulumi.Input<inputs.K8sClusterKubeconfig>[]>;
    /**
     * The name for the Kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The OpenID Connect configuration of the cluster
     */
    openIdConnectConfig?: pulumi.Input<inputs.K8sClusterOpenIdConnectConfig>;
    /**
     * The organization ID the cluster is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The ID of the private network of the cluster.
     *
     * > **Important:** This field can only be set at cluster creation and cannot be updated later.
     * Changes to this field will cause the cluster to be destroyed then recreated.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the cluster is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the cluster should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the Kubernetes cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags associated with the Kubernetes cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
     */
    type?: pulumi.Input<string>;
    /**
     * The last update date of the cluster.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Set to `true` if a newer Kubernetes version is available.
     */
    upgradeAvailable?: pulumi.Input<boolean>;
    /**
     * The version of the Kubernetes cluster.
     */
    version?: pulumi.Input<string>;
    /**
     * The DNS wildcard that points to all ready nodes.
     */
    wildcardDns?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a K8sCluster resource.
 */
export interface K8sClusterArgs {
    /**
     * The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
     */
    admissionPlugins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Additional Subject Alternative Names for the Kubernetes API server certificate
     */
    apiserverCertSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The auto upgrade configuration.
     */
    autoUpgrade?: pulumi.Input<inputs.K8sClusterAutoUpgrade>;
    /**
     * The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
     */
    autoscalerConfig?: pulumi.Input<inputs.K8sClusterAutoscalerConfig>;
    /**
     * The Container Network Interface (CNI) for the Kubernetes cluster.
     * > **Important:** Updates to this field will recreate a new resource.
     */
    cni: pulumi.Input<string>;
    /**
     * Delete additional resources like block volumes, loadbalancers and the cluster private network (if empty) that were created in Kubernetes on cluster deletion.
     * > **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
     * If you prefer keeping it, you should instead set it as `false`.
     */
    deleteAdditionalResources: pulumi.Input<boolean>;
    /**
     * A description for the Kubernetes cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
     */
    featureGates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for the Kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The OpenID Connect configuration of the cluster
     */
    openIdConnectConfig?: pulumi.Input<inputs.K8sClusterOpenIdConnectConfig>;
    /**
     * The ID of the private network of the cluster.
     *
     * > **Important:** This field can only be set at cluster creation and cannot be updated later.
     * Changes to this field will cause the cluster to be destroyed then recreated.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the cluster is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the cluster should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the Kubernetes cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
     */
    type?: pulumi.Input<string>;
    /**
     * The version of the Kubernetes cluster.
     */
    version: pulumi.Input<string>;
}
