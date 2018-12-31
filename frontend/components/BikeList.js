import React from 'react'

class BikeList extends React.Component {
    render() {
        const orderedBikes = [...this.props.bikes].sort((a, b) => a.id > b.id)
        return (
            <div className="bikes-list">
                <ul>
                    {orderedBikes.map(bike => {
                        const active = bike.id === this.props.bikeId ? 'active' : '';
                        return (
                            <li key={bike.id} className={"bike " + active}>
                                <a
                                    onClick={() => this.props.viewBike(bike.id)}
                                    href="#">
                                    {bike.make + " " + bike.model}
                                </a>
                            </li>
                        )
                    })}

                </ul>
            </div>
        )
    }
}

export default BikeList