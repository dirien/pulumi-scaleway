// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLbIpsIp {
    /**
     * @return The associated IP ID.
     * 
     */
    private String id;
    /**
     * @return The IP Address
     * 
     */
    private String ipAddress;
    /**
     * @return The associated load-balancer ID if any
     * 
     */
    private String lbId;
    /**
     * @return The organization ID the load-balancer is associated with.
     * 
     */
    private String organizationId;
    /**
     * @return The ID of the project the load-balancer is associated with.
     * 
     */
    private String projectId;
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    private String reverse;
    /**
     * @return `zone`) The zone in which IPs exist.
     * 
     */
    private String zone;

    private GetLbIpsIp() {}
    /**
     * @return The associated IP ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The IP Address
     * 
     */
    public String ipAddress() {
        return this.ipAddress;
    }
    /**
     * @return The associated load-balancer ID if any
     * 
     */
    public String lbId() {
        return this.lbId;
    }
    /**
     * @return The organization ID the load-balancer is associated with.
     * 
     */
    public String organizationId() {
        return this.organizationId;
    }
    /**
     * @return The ID of the project the load-balancer is associated with.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    public String reverse() {
        return this.reverse;
    }
    /**
     * @return `zone`) The zone in which IPs exist.
     * 
     */
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbIpsIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String ipAddress;
        private String lbId;
        private String organizationId;
        private String projectId;
        private String reverse;
        private String zone;
        public Builder() {}
        public Builder(GetLbIpsIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ipAddress = defaults.ipAddress;
    	      this.lbId = defaults.lbId;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.reverse = defaults.reverse;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddress(String ipAddress) {
            this.ipAddress = Objects.requireNonNull(ipAddress);
            return this;
        }
        @CustomType.Setter
        public Builder lbId(String lbId) {
            this.lbId = Objects.requireNonNull(lbId);
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder reverse(String reverse) {
            this.reverse = Objects.requireNonNull(reverse);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetLbIpsIp build() {
            final var o = new GetLbIpsIp();
            o.id = id;
            o.ipAddress = ipAddress;
            o.lbId = lbId;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.reverse = reverse;
            o.zone = zone;
            return o;
        }
    }
}
