//表头效果
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
//鼠标悬停出现框
// $(".menu_item").eq(0).mouseover(function (){
//     $(".left_alert").show();
// })
// $(".menu_item").eq(0).mouseleave(function (){
//     $(".left_alert").hide();
// })
// $(".left_alert").mouseover(function (){
//     $(".left_alert").show();
// })
// $(".left_alert").mouseleave(function (){
//     $(".left_alert").hide();
// })
//轮播图代码
//定义全局变量和定时器
var i = 0 ;
var timer;
var count;

$(function(){
    //获取图片的个数
    count = $(".nav_middle img").length;
    //用jquery方法设置第一张图片显示，其余隐藏
    $('.ig').eq(0).show().siblings('.ig').hide();
    //调用showTime()函数（轮播函数）
    showTime();
    //当鼠标经过下面的数字时，触发两个事件（鼠标悬停和鼠标离开）
    $('.tab').hover(function(){
        //获取当前i的值，并显示，同时还要清除定时器
        i = $(this).index();
        Show();
        clearInterval(timer);
    },function(){
        showTime();
    });
    //鼠标点击左侧的箭头
    $('.btn1').click(function(){
        clearInterval(timer);
        if(i == 0){
            i = count;//注意此时i的值
        }
        i--;
        Show();
        showTime();
    });
    //鼠标点击右侧的箭头
    $('.btn2').click(function(){
        clearInterval(timer);
        //console.log(count-1);
        if(i == count-1){
            i = -1;//注意此时i的值
        }
        i++;
        Show();
        showTime();
    });
});
//创建一个showTime函数
function showTime(){
    //定时器
    timer = setInterval(function(){
        //调用一个Show()函数
        Show();
        i++;
        //当图片是最后一张的后面时，设置图片为第一张
        if(i==count){
            i=0;
        }
    },3000);
}
//创建一个Show函数
function Show(){
    //在这里可以用其他jquery的动画
    $('.ig').eq(i).fadeIn(300).siblings('.ig').fadeOut(300);
    //给.tab创建一个新的Class为其添加一个新的样式，并且要在css代码中设置该样式
    $('.tab').eq(i).addClass('bg').siblings('.tab').removeClass('bg');

}
//按钮点击变色
$(function (){
    $('.nav_right a').click(function (){
        $(this).addClass("changebox");
        $('#username').removeClass("changebox")
    })
})
//城市站入口选择框选择样式
$(function (){
    $(".checkbox ul a").click(function (){
        $(this).children('li').addClass('changebox');
        $(this).siblings().children('li').removeClass('changebox');
        var index = $(this).index();
        console.log(index);
        $(".bodycontent .content").eq(index).show().siblings().hide();
    })
})



//页面弹出框
var exploit = ["Java开发工程师","PHP开发工程师","C/C++开发工程师", "Python开发工程师",".NET开发工程师","C#开发工程师", "Ruby开发工程师","Go开发工程师","大数据开发工程师"
    ,"Hadoop工程师","爬虫开发工程师","脚本开发工程师","多媒体开发工程师","GIS工程师","全栈工程师", "ERP技术开发","区块链开发","高级软件工程师","软件工程师","系统架构设计师",
    "系统分析员", "技术文员/助理","技术文档工程师", "Android开发工程师","iOS开发工程师","小程序开发工程师","移动开发工程师","Web前端开发","HTML5开发工程师","其他"]

var robot = ["自动化测试工程师","硬件测试工程师","项目经理","SLAM定位导航高级算法工程师","云端UE4高级算法工程师","云端UE4高级建模工程师",
    "机械臂控制算法工程师","视觉算法工程师","结构工程师","IQC来料检验工程师","UE4高级3D建模工程师","高级机器人项目策划","私域增长专家","生产信息化系统工程师","BOM工程师",
    "IQC检验员","NPI测试工程师","工厂信息化技术","工厂产业自动化"]
var ai = ["机器学习工程师","深度学习工程师","图像算法工程师","图像处理工程师","图像识别工程师","语音识别工程师","机器视觉工程师","自然语言处理(NLP)","算法工程师"
    ,"推荐算法工程师","搜索算法工程师","其他"]

var game = ["游戏策划师","游戏系统策划","游戏数值策划","游戏关卡策划","游戏文案策划/剧情策划","游戏界面设计师","游戏角色设计师","游戏特效设计师","UE4特效师"
    ,"游戏动作设计师","游戏场景设计师","游戏原画师","游戏动画师","游戏开发工程师","Cocos2d-x开发工程师","Unity3d开发工程师","UE4开发工程师","游戏客户端开发工程师"
    ,"游戏服务端开发工程师","游戏测试",	"游戏运营","电子竞技运营","其他"]

var test = ["软件测试工程师","功能测试","性能测试","安全测试","自动化测试","移动端测试", "测试开发","测试总监","测试经理","测试主管","系统测试","标准化工程师","测试工程师","其他"]

var technology = ["运维工程师","自动化运维工程师","系统工程师", "数据库工程师(DBA)","系统集成工程师","ERP实施顾问", "网络安全工程师","运维开发","网站维护工程师","技术支持/维护经理",
    "技术支持/维护工程师","配置管理工程师","IT经理/IT主管","网络工程师(IT工程师)","网络管理(Helpdesk)","网络维修","手机维修","电脑维修","其他"]

var bigdata = ["数据分析经理/主管","数据分析师","ETL开发工程师", "BI工程师","数据仓库工程师","数据采集工程师","数据建模工程师","数据治理工程师","数据标注师","密码技术应用员",
    "电子数据取证分析师","其他"]

var produce = ["产品总监","产品经理/主管","互联网产品经理","移动产品经理","用户产品经理","电商产品经理","产品专员","产品助理","需求工程师","其他"]

var operation = ["运营总监","运营经理","运营主管","运营专员"	,"运营助理","网站运营总监","网站运营经理/主管","网站运营专员","网络推广总监","网络推广经理/主管",
    "网络推广专员",	"SEO/SEM","信息流优化师","新媒体运营","直播运营","微信运营","微博运营","用户运营","社区/社群运营","活动运营","内容运营","品类运营",
    "数据运营","线下运营", "产品运营","网站编辑"	,"内容审核","网站策划","其他"]
var management = ["首席技术执行官CTO","首席信息官CIO","技术总监/经理","项目总监","项目经理","项目主管","项目执行/协调人员","项目助理","其他"]

var str = ""
for (var i = 0;i < exploit.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+exploit[i]+"</a></li>"
    $("#exploit ul").append(str);
}
for (var i = 0;i < robot.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+robot[i]+"</a></li>"
    $("#robot ul").append(str);
}
for (var i = 0;i < ai.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+ai[i]+"</a></li>"
    $("#ai ul").append(str);
}
for (var i = 0;i < game.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+game[i]+"</a></li>"
    $("#game ul").append(str);
}
for (var i = 0;i < test.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+test[i]+"</a></li>"
    $("#test ul").append(str);
}
for (var i = 0;i < technology.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+technology[i]+"</a></li>"
    $("#technology ul").append(str);
}
for (var i = 0;i < bigdata.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+bigdata[i]+"</a></li>"
    $("#bigdata ul").append(str);
}
for (var i = 0;i < produce.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+produce[i]+"</a></li>"
    $("#produce ul").append(str);
}
for (var i = 0;i < operation.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+operation[i]+"</a></li>"
    $("#operation ul").append(str);
}
for (var i = 0;i < management.length;i++){
    str = "<li class=\"left_alert_item\"><a href=\"#\">"+management[i]+"</a></li>"
    $("#management ul").append(str);
}

$(function (){
    $('.menu_item').mouseover(function (){
        var index =$(this).index();
        $('.left_alert').eq(index).show().siblings().hide();
    })
    $('.menu_item').mouseleave(function (){
        $('.left_alert').hide()
    })
    $('.left_alert').mouseover(function (){
        var index =$(this).index();
        $('.left_alert').eq(index).show().siblings().hide();
    })
    $('.left_alert').mouseleave(function (){
        $('.left_alert').hide()
    })
})


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

var user = document.getElementById("user")
var username = document.getElementById("username")
var pid = getQueryVariable("pid")
function mergeData(struct){
    var str = "<li>" +
        "<span>"+struct.title+"</span>"+
        "<span>"+struct.salary+"</span>"+
        "<span>"+struct.address+"</span>"+
        "<span>"+struct.experience+"</span>"+
        "<span>"+struct.education+"</span>"+
        "<span>"+struct.shortname+"</span>"+
        "<span>"+struct.companytype +"||" +struct.scale+"</span>"+
        "<img src=\"../statics/images/面试邀请（空白/u2042.svg\">"+
        "<button class='look' onclick='ViewJobPerson(\""+""+struct.jid+""+"\")'>查看</button>"+"</li>"
    return str
}

function showData1(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = mergeData(jsonList[i])
        $("#content").append(str);
    }
}
function showData2(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = mergeData(jsonList[i])
        $("#content1").append(str);
    }
}
function showData3(jsonList){
    var str = ""
    for (var i = 0;i<jsonList.length;i++){
        str = mergeData(jsonList[i])
        $("#content2").append(str);
    }
}
$("#mainbody").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{"ID": pid,},
        url:'/v4/home',
        success:function(result){
            user.innerHTML = result.name
            username.innerHTML = "Hi," + result.name
            showData1(result.rank1)
            showData2(result.rank3)
            showData3(result.rank2)
        },
    })
})


function jump1(){
    var url = "/v4/home?pid=" + pid
    console.log(url)
    window.location.href=url
}
function jump2(){
    var url = "/v4/city?pid=" + pid
    console.log(url)
    window.location.href=url
}
function jump3(){
    var url = "/v4/stu?pid=" + pid
    window.location.href=url
}
function jump4(){
    var url = "/v4/scr?pid=" + pid
    window.location.href=url
}

function ViewJobPerson(jid) {
    console.log("pid=", pid)
    console.log("jid=", jid)
    var url = "/v2/jobinfo?pid=" + pid + "&jid=" + jid
    window.location.href = url
}
var index
$("#head_select").change(function (){
    index = $("#head_select option:selected").val();
    console.log(index);
})

var head = document.querySelector(".head_searchInput2")
head.addEventListener('click', function () {
    $('.content').css("display","block");
    var message = document.querySelector(".head_searchInput1").value
    console.log(message)
    var new_url = "/v4/search?pid="+pid+"&q="+message+"&i="+index
    window.location.href=new_url
});

$("#nav").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{"ID":pid},
        url:"/v4/home",
        success:function(result){
            var user = document.getElementById("user")
            user.innerHTML = result.name
        },
    })
})

var userId = document.getElementById("userId")
userId.addEventListener("click",function () {
    $.ajax({
        url:'/v2/resumeModify',
        type:'post',
        dataType:"json",
        data:{
            "ID":pid,
        },
        error:function (returnResult) {
            var url = '/v2/resumeModify?pid=' + pid
            window.location.href = url
        }
    })
})


var username = document.getElementById("username");
username.addEventListener("click",function () {
    $.ajax({
        url:'/v2/resumeModify',
        type:'post',
        dataType:"json",
        data:{
            "ID":pid,
        },
        error:function (returnResult) {
            var url = '/v2/resumeModify?pid=' + pid
            window.location.href = url
        }
    })
})

var myResume = document.getElementById("myResume");
myResume.addEventListener("click",function () {
    $.ajax({
        url:'/v2/resumeModify',
        type:'post',
        dataType:"json",
        data:{
            "pid":pid,
        },
        error:function (returnResult) {
            var url = '/v2/resumeModify?pid=' + pid
            window.location.href = url
        }
    })
})

var myPost = document.getElementById("myPost")
myPost.addEventListener("click",function () {
    $.ajax({
        url:'/v2/application',
        type:'post',
        dataType:"json",
        data:{
            "pid":pid,
        },
        success:function (returnResult) {
            var url = '/v2/application?pid=' + pid
            console.log(url)
            window.location.href = url
        }
    })
})

var preview = document.getElementById("preview")
preview.addEventListener("click",function () {
    $.ajax({
        url:'/v2/resumePreview',
        type:'post',
        dataType:"json",
        data:{
            "pid":pid,
        },
        error:function (returnResult) {
            console.log(pid)
            var url = '/v2/resumePreview?pid=' + pid
            window.location.href = url
        }
    })
})
