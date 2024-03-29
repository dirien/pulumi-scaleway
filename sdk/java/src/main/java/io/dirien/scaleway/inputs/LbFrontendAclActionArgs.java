// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.LbFrontendAclActionRedirectArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LbFrontendAclActionArgs extends com.pulumi.resources.ResourceArgs {

    public static final LbFrontendAclActionArgs Empty = new LbFrontendAclActionArgs();

    /**
     * Redirect parameters when using an ACL with `redirect` action.
     * 
     */
    @Import(name="redirects")
    private @Nullable Output<List<LbFrontendAclActionRedirectArgs>> redirects;

    /**
     * @return Redirect parameters when using an ACL with `redirect` action.
     * 
     */
    public Optional<Output<List<LbFrontendAclActionRedirectArgs>>> redirects() {
        return Optional.ofNullable(this.redirects);
    }

    /**
     * The redirect type. Possible values are: `location` or `scheme`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The redirect type. Possible values are: `location` or `scheme`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private LbFrontendAclActionArgs() {}

    private LbFrontendAclActionArgs(LbFrontendAclActionArgs $) {
        this.redirects = $.redirects;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LbFrontendAclActionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LbFrontendAclActionArgs $;

        public Builder() {
            $ = new LbFrontendAclActionArgs();
        }

        public Builder(LbFrontendAclActionArgs defaults) {
            $ = new LbFrontendAclActionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param redirects Redirect parameters when using an ACL with `redirect` action.
         * 
         * @return builder
         * 
         */
        public Builder redirects(@Nullable Output<List<LbFrontendAclActionRedirectArgs>> redirects) {
            $.redirects = redirects;
            return this;
        }

        /**
         * @param redirects Redirect parameters when using an ACL with `redirect` action.
         * 
         * @return builder
         * 
         */
        public Builder redirects(List<LbFrontendAclActionRedirectArgs> redirects) {
            return redirects(Output.of(redirects));
        }

        /**
         * @param redirects Redirect parameters when using an ACL with `redirect` action.
         * 
         * @return builder
         * 
         */
        public Builder redirects(LbFrontendAclActionRedirectArgs... redirects) {
            return redirects(List.of(redirects));
        }

        /**
         * @param type The redirect type. Possible values are: `location` or `scheme`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The redirect type. Possible values are: `location` or `scheme`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public LbFrontendAclActionArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
