import conf from './config'
import {app} from './app'


app.listen(conf.port, () => console.log(`Listening on ${conf.port}`))

console.log(process.env.SECRET_KEY)

