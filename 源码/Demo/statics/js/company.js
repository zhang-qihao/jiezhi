/* 获取验证码按钮 */
var getcode = document.getElementById("code");



// btns = document.querySelectorAll(".companyBtn")
// for (let index = 0 ;index < btns.length;index++){
//     const btn = btns[index];
//     btn.onclick = () =>{
//         if(btn.className == 'companyBtn'){
//             btn.className = 'resumitBtn';
//         }
//         else{
//             btn.className = 'companyBtn';
//         }
//     }
// }
// //边  三选一
// one_side.addEventListener('click',function(){
//     one_side.className = 'resumeside';
//     two_side.className = 'orgside';
//     three_side.className = 'orgside';
// })
// two_side.addEventListener('click',function(){
//     two_side.className = 'resumeside';
//     one_side.className = 'orgside';
//     three_side.className = 'orgside';
// })
// three_side.addEventListener('click',function(){
//     three_side.className = 'resumeside';
//     two_side.className = 'orgside';
//     one_side.className = 'orgside';
// })
//
// company_type1.addEventListener('click',function(){
//     company_type1.className = 'resumitBtn';
//     company_type2.className = 'companyBtn';
//     company_type3.className = 'companyBtn';
// })
// company_type2.addEventListener('click',function(){
//     company_type2.className = 'resumitBtn';
//     company_type1.className = 'companyBtn';
//     company_type3.className = 'companyBtn';
// })
// company_type3.addEventListener('click',function(){
//     company_type3.className = 'resumitBtn';
//     company_type2.className = 'companyBtn';
//     company_type1.className = 'companyBtn';
// })
//
//
//
//
//
//
// //跳出居住城市界面
// province.addEventListener('click', function () {
//     if (userPlaceAll.style.display = 'none') {
//         userPlaceAll.style.display = 'block';
//     }
// })
// city.addEventListener('click', function () {
//     if (userPlaceAll.style.display = 'none') {
//         userPlaceAll.style.display = 'block';
//     }
// })
// //居住城市弹出框点击变色与跳转界面
// var tds = document.querySelectorAll(".userPlaceMore2");
// var divs = document.querySelectorAll(".userPlaceMoreAll")
// for (var i = 0; i < tds.length; i++) {
//     tds[i].setAttribute('index', i);
//     tds[i].onclick = function () {
//         for (var i = 0; i < tds.length; i++) {
//             tds[i].className = ' ';
//         }
//         // 留下我自己
//         this.className = 'canPickTex1';
//         //var ss = this.getAttribute('value');
//         document.getElementById('province').value = this.getAttribute('value');
//         // 2. 下面的显示内容模块
//         var index = this.getAttribute('index');
//         // 干掉所有人 让其余的item 这些div 隐藏
//         for (var i = 0; i < divs.length; i++) {
//             divs[i].style.display = 'none';
//         }
//         // 留下我自己 让对应的item 显示出来
//         divs[index].style.display = 'block';
//
//     }
// }
//
// function changeUserPlace(value){
//     console.log(value);
//     userPlaceAll.style.display = 'none';
//     document.getElementById('city').value = value;
// }
//
// //公司规模
// var mySelect=$("#company_size option");
// var num="请输入公司规模";//某个值
// mySelect.each(function (i,el) {
//     if($(el).text()==num){
//         $(this).hide();
//     }
// })
//
// //上传图片
// document.getElementById('com_image').addEventListener('change',function(e){
//     var files = this.files;
//     var img = new Image();
//     var reader = new FileReader();
//     reader.readAsDataURL(files[0]);
//     reader.onload = function(e){
//         var mb = (e.total/1024)/1024;
//         if(mb>= 0.5){
//             alert('文件大小大于500k');
//             return;
//         }
//         img.src = this.result;
//         img.style.width = "101px";
//         img.style.height = "101px";
//         document.getElementById('click').style.width="101px";
//         document.getElementById('click').style.height="101px";
//         document.getElementById('click').innerHTML = '';
//         document.getElementById('click').appendChild(img);
//     }
// });



//进入密码管理
function repword(){
    document.getElementById('password').style.display='block';
}
//清空input
function returnPage(){
    document.getElementById('password').style.display='none';
    $("#cen input").val("");
    document.getElementById('tips1').style.display='none';
    document.getElementById('tips2').style.display='none';
    document.getElementById('tips3').style.display='none';
    document.getElementById('tips4').style.display='none';
}

//密码格式
function cheeckPassword(){
    var _password = $("#InputUserPwd").val();
    if(!(/^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/.test(_password))){
        document.getElementById('tips1').style.display='block';
        return false;
    }else{
        document.getElementById('tips1').style.display='none';
        return false;
    }

}
//确认密码
function conPassword(){
    var pass1 = document.getElementById("InputUserPwd").value;
    var pass2 = document.getElementById("ConfirmPassword").value;
    console.log(pass1);
    console.log(pass2);
    if(pass1 == pass2){
        document.getElementById('tips2').style.display='none';
        return false;
    }else{
        document.getElementById('tips2').style.display='block';
        return false;
    }
}
//确认手机号
function checkTel(){
    var _phone = $("#InputTelephone").val();
    if(!(/^1[3|5|6|7|8|9][0-9]\d{8}$/.test(_phone))){
        document.getElementById('tips3').style.display='block';
        return false;
    }else{
        document.getElementById('tips3').style.display='none';
        return false;
    }
}

/*密码复用js*/
//进入密码管理
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
var cid =  getQueryVariable("cid")
console.log(cid)
function repword(){
    document.getElementById('password').style.display='block';
}
//清空input
function returnPage(){
    document.getElementById('password').style.display='none';
    $("#cen input").val("");
    document.getElementById('tips1').style.display='none';
    document.getElementById('tips2').style.display='none';
    document.getElementById('tips3').style.display='none';
    document.getElementById('tips4').style.display='none';
}

//密码格式
function cheeckPassword(){
    var _password = $("#InputUserPwd").val();
    if(!(/^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/.test(_password))){
        document.getElementById('tips1').style.display='block';
        return false;
    }else{
        document.getElementById('tips1').style.display='none';
        return true;
    }

}
//确认密码
function conPassword(){
    var pass1 = document.getElementById("InputUserPwd").value;
    var pass2 = document.getElementById("ConfirmPassword").value;
    console.log(pass1);
    console.log(pass2);
    if(pass1 == pass2){
        document.getElementById('tips2').style.display='none';
        return true;
    }else{
        document.getElementById('tips2').style.display='block';
        return false;
    }
}
//确认手机号
function checkTel(){
    var _phone = $("#InputTelephone").val();
    if(!(/^1[3|5|6|7|8|9][0-9]\d{8}$/.test(_phone))){
        document.getElementById('tips3').style.display='block';
        return false;
    }else{
        document.getElementById('tips3').style.display='none';
        return true;
    }
}
//获取验证码
var btn = document.querySelector('#code');
btn.addEventListener('click', function() {
    var time = 60;
    var timer = null;
    timer = setInterval(function(){
        btn.disabled = true;
        btn.innerText = time + "秒后重新发送";
        time--;
        if(time == 0 ){
            btn.innerText = '重新发送验证码';
            btn.disabled = false;
            clearInterval(timer);
        }

    },1000)

})
//确认验证码
function confirmBtn(){
    inputphone = document.getElementById("InputTelephone").value;
    inputpassword = document.getElementById("InputUserPwd").value;
    comfirmpassword = document.getElementById("ConfirmPassword").value;
    codebutton = document.getElementById("code").value;
    inputcode = document.getElementById("InputCode").value;
    console.log(codebutton)
    if (inputphone == "" || inputpassword == "" || comfirmpassword == "") {
        alert("手机号密码请填写完整")
    }else {
        if (!(/^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/.test(inputpassword))) {
            alert("密码格式错误")
        } else {
            if (!(/^1[345789]\d{9}$/.test(inputphone))) {
                alert("手机号格式错误")
            } else {
                if (inputpassword != comfirmpassword) {
                    alert("密码不一致")
                } else {
                    if (inputcode == ""){
                        alert("请输入验证码")
                    }else{
                        $.ajax({
                            url: "/v3/companyInfoEdit",
                            type: "POST",
                            dataType: "json",
                            data: {"inputphone": inputphone, "inputpassword": inputpassword, "inputcode":inputcode, "ButtonType": "SetPwd"},
                            success: function (returnPoint) {
                                console.log(returnPoint)
                                if (returnPoint.msg == "验证码错误") {
                                    alert("验证码错误")
                                }
                                if (returnPoint.msg == "手机号码不存在") {
                                    alert("手机号码不存在")
                                }
                                if(returnPoint.msg == "信息更新成功") {
                                    alert("信息更新成功")
                                    var CID = getQueryVariable("cid")
                                    var url = "/v3/company?cid=" + CID
                                    window.location.href = url
                                }
                            },
                            error: function () {
                                console.log("ERROR：ERROR：后端没有传JSON数据")
                            }
                        })
                    }
                }
            }
        }
    }
}
var fullname = document.querySelector('#companyName');
var companytype = document.querySelector('#companyType1');
var shortname = document.querySelector('#shortname');
var scale = document.querySelector('#scale');
var address = document.querySelector('#address');
var brief = document.querySelector('#brief');

$("overall").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid
        },
        url: "/v3/company",
        success:function(result){
            console.log(result)
            console.log(result.title)
            fullname.innerHTML = result.fullname;
            companytype.innerHTML = result.companytype
            shortname.innerHTML = result.shortname;
            scale.innerHTML = result.scale;
            address.innerHTML = result.address;
            brief.innerHTML = result.brief;
        }
    })
})

var inputphone
var comfirmpassword
var inputpassword
var codebutton
var getcode = document.getElementById("code")
var comfirmpassword
var inputcode = document.getElementById("InputCode")
function CreatCode() {
    inputphone = document.getElementById("InputTelephone").value;
    inputpassword = document.getElementById("InputUserPwd").value;
    comfirmpassword = document.getElementById("ConfirmPassword").value;
    codebutton = document.getElementById("code").value;
    console.log(codebutton)
    if (inputphone == "" || inputpassword == "" || comfirmpassword == "") {
        alert("手机号密码请填写完整")
    }else {
        if (!(/^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/.test(inputpassword))){
            alert("密码格式错误")
        }else{
            if (!(/^1[345789]\d{9}$/.test(inputphone))){
                alert("手机号格式错误")
            }else {
                if (inputpassword != comfirmpassword) {
                    alert("密码不一致")
                } else {
                    /* 手机号输入成功 */
                    var time = 60;
                    // getcode.disabled = false;
                    var timer = setInterval(function () {
                        if (time == 0) {
                            clearInterval(timer);
                            //这里时设置当时间到0的时候重新设置点击事件，并且默认time修改为60
                            getcode.disabled = false;
                            getcode.innerHTML = "<span>获取验证码</span>";
                            time = 60;
                        } else {
                            getcode.disabled = true;
                            getcode.innerHTML = "<span>" + time + "s</span>";
                            time--;
                        }
                    }, 1000);
                    /* 向后端传手机号框和获取验证码按钮的value值 */
                    $.ajax({
                        url: "/v0/getcode",
                        type: "POST",
                        dataType: "json",
                        data: {"inputphone": inputphone, "inputpassword": inputpassword},
                        success: function (returnPoint) {
                            console.log(returnPoint)
                        },
                        error: function (){
                            console.log("ERROR：ERROR：后端没有传JSON数据")
                        }
                    })
                }
            }
        }
    }
}

function ViewJob(jid) {
    var url = "/v3/jobEdit?cid="+cid+"&jid="+jid
    console.log(url)
    window.location.href = url
}


function logout(){
    alert("退出成功")
    $.ajax({
        url: "/v3/companyInfoEdit",
        type: "POST",
        dataType: "json",
        data: {"exit": "exit"}
    })
    var url = "/v1/cplogin"
    window.location.href = url;
}

function jump1(){
    var url = "/v3/company?cid=" + cid
    console.log(url)
    window.location.href=url
}
function jump2(){
    var url = "/v3/positionRelease?cid=" + cid
    console.log(url)
    window.location.href=url
}
function jump3(){
    var url = "/v3/manageApply?cid=" + cid
    window.location.href=url
}

function companyjump() {
    var url = "/v3/companyInfoEdit?cid=" + cid
    window.location.href=url
}


function addJob(){
    var url = "/v3/jobCreate?cid=" + cid
    window.location.href=url
}
//保存公司基本信息按钮事件
// var SubmitCompanyInfo = document.getElementById("submit");   //获取提交信息按钮
// SubmitCompanyInfo.addEventListener('click',function () {
//     fullName = document.getElementsByName("fullName")
//     shortName = document.getElementsByName("shortName")
//     companyType = document.getElementsByName("companyType")
//     scale = document.getElementsByName("scale")
//     province = document.getElementsByName("province")
//     city = document.getElementsByName("city")
//     specificAddress = document.getElementsByName("specificAddress")
//     logUrl = document.getElementsByName("logUrl")
//     brief = document.getElementsByName("brief")
//
//     if (fullName == null || companyType == null || scale == null || province == null || city == null){
//         alert("信息未输全，请重新输入")
//     }else{
//         $.ajax({
//             url: "v3/company",
//             type: "POST",
//             dataType:"json",
//             data:{
//                 "type": "SaveInfo",
//                 "fullName":fullName,
//                 "shortName":shortName,
//                 "companyType":companyType,
//                 "scale":scale,
//                 "province":province,
//                 "city":city,
//                 "specificAddress":specificAddress,
//                 "logUrl":logUrl,
//                 "brief":brief,},
//             success:function (returnPoint) {
//                 console.log("ERROR:向后端传输信息成功")
//             },
//             error: function(){
//                 console.log("ERROR:向后端传输信息失败")
//             }
//         })
//     }
//
// })