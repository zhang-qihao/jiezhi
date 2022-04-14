
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

// $('.head_searchInput2').click(function (){
//     $('.content').css("display","block");
// })
var pid = getQueryVariable("pid")
var index
$("#head_select").change(function (){
    index = $("#head_select option:selected").val();
    console.log(index);
})
var oldmessage = getQueryVariable("q")
var oldindex = getQueryVariable("i")
console.log(oldmessage)
if (oldmessage!=""){
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "pid":pid,
            "index":oldindex,
            "message":oldmessage,
        },
        url:"/v4/search",
        success:function(result){
            var jsonlist = result.joblist
            showData(jsonlist)
            console.log(jsonlist)
            oldmessage=""
        },
    })
}


var head = document.querySelector(".head_searchInput2")
head.addEventListener('click', function () {
    $('.content').css("display","block");
    var message = document.getElementById("head_searchInput1").value
    console.log(message)
    $.ajax({
        type:'post',
        dataType:'json',
        data:{
            "pid":pid,
            "index":index,
            "message":message,
        },
        url:"/v4/search",
        success:function(result){
            var jsonlist = result.joblist
            showData(jsonlist)
            console.log(result)
            var new_url = "/v4/search?pid="+pid+"&q="+message+"&i="+index
            window.location.href=new_url
        },
    })
});

function showData(jsonlist){
    var str =""
    var obj = document.getElementById("jobList");
    obj.innerHTML="";
    for (var i=0;i<jsonlist.length;i++){
        str ="<li>\n" +
            "<p class=\"positionName\">"+jsonlist[i].title+"</p>\n" +
            "<p class=\"wages\">"+jsonlist[i].salary+"</p>\n" +
            "<span class=\"place\">"+jsonlist[i].address+"</span>\n" +
            "<span class=\"time\">"+jsonlist[i].experience+"</span>\n" +
            "<span class=\"education\">"+jsonlist[i].education+"</span>\n" +
            "<img class=\"pic\" src=\"../statics/images/java职位下级页面/u274.svg\">\n" +
            "<p class=\"company\">"+jsonlist[i].company+"</p>\n" +
            "<span class=\"type\">"+jsonlist[i].nature+"</span>\n" +
            "<span class=\"separate\">|</span>\n" +
            "<span class=\"scale\">"+jsonlist[i].salary+"</span>\n" +
            "<button class=\"write\" onclick='ViewJobPerson(\""+""+jsonlist[i].jid+""+"\")'>查看</button>"+"</li>"
            "</li>"
        $("#jobList").append(str);
    }
}

function ViewJobPerson(jid) {
    console.log("pid=", pid)
    console.log("jid=", jid)
    var url = "/v2/jobinfo?pid=" + pid + "&jid=" + jid
    window.location.href = url
}
// "<div class=\"write\"><span>查看</span></div>\n" +
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

$(function (){
    $('.checkboxli').click(function (){
        $(this).addClass("choose");
        $('.checkboxli2').removeClass("choose");
        $('.column_side1').css("display","block");
        $('.column_side3').css("display","none");
    })
    $('.checkboxli2').click(function (){
        $(this).addClass("choose");
        $('.checkboxli').removeClass("choose");
        $('.column_side3').css("display","block");
        $('.column_side1').css("display","none");
    })
})
