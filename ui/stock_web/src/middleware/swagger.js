import {DefaultApi} from '../stock-api-client-gen/src/index'

export const SWAGGER_API="SWAGGER_API"

export const STOCK_LIST_API='STOCK_LIST_API'

function callApi(apiName,request,callback) {
    let api = new DefaultApi()

    switch (apiName) {
        case STOCK_LIST_API:
            return api.stocksList({exchangeId: request["exchangeId"]}, callback);
        default:
            console.error("unknown api ", apiName)
    }
}

export default store => next => action => {
    const swaggerApi = action[SWAGGER_API]
    if (typeof swaggerApi === 'undefined') {
        return next(action)
    }

    const apiName = swaggerApi["api"]
    if (typeof apiName === 'undefined') {
        return console.error("api undefined")
    }

    const [REQUEST_ACTION, SUCCESS_ACTION, FAILED_ACTION] = swaggerApi.types

    const {request} = swaggerApi

    let callback = (error, data, response) => {
        if (error) {
            return next({type: FAILED_ACTION, error: true, payload: error})
        }

        return next({type: SUCCESS_ACTION, payload: data})
    }

    next({type: REQUEST_ACTION})

    callApi(apiName, request, callback)
}
