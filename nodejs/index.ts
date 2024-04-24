const express = require('express')

const app = express();
const port = 4000;

app.get('/', function(req, res) {
	res.send('GET request to homepage')
})

app.get('/hello', function(req, res) {
	res.sendFile('index.html', {root: __dirname})
})

app.listen(port, () => { 
	console.log(`Now listening on port ${port}`)
})

type User = {
    name: string;
    age: number;
  };
  
  function isAdult(user: User): boolean {
    return user.age >= 18;
  }
  
  const justine: User = {
    name: 'Justine',
    age: 20,
  };
  
  const isJustineAnAdult: boolean = isAdult(justine);
