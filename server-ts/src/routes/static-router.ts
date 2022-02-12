import Router from 'koa-router';


// Constants
const router = new Router();


/**
 * Serve html.
 */
router.get('', (ctx) => {
    ctx.type = 'html';
    ctx.body = '<div>Hello, you are running Koa in development mode.</div>';
});


// Export default
export default router;
