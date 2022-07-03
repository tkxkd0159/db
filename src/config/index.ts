import 'dotenv/config'

const mongoDBname = 'myapp' + (process.env.NODE_ENV !== 'production' ? '-test' : '');
const mongoose = {
    url: `mongodb://${process.env.MONGO_USER}:${process.env.MONGO_PW}@${process.env.MONGO_HOST}/${mongoDBname}`,
    options: {}
}


export default {
    env: process.env.NODE_ENV,
    port: process.env.PORT,
    morgan: process.env.NODE_ENV === 'production' ? 'combined' : 'dev',
    mongoose
}