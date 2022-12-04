import sql from "../db/sql.js";

class Rating {
  constructor(row) {
    this.id = row.rating_id;
    this.bizID = row.fk_business_id;
    this.stars = row.stars;
    this.created = row.created_kor;
    this.comment = row.comment;
  }
}

async function findByBizID(id) {
  const q = `
  SELECT rating_id, stars, comment,
  DATE_FORMAT(created, "%y년 %m월 %d일 %p %h시 %i분 %s초") AS created_kor
  FROM ratings
  WHERE fk_business_id = ${id}
  `;
  const [rows] = await sql.query(q);
  return rows.map((item) => new Rating(item));
}

async function addRating(bizid, stars, comment) {
  const q = `
  INSERT INTO ratings(fk_business_id, stars, comment)
  VALUES (${bizid}, '${stars}', '${comment}')
  `;
  return sql.query(q);
}

async function removeRating(id) {
  const q = `
  DELETE FROM ratings
  WHERE rating_id = ${id}
  `;
  return sql.query(q);
}

export default {
  findByBizID,
  addRating,
  removeRating,
};
