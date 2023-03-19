// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class K8sClusterKubeconfig {
    /**
     * @return The CA certificate of the Kubernetes API server.
     * 
     */
    private @Nullable String clusterCaCertificate;
    /**
     * @return The raw kubeconfig file.
     * 
     */
    private @Nullable String configFile;
    /**
     * @return The URL of the Kubernetes API server.
     * 
     */
    private @Nullable String host;
    /**
     * @return The token to connect to the Kubernetes API server.
     * 
     */
    private @Nullable String token;

    private K8sClusterKubeconfig() {}
    /**
     * @return The CA certificate of the Kubernetes API server.
     * 
     */
    public Optional<String> clusterCaCertificate() {
        return Optional.ofNullable(this.clusterCaCertificate);
    }
    /**
     * @return The raw kubeconfig file.
     * 
     */
    public Optional<String> configFile() {
        return Optional.ofNullable(this.configFile);
    }
    /**
     * @return The URL of the Kubernetes API server.
     * 
     */
    public Optional<String> host() {
        return Optional.ofNullable(this.host);
    }
    /**
     * @return The token to connect to the Kubernetes API server.
     * 
     */
    public Optional<String> token() {
        return Optional.ofNullable(this.token);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(K8sClusterKubeconfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clusterCaCertificate;
        private @Nullable String configFile;
        private @Nullable String host;
        private @Nullable String token;
        public Builder() {}
        public Builder(K8sClusterKubeconfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterCaCertificate = defaults.clusterCaCertificate;
    	      this.configFile = defaults.configFile;
    	      this.host = defaults.host;
    	      this.token = defaults.token;
        }

        @CustomType.Setter
        public Builder clusterCaCertificate(@Nullable String clusterCaCertificate) {
            this.clusterCaCertificate = clusterCaCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder configFile(@Nullable String configFile) {
            this.configFile = configFile;
            return this;
        }
        @CustomType.Setter
        public Builder host(@Nullable String host) {
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder token(@Nullable String token) {
            this.token = token;
            return this;
        }
        public K8sClusterKubeconfig build() {
            final var o = new K8sClusterKubeconfig();
            o.clusterCaCertificate = clusterCaCertificate;
            o.configFile = configFile;
            o.host = host;
            o.token = token;
            return o;
        }
    }
}
