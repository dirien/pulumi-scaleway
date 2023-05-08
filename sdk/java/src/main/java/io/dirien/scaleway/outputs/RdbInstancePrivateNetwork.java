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
public final class RdbInstancePrivateNetwork {
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
    private @Nullable String ipNet;
    /**
     * @return The name of the Database Instance.
     * 
     */
    private @Nullable String name;
    private String pnId;
    /**
     * @return Port of the endpoint.
     * 
     */
    private @Nullable Integer port;
    private @Nullable String zone;

    private RdbInstancePrivateNetwork() {}
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
    public Optional<String> ipNet() {
        return Optional.ofNullable(this.ipNet);
    }
    /**
     * @return The name of the Database Instance.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public String pnId() {
        return this.pnId;
    }
    /**
     * @return Port of the endpoint.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RdbInstancePrivateNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String endpointId;
        private @Nullable String hostname;
        private @Nullable String ip;
        private @Nullable String ipNet;
        private @Nullable String name;
        private String pnId;
        private @Nullable Integer port;
        private @Nullable String zone;
        public Builder() {}
        public Builder(RdbInstancePrivateNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointId = defaults.endpointId;
    	      this.hostname = defaults.hostname;
    	      this.ip = defaults.ip;
    	      this.ipNet = defaults.ipNet;
    	      this.name = defaults.name;
    	      this.pnId = defaults.pnId;
    	      this.port = defaults.port;
    	      this.zone = defaults.zone;
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
        public Builder ipNet(@Nullable String ipNet) {
            this.ipNet = ipNet;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder pnId(String pnId) {
            this.pnId = Objects.requireNonNull(pnId);
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder zone(@Nullable String zone) {
            this.zone = zone;
            return this;
        }
        public RdbInstancePrivateNetwork build() {
            final var o = new RdbInstancePrivateNetwork();
            o.endpointId = endpointId;
            o.hostname = hostname;
            o.ip = ip;
            o.ipNet = ipNet;
            o.name = name;
            o.pnId = pnId;
            o.port = port;
            o.zone = zone;
            return o;
        }
    }
}
