import express from 'express';

const router = express.Router();

router.get("/", function(req, res) {
    res.json({msg: "I am hello from cookie", csrfToken: req.csrfToken()})
})

router.get("/form", function(req, res) {

    res.render("send/send-cookie", { csrfToken: req.csrfToken() });
});

router.post("/process", function(req, res) {
    res.send("data is being processed");
});


export default router;