import axios from "axios";

export const API_BASE_URL = "http://localhost:8080";

function storeToken(token, userData) {
	localStorage.setItem("token", token);
	localStorage.setItem("userData", JSON.stringify(userData));
}

function getToken() {
	return localStorage.getItem("token");
}

export function isAuthenticated() {
	return getToken() !== null;
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

