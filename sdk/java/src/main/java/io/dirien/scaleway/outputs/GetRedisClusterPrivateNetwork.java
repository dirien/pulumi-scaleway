// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRedisClusterPrivateNetwork {
    private String endpointId;
    /**
     * @return The ID of the Redis cluster.
     * 
     */
    private String id;
    private List<String> serviceIps;
    /**
     * @return `region`) The zone in which the server exists.
     * 
     */
    private String zone;

    private GetRedisClusterPrivateNetwork() {}
    public String endpointId() {
        return this.endpointId;
    }
    /**
     * @return The ID of the Redis cluster.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> serviceIps() {
        return this.serviceIps;
    }
    /**
     * @return `region`) The zone in which the server exists.
     * 
     */
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRedisClusterPrivateNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String endpointId;
        private String id;
        private List<String> serviceIps;
        private String zone;
        public Builder() {}
        public Builder(GetRedisClusterPrivateNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointId = defaults.endpointId;
    	      this.id = defaults.id;
    	      this.serviceIps = defaults.serviceIps;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder endpointId(String endpointId) {
            this.endpointId = Objects.requireNonNull(endpointId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder serviceIps(List<String> serviceIps) {
            this.serviceIps = Objects.requireNonNull(serviceIps);
            return this;
        }
        public Builder serviceIps(String... serviceIps) {
            return serviceIps(List.of(serviceIps));
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetRedisClusterPrivateNetwork build() {
            final var o = new GetRedisClusterPrivateNetwork();
            o.endpointId = endpointId;
            o.id = id;
            o.serviceIps = serviceIps;
            o.zone = zone;
            return o;
        }
    }
}
