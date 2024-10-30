import axios from "axios";

export default async function getAddress(longitude: string, latitude: string) {
    try {
        const response = await axios.get(
            `https://api.mapbox.com/search/geocode/v6/reverse?longitude=${longitude}&latitude=${latitude}`,
            {
                params: {
                    access_token: import.meta.env.VITE_MAPBOX_TOKEN,
                },
            }
        );
        return response;
    } catch (error) {
        throw new Error("There was an error while fetching places:")
    }
}