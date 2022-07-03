import fs from "fs/promises";
import path from "path";
import {Client} from "pg";
import {delay} from "../utils";

async function insertSQL(dbname: string, mode?: string) {
    if (mode === "createdb") {
        const initfp = path.join(__dirname, "..", "db/schema/init.db.sql");
        const tmp = await fs.readFile(initfp, { encoding: "utf8" });
        const data = tmp.split(";");
        const initq = [];
        for (let elem of data) {
            initq.push(elem.trim());
        }
        initq.pop();

        const client = new Client({
            user: "ljs",
            password: "secret",
            host: "localhost",
            port: 5444,
            database: "postgres",
        });
        await client.connect();

        for (let i = 0; i < initq.length; i++) {
            if (i === initq.length - 1) {
                client.query(initq[i], (err, res) => {
                    console.log("ERROR during initdb : ", null);
                    client.end();
                });
            } else {
                client.query(initq[i], (err, res) => {
                    console.log("ERROR during initdb : ", null);
                });
            }
        }
    }

    await delay(1000);

    const filepath = path.join(__dirname, "..", "db/schema/init.schema.sql");
    const tmp = await fs.readFile(filepath, { encoding: "utf8" });
    const data = tmp.split(";");
    const q = [];
    for (let elem of data) {
        q.push(elem.trim());
    }
    q.pop();

    const client = new Client({
        user: "ljs",
        password: "secret",
        host: "localhost",
        port: 5444,
        database: dbname,
    });
    await client.connect();

    for (let target of q) {
        client.query(target, (err, res) => {
            console.log("ERROR : ", err);
        });
    }

    setTimeout(() => {
        client.end();
    }, 1000);
}

export default insertSQL;