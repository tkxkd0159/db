
import express from 'express';
import authR from './auth.route'

const router = express.Router();
router.use('/auth', authR);

router.get('/status', (req, res) => res.send('OK : /v1/status'))

export default router