const Koa = require('koa');
const Bodyparser = require('koa-bodyparser');
const router = new (require('koa-router'));
const app = new Koa();
const port = 5001;

app.use(Bodyparser());  // add body-parser to koa

const Redis = require('ioredis');
const redis = new Redis({
    host: '172.16.88.145',
    port: 6379,
    keepAlive: true,
});

router.get('/get/:key', async (ctx) => {
    const key = ctx.params.key;
    const value = await redis.get(key);
    ctx.body = {
        key,
        value,
    }
    return ctx.status = 200;
});
router.post('/set', async (ctx) => {
    const { key, value } = ctx.request.body;
    console.log(`set ${key}: ${value}`);
    const result = await redis.set(key, value);
    ctx.body = {
        key,
        value,
        result,
    }
    return ctx.status = 200;
});
app.use(router.routes());

app.listen(port, () => {
    console.log(`listening ${port}`);
});
