import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

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

export function register(user) {
	return axios.post(`${API_BASE_URL}/auth/register`, user);
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

