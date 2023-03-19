// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLbsLbInstance {
    /**
     * @return Date at which the Load balancer was created.
     * 
     */
    private String createdAt;
    /**
     * @return The ID of the load-balancer.
     * 
     */
    private String id;
    private String ipAddress;
    /**
     * @return The state of the LB&#39;s instance. Possible values are: `unknown`, `ready`, `pending`, `stopped`, `error`, `locked` and `migrating`.
     * 
     */
    private String status;
    /**
     * @return Date at which the Load balancer was updated.
     * 
     */
    private String updatedAt;
    /**
     * @return `zone`) The zone in which LBs exist.
     * 
     */
    private String zone;

    private GetLbsLbInstance() {}
    /**
     * @return Date at which the Load balancer was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The ID of the load-balancer.
     * 
     */
    public String id() {
        return this.id;
    }
    public String ipAddress() {
        return this.ipAddress;
    }
    /**
     * @return The state of the LB&#39;s instance. Possible values are: `unknown`, `ready`, `pending`, `stopped`, `error`, `locked` and `migrating`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Date at which the Load balancer was updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return `zone`) The zone in which LBs exist.
     * 
     */
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbsLbInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private String ipAddress;
        private String status;
        private String updatedAt;
        private String zone;
        public Builder() {}
        public Builder(GetLbsLbInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.ipAddress = defaults.ipAddress;
    	      this.status = defaults.status;
    	      this.updatedAt = defaults.updatedAt;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
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
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetLbsLbInstance build() {
            final var o = new GetLbsLbInstance();
            o.createdAt = createdAt;
            o.id = id;
            o.ipAddress = ipAddress;
            o.status = status;
            o.updatedAt = updatedAt;
            o.zone = zone;
            return o;
        }
    }
}
