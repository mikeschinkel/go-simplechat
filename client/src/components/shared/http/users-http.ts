import http from './http';

// User routes
const prefix = '/api/users',
    routes = {
        add: ['POST', prefix],
        fetchAll: ['GET', prefix],
    } as const;


export interface IUser {
    id: number;
    email: string;
    name: string;
}


/**
 * Login user. Return a boolean letting user know if passed or failed.
 * 
 * @param email
 * @param name
 * @param password 
 * @returns 
 */
function add(email: string, name: string, password: string): Promise<void> {
    return http(routes.add, {email, name, password});
}


/**
 * Fetch all users.
 * 
 * @returns 
 */
function fetchAll(): Promise<{users: IUser[]}> {
    return http(routes.fetchAll)
}


// Export default
export default {
    add,
    fetchAll,
} as const;
