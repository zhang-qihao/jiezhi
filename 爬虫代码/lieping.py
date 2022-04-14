import re
import urllib.request
import urllib.parse
from bs4 import BeautifulSoup
import xlwt
import random
key = "cs"

# 查找公司信息
# 查找公司链接
findcompany_link = re.compile(r'<a href="(.*?)"')
# 查找公司简称
findsimple_name = re.compile(r'data-selector="company-name">(.*)</span>')
# 查找公司全称
findname = re.compile(r'<p class="text">(.*)</p>')
# 查找公司LOGO
findlogo = re.compile(r'class="logo.*src="(.*)"/>')
# 查找公司近照
findphoto = re.compile(r'<img src="(.*)"/>')
# 查找公司行业
findcompany_type = re.compile(r'<p>(.*)')
# 查找公司规模
findcompany_people = re.compile(r'<span>·</span>(.*)')
# 查找公司类型
findcompany_who = re.compile(r'<span>·</span>(.*)')
# 查找公司简介
findcompany_text = re.compile(r'lass="inner-text">(.*)')
# 查找公司地区
findcompany_place = re.compile(r'<p class="text">(.*)</p>')


# 查找职业信息
# 查找工作链接
findjob_link = re.compile(r'<a da.* href="(.*?)"')#创建正则表达式对象，表示规则（字符串模式）
# 查找工作薪资
findmoney =re.compile(r'<span class="text-warning">(.*?)<', re.S)
# 查找工作名称
findTitle = re.compile(r'\t\t\t\t\t\t\t\t\t\t\t\t\t(.*) </a>')
# 查找工作地点
findplace = re.compile(r'area".*>(.*)<')
# 查找工作教育要求
findedu = re.compile(r'edu">(.*)<')
# 查找工作时间
findyear = re.compile(r'<span>(.*)</span>')
# 查找工作企业
findcompany = re.compile(r'target="_blank" title=.*>(.*)<')
# 查找工作职责
findzhize = re.compile(r'(.*)(岗位要求|岗位需求|任职需求|任职资格|基本要求|Requirements|任职要求|Qualifications|Required Skills|Essential Responsibilities)')
# 查找工作要求
findyaoqiu = re.compile(r'(岗位要求|岗位需求|任职需求|任职资格|基本要求|Requirements|任职要求|Qualifications|Required Skills|Essential Responsibilities)(.*)',)


def main():
    baseurl = "https://www.liepin.com/zhaopin/?compkind=&dqs=&pubTime=&pageSize=40&salary=&compTag=&sortFlag=&degradeFlag=0&compIds=&subIndustry=&jobKind=&industries=&compscale=&key=" + key + "&siTag=I-7rQ0e90mv8a37po7dV3Q~fA9rXquZc5IkJpXC-Ycixw&d_sfrom=search_fp_bar&d_ckId=4cddcd91ae54287937d4ccf7e3a6f75e&d_curPage=1&d_pageSize=40&d_headId=4cddcd91ae54287937d4ccf7e3a6f75e&curPage="
    getData(baseurl)


# 爬取页面信息
def getData(baseurl):
    datalist = []
    company_list = []
    job_list = []
    num = 1
    for i in range(0,1): # 获取页面信息26页
        url = baseurl + str(i)
        # print(i)
        html = askURL(url)# 保存源码
        print(url)
        # 2.解析数据
        soup = BeautifulSoup(html,"html.parser")
        for item in soup.find_all('div',class_="sojob-item-main clearfix"):   #查找符合要求的字符串，形成链表  <div class="sojob-item-main clearfix">
            # print(item)
            data = []

            item = str(item)
            num = str(num)
            data.append(num)
            num = int(num)
            num += 1

            titles = re.findall(findTitle,item)[0]
            # print(titles)
            data.append(titles)

            company = re.findall(findcompany,item)[0]
            # print(company)
            data.append(company)

            money = re.findall(findmoney,item)[0]
            # print(money)
            data.append(money)

            try:
                edu = re.findall(findedu,item)[0]
                # print(edu)
                data.append(edu)
            except BaseException:
                data.append(" ")

            flog = random.random()
            if(flog<0.4):
                data.append("应届")
            else:
                data.append("实习")

            try:
                year = re.findall(findyear,item)[0]
                # print(year)
                data.append(year)
            except BaseException:
                data.append(" ")

            try:
                place = re.findall(findplace,item)[0]
                # print(place)
                data.append(place)
            except BaseException:
                data.append(" ")

            try:
                company_link = re.findall(findcompany_link,item)[0] #re库用来通过正则表达式查找特定的字符串
                data.append(company_link)
                # print(company_link)
            except BaseException:
                data.append(" ")

            try:
                job_link = re.findall(findjob_link,item)[0] #re库用来通过正则表达式查找特定的字符串
                # print(job_link)
                data.append(job_link)
            except BaseException:
                data.append(" ")

            datalist.append(data)

            company_list = getcompanyData(company_list, company_link)   #爬取公司信息

            job_list = getjobData(job_list, job_link)   #爬取职业具体信息




    # print(datalist)
    saveData(datalist,job_list,company_list)
    return 0


# 获取一个页面的内容
def askURL(url):
    headers = {##假装自己是浏览器
        "user-agent":"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.89 Safari/537.36",
        "Cookies":"__uuid=1626764902671.75; need_bind_tel=false; c_flag=47165b0e8029d7a2db70e6182f9bb6d9; grwng_uid=3c41274a-0b48-4987-b3c5-d2e2737cd83c; __s_bid=efe24eafd71ff9533dfe9d6b4178f0f687d6; _fecdn_=1; new_user=false; vas_userc_exixt=d4242c71dbc888b541a4dd3b4ddde3a1; ADHOC_MEMBERSHIP_CLIENT_ID1.0=25477746-e847-b409-8680-2ca7c958690d; gr_session_id_bad1b2d9162fab1f80dde1897f7a2972=4a83daf0-414b-4e32-bc6d-6c792f7d3b05; gr_cs1_4a83daf0-414b-4e32-bc6d-6c792f7d3b05=UniqueKey%3A423046c584e7f91841344312b1f2ca2b; gr_session_id_97dcf586237881ba=513b4419-6de9-43f1-ac71-8c8ac2cf1eec; __tlog=1627000467770.90%7C00000000%7C00000000%7Cs_00_pz0%7Cs_00_t00; UniqueKey=423046c584e7f91841344312b1f2ca2b; lt_auth=vewObiZXnl3xtXjdjGJd4%2FlJ3N%2BuVTqY8XhZhxpV1Ye0D%2FTj4P%2FiQA%2BEqbEF%2BioIqx0mc%2FszMLb8%0D%0AMuj%2ByXdD40Ab%2BFGnkp2uvPu7z34CSPpnJvSflMXuqtzpQJwnk3k6ykpjn3ki0HU%3D%0D%0A; acw_tc=2760827c16270004995778315e55095172b6c1735c4171327d7e57bffc4a0a; inited_user=d4242c71dbc888b541a4dd3b4ddde3a1; user_roles=0; imClientId=48db9ec3c0caa03ad09bd359c20a8386; imId=48db9ec3c0caa03ac3212a2e215190f1; imClientId_0=48db9ec3c0caa03ad09bd359c20a8386; imId_0=48db9ec3c0caa03ac3212a2e215190f1; fe_im_connectJson_0=%7B%220_423046c584e7f91841344312b1f2ca2b%22%3A%7B%22socketConnect%22%3A%221%22%2C%22connectDomain%22%3A%22liepin.com%22%7D%7D; fe_se=-1627000710982; JSESSIONID=C334556AF08AF8969B5354EFF196A3BC; __session_seq=5; __uv_seq=5; Hm_lvt_a2647413544f5a04f00da7eee0d5e200=1626827236,1626924753,1627000509; Hm_lpvt_a2647413544f5a04f00da7eee0d5e200=1627000741; gr_user_id=50ffcf56-7de0-4b9f-82e0-9de2e61d3720; bad1b2d9162fab1f80dde1897f7a2972_gr_session_id_57772dc1-3980-461d-8647-caea29c53ad2=true; bad1b2d9162fab1f80dde1897f7a2972_gr_session_id=57772dc1-3980-461d-8647-caea29c53ad2; bad1b2d9162fab1f80dde1897f7a2972_gr_last_sent_sid_with_cs1=57772dc1-3980-461d-8647-caea29c53ad2;bad1b2d9162fab1f80dde1897f7a2972_gr_last_sent_cs1=423046c584e7f91841344312b1f2ca2b; bad1b2d9162fab1f80dde1897f7a2972_gr_cs1=423046c584e7f91841344312b1f2ca2b; imApp_0=1; user_photo=5f8fa3a78dbe6273dcf85e2608u.png; user_name=%E5%BD%AD%E6%B5%A9; fe_im_socketSequence_new_0=3_3_1",
                }
    requests = urllib.request.Request(url,headers=headers,method="GET")

    try:
        html = ""
        response = urllib.request.urlopen(requests)
        html = response.read().decode("utf-8")
        # print(html)
    except urllib.error.URLError as e:
        if hasattr(e,"code"):
            print(e.code)
        if hasattr(e,"reason"):
            print(e.reason)
    return html

# 爬取公司页面
def getcompanyData(company_list,url):
    html = askURL(url)
    print(url)
    # 2.解析数据
    data = []
    soup = BeautifulSoup(html, "html.parser")
    for item in soup.find_all('div',class_="main"):  # 查找符合要求的字符串，形成链表  <div class="sojob-item-main clearfix">
        # print(item)

        data = []
        item = str(item)

        try:
            simple_name = re.findall(findsimple_name, item)[0]
            # print(simple_name)
            data.append(simple_name)
        except BaseException:
            data.append(" ")

        try:
            name = re.findall(findname, item)[0]
            # print(name)
            data.append(name)
        except BaseException:
            data.append(" ")

        try:
            logo = re.findall(findlogo, item)[0]
            # print(logo)
            data.append(logo)
        except BaseException:
            data.append(" ")

        try:
            photo = re.findall(findphoto, item)[0]
            # print(photo)
            data.append(photo)
        except BaseException:
            data.append(" ")

        try:
            type = re.findall(findcompany_type, item)[0]
            # print(type)
            data.append(type)
        except BaseException:
            data.append(" ")

        try:
            people = re.findall(findcompany_people, item)[0]
            # print(people)
            data.append(people)
        except BaseException:
            data.append(" ")

        try:
            who = re.findall(findcompany_who, item)
            # print(who[len(who)-1])
            data.append(who[len(who)-1])
        except BaseException:
            data.append(" ")


        try:
            # 公司简介爬取
            for item1 in soup.find_all('div', class_="inner-text"):
                item1 = str(item1)
                text = item1.replace('<div class="inner-text">', '').replace('</div>', '').replace('\n', '')
                # print(text)
            data.append(text)
        except BaseException:
            data.append(" ")


        try:
            place = re.findall(findcompany_place, item)[4]
            # print(place)
            data.append(place)
        except BaseException:
            data.append(" ")

    company_list.append(data)
    # print(data)

    return company_list

# 爬取工作页面
def getjobData(job_list, url):
    data = []
    if (len(url)>25):
        html = askURL(url)  # 保存源码
        print(url)
        # 2.解析数据
        soup = BeautifulSoup(html, "html.parser")
        for item in soup.find_all('div', class_="job-item main-message job-description"):
            # print(item)
            data = []
            item = str(item)

        try:
            zhize = re.findall(findzhize, item)[0]
            zhize = str(zhize)
            zhize = zhize.replace("<br/>","").replace("</div>","").replace("<div>","").replace('<div class="content content-word">','').replace("('","").replace(")","").replace("head.Requirements/","").replace(r"\uf06c","").replace("Skills /","").replace(r".\uf0a7","")
            zhize = zhize.replace("', '任职资格'","").replace("', 'Qualifications'","").replace("', '岗位要求'","").replace("', '任职要求'","").replace("二、","").replace("', '基本要求'","").replace("', 'Requirements'","").replace(r"\xa0","").replace(r"\t","").replace("Job Requirements/","")
            # print(zhize)
            data.append(zhize)
        except IndexError:
            data.append(" ")

        try:
            yaoqiu = re.findall(findyaoqiu, item)[0]
            yaoqiu = str(yaoqiu)
            yaoqiu = yaoqiu.replace("('","").replace("<br/>","").replace(r"\xa0","").replace("')","").replace("', '","").replace(r"\t","").replace("微信分享","").replace(r"\uf06c","").replace(r"\uf0a7","").replace('")','').replace(', "','').replace("'","")
            # print(yaoqiu)
            data.append(yaoqiu)
        except IndexError:
            data.append(" ")
    else:
        data.append("")
        data.append("")

    job_list.append(data)
    return job_list


# 保存网页信息进excel
def saveData(joblist,job_list,company_list):
    print("save....")
    book = xlwt.Workbook(encoding="utf-8", style_compression=0)
    sheet = book.add_sheet('job_info', cell_overwrite_ok=False)
    col = ("序号","职位","公司","薪资","是否应届","学历要求", "工作经验","工作地址","公司链接","职业链接","职位描述","职位要求")
    for i in range(0,12):
        sheet.write(0,i,col[i])
    for i in range(0,40):
        print("这是第%d条",i+1)
        data = joblist[i]
        dada = job_list[i]
        for j in range(0,10):
            # print(data[j])
            sheet.write(i+1,j,data[j])
        for j in range(0,2):
            sheet.write(i+1,10+j,dada[j])
    book.save('job.xls')

    print("save...company....")

    sheet = book.add_sheet('com_info')
    col = ("公司简称","公司全称","logo","近照","经营方向","人数规模","公司类型", "公司简介","公司地址")
    for i in range(0,9):
        sheet.write(0,i,col[i])
    for i in range(0,40):
        print("这是第%d条",i+1)
        # print(company_list[i])
        data = company_list[i]
        try:
            for j in range(0,9):
                sheet.write(i+1,j,data[j])
        except BaseException:
            sheet.write(i + 1, j,"")
    book.save('job.xls')


if __name__  == "__main__":
    # 调用函数
    main()
    print("爬取完毕！")
