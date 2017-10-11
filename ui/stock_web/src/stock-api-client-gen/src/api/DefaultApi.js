/**
 * Stock
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import Stock from '../model/Stock';

/**
* Default service.
* @module api/DefaultApi
* @version v1
*/
export default class DefaultApi {

    /**
    * Constructs a new DefaultApi. 
    * @alias module:api/DefaultApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the stocksGet operation.
     * @callback module:api/DefaultApi~stocksGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Stock} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get stock,sh_{code},sz_{code}
     * @param {String} stockId Stock id
     * @param {module:api/DefaultApi~stocksGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Stock}
     */
    stocksGet(stockId, callback) {
      let postBody = null;

      // verify the required parameter 'stockId' is set
      if (stockId === undefined || stockId === null) {
        throw new Error("Missing the required parameter 'stockId' when calling stocksGet");
      }


      let pathParams = {
        'stockId': stockId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json;charset=utf-8'];
      let accepts = ['application/json;charset=utf-8'];
      let returnType = Stock;

      return this.apiClient.callApi(
        '/stocks/{stockId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the stocksList operation.
     * @callback module:api/DefaultApi~stocksListCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Stock>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get stock list
     * @param {Object} opts Optional parameters
     * @param {String} opts.exchangeId Exchange id.eg sz,sh...
     * @param {module:api/DefaultApi~stocksListCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Stock>}
     */
    stocksList(opts, callback) {
      opts = opts || {};
      let postBody = null;


      let pathParams = {
      };
      let queryParams = {
        'exchangeId': opts['exchangeId']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json;charset=utf-8'];
      let accepts = ['application/json;charset=utf-8'];
      let returnType = [Stock];

      return this.apiClient.callApi(
        '/stocks', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }


}
