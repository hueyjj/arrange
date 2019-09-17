from django.contrib.auth.models import User, Group
from rest_framework import serializers

from arrange.guild.models import Guild


class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ['url', 'username', 'email', 'groups']


class GroupSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Group
        fields = ['url', 'name']

class GuildSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Guild
        fields = ['created']