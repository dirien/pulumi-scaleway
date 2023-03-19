// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIotDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIotDevicePlainArgs Empty = new GetIotDevicePlainArgs();

    /**
     * The device ID.
     * Only one of the `name` and `device_id` should be specified.
     * 
     */
    @Import(name="deviceId")
    private @Nullable String deviceId;

    /**
     * @return The device ID.
     * Only one of the `name` and `device_id` should be specified.
     * 
     */
    public Optional<String> deviceId() {
        return Optional.ofNullable(this.deviceId);
    }

    /**
     * The hub ID.
     * 
     */
    @Import(name="hubId")
    private @Nullable String hubId;

    /**
     * @return The hub ID.
     * 
     */
    public Optional<String> hubId() {
        return Optional.ofNullable(this.hubId);
    }

    /**
     * The name of the Hub.
     * Only one of the `name` and `device_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the Hub.
     * Only one of the `name` and `device_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the hub exists.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return `region`) The region in which the hub exists.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetIotDevicePlainArgs() {}

    private GetIotDevicePlainArgs(GetIotDevicePlainArgs $) {
        this.deviceId = $.deviceId;
        this.hubId = $.hubId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIotDevicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIotDevicePlainArgs $;

        public Builder() {
            $ = new GetIotDevicePlainArgs();
        }

        public Builder(GetIotDevicePlainArgs defaults) {
            $ = new GetIotDevicePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deviceId The device ID.
         * Only one of the `name` and `device_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(@Nullable String deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param hubId The hub ID.
         * 
         * @return builder
         * 
         */
        public Builder hubId(@Nullable String hubId) {
            $.hubId = hubId;
            return this;
        }

        /**
         * @param name The name of the Hub.
         * Only one of the `name` and `device_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region `region`) The region in which the hub exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetIotDevicePlainArgs build() {
            return $;
        }
    }

}
