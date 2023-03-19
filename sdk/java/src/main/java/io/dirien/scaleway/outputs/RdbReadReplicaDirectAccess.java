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
public final class RdbReadReplicaDirectAccess {
    /**
     * @return The ID of the endpoint of the read replica.
     * 
     */
    private @Nullable String endpointId;
    /**
     * @return Hostname of the endpoint. Only one of ip and hostname may be set.
     * 
     */
    private @Nullable String hostname;
    /**
     * @return IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     * 
     */
    private @Nullable String ip;
    /**
     * @return Name of the endpoint.
     * 
     */
    private @Nullable String name;
    /**
     * @return TCP port of the endpoint.
     * 
     */
    private @Nullable Integer port;

    private RdbReadReplicaDirectAccess() {}
    /**
     * @return The ID of the endpoint of the read replica.
     * 
     */
    public Optional<String> endpointId() {
        return Optional.ofNullable(this.endpointId);
    }
    /**
     * @return Hostname of the endpoint. Only one of ip and hostname may be set.
     * 
     */
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }
    /**
     * @return IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return Name of the endpoint.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return TCP port of the endpoint.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RdbReadReplicaDirectAccess defaults) {
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
        public Builder(RdbReadReplicaDirectAccess defaults) {
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
        public RdbReadReplicaDirectAccess build() {
            final var o = new RdbReadReplicaDirectAccess();
            o.endpointId = endpointId;
            o.hostname = hostname;
            o.ip = ip;
            o.name = name;
            o.port = port;
            return o;
        }
    }
}
