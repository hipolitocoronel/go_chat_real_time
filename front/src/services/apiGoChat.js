import axios from "axios";

const baseURL = meta.process.env.VITE_APP_BACK_URL;

const apiGoChat = axios.create({ baseURL });

export default apiGoChat;
