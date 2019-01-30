import React from 'react';
import { View, Text, Image, Linking } from 'react-native';
// import { render } from 'react-dom'
import MapView from 'react-native-maps';
import Card from './Card';
import CardSection from './CardSection'
import Button from './Button'

const VetDetail = ({vet, coords}) => {
    //TODO: will want to get lat long to calculate nearest vet
    const { formattedName, address, distance, thumbnailImage, location } = vet;
    const {latitude, longitude} = coords
    const position = [location.latitude, location.longitude]
    const { 
        thumbnailStyle, 
        headerContentStyle, 
        thumbnailContainerStyle,
        headerTextStyle,
        mapContainerStyle,
        mapContentStyle,
        imageStyle
    } = styles; 

    return (
        <Card flexDirection='row' justifyContent='space-between'>
            <CardSection>
                <View style={thumbnailContainerStyle}>
                    <Image
                        style={thumbnailStyle}
                        source={{ uri: "https://imgur.com/tRumzvt.png"}}
                    />
                </View>
                <View style={headerContentStyle}>
                    <Text style={headerTextStyle}>{formattedName}</Text>
                    <Text>{address.line1}</Text>
                    <Text>{address.city}, {address.state}</Text>
                    <Text>{distance} mi</Text>
                </View>
                </CardSection>

            <CardSection>
                <MapView
                    style={styles.mapContentStyle}
                    region={{
                    latitude: location.latitude,
                    longitude: location.longitude,
                    latitudeDelta: 0.0922,
                    longitudeDelta: 0.0421,
                    }}
                >
                    <MapView.Marker
                        coordinate={{
                            latitude: latitude,
                            longitude: longitude}}
                        pinColor={"blue"}
                        title={"Your location"}
                    /> 
                    <MapView.Marker
                        coordinate={{
                            latitude: location.latitude,
                            longitude: location.longitude}}
                        title={"Vet location"}
                    /> 
                </MapView>
            </CardSection>
            {/* <CardSection>
                <Button onPress={() => console.log("pressd") }>
                    Go
                </Button>
            </CardSection> */}
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
    },
    mapContainerStyle: {
        position: 'relative',
        top: 0,
        left: 0,
        bottom: 0,
        right: 0,
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center'
    },
    mapContentStyle: {
        position: 'absolute',
        top: 0,
        left: 0,
        bottom: 0,
        right: 0,
        width:360,
        height: 360
    }
};

export default VetDetail;