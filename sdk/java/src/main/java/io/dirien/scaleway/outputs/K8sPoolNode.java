// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class K8sPoolNode {
    /**
     * @return The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    private @Nullable String name;
    /**
     * @return The public IPv4.
     * 
     */
    private @Nullable String publicIp;
    /**
     * @return The public IPv6.
     * 
     */
    private @Nullable String publicIpV6;
    /**
     * @return The status of the node.
     * 
     */
    private @Nullable String status;

    private K8sPoolNode() {}
    /**
     * @return The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The public IPv4.
     * 
     */
    public Optional<String> publicIp() {
        return Optional.ofNullable(this.publicIp);
    }
    /**
     * @return The public IPv6.
     * 
     */
    public Optional<String> publicIpV6() {
        return Optional.ofNullable(this.publicIpV6);
    }
    /**
     * @return The status of the node.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(K8sPoolNode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String publicIp;
        private @Nullable String publicIpV6;
        private @Nullable String status;
        public Builder() {}
        public Builder(K8sPoolNode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.publicIp = defaults.publicIp;
    	      this.publicIpV6 = defaults.publicIpV6;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder publicIp(@Nullable String publicIp) {
            this.publicIp = publicIp;
            return this;
        }
        @CustomType.Setter
        public Builder publicIpV6(@Nullable String publicIpV6) {
            this.publicIpV6 = publicIpV6;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public K8sPoolNode build() {
            final var o = new K8sPoolNode();
            o.name = name;
            o.publicIp = publicIp;
            o.publicIpV6 = publicIpV6;
            o.status = status;
            return o;
        }
    }
}
