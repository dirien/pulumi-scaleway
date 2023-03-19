// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetLbIpResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String ipAddress;
    private @Nullable String ipId;
    /**
     * @return The associated load-balancer ID if any
     * 
     */
    private String lbId;
    /**
     * @return (Defaults to provider `organization_id`) The ID of the organization the LB IP is associated with.
     * 
     */
    private String organizationId;
    private String projectId;
    private String region;
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    private String reverse;
    private String zone;

    private GetLbIpResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    public Optional<String> ipId() {
        return Optional.ofNullable(this.ipId);
    }
    /**
     * @return The associated load-balancer ID if any
     * 
     */
    public String lbId() {
        return this.lbId;
    }
    /**
     * @return (Defaults to provider `organization_id`) The ID of the organization the LB IP is associated with.
     * 
     */
    public String organizationId() {
        return this.organizationId;
    }
    public String projectId() {
        return this.projectId;
    }
    public String region() {
        return this.region;
    }
    /**
     * @return The reverse domain associated with this IP.
     * 
     */
    public String reverse() {
        return this.reverse;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbIpResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String ipAddress;
        private @Nullable String ipId;
        private String lbId;
        private String organizationId;
        private String projectId;
        private String region;
        private String reverse;
        private String zone;
        public Builder() {}
        public Builder(GetLbIpResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ipAddress = defaults.ipAddress;
    	      this.ipId = defaults.ipId;
    	      this.lbId = defaults.lbId;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.reverse = defaults.reverse;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddress(@Nullable String ipAddress) {
            this.ipAddress = ipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder ipId(@Nullable String ipId) {
            this.ipId = ipId;
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
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
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
        public GetLbIpResult build() {
            final var o = new GetLbIpResult();
            o.id = id;
            o.ipAddress = ipAddress;
            o.ipId = ipId;
            o.lbId = lbId;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.region = region;
            o.reverse = reverse;
            o.zone = zone;
            return o;
        }
    }
}
