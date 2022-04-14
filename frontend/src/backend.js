import axios from "axios";

const BASE_URL = "http://localhost:5000";

function getCookie(name) {
  const cookies = document.cookie.split(";");
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i].trim();
    if (cookie.startsWith(name + "=")) {
      return cookie.substring(name.length + 1);
    }
  }
  return "";
}

// Helper Functions
async function makeRequest(path, method, data = {}) {
    const headers = {
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Authorization": `Bearer ${getCookie("session")}`
    };

    const req = {
        method: method,
        url: `${BASE_URL}${path}`,
        headers: headers,
        data: data
    };
    
    return await (
        await axios.request(req)
    ).data;
}


// Use Endpoints

export async function GetSelf() {
    return await makeRequest("/api/self", "GET")
} 

