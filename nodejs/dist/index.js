"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = (0, express_1.default)();
const port = process.env.PORT || "4000";
app.get('/', function (req, res) {
    res.send('GET request to homepage');
});
app.get('/hello', function (req, res) {
    res.sendFile('index.html', { root: __dirname });
});
app.listen(port, () => {
    console.log(`Now listening on port ${port}`);
});
function isAdult(user) {
    return user.age >= 18;
}
const justine = {
    name: 'Justine',
    age: 20,
};
const isJustineAnAdult = isAdult(justine);
