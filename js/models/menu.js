import sql from "../db/sql.js";

class Menu {
  constructor(row) {
    this.id = row.menu_id;
    this.bizID = row.fk_business_id;
    this.name = row.menu_name;
    this.kilocalories = row.kilocalories;
    this.price = row.price;
    this.likes = row.likes;
  }
}

async function findByBizID(id) {
  const q = `
  SELECT * FROM menus
  WHERE fk_business_id = ${id}
  `;
  const [rows] = await sql.query(q);
  return rows.map((item) => new Menu(item));
}

async function updateMenuLikes(id, like) {
  return sql.query(`
    UPDATE menus
    SET likes = likes + ${like}
    WHERE menu_id = ${id}
  `);
}

export default {
  findByBizID,
  updateMenuLikes,
};
