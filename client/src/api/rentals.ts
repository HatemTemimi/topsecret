import axios from "axios"

export async function getRentals(){
try {
    const response = await axios.get('http://localhost:3001/api/rental/list') // Update with your actual API endpoint
    return response.data
  } catch (error) {
    console.error("Failed to fetch rentals:", error)
  }
}


export async function getRentalById(id:string){
try {
    const response = await axios.get(`http://localhost:3001/api/rental/:${id}`) // Update with your actual API endpoint
    return response.data
  } catch (error) {
    console.error("Failed to fetch rentals with id:",id, error)
  }
}
