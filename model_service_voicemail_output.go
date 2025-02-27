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
 * ServiceVoicemailOutput
 */
@.annotation.Generated(value = "org.openapitools.codegen.languages.GoClientCodegen", comments = "Generator version: 7.11.0-SNAPSHOT")
public class ServiceVoicemailOutput {
  @.annotation.Nullable
  private string Id;

  @.annotation.Nullable
  private string Mailbox;

  @.annotation.Nullable
  private ServiceVoicemailMedia Media;

  @.annotation.Nullable
  private string MediaExtension;

  @.annotation.Nullable
  private int32 Messages;

  @.annotation.Nullable
  private string Name;

  @.annotation.Nullable
  private string OwnerId;

  @.annotation.Nullable
  private string Pin;

  @.annotation.Nullable
  private bool RequirePin;

  @.annotation.Nullable
  private string Timezone;

  public ServiceVoicemailOutput() {
  }

  public ServiceVoicemailOutput Id(@.annotation.Nullable string Id) {
    
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

  public ServiceVoicemailOutput Mailbox(@.annotation.Nullable string Mailbox) {
    
    this.Mailbox = Mailbox;
    return this;
  }

  /**
   * Get Mailbox
   * @return Mailbox
   */
  @.annotation.Nullable

  public string GetMailbox() {
    return Mailbox;
  }


  public void setMailbox(@.annotation.Nullable string Mailbox) {
    this.Mailbox = Mailbox;
  }

  public ServiceVoicemailOutput Media(@.annotation.Nullable ServiceVoicemailMedia Media) {
    
    this.Media = Media;
    return this;
  }

  /**
   * Get Media
   * @return Media
   */
  @.annotation.Nullable

  public ServiceVoicemailMedia GetMedia() {
    return Media;
  }


  public void setMedia(@.annotation.Nullable ServiceVoicemailMedia Media) {
    this.Media = Media;
  }

  public ServiceVoicemailOutput MediaExtension(@.annotation.Nullable string MediaExtension) {
    
    this.MediaExtension = MediaExtension;
    return this;
  }

  /**
   * Get MediaExtension
   * @return MediaExtension
   */
  @.annotation.Nullable

  public string GetMediaExtension() {
    return MediaExtension;
  }


  public void setMediaExtension(@.annotation.Nullable string MediaExtension) {
    this.MediaExtension = MediaExtension;
  }

  public ServiceVoicemailOutput Messages(@.annotation.Nullable int32 Messages) {
    
    this.Messages = Messages;
    return this;
  }

  /**
   * Get Messages
   * @return Messages
   */
  @.annotation.Nullable

  public int32 GetMessages() {
    return Messages;
  }


  public void setMessages(@.annotation.Nullable int32 Messages) {
    this.Messages = Messages;
  }

  public ServiceVoicemailOutput Name(@.annotation.Nullable string Name) {
    
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

  public ServiceVoicemailOutput OwnerId(@.annotation.Nullable string OwnerId) {
    
    this.OwnerId = OwnerId;
    return this;
  }

  /**
   * Get OwnerId
   * @return OwnerId
   */
  @.annotation.Nullable

  public string GetOwnerId() {
    return OwnerId;
  }


  public void setOwnerId(@.annotation.Nullable string OwnerId) {
    this.OwnerId = OwnerId;
  }

  public ServiceVoicemailOutput Pin(@.annotation.Nullable string Pin) {
    
    this.Pin = Pin;
    return this;
  }

  /**
   * Get Pin
   * @return Pin
   */
  @.annotation.Nullable

  public string GetPin() {
    return Pin;
  }


  public void setPin(@.annotation.Nullable string Pin) {
    this.Pin = Pin;
  }

  public ServiceVoicemailOutput RequirePin(@.annotation.Nullable bool RequirePin) {
    
    this.RequirePin = RequirePin;
    return this;
  }

  /**
   * Get RequirePin
   * @return RequirePin
   */
  @.annotation.Nullable

  public bool getRequirePin() {
    return RequirePin;
  }


  public void setRequirePin(@.annotation.Nullable bool RequirePin) {
    this.RequirePin = RequirePin;
  }

  public ServiceVoicemailOutput Timezone(@.annotation.Nullable string Timezone) {
    
    this.Timezone = Timezone;
    return this;
  }

  /**
   * Get Timezone
   * @return Timezone
   */
  @.annotation.Nullable

  public string GetTimezone() {
    return Timezone;
  }


  public void setTimezone(@.annotation.Nullable string Timezone) {
    this.Timezone = Timezone;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServiceVoicemailOutput ServiceVoicemailOutput = (ServiceVoicemailOutput) o;
    return Objects.equals(this.Id, ServiceVoicemailOutput.Id) &&
        Objects.equals(this.Mailbox, ServiceVoicemailOutput.Mailbox) &&
        Objects.equals(this.Media, ServiceVoicemailOutput.Media) &&
        Objects.equals(this.MediaExtension, ServiceVoicemailOutput.MediaExtension) &&
        Objects.equals(this.Messages, ServiceVoicemailOutput.Messages) &&
        Objects.equals(this.Name, ServiceVoicemailOutput.Name) &&
        Objects.equals(this.OwnerId, ServiceVoicemailOutput.OwnerId) &&
        Objects.equals(this.Pin, ServiceVoicemailOutput.Pin) &&
        Objects.equals(this.RequirePin, ServiceVoicemailOutput.RequirePin) &&
        Objects.equals(this.Timezone, ServiceVoicemailOutput.Timezone);
  }

  @Override
  public int hashCode() {
    return Objects.hash(Id, Mailbox, Media, MediaExtension, Messages, Name, OwnerId, Pin, RequirePin, Timezone);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServiceVoicemailOutput {\n");
    sb.append("    Id: ").append(toIndentedString(Id)).append("\n");
    sb.append("    Mailbox: ").append(toIndentedString(Mailbox)).append("\n");
    sb.append("    Media: ").append(toIndentedString(Media)).append("\n");
    sb.append("    MediaExtension: ").append(toIndentedString(MediaExtension)).append("\n");
    sb.append("    Messages: ").append(toIndentedString(Messages)).append("\n");
    sb.append("    Name: ").append(toIndentedString(Name)).append("\n");
    sb.append("    OwnerId: ").append(toIndentedString(OwnerId)).append("\n");
    sb.append("    Pin: ").append(toIndentedString(Pin)).append("\n");
    sb.append("    RequirePin: ").append(toIndentedString(RequirePin)).append("\n");
    sb.append("    Timezone: ").append(toIndentedString(Timezone)).append("\n");
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

