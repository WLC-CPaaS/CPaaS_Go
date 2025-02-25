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
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource();
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

  describe('ServiceCallRecordingSource', function() {
    it('should create an instance of ServiceCallRecordingSource', function() {
      // uncomment below and update the code to test ServiceCallRecordingSource
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource();
      //expect(instance).to.be.a(WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource);
    });

    it('should have the property any (base name: "any")', function() {
      // uncomment below and update the code to test the property any
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource();
      //expect(instance).to.be();
    });

    it('should have the property offnet (base name: "offnet")', function() {
      // uncomment below and update the code to test the property offnet
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource();
      //expect(instance).to.be();
    });

    it('should have the property onnet (base name: "onnet")', function() {
      // uncomment below and update the code to test the property onnet
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceCallRecordingSource();
      //expect(instance).to.be();
    });

  });

}));
