import jsonwebtoken, { JwtPayload } from 'jsonwebtoken'; 
import randomstring from 'randomstring';

import userDao from '@daos/user-dao';
import pwdUtil from '@util/pwd-util';


interface ILoginResp {
    passed: boolean;
    jwt?: string;
    error?: string;
}


/**
 * Check user creds and return a jwt if they passed.
 * 
 * @param email 
 * @param password 
 * @returns 
 */
async function login(email: string, password: string): Promise<ILoginResp> {
    // Fetch user
    const user = await userDao.findByEmail(email);
    if (!user) {
        return {passed: false};
    }
    // Fetch password-hash
    const pwdHash = await userDao.fetchPwdHash(user.id);
    // Check password
    const pwdPassed = await pwdUtil.verify(password, pwdHash);
    if (!pwdPassed) {
        return {passed: false};
    }
    // Create the jwt
    const jwt = await sign({
        id: user.id,
        email: user.email,
        name: user.name,
    });
    // Return
    return {passed: true, jwt};
}


/**
 * Encrypt data and return jwt.
 *
 * @param data
 */
function sign(data: JwtPayload): Promise<string> {
    // Setup secret and options
    const secret = (process.env.JWT_SECRET ?? randomstring.generate(100)),
        options = {expiresIn: process.env.COOKIE_EXP};
    // Return promise
    return new Promise((resolve, reject) => {
        return jsonwebtoken.sign(data, secret, options, (err, token) => {
            return err ? reject(err) : resolve(token ?? '');
        });
    });
}


// Export default
export default {
    login,
} as const;
