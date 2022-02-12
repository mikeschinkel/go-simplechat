import userDao from '@daos/user-dao';
import User, { IUser, IUserCreds } from '@models/user-model';
import pwdUtil from '@util/pwd-util';


const errors = {
    addOne: 'User cound not be saved',
} as const;


/**
 * Add one user.
 * 
 * @param email 
 * @param name 
 * @param password 
 */
async function addOne(email: string, name: string, password: string): Promise<void> {
    // Save the user
    const newUser = User.new(email, name);
    const newUserId = await userDao.addOne(newUser);
    if (!newUserId) {
        throw Error(errors.addOne);
    }
    newUser.id = newUserId;
    // Once we have we have the new user's id, insert the password
    const pwdHash = await pwdUtil.encrypt(password);
    const creds: IUserCreds = {
        pwdHash,
        userId: newUser.id,
    };
    await userDao.addCreds(creds);
}


/**
 * Fetch all users.
 * 
 * @returns
 */
function fetchAll(): Promise<IUser[]> {
    return userDao.fetchAll();
}


// Export default
export default {
    addOne,
    fetchAll,
} as const;
