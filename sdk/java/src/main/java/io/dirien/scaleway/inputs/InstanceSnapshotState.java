// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.InstanceSnapshotImportArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceSnapshotState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceSnapshotState Empty = new InstanceSnapshotState();

    /**
     * The snapshot creation time.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The snapshot creation time.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Import a snapshot from a qcow2 file located in a bucket
     * 
     */
    @Import(name="import")
    private @Nullable Output<InstanceSnapshotImportArgs> import_;

    /**
     * @return Import a snapshot from a qcow2 file located in a bucket
     * 
     */
    public Optional<Output<InstanceSnapshotImportArgs>> import_() {
        return Optional.ofNullable(this.import_);
    }

    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the snapshot. If not provided it will be randomly generated.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The organization ID the snapshot is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return The organization ID the snapshot is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * `project_id`) The ID of the project the snapshot is
     * associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the snapshot is
     * associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * (Optional) The size of the snapshot.
     * 
     */
    @Import(name="sizeInGb")
    private @Nullable Output<Integer> sizeInGb;

    /**
     * @return (Optional) The size of the snapshot.
     * 
     */
    public Optional<Output<Integer>> sizeInGb() {
        return Optional.ofNullable(this.sizeInGb);
    }

    /**
     * A list of tags to apply to the snapshot.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tags to apply to the snapshot.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The snapshot&#39;s volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
     * Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The snapshot&#39;s volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
     * Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The ID of the volume to take a snapshot from.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The ID of the volume to take a snapshot from.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * `zone`) The zone in which
     * the snapshot should be created.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which
     * the snapshot should be created.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private InstanceSnapshotState() {}

    private InstanceSnapshotState(InstanceSnapshotState $) {
        this.createdAt = $.createdAt;
        this.import_ = $.import_;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.projectId = $.projectId;
        this.sizeInGb = $.sizeInGb;
        this.tags = $.tags;
        this.type = $.type;
        this.volumeId = $.volumeId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceSnapshotState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceSnapshotState $;

        public Builder() {
            $ = new InstanceSnapshotState();
        }

        public Builder(InstanceSnapshotState defaults) {
            $ = new InstanceSnapshotState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt The snapshot creation time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The snapshot creation time.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param import_ Import a snapshot from a qcow2 file located in a bucket
         * 
         * @return builder
         * 
         */
        public Builder import_(@Nullable Output<InstanceSnapshotImportArgs> import_) {
            $.import_ = import_;
            return this;
        }

        /**
         * @param import_ Import a snapshot from a qcow2 file located in a bucket
         * 
         * @return builder
         * 
         */
        public Builder import_(InstanceSnapshotImportArgs import_) {
            return import_(Output.of(import_));
        }

        /**
         * @param name The name of the snapshot. If not provided it will be randomly generated.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the snapshot. If not provided it will be randomly generated.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId The organization ID the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId The organization ID the snapshot is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param projectId `project_id`) The ID of the project the snapshot is
         * associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the snapshot is
         * associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param sizeInGb (Optional) The size of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder sizeInGb(@Nullable Output<Integer> sizeInGb) {
            $.sizeInGb = sizeInGb;
            return this;
        }

        /**
         * @param sizeInGb (Optional) The size of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder sizeInGb(Integer sizeInGb) {
            return sizeInGb(Output.of(sizeInGb));
        }

        /**
         * @param tags A list of tags to apply to the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tags to apply to the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of tags to apply to the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type The snapshot&#39;s volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
         * Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The snapshot&#39;s volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
         * Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param volumeId The ID of the volume to take a snapshot from.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The ID of the volume to take a snapshot from.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        /**
         * @param zone `zone`) The zone in which
         * the snapshot should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which
         * the snapshot should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public InstanceSnapshotState build() {
            return $;
        }
    }

}
