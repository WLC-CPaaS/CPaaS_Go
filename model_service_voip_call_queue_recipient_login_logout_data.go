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
import bytes;
import fmt;

/**
 * ServiceVOIPCallQueueRecipientLoginLogoutData
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceVOIPCallQueueRecipientLoginLogoutData {
  /**
   * Gets or Sets Action
   */
  public enum ACTION {
    LOGIN(string.valueOf("login")),
    
    LOGOUT(string.valueOf("logout"));

    private string value;

    ACTION(string value) {
      this.value = value;
    }

    public string getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ACTION fromValue(string value) {
      for (ACTION b : ACTION.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  @.annotation.Nonnull
  private ACTION Action;

  public ServiceVOIPCallQueueRecipientLoginLogoutData() {
  }

  public ServiceVOIPCallQueueRecipientLoginLogoutData Action(@.annotation.Nonnull ACTION Action) {
    
    this.Action = Action;
    return this;
  }

  /**
   * Get Action
   * @return Action
   */
  @.annotation.Nonnull

  public ACTION GetAction() {
    return Action;
  }


  public void setAction(@.annotation.Nonnull ACTION Action) {
    this.Action = Action;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceVOIPCallQueueRecipientLoginLogoutData ServiceVOIPCallQueueRecipientLoginLogoutData = (ServiceVOIPCallQueueRecipientLoginLogoutData) o;
    return Objects.equals(this.Action, ServiceVOIPCallQueueRecipientLoginLogoutData.Action);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Action);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceVOIPCallQueueRecipientLoginLogoutData {\n");
    sb.append("    Action: ").append(toIndentedString(Action)).append("\n");
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

