import { Server } from "http";
import mongoose from "mongoose";
import { Client } from "pg";

import conf from "./config";
import { app } from "./app";
import pgdb from "./db";
import redis from "./db/redis"


let server: Server;

async function init() {
    try {
        await mongoose.connect(conf.mongoose.url, conf.mongoose.options);
        console.log(`Connected to MongoDB: ${conf.mongoose.url}`);
        server = app.listen(conf.port, () =>
            console.log(`Listening on ${conf.port}`)
        );
    } catch (err) {
        console.error(err);
    }
}

const unexpectedErrHandler = (err: Error) => {
    console.error(err);
    if (server) {
        server.close(() => {
            console.log("Server closed");
            process.exit(1);
        });
    } else {
        process.exit(1);
    }
};

init();

process.on("uncaughtException", unexpectedErrHandler);
process.on("unhandledRejection", unexpectedErrHandler);
process.on("SIGTERM", () => {
    console.log("SIGTERM triggered");
    if (server) {
        server.close(() => {
            console.log("Server closed");
            process.exit(0);
        });
    }
});

// DBTest()
async function DBTest() {
    const userSchema = new mongoose.Schema({
        name: String,
    });
    const appUsers = mongoose.model("users", userSchema);
    console.log(await appUsers.find());

    // PostgreSQL test
    await pgdb.query(
        "SELECT * FROM cosmo.contacts WHERE name = $1",
        ["ljs"],
        (err, result) => {
            if (err) {
                process.exit(1);
            }
            console.log(result.rows);
        }
    );

    const client = new Client({
        user: "ljs",
        password: "secret",
        host: "localhost",
        port: 5444,
        database: "myapp-test",
    });
    await client.connect();
    client.query("SELECT * FROM cosmo.contacts", (err, res) => {
        console.log("FROM new pg client without pool : ", res.rows);
        client.end();
    });

    // => PosgreSQL Pool & Transaction test (batch processing)
    const poolClient = await pgdb.getClient();
    try {
        await poolClient.query("BEGIN"); // Begin transaction
        const queryText =
            "INSERT INTO cosmo.contacts(name,phone) VALUES($1, $2)";
        await poolClient.query(queryText, ["ljs2", "010-4444-5555"]);
        await poolClient.query(queryText, ["ljs3", "010-6666-7777"]);
        await poolClient.query("COMMIT"); // Finalize our transaction

        poolClient.query("SELECT * FROM cosmo.contacts", (err, res) => {
            console.log("FROM pool client : ", res.rows);
        });
    } catch (e) {
        await poolClient.query("ROLLBACK");
        if (e instanceof Error) {
            console.log(e.stack);
        } else {
            console.error(e);
        }
    } finally {
        poolClient.release();
    }
}

// redisTest();
async function redisTest() {
    const redisC = await redis.getClient();
    await redisC.set("first-key", "first-val");
    const value = await redisC.get('first-key');
    console.log("Get value from Redis : ", value);
    const newRedisC = await redis.getClient();
    let value2 = await newRedisC.get('first-key');
    console.log("Get value from Redis 2nd: ", value2);
    newRedisC.quit();

    // ClientClosedError: The client is closed
    // let value3 = await redisC.get('first-key');
    // console.log("Get value from Redis 3rd: ", value3);

}
