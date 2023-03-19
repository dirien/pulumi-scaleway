// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRedisClusterPublicNetwork {
    /**
     * @return The ID of the Redis cluster.
     * 
     */
    private String id;
    private List<String> ips;
    private Integer port;

    private GetRedisClusterPublicNetwork() {}
    /**
     * @return The ID of the Redis cluster.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ips() {
        return this.ips;
    }
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRedisClusterPublicNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ips;
        private Integer port;
        public Builder() {}
        public Builder(GetRedisClusterPublicNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ips = defaults.ips;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ips(List<String> ips) {
            this.ips = Objects.requireNonNull(ips);
            return this;
        }
        public Builder ips(String... ips) {
            return ips(List.of(ips));
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public GetRedisClusterPublicNetwork build() {
            final var o = new GetRedisClusterPublicNetwork();
            o.id = id;
            o.ips = ips;
            o.port = port;
            return o;
        }
    }
}
