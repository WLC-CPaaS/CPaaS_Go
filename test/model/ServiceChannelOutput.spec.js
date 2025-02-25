/**
 * White Label Communications CPaas API Documentation
 * A CPaaS platform API
 *
 * The version of the OpenAPI document: 1.1
 * Contact: support@whitelabelcomm.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.WhiteLabelCommunicationsCPaasApiDocumentation);
  }
}(this, function(expect, WhiteLabelCommunicationsCPaasApiDocumentation) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ServiceChannelOutput', function() {
    it('should create an instance of ServiceChannelOutput', function() {
      // uncomment below and update the code to test ServiceChannelOutput
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be.a(WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput);
    });

    it('should have the property answered (base name: "answered")', function() {
      // uncomment below and update the code to test the property answered
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property authorizingId (base name: "authorizing_id")', function() {
      // uncomment below and update the code to test the property authorizingId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property authorizingType (base name: "authorizing_type")', function() {
      // uncomment below and update the code to test the property authorizingType
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property callflowId (base name: "callflow_id")', function() {
      // uncomment below and update the code to test the property callflowId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property channelAuthorized (base name: "channel_authorized")', function() {
      // uncomment below and update the code to test the property channelAuthorized
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property customApplicationVars (base name: "custom_application_vars")', function() {
      // uncomment below and update the code to test the property customApplicationVars
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property customAuthHeaders (base name: "custom_auth_headers")', function() {
      // uncomment below and update the code to test the property customAuthHeaders
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property customChannelVars (base name: "custom_channel_vars")', function() {
      // uncomment below and update the code to test the property customChannelVars
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property customSipHeaders (base name: "custom_sip_headers")', function() {
      // uncomment below and update the code to test the property customSipHeaders
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property destination (base name: "destination")', function() {
      // uncomment below and update the code to test the property destination
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property direction (base name: "direction")', function() {
      // uncomment below and update the code to test the property direction
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property elapsedS (base name: "elapsed_s")', function() {
      // uncomment below and update the code to test the property elapsedS
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property fromTag (base name: "from_tag")', function() {
      // uncomment below and update the code to test the property fromTag
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property interactionId (base name: "interaction_id")', function() {
      // uncomment below and update the code to test the property interactionId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property isLoopback (base name: "is_loopback")', function() {
      // uncomment below and update the code to test the property isLoopback
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property isOnhold (base name: "is_onhold")', function() {
      // uncomment below and update the code to test the property isOnhold
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property otherLeg (base name: "other_leg")', function() {
      // uncomment below and update the code to test the property otherLeg
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property ownerId (base name: "owner_id")', function() {
      // uncomment below and update the code to test the property ownerId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property presenceId (base name: "presence_id")', function() {
      // uncomment below and update the code to test the property presenceId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property request (base name: "request")', function() {
      // uncomment below and update the code to test the property request
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property resellerId (base name: "reseller_id")', function() {
      // uncomment below and update the code to test the property resellerId
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property timestamp (base name: "timestamp")', function() {
      // uncomment below and update the code to test the property timestamp
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property toTag (base name: "to_tag")', function() {
      // uncomment below and update the code to test the property toTag
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property username (base name: "username")', function() {
      // uncomment below and update the code to test the property username
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

    it('should have the property uuid (base name: "uuid")', function() {
      // uncomment below and update the code to test the property uuid
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceChannelOutput();
      //expect(instance).to.be();
    });

  });

}));
