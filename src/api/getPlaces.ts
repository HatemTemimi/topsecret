import axios from "axios";

export default async function getPlaces(query: string) {
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
        console.error("There was an error while fetching places:", error);
    }
}