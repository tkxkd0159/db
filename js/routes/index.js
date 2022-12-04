import { Router } from "express";
import boardC from "../controllers/board.js";

const router = Router();

router.get("/", (_req, res, _next) => {
  res.render("index", { title: "Express" });
});

router.get("/sections", boardC.foodView);

export default router;
