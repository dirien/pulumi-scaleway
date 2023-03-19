// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class K8sClusterKubeconfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final K8sClusterKubeconfigArgs Empty = new K8sClusterKubeconfigArgs();

    /**
     * The CA certificate of the Kubernetes API server.
     * 
     */
    @Import(name="clusterCaCertificate")
    private @Nullable Output<String> clusterCaCertificate;

    /**
     * @return The CA certificate of the Kubernetes API server.
     * 
     */
    public Optional<Output<String>> clusterCaCertificate() {
        return Optional.ofNullable(this.clusterCaCertificate);
    }

    /**
     * The raw kubeconfig file.
     * 
     */
    @Import(name="configFile")
    private @Nullable Output<String> configFile;

    /**
     * @return The raw kubeconfig file.
     * 
     */
    public Optional<Output<String>> configFile() {
        return Optional.ofNullable(this.configFile);
    }

    /**
     * The URL of the Kubernetes API server.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The URL of the Kubernetes API server.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * The token to connect to the Kubernetes API server.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The token to connect to the Kubernetes API server.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    private K8sClusterKubeconfigArgs() {}

    private K8sClusterKubeconfigArgs(K8sClusterKubeconfigArgs $) {
        this.clusterCaCertificate = $.clusterCaCertificate;
        this.configFile = $.configFile;
        this.host = $.host;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(K8sClusterKubeconfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private K8sClusterKubeconfigArgs $;

        public Builder() {
            $ = new K8sClusterKubeconfigArgs();
        }

        public Builder(K8sClusterKubeconfigArgs defaults) {
            $ = new K8sClusterKubeconfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterCaCertificate The CA certificate of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder clusterCaCertificate(@Nullable Output<String> clusterCaCertificate) {
            $.clusterCaCertificate = clusterCaCertificate;
            return this;
        }

        /**
         * @param clusterCaCertificate The CA certificate of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder clusterCaCertificate(String clusterCaCertificate) {
            return clusterCaCertificate(Output.of(clusterCaCertificate));
        }

        /**
         * @param configFile The raw kubeconfig file.
         * 
         * @return builder
         * 
         */
        public Builder configFile(@Nullable Output<String> configFile) {
            $.configFile = configFile;
            return this;
        }

        /**
         * @param configFile The raw kubeconfig file.
         * 
         * @return builder
         * 
         */
        public Builder configFile(String configFile) {
            return configFile(Output.of(configFile));
        }

        /**
         * @param host The URL of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The URL of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param token The token to connect to the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The token to connect to the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public K8sClusterKubeconfigArgs build() {
            return $;
        }
    }

}
