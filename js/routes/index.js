import { Router } from 'express';
var router = Router();
import sql from '../db/sql.js';

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

const sectionIcons = [
  '🍚', '🍿', '🍜', '🍣', '🥩', '☕', '🍰'
]

router.get('/sections', async function(req, res, next) {

  const sections = await sql.getSections()

  sections.map((item) => {
    item.icon = sectionIcons[item.section_id - 1]
  })


  res.render('sections', { 
    title: '섹션 목록',
    sections
  });
});

export default router;
