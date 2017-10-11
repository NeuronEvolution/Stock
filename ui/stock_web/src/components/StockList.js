import React from 'react'
import PropTypes from 'prop-types'
import List, { ListItem } from 'material-ui/List';
import "./StockList.css"

class StockList extends React.Component {
    renderItem(stockItem) {
        //return <h1 key={stockItem.stockId}
        //         onClick={() => {
        //           this.props.onItemClick(stockItem.stockId)
        //     }}
        //>{stockItem.stockId}+{stockItem.stockCode}</h1>
        return (
            <ListItem key={stockItem.stockId} divider="true" button="true" onClick={() => {
                this.props.onItemClick(stockItem.stockId)
            }}>
                <label className="StockList-item-code">{stockItem.stockCode}</label>
                <label className="StockList-item-nameCN">{stockItem.stockNameCN}</label>
                <label className="StockList-item-industryName">{stockItem.industryName}</label>
            </ListItem>
        )
    }

    render() {
        return (
            <div>
                <List>
                    {this.props.items.map(this.renderItem.bind(this))}
                </List>
            </div>
        )
    }
}

StockList.propTypes= {
    items: PropTypes.array.isRequired,
    onItemClick: PropTypes.func.isRequired,
    onLoadMore: PropTypes.func.isRequired,
}

export default StockList
