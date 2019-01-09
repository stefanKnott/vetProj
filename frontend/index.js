import React from 'react';
import { View, Text, AppRegistry } from 'react-native';
import Header from './src/components/Header';
import VetList from './src/components/VetList';

//this is the first component to be shown
//component: a (React) javascript function that returns some sort of JSX (HTML lookin stuff)
const App = () => {
    return (
        <View style={{ flex: 1 }}>
            <Header headerText={'Vets Near Me'}/>
            <VetList />
        </View>
    );
};

//name must match up with project directory name
AppRegistry.registerComponent('frontend', () => App);