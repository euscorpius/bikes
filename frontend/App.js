import React from 'react'
import i18n from 'i18next';
import { withI18n, reactI18nextModule } from "react-i18next";

import BikeList from './components/BikeList'

i18n
    .use(reactI18nextModule) // passes i18n down to react-i18next
    .init({
        resources: {
            en: {
                translation: {
                    "bikesHeader": "Bikes"
                }
            },
            fi: {
                translation: {
                    "bikesHeader": "Pyörät"
                }
            }
        },
        lng: "en",
        fallbackLng: "en",

        interpolation: {
            escapeValue: false
        }
    });


const DUMMY_BIKES = [
    {
        bikeId: 1,
        make: "Canyon",
        model: "Aeroad"
    },
    {
        bikeId: 2,
        make: "Scorpion",
        model: "Euscorpius"
    }
]

class App extends React.Component {

    constructor() {
        super()
        this.state = {
            bikeId: null,
            bikes: DUMMY_BIKES
        }
        this.viewBike = this.viewBike.bind(this)
    }

    viewBike(bike) {
        this.setState({ bikeId: bike.id })
    }

    render() {
        const { t } = this.props;
        return (
            <div className="bikes">
                <h1>{t('bikesHeader')}</h1>
                <BikeList
                    viewBike={this.viewBike}
                    bikes={this.state.bikes}
                    bikeId={this.state.bikeId} />
            </div>
        );
    }
}

export default withI18n()(App);