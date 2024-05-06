import express, { Express, Request, Response } from "express";

const app: Express = express();
const port: string = process.env.PORT || "4000";

app.get('/', function(req: Request, res: Response) {
	res.send('GET request to homepage')
})

app.get('/hello', function(req: Request, res: Response) {
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

const foo : number[] | string[] = ["&", "3"]

const isJustineAnAdult = isAdult(justine);
