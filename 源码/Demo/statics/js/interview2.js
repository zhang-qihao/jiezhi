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
var pid = getQueryVariable("pid")
console.log(pid)


function JobMergeSting(struct){
    var str = "<li>" +
        "<span>"+struct.title+"</span>"+
        "<h1>"+struct.salary+"</h1>"+
        "<span>"+struct.experience+"</span>"+
        "<span>"+struct.education+"</span>"+
        "<span>"+struct.address+"</span>"+
        "<img src=\"../statics/images/面试邀请/u2118.svg\">"+
        "<button onclick='ViewJobPerson(\""+""+struct.Jid+""+"\")'>查看</button>"+"</li>"
    return str
}
function ViewJobPerson(jid) {
    console.log("pid=", pid)
    console.log("jid=", jid)
    var url = "/v2/jobinfo?pid=" + pid + "&jid=" + jid
    window.location.href = url
}
function showData(jsonList){
    var str = ''
    var jsonList0 = jsonList.rank0
    for (var i = 0;i<jsonList0.length;i++){
        str = JobMergeSting(jsonList0[i])
        $("#alllist5").append(str);
    }
    var jsonList1 = jsonList.rank1
    for (var i = 0;i<jsonList1.length;i++){
        str = JobMergeSting(jsonList1[i])
        $("#alllist6").append(str);
    }
    var jsonList2 = jsonList.rank2
    for (var i = 0;i<jsonList2.length;i++){
        str = JobMergeSting(jsonList2[i])
        $("#alllist7").append(str);
    }
    var jsonList3 = jsonList.rank3
    for (var i = 0;i<jsonList3.length;i++){
        str = JobMergeSting(jsonList3[i])
        $("#alllist8").append(str);
    }
    var jsonList4 = jsonList.rank4
    for (var i = 0;i<jsonList4.length;i++){
        str = JobMergeSting(jsonList4[i])
        $("#alllist9").append(str);
    }
}
$("#head_body").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "pid":pid,
        },
        url:"/v2/application",
        success:function(result){
            showData(result)
            console.log(result)
        },
    })
})

