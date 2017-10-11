import {SWAGGER_API,STOCK_LIST_API} from "../middleware/swagger";

export const STOCK_LIST_REQUEST='STOCK_LIST_REQUEST'
export const STOCK_LIST_SUCCESS='STOCK_LIST_SUCCESS'
export const STOCK_LIST_FAILED='STOCK_LIST_FAILED'

export function apiStockList(exchangeId) {
    return {
        [SWAGGER_API]: {
            api: STOCK_LIST_API,
            types: [STOCK_LIST_REQUEST, STOCK_LIST_SUCCESS, STOCK_LIST_FAILED],
            request: {
                exchangeId: exchangeId
            }
        }
    }
}

export const STOCK_LIST_ITEM_CLICK='STOCK_LIST_ITEM_CLICK'

export function onStockItemClick(stockId) {
    return {
        type: STOCK_LIST_ITEM_CLICK,
        payload: {
            stockId: stockId
        }
    }
}