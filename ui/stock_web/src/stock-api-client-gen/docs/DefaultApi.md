# Stock.DefaultApi

All URIs are relative to *https://localhost/api/stock/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**stocksGet**](DefaultApi.md#stocksGet) | **GET** /stocks/{stockId} | Get stock,sh_{code},sz_{code}
[**stocksList**](DefaultApi.md#stocksList) | **GET** /stocks | Get stock list


<a name="stocksGet"></a>
# **stocksGet**
> Stock stocksGet(stockId)

Get stock,sh_{code},sz_{code}

### Example
```javascript
import Stock from 'stock';

let apiInstance = new Stock.DefaultApi();

let stockId = "stockId_example"; // String | Stock id


apiInstance.stocksGet(stockId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockId** | **String**| Stock id | 

### Return type

[**Stock**](Stock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

<a name="stocksList"></a>
# **stocksList**
> [Stock] stocksList(opts)

Get stock list

### Example
```javascript
import Stock from 'stock';

let apiInstance = new Stock.DefaultApi();

let opts = { 
  'exchangeId': "exchangeId_example" // String | Exchange id.eg sz,sh...
};

apiInstance.stocksList(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeId** | **String**| Exchange id.eg sz,sh... | [optional] 

### Return type

[**[Stock]**](Stock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

