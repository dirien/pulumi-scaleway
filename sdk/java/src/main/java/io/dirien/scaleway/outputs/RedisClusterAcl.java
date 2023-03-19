// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RedisClusterAcl {
    /**
     * @return A text describing this rule. Default description: `Allow IP`
     * 
     */
    private @Nullable String description;
    /**
     * @return The UUID of the private network resource.
     * 
     */
    private @Nullable String id;
    /**
     * @return The ip range to whitelist
     * in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
     * 
     */
    private String ip;

    private RedisClusterAcl() {}
    /**
     * @return A text describing this rule. Default description: `Allow IP`
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The UUID of the private network resource.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The ip range to whitelist
     * in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
     * 
     */
    public String ip() {
        return this.ip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RedisClusterAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String id;
        private String ip;
        public Builder() {}
        public Builder(RedisClusterAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        public RedisClusterAcl build() {
            final var o = new RedisClusterAcl();
            o.description = description;
            o.id = id;
            o.ip = ip;
            return o;
        }
    }
}
