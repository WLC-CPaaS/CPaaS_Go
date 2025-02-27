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
 * ServiceDocsPhonenumberAssignPayload
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceDocsPhonenumberAssignPayload {
  @.annotation.Nullable
  private string Phonenumber;

  @.annotation.Nullable
  private string ToAccount;

  public ServiceDocsPhonenumberAssignPayload() {
  }

  public ServiceDocsPhonenumberAssignPayload Phonenumber(@.annotation.Nullable string Phonenumber) {
    
    this.Phonenumber = Phonenumber;
    return this;
  }

  /**
   * Get Phonenumber
   * @return Phonenumber
   */
  @.annotation.Nullable

  public string GetPhonenumber() {
    return Phonenumber;
  }


  public void setPhonenumber(@.annotation.Nullable string Phonenumber) {
    this.Phonenumber = Phonenumber;
  }

  public ServiceDocsPhonenumberAssignPayload ToAccount(@.annotation.Nullable string ToAccount) {
    
    this.ToAccount = ToAccount;
    return this;
  }

  /**
   * Get ToAccount
   * @return ToAccount
   */
  @.annotation.Nullable

  public string GetToAccount() {
    return ToAccount;
  }


  public void setToAccount(@.annotation.Nullable string ToAccount) {
    this.ToAccount = ToAccount;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceDocsPhonenumberAssignPayload ServiceDocsPhonenumberAssignPayload = (ServiceDocsPhonenumberAssignPayload) o;
    return Objects.equals(this.Phonenumber, ServiceDocsPhonenumberAssignPayload.Phonenumber) &&
        Objects.equals(this.ToAccount, ServiceDocsPhonenumberAssignPayload.ToAccount);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Phonenumber, ToAccount);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceDocsPhonenumberAssignPayload {\n");
    sb.append("    Phonenumber: ").append(toIndentedString(Phonenumber)).append("\n");
    sb.append("    ToAccount: ").append(toIndentedString(ToAccount)).append("\n");
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

