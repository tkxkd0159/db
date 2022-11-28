import { Router } from "express";
import sql from "../db/sql.js";

const router = Router();

/* GET home page. */
router.get("/", (_req, res, _next) => {
  res.render("index", { title: "Express" });
});

const sectionIcons = [
  "🍚", "🍿", "🍜", "🍣", "🥩", "☕", "🍰",
];

router.get("/sections", async (req, res, _next) => {
  const sections = await sql.getSections();

  const sectionsWithIcon = sections.map((item) => {
    const i = item;
    i.icon = sectionIcons[item.section_id - 1];
    return i;
  });

  res.render("sections", {
    title: "섹션 목록",
    sectionsWithIcon,
  });
});

export default router;
