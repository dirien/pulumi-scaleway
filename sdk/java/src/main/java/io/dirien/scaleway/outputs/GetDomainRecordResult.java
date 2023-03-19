// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.GetDomainRecordGeoIp;
import io.dirien.scaleway.outputs.GetDomainRecordHttpService;
import io.dirien.scaleway.outputs.GetDomainRecordView;
import io.dirien.scaleway.outputs.GetDomainRecordWeighted;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDomainRecordResult {
    private @Nullable String data;
    private @Nullable String dnsZone;
    /**
     * @return Dynamic record base on user geolocalisation (More information about dynamic records)
     * 
     */
    private List<GetDomainRecordGeoIp> geoIps;
    /**
     * @return Dynamic record base on URL resolve (More information about dynamic records)
     * 
     */
    private List<GetDomainRecordHttpService> httpServices;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean keepEmptyZone;
    private @Nullable String name;
    /**
     * @return The priority of the record (mostly used with an `MX` record)
     * 
     */
    private Integer priority;
    private String projectId;
    private @Nullable String recordId;
    private Boolean rootZone;
    /**
     * @return Time To Live of the record in seconds.
     * 
     */
    private Integer ttl;
    private @Nullable String type;
    /**
     * @return Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
     * 
     */
    private List<GetDomainRecordView> views;
    /**
     * @return Dynamic record base on IP weights (More information about dynamic records)
     * 
     */
    private List<GetDomainRecordWeighted> weighteds;

    private GetDomainRecordResult() {}
    public Optional<String> data() {
        return Optional.ofNullable(this.data);
    }
    public Optional<String> dnsZone() {
        return Optional.ofNullable(this.dnsZone);
    }
    /**
     * @return Dynamic record base on user geolocalisation (More information about dynamic records)
     * 
     */
    public List<GetDomainRecordGeoIp> geoIps() {
        return this.geoIps;
    }
    /**
     * @return Dynamic record base on URL resolve (More information about dynamic records)
     * 
     */
    public List<GetDomainRecordHttpService> httpServices() {
        return this.httpServices;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean keepEmptyZone() {
        return this.keepEmptyZone;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The priority of the record (mostly used with an `MX` record)
     * 
     */
    public Integer priority() {
        return this.priority;
    }
    public String projectId() {
        return this.projectId;
    }
    public Optional<String> recordId() {
        return Optional.ofNullable(this.recordId);
    }
    public Boolean rootZone() {
        return this.rootZone;
    }
    /**
     * @return Time To Live of the record in seconds.
     * 
     */
    public Integer ttl() {
        return this.ttl;
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
     * 
     */
    public List<GetDomainRecordView> views() {
        return this.views;
    }
    /**
     * @return Dynamic record base on IP weights (More information about dynamic records)
     * 
     */
    public List<GetDomainRecordWeighted> weighteds() {
        return this.weighteds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainRecordResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String data;
        private @Nullable String dnsZone;
        private List<GetDomainRecordGeoIp> geoIps;
        private List<GetDomainRecordHttpService> httpServices;
        private String id;
        private Boolean keepEmptyZone;
        private @Nullable String name;
        private Integer priority;
        private String projectId;
        private @Nullable String recordId;
        private Boolean rootZone;
        private Integer ttl;
        private @Nullable String type;
        private List<GetDomainRecordView> views;
        private List<GetDomainRecordWeighted> weighteds;
        public Builder() {}
        public Builder(GetDomainRecordResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.data = defaults.data;
    	      this.dnsZone = defaults.dnsZone;
    	      this.geoIps = defaults.geoIps;
    	      this.httpServices = defaults.httpServices;
    	      this.id = defaults.id;
    	      this.keepEmptyZone = defaults.keepEmptyZone;
    	      this.name = defaults.name;
    	      this.priority = defaults.priority;
    	      this.projectId = defaults.projectId;
    	      this.recordId = defaults.recordId;
    	      this.rootZone = defaults.rootZone;
    	      this.ttl = defaults.ttl;
    	      this.type = defaults.type;
    	      this.views = defaults.views;
    	      this.weighteds = defaults.weighteds;
        }

        @CustomType.Setter
        public Builder data(@Nullable String data) {
            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder dnsZone(@Nullable String dnsZone) {
            this.dnsZone = dnsZone;
            return this;
        }
        @CustomType.Setter
        public Builder geoIps(List<GetDomainRecordGeoIp> geoIps) {
            this.geoIps = Objects.requireNonNull(geoIps);
            return this;
        }
        public Builder geoIps(GetDomainRecordGeoIp... geoIps) {
            return geoIps(List.of(geoIps));
        }
        @CustomType.Setter
        public Builder httpServices(List<GetDomainRecordHttpService> httpServices) {
            this.httpServices = Objects.requireNonNull(httpServices);
            return this;
        }
        public Builder httpServices(GetDomainRecordHttpService... httpServices) {
            return httpServices(List.of(httpServices));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder keepEmptyZone(Boolean keepEmptyZone) {
            this.keepEmptyZone = Objects.requireNonNull(keepEmptyZone);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder recordId(@Nullable String recordId) {
            this.recordId = recordId;
            return this;
        }
        @CustomType.Setter
        public Builder rootZone(Boolean rootZone) {
            this.rootZone = Objects.requireNonNull(rootZone);
            return this;
        }
        @CustomType.Setter
        public Builder ttl(Integer ttl) {
            this.ttl = Objects.requireNonNull(ttl);
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder views(List<GetDomainRecordView> views) {
            this.views = Objects.requireNonNull(views);
            return this;
        }
        public Builder views(GetDomainRecordView... views) {
            return views(List.of(views));
        }
        @CustomType.Setter
        public Builder weighteds(List<GetDomainRecordWeighted> weighteds) {
            this.weighteds = Objects.requireNonNull(weighteds);
            return this;
        }
        public Builder weighteds(GetDomainRecordWeighted... weighteds) {
            return weighteds(List.of(weighteds));
        }
        public GetDomainRecordResult build() {
            final var o = new GetDomainRecordResult();
            o.data = data;
            o.dnsZone = dnsZone;
            o.geoIps = geoIps;
            o.httpServices = httpServices;
            o.id = id;
            o.keepEmptyZone = keepEmptyZone;
            o.name = name;
            o.priority = priority;
            o.projectId = projectId;
            o.recordId = recordId;
            o.rootZone = rootZone;
            o.ttl = ttl;
            o.type = type;
            o.views = views;
            o.weighteds = weighteds;
            return o;
        }
    }
}
