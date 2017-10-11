import React from 'react'
import { Provider } from 'react-redux'
import {BrowserRouter,Route} from 'react-router-dom'
import PropTypes from 'prop-types'
import App from './App'

class Root extends React.Component{
    render(){
        return(
            <Provider store={this.props.store}>
                <BrowserRouter>
                    <Route path='/' component={App}/>
                </BrowserRouter>
            </Provider>
        )
    }
}

Root.propTypes={
    store:PropTypes.object.isRequired
}

export default Root