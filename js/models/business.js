import sql from "../db/sql.js";
import { sectionIcons } from "./food-section.js";

const statusKorMap = {
  OPN: "영업중",
  CLS: "폐업",
  VCT: "휴가중",
  RMD: "리모델링",
};

class FoodBusiness {
  constructor(row) {
    this.id = row.business_id;
    this.name = row.business_name;
    this.status = row.status;
    this.sectionName = row.section_name;
    this.sectionID = row.section_id;
    this.floor = row.floor;
    this.canTakeout = row.can_takeout;
    this.statusKor = statusKorMap[this.status];
    this.icon = sectionIcons[this.sectionID - 1];
  }
}

async function findAllByQuery(query) {
  const q = `
      SELECT * FROM sections S
      LEFT JOIN businesses B
        ON S.section_id = B.fk_section_id
      WHERE TRUE
        ${query.section
    ? (`AND section_id = ${query.section}`) : ""}
        ${query.floor
    ? (`AND floor = ${query.floor}`) : ""}
        ${query.status
    ? (`AND status = '${query.status}'`) : ""}
      ORDER BY
         ${query.order
    ? query.order : "business_id"}
    `;
  const [rows] = await sql.query(q);
  return rows.map((item) => new FoodBusiness(item));
}

async function findByID(id) {
  const q = `
  SELECT * FROM sections S
  LEFT JOIN businesses B
  ON S.section_id = B.fk_section_id
  WHERE business_id = ${id}
  `;
  const [rows] = await sql.query(q);
  return new FoodBusiness(rows[0]);
}

export default {
  findAllByQuery,
  findByID,
};

export {
  FoodBusiness,
};
