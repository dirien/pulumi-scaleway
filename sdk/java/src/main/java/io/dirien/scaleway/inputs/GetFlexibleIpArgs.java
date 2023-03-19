// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFlexibleIpArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFlexibleIpArgs Empty = new GetFlexibleIpArgs();

    @Import(name="flexibleIpId")
    private @Nullable Output<String> flexibleIpId;

    public Optional<Output<String>> flexibleIpId() {
        return Optional.ofNullable(this.flexibleIpId);
    }

    /**
     * The IP address.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The IP address.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return (Defaults to provider `project_id`) The ID of the project the IP is in.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    private GetFlexibleIpArgs() {}

    private GetFlexibleIpArgs(GetFlexibleIpArgs $) {
        this.flexibleIpId = $.flexibleIpId;
        this.ipAddress = $.ipAddress;
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFlexibleIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFlexibleIpArgs $;

        public Builder() {
            $ = new GetFlexibleIpArgs();
        }

        public Builder(GetFlexibleIpArgs defaults) {
            $ = new GetFlexibleIpArgs(Objects.requireNonNull(defaults));
        }

        public Builder flexibleIpId(@Nullable Output<String> flexibleIpId) {
            $.flexibleIpId = flexibleIpId;
            return this;
        }

        public Builder flexibleIpId(String flexibleIpId) {
            return flexibleIpId(Output.of(flexibleIpId));
        }

        /**
         * @param ipAddress The IP address.
         * Only one of `ip_address` and `ip_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The IP address.
         * Only one of `ip_address` and `ip_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the IP is in.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId (Defaults to provider `project_id`) The ID of the project the IP is in.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public GetFlexibleIpArgs build() {
            return $;
        }
    }

}
