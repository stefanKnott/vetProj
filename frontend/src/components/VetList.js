import React, { Component } from 'react';
import { ScrollView } from 'react-native';
import axios from "axios";
import VetDetail from './VetDetail'

//props in parent->child, state for local

class VetList extends Component {
    //define initial state of our AlbumList component
    state = {vets: [], coords: {}, error: ''};

    componentWillMount(){
        //http get req is a promise..use .then() to handle result of promise
        axios.get('http://localhost:8000/getAll')
            .then(response => this.setState({vets: response.data}));

        navigator.geolocation.getCurrentPosition(
            (position) => {
                this.setState({
                    coords: position.coords,
                    error: null,
                });
            },
            (error) => this.setState({ error: error.message }),
            {enableHighAccuracy: true, timeout: 20000, maximumAge: 1000},
        );
    }

    //create text jsx for each album
    renderVets() {
        return this.state.vets.map(vet => 
            <VetDetail vet={vet} />
        );
    }

    render() {
        return(
            <ScrollView>
                {this.renderVets()}
            </ScrollView>
        );
    }
}

export default VetList;