// lib/api.js
import axios from 'axios';

const api = axios.create({
  baseURL: 'https://jsonplaceholder.typicode.com/posts', // Replace with your API base URL
});

export const fetchData = async () => {
  try {
    const response = await api.get(); // Replace with your API endpoint
    return response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;
  }
};