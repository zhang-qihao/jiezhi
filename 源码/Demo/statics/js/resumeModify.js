function showEducation(jsonlist){
    var str =''
    for (var i=0;i<jsonlist.length;i++){
        str = '<div className="study1"><p>'+jsonlist[i].school_name+'</p><p>'+jsonlist[i].highest_education+
            '|<span'+jsonlist[i].major_studied+'</span></p></div>'
        $("#study").append(str);
    }
}

$("#modifyStudy").ready(function (){
    $.ajax({
        type:'post',
        dataType:'json',
        url:'http://127.0.0.1:8080/v2/resumeModify.html',
        success:function(result){
            var jsonlist = result.education
            showEducation(jsonlist)
            console.log(result)
        },
    })
})


