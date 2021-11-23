<div id="top"></div>
<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <a href="https://github.com/alfiancikoa/alta-gproject-restful-api/">
    <img src="images/logo.gif" alt="Logo" width="700" height="275">
  </a>

  <h3 align="center">Project#1 "Alta-Shop" E-Commerce </h3>

  <p align="center">
    Projek Pertama Pembangunan RESTful API Program Immersive Back End Batch 4
    <br />
    <a href="https://github.com/alfiancikoa/alta-gproject-restful-api"><strong>Kunjungi kami »</strong></a>
    <br />
  </p>
</div>

### 🛠 &nbsp;Build App & Database

![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;

<!-- ABOUT THE PROJECT -->
### 💻 &nbsp;About The Project

Alta-Shop merupakan projek pertama kami untuk membangun sebuah RESTful API E-commerce dengan menggunakan bahasa Golang.    
dilengkapi dengan berbagai fitur yang memungkinkan user untuk mengakses data yang ada didalam server. Adapun fitur yang ada dalam RESTful API kami antara lain :
<div>
      <details>
<summary>🙎 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Di User terdapat fitur untuk membuat Akun dan Login agar mendapat legalitas untuk mengakses berbagai fitur lain di aplikasi, 
 terdapat juga fitur Update untuk mengedit data yang berkaitan dengan user, serta fitur delete berfungsi jika user menginginkan hapus akun.
 
<div>
  
| Feature User | Format JSON |
| --- | --- |
| [e.POST("/users", user.CreateUserController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/user/user.go) |  |
| [e.POST("/login", user.LoginUsersController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/user/user.go) |  |
| [eJWT.GET("/users/:id", user.GetUserByIdController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/user/user.go) |  |
| [eJWT.PUT("/users/:id", user.UpdateUserController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/user/user.go) |  |
| [eJWT.DELETE("/users/:id", user.DeleteUserController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/user/user.go) |  |

</details>  

<details>
<summary>🏷&nbsp;Category</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
  Category berfungsi untuk mengelompokkan berbagai product agar user lebih mudah mencari barang yang dibutuhkan, terdapat fitur Insert untuk membuat category product,
  dan GET merupakan fitur untuk user mendapatkan product sesuai Category.
  
| Feature Category | Format JSON |
| --- | --- |
| [e.POST("/products/category", category.InsertCategoryController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/category/category.go) | |
| [e.GET("/products/category", category.GetAllCategorysController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/category/category.go) | |

</details>

<details>
<summary>📦&nbsp;Products</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
| Feature Products | Format JSON |
| --- | --- |
| [e.GET("/products", product.GetAllProductsController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |
| [eJWT.GET("/products/:id", product.GetProductController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |
| [eJWT.GET("/products/my", product.GetMyProductController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |
| [eJWT.POST("/products", product.CreateProductsController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |
| [eJWT.DELETE("/products/:id", product.DeleteProductController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |
| [eJWT.PUT("/products/:id", product.UpdateProductController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/product/product.go) | |

</details>

<details>
<summary>🛒&nbsp;Cart</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
| Feature Cart | Format JSON |
| --- | --- |
| [eJWT.POST("/carts", cart.CreateCartController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/cart/cart.go) | |
| [eJWT.GET("/carts/my", cart.GetCartController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/cart/cart.go) | |
| [eJWT.PUT("/carts/:id", cart.UpdateCartController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/cart/cart.go) | |
| [eJWT.DELETE("/carts/:id", cart.DeleteCartController)](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/cart/cart.go) | |

</details>

<details>
<summary>💳&nbsp;Order</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
| Feature Order | Format JSON |
| --- | --- |
| [eJWT.POST("/orders", order.CreateNewOrderController))](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/order/order.go) | |
| [eJWT.GET("/orders", order.GetOrderController))](https://github.com/alfiancikoa/alta-gproject-restful-api/blob/main/controllers/order/order.go) | |

</details>
      

<!-- ERD -->
### ERD
<img src="images/erd.jpg">

<!-- CONTACT -->
### Contact

[![Gmail: Fian](https://img.shields.io/badge/-Muhammad%20Alfian-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=CllgCHrjmjRlSpLttDDmhqnRQTQVTSQCjFvQxCSSqGDHvQjrjJvvzKMvnlWTrWwkcGdSzfJPXnV)
[![GitHub Fian](https://img.shields.io/badge/-alfiancikoa-white?style=flat&logo=github&logoColor=black)](https://github.com/alfiancikoa)

[![Gmail: Fafa](https://img.shields.io/badge/-Naufal%20Muhammad-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=DmwnWslzCnrLrhrlnrRWdpHqsBmRtbbtZSKxXFrdGHmhLVLjLDmVfNRxdBShrxQNTBBHFgDdLfKQ)
[![GitHub Fafa](https://img.shields.io/badge/-DylanRipper-white?style=flat&logo=github&logoColor=black)](https://github.com/DylanRipper)

[![Gmail: Supriadi](https://img.shields.io/badge/-Supriadi-maroon?style=flat&logo=gmail)]()
[![GitHub Supriadi](https://img.shields.io/badge/-sprdx-white?style=flat&logo=github&logoColor=black)](https://github.com/sprdx)


<p align="center">:copyright: 2021 | ANS</p>
</h3>
