import React from 'react'
import PropTypes from 'prop-types'

class StockList extends React.Component {
    renderItem(stockItem) {
        return <h1 key={stockItem.stockId}
                   onClick={() => {
                       this.props.onItemClick(stockItem.stockId)
                   }}
        >{stockItem.stockId}+{stockItem.stockCode}</h1>
    }

    render() {
        return (
            <div>
                {this.props.items.map(this.renderItem.bind(this))}
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
