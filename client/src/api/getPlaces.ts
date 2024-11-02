import axios from "axios";

export async function getPlacesMapBox(query: string) {
    try {
        const response = await axios.get(
            `https://api.mapbox.com/geocoding/v5/mapbox.places/${query}.json`,
            {
                params: {
                    access_token: import.meta.env.VITE_MAPBOX_TOKEN,
                    country: 'tn'
                },
            }
        );

        return response.data.features;
    } catch (error) {
        throw new Error("There was an error while fetching places")
    }
}

export async function getPlacesGoogle(query: string) {
    const response = await axios.get(
      `http://localhost:3001/api/places`,
      {
          params: {
              input: query,
          }
      }
    );
    return response.data.predictions;
}

export async function getPlaceDetails(placeID: string){
    const response = await axios.get(
        `http://localhost:3001/api/placeDetails`,
        {
            params: {
                place_id: placeID
            }
        }
    )
    return response.data.result
}

export async function getAddressFromLatLng(latitude: string, longitude: string){
    const response = await axios.get(
        `http://localhost:3001/api/address/lookup`,
        {
            params: {
                latitude: latitude,
                longitude: longitude
            }
        }
    )
    return response.data
}

