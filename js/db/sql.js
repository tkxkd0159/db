import { createPool } from "mysql2";

const pool = createPool(
  process.env.JAWSDB_URL ?? {
    host: "localhost",
    user: "root",
    database: "alco",
    password: "secret",
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0,
  },
);
const promisePool = pool.promise();

export default promisePool;
