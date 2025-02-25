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

import ApiClient from '../ApiClient';
import ModelMonthlySummaryComponentsData from './ModelMonthlySummaryComponentsData';

/**
 * The ModelMonthlySummaryReport model module.
 * @module model/ModelMonthlySummaryReport
 * @version 1.1
 */
class ModelMonthlySummaryReport {
    /**
     * Constructs a new <code>ModelMonthlySummaryReport</code>.
     * @alias module:model/ModelMonthlySummaryReport
     */
    constructor() { 
        
        ModelMonthlySummaryReport.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ModelMonthlySummaryReport</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ModelMonthlySummaryReport} obj Optional instance to populate.
     * @return {module:model/ModelMonthlySummaryReport} The populated <code>ModelMonthlySummaryReport</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ModelMonthlySummaryReport();

            if (data.hasOwnProperty('components')) {
                obj['components'] = ApiClient.convertToType(data['components'], [ModelMonthlySummaryComponentsData]);
            }
            if (data.hasOwnProperty('transaction_month')) {
                obj['transaction_month'] = ApiClient.convertToType(data['transaction_month'], 'Number');
            }
            if (data.hasOwnProperty('transaction_year')) {
                obj['transaction_year'] = ApiClient.convertToType(data['transaction_year'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ModelMonthlySummaryReport</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ModelMonthlySummaryReport</code>.
     */
    static validateJSON(data) {
        if (data['components']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['components'])) {
                throw new Error("Expected the field `components` to be an array in the JSON data but got " + data['components']);
            }
            // validate the optional field `components` (array)
            for (const item of data['components']) {
                ModelMonthlySummaryComponentsData.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Array.<module:model/ModelMonthlySummaryComponentsData>} components
 */
ModelMonthlySummaryReport.prototype['components'] = undefined;

/**
 * @member {Number} transaction_month
 */
ModelMonthlySummaryReport.prototype['transaction_month'] = undefined;

/**
 * @member {Number} transaction_year
 */
ModelMonthlySummaryReport.prototype['transaction_year'] = undefined;






export default ModelMonthlySummaryReport;

