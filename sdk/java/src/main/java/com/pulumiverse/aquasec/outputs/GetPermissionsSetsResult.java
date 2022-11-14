// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumiverse.aquasec.outputs.GetPermissionsSetsPermissionsSet;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetPermissionsSetsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<GetPermissionsSetsPermissionsSet> permissionsSets;

    private GetPermissionsSetsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetPermissionsSetsPermissionsSet> permissionsSets() {
        return this.permissionsSets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPermissionsSetsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetPermissionsSetsPermissionsSet> permissionsSets;
        public Builder() {}
        public Builder(GetPermissionsSetsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.permissionsSets = defaults.permissionsSets;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder permissionsSets(List<GetPermissionsSetsPermissionsSet> permissionsSets) {
            this.permissionsSets = Objects.requireNonNull(permissionsSets);
            return this;
        }
        public Builder permissionsSets(GetPermissionsSetsPermissionsSet... permissionsSets) {
            return permissionsSets(List.of(permissionsSets));
        }
        public GetPermissionsSetsResult build() {
            final var o = new GetPermissionsSetsResult();
            o.id = id;
            o.permissionsSets = permissionsSets;
            return o;
        }
    }
}
