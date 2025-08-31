import api_client from "@/backend/api_client.js";

export function exitApp() {
    api_client.post('/v1/app/exit_app');
}