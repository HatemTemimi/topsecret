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
      `http://localhost:3001/api/google-places`,
      {
          params: {
              input: query,
          }
      }
    );
    return response.data.predictions;
}

