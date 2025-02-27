/*
 * White Label Communications CPaas API Documentation
 * A CPaaS platform API
 *
 * The version of the OpenAPI document: 1.1
 * Contact: support@whitelabelcomm.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package openapi;

import java.util.Objects;
import java.util.Arrays;

/**
 * ServiceTemporalRuleOutputShort
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceTemporalRuleOutputShort {
  @.annotation.Nullable
  private bool Enabled;

  @.annotation.Nullable
  private []string Features;

  @.annotation.Nullable
  private []string Flags;

  @.annotation.Nullable
  private string Id;

  @.annotation.Nullable
  private string Name;

  public ServiceTemporalRuleOutputShort() {
  }

  public ServiceTemporalRuleOutputShort Enabled(@.annotation.Nullable bool Enabled) {
    
    this.Enabled = Enabled;
    return this;
  }

  /**
   * Get Enabled
   * @return Enabled
   */
  @.annotation.Nullable

  public bool getEnabled() {
    return Enabled;
  }


  public void setEnabled(@.annotation.Nullable bool Enabled) {
    this.Enabled = Enabled;
  }

  public ServiceTemporalRuleOutputShort Features(@.annotation.Nullable []string Features) {
    
    this.Features = Features;
    return this;
  }

  public ServiceTemporalRuleOutputShort addFeaturesItem(string FeaturesItem) {
    if (this.Features == null) {
      this.Features = new ArrayList<>();
    }
    this.Features.add(FeaturesItem);
    return this;
  }

  /**
   * Get Features
   * @return Features
   */
  @.annotation.Nullable

  public []string GetFeatures() {
    return Features;
  }


  public void setFeatures(@.annotation.Nullable []string Features) {
    this.Features = Features;
  }

  public ServiceTemporalRuleOutputShort Flags(@.annotation.Nullable []string Flags) {
    
    this.Flags = Flags;
    return this;
  }

  public ServiceTemporalRuleOutputShort addFlagsItem(string FlagsItem) {
    if (this.Flags == null) {
      this.Flags = new ArrayList<>();
    }
    this.Flags.add(FlagsItem);
    return this;
  }

  /**
   * Get Flags
   * @return Flags
   */
  @.annotation.Nullable

  public []string GetFlags() {
    return Flags;
  }


  public void setFlags(@.annotation.Nullable []string Flags) {
    this.Flags = Flags;
  }

  public ServiceTemporalRuleOutputShort Id(@.annotation.Nullable string Id) {
    
    this.Id = Id;
    return this;
  }

  /**
   * Get Id
   * @return Id
   */
  @.annotation.Nullable

  public string GetId() {
    return Id;
  }


  public void setId(@.annotation.Nullable string Id) {
    this.Id = Id;
  }

  public ServiceTemporalRuleOutputShort Name(@.annotation.Nullable string Name) {
    
    this.Name = Name;
    return this;
  }

  /**
   * Get Name
   * @return Name
   */
  @.annotation.Nullable

  public string GetName() {
    return Name;
  }


  public void setName(@.annotation.Nullable string Name) {
    this.Name = Name;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceTemporalRuleOutputShort ServiceTemporalRuleOutputShort = (ServiceTemporalRuleOutputShort) o;
    return Objects.equals(this.Enabled, ServiceTemporalRuleOutputShort.Enabled) &&
        Objects.equals(this.Features, ServiceTemporalRuleOutputShort.Features) &&
        Objects.equals(this.Flags, ServiceTemporalRuleOutputShort.Flags) &&
        Objects.equals(this.Id, ServiceTemporalRuleOutputShort.Id) &&
        Objects.equals(this.Name, ServiceTemporalRuleOutputShort.Name);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Enabled, Features, Flags, Id, Name);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceTemporalRuleOutputShort {\n");
    sb.append("    Enabled: ").append(toIndentedString(Enabled)).append("\n");
    sb.append("    Features: ").append(toIndentedString(Features)).append("\n");
    sb.append("    Flags: ").append(toIndentedString(Flags)).append("\n");
    sb.append("    Id: ").append(toIndentedString(Id)).append("\n");
    sb.append("    Name: ").append(toIndentedString(Name)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

