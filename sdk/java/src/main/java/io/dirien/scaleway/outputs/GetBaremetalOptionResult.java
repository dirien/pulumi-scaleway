// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBaremetalOptionResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Is false if the option could not be added or removed.
     * 
     */
    private Boolean manageable;
    /**
     * @return The name of the option.
     * 
     */
    private @Nullable String name;
    private @Nullable String optionId;
    private String zone;

    private GetBaremetalOptionResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Is false if the option could not be added or removed.
     * 
     */
    public Boolean manageable() {
        return this.manageable;
    }
    /**
     * @return The name of the option.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> optionId() {
        return Optional.ofNullable(this.optionId);
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBaremetalOptionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Boolean manageable;
        private @Nullable String name;
        private @Nullable String optionId;
        private String zone;
        public Builder() {}
        public Builder(GetBaremetalOptionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.manageable = defaults.manageable;
    	      this.name = defaults.name;
    	      this.optionId = defaults.optionId;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder manageable(Boolean manageable) {
            this.manageable = Objects.requireNonNull(manageable);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder optionId(@Nullable String optionId) {
            this.optionId = optionId;
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetBaremetalOptionResult build() {
            final var o = new GetBaremetalOptionResult();
            o.id = id;
            o.manageable = manageable;
            o.name = name;
            o.optionId = optionId;
            o.zone = zone;
            return o;
        }
    }
}
