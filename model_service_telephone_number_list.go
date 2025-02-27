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
 * ServiceTelephoneNumberList
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceTelephoneNumberList {
  @.annotation.Nullable
  private []string TelephoneNumber;

  public ServiceTelephoneNumberList() {
  }

  public ServiceTelephoneNumberList TelephoneNumber(@.annotation.Nullable []string TelephoneNumber) {
    
    this.TelephoneNumber = TelephoneNumber;
    return this;
  }

  public ServiceTelephoneNumberList addTelephoneNumberItem(string TelephoneNumberItem) {
    if (this.TelephoneNumber == null) {
      this.TelephoneNumber = new ArrayList<>();
    }
    this.TelephoneNumber.add(TelephoneNumberItem);
    return this;
  }

  /**
   * Get TelephoneNumber
   * @return TelephoneNumber
   */
  @.annotation.Nullable

  public []string GetTelephoneNumber() {
    return TelephoneNumber;
  }


  public void setTelephoneNumber(@.annotation.Nullable []string TelephoneNumber) {
    this.TelephoneNumber = TelephoneNumber;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceTelephoneNumberList ServiceTelephoneNumberList = (ServiceTelephoneNumberList) o;
    return Objects.equals(this.TelephoneNumber, ServiceTelephoneNumberList.TelephoneNumber);
  }

  @Override
  public int hashCode() {
    return Objects.hash(TelephoneNumber);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceTelephoneNumberList {\n");
    sb.append("    TelephoneNumber: ").append(toIndentedString(TelephoneNumber)).append("\n");
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

