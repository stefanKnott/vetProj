import React from 'react';
import { View, Text, Image, Linking } from 'react-native';
import Card from './Card';
import CardSection from './CardSection'
import Button from './Button'

const VetDetail = ({vet}) => {
    //TODO: will want to get lat long to calculate nearest vet
    const { formattedName, address, thumbnailImage } = vet;

    const { 
        thumbnailStyle, 
        headerContentStyle, 
        thumbnailContainerStyle,
        headerTextStyle,
        imageStyle
    } = styles; 

    return (
        <Card>
            <CardSection>
                <View style={thumbnailContainerStyle}>
                    <Image
                        style={thumbnailStyle}
                    />
                </View>
                <View style={headerContentStyle}>
                    <Text style={headerTextStyle}>{formattedName}</Text>
                    <Text>{address.line1}</Text>
                </View>
            </CardSection>
            {/* <CardSection>
                {/* TODO: this is where the google maps insert may be 
                <Image 
                    style={imageStyle}
                    source={{ uri: image }} 
                />
            </CardSection> */}
            <CardSection>
                <Button onPress={() => console.log("pressd") }>
                    See Details
                </Button>
            </CardSection>
        </Card>
    )
}

const styles= {
    headerContentStyle: {
        flexDirection: 'column',
        justifyContent: 'space-around',
    },
    headerTextStyle: {
        fontSize: 18
    },
    thumbnailStyle: {
        width: 50,
        height: 50
    },
    thumbnailContainerStyle: {
        justifyContent: 'center',
        alignItems: 'center',
        marginLeft: 10,
        marginRight: 10
    },
    imageStyle: {
        height: 300,
        flex: 1,
        width: null
    }
};

export default VetDetail;