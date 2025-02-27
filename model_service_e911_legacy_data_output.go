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
 * ServiceE911LegacyDataOutput
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceE911LegacyDataOutput {
  @.annotation.Nullable
  private string HouseNumber;

  @.annotation.Nullable
  private string Predirectional;

  @.annotation.Nullable
  private string StreetName;

  @.annotation.Nullable
  private string Suite;

  public ServiceE911LegacyDataOutput() {
  }

  public ServiceE911LegacyDataOutput HouseNumber(@.annotation.Nullable string HouseNumber) {
    
    this.HouseNumber = HouseNumber;
    return this;
  }

  /**
   * Get HouseNumber
   * @return HouseNumber
   */
  @.annotation.Nullable

  public string GetHouseNumber() {
    return HouseNumber;
  }


  public void setHouseNumber(@.annotation.Nullable string HouseNumber) {
    this.HouseNumber = HouseNumber;
  }

  public ServiceE911LegacyDataOutput Predirectional(@.annotation.Nullable string Predirectional) {
    
    this.Predirectional = Predirectional;
    return this;
  }

  /**
   * Get Predirectional
   * @return Predirectional
   */
  @.annotation.Nullable

  public string GetPredirectional() {
    return Predirectional;
  }


  public void setPredirectional(@.annotation.Nullable string Predirectional) {
    this.Predirectional = Predirectional;
  }

  public ServiceE911LegacyDataOutput StreetName(@.annotation.Nullable string StreetName) {
    
    this.StreetName = StreetName;
    return this;
  }

  /**
   * Get StreetName
   * @return StreetName
   */
  @.annotation.Nullable

  public string GetStreetName() {
    return StreetName;
  }


  public void setStreetName(@.annotation.Nullable string StreetName) {
    this.StreetName = StreetName;
  }

  public ServiceE911LegacyDataOutput Suite(@.annotation.Nullable string Suite) {
    
    this.Suite = Suite;
    return this;
  }

  /**
   * Get Suite
   * @return Suite
   */
  @.annotation.Nullable

  public string GetSuite() {
    return Suite;
  }


  public void setSuite(@.annotation.Nullable string Suite) {
    this.Suite = Suite;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceE911LegacyDataOutput ServiceE911LegacyDataOutput = (ServiceE911LegacyDataOutput) o;
    return Objects.equals(this.HouseNumber, ServiceE911LegacyDataOutput.HouseNumber) &&
        Objects.equals(this.Predirectional, ServiceE911LegacyDataOutput.Predirectional) &&
        Objects.equals(this.StreetName, ServiceE911LegacyDataOutput.StreetName) &&
        Objects.equals(this.Suite, ServiceE911LegacyDataOutput.Suite);
  }

  @Override
  public int hashCode() {
    return Objects.hash(HouseNumber, Predirectional, StreetName, Suite);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceE911LegacyDataOutput {\n");
    sb.append("    HouseNumber: ").append(toIndentedString(HouseNumber)).append("\n");
    sb.append("    Predirectional: ").append(toIndentedString(Predirectional)).append("\n");
    sb.append("    StreetName: ").append(toIndentedString(StreetName)).append("\n");
    sb.append("    Suite: ").append(toIndentedString(Suite)).append("\n");
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

