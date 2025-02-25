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
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
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

  describe('ServiceTemporalRuleOutputFull', function() {
    it('should create an instance of ServiceTemporalRuleOutputFull', function() {
      // uncomment below and update the code to test ServiceTemporalRuleOutputFull
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be.a(WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull);
    });

    it('should have the property cycle (base name: "cycle")', function() {
      // uncomment below and update the code to test the property cycle
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property days (base name: "days")', function() {
      // uncomment below and update the code to test the property days
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property interval (base name: "interval")', function() {
      // uncomment below and update the code to test the property interval
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property month (base name: "month")', function() {
      // uncomment below and update the code to test the property month
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property ordinal (base name: "ordinal")', function() {
      // uncomment below and update the code to test the property ordinal
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "start_date")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property timeWindowStart (base name: "time_window_start")', function() {
      // uncomment below and update the code to test the property timeWindowStart
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property timeWindowStop (base name: "time_window_stop")', function() {
      // uncomment below and update the code to test the property timeWindowStop
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

    it('should have the property wdays (base name: "wdays")', function() {
      // uncomment below and update the code to test the property wdays
      //var instance = new WhiteLabelCommunicationsCPaasApiDocumentation.ServiceTemporalRuleOutputFull();
      //expect(instance).to.be();
    });

  });

}));
