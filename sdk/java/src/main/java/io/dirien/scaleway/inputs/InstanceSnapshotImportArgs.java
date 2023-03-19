// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class InstanceSnapshotImportArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceSnapshotImportArgs Empty = new InstanceSnapshotImportArgs();

    /**
     * Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * Key of the object to import
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return Key of the object to import
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    private InstanceSnapshotImportArgs() {}

    private InstanceSnapshotImportArgs(InstanceSnapshotImportArgs $) {
        this.bucket = $.bucket;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceSnapshotImportArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceSnapshotImportArgs $;

        public Builder() {
            $ = new InstanceSnapshotImportArgs();
        }

        public Builder(InstanceSnapshotImportArgs defaults) {
            $ = new InstanceSnapshotImportArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket Bucket name containing [qcow2](https://en.wikipedia.org/wiki/Qcow) to import
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param key Key of the object to import
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Key of the object to import
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public InstanceSnapshotImportArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
