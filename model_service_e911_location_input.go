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
 * ServiceE911LocationInput
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceE911LocationInput {
  @.annotation.Nonnull
  private string Address1;

  @.annotation.Nullable
  private string Address2;

  @.annotation.Nonnull
  private string Community;

  @.annotation.Nullable
  private string PlusFour;

  @.annotation.Nonnull
  private string PostalCode;

  @.annotation.Nonnull
  private string State;

  /**
   * Gets or Sets Type
   */
  public enum TYPE {
    ADDRESS(string.valueOf("ADDRESS"));

    private string value;

    TYPE(string value) {
      this.value = value;
    }

    public string getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TYPE fromValue(string value) {
      for (TYPE b : TYPE.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  @.annotation.Nullable
  private TYPE Type;

  public ServiceE911LocationInput() {
  }

  public ServiceE911LocationInput Address1(@.annotation.Nonnull string Address1) {
    
    this.Address1 = Address1;
    return this;
  }

  /**
   * Get Address1
   * @return Address1
   */
  @.annotation.Nonnull

  public string GetAddress1() {
    return Address1;
  }


  public void setAddress1(@.annotation.Nonnull string Address1) {
    this.Address1 = Address1;
  }

  public ServiceE911LocationInput Address2(@.annotation.Nullable string Address2) {
    
    this.Address2 = Address2;
    return this;
  }

  /**
   * Get Address2
   * @return Address2
   */
  @.annotation.Nullable

  public string GetAddress2() {
    return Address2;
  }


  public void setAddress2(@.annotation.Nullable string Address2) {
    this.Address2 = Address2;
  }

  public ServiceE911LocationInput Community(@.annotation.Nonnull string Community) {
    
    this.Community = Community;
    return this;
  }

  /**
   * Get Community
   * @return Community
   */
  @.annotation.Nonnull

  public string GetCommunity() {
    return Community;
  }


  public void setCommunity(@.annotation.Nonnull string Community) {
    this.Community = Community;
  }

  public ServiceE911LocationInput PlusFour(@.annotation.Nullable string PlusFour) {
    
    this.PlusFour = PlusFour;
    return this;
  }

  /**
   * Get PlusFour
   * @return PlusFour
   */
  @.annotation.Nullable

  public string GetPlusFour() {
    return PlusFour;
  }


  public void setPlusFour(@.annotation.Nullable string PlusFour) {
    this.PlusFour = PlusFour;
  }

  public ServiceE911LocationInput PostalCode(@.annotation.Nonnull string PostalCode) {
    
    this.PostalCode = PostalCode;
    return this;
  }

  /**
   * Get PostalCode
   * @return PostalCode
   */
  @.annotation.Nonnull

  public string GetPostalCode() {
    return PostalCode;
  }


  public void setPostalCode(@.annotation.Nonnull string PostalCode) {
    this.PostalCode = PostalCode;
  }

  public ServiceE911LocationInput State(@.annotation.Nonnull string State) {
    
    this.State = State;
    return this;
  }

  /**
   * Get State
   * @return State
   */
  @.annotation.Nonnull

  public string GetState() {
    return State;
  }


  public void setState(@.annotation.Nonnull string State) {
    this.State = State;
  }

  public ServiceE911LocationInput Type(@.annotation.Nullable TYPE Type) {
    
    this.Type = Type;
    return this;
  }

  /**
   * Get Type
   * @return Type
   */
  @.annotation.Nullable

  public TYPE GetType() {
    return Type;
  }


  public void setType(@.annotation.Nullable TYPE Type) {
    this.Type = Type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceE911LocationInput ServiceE911LocationInput = (ServiceE911LocationInput) o;
    return Objects.equals(this.Address1, ServiceE911LocationInput.Address1) &&
        Objects.equals(this.Address2, ServiceE911LocationInput.Address2) &&
        Objects.equals(this.Community, ServiceE911LocationInput.Community) &&
        Objects.equals(this.PlusFour, ServiceE911LocationInput.PlusFour) &&
        Objects.equals(this.PostalCode, ServiceE911LocationInput.PostalCode) &&
        Objects.equals(this.State, ServiceE911LocationInput.State) &&
        Objects.equals(this.Type, ServiceE911LocationInput.Type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Address1, Address2, Community, PlusFour, PostalCode, State, Type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceE911LocationInput {\n");
    sb.append("    Address1: ").append(toIndentedString(Address1)).append("\n");
    sb.append("    Address2: ").append(toIndentedString(Address2)).append("\n");
    sb.append("    Community: ").append(toIndentedString(Community)).append("\n");
    sb.append("    PlusFour: ").append(toIndentedString(PlusFour)).append("\n");
    sb.append("    PostalCode: ").append(toIndentedString(PostalCode)).append("\n");
    sb.append("    State: ").append(toIndentedString(State)).append("\n");
    sb.append("    Type: ").append(toIndentedString(Type)).append("\n");
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

