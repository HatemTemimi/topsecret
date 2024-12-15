import axios from "./axios"

export async function getRentals(){
try {
    const response = await axios.get('http://localhost:3001/api/rental/list')
    return response.data
  } catch (error) {
    console.error("Failed to fetch rentals:", error)
  }
}

export async function getRentalById(id:string){
try {
    const response = await axios.get(`http://localhost:3001/api/rental/${id}`)
    return response.data
  } catch (error) {
    console.error("Failed to fetch rental with id:",id, error)
  }
}

export async function getRentalsByUserId(id:string){
try {
    const response = await axios.get(`http://localhost:3001/api/rental/user/${id}`)
    return response.data
  } catch (error) {
    console.error("Failed to fetch rental with id:",id, error)
  }
}

export async function addRental(data){
  const response = await axios.post("http://localhost:3001/api/rental/add", data, {
    headers: { "Content-Type": "multipart/form-data" },
});
return response
}

export async function updateRental(id, data){
  const response = await axios.put(`http://localhost:3001/api/rental/${id}`, data);
  return response
}

export async function deleteRental(id){
  const response = await axios.delete(`http://localhost:3001/api/rental/${id}`);
  return response
}
