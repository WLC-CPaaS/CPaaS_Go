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
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
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

  describe('ServiceStoragePlanDatabaseTypes', function() {
    it('should create an instance of ServiceStoragePlanDatabaseTypes', function() {
      // uncomment below and update the code to test ServiceStoragePlanDatabaseTypes
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be.a(WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes);
    });

    it('should have the property callRecording (base name: "call_recording")', function() {
      // uncomment below and update the code to test the property callRecording
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be();
    });

    it('should have the property fax (base name: "fax")', function() {
      // uncomment below and update the code to test the property fax
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be();
    });

    it('should have the property _function (base name: "function")', function() {
      // uncomment below and update the code to test the property _function
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be();
    });

    it('should have the property mailboxMessage (base name: "mailbox_message")', function() {
      // uncomment below and update the code to test the property mailboxMessage
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be();
    });

    it('should have the property media (base name: "media")', function() {
      // uncomment below and update the code to test the property media
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceStoragePlanDatabaseTypes();
      //expect(instance).to.be();
    });

  });

}));
