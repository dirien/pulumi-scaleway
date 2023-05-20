// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.K8sPoolArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.K8sPoolState;
import io.dirien.scaleway.outputs.K8sPoolNode;
import io.dirien.scaleway.outputs.K8sPoolUpgradePolicy;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Kubernetes pools can be imported using the `{region}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/k8sPool:K8sPool mypool fr-par/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/k8sPool:K8sPool")
public class K8sPool extends com.pulumi.resources.CustomResource {
    /**
     * Enables the autohealing feature for this pool.
     * 
     */
    @Export(name="autohealing", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autohealing;

    /**
     * @return Enables the autohealing feature for this pool.
     * 
     */
    public Output<Optional<Boolean>> autohealing() {
        return Codegen.optional(this.autohealing);
    }
    /**
     * Enables the autoscaling feature for this pool.
     * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
     * 
     */
    @Export(name="autoscaling", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoscaling;

    /**
     * @return Enables the autoscaling feature for this pool.
     * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
     * 
     */
    public Output<Optional<Boolean>> autoscaling() {
        return Codegen.optional(this.autoscaling);
    }
    /**
     * The ID of the Kubernetes cluster on which this pool will be created.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of the Kubernetes cluster on which this pool will be created.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The container runtime of the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="containerRuntime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> containerRuntime;

    /**
     * @return The container runtime of the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<Optional<String>> containerRuntime() {
        return Codegen.optional(this.containerRuntime);
    }
    /**
     * The creation date of the pool.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The creation date of the pool.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The actual size of the pool
     * 
     */
    @Export(name="currentSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> currentSize;

    /**
     * @return The actual size of the pool
     * 
     */
    public Output<Integer> currentSize() {
        return this.currentSize;
    }
    /**
     * The Kubelet arguments to be used by this pool
     * 
     */
    @Export(name="kubeletArgs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> kubeletArgs;

    /**
     * @return The Kubelet arguments to be used by this pool
     * 
     */
    public Output<Optional<Map<String,String>>> kubeletArgs() {
        return Codegen.optional(this.kubeletArgs);
    }
    /**
     * The maximum size of the pool, used by the autoscaling feature.
     * 
     */
    @Export(name="maxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxSize;

    /**
     * @return The maximum size of the pool, used by the autoscaling feature.
     * 
     */
    public Output<Integer> maxSize() {
        return this.maxSize;
    }
    /**
     * The minimum size of the pool, used by the autoscaling feature.
     * 
     */
    @Export(name="minSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minSize;

    /**
     * @return The minimum size of the pool, used by the autoscaling feature.
     * 
     */
    public Output<Optional<Integer>> minSize() {
        return Codegen.optional(this.minSize);
    }
    /**
     * The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
     * 
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="nodeType", refs={String.class}, tree="[0]")
    private Output<String> nodeType;

    /**
     * @return The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
     * 
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<String> nodeType() {
        return this.nodeType;
    }
    /**
     * (List of) The nodes in the default pool.
     * 
     */
    @Export(name="nodes", refs={List.class,K8sPoolNode.class}, tree="[0,1]")
    private Output<List<K8sPoolNode>> nodes;

    /**
     * @return (List of) The nodes in the default pool.
     * 
     */
    public Output<List<K8sPoolNode>> nodes() {
        return this.nodes;
    }
    /**
     * The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="placementGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> placementGroupId;

    /**
     * @return The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<Optional<String>> placementGroupId() {
        return Codegen.optional(this.placementGroupId);
    }
    /**
     * `region`) The region in which the pool should be created.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return `region`) The region in which the pool should be created.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The size of the system volume of the nodes in gigabyte
     * 
     */
    @Export(name="rootVolumeSizeInGb", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rootVolumeSizeInGb;

    /**
     * @return The size of the system volume of the nodes in gigabyte
     * 
     */
    public Output<Optional<Integer>> rootVolumeSizeInGb() {
        return Codegen.optional(this.rootVolumeSizeInGb);
    }
    /**
     * System volume type of the nodes composing the pool
     * 
     */
    @Export(name="rootVolumeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rootVolumeType;

    /**
     * @return System volume type of the nodes composing the pool
     * 
     */
    public Output<Optional<String>> rootVolumeType() {
        return Codegen.optional(this.rootVolumeType);
    }
    /**
     * The size of the pool.
     * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The size of the pool.
     * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The status of the node.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the node.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags associated with the pool.
     * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags associated with the pool.
     * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The last update date of the pool.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The last update date of the pool.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * The Pool upgrade policy
     * 
     */
    @Export(name="upgradePolicy", refs={K8sPoolUpgradePolicy.class}, tree="[0]")
    private Output<K8sPoolUpgradePolicy> upgradePolicy;

    /**
     * @return The Pool upgrade policy
     * 
     */
    public Output<K8sPoolUpgradePolicy> upgradePolicy() {
        return this.upgradePolicy;
    }
    /**
     * The version of the pool.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version of the pool.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * Whether to wait for the pool to be ready.
     * 
     */
    @Export(name="waitForPoolReady", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitForPoolReady;

    /**
     * @return Whether to wait for the pool to be ready.
     * 
     */
    public Output<Optional<Boolean>> waitForPoolReady() {
        return Codegen.optional(this.waitForPoolReady);
    }
    /**
     * `zone`) The zone in which the pool should be created.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the pool should be created.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public K8sPool(String name) {
        this(name, K8sPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public K8sPool(String name, K8sPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public K8sPool(String name, K8sPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/k8sPool:K8sPool", name, args == null ? K8sPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private K8sPool(String name, Output<String> id, @Nullable K8sPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/k8sPool:K8sPool", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static K8sPool get(String name, Output<String> id, @Nullable K8sPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new K8sPool(name, id, state, options);
    }
}
