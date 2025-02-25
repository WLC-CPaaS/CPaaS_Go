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
    instance = new WhiteLabelCommunicationsCPaasApiDocumentation.CallQueueMembershipApi();
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

  describe('CallQueueMembershipApi', function() {
    describe('v1AccountAccountIDQueuemembershipPost', function() {
      it('should call v1AccountAccountIDQueuemembershipPost successfully', function(done) {
        //uncomment below and update the code to test v1AccountAccountIDQueuemembershipPost
        //instance.v1AccountAccountIDQueuemembershipPost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('v1AccountAccountIDQueuemembershipRecipientIDDisablePost', function() {
      it('should call v1AccountAccountIDQueuemembershipRecipientIDDisablePost successfully', function(done) {
        //uncomment below and update the code to test v1AccountAccountIDQueuemembershipRecipientIDDisablePost
        //instance.v1AccountAccountIDQueuemembershipRecipientIDDisablePost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('v1AccountAccountIDQueuemembershipRecipientIDEnablePost', function() {
      it('should call v1AccountAccountIDQueuemembershipRecipientIDEnablePost successfully', function(done) {
        //uncomment below and update the code to test v1AccountAccountIDQueuemembershipRecipientIDEnablePost
        //instance.v1AccountAccountIDQueuemembershipRecipientIDEnablePost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
