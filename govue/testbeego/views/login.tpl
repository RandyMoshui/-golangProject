<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <title>用户登录</title>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />

</head>
<body>
    <div class="" id="login">
        <h1>Login</h1>
        <script src="https://unpkg.com/axios/dist/axios.min.js"></script>

            <form method="post">
            <input type="text" id="name" required="required" placeholder="用户名" name="name"></input>
            <input type="password" id="pass" required="required" placeholder="密码" name="pass"></input>
            </form>
            <button class="but" name="button" onclick=" sub();">登录</button>
    </div>

    <script>
        function sub(){
            // console.log('ok!');
            axios({
                    method: 'post',
                    url: '/login',
                    data: {
                        name: 'luke',
                        pass: '123456'
                    }
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });

        }
    </script>


</body>
</html>