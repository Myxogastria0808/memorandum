from django.http import HttpResponse

#urls.pyに渡す関数
#urls.pyの時に、views.関数名で渡す
#多くの場合のテンプレートパターン
# (from django.http import HttpResponse)
# def index(request):
#   return HttpResponse("Hello")
#流れとしては、views.py -> urls.py -> routeのurls.py

def index(request):
    return HttpResponse("Hello")