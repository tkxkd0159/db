import express from 'express';

const router = express.Router();

router.get("/", function(req, res) {
    res.json({msg: "I am hello from session", csrfToken: req.csrfToken(), expired: req.session.cookie.maxAge})
})

router.get("/csrf", async (req, res) => {
    res.json({value: req.csrfToken()});
});

router.get("/form", async (req, res) => {
    res.render("send/send-sess");
});

router.post("/process", async (req, res) => {
    req.session.jwt = "newValidatedJWT"
    res.send(`data is being processed: ${req.body['favoriteColor']}`);
});

router.get("/logout", async (req, res) => {
    await req.app.locals.redis.del(`sess:${req.sessionID}`);
    res.send("logout & remove session");
})


export default router;