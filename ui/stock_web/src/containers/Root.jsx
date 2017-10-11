import React from 'react'
import { Provider } from 'react-redux'
import {BrowserRouter,Route} from 'react-router-dom'
import PropTypes from 'prop-types'
import App from './App'
import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles';

const theme = createMuiTheme();

class Root extends React.Component {
    render() {
        return (
            <Provider store={this.props.store}>
                <MuiThemeProvider theme={theme}>
                    <BrowserRouter>
                        <Route path='/' component={App}/>
                    </BrowserRouter>
                </MuiThemeProvider>
            </Provider>
        )
    }
}

Root.propTypes={
    store:PropTypes.object.isRequired
}

export default Root