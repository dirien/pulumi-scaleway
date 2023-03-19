// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceServerPrivateNetwork {
    private @Nullable String macAddress;
    private String pnId;
    private @Nullable String status;
    /**
     * @return `zone`) The zone in which the server should be created.
     * 
     */
    private @Nullable String zone;

    private InstanceServerPrivateNetwork() {}
    public Optional<String> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }
    public String pnId() {
        return this.pnId;
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return `zone`) The zone in which the server should be created.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceServerPrivateNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String macAddress;
        private String pnId;
        private @Nullable String status;
        private @Nullable String zone;
        public Builder() {}
        public Builder(InstanceServerPrivateNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.macAddress = defaults.macAddress;
    	      this.pnId = defaults.pnId;
    	      this.status = defaults.status;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder macAddress(@Nullable String macAddress) {
            this.macAddress = macAddress;
            return this;
        }
        @CustomType.Setter
        public Builder pnId(String pnId) {
            this.pnId = Objects.requireNonNull(pnId);
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder zone(@Nullable String zone) {
            this.zone = zone;
            return this;
        }
        public InstanceServerPrivateNetwork build() {
            final var o = new InstanceServerPrivateNetwork();
            o.macAddress = macAddress;
            o.pnId = pnId;
            o.status = status;
            o.zone = zone;
            return o;
        }
    }
}
