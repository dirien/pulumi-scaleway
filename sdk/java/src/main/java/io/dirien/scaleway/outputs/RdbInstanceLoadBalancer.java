// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RdbInstanceLoadBalancer {
    /**
     * @return The ID of the endpoint of the private network.
     * 
     */
    private @Nullable String endpointId;
    /**
     * @return Name of the endpoint.
     * 
     */
    private @Nullable String hostname;
    /**
     * @return IP of the endpoint.
     * 
     */
    private @Nullable String ip;
    /**
     * @return The name of the Database Instance.
     * 
     */
    private @Nullable String name;
    /**
     * @return Port of the endpoint.
     * 
     */
    private @Nullable Integer port;

    private RdbInstanceLoadBalancer() {}
    /**
     * @return The ID of the endpoint of the private network.
     * 
     */
    public Optional<String> endpointId() {
        return Optional.ofNullable(this.endpointId);
    }
    /**
     * @return Name of the endpoint.
     * 
     */
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }
    /**
     * @return IP of the endpoint.
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return The name of the Database Instance.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Port of the endpoint.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RdbInstanceLoadBalancer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String endpointId;
        private @Nullable String hostname;
        private @Nullable String ip;
        private @Nullable String name;
        private @Nullable Integer port;
        public Builder() {}
        public Builder(RdbInstanceLoadBalancer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointId = defaults.endpointId;
    	      this.hostname = defaults.hostname;
    	      this.ip = defaults.ip;
    	      this.name = defaults.name;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder endpointId(@Nullable String endpointId) {
            this.endpointId = endpointId;
            return this;
        }
        @CustomType.Setter
        public Builder hostname(@Nullable String hostname) {
            this.hostname = hostname;
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        public RdbInstanceLoadBalancer build() {
            final var o = new RdbInstanceLoadBalancer();
            o.endpointId = endpointId;
            o.hostname = hostname;
            o.ip = ip;
            o.name = name;
            o.port = port;
            return o;
        }
    }
}
