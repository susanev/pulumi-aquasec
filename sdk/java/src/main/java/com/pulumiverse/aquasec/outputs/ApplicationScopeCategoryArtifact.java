// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryArtifactCf;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryArtifactFunction;
import com.pulumiverse.aquasec.outputs.ApplicationScopeCategoryArtifactImage;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationScopeCategoryArtifact {
    private @Nullable List<ApplicationScopeCategoryArtifactCf> cfs;
    private @Nullable List<ApplicationScopeCategoryArtifactFunction> functions;
    private @Nullable List<ApplicationScopeCategoryArtifactImage> images;

    private ApplicationScopeCategoryArtifact() {}
    public List<ApplicationScopeCategoryArtifactCf> cfs() {
        return this.cfs == null ? List.of() : this.cfs;
    }
    public List<ApplicationScopeCategoryArtifactFunction> functions() {
        return this.functions == null ? List.of() : this.functions;
    }
    public List<ApplicationScopeCategoryArtifactImage> images() {
        return this.images == null ? List.of() : this.images;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationScopeCategoryArtifact defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<ApplicationScopeCategoryArtifactCf> cfs;
        private @Nullable List<ApplicationScopeCategoryArtifactFunction> functions;
        private @Nullable List<ApplicationScopeCategoryArtifactImage> images;
        public Builder() {}
        public Builder(ApplicationScopeCategoryArtifact defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cfs = defaults.cfs;
    	      this.functions = defaults.functions;
    	      this.images = defaults.images;
        }

        @CustomType.Setter
        public Builder cfs(@Nullable List<ApplicationScopeCategoryArtifactCf> cfs) {
            this.cfs = cfs;
            return this;
        }
        public Builder cfs(ApplicationScopeCategoryArtifactCf... cfs) {
            return cfs(List.of(cfs));
        }
        @CustomType.Setter
        public Builder functions(@Nullable List<ApplicationScopeCategoryArtifactFunction> functions) {
            this.functions = functions;
            return this;
        }
        public Builder functions(ApplicationScopeCategoryArtifactFunction... functions) {
            return functions(List.of(functions));
        }
        @CustomType.Setter
        public Builder images(@Nullable List<ApplicationScopeCategoryArtifactImage> images) {
            this.images = images;
            return this;
        }
        public Builder images(ApplicationScopeCategoryArtifactImage... images) {
            return images(List.of(images));
        }
        public ApplicationScopeCategoryArtifact build() {
            final var o = new ApplicationScopeCategoryArtifact();
            o.cfs = cfs;
            o.functions = functions;
            o.images = images;
            return o;
        }
    }
}
