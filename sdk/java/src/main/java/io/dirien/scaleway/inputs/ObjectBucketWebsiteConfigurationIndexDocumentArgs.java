// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ObjectBucketWebsiteConfigurationIndexDocumentArgs extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketWebsiteConfigurationIndexDocumentArgs Empty = new ObjectBucketWebsiteConfigurationIndexDocumentArgs();

    @Import(name="suffix", required=true)
    private Output<String> suffix;

    public Output<String> suffix() {
        return this.suffix;
    }

    private ObjectBucketWebsiteConfigurationIndexDocumentArgs() {}

    private ObjectBucketWebsiteConfigurationIndexDocumentArgs(ObjectBucketWebsiteConfigurationIndexDocumentArgs $) {
        this.suffix = $.suffix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketWebsiteConfigurationIndexDocumentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketWebsiteConfigurationIndexDocumentArgs $;

        public Builder() {
            $ = new ObjectBucketWebsiteConfigurationIndexDocumentArgs();
        }

        public Builder(ObjectBucketWebsiteConfigurationIndexDocumentArgs defaults) {
            $ = new ObjectBucketWebsiteConfigurationIndexDocumentArgs(Objects.requireNonNull(defaults));
        }

        public Builder suffix(Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        public ObjectBucketWebsiteConfigurationIndexDocumentArgs build() {
            $.suffix = Objects.requireNonNull($.suffix, "expected parameter 'suffix' to be non-null");
            return $;
        }
    }

}
