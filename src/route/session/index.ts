import express from 'express';
import redis from "../../db/redis"


const router = express.Router();

router.get("/", function(req, res) {
    res.json({msg: "I am hello from session", csrfToken: req.csrfToken()})
})

router.get("/form", function(req, res) {

    res.render("send/send-sess", { csrfToken: req.csrfToken() });
});

router.post("/process", function(req, res) {
    res.send("data is being processed");
});

router.get("/logout", async (req, res) => {
    const redisClient = await redis.getClient();
    await redisClient.del(`sess:${req.sessionID}`);
    res.send("logout & remove session");
})


export default router;