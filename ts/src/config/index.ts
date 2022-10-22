import * as dotenv from 'dotenv';
import path from 'path';

declare module 'express-session' {
  export interface SessionData {
    user: string;
    jwt: string;
  }
}

dotenv.config({path: path.join(__dirname, '..', '..', '.env.dev')});
// dotenv.config({path: path.join(__dirname, '..', '..', '.env')});

const mongoDBname = 'myapp' + (process.env.NODE_ENV !== 'production' ? '-test' : '');
const mongoose = {
    url: `mongodb://${process.env.MONGO_USER}:${process.env.MONGO_PW}@${process.env.MONGO_HOST}/${mongoDBname}`,
    options: {}
}


export default {
    env: process.env.NODE_ENV,
    port: process.env.PORT,
    secret: process.env.SECRET_KEY,
    morgan: process.env.NODE_ENV === 'production' ? 'combined' : 'dev',
    mongoose
}