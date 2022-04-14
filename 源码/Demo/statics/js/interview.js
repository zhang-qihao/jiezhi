//图片颜色更改
var user = document.getElementById("user");
var head_headRigth_img = document.getElementById("head_headRigth_img");
user.addEventListener('mouseover',function (){
    this.style.color = "rgb(46,176,252)";
    head_headRigth_img.src = "../statics/images/head__/u288_selected.svg";
})
user.addEventListener('mouseout',function (){
    this.style.color = "";
    head_headRigth_img.src = "../statics/images/head__/u288.svg";
})
//页面选择
var head_checkbox = document.querySelector(".head_checkbox");
var lis = head_checkbox.querySelectorAll('li');
    var items = document.querySelectorAll('.head_body1_list2');
    for (var i = 0; i < lis.length; i++) {
        // 开始给4个小li 设置索引号
        lis[i].setAttribute('index', i);
        lis[i].onclick = function() {
            // 1. 上的模块选项卡，点击某一个，当前这一个底色会是红色，其余不变（排他思想） 修改类名的方式
            // 干掉所有人 其余的li清除 class 这个类
            for (var i = 0; i < lis.length; i++) {
                lis[i].className = '';
            }
            // 留下我自己
            this.className = 'current';
            // 2. 下面的显示内容模块
            var index = this.getAttribute('index');
            // 干掉所有人 让其余的item 这些div 隐藏
            for (var i = 0; i < items.length; i++) {
                items[i].style.display = 'none';
            }
            // 留下我自己 让对应的item 显示出来
            items[index].style.display = 'block';
        }
    }
    //数据传输函数
function showData1(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companytype+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist1").append(str);
    }
}
function showData2(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companytype+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist2").append(str);
    }
}
function showData3(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companytype+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist3").append(str);
    }
}
function showData4(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companytype+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist4").append(str);
    }
}

function showData5(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companyType+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist5").append(str);
    }
}
function showData6(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companyType+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist6").append(str);
    }
}
function showData7(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companyType+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist7").append(str);
    }
}
function showData8(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companyType+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist8").append(str);
    }
}
function showData9(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = "<li>" +
            "<span>"+jsonList[i].title+"</span>"+
            "<h1>"+jsonList[i].salary+"</h1>"+
            "<span>"+jsonList[i].education+"</span>"+
            "<span>"+jsonList[i].experience+"</span>"+
            "<span>"+jsonList[i].shortname+"</span>"+
            "<span>"+jsonList[i].address+"</span>"+
            "<span>"+jsonList[i].companyType+"</span>"+
            "<span>"+jsonList[i].scale+"</span>"+
            "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
            "<button>查看</button>"+"</li>"
        $("#alllist9").append(str);
    }
}

$("#test").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid
        },
        url:'http://127.0.0.1:8080/v4/laceinvite',
        success:function(result){
            var jsonlist = result.rank1
            showData1(jsonlist)
            console.log(result)
        },
    })
})
$("#test1").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid,
        },
        url:'http://127.0.0.1:8080/v4/laceinvite',
        success:function(result){
            var jsonlist = result.rank1
            showData1(jsonlist)
            console.log(result)
        },
    })
});
$("#test2").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid,
        },
        url:'http://127.0.0.1:8080/v4/laceinvite',
        success:function(result){
            var jsonlist = result.rank2
            showData2(jsonlist)
            console.log(result)
        },
    })
});

$("#test3").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid,
        },
        url:'http://127.0.0.1:8080/v4/laceinvite',
        success:function(result){
            var jsonlist = result.rank3
            showData3(jsonlist)
            console.log(result)
        },
    })
});

$("#test4").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "cid":cid,
        },
        url:'http://127.0.0.1:8080/v4/laceinvite',
        success:function(result){
            var jsonlist = result.rank4
            showData4(jsonlist)
            console.log(result)
        },
    })
});


$("#test6").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank5
            showData5(jsonlist)
            console.log(result)
        },
    })
})
$("#test5").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank5
            showData5(jsonlist)
            console.log(result)
        },
    })
});
$("#test7").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank6
            showData6(jsonlist)
            console.log(result)
        },
    })
});

$("#test8").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank7
            showData7(jsonlist)
            console.log(result)
        },
    })
});

$("#test9").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank8
            showData8(jsonlist)
            console.log(result)
        },
    })
});

$("#test10").one("click",function(){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v7/application',
        success:function(result){
            var jsonlist = result.rank9
            showData9(jsonlist)
            console.log(result)
        },
    })
});

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
var pid = getQueryVariable("pid")
console.log(pid)
$("head_body").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "pid":pid,
        },
        url: 'http://127.0.0.1:8080/v7/application',
    })
})
