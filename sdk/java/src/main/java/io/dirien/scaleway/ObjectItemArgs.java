// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectItemArgs extends com.pulumi.resources.ResourceArgs {

    public static final ObjectItemArgs Empty = new ObjectItemArgs();

    /**
     * The name of the bucket.
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * The name of the file to upload, defaults to an empty file
     * 
     */
    @Import(name="file")
    private @Nullable Output<String> file;

    /**
     * @return The name of the file to upload, defaults to an empty file
     * 
     */
    public Optional<Output<String>> file() {
        return Optional.ofNullable(this.file);
    }

    /**
     * Hash of the file, used to trigger upload on file change
     * 
     */
    @Import(name="hash")
    private @Nullable Output<String> hash;

    /**
     * @return Hash of the file, used to trigger upload on file change
     * 
     */
    public Optional<Output<String>> hash() {
        return Optional.ofNullable(this.hash);
    }

    /**
     * The path of the object.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The path of the object.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * Map of metadata used for the object, keys must be lowercase
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return Map of metadata used for the object, keys must be lowercase
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the bucket is associated with.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The Scaleway region this bucket resides in.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The Scaleway region this bucket resides in.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    @Import(name="storageClass")
    private @Nullable Output<String> storageClass;

    /**
     * @return Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    public Optional<Output<String>> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }

    /**
     * Map of tags
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Visibility of the object, `public-read` or `private`
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return Visibility of the object, `public-read` or `private`
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private ObjectItemArgs() {}

    private ObjectItemArgs(ObjectItemArgs $) {
        this.bucket = $.bucket;
        this.file = $.file;
        this.hash = $.hash;
        this.key = $.key;
        this.metadata = $.metadata;
        this.projectId = $.projectId;
        this.region = $.region;
        this.storageClass = $.storageClass;
        this.tags = $.tags;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectItemArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectItemArgs $;

        public Builder() {
            $ = new ObjectItemArgs();
        }

        public Builder(ObjectItemArgs defaults) {
            $ = new ObjectItemArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param file The name of the file to upload, defaults to an empty file
         * 
         * @return builder
         * 
         */
        public Builder file(@Nullable Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The name of the file to upload, defaults to an empty file
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param hash Hash of the file, used to trigger upload on file change
         * 
         * @return builder
         * 
         */
        public Builder hash(@Nullable Output<String> hash) {
            $.hash = hash;
            return this;
        }

        /**
         * @param hash Hash of the file, used to trigger upload on file change
         * 
         * @return builder
         * 
         */
        public Builder hash(String hash) {
            return hash(Output.of(hash));
        }

        /**
         * @param key The path of the object.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The path of the object.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param metadata Map of metadata used for the object, keys must be lowercase
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Map of metadata used for the object, keys must be lowercase
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param projectId `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId `project_id`) The ID of the project the bucket is associated with.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The Scaleway region this bucket resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The Scaleway region this bucket resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(@Nullable Output<String> storageClass) {
            $.storageClass = storageClass;
            return this;
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(String storageClass) {
            return storageClass(Output.of(storageClass));
        }

        /**
         * @param tags Map of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param visibility Visibility of the object, `public-read` or `private`
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Visibility of the object, `public-read` or `private`
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public ObjectItemArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
