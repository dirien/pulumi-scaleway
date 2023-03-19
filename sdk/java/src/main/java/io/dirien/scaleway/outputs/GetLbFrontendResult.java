// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.GetLbFrontendAcl;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetLbFrontendResult {
    private List<GetLbFrontendAcl> acls;
    private String backendId;
    private String certificateId;
    private List<String> certificateIds;
    private Boolean enableHttp3;
    private @Nullable String frontendId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer inboundPort;
    private @Nullable String lbId;
    private @Nullable String name;
    private String timeoutClient;

    private GetLbFrontendResult() {}
    public List<GetLbFrontendAcl> acls() {
        return this.acls;
    }
    public String backendId() {
        return this.backendId;
    }
    public String certificateId() {
        return this.certificateId;
    }
    public List<String> certificateIds() {
        return this.certificateIds;
    }
    public Boolean enableHttp3() {
        return this.enableHttp3;
    }
    public Optional<String> frontendId() {
        return Optional.ofNullable(this.frontendId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer inboundPort() {
        return this.inboundPort;
    }
    public Optional<String> lbId() {
        return Optional.ofNullable(this.lbId);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public String timeoutClient() {
        return this.timeoutClient;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLbFrontendResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetLbFrontendAcl> acls;
        private String backendId;
        private String certificateId;
        private List<String> certificateIds;
        private Boolean enableHttp3;
        private @Nullable String frontendId;
        private String id;
        private Integer inboundPort;
        private @Nullable String lbId;
        private @Nullable String name;
        private String timeoutClient;
        public Builder() {}
        public Builder(GetLbFrontendResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acls = defaults.acls;
    	      this.backendId = defaults.backendId;
    	      this.certificateId = defaults.certificateId;
    	      this.certificateIds = defaults.certificateIds;
    	      this.enableHttp3 = defaults.enableHttp3;
    	      this.frontendId = defaults.frontendId;
    	      this.id = defaults.id;
    	      this.inboundPort = defaults.inboundPort;
    	      this.lbId = defaults.lbId;
    	      this.name = defaults.name;
    	      this.timeoutClient = defaults.timeoutClient;
        }

        @CustomType.Setter
        public Builder acls(List<GetLbFrontendAcl> acls) {
            this.acls = Objects.requireNonNull(acls);
            return this;
        }
        public Builder acls(GetLbFrontendAcl... acls) {
            return acls(List.of(acls));
        }
        @CustomType.Setter
        public Builder backendId(String backendId) {
            this.backendId = Objects.requireNonNull(backendId);
            return this;
        }
        @CustomType.Setter
        public Builder certificateId(String certificateId) {
            this.certificateId = Objects.requireNonNull(certificateId);
            return this;
        }
        @CustomType.Setter
        public Builder certificateIds(List<String> certificateIds) {
            this.certificateIds = Objects.requireNonNull(certificateIds);
            return this;
        }
        public Builder certificateIds(String... certificateIds) {
            return certificateIds(List.of(certificateIds));
        }
        @CustomType.Setter
        public Builder enableHttp3(Boolean enableHttp3) {
            this.enableHttp3 = Objects.requireNonNull(enableHttp3);
            return this;
        }
        @CustomType.Setter
        public Builder frontendId(@Nullable String frontendId) {
            this.frontendId = frontendId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder inboundPort(Integer inboundPort) {
            this.inboundPort = Objects.requireNonNull(inboundPort);
            return this;
        }
        @CustomType.Setter
        public Builder lbId(@Nullable String lbId) {
            this.lbId = lbId;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder timeoutClient(String timeoutClient) {
            this.timeoutClient = Objects.requireNonNull(timeoutClient);
            return this;
        }
        public GetLbFrontendResult build() {
            final var o = new GetLbFrontendResult();
            o.acls = acls;
            o.backendId = backendId;
            o.certificateId = certificateId;
            o.certificateIds = certificateIds;
            o.enableHttp3 = enableHttp3;
            o.frontendId = frontendId;
            o.id = id;
            o.inboundPort = inboundPort;
            o.lbId = lbId;
            o.name = name;
            o.timeoutClient = timeoutClient;
            return o;
        }
    }
}
