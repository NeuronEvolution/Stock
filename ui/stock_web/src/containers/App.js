import React, { Component } from 'react';
import './App.css';
import { connect } from 'react-redux';
import StockList from "../components/StockList";
import {apiStockList,onStockItemClick} from '../actions/actions'

class App extends Component {
    loadStockList() {
        this.props.apiStockList("szA")
    }

    componentWillMount() {
        this.loadStockList()

        setInterval(() => {
            //this.props.apiStockList("szA")
        }, 10000)
    }

    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <h1 className="App-title">助理巴菲特</h1>
                </header>
                <div>
                    <StockList
                        items={this.props.stockListItems}
                        onItemClick={(stockId) => {
                            this.props.onStockItemClick(stockId)
                        }}
                        onLoadMore={() => {
                            console.log("onLoadMore")
                        }}
                    />
                </div>
            </div>
        );
    }
}

function connectProps(state) {
    return state
}

export default connect(connectProps,{apiStockList,onStockItemClick})(App);
