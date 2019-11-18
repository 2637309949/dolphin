package tempalte

// TmplAuth defined template
var TmplAuth = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Auth</title>
    <link
      rel="stylesheet"
      href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"
    />
    <script src="//code.jquery.com/jquery-2.2.4.min.js"></script>
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
  </head>

  <body>
    <div class="container">
      <div class="jumbotron">
        <form action="/api/oauth2/affirm" method="POST">
          <h1>Authorize</h1>
          <p>The client would like to perform actions on your behalf.</p>
          <p>
            <button
              type="submit"
              class="btn btn-primary btn-lg"
              style="width:200px;"
            >
              Allow
            </button>
          </p>
        </form>
      </div>
    </div>
  </body>
</html>
`

// TmplLogin defined template
var TmplLogin = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css">
    <script src="//code.jquery.com/jquery-2.2.4.min.js"></script>
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
</head>

<body>
    <div class="container">
        <h1>Login In</h1>
        <form action="/api/oauth2/login" method="POST">
            <div class="form-group">
                <label for="username">User Name</label>
                <input type="text" class="form-control" name="username" placeholder="Please enter your user name">
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" class="form-control" name="password" placeholder="Please enter your password">
            </div>
            <button type="submit" class="btn btn-success">Login</button>
        </form>
    </div>
</body>

</html>
`
