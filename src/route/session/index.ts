import express from 'express';

const router = express.Router();

router.get("/", function(req, res) {
    res.json({msg: "I am hello from session", csrfToken: req.csrfToken(), expired: req.session.cookie.maxAge})
})

router.get("/form", async (req, res) => {
    res.render("send/send-sess");
});

router.post("/process", async (req, res) => {
    let sessData = await req.app.locals.redis.v4.get(`sess:${req.sessionID}`)
    sessData = JSON.parse(sessData)
    sessData['jwt'] = 'validatedJWT'
    await req.app.locals.redis.v4.set(`sess:${req.sessionID}`, JSON.stringify(sessData))
    res.send(`data is being processed: ${req.body['favoriteColor']}`);
});

router.get("/logout", async (req, res) => {
    await req.app.locals.redis.del(`sess:${req.sessionID}`);
    res.send("logout & remove session");
})


export default router;