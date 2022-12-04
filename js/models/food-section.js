import sql from "../db/sql.js";

const sectionIcons = ["🍚", "🍿", "🍜", "🍣", "🥩", "☕", "🍰"];

class FoodSection {
  icon;

  constructor(row) {
    this.id = row.section_id;
    this.name = row.section_name;
  }
}

async function findAll() {
  const [rows] = await sql.query(`
      SELECT * FROM sections
    `);
  return rows.map((item) => new FoodSection(item));
}

async function findByID(id) {
  const [rows] = await sql.query(`
  SELECT * FROM sections where = ${id}
`);
  return new FoodSection(rows[0]);
}

async function sectionsWithIcon() {
  const sections = await findAll();
  return sections.map((item) => {
    const i = item;
    i.icon = sectionIcons[item.id - 1];
    return i;
  });
}

export {
  FoodSection, findAll, findByID, sectionsWithIcon,
};
