from django.db import models

# Create your models here.
#Tableの設計をここに記述する。
#詳細については、以下のリンク参照
# https://qiita.com/kotayanagi/items/3cfadae951c407ac044a
#Primary keyは、自動で生成される。
class Question(models.Model):
    #models.CharFieldは、必ずmax_length=を指定する。
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField("date published")
    def __str__(self):
        return self.question_text

class Choice(models.Model):
    #1対他の関係に必要なリレーションキー on_delete=models.CASCADEは、親が消えると子も消える
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)
    def __str__(self):
        return self.choice_text
