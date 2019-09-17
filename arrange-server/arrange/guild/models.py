from django.db import models

class Guild(models.Model):
    created = models.DateTimeField(auto_now_add=True)
    description = models.TextField(default="")