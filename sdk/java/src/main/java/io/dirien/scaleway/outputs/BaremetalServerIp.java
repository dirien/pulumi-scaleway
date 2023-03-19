// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BaremetalServerIp {
    /**
     * @return The address of the IP.
     * 
     */
    private @Nullable String address;
    /**
     * @return The id of the private network to attach.
     * 
     */
    private @Nullable String id;
    /**
     * @return The reverse of the IP.
     * 
     */
    private @Nullable String reverse;
    /**
     * @return The type of the IP.
     * 
     */
    private @Nullable String version;

    private BaremetalServerIp() {}
    /**
     * @return The address of the IP.
     * 
     */
    public Optional<String> address() {
        return Optional.ofNullable(this.address);
    }
    /**
     * @return The id of the private network to attach.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The reverse of the IP.
     * 
     */
    public Optional<String> reverse() {
        return Optional.ofNullable(this.reverse);
    }
    /**
     * @return The type of the IP.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BaremetalServerIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String address;
        private @Nullable String id;
        private @Nullable String reverse;
        private @Nullable String version;
        public Builder() {}
        public Builder(BaremetalServerIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.id = defaults.id;
    	      this.reverse = defaults.reverse;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder address(@Nullable String address) {
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder reverse(@Nullable String reverse) {
            this.reverse = reverse;
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public BaremetalServerIp build() {
            final var o = new BaremetalServerIp();
            o.address = address;
            o.id = id;
            o.reverse = reverse;
            o.version = version;
            return o;
        }
    }
}
