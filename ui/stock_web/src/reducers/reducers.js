import { combineReducers } from 'redux'
import {STOCK_LIST_REQUEST,STOCK_LIST_SUCCESS,STOCK_LIST_FAILED} from "../actions/actions";

function onStockListCall(stockListItems=[],action) {
    switch (action.type) {
        case STOCK_LIST_REQUEST:
            return stockListItems
        case STOCK_LIST_SUCCESS:
            return action.payload;
        case STOCK_LIST_FAILED:
            return [];
        default:
            return stockListItems
    }
}

const rootReducer=combineReducers({
    stockListItems: onStockListCall,
})

export default rootReducer