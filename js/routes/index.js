import { Router } from "express";
import bizC from "../controllers/business.js";

const router = Router();

router.get("/", bizC.foodView);
router.get("/biz-simple", bizC.simpBizView);
router.get("/biz-adv", bizC.advBizView);
router.get("/business/:id", bizC.BizIdView);
router.put("/menus/:id", bizC.addStarToMenu);
router.post("/ratings", bizC.addRating);
router.delete("/ratings/:id", bizC.removeRating);

export default router;
