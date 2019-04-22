# -*- coding:UTF8 -*-

import re
headers = {'Host': 'riji.bozhong.com',
"Connection": "keep-alive",
"Cache-Control": "max-age=0",
"User-Agent": "Baiduspider+(+http://www.baidu.com/search/spider.htm)",
"Upgrade-Insecure-Requests": "1",
"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
"Accept-Language":"zh,zh-CN;q=0.9,en-US;q=0.8,en;q=0.7"}

from urllib.parse import quote
import urllib.request, urllib.parse, urllib.error, urllib.request, urllib.error, urllib.parse, http.cookiejar
import os
import string

#from urlparse import urljoin



def sendwithhead(url,data,headers):
    ret = ""
    try:
        if headers is None:
            request = urllib.request.Request(url, data)
        else:
            request = urllib.request.Request(url, data, headers)
        ret = urllib.request.urlopen(request).read().decode('utf-8')
        #print(ret)
    except Exception as e:
        print(e)
    return ret

def write(mydir,filename,content):
    if not os.path.exists(mydir):
        os.makedirs(mydir)
    mypath = (mydir+filename)
    #print(mydir + filename,mypath )
    output = open(mypath, 'w',encoding="utf-8")
    output.write(content)
    output.close()

def saveurl(url,mydir,filename,headers):
#    mydir = mydir.replace("\\\\","\\")
    if not os.path.exists( (mydir+"/"+filename) )  \
            or os.path.getsize( (mydir+"/"+filename) )==0:
        url = quote(url, safe=string.printable)
        content = sendwithhead(url, None, headers)
        write(mydir,filename,content)

def saveimg(url,localpath):
    try:
        if not os.path.exists(localpath) \
                or os.path.getsize(localpath) == 0:
            dirname = os.path.dirname(localpath)
            if not os.path.exists(dirname):
                os.makedirs(dirname)
            url = quote(url, safe=string.printable)
            urllib.request.urlretrieve(url,localpath)
    except Exception as e:
        print(e)

def read(filepath):
    file_object = open(filepath,encoding="utf8")
    try:
        all_the_text = file_object.read()
    finally:
        file_object.close()
    return all_the_text

def getUrl(html):
    patternjs = '<script.*?src="(.*?)"'
    patternimg = '<img.*?src="(.*?)"'
    patterncss = '<link.*?href="(.*?)"'
    patternimg2 = '<div.*?data-src="(.*?)"'
    href = re.compile(patternjs, re.S | re.I ).findall(html)
    href += re.compile(patternimg, re.S | re.I).findall(html)
    href += re.compile(patterncss, re.S | re.I).findall(html)
    href += re.compile(patternimg2, re.S | re.I).findall(html)
    return href

rootpath = "/Users/tom-mac/WorkSpace/tools/html_download/sumec"
url = "https://www.sumec.com/"
filename = url.split("/")[-1]
#saveurl(url, rootpath , filename,headers)
#html = read(rootpath+filename)
#urls = getUrl(html)

for item in open("foo.txt"):
    item = item.strip()
    myurl = urllib.parse.urljoin(url,item)
    print(myurl)
    if "http" in myurl:
        o = urllib.parse.urlparse(item)
        filename = myurl.split("/")[-1]
        filedir = o.path.replace(filename,"")
        if item.endswith(".jpg") or item.endswith(".png") or item.endswith(".gif") :
            print("img")
            saveimg(myurl, rootpath + o.path)
        else:
            print("notimg")
            saveurl(myurl, rootpath + filedir, filename, None)
