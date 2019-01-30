import React, { Component } from 'react';
import { ScrollView } from 'react-native';
import axios from "axios";
import VetDetail from './VetDetail'

//props in parent->child, state for local
class VetList extends Component {
    //define initial state of our AlbumList component
    state = {vets: [], coords: {}, error: ''};

    //future implementation for offline situation
    // orderByProximity(vets){
    //     R = 3959
    //     //will need to iterate thru vets array and determine distance from
    //     for (var i = 0; i < vets.length; i++){
    //         //calculate distance using Haversine
    //         lat1 = this.state.coords.latitude
    //         lat2 = vets[i].location.latitude
    //         lon1 = this.state.coords.longitude
    //         lon2 = vets[i].location.longitude

    //         var R = 6371e3; // metres
    //         var φ1 = lat1.toRadians();
    //         var φ2 = lat2.toRadians();
    //         var Δφ = (lat2-lat1).toRadians();
    //         var Δλ = (lon2-lon1).toRadians();

    //         var a = Math.sin(Δφ/2) * Math.sin(Δφ/2) +
    //                 Math.cos(φ1) * Math.cos(φ2) *
    //                 Math.sin(Δλ/2) * Math.sin(Δλ/2);
    //         var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));

    //         vets[i].distance = R * c;
    //     }

    //     return vets
    // }

    componentWillMount(){
        navigator.geolocation.getCurrentPosition(
            (position) => {
                //for now hardcode denver lat long -- NOTE THIS ALL DOES WORK
                axios.get('http://localhost:8000/getAll', {params: {latitude: 39.7392, longitude: -104.9903}} )
                .then(response => {
                    // vetData = this.orderByProximity(response.data)
                    this.setState({vets: response.data, coords: {latitude: 39.7392, longitude: -104.9903}})
                });
            },
            (error) => this.setState({ error: error.message }),
            {enableHighAccuracy: true, timeout: 20000, maximumAge: 1000},
        );
       
    }

    //create text jsx for each album
    renderVets() {
        return this.state.vets.map(vet => 
            <VetDetail vet={vet} coords={this.state.coords} />
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