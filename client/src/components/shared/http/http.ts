import axios, { Method } from 'axios';


/**
 * Wrapper for the axios function.
 * 
 * @param route
 * @param data 
 * @returns 
 */
async function http(route: Readonly<string[]>, data?: Record<string, any>) {
    const resp = await axios({method: route[0] as Method, url: route[1], data});
    return resp.data;
}


export default http;
