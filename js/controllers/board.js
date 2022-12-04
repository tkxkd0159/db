import { sectionsWithIcon } from "../models/food-section.js";

const foodView = async (req, res, _next) => {
  res.render("sections", {
    title: "섹션 목록",
    list: await sectionsWithIcon(),
  });
};

export default {
  foodView,
};
