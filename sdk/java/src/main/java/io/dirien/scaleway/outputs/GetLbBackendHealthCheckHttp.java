// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLbBackendHealthCheckHttp {
    private Integer code;
    private String hostHeader;
    private String method;
    private String sni;
    private String uri;

    private GetLbBackendHealthCheckHttp() {}
    public Integer code() {
        return this.code;
    }
    public String hostHeader() {
        return this.hostHeader;
    }
    public String method() {
        return this.method;
    }
    public String sni() {
        return this.sni;
    }
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbBackendHealthCheckHttp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer code;
        private String hostHeader;
        private String method;
        private String sni;
        private String uri;
        public Builder() {}
        public Builder(GetLbBackendHealthCheckHttp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.code = defaults.code;
    	      this.hostHeader = defaults.hostHeader;
    	      this.method = defaults.method;
    	      this.sni = defaults.sni;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder code(Integer code) {
            this.code = Objects.requireNonNull(code);
            return this;
        }
        @CustomType.Setter
        public Builder hostHeader(String hostHeader) {
            this.hostHeader = Objects.requireNonNull(hostHeader);
            return this;
        }
        @CustomType.Setter
        public Builder method(String method) {
            this.method = Objects.requireNonNull(method);
            return this;
        }
        @CustomType.Setter
        public Builder sni(String sni) {
            this.sni = Objects.requireNonNull(sni);
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public GetLbBackendHealthCheckHttp build() {
            final var o = new GetLbBackendHealthCheckHttp();
            o.code = code;
            o.hostHeader = hostHeader;
            o.method = method;
            o.sni = sni;
            o.uri = uri;
            return o;
        }
    }
}
