import axios from "axios";
import {Base64} from 'js-base64';

export const API_BASE_URL = "http://localhost:8080";

function storeToken(token, userData) {
    localStorage.setItem("token", token);
    localStorage.setItem("userData", JSON.stringify(userData));
}

export function getToken() {
    return localStorage.getItem("token");
}

export function isAuthenticated() {
    const token = getToken();

    if (token === null) {
        return false;
    }

    const tokenData = JSON.parse(Base64.decode(token.split('.')[1]));

    const isExpired = Math.round(Date.now() / 1000) > tokenData.exp;

    if (isExpired) {
        localStorage.removeItem("token");
        return false;
    }

    return true;
}

export function extractError(e) {
    if (e && e.response && e.response.data && e.response.data.code && e.response.data.message) {
        return e.response.data;
    } else if (e && e.name && e.message) {
        return {
            id: e.name,
            message: e.message,
            code: 0,
        };
    } else {
        return {
            id: 'unknown-error',
            message: 'An unknown error has occured',
            details: e,
        }
    }
}

async function handleAPIError(err) {
	throw extractError(err);
}

export async function register(user) {
    const result = await axios.post(`${API_BASE_URL}/auth/register`, user);

    if (result.status === 200) {
        storeToken(result.data.jwt, result.data.user);
        return result;
    } else {
        return result;
    }
}

export function logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("userData");
}

export async function loginUsernamePassword(login, password) {
    const result = await axios.post(`${API_BASE_URL}/auth/loginup`, {login, password});

    if (result.status === 200) {
        storeToken(result.data.jwt, result.data.user);
        return result;
    } else {
        return result;
    }
}

export function usersMe() {
    return axios.get(`${API_BASE_URL}/users/me`, {
        headers: {
            "Authorization": `Bearer ${getToken()}`,
        },
    });
}

export function getAdminSchema() {
    return axios.get(`${API_BASE_URL}/crud/_schema`, {
        headers: {
            "Authorization": `Bearer ${getToken()}`,
        }
    }).catch(handleAPIError);
}
