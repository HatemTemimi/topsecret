import axios from "axios";

export default async function getAddress(longitude: string, latitude: string) {
    console.log('received:', longitude, latitude)
    try {
        const response = await axios.get(
            `https://api.mapbox.com/search/geocode/v6/reverse?longitude=${longitude}&latitude=${latitude}`,
            {
                params: {
                    access_token: import.meta.env.VITE_MAPBOX_TOKEN,
                },
            }
        );
        console.log(response)
        return response;
    } catch (error) {
        console.error("There was an error while fetching places:", error);
    }
}