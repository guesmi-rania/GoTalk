import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL || "https://gotalk-ykyu.onrender.com";

export const healthCheck = async () => {
  const res = await axios.get(`${API_URL}/health`);
  return res.data;
};
