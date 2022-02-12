import bcrypt from 'bcrypt';


const pwdSaltRounds = 12;


/**
 * Hash the password.
 * 
 * @param password 
 * @returns 
 */
function encrypt(password: string): Promise<string> {
    return bcrypt.hash(password, pwdSaltRounds);
}


/**
 * Hash the password synchronously. Useful for testing.
 * 
 * @param password 
 * @returns 
 */
function encryptSync(password: string): string {
    return bcrypt.hashSync(password, pwdSaltRounds);
}


/**
 * See if a password passed.
 * 
 * @param password 
 * @param pwdHash 
 * @returns 
 */
function verify(password: string, pwdHash: string): Promise<boolean> {
    return bcrypt.compare(password, pwdHash);
}


// Export default
export default {
    encrypt,
    encryptSync,
    verify,
};
