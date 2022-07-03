import express from "express";
import helmet from "helmet";
import cors from "cors";
import morgan from "morgan";
import compress from "compression";
import mongoSanitize from "express-mongo-sanitize";

import conf from './config'
import v1route from './route/v1'
import { APIError, errorConverter, errorHandler } from "./middleware/error";
import httpStatus from "http-status";



const app = express();

app.use(morgan(conf.morgan)); //log to console on development
app.use(express.urlencoded({extended: true}));
app.use(express.json());
app.use(mongoSanitize());

app.use(helmet());
app.use(compress());
app.use(cors())

app.use('/v1', v1route)

// send back a 404 error for any unknown api request
app.use((req, res, next) => {
    next(new APIError(httpStatus.NOT_FOUND, 'Not found'))
});
// convert error to ApiError, if needed
app.use(errorConverter);
app.use(errorHandler);

export {
    app
}