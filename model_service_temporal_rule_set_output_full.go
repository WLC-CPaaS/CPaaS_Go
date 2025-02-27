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
 * ServiceTemporalRuleSetOutputFull
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceTemporalRuleSetOutputFull {
  @.annotation.Nullable
  private string Id;

  @.annotation.Nullable
  private string Name;

  @.annotation.Nullable
  private []string TemporalRules;

  public ServiceTemporalRuleSetOutputFull() {
  }

  public ServiceTemporalRuleSetOutputFull Id(@.annotation.Nullable string Id) {
    
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

  public ServiceTemporalRuleSetOutputFull Name(@.annotation.Nullable string Name) {
    
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

  public ServiceTemporalRuleSetOutputFull TemporalRules(@.annotation.Nullable []string TemporalRules) {
    
    this.TemporalRules = TemporalRules;
    return this;
  }

  public ServiceTemporalRuleSetOutputFull addTemporalRulesItem(string TemporalRulesItem) {
    if (this.TemporalRules == null) {
      this.TemporalRules = new ArrayList<>();
    }
    this.TemporalRules.add(TemporalRulesItem);
    return this;
  }

  /**
   * Get TemporalRules
   * @return TemporalRules
   */
  @.annotation.Nullable

  public []string GetTemporalRules() {
    return TemporalRules;
  }


  public void setTemporalRules(@.annotation.Nullable []string TemporalRules) {
    this.TemporalRules = TemporalRules;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceTemporalRuleSetOutputFull ServiceTemporalRuleSetOutputFull = (ServiceTemporalRuleSetOutputFull) o;
    return Objects.equals(this.Id, ServiceTemporalRuleSetOutputFull.Id) &&
        Objects.equals(this.Name, ServiceTemporalRuleSetOutputFull.Name) &&
        Objects.equals(this.TemporalRules, ServiceTemporalRuleSetOutputFull.TemporalRules);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Id, Name, TemporalRules);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceTemporalRuleSetOutputFull {\n");
    sb.append("    Id: ").append(toIndentedString(Id)).append("\n");
    sb.append("    Name: ").append(toIndentedString(Name)).append("\n");
    sb.append("    TemporalRules: ").append(toIndentedString(TemporalRules)).append("\n");
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

