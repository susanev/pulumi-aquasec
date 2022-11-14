// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGatewaysGateway {
    private String description;
    private String grpcAddress;
    private String hostname;
    /**
     * @return The ID of this resource.
     * 
     */
    private String id;
    private String logicalname;
    private String publicAddress;
    private String status;
    private String version;

    private GetGatewaysGateway() {}
    public String description() {
        return this.description;
    }
    public String grpcAddress() {
        return this.grpcAddress;
    }
    public String hostname() {
        return this.hostname;
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String logicalname() {
        return this.logicalname;
    }
    public String publicAddress() {
        return this.publicAddress;
    }
    public String status() {
        return this.status;
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGatewaysGateway defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String grpcAddress;
        private String hostname;
        private String id;
        private String logicalname;
        private String publicAddress;
        private String status;
        private String version;
        public Builder() {}
        public Builder(GetGatewaysGateway defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.grpcAddress = defaults.grpcAddress;
    	      this.hostname = defaults.hostname;
    	      this.id = defaults.id;
    	      this.logicalname = defaults.logicalname;
    	      this.publicAddress = defaults.publicAddress;
    	      this.status = defaults.status;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder grpcAddress(String grpcAddress) {
            this.grpcAddress = Objects.requireNonNull(grpcAddress);
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            this.hostname = Objects.requireNonNull(hostname);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logicalname(String logicalname) {
            this.logicalname = Objects.requireNonNull(logicalname);
            return this;
        }
        @CustomType.Setter
        public Builder publicAddress(String publicAddress) {
            this.publicAddress = Objects.requireNonNull(publicAddress);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetGatewaysGateway build() {
            final var o = new GetGatewaysGateway();
            o.description = description;
            o.grpcAddress = grpcAddress;
            o.hostname = hostname;
            o.id = id;
            o.logicalname = logicalname;
            o.publicAddress = publicAddress;
            o.status = status;
            o.version = version;
            return o;
        }
    }
}
