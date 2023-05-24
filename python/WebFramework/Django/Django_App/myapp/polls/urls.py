from django.urls import path
from . import views
#各Appのルーティングは、各アプリケーションごとに記述する。
# from django.urls import path
# from . import views
# urlpatterns = [
#    path("URLパターン", views.関数名, kwargs=渡す引数, name=グローバルに参照できる),
#    最小の記述は、path("URLのパターン", views.関数名),
# ]
urlpatterns = [
    path("", views.index, name="index"),
]