import express from "express";
import helmet from "helmet";
import cors from "cors";
import morgan from "morgan";
import compress from "compression";

import conf from './config'
import v1route from './route/v1'



const app = express();

app.use(morgan(conf.morgan)); //log to console on development
app.use(express.urlencoded({extended: true}));
app.use(express.json());

app.use(compress());
app.use(helmet());
app.use(cors())

app.use('/v1', v1route)


export {
    app
}