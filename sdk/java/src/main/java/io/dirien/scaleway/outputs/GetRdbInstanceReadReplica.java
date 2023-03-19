// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRdbInstanceReadReplica {
    private String ip;
    /**
     * @return The name of the RDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    private String name;
    private Integer port;

    private GetRdbInstanceReadReplica() {}
    public String ip() {
        return this.ip;
    }
    /**
     * @return The name of the RDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    public String name() {
        return this.name;
    }
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRdbInstanceReadReplica defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private String name;
        private Integer port;
        public Builder() {}
        public Builder(GetRdbInstanceReadReplica defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.name = defaults.name;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public GetRdbInstanceReadReplica build() {
            final var o = new GetRdbInstanceReadReplica();
            o.ip = ip;
            o.name = name;
            o.port = port;
            return o;
        }
    }
}
