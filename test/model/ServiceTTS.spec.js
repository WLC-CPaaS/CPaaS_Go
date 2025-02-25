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
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTTS();
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

  describe('ServiceTTS', function() {
    it('should create an instance of ServiceTTS', function() {
      // uncomment below and update the code to test ServiceTTS
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTTS();
      //expect(instance).to.be.a(WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTTS);
    });

    it('should have the property text (base name: "text")', function() {
      // uncomment below and update the code to test the property text
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTTS();
      //expect(instance).to.be();
    });

    it('should have the property voice (base name: "voice")', function() {
      // uncomment below and update the code to test the property voice
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTTS();
      //expect(instance).to.be();
    });

  });

}));
