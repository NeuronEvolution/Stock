import React from 'react';
import ReactDOM from 'react-dom';
import {createStore,applyMiddleware} from 'redux'
import thunk from 'redux-thunk'
import './index.css';
import Root from './containers/Root';
import registerServiceWorker from './registerServiceWorker';
import rootReducer from "./reducers/reducers";
import swagger from './middleware/swagger'
import logger from 'redux-logger'

let store=createStore(rootReducer,{},applyMiddleware(thunk,swagger,logger))

ReactDOM.render(<Root store={store}/>, document.getElementById('root'));
registerServiceWorker();
