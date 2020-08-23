import axios from "axios";

export default axios.create({
    baseURL: "http://localhost:9000/animal?id=3",
    responseType: "json"
});
