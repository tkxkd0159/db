import express from 'express';

const router = express.Router();

router.route('/info')
    .get((req, res) => res.send('I am /v1/auth/info'))

export default router