import jwt from 'koa-jwt';
import randomstring from 'randomstring';


/**
 * Parse jsonwebtoken, token data is available at "ctx.state.user".
 * 
 * @returns 
 */
export function getApiMw() {
    return jwt({
        secret: (process.env.JWT_SECRET ?? randomstring.generate(100)),
        cookie: process.env.COOKIE_NAME,
    }).unless({path: [/^\/api\/auth/]});
}


/**
 * Parse jsonwebtoken, token data is available at "ctx.state.user".
 * 
 * @returns 
 */
export function getSessionMw() {
    return jwt({
        secret: (process.env.JWT_SECRET ?? randomstring.generate(100)),
        cookie: process.env.COOKIE_NAME,
        passthrough: true,
    });
}
