//header component for our app screen

//import libraries
import React from 'react';
import { Text, View } from 'react-native';

//make component
//design note: always create component with same name as file
const Header = (props) => {
    //destructure styles to eliminate styles. repitition
    const { textStyle, viewStyle } = styles;

    //make sure to set prop style!
    return (
        <View style={viewStyle}>
            <Text style={textStyle}>{ props.headerText }</Text>
        </View>
    );
}

//keep styling for component here
const styles = {
    viewStyle: {
        backgroundColor: '#F8F8F8',
        justifyContent: 'center',
        alignItems: 'center',
        height: 60,
        paddingTop: 35,
        shadowColor: '#000',
        shadowOffset: {width: 0, height: 2},
        shadowOpacity: 0.3,
        elevation: 2,
        position: 'relative'
    },
    textStyle: {
        fontSize: 20
    }
};

//make component available to other parts of the app
export default Header;