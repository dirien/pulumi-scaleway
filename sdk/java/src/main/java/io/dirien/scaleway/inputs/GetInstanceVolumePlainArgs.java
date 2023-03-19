// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceVolumePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceVolumePlainArgs Empty = new GetInstanceVolumePlainArgs();

    /**
     * The volume name.
     * Only one of `name` and `volume_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The volume name.
     * Only one of `name` and `volume_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The volume id.
     * Only one of `name` and `volume_id` should be specified.
     * 
     */
    @Import(name="volumeId")
    private @Nullable String volumeId;

    /**
     * @return The volume id.
     * Only one of `name` and `volume_id` should be specified.
     * 
     */
    public Optional<String> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * `zone`) The zone in which the volume exists.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which the volume exists.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetInstanceVolumePlainArgs() {}

    private GetInstanceVolumePlainArgs(GetInstanceVolumePlainArgs $) {
        this.name = $.name;
        this.volumeId = $.volumeId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceVolumePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceVolumePlainArgs $;

        public Builder() {
            $ = new GetInstanceVolumePlainArgs();
        }

        public Builder(GetInstanceVolumePlainArgs defaults) {
            $ = new GetInstanceVolumePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The volume name.
         * Only one of `name` and `volume_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param volumeId The volume id.
         * Only one of `name` and `volume_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable String volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the volume exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetInstanceVolumePlainArgs build() {
            return $;
        }
    }

}
