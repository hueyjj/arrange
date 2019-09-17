from django.conf.urls import include, url
from rest_framework.routers import DefaultRouter

from arrange.guild import views

router = DefaultRouter()
router.register(r'users', views.UserViewSet)
router.register(r'guilds', views.GuildViewSet)

urlpatterns = [
    url('', include(router.urls))
]