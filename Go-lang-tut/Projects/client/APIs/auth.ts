import axios from "axios";
import { SigninUser, SignupUser } from "../models/user";
import AsyncStorage from "@react-native-async-storage/async-storage";

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



const API_URL = "http://192.168.0.123:8000/";
async function register(signupUser: SignupUser) {
    return await axios.post(API_URL + "register", signupUser);
}

async function login(signinUser: SigninUser) {
    return await axios.post(API_URL + "login", signinUser);
}




export {
    register,
    login
}