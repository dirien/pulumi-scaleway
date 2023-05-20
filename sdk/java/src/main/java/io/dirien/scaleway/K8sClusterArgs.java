// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.K8sClusterAutoUpgradeArgs;
import io.dirien.scaleway.inputs.K8sClusterAutoscalerConfigArgs;
import io.dirien.scaleway.inputs.K8sClusterOpenIdConnectConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class K8sClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final K8sClusterArgs Empty = new K8sClusterArgs();

    /**
     * The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
     * 
     */
    @Import(name="admissionPlugins")
    private @Nullable Output<List<String>> admissionPlugins;

    /**
     * @return The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
     * 
     */
    public Optional<Output<List<String>>> admissionPlugins() {
        return Optional.ofNullable(this.admissionPlugins);
    }

    /**
     * Additional Subject Alternative Names for the Kubernetes API server certificate
     * 
     */
    @Import(name="apiserverCertSans")
    private @Nullable Output<List<String>> apiserverCertSans;

    /**
     * @return Additional Subject Alternative Names for the Kubernetes API server certificate
     * 
     */
    public Optional<Output<List<String>>> apiserverCertSans() {
        return Optional.ofNullable(this.apiserverCertSans);
    }

    /**
     * The auto upgrade configuration.
     * 
     */
    @Import(name="autoUpgrade")
    private @Nullable Output<K8sClusterAutoUpgradeArgs> autoUpgrade;

    /**
     * @return The auto upgrade configuration.
     * 
     */
    public Optional<Output<K8sClusterAutoUpgradeArgs>> autoUpgrade() {
        return Optional.ofNullable(this.autoUpgrade);
    }

    /**
     * The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
     * 
     */
    @Import(name="autoscalerConfig")
    private @Nullable Output<K8sClusterAutoscalerConfigArgs> autoscalerConfig;

    /**
     * @return The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
     * 
     */
    public Optional<Output<K8sClusterAutoscalerConfigArgs>> autoscalerConfig() {
        return Optional.ofNullable(this.autoscalerConfig);
    }

    /**
     * The Container Network Interface (CNI) for the Kubernetes cluster.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="cni", required=true)
    private Output<String> cni;

    /**
     * @return The Container Network Interface (CNI) for the Kubernetes cluster.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Output<String> cni() {
        return this.cni;
    }

    /**
     * Delete additional resources like block volumes, IPs and loadbalancers that were created in Kubernetes on cluster deletion.
     * &gt; **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
     * If you prefer keeping it, you should instead set it as `false`.
     * 
     */
    @Import(name="deleteAdditionalResources", required=true)
    private Output<Boolean> deleteAdditionalResources;

    /**
     * @return Delete additional resources like block volumes, IPs and loadbalancers that were created in Kubernetes on cluster deletion.
     * &gt; **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
     * If you prefer keeping it, you should instead set it as `false`.
     * 
     */
    public Output<Boolean> deleteAdditionalResources() {
        return this.deleteAdditionalResources;
    }

    /**
     * A description for the Kubernetes cluster.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the Kubernetes cluster.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
     * 
     */
    @Import(name="featureGates")
    private @Nullable Output<List<String>> featureGates;

    /**
     * @return The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
     * 
     */
    public Optional<Output<List<String>>> featureGates() {
        return Optional.ofNullable(this.featureGates);
    }

    /**
     * The name for the Kubernetes cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name for the Kubernetes cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The OpenID Connect configuration of the cluster
     * 
     */
    @Import(name="openIdConnectConfig")
    private @Nullable Output<K8sClusterOpenIdConnectConfigArgs> openIdConnectConfig;

    /**
     * @return The OpenID Connect configuration of the cluster
     * 
     */
    public Optional<Output<K8sClusterOpenIdConnectConfigArgs>> openIdConnectConfig() {
        return Optional.ofNullable(this.openIdConnectConfig);
    }

    /**
     * The ID of the private network of the cluster.
     * 
     * &gt; **Important:** This field can only be set at cluster creation and cannot be updated later.
     * Changes to this field will cause the cluster to be destroyed then recreated.
     * 
     */
    @Import(name="privateNetworkId")
    private @Nullable Output<String> privateNetworkId;

    /**
     * @return The ID of the private network of the cluster.
     * 
     * &gt; **Important:** This field can only be set at cluster creation and cannot be updated later.
     * Changes to this field will cause the cluster to be destroyed then recreated.
     * 
     */
    public Optional<Output<String>> privateNetworkId() {
        return Optional.ofNullable(this.privateNetworkId);
    }

    /**
     * `project_id`) The ID of the project the cluster is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the cluster is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `region`) The region in which the cluster should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the cluster should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The tags associated with the Kubernetes cluster.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the Kubernetes cluster.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The version of the Kubernetes cluster.
     * 
     */
    @Import(name="version", required=true)
    private Output<String> version;

    /**
     * @return The version of the Kubernetes cluster.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    private K8sClusterArgs() {}

    private K8sClusterArgs(K8sClusterArgs $) {
        this.admissionPlugins = $.admissionPlugins;
        this.apiserverCertSans = $.apiserverCertSans;
        this.autoUpgrade = $.autoUpgrade;
        this.autoscalerConfig = $.autoscalerConfig;
        this.cni = $.cni;
        this.deleteAdditionalResources = $.deleteAdditionalResources;
        this.description = $.description;
        this.featureGates = $.featureGates;
        this.name = $.name;
        this.openIdConnectConfig = $.openIdConnectConfig;
        this.privateNetworkId = $.privateNetworkId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.tags = $.tags;
        this.type = $.type;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(K8sClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private K8sClusterArgs $;

        public Builder() {
            $ = new K8sClusterArgs();
        }

        public Builder(K8sClusterArgs defaults) {
            $ = new K8sClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param admissionPlugins The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder admissionPlugins(@Nullable Output<List<String>> admissionPlugins) {
            $.admissionPlugins = admissionPlugins;
            return this;
        }

        /**
         * @param admissionPlugins The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder admissionPlugins(List<String> admissionPlugins) {
            return admissionPlugins(Output.of(admissionPlugins));
        }

        /**
         * @param admissionPlugins The list of [admission plugins](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder admissionPlugins(String... admissionPlugins) {
            return admissionPlugins(List.of(admissionPlugins));
        }

        /**
         * @param apiserverCertSans Additional Subject Alternative Names for the Kubernetes API server certificate
         * 
         * @return builder
         * 
         */
        public Builder apiserverCertSans(@Nullable Output<List<String>> apiserverCertSans) {
            $.apiserverCertSans = apiserverCertSans;
            return this;
        }

        /**
         * @param apiserverCertSans Additional Subject Alternative Names for the Kubernetes API server certificate
         * 
         * @return builder
         * 
         */
        public Builder apiserverCertSans(List<String> apiserverCertSans) {
            return apiserverCertSans(Output.of(apiserverCertSans));
        }

        /**
         * @param apiserverCertSans Additional Subject Alternative Names for the Kubernetes API server certificate
         * 
         * @return builder
         * 
         */
        public Builder apiserverCertSans(String... apiserverCertSans) {
            return apiserverCertSans(List.of(apiserverCertSans));
        }

        /**
         * @param autoUpgrade The auto upgrade configuration.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgrade(@Nullable Output<K8sClusterAutoUpgradeArgs> autoUpgrade) {
            $.autoUpgrade = autoUpgrade;
            return this;
        }

        /**
         * @param autoUpgrade The auto upgrade configuration.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgrade(K8sClusterAutoUpgradeArgs autoUpgrade) {
            return autoUpgrade(Output.of(autoUpgrade));
        }

        /**
         * @param autoscalerConfig The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
         * 
         * @return builder
         * 
         */
        public Builder autoscalerConfig(@Nullable Output<K8sClusterAutoscalerConfigArgs> autoscalerConfig) {
            $.autoscalerConfig = autoscalerConfig;
            return this;
        }

        /**
         * @param autoscalerConfig The configuration options for the [Kubernetes cluster autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler).
         * 
         * @return builder
         * 
         */
        public Builder autoscalerConfig(K8sClusterAutoscalerConfigArgs autoscalerConfig) {
            return autoscalerConfig(Output.of(autoscalerConfig));
        }

        /**
         * @param cni The Container Network Interface (CNI) for the Kubernetes cluster.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder cni(Output<String> cni) {
            $.cni = cni;
            return this;
        }

        /**
         * @param cni The Container Network Interface (CNI) for the Kubernetes cluster.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder cni(String cni) {
            return cni(Output.of(cni));
        }

        /**
         * @param deleteAdditionalResources Delete additional resources like block volumes, IPs and loadbalancers that were created in Kubernetes on cluster deletion.
         * &gt; **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
         * If you prefer keeping it, you should instead set it as `false`.
         * 
         * @return builder
         * 
         */
        public Builder deleteAdditionalResources(Output<Boolean> deleteAdditionalResources) {
            $.deleteAdditionalResources = deleteAdditionalResources;
            return this;
        }

        /**
         * @param deleteAdditionalResources Delete additional resources like block volumes, IPs and loadbalancers that were created in Kubernetes on cluster deletion.
         * &gt; **Important:** Setting this field to `true` means that you will lose all your cluster data and network configuration when you delete your cluster.
         * If you prefer keeping it, you should instead set it as `false`.
         * 
         * @return builder
         * 
         */
        public Builder deleteAdditionalResources(Boolean deleteAdditionalResources) {
            return deleteAdditionalResources(Output.of(deleteAdditionalResources));
        }

        /**
         * @param description A description for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param featureGates The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder featureGates(@Nullable Output<List<String>> featureGates) {
            $.featureGates = featureGates;
            return this;
        }

        /**
         * @param featureGates The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder featureGates(List<String> featureGates) {
            return featureGates(Output.of(featureGates));
        }

        /**
         * @param featureGates The list of [feature gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) to enable on the cluster.
         * 
         * @return builder
         * 
         */
        public Builder featureGates(String... featureGates) {
            return featureGates(List.of(featureGates));
        }

        /**
         * @param name The name for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param openIdConnectConfig The OpenID Connect configuration of the cluster
         * 
         * @return builder
         * 
         */
        public Builder openIdConnectConfig(@Nullable Output<K8sClusterOpenIdConnectConfigArgs> openIdConnectConfig) {
            $.openIdConnectConfig = openIdConnectConfig;
            return this;
        }

        /**
         * @param openIdConnectConfig The OpenID Connect configuration of the cluster
         * 
         * @return builder
         * 
         */
        public Builder openIdConnectConfig(K8sClusterOpenIdConnectConfigArgs openIdConnectConfig) {
            return openIdConnectConfig(Output.of(openIdConnectConfig));
        }

        /**
         * @param privateNetworkId The ID of the private network of the cluster.
         * 
         * &gt; **Important:** This field can only be set at cluster creation and cannot be updated later.
         * Changes to this field will cause the cluster to be destroyed then recreated.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(@Nullable Output<String> privateNetworkId) {
            $.privateNetworkId = privateNetworkId;
            return this;
        }

        /**
         * @param privateNetworkId The ID of the private network of the cluster.
         * 
         * &gt; **Important:** This field can only be set at cluster creation and cannot be updated later.
         * Changes to this field will cause the cluster to be destroyed then recreated.
         * 
         * @return builder
         * 
         */
        public Builder privateNetworkId(String privateNetworkId) {
            return privateNetworkId(Output.of(privateNetworkId));
        }

        /**
         * @param projectId `project_id`) The ID of the project the cluster is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the cluster is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region `region`) The region in which the cluster should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the cluster should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tags The tags associated with the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of Kubernetes cluster. Possible values are: `kapsule` or `multicloud`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param version The version of the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder version(Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public K8sClusterArgs build() {
            $.cni = Objects.requireNonNull($.cni, "expected parameter 'cni' to be non-null");
            $.deleteAdditionalResources = Objects.requireNonNull($.deleteAdditionalResources, "expected parameter 'deleteAdditionalResources' to be non-null");
            $.version = Objects.requireNonNull($.version, "expected parameter 'version' to be non-null");
            return $;
        }
    }

}
