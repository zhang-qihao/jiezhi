var loginbodybox_1_text4 = document.querySelector(".loginbodybox_1_text4")
var loginbodyboximg = document.querySelector(".loginbodybox_3")
var inputphone      /* 手机号框 */
var getcode = document.getElementById("getcode");   /* 获取验证码按钮 */
var inputcode       /* 验证码 */
var argument        /* 用户协议选项 */
var codebutton      /* 获取验证码按钮 */
var inputpassword   /* 密码框 */
var loginbodybox_2_text3 = document.querySelector('.loginbodybox_2_text3');
var changeweiht2 = document.querySelector('.loginbodybox_2_title2 a');
var changeweiht1 = document.querySelector('.loginbodybox_2_title1 a');
var loginbodybox_2_input2 = document.querySelector('.loginbodybox_2_input2')
var loginbodybox_2_input3 =document.querySelector('.loginbodybox_2_input3');
var loginbodybox_2_button1 = document.querySelector('.loginbodybox_2_button1');
var loginimg = document.querySelector('#line');
/* 定义标签 */

/* 公众号显示隐藏 */
loginbodybox_1_text4.addEventListener('mouseover', function () {
    loginbodyboximg.style.display = "block";
});
loginbodybox_1_text4.addEventListener('mouseout', function () {
    loginbodyboximg.style.display = "none";
})


/* 获取验证码按钮点击事件 */
getcode.addEventListener('click',function (){
    inputphone = document.getElementById("inputphone").value;
    codebutton = document.getElementById("getcode").value;
    if (inputphone == "") {
        alert("请输入手机号")
    }else {
        if (!(/^1[345789]\d{9}$/.test(inputphone))){
            $('.loginbodybox_2_text1').show();
        }else{
            /* 手机号输入成功 */
            $('.loginbodybox_2_text1').hide();

            var time = 60;
            getcode.disabled = false;
            var timer = setInterval(function () {
                if (time == 0) {
                    clearInterval(timer);
                    //这里时设置当时间到0的时候重新设置点击事件，并且默认time修改为60
                    getcode.disabled = false;
                    loginbodybox_2_text3.style.display = "block";
                    getcode.innerHTML = "<span>获取验证码</span>";
                    time = 60;
                } else {
                    getcode.disabled = true;
                    loginbodybox_2_text3.style.display = "none";
                    getcode.innerHTML = "<span>" + time + "s</span>";
                    time--;
                }
            }, 1000);
            /* 向后端传手机号框和获取验证码按钮的value值 */
            $.ajax({
                url: "/v0/getcode",
                type: "POST",
                dataType: "json",
                data: {"inputphone": inputphone, "codebutton": codebutton},
                success: function (returnPoint) {
                    console.log(returnPoint)
                },
                error: function (){
                    console.log("ERROR：ERROR：后端没有传JSON数据")
                }
            })
        }
    }
})


/* 个人用户注册/登陆页面 */
function personRegister(){
    var ID = getQueryVariable("id")
    inputphone = document.getElementById("inputphone").value;
    inputcode = document.getElementById("inputcode").value;
    argument = document.getElementById("argument").value;
    if (inputphone == "") {
        alert("请输入手机号")
    }else{
        if (inputcode == "") {
            alert("请输入验证码")
        }else {
            var val=$('input:radio[name="loginbody_2choose2"]:checked').val();
            if (val == null){
                alert("请勾选用户协议");
            }else {
                /* 手机号、验证码、用户协议输入完毕 */
                /* 向后段传手机号、验证码输入框的value值 */
                $.ajax({
                    url: "/v1/personlogin",
                    type: "POST",
                    dataType: "json",
                    data: {
                        "inputphone": inputphone,
                        "inputcode": inputcode,
                    },
                    success: function (returnPoint){
                        console.log(returnPoint)
                        if (returnPoint.msg == "新用户注册成功"){
                            alert("新用户注册成功");
                            var url = "/v2/recruit"
                            window.location.href = url
                        }
                        if (returnPoint.msg == "登陆成功"){
                            alert("登陆成功");
                            var url = "/v4/home?pid=" + returnPoint.ID
                            window.location.href = url
                        }
                        if (returnPoint.msg == "验证码错误"){
                            alert("验证码错误");                         /* 不发生页面跳转 */
                        }
                        if (returnPoint.msg == "未填写个人简历"){
                            alert("登陆成功")
                            var url = "/v2/recruit"
                            window.location.href = url
                        }
                    },
                    error: function (){
                        console.log("ERROR：后端没有传JSON数据")
                    }
                })
            }
        }
    }
}

/* 企业注册页面 */
function cpRegister(){
    inputphone = document.getElementById("inputphone").value;
    inputcode = document.getElementById("inputcode").value;
    argument = document.getElementById("argument").value;
    if (inputphone == "") {
        alert("请输入手机号")
    }else{
        if (inputcode == "") {
            alert("请输入验证码")
        }else {
            var val=$('input:radio[name="loginbody_2choose2"]:checked').val();
            if (val == null){
                alert("请勾选用户协议");
            }else {
                /* 手机号、验证码、用户协议输入完毕 */
                /* 向后段传手机号、验证码输入框的value值 */
                $.ajax({
                    url: "/v1/cpregister",
                    type: "POST",
                    dataType: "json",
                    data: {
                        "inputphone": inputphone,
                        "inputcode": inputcode,
                    },
                    success: function (returnPoint){
                        console.log(returnPoint)
                        if (returnPoint.msg == "新用户注册成功"){
                            alert("注册成功");
                            window.location.href="/v1/cplogin";       /* 跳转到企业登陆主页 */
                        }
                        if (returnPoint.msg == "验证码错误"){
                            alert("验证码错误");                       /* 不发生页面跳转 */
                        }
                        if (returnPoint.msg == "手机号码已注册"){
                            alert("手机号码已注册");                   /* 不发生页面跳转 */
                        }
                    },
                    error: function (){
                        console.log("ERROR：后端没有传JSON数据")
                    }
                })
            }
        }
    }
}



/* 企业短信登陆和验证码登陆页面 */
function cpLogin(){
    var ID = getQueryVariable("id")
    inputphone = document.getElementById("inputphone").value;
    inputcode = document.getElementById("inputcode").value;
    argument = document.getElementById("argument").value;
    inputpassword = document.getElementById('inputpassword').value;
    if (inputphone == ""){
        alert("请输入手机号")
    }else {
        if ( loginbodybox_2_button1.style.display == "block"){
            /* 进行短信登录 */
            if (inputcode == ""){
                alert("请输入验证码")
            }else {
                var val=$('input:radio[name="loginbody_2choose2"]:checked').val();
                if (val == null){
                    alert("请勾选用户协议");
                }else {
                    /* 向后端传手机号和验证码 */
                    $.ajax({
                        url: "/v1/cplogin",
                        type: "POST",
                        dataType: "json",
                        data: {
                            "inputphone": inputphone,
                            "inputcode": inputcode,
                            "ID": ID,
                        },
                        success: function (returnPoint){
                            console.log(returnPoint)
                            if (returnPoint.msg == "登陆成功"){
                                alert("登陆成功");
                                if (returnPoint.ID == "") {
                                    window.location.href = "/v3/companyInfoEdit?tel=" + inputphone
                                }else {
                                    var url = "/v3/company?cid=" + returnPoint.ID
                                    window.location.href = url
                                }
                            }
                            if (returnPoint.msg == "验证码错误"){
                                alert("验证码错误");   /* 不发生页面跳转 */
                                inputcode = "";
                            }
                            if (returnPoint.msg == "手机号未注册"){
                                alert("手机号未注册");    /* 不发生跳转 */
                                inputphone = "";
                                inputpassword = "";
                            }
                        },
                        error: function (returnPoint){
                            console.log("ERROR：后端没有传JSON数据")
                        }
                    })
                }
            }
        }else {
            /* 进行密码登录 */
            if (inputpassword == ""){
                alert("请输入密码")
            }else {
                var val=$('input:radio[name="loginbody_2choose2"]:checked').val();
                if (val == null){
                    alert("请勾选用户协议");
                }else {
                    /* 向后端传手机号码和密码 */
                    $.ajax({
                        url: "/v1/cplogin",
                        type: "POST",
                        dataType: "json",
                        data: {
                            "inputphone": inputphone,
                            "inputpassword": inputpassword,
                            "ID": ID,
                        },
                        success: function (returnPoint){
                            console.log(returnPoint)
                            if (returnPoint.msg == "登陆成功"){
                                alert("登陆成功");
                                if (returnPoint.ID == "") {
                                    window.location.href = "/v3/companyInfoEdit"
                                }else {
                                    var url = "/v3/company?cid=" + returnPoint.ID
                                    window.location.href = url
                                }
                            }
                            if (returnPoint.msg == "密码错误"){
                                alert("密码错误");   /* 不发生页面跳转 */
                            }
                            if (returnPoint.msg == "手机号未注册"){
                                alert("手机号未注册");    /* 不发生跳转 */
                            }
                        },
                        error: function (returnPoint){
                            console.log("ERROR：后端没有传JSON数据")
                        }
                    })
                }
            }
        }
    }

}

changeweiht1.className = 'changeweight';
loginbodybox_2_input3.style.display = "none";       //密码
loginbodybox_2_button1.style.display = "block";     //验证码

/* 短信登录和密码登录切换 */
function pwdLogin(){
    changeweiht2.className = 'changeweight';
    loginbodybox_2_input3.style.display = "block";      //密码
    loginbodybox_2_button1.style.display = "none";      //验证码
    loginimg.style.display = "none";
    changeweiht1.className = '';
    loginbodybox_2_input2.value = ''
}
function codeLogin(){
    changeweiht1.className = 'changeweight';
    loginbodybox_2_input3.style.display = "none";       //密码
    loginbodybox_2_button1.style.display = "block";     //验证码
    changeweiht2.className = '';
    loginimg.style.display = "block";
    loginbodybox_2_input3.value = ''
}


function SetCookies(cookieName, cookieValue){
    document.cookie = cookieName + "=" + cookieValue
}

function getQueryVariable(variable)
{
    var query = window.location.search.substring(1);
    var vars = query.split("&");
    for (var i=0;i<vars.length;i++) {
        var pair = vars[i].split("=");
        if(pair[0] == variable){return pair[1];}
    }
    return(false);
}


