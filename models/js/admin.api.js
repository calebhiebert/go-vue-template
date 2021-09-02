import axios from "axios";
import { getToken } from "./api.js";

export const API_BASE_URL = "http://localhost:8080";

function extractError(e) {
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


