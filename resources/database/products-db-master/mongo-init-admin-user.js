use admin
db.createUser(
  {
    user: "productListUser",
    pwd: "productListPassword",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)