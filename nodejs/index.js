"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = require("express");
var app = (0, express_1.default)();
var port = process.env.PORT || "4000";
app.get('/', function (req, res) {
    res.send('GET request to homepage');
});
app.get('/hello', function (req, res) {
    res.sendFile('index.html', { root: __dirname });
});
app.listen(port, function () {
    console.log("Now listening on port ".concat(port));
});
function isAdult(user) {
    return user.age >= 18;
}
var justine = {
    name: 'Justine',
    age: 20,
};
var isJustineAnAdult = isAdult(justine);
