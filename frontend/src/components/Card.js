import React from 'react';
import { View } from 'react-native';

const Card = (props) =>{
    return(
        <View style={styles.containerStyle}>
            {props.children}
        </View>
    );
};

const styles = {
    containerStyle: {
        borderWidth: 1,
        borderRadius: 2,
        borderColor: '#ddd',
        borderBottomWidth: 0,
        shadowColor: '#000',
        shadowOffset: {width: 0, height: 2},
        shadowOpacity: 0.2,
        shadowRadius: 2,
        marginLeft: 5,
        marginRight: 5,
        marginBottom: 350,
        // position: 'relative',
        // flexDirection: 'column',
        // justifyContent: 'space-around'
    }
};

export default Card;