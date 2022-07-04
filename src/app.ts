import path from "path";
import express from "express";
import cookieParser from "cookie-parser";
import csrf from "csurf";
import helmet from "helmet";
import cors from "cors";
import morgan from "morgan";
import compress from "compression";
import mongoSanitize from "express-mongo-sanitize";
import httpStatus from "http-status";
import session from "express-session";
import connectredis from "connect-redis"
import {createClient} from "redis";

import conf from "./config";
import v1route from "./route/v1";
import ckroute from "./route/cookie";
import sessroute from "./route/session";
import { APIError, errorConverter, errorHandler } from "./middleware/error";


const app = express();

app.set("views", path.join(__dirname, "views"));
app.set("view engine", "ejs");
app.use(express.static(path.join(__dirname, 'public')));

app.use(morgan(conf.morgan)); //log to console on development
app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.use(cookieParser());
const csrfCookie = csrf({ cookie: {key: '_csrf', maxAge: 60 * 10} }) // send _csrf cookie to client to use as seed

const csrfSess = csrf();
const RedisStore = connectredis(session)
let redisClient = createClient({ legacyMode: true })
redisClient.connect().catch(console.error)
const sessOpt: session.SessionOptions = {
    store: new RedisStore({ client: redisClient }),
    saveUninitialized: false,
    secret: conf.secret as string,
    resave: false,
    cookie: {}
}
if (app.get('env') === 'production') {
    app.set('trust proxy', 1);
    if (sessOpt.cookie) {
        sessOpt.cookie.maxAge = 1000 * 60 * 60 * 2;
        sessOpt.cookie.secure = true;
    }
    const whiteList = ['http://localhost:3333', /\.google\.com$/];
    app.use(cors({origin: whiteList}));
} else {
    app.use(cors());
}
app.use(session(sessOpt))
app.locals.redis = redisClient;


app.use(mongoSanitize());
app.use(helmet());
app.use(compress());

app.use("/v1", v1route);
app.use("/cookie", csrfCookie, ckroute);
app.use("/session", csrfSess, (req, res, next) => {
    res.cookie('XSRF-TOKEN', req.csrfToken(), {httpOnly: false}); // for SPA
    res.locals.csrf = req.csrfToken();
    res.locals.sessid = req.sessionID;
    next();
},sessroute);

// send back a 404 error for any unknown api request
app.use((req, res, next) => {
    next(new APIError(httpStatus.NOT_FOUND, "Not found"));
});
// convert error to ApiError, if needed
app.use(errorConverter);
app.use(errorHandler);

export { app };
