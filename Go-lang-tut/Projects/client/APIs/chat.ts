
import AsyncStorage from "@react-native-async-storage/async-storage";
import axios from "axios";
import { Chat } from "../models/chat";
import { jwtDecode } from "jwt-decode";
import "core-js/stable/atob";

const API_URL = "http://192.168.0.123:8000/";

function getJWTData(token: string) {
    const data = jwtDecode(token);
    return data as { id: string, email: string, exp: number };
}

axios.interceptors.request.use(
    async (config) => {
        const blackListedEndPoints = [
            "login",
            "register",
            "refresh",
        ]
        const url = config.url?.split("/")[config.url?.split("/").length - 1] || "";
        if (blackListedEndPoints.includes(url)) {
            return config;
        }
        const token = await AsyncStorage.getItem("userToken");
        if (token) {
            config.headers["Authorization"] = `${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

async function getGPT(question: string) {
    return await axios.post(API_URL + "getResponse", {
        "query": question
    });
}

async function createChat(chat: Chat) {
    const token = await AsyncStorage.getItem("userToken");
    if (token) {
        const data = getJWTData(token);
        return await axios.post(API_URL + "createChat", {
            ...chat,
            userId: data.id
        });
    }
}

async function getAllChats(): Promise<Chat[]> {
    const token = await AsyncStorage.getItem("userToken");
    if (token) {
        const data = getJWTData(token);
        const res = await axios.get(API_URL + "getAllChats/" + data.id);
        if (res.data) {
            return (res.data as any[]).map((chat) => new Chat(chat));
        }
    }
    return [];
}

export {
    getGPT,
    createChat,
    getAllChats
}


