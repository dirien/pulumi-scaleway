// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIotDeviceMessageFilterPublish {
    private String policy;
    private List<String> topics;

    private GetIotDeviceMessageFilterPublish() {}
    public String policy() {
        return this.policy;
    }
    public List<String> topics() {
        return this.topics;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIotDeviceMessageFilterPublish defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String policy;
        private List<String> topics;
        public Builder() {}
        public Builder(GetIotDeviceMessageFilterPublish defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.policy = defaults.policy;
    	      this.topics = defaults.topics;
        }

        @CustomType.Setter
        public Builder policy(String policy) {
            this.policy = Objects.requireNonNull(policy);
            return this;
        }
        @CustomType.Setter
        public Builder topics(List<String> topics) {
            this.topics = Objects.requireNonNull(topics);
            return this;
        }
        public Builder topics(String... topics) {
            return topics(List.of(topics));
        }
        public GetIotDeviceMessageFilterPublish build() {
            final var o = new GetIotDeviceMessageFilterPublish();
            o.policy = policy;
            o.topics = topics;
            return o;
        }
    }
}
