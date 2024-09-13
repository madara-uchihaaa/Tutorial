const express = require("express");

const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.post("/login", (req, res) => {
  console.log("req.body", req.body);
  res.send({
    message: "Login success",
    data: req.body,
  });
});

app.post("/postForm", (req, res) => {
  console.log("req.body", req.body);
  res.send({
    message: "Post form success",
    data: req.body,
  });
});

app.listen(3000, () => {
  console.log("Server is running on port 3000");
});
