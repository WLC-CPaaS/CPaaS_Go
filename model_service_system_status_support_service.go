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
 * ServiceSystemStatusSupportService
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceSystemStatusSupportService {
  @.annotation.Nullable
  private string E911Server;

  @.annotation.Nullable
  private string PhoneNumberServer;

  public ServiceSystemStatusSupportService() {
  }

  public ServiceSystemStatusSupportService E911Server(@.annotation.Nullable string E911Server) {
    
    this.E911Server = E911Server;
    return this;
  }

  /**
   * Get E911Server
   * @return E911Server
   */
  @.annotation.Nullable

  public string GetE911Server() {
    return E911Server;
  }


  public void setE911Server(@.annotation.Nullable string E911Server) {
    this.E911Server = E911Server;
  }

  public ServiceSystemStatusSupportService PhoneNumberServer(@.annotation.Nullable string PhoneNumberServer) {
    
    this.PhoneNumberServer = PhoneNumberServer;
    return this;
  }

  /**
   * Get PhoneNumberServer
   * @return PhoneNumberServer
   */
  @.annotation.Nullable

  public string GetPhoneNumberServer() {
    return PhoneNumberServer;
  }


  public void setPhoneNumberServer(@.annotation.Nullable string PhoneNumberServer) {
    this.PhoneNumberServer = PhoneNumberServer;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceSystemStatusSupportService ServiceSystemStatusSupportService = (ServiceSystemStatusSupportService) o;
    return Objects.equals(this.E911Server, ServiceSystemStatusSupportService.E911Server) &&
        Objects.equals(this.PhoneNumberServer, ServiceSystemStatusSupportService.PhoneNumberServer);
  }

  @Override
  public int hashCode() {
    return Objects.hash(E911Server, PhoneNumberServer);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceSystemStatusSupportService {\n");
    sb.append("    E911Server: ").append(toIndentedString(E911Server)).append("\n");
    sb.append("    PhoneNumberServer: ").append(toIndentedString(PhoneNumberServer)).append("\n");
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

