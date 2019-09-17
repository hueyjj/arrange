from django.db import models

class Guild(models.Model):
    create = models.DateTimeField(auto_now_add=True)