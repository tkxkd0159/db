import foodSections from "../models/food-section.js";
import foodBiz from "../models/business.js";
import foodMenu from "../models/menu.js";
import foodRating from "../models/rating.js";

const foodView = async (req, res, _next) => {
  res.render("food_biz/sections", {
    title: "섹션 목록",
    lists: await foodSections.sectionsWithIcon(),
  });
};

const simpBizView = async (req, res, _next) => {
  res.render("food_biz/biz-simple", {
    title: "단순 식당 목록",
    lists: await foodBiz.findAllByQuery(req.query),
  });
};

const advBizView = async (req, res, _next) => {
  res.render("food_biz/biz-adv", {
    title: "고급 식당 목록",
    q: req.query,
    lists: await foodBiz.findAllByQuery(req.query),
  });
};

const BizIdView = async (req, res, _next) => {
  const biz = await foodBiz.findByID(req.params.id);
  const menus = await foodMenu.findByBizID(req.params.id);
  const ratings = await foodRating.findByBizID(req.params.id);

  res.render("food_biz/detail", {
    biz,
    menus,
    ratings,
  });
};

const addStarToMenu = async (req, res, _next) => {
  const updateRes = await foodMenu.updateMenuLikes(req.params.id, req.body.like);
  res.send(updateRes);
};

const addRating = async (req, res, _next) => {
  const addRes = await foodRating.addRating(req.body.business_id, req.body.stars, req.body.comment);
  res.send(addRes);
};

const removeRating = async (req, res, _next) => {
  const removeRes = await foodRating.removeRating(req.params.id);
  res.send(removeRes);
};

export default {
  foodView,
  simpBizView,
  advBizView,
  BizIdView,
  addStarToMenu,
  addRating,
  removeRating,
};
