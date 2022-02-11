const express = require('express');
const app = express();
app.use(express.json());

const port = process.env.PORT || 8080;
app.listen(port, () => console.log(`listening on port ${port}`));

var users = [
    {
        id: 1,
        name: "foo",
    }, {
        id: 2,
        name: "bar"
    }
]

app.get('/', (req, res) => {
    res.send('HTTP GET request called');
});

app.get("/api/users", (req, res) => {
    res.json(users);
});

app.post('/', (req, res) => {
    res.send('HTTP POST request called');
});

app.post('/api/users', (req, res) => {
    const user = {
        id: users.length + 1,
        name: req.body.name
    };
    users.push(user);
    res.send(user);
});